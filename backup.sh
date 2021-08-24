# Flag
while getopts h:d:f:k: flag
do
   case "${flag}" in
      h) host_name=${OPTARG};;
      d) db_name=${OPTARG};;
      f) backupfolder=${OPTARG};;
      k) keep_day=${OPTARG};;
   esac
done

sqlfile=$backupfolder/all-database-$(date +%d-%m-%Y_%H-%M-%S).sql

#create backup folder
mkdir -p $backupfolder

# Create a backup
if sudo pg_dump -U postgres -h $host_name $db_name > $sqlfile ; then
   echo 'Sql dump created'
else
   echo 'pg_dump return non-zero code' 
   exit
fi

echo $sqlfile 

# Delete old backups 
find $backupfolder -mtime +$keep_day -delete