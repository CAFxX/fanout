package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint16; F1 int64; F2 int64 }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int32 }

func main() {
    v1 := []uint32{1, 334}
    v2 := []uint32{42, 923, 42}
    var v3 bool = (true && true)
    v4 := []uint32{42, 770, 2147483648, 538}
    v5 := GenS1{F0: ("" + "a"), F1: 2}
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = 1
    }
    switch 32767 {
    case 106:
        v6 := GenS1{F0: ("a" + "foo bar"), F1: 18446744073709551615}
        _ = v6
    case 1:
        var v7 uint16 = 356
        _ = v7
    case 32767:
        v3 = v3
    default:
        var v8 string = ((("foo bar" + "a") + "a") + "\n\t")
        _ = v8
    }
    switch "foo bar" {
    case "\n\t":
        for fi9 := 0; fi9 < 3; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                if v3 {
                    v3 = v3
                    v3 = v3
                    v3 = (!true)
                } else {
                    v11 := GenS1{F0: (("" + "\n\t") + ("" + "\n\t")), F1: 367}
                    var v12 bool = v3
                    _ = v11
                    _ = v12
                }
                v13 := GenS2{F0: 611}
                var v14 bool = ((((v3 || v3) || (v3 || v3)) || (!false)) && (!(true && (true && (!(v3 || false))))))
                _ = fi10
                _ = v13
                _ = v14
            }
            var v15 string = "\n\t"
            v16 := GenS3{F0: 114, F1: v3}
            _ = fi9
            _ = v15
            _ = v16
        }
    case "foo bar":
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int = (fi17 / (((-(fi17))) | 1))
            v3 = v3
            if v3 {
                if (832 == 0) {
                    v18 = v18
                    v19 := GenS4{F0: fi17, F1: 325}
                    _ = v19
                }
            } else {
                var v20 uint16 = 0
                for fi21 := 0; fi21 < 13; fi21++ {
                    v3 = true
                    _ = fi21
                }
                _ = v20
            }
            _ = fi17
            _ = v18
        }
    default:
        var v22 uint8 = 1
        _ = v22
    }
    v23 := GenS5{F0: -9223372036854775808, F1: 852, F2: 255}
    v3 = (1e10 > -2.25)
    switch "" {
    case "hello":
        for fi24 := 0; fi24 < 6; fi24++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3 = v3
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi24
        }
    case "\n\t":
        var v25 float32 = -2.25
        _ = v25
    }
    var v26 int64 = 1
    if false {
        switch -47 {
        case 547:
            var v27 float32 = (((((0.0 * 0.0) * 1.5) * 1e10) + -2.25) + -2.25)
            _ = v27
        case -330:
            if true {
                var v28 int8 = -99
                v29 := GenS6{F0: 65535, F1: (^(v26)), F2: v26}
                _ = v28
                _ = v29
            }
        case 32767:
            v30 := GenS7{F0: 42, F1: "\n\t"}
            _ = v30
        }
    } else {
        switch 132 {
        case 293:
            for fi31 := 0; fi31 < 11; fi31++ {
                if (-2.25 <= (1.5 - 3.14159)) {
                    fmt.Println(v3)
                    var v32 int64 = v26
                    _ = v32
                }
                for fi33 := 0; fi33 < 20; fi33++ {
                    v26 += (v26 >> ((255) & 63))
                    if (fi31 == (fi33 % ((fi33) | 1))) {
                        v26 -= 100
                        var v34 uint8 = 68
                        v35 := GenS7{F0: 588, F1: ("\n\t" + "foo bar")}
                        _ = v34
                        _ = v35
                    }
                    _ = fi33
                }
                _ = fi31
            }
        }
        for fi36 := 0; fi36 < 15; fi36++ {
            if (0 == 18446744073709551615) {
                v3 = (1 < 18446744073709551615)
                v2 = append(v2, 42)
                v3 = (!v3)
            } else {
                for fi37 := 0; fi37 < 20; fi37++ {
                    v38 := GenS4{F0: fi37, F1: 356}
                    _ = fi37
                    _ = v38
                }
                switch 0 {
                case 42:
                    for fi39 := 0; fi39 < 5; fi39++ {
                        var v40 uint32 = 995
                        if ((953 <= -704) || v3) {
                            var v41 bool = (0 <= 1)
                            v42 := GenS3{F0: uint8(v40), F1: (9223372036854775807 < v26)}
                            _ = v41
                            _ = v42
                        } else {
                            var v43 int8 = -128
                            v44 := GenS8{F0: 291}
                            _ = v43
                            _ = v44
                        }
                        _ = fi39
                        _ = v40
                    }
                case 128:
                    for fi45 := 0; fi45 < 11; fi45++ {
                        var v46 float64 = 1e10
                        if false {
                            switch "hello" {
                            case "":
                                v3 = false
                            }
                            _ = len(v2)
                        } else {
                            var v47 uint32 = 776
                            v48 := GenS7{F0: 42, F1: "\n\t"}
                            v26 = 857
                            _ = v47
                            _ = v48
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(fi36)
                        v26 = 9223372036854775807
                        _ = fi45
                        _ = v46
                    }
                case 110:
                    v49 := GenS9{F0: 1}
                    _ = v49
                }
            }
            for fi50 := 0; fi50 < 15; fi50++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = len(v4)
                _ = fi50
            }
            var v51 float64 = 0.0
            _ = fi36
            _ = v51
        }
    }
    if v3 {
        var v52 uint32 = 2147483648
        if len(v4) > 0 {
            v4[0] = (v52 ^ (v52 >> ((386) & 63)))
        }
        _ = v52
    } else {
        var v53 int = 32767
        _ = v53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
