       IDENTIFICATION DIVISION.
       PROGRAM-ID. TEST.

       ENVIRONMENT DIVISION.
       INPUT-OUTPUT SECTION.
       FILE-CONTROL.
       SELECT INPUT-FILE ASSIGN TO "../inputs/real.txt"
             ORGANIZATION LINE SEQUENTIAL
             FILE STATUS INPUT-FILE-STATUS.
       
       DATA DIVISION.
       FILE SECTION.
       FD INPUT-FILE.
       01 INPUT-RECORD PIC X(10).

       WORKING-STORAGE SECTION.
       01 INPUT-FILE-STATUS PIC 99.
          88 FILE-IS-OK     VALUE 0.
          88 END-OF-FILE    VALUE 10.

       01 LINE-COUNT        PIC 9(6).
       01 TOTAL             PIC 9(12).
       01 TEMP              PIC 9(10).

       PROCEDURE DIVISION.
           PERFORM PART-A THRU PART-A-FN.
           PERFORM PART-B THRU PART-B-FN.
           STOP RUN.

      *============================================================*

       PART-A.
      *-------*
           OPEN INPUT INPUT-FILE.
           IF NOT FILE-IS-OK
              DISPLAY "File could not be opened"
              EXIT
           END-IF.
              
           PERFORM VARYING LINE-COUNT FROM 1 BY 1 UNTIL END-OF-FILE
              READ INPUT-FILE
              IF NOT END-OF-FILE
                 MOVE FUNCTION TRIM(INPUT-RECORD) TO TEMP
                 DIVIDE TEMP BY 3 GIVING TEMP
                 SUBTRACT 2 FROM TEMP GIVING TEMP
                 ADD TEMP TO TOTAL GIVING TOTAL
              END-IF
           END-PERFORM.

           CLOSE INPUT-FILE.

           DISPLAY "Part A: " TOTAL.

       PART-A-FN.
      *----------*
          EXIT.

       PART-B.
      *-------*
          OPEN INPUT INPUT-FILE.
          IF NOT FILE-IS-OK
             DISPLAY "File could not be opened"
             EXIT
          END-IF.

          MOVE ZEROES TO TOTAL
          PERFORM VARYING LINE-COUNT FROM 1 BY 1 UNTIL END-OF-FILE
             READ INPUT-FILE
             IF NOT END-OF-FILE
                MOVE FUNCTION TRIM(INPUT-RECORD) TO TEMP
                PERFORM UNTIL TEMP <= 8
                   DIVIDE TEMP BY 3 GIVING TEMP
                   SUBTRACT 2 FROM TEMP GIVING TEMP
                   ADD TEMP TO TOTAL GIVING TOTAL
                END-PERFORM
             END-IF
          END-PERFORM

          CLOSE INPUT-FILE.

          DISPLAY "Part B: " TOTAL.

       PART-B-FN.
      *----------*
          EXIT.
