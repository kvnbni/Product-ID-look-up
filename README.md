# Product Id look up using Slack

Product information like product id is usually stored in a backend database and it is a pain to look it up for internal use
without production access. Slack is now a widely used collaboration app used to connect with team mates by the means of 
workspaces and further by channels and dms. This project leverages Slack API's to look up product id's in a Slack channel.

Getting Started
1. Build a Slack app with sufficient permission to read and write into channels. 
2. The server code needs to be hosted in a cloud service (I used GCP) so that Slack has a publicly available 
domain to respond to.
3. Replace the product_list.txt file with a json file that has the product information with the same tags.
4. Run the insertRedis.go program to insert the product details into Redis. Note that the product names should be of the format 
<productName.amazon.com>

Prerequisites:
1. Install Go in your public cloud instance.
2. Install Redis client and server.


Running the tests:
I've included a giphy of how the app works.

Bult with:
Go, Redis.

Acknowledgements:
godoc.org and Slack API's documentation.








