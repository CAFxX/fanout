package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int8; F2 int8 }
type GenS4 struct { F0 int64; F1 uint8 }

func main() {
    var v1 int = 83
    v2 := []int{}
    var v3 int8 = 127
    var v4 uint32 = 65535
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: (fi5 ^ (-(fi5))), F1: 665, F2: 42}
        v4 = (v4 ^ (v4 & 65535))
        _ = fi5
        _ = v6
    }
    _ = len(v2)
    v4 = v4
    v7 := GenS1{F0: (v1 + -1), F1: (v4 | (v4 + v4)), F2: (^(v3))}
    var v8 uint64 = 140
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        v4 = 42
        v9 := GenS1{F0: 32767, F1: (v4 + (^(v4))), F2: (v3 | 127)}
        _ = v9
    } else {
        var v10 int16 = -32768
        if (-2.25 == -2.25) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        if (!(("foo bar" + "\n\t") == ("" + (("hello" + ((("hello" + "foo bar") + "a") + ("\n\t" + ("hello" + ("" + ("" + "a")))))) + "a")))) {
            v11 := GenS1{F0: (v1 << ((2) & 63)), F1: v4, F2: v3}
            v4 += (v4 >> ((732) & 63))
            _ = v11
        }
        _ = v10
    }
    if true {
        v12 := GenS1{F0: 32767, F1: (v4 % ((789) | 1)), F2: v3}
        for fi13 := 0; fi13 < 17; fi13++ {
            v1 = v1
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 float64 = 0.0
                switch fi14 {
                case -1:
                    if false {
                        var v16 float64 = 1.5
                        v17 := GenS1{F0: (-(fi13)), F1: v4, F2: (v3 | 42)}
                        _ = v16
                        _ = v17
                    }
                case 32767:
                    if true {
                        v3 = (-(v3))
                        if true {
                            _ = len(v2)
                            v18 := GenS1{F0: 181, F1: v4, F2: v3}
                            v19 := GenS2{F0: (false || false)}
                            _ = v18
                            _ = v19
                        } else {
                            _ = len(v2)
                            switch (^(fi14)) {
                            case -343:
                                var v20 uint8 = 161
                                _ = v20
                            case 357:
                                switch 32768 {
                                case -9223372036854775808:
                                    _ = len(v2)
                                }
                            }
                        }
                        var v21 uint16 = 255
                        _ = v21
                    } else {
                        switch v1 {
                        case 32767:
                            if ((v4 + (^(v4))) != (^(v4))) {
                                v1 += (^(fi14))
                                var v22 uint = 955
                                _ = v22
                            }
                        default:
                            var v23 float32 = 0.0
                            _ = v23
                        }
                    }
                case -224:
                    if true {
                        var v24 uint16 = 0
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                        var v25 uint64 = (v8 + v8)
                        _ = v24
                        _ = v25
                    }
                }
                _ = fi14
                _ = v15
            }
            v26 := GenS2{F0: false}
            _ = fi13
            _ = v26
        }
        if true {
            var v27 uint16 = 304
            if (("" <= "foo bar") || false) {
                var v28 uint = 157
                v29 := GenS2{F0: true}
                _ = v28
                _ = v29
            }
            var v30 uint8 = 1
            _ = v27
            _ = v30
        }
        _ = v12
    }
    if (-505 > -342) {
        fmt.Println(v3, v8)
        switch (^(v1)) {
        case 0:
            v8 = v8
        case 2147483647:
            for fi31 := 0; fi31 < 1; fi31++ {
                v1 -= v1
                _ = fi31
            }
        case 1:
            var v32 string = ("\n\t" + "foo bar")
            _ = v32
        default:
            v3 = (v3 & -74)
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 int32 = 1
            _ = fi33
            _ = v34
        }
    } else {
        v35 := GenS1{F0: v1, F1: (v4 & 0), F2: -21}
        for fi36 := 0; fi36 < 7; fi36++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi37 := 0; fi37 < 10; fi37++ {
                var v38 float64 = 3.14159
                if (1 < 42) {
                    v2 = append(v2, fi36)
                    v1 = fi36
                    var v39 int = fi36
                    v38 = (v38 * (v38 + 1.5))
                    _ = v39
                }
                switch ("foo bar" + ("\n\t" + "\n\t")) {
                case "foo bar":
                    switch ("hello" + "hello") {
                    case "\n\t":
                        v40 := GenS1{F0: 32768, F1: (v4 % ((v4) | 1)), F2: (v3 + (^(v3)))}
                        _ = v40
                    case "":
                        switch -2147483648 {
                        case 0:
                            v41 := GenS3{F0: -2147483648, F1: int8(fi36), F2: -22}
                            _ = v41
                        case -9223372036854775808:
                            v3 = 4
                        }
                    case "a":
                        switch "foo bar" {
                        case "":
                            v4 = v4
                        case "a":
                            v42 := GenS1{F0: v1, F1: v4, F2: (v3 % ((v3) | 1))}
                            _ = v42
                        case "hello":
                            v3 += v3
                        }
                    }
                case "\n\t":
                    v2 = append(v2, 0)
                case "":
                    v2 = append(v2, -9223372036854775808)
                }
                _ = fi37
                _ = v38
            }
            v43 := GenS4{F0: -9223372036854775808, F1: 42}
            _ = fi36
            _ = v43
        }
        v44 := GenS4{F0: 1, F1: 0}
        _ = v35
        _ = v44
    }
    v8 -= 860
    v45 := GenS3{F0: 42, F1: 100, F2: 127}
    fmt.Println(v1)
    v46 := GenS2{F0: false}
    for fi47 := 0; fi47 < 16; fi47++ {
        v3 = v3
        for fi48 := 0; fi48 < 19; fi48++ {
            var v49 uint8 = 2
            _ = fi48
            _ = v49
        }
        v1 = (^(fi47))
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v7
    _ = v8
}
