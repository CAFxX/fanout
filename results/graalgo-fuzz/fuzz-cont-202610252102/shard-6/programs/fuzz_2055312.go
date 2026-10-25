package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint8) bool {
    fmt.Println(p0, p0)
    var v5 float64 = 0.0
    switch 1 {
    case 32768:
        var v6 uint32 = (p0 + (^(p0)))
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return (0 == 338)
}

func main() {
    v1 := []uint32{4294967295, 0, 831}
    var v2 uint32 = 216
    var v3 int = 40
    v4 := []int{(v3 & 127), (v3 | (v3 * (-(v3))))}
    if (1 >= 0) {
        for fi7 := 0; fi7 < 20; fi7++ {
            if len(v4) > 0 {
                v4[0] = (v3 * 4294967295)
            }
            if len(v1) > 0 {
                v1[0] = 2
            }
            v3 = (fi7 >> ((65535) & 63))
            _ = fi7
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            if false {
                switch "hello" {
                case "":
                    v3 = 32768
                case "foo bar":
                    fmt.Println(v3, v2, v2)
                case "hello":
                    var v9 float32 = 1e10
                    _ = v9
                default:
                    if len(v1) > 0 {
                        v1[0] = (v2 + (^(v2)))
                    }
                }
            } else {
                v10 := GenS1{F0: 793, F1: (575 != 2), F2: 2}
                if false {
                    v1 = append(v1, v2)
                }
                _ = v10
            }
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 bool = (0.1 > -2.25)
                if (v12 || true) {
                    v12 = false
                    v12 = (!false)
                    if (66 < 15) {
                        var v13 string = "hello"
                        _ = v13
                    }
                }
                fmt.Println(v2, fi11)
                _ = fi11
                _ = v12
            }
            _ = fi8
        }
        var v14 uint = 65535
        _ = v14
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (0.0 != (0.0 + ((3.14159 - 1e10) / ((-2.25 + 1e10) * 1.5)))) {
            for fi15 := 0; fi15 < 7; fi15++ {
                v2 -= v2
                _ = fi15
            }
            v2 -= 1
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = append(v4, v3)
    if false {
        if true {
            v3 -= 2147483647
            v3 = 803
            var v16 string = ("foo bar" + "a")
            _ = v16
        } else {
            v2 = v2
            var v17 float32 = 1.5
            _ = v17
        }
        v2 += v2
        v3 = (v3 >> ((0) & 63))
    } else {
        fmt.Println(v3)
    }
    var v18 int64 = -1
    v2 = (v2 ^ (^(v2)))
    var v19 string = (("" + "foo bar") + (("hello" + ("" + ("foo bar" + ""))) + "foo bar"))
    var v20 uint8 = 0
    v21 := GenS2{F0: 2147483647}
    fmt.Println(v19, v20, v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
