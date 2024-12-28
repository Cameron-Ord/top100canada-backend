DELIMITER ;;
CREATE DEFINER=`cameron`@`localhost` PROCEDURE `get_default_list`()
BEGIN
    SELECT id, name, company FROM top100list order by id desc;
END ;;
DELIMITER ;

