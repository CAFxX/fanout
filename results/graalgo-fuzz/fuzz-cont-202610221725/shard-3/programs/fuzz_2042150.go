package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint32; F1 int; F2 uint16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int64, p1 int16) int {
    v6 := GenS1{F0: ("foo bar" + ("a" + "foo bar")), F1: 42, F2: (-(p0))}
    var v7 uint8 = 96
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -193
}
func genFn1(p0 int) bool {
    p0 = (p0 % ((p0) | 1))
    p0 += (p0 ^ p0)
    p0 = -551
    _ = p0
    return false
}

func main() {
    var v1 int8 = 42
    var v2 int32 = -1
    var v3 uint8 = 127
    var v4 string = "a"
    var v5 int = -228
    for fi8 := 0; fi8 < 11; fi8++ {
        if false {
            fmt.Println(v5, v3, v5)
            switch (fi8 + genFn0(9223372036854775807, 686)) {
            case 32768:
                fmt.Println(v4, v2)
            case 909:
                fmt.Println(v3)
            default:
                var v9 uint32 = 0
                _ = v9
            }
        } else {
            switch v4 {
            case "a":
                for fi10 := 0; fi10 < 11; fi10++ {
                    fmt.Println(fi8)
                    fmt.Println(v2, v4, v3)
                    _ = fi10
                }
            case "hello":
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := GenS1{F0: v4, F1: 4294967295, F2: 302}
                    if ((true || ((-43 > 42) || false)) && true) {
                        var v13 int = (fi11 / ((fi8) | 1))
                        switch 4294967295 {
                        case 32767:
                            var v14 int64 = 42
                            _ = v14
                        default:
                            var v15 int32 = v2
                            _ = v15
                        }
                        _ = v13
                    }
                    _ = fi11
                    _ = v12
                }
            }
        }
        v3 = (v3 % ((v3) | 1))
        v1 = (v1 - -1)
        _ = fi8
    }
    v5 = (v5 * (v5 >> ((uint(v3)) & 63)))
    switch "" {
    case "hello":
        v2 += 756
    case "":
        fmt.Println(v3, v3)
    case "a":
        switch (v5 % (((v5 + (v5 >> ((332) & 63)))) | 1)) {
        case 128:
            switch "\n\t" {
            case "\n\t":
                v16 := GenS1{F0: "a", F1: 42, F2: -1}
                _ = v16
            case "foo bar":
                v5 = v5
            }
        case 4294967295:
            v17 := GenS1{F0: "foo bar", F1: 398, F2: 365}
            _ = v17
        case 0:
            switch v4 {
            case "hello":
                switch (v5 << ((811) & 63)) {
                case 813:
                    v2 -= (v2 ^ (v2 << ((0) & 63)))
                case 719:
                    var v18 int16 = 0
                    _ = v18
                case -8:
                    var v19 uint64 = 925
                    _ = v19
                default:
                    var v20 int32 = (v2 | 42)
                    _ = v20
                }
            case "a":
                fmt.Println(v3, v3)
            case "foo bar":
                v21 := GenS2{F0: 0, F1: -281, F2: 0}
                _ = v21
            }
        }
    }
    v22 := GenS3{F0: 100}
    v5 = v5
    fmt.Println(v1)
    v1 = v1
    fmt.Println(v5)
    var v23 int32 = (v2 >> ((255) & 63))
    var v24 int = int(v3)
    v24 = (^(v24))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
