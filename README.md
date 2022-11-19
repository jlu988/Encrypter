This is a sample encoder and decoder service that takes two parameters

  * key - string that contains only unique letters (case-sensitive) and digits
  * message - string to encode/decode (shouldn't be shorter than key)


	* Based on the provided key, two numerical key will be genrated, a original numerical and transposed key
		* Original numerical key which will be use for encoding the message
		* Transpos key whcih will be use for decoding the message

* Example message : 'important message'
* Example key : '6Cat9w'

|           | original | 	sorted |
|-----------|----------|---------|
| Key       | 6Cat9w   | 	Catw69 |
| Numerical | 123456   | 	234615 |


| Before encoding                                                                        | After encoding                                                                          |
|----------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| <pre>2 3 4 6 1 5 <br>-----------<br>i m p o r t<br>a n t   m e<br>s s a g	e	<br></pre> | <pre>1 2 3 4 5 6 <br>----------- <br>r i m p t o <br>m a n t e <br>e s s a g <br></pre> |


* Output

The result for encoding the message 'important message' using key 6Cat9w is 'rimptomante essa g'

    * AWS dynamoDb is required for internal key