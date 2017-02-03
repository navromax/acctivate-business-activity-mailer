USE MaxFormEmail;
GO

DROP PROCEDURE [dbo].[spAddBusinessActivity];
CREATE PROCEDURE [dbo].[spAddBusinessActivity]
    @ActivityID varchar(15) = '',
    @Type varchar(3) = '',
    @Status varchar(30) = '',
    @Code varchar(3) = '',
    @Priority varchar(3) = '',
    @AssignedTo varchar(3) = ' ',
    @Customer varchar(41) = '' ,
    @Contact varchar(209) = '',
    @Phone varchar(21) ='',
    @Fax varchar(21) ='',
    @Email varchar(99) ='',
    @AddressName varchar(209) = '',
    @Address varchar(41) = '',
    @Address2 varchar(41) = '',
    @City varchar(31) = '',
    @State varchar(21)='',
    @Zip varchar(13)='',
    @Country varchar(31) = '',
    @DateOpened datetime = Null,
    @OpenedBy varchar(3) = '',
    @DueDate datetime = Null,
    @Description varchar(255) = '',
    @Discussion varchar(8000) = '',
    @Resolution text = '',
    @Reference varchar(20) = '',
    @Reference2 varchar(20) = ''
AS
  BEGIN
    SELECT * FROM tbIssue;
  END
GO

DROP TABLE  tbIssue;
CREATE TABLE tbIssue(IssueID char(16));
INSERT INTO tbIssue values('0123456789012345')


execute spAddBusinessActivity @Reference2 = '';