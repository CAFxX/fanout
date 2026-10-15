package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int64; F1 int64; F2 int64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int32, p1 int32, p2 uint16) string {
    var v4 int64 = 9223372036854775807
    for fi5 := 0; fi5 < 8; fi5++ {
        if (42 == 2147483648) {
            if ((p2 >= 32767) || false) {
                p1 -= p0
            }
            var v6 uint = 255
            v7 := GenS1{F0: 32767}
            _ = v6
            _ = v7
        } else {
            if false {
                v8 := []int{fi5, fi5, fi5}
                var v9 uint8 = 42
                v10 := GenS2{F0: 127, F1: 0}
                _ = v10
                _ = v8
                _ = v9
            }
        }
        var v11 float32 = -2.25
        v12 := GenS1{F0: fi5}
        _ = fi5
        _ = v11
        _ = v12
    }
    if false {
        if (18446744073709551615 < uint64(p2)) {
            fmt.Println(v4, p0)
            for fi13 := 0; fi13 < 17; fi13++ {
                fmt.Println(p0, p1, p2)
                switch "\n\t" {
                case "foo bar":
                    v14 := GenS1{F0: 0}
                    _ = v14
                default:
                    v15 := GenS2{F0: -128, F1: 42}
                    _ = v15
                }
                _ = fi13
            }
            v16 := GenS3{F0: v4, F1: (v4 & (v4 / ((9223372036854775807) | 1))), F2: v4}
            _ = v16
        }
    }
    if ((18446744073709551615 != 65535) || (1 <= v4)) {
        if false {
            v17 := []int{}
            p1 = -2147483648
            var v18 string = "foo bar"
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                if (-713 < 83) {
                    switch 3 {
                    case 1:
                        p1 += (p1 ^ p1)
                    default:
                        fmt.Println(fi19, fi19)
                    }
                    if false {
                        var v21 int = (fi19 / ((9223372036854775807) | 1))
                        var v22 uint = 145
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS2{F0: 1, F1: 787}
                    _ = v23
                }
                _ = fi20
            }
            _ = fi19
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi24 := 0; fi24 < 6; fi24++ {
            v25 := GenS4{F0: p2}
            for fi26 := 0; fi26 < 13; fi26++ {
                v4 = v4
                _ = fi26
            }
            _ = fi24
            _ = v25
        }
    case "\n\t":
        v27 := []int{343}
        _ = v27
    case "":
        v28 := GenS4{F0: p2}
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "hello"
}

func main() {
    var v1 int = -1
    var v2 bool = false
    v3 := []string{"\n\t"}
    v2 = true
    v29 := GenS4{F0: 255}
    v2 = v2
    v3 = append(v3, ("foo bar" + "foo bar"))
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 int16 = -1
        v32 := GenS4{F0: 255}
        _ = fi30
        _ = v31
        _ = v32
    }
    v1 = 128
    v2 = (-509 < -2147483648)
    v1 = -567
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
}
