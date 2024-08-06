# rolling-code-generator

With all of the AI deep fakes in the news, this project was initiated after hearing from a well-known security company that they used rolling codes internally (only accessible via LAN) to verify everyone in a Teams/Zoom/Phone meeting was really who they said they were. They would deploy the service only on an internally available and secure LAN so that one could only access the rolling code service to verify with others if all were already on the same internal network. This project aims to provide similar functionality for any organization that would like to experiment or use this in production.

# use cases

- if you have suspicion of any meeting participants, you can ask them to relay over Teams/Zoom/Phone the rolling-code
- at the start of a meeting, all present could respond to the others with the current rolling code to verify themselves

# screenshot

![image](https://github.com/user-attachments/assets/f409c898-d20c-46fb-bd18-2830b4988f00)

# demo

https://rolling-code-generator-app-kd8ka.ondigitalocean.app/

# deploy via Docker

```docker run -p 80:1423 ghcr.io/gsantella/rolling-code-generator:main```
