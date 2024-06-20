#!/bin/bash
# wait-for-it.sh

# Пример использования:
# ./wait-for-it.sh localhost:5432 -- echo "База данных доступна."

TIMEOUT=30
QUIET=0

echoerr() {
  if [ "$QUIET" -eq 0 ]; then echo "$@" 1>&2; fi
}

usage() {
  exitcode="$1"
  cat << USAGE >&2
Usage:
$cmdname host:port [-t timeout] [-- command args]
-t TIMEOUT | --timeout=TIMEOUT
    Максимальное время ожидания в секундах, по умолчанию 15
-q | --quiet
    Не выводить сообщения
-- COMMAND ARGS
    Команда и ее аргументы для выполнения после того, как порт станет доступным
USAGE
  exit "$exitcode"
}

wait_for() {
  for i in `seq $TIMEOUT` ; do
    nc -z "$HOST" "$PORT" > /dev/null 2>&1

    result=$?
    if [ $result -eq 0 ] ; then
      if [ $# -gt 0 ] ; then
        exec "$@"
      fi
      exit 0
    fi
    sleep 1
  done
  echo "Операция превысила время ожидания: $TIMEOUT секунд"
  exit 1
}

while [ $# -gt 0 ]
do
  case "$1" in
    *:* )
    HOST=`echo $1 | sed -e 's/:.*//'`
    PORT=`echo $1 | sed -e 's/.*://'`
    shift 1
    ;;
    -q | --quiet)
    QUIET=1
    shift 1
    ;;
    -t)
    TIMEOUT="$2"
    if [ "$TIMEOUT" = "" ]; then break; fi
    shift 2
    ;;
    --timeout=*)
    TIMEOUT="${1#*=}"
    shift 1
    ;;
    --)
    shift
    wait_for "$@"
    break
    ;;
    --help)
    usage 0
    ;;
    *)
    echoerr "Неизвестный аргумент: $1"
    usage 1
    ;;
  esac
done

if [ "$HOST" = "" -o "$PORT" = "" ]; then
  echoerr "Ошибка: необходимо указать хост и порт."
  usage 2
fi

wait_for "$@"
