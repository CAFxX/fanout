package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 string }
func genFn0() uint64 {
    var v5 float64 = 1.5
    v5 += ((1.5 - v5) - v5)
    var v6 float32 = 1e10
    fmt.Println(v6)
    _ = v5
    _ = v6
    return 765
}

func main() {
    var v1 uint = 261
    v2 := []int{-1, 593, 727, 9223372036854775807}
    var v3 int64 = -902
    v4 := make(map[string]uint32)
    switch 186 {
    case 32767:
        v7 := v4[("hello" + ("hello" + (("" + "foo bar") + "\n\t")))]
        _ = v7
    case -2147483648:
        fmt.Println(v3)
    }
    v3 = -1
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = 0
    }
    v3 -= 9223372036854775807
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v8 string = ""
    var v9 uint32 = 42
    fmt.Println(v8)
    if (58 < 205) {
        v1 -= 131
        var v10 int32 = 2147483647
        if (true || (2 != 65535)) {
            delete(v4, v8)
            if (2 < 2) {
                v11 := GenS1{F0: (v3 >> ((v1) & 63)), F1: (^(v1)), F2: (!(58 < -37))}
                _ = v11
            }
        }
        _ = v10
    } else {
        v8 = (v8 + ("foo bar" + "a"))
        var v12 float64 = float64((0.1 * (0.1 * ((1e10 / 3.14159) * float32(1.5)))))
        _ = v12
    }
    var v13 bool = (238 >= 128)
    switch v8 {
    case "a":
        if (!((v1 ^ (v1 % (((^(v1))) | 1))) >= v1)) {
            switch "\n\t" {
            case "":
                v9 = v9
            case "foo bar":
                var v14 float64 = float64(0.1)
                _ = v14
            case "a":
                v15 := GenS1{F0: v3, F1: (v1 << ((v1) & 63)), F2: false}
                _ = v15
            }
            v13 = v13
            _ = len(v2)
        }
    case "":
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint = v1
            if (true && v13) {
                v18 := GenS1{F0: 267, F1: 698, F2: false}
                for fi19 := 0; fi19 < 10; fi19++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v20 string = v8
                    v8 = ""
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                v21 := GenS2{F0: fi16, F1: 377, F2: 32767}
                if ((fi16 & fi16) >= (^(fi16))) {
                    if len(v2) > 0 {
                        v2[0] = (fi16 - fi16)
                    }
                    var v22 uint32 = v9
                    _ = v22
                }
                v1 += 0
                _ = v21
            }
            _ = fi16
            _ = v17
        }
    case "\n\t":
        v23 := GenS1{F0: 1, F1: v1, F2: true}
        _ = v23
    default:
        if v13 {
            if (0.1 <= 0.0) {
                v24 := GenS1{F0: v3, F1: (v1 / ((v1) | 1)), F2: (-9223372036854775808 >= (v3 / ((v3) | 1)))}
                _ = v24
            } else {
                for fi25 := 0; fi25 < 5; fi25++ {
                    _ = len(v2)
                    _ = fi25
                }
                switch "a" {
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "foo bar":
                    v26 := GenS3{F0: 0}
                    _ = v26
                case "\n\t":
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v9 = 724
                        v13 = ("\n\t" == ((v8 + v8) + ((v8 + v8) + v8)))
                        _ = fi27
                    }
                }
                var v28 int64 = (v3 + v3)
                _ = v28
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            if (!(false && v13)) {
                if (!false) {
                    v9 = (v9 << (((v1 % (((v1 - (v1 / ((v1) | 1)))) | 1))) & 63))
                }
            }
            v29 := GenS4{F0: 65535}
            switch 198 {
            case 128:
                var v30 float32 = (-2.25 - 0.1)
                _ = v30
            default:
                v31 := GenS5{F0: 1, F1: (v8 + v8)}
                _ = v31
            }
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
