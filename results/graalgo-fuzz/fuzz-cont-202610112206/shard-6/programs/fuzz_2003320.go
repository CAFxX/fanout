package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 int16 = 1
    var v2 uint = 18446744073709551615
    var v3 uint = 4294967295
    var v4 int = -110
    var v5 uint16 = 65535
    v6 := []string{(("a" + "\n\t") + ("foo bar" + "foo bar")), "", ("foo bar" + "")}
    v7 := GenS1{F0: (v1 >> (((v2 % (((v2 ^ 1)) | 1))) & 63)), F1: (v1 + (v1 >> (((v3 | (v3 << (((v2 * (v2 & (v3 * (v3 & 323))))) & 63)))) & 63))), F2: -9223372036854775808}
    if false {
        if ((!((!((v2 >> ((uint(v5)) & 63)) > 4294967295)) || false)) || true) {
            v8 := GenS2{F0: 63, F1: 617, F2: uint(v5)}
            _ = v8
        } else {
            fmt.Println(v5)
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 string = ("foo bar" + "foo bar")
                var v11 int = 0
                if (100 <= 1) {
                    var v12 uint32 = 1
                    switch (-(v11)) {
                    case 42:
                        v5 += (v5 & (v5 << ((v3) & 63)))
                    case 0:
                        v5 = (^(v5))
                    case 127:
                        v13 := GenS1{F0: v1, F1: v1, F2: 9223372036854775807}
                        _ = v13
                    }
                    var v14 int32 = 1
                    _ = v12
                    _ = v14
                }
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        switch v4 {
        case -9223372036854775808:
            v2 -= (^(v2))
        case 991:
            var v15 bool = false
            _ = v15
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v4 += -24
    for fi16 := 0; fi16 < 14; fi16++ {
        switch v4 {
        case 1:
            var v17 uint8 = 1
            _ = v17
        case 4294967295:
            v18 := GenS2{F0: 0, F1: -2147483648, F2: (v3 % ((255) | 1))}
            _ = v18
        }
        if false {
            var v19 float32 = -2.25
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v19
        } else {
            if false {
                for fi20 := 0; fi20 < 14; fi20++ {
                    v2 = (v2 % (((v3 % (((v3 | (^(v2)))) | 1))) | 1))
                    if len(v6) > 0 {
                        v6[0] = ("" + "\n\t")
                    }
                    var v21 float32 = 0.1
                    if len(v6) > 0 {
                        v6[0] = ("hello" + "a")
                    }
                    var v22 int16 = v1
                    var v23 float32 = v21
                    var v24 int = 250
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            } else {
                switch (("foo bar" + "\n\t") + ("foo bar" + "\n\t")) {
                case "\n\t":
                    fmt.Println(v1)
                case "hello":
                    var v25 int8 = -1
                    _ = v25
                case "foo bar":
                    if (("hello" >= ("foo bar" + "a")) || false) {
                        var v26 uint32 = 795
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 3; fi27++ {
                            v28 := GenS3{F0: -1, F1: v4}
                            _ = fi27
                            _ = v28
                        }
                        v29 := GenS1{F0: v1, F1: v1, F2: 1}
                        var v30 string = "\n\t"
                        _ = v29
                        _ = v30
                    }
                default:
                    var v31 uint8 = 2
                    _ = v31
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v33 := GenS4{F0: 2, F1: 0}
                    v5 += (v5 << (((v2 / ((v3) | 1))) & 63))
                    var v34 uint16 = 42
                    fmt.Println(v34)
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                switch "a" {
                case "":
                    var v35 int = fi16
                    _ = v35
                }
            }
        }
        var v36 uint = v3
        _ = fi16
        _ = v36
    }
    v5 += v5
    for fi37 := 0; fi37 < 9; fi37++ {
        var v38 uint8 = 0
        if false {
            for fi39 := 0; fi39 < 19; fi39++ {
                v40 := GenS5{F0: 497}
                _ = fi39
                _ = v40
            }
            _ = len(v6)
            v2 = (^(v3))
        }
        _ = fi37
        _ = v38
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        v42 := GenS6{F0: 42, F1: 42}
        for fi43 := 0; fi43 < 17; fi43++ {
            v44 := GenS7{F0: -1}
            _ = fi43
            _ = v44
        }
        _ = fi41
        _ = v42
    }
    var v45 float32 = 1e10
    var v46 int16 = 0
    var v47 int64 = 1
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if len(v6) > 0 {
            v6[0] = "\n\t"
        }
    } else {
        fmt.Println(v45)
        fmt.Println(v4, v5)
    }
    v46 -= (v1 >> (((^(v3))) & 63))
    var v48 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
