package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 bool }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
func genFn0(p0 uint16) uint16 {
    if (false || (!(true || true))) {
        p0 = p0
        v7 := []int{127}
        if (128 == 140) {
            for fi8 := 0; fi8 < 19; fi8++ {
                fmt.Println(fi8)
                v9 := GenS1{F0: 42, F1: p0, F2: false}
                _ = fi8
                _ = v9
            }
        } else {
            if ((9223372036854775807 >= 1) && (-1 >= 0)) {
                if (-850 > -32768) {
                    v10 := GenS2{F0: 18446744073709551615, F1: 2147483647}
                    _ = v10
                }
            }
        }
        _ = v7
    } else {
        fmt.Println(p0)
        var v11 string = "hello"
        _ = v11
    }
    p0 = (p0 & p0)
    var v12 float64 = float64((0.0 - float32(0.0)))
    _ = p0
    _ = v12
    return (^(p0))
}

func main() {
    v1 := []int{42, 42, 9223372036854775807, -157}
    var v2 uint8 = 127
    v3 := make(map[string]string)
    v4 := []string{}
    v5 := []string{("\n\t" + "hello")}
    var v6 bool = true
    _ = len(v5)
    var v13 uint16 = 265
    if (-9223372036854775808 == -79) {
        for fi14 := 0; fi14 < 4; fi14++ {
            switch ("" + "foo bar") {
            case "foo bar":
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 float64 = 1.5
                    _ = fi15
                    _ = v16
                }
            case "":
                if (0.1 == 1.5) {
                    v2 -= (v2 << ((255) & 63))
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 int64 = 1
                        v2 -= 2
                        _ = len(v1)
                        _ = fi17
                        _ = v18
                    }
                } else {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        fmt.Println(fi14, fi14)
                        if (("foo bar" + "\n\t") > "\n\t") {
                            var v20 int64 = -1
                            _ = v20
                        }
                        _ = fi19
                    }
                    var v21 int64 = -780
                    fmt.Println(v6, v13)
                    _ = v21
                }
            case "hello":
                if v6 {
                    var v22 int32 = 502
                    _ = v22
                }
            }
            fmt.Println(fi14, v2)
            if len(v1) > 0 {
                v1[0] = (fi14 - 782)
            }
            _ = fi14
        }
        _ = len(v3)
        v13 = v13
    } else {
        var v23 bool = (!false)
        v23 = false
        _ = v23
    }
    if (1 == 404) {
        for fi24 := 0; fi24 < 12; fi24++ {
            v25 := GenS1{F0: (fi24 | fi24), F1: (v13 ^ v13), F2: ("a" == "\n\t")}
            _ = fi24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v2 = v2
        _ = len(v3)
        _ = fi26
    }
    v13 -= genFn0(1)
    switch -112 {
    case 42:
        v6 = (1 > 0)
    case 532:
        var v27 uint = 18446744073709551615
        _ = v27
    default:
        var v28 int32 = 1
        _ = v28
    }
    switch 964 {
    case 614:
        v2 = 44
    case 2147483647:
        v3["hello"] = "foo bar"
    case -301:
        v29 := GenS3{F0: 1, F1: 0, F2: 1}
        _ = v29
    }
    v30 := GenS2{F0: 590, F1: 0}
    fmt.Println(v13, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
