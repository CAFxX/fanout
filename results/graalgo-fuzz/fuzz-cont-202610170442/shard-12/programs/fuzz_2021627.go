package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int32 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []string{}
    var v2 bool = (1 == 2147483647)
    var v3 uint16 = 728
    var v4 uint64 = uint64(v3)
    var v5 int64 = int64(v3)
    var v6 int32 = 1
    switch ("" + "hello") {
    case "a":
        v6 = v6
    case "hello":
        _ = len(v1)
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v2 = (v2 && (2147483648 <= 0))
        _ = fi7
    }
    var v8 uint32 = 1
    switch ("foo bar" + "a") {
    case "":
        for fi9 := 0; fi9 < 9; fi9++ {
            if true {
                fmt.Println(v5, v8)
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                v10 := GenS1{F0: (v5 + (v5 >> ((42) & 63))), F1: 2}
                _ = v10
            } else {
                if (!false) {
                    switch fi9 {
                    case 1:
                        v11 := GenS1{F0: (v5 >> ((18446744073709551615) & 63)), F1: v8}
                        _ = v11
                    case 32768:
                        _ = len(v1)
                    default:
                        if (1.5 == (0.0 / 0.1)) {
                            v12 := GenS1{F0: (-(v5)), F1: (v8 ^ (v8 >> ((682) & 63)))}
                            _ = len(v1)
                            v13 := GenS2{F0: 42, F1: 14, F2: (v6 * -2147483648)}
                            _ = v12
                            _ = v13
                        } else {
                            _ = len(v1)
                        }
                    }
                } else {
                    v14 := GenS1{F0: (v5 & (v5 - (-(v5)))), F1: 1}
                    _ = v14
                }
                switch (fi9 + 190) {
                case -2147483648:
                    v15 := GenS3{F0: (v5 & (v5 << ((335) & 63)))}
                    _ = v15
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            v6 = v6
            _ = fi9
        }
    case "a":
        if len(v1) > 0 {
            v1[0] = ("a" + "")
        }
    }
    v2 = v2
    var v16 uint16 = 431
    if ((v4 + 0) >= v4) {
        _ = len(v1)
    }
    v5 -= v5
    if ((!false) && (1e10 >= 3.14159)) {
        var v17 int8 = 0
        fmt.Println(v3, v5, v17)
        _ = v17
    } else {
        var v18 uint32 = 1
        for fi19 := 0; fi19 < 17; fi19++ {
            switch "\n\t" {
            case "foo bar":
                v5 += (v5 % (((^(v5))) | 1))
            case "hello":
                v3 = 2
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                if v2 {
                    v5 = v5
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v6 -= -2147483648
                    v21 := GenS3{F0: (v5 / ((9223372036854775807) | 1))}
                    _ = v21
                }
                if (int(v18) >= (fi19 / ((-2147483648) | 1))) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v22 bool = (!false)
                    _ = v22
                }
                _ = fi20
            }
            _ = fi19
        }
        _ = v18
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        v1 = append(v1, ("foo bar" + ("hello" + "\n\t")))
        var v24 int16 = 767
        var v25 string = "foo bar"
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
