def response(hey_bob):
    hey_bob_striped = hey_bob.strip()
    
    if '?' in hey_bob_striped and hey_bob_striped.isupper():
        return "Calm down, I know what I'm doing!"
        
    if '?' in hey_bob_striped and hey_bob_striped[-1] is '?':
        return "Sure."
        
    if hey_bob_striped.isupper():
        return "Whoa, chill out!" 
        
    if not hey_bob_striped:
        return "Fine. Be that way!" 
        
    return "Whatever."
    
