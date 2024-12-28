DELIMITER ;;
CREATE DEFINER=`cameron`@`localhost` PROCEDURE `insert_into_list`(name_str varchar(255), company_str varchar(255))
BEGIN
    insert into top100list (name, company) values (name_str, company_str);
    commit;
END ;;
DELIMITER ;

