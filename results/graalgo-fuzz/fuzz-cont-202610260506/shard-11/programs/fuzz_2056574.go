package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int }
func genFn0(p0 bool, p1 bool) int32 {
    p0 = (p0 && false)
    v5 := GenS1{F0: 100}
    _ = p0
    _ = p1
    _ = v5
    return 2147483647
}
func genFn1() int8 {
    var v6 uint = 4294967295
    v6 = v6
    _ = v6
    return -21
}

func main() {
    var v1 float32 = 0.1
    var v2 uint32 = 2147483648
    var v3 int64 = -1
    var v4 int16 = 1000
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS2{F0: genFn1(), F1: fi7}
        v2 = (v2 >> ((2) & 63))
        _ = fi7
        _ = v8
    }
    fmt.Println(v4, v1, v4)
    fmt.Println(v2, v3)
    fmt.Println(v3)
    fmt.Println(v3, v2)
    switch "foo bar" {
    case "a":
        fmt.Println(v4)
    case "\n\t":
        var v9 uint16 = 2
        _ = v9
    case "hello":
        if (true || ((100 >= 1) || true)) {
            if (!false) {
                for fi10 := 0; fi10 < 13; fi10++ {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v12 := []int{}
                        switch "hello" {
                        case "\n\t":
                            v13 := GenS1{F0: genFn1()}
                            _ = v13
                        case "":
                            v1 -= v1
                        }
                        v14 := GenS1{F0: -71}
                        _ = fi11
                        _ = v12
                        _ = v14
                    }
                    _ = fi10
                }
                v3 -= v3
            }
        }
    default:
        var v15 uint8 = 2
        _ = v15
    }
    fmt.Println(v1, v2, v3)
    v16 := []int{}
    switch "hello" {
    case "a":
        fmt.Println(v3, v4, v4)
    case "":
        fmt.Println(v2, v1)
    case "\n\t":
        switch ("\n\t" + (("hello" + "") + (("hello" + (("hello" + "foo bar") + ("\n\t" + "\n\t"))) + ("" + ("" + "\n\t"))))) {
        case "a":
            if (v4 > -975) {
                _ = len(v16)
                var v17 int16 = v4
                var v18 float32 = v1
                _ = v17
                _ = v18
            } else {
                switch 641 {
                case 540:
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v2 += v2
                        _ = fi19
                    }
                case 144:
                    _ = len(v16)
                case 32768:
                    fmt.Println(v1, v3)
                default:
                    if (((("foo bar" + "\n\t") + "hello") + ((("a" + "") + "a") + "a")) < "a") {
                        fmt.Println(v3, v3)
                    }
                }
                var v20 uint64 = 961
                v21 := GenS2{F0: 100, F1: 32768}
                _ = v20
                _ = v21
            }
        default:
            v22 := GenS1{F0: -128}
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 6; fi23++ {
            v16 = append(v16, (fi23 / ((-9223372036854775808) | 1)))
            _ = fi23
        }
    }
    var v24 bool = true
    var v25 int8 = 42
    v16 = append(v16, -988)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
