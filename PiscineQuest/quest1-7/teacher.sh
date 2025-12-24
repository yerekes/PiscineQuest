#!/bin/bash
INTERVIEW_FILE=$(find . -type f -name "interview*" | head -n 1)
KEY_INTERVIEW
KEY_INTERVIEW=$(echo "$INTERVIEW_FILE" | grep -o '[0-9]\+')
export KEY_INTERVIEW
echo "$KEY_INTERVIEW"
cat "$INTERVIEW_FILE"
MAIN_SUSPECT
echo "$MAIN_SUSPECT"