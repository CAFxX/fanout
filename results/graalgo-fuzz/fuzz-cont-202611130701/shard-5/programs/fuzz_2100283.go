package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 uint64; F2 uint32 }
func genFn0(p0 int) int16 {
    switch 4294967295 {
    case 42:
        p0 += 42
    case 1:
        p0 = p0
    case -2147483648:
        p0 += 1
    default:
        p0 = (^(p0))
    }
    p0 = (-(p0))
    v4 := []int{p0, p0}
    _ = p0
    _ = v4
    return -544
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := []uint32{}
    if (-539 > 961) {
        var v5 uint8 = 255
        if (!((1e10 / 1.0) == 0.0)) {
            switch -9223372036854775808 {
            case 32767:
                v5 -= (^(v5))
            case 0:
                delete(v2, 290)
            default:
                var v6 float64 = -2.25
                _ = v6
            }
        } else {
            if false {
                if (("a" + "") > ("\n\t" + "")) {
                    if true {
                        var v7 uint32 = 65535
                        _ = v7
                    }
                    delete(v2, 568)
                    v5 = v5
                }
                switch ("hello" + "") {
                case "":
                    v5 -= (v5 << ((2) & 63))
                default:
                    var v8 int16 = -1
                    _ = v8
                }
                v5 = v5
            } else {
                switch 0 {
                case -9223372036854775808:
                    var v9 uint16 = 846
                    _ = v9
                case 4294967295:
                    v10 := GenS1{F0: 9223372036854775807, F1: 32767, F2: ("" + (("a" + "\n\t") + ("" + "")))}
                    _ = v10
                }
                v5 += v5
            }
        }
        v11 := GenS2{F0: 1, F1: 42}
        _ = v11
        _ = v5
    } else {
        var v12 float64 = (1e10 - 1.5)
        _ = v12
    }
    switch ("hello" + (("a" + ("foo bar" + ("hello" + "hello"))) + "foo bar")) {
    case "":
        var v13 uint8 = 255
        _ = v13
    }
    v14 := GenS3{F0: 9223372036854775807, F1: 0}
    v15 := GenS4{F0: "\n\t"}
    var v16 int = 2147483647
    _ = len(v3)
    switch v16 {
    case -2147483648:
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS4{F0: ""}
            _ = fi17
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS5{F0: 4294967295}
            v21 := GenS6{F0: 127, F1: 0, F2: 2147483647}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v22 := GenS3{F0: (v16 ^ (v16 | v16)), F1: -74}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
}
