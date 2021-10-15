# Enigma

A digital clone of a 3-rotor German Enigma machine.

Usage as a Go package should be easy enough from the code documentation. A binary that encrypt from
standard input and decrypt to standard output is available in `cmd/cli/`. For now, it defaults to a
hardcoded key.

## TODO

### Package

- [ ] Every input character that is not in [A-Z] is a no-op and is output as it is. A warning should be sent. This requires to properly convert the character to int
- [ ] Refactor a bit, and make the package exposed API sound and clear.
- [ ] Log only if asked for
- [ ] Benchmark
- [ ] Log state if needed in the following form:

```text
0001 F > KGWNT(R)BLQPAHYDVJIFXEZOCSMU CDTK 25 15 16 26
0002 O > UORYTQSLWXZHNM(B)VFCGEAPIJDK CDTL 25 15 16 01
0003 L > HLNRSKJAMGF(B)ICUQPDEYOZXWTV CDTM 25 15 16 02
0004 G > KPTXIG(F)MESAUHYQBOVJCLRZDNW CDUN 25 15 17 03
...
```

- [ ] Log a character encryption in the following form:

```text
 G > ABCDEF(G)HIJKLMNOPQRSTUVWXYZ
   P EFMQAB(G)UINKXCJORDPZTHWVLYS         AE.BF.CM.DQ.HU.JN.LX.PR.SZ.VW
   1 OFRJVM(A)ZHQNBXPYKCULGSWETDI  N  03  VIII
   2 (N)UKCHVSMDGTZQFYEWPIALOXRJB  U  17  VI
   3 XJMIYVCARQOWH(L)NDSUFKGBEPZT  D  15  V
   4 QUNGALXEPKZ(Y)RDSOFTVCMBIHWJ  C  25  β
   R RDOBJNTKVEHMLFCWZAXGYIPS(U)Q         c
   4 EVTNHQDXWZJFUCPIAMOR(B)SYGLK         β
   3 H(V)GPWSUMDBTNCOKXJIQZRFLAEY         V
   2 TZDIPNJESYCUHAVRMXGKB(F)QWOL         VI
   1 GLQYW(B)TIZDPSFKANJCUXREVMOH         VIII
   P E(F)MQABGUINKXCJORDPZTHWVLYS         AE.BF.CM.DQ.HU.JN.LX.PR.SZ.VW
 F < KPTXIG(F)MESAUHYQBOVJCLRZDNW
```

### CLI

- [ ] Read the key form a flag or input
- [ ] Possibility to auto encrypt/decrypt message in the form of Enigma original messages.
