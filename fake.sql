CREATE TABLE [dbo].[tbIssue](
  [GUIDIssue] [uniqueidentifier] ROWGUIDCOL  NOT NULL,
  [IssueID] [varchar](40) NULL,
  [GUIDIssueType] [uniqueidentifier] NULL,
  [GUIDStatusCode] [uniqueidentifier] NULL,
  [StatusChangedDate] [datetime] NULL,
  [Completed] [bit] NOT NULL,
  [GUIDIssueCode] [uniqueidentifier] NULL,
  [GUIDResolutionCode] [uniqueidentifier] NULL,
  [AssignedTo] [varchar](3) NULL,
  [AssignedOrder] [int] NULL,
  [DateOpened] [datetime] NULL,
  [OpenedBy] [varchar](3) NULL,
  [DateDue] [datetime] NULL,
  [DatePromised] [datetime] NULL,
  [PromisedBy] [varchar](3) NULL,
  [EstimatedHours] [decimal](19, 7) NULL,
  [WorkAroundDate] [datetime] NULL,
  [WorkAroundBy] [varchar](3) NULL,
  [DateResolved] [datetime] NULL,
  [ResolvedBy] [varchar](3) NULL,
  [DateClosed] [datetime] NULL,
  [ClosedBy] [varchar](3) NULL,
  [Description] [varchar](255) NULL,
  [Discussion] [text] NULL,
  [ResolutionDiscussion] [text] NULL,
  [GUIDPriorityCode] [uniqueidentifier] NULL,
  [EstimatedAmount] [money] NULL,
  [Reference] [varchar](20) NULL,
  [Reference2] [varchar](20) NULL,
  [TaxIncluded] [bit] NOT NULL,
  [_StartDate] [datetime] NULL,
  [_EndDate] [datetime] NULL,
  [_PMDue] [datetime] NULL,
  [_ContractType] [varchar](15) NULL,
  [_Disposition] [varchar](60) NULL,
  CONSTRAINT [PK_tbIssue] PRIMARY KEY NONCLUSTERED
    (
      [GUIDIssue] ASC
    )WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]

GO

ALTER TABLE [dbo].[tbIssue] ADD  DEFAULT ((0)) FOR [Completed]
GO

ALTER TABLE [dbo].[tbIssue] ADD  DEFAULT ((0)) FOR [TaxIncluded]
GO

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
  DECLARE @GUIDIssue uniqueidentifier
  DECLARE @GUIDCustomer uniqueidentifier
  DECLARE @XrefType VARCHAR(1)
  DECLARE @GUIDIssueType UNIQUEIDENTIFIER, @GUIDIssueCode UNIQUEIDENTIFIER
  DECLARE @GUIDStatusCode UNIQUEIDENTIFIER, @GUIDPriorityCode UNIQUEIDENTIFIER
  DECLARE @NewContact VARCHAR(209)

  SET NOCOUNT ON

  IF @ActivityID = ''
    select @ActivityID = NEXT VALUE FOR seqIssue

  IF @DateOpened IS NULL
    SELECT @DateOpened = GetDate()

  SELECT @GUIDIssue = NEWID()

  INSERT INTO tbIssue (GUIDIssue, IssueID, GUIDIssueType, GUIDStatusCode, GUIDIssueCode, GUIDPriorityCode, AssignedTo, DateOpened, OpenedBy, DateDue, Description, Discussion, ResolutionDiscussion, Reference, Reference2)
  VALUES (@GUIDIssue, @ActivityID, @GUIDIssueType, @GUIDStatusCode, @GUIDIssueCode, @GUIDPriorityCode, @AssignedTo,
                      @DateOpened, @OpenedBy, @DueDate, @Description, @Discussion, @Resolution, @Reference, @Reference2)

  SET NOCOUNT OFF

  SELECT * FROM tbIssue WHERE GUIDIssue = @GUIDIssue
GO

CREATE SEQUENCE seqIssue
START WITH 1
INCREMENT BY 1;
GO

execute spAddBusinessActivity @ActivityID='ZZZ';
