# Projekat Polovni Automobili - NTP

## Opis aplikacije
- Postoje 3 uloge u sistemu: Neulogovani korisnik, ulogovano korisnik (oglasivac) i admin.
- Neulogovani korisnici mogu samo da pregledaju oglase. 
- Ulogovani korisnici mogu i da pregledaju, kace i brisu oglase. Mogu da komentarisu i da ocene oglas.
- Admin moze da dodaje, izmeni i obrise marke i modele vozila.
- Svi korisnici mogu da pretrazuju oglase po markama i modelima, po lokaciji, snazi motora i ostalim karakteristikama vozila. Mogu da sortiraju oglase po ceni i datumu postavljanja oglasa.

## Tehnologije

### Postgress
- Baza podataka

### Spring boot
- Koristi se za prvi deo backend aplikacije.
- CRUD operacije
- Komuniciranje sa bazom.

### Golang
- Koristi se za drugi deo backend aplikacije, kao sto su pretrage, filtriranje i sortiranje. 
- Podatke dobavlja sa prve backend aplikacije.

### Pharo
- Korsiti se za razvoj klijentskog dela aplikacije.
- 
