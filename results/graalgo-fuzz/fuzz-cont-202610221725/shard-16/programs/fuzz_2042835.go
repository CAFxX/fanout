package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 bool; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint32 }

func main() {
    var v1 int16 = 1000
    var v2 string = "hello"
    v3 := []string{"", ("hello" + v2), ""}
    var v4 int32 = -164
    fmt.Println(v1, v4, v1)
    v2 = (v2 + (v2 + "a"))
    switch 128 {
    case 9223372036854775807:
        var v5 float64 = 1e10
        _ = v5
    case -1:
        v6 := GenS1{F0: v2, F1: true}
        _ = v6
    case 2147483647:
        if (!(("" + v2) > (v2 + v2))) {
            fmt.Println(v1)
            var v7 int = -987
            for fi8 := 0; fi8 < 20; fi8++ {
                v4 += v4
                for fi9 := 0; fi9 < 4; fi9++ {
                    v2 = v2
                    if (!false) {
                        _ = len(v3)
                        v10 := GenS1{F0: v2, F1: ((2147483647 > v4) || (833 == 0))}
                        _ = v10
                    }
                    if true {
                        if true {
                            _ = len(v3)
                        }
                    }
                    _ = fi9
                }
                var v11 bool = (0.0 >= 1.5)
                _ = fi8
                _ = v11
            }
            _ = v7
        }
    default:
        v2 = (("a" + "foo bar") + "a")
    }
    v1 = v1
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS2{F0: (true || false), F1: 2, F2: (v4 / (((v4 * v4)) | 1))}
        v14 := GenS2{F0: false, F1: 616, F2: 2147483647}
        var v15 uint = 18446744073709551615
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := GenS3{F0: 65535}
    fmt.Println(v2, v1, v1)
    v1 = (v1 | v1)
    for fi17 := 0; fi17 < 3; fi17++ {
        v3 = append(v3, (v2 + v2))
        if (0 <= 688) {
            v2 = (v2 + v2)
        } else {
            for fi18 := 0; fi18 < 9; fi18++ {
                fmt.Println(v1, v2, fi17)
                v19 := GenS4{F0: 990, F1: 65535}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi18
                _ = v19
            }
        }
        _ = fi17
    }
    _ = len(v3)
    for fi20 := 0; fi20 < 12; fi20++ {
        if false {
            for fi21 := 0; fi21 < 10; fi21++ {
                v4 += -661
                _ = fi21
            }
            var v22 uint16 = 1
            _ = v22
        }
        var v23 string = ("" + v2)
        var v24 bool = false
        _ = fi20
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
