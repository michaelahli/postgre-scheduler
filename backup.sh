# Flag
while getopts h:d:f:k: flag
do
   case "${flag}" in
      h) host_name=${OPTARG};;
      d) db_name=${OPTARG};;
      f) backupfolder=${OPTARG};;
      k) keep_day=${OPTARG};;
      u) user_name=${OPTARG};;
      p) password=${OPTARG};;
   esac
done

sqlfile=$backupfolder/database-$(date +%d-%m-%Y).sql

#create backup folder
mkdir -p $backupfolder

# Create a backup
if pg_dump postgresql://$user_name:$password@$host_name:5432/$db_name > $sqlfile ; then
   echo 'Sql dump created'
else
   echo 'pg_dump return non-zero code' 
   exit
fi

echo $sqlfile 

# Delete old backups 
find $backupfolder -mtime +$keep_day -delete