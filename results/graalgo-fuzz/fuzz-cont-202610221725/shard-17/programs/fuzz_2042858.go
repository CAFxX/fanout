package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint8; F1 uint; F2 int8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint = 811
    v2 := []string{("foo bar" + ("foo bar" + ("hello" + ""))), "\n\t"}
    var v3 int = -1
    var v4 uint64 = 4294967295
    var v5 uint = 1
    var v6 int64 = 1
    var v7 uint16 = 983
    if len(v2) > 0 {
        v2[0] = ""
    }
    v6 -= 508
    switch ("\n\t" + "hello") {
    case "hello":
        v8 := GenS1{F0: 4294967295, F1: v3}
        _ = v8
    default:
        v9 := GenS2{F0: 144, F1: 892, F2: -128}
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v5 = v1
        _ = fi10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS3{F0: 1, F1: 203}
        if false {
            _ = len(v2)
        }
        var v13 int8 = -95
        _ = fi11
        _ = v12
        _ = v13
    }
    switch v3 {
    case 2147483647:
        if true {
            fmt.Println(v7, v7, v7)
            v6 -= 1
        } else {
            v14 := GenS2{F0: 246, F1: v1, F2: -51}
            _ = len(v2)
            var v15 int64 = v6
            _ = v14
            _ = v15
        }
    default:
        var v16 bool = (false || (!(v7 <= v7)))
        _ = v16
    }
    if (504 == -560) {
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 bool = true
            v18 = true
            var v19 float64 = ((0.1 * (3.14159 - ((1.5 + 0.1) + float64(-2.25)))) * 3.14159)
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS4{F0: (v5 + (^(v5)))}
            switch (fi20 >> ((v5) & 63)) {
            case -2147483648:
                for fi22 := 0; fi22 < 17; fi22++ {
                    v2 = append(v2, "")
                    _ = fi22
                }
            case 9223372036854775807:
                v4 = v4
            case 32767:
                fmt.Println(v3, v5)
            default:
                if ((1 == (^(fi20))) || false) {
                    v3 = -2147483648
                }
            }
            v7 += 32767
            _ = fi20
            _ = v21
        }
    }
    var v23 int8 = 1
    for fi24 := 0; fi24 < 18; fi24++ {
        fmt.Println(v1, v4, fi24)
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
