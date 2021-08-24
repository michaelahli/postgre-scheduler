# Flag
while getopts d:f:k: flag
do
   case "${flag}" in
      d) db_name=${OPTARG};;
      f) backupfolder=${OPTARG};;
      k) keep_day=${OPTARG};;
   esac
done

sqlfile=$backupfolder/all-database-$(date +%d-%m-%Y_%H-%M-%S).sql

#create backup folder
mkdir -p $backupfolder

# Create a backup
if sudo -u postgres pg_dump $db_name > $sqlfile ; then
   echo 'Sql dump created'
else
   echo $backupfolder $db_name $keep_day
   exit
fi

echo $sqlfile 

# Delete old backups 
find $backupfolder -mtime +$keep_day -delete