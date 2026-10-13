package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint }
type GenS5 struct { F0 int; F1 uint64; F2 int8 }
type GenS6 struct { F0 bool; F1 string; F2 bool }
func genFn0(p0 int8, p1 bool, p2 uint) uint16 {
    var v7 string = "\n\t"
    var v8 bool = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32767
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 0.0
    var v3 int64 = 447
    var v4 uint = 65535
    v5 := []int{-516}
    v6 := []int{0, int(v3), 9223372036854775807}
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            v4 = v4
            _ = fi10
        }
        v3 = (v3 >> (((v4 ^ v4)) & 63))
        _ = fi9
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(v4)
        fmt.Println(v4, v2, v3)
        _ = fi11
    }
    switch "a" {
    case "foo bar":
        if (v4 < v4) {
            if (127 != 1) {
                v4 -= v4
            }
        }
    }
    if true {
        switch 42 {
        case 2147483647:
            fmt.Println(v3, v3)
        default:
            v2 = (float32(1e10) / 1.5)
        }
        if ((0.1 + float64(0.1)) == float64(float32(1e10))) {
            if true {
                v12 := GenS1{F0: 1, F1: -18}
                switch "" {
                case "foo bar":
                    for fi13 := 0; fi13 < 9; fi13++ {
                        if false {
                            v14 := GenS2{F0: (v4 & v4), F1: (("a" + "\n\t") + (((("\n\t" + "hello") + "hello") + "\n\t") + "hello"))}
                            fmt.Println(v4, v3, fi13)
                            delete(v1, "")
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 13; fi15++ {
                            v16 := GenS3{F0: 2}
                            var v17 int8 = -62
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        _ = len(v6)
                        _ = fi13
                    }
                }
                v4 = (v4 & (v4 >> ((v4) & 63)))
                _ = v12
            }
        } else {
            fmt.Println(v4)
            var v18 float32 = v2
            _ = v18
        }
        var v19 float32 = (-2.25 + v2)
        _ = v19
    } else {
        v4 = v4
        if false {
            v4 = v4
        }
        v1["foo bar"] = false
    }
    _ = len(v5)
    if false {
        fmt.Println(v3, v4)
        var v20 uint = 0
        for fi21 := 0; fi21 < 1; fi21++ {
            if (-100 >= 100) {
                v22 := GenS4{F0: 255, F1: 42, F2: v4}
                var v23 bool = false
                _ = v22
                _ = v23
            }
            v24 := GenS5{F0: fi21, F1: 759, F2: 100}
            _ = fi21
            _ = v24
        }
        _ = v20
    } else {
        switch -117 {
        case -1:
            _ = len(v6)
        case 32768:
            v25 := GenS6{F0: (!(v4 >= v4)), F1: "", F2: (!true)}
            _ = v25
        case 913:
            fmt.Println(v4)
        default:
            v2 += 3.14159
        }
        switch 128 {
        case 583:
            for fi26 := 0; fi26 < 17; fi26++ {
                v2 = 1.5
                _ = fi26
            }
        case 748:
            switch "foo bar" {
            case "a":
                var v27 uint16 = 208
                _ = v27
            case "hello":
                v4 = (^(v4))
            case "":
                _ = len(v5)
            default:
                v3 = (v3 & (v3 >> ((v4) & 63)))
            }
        }
    }
    if (v4 <= 791) {
        var v28 float32 = (v2 / ((1.5 + (-2.25 / 0.1)) - v2))
        v4 = (^(v4))
        _ = v28
    }
    v29 := v1[(("a" + ("hello" + "hello")) + ("\n\t" + "\n\t"))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
