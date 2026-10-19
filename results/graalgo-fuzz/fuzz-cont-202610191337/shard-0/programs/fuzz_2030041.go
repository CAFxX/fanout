package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint16; F2 uint }
func genFn0(p0 uint8) uint8 {
    p0 -= 0
    var v5 string = ("foo bar" + (("foo bar" + ("" + "a")) + ("" + "\n\t")))
    var v6 uint64 = 561
    fmt.Println(v6, v6, p0)
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int64) int8 {
    for fi7 := 0; fi7 < 9; fi7++ {
        switch "a" {
        case "":
            var v8 uint16 = 65535
            _ = v8
        case "a":
            fmt.Println(p0, fi7)
        case "\n\t":
            v9 := GenS1{F0: 0}
            _ = v9
        default:
            var v10 int16 = 0
            _ = v10
        }
        for fi11 := 0; fi11 < 7; fi11++ {
            p0 = p0
            p0 += -698
            var v12 int = fi7
            _ = fi11
            _ = v12
        }
        _ = fi7
    }
    v13 := []int{}
    v13 = append(v13, 32768)
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            p0 += p0
            fmt.Println(p0, fi15, fi14)
            v16 := GenS1{F0: 4294967295}
            _ = fi15
            _ = v16
        }
        switch (("a" + "\n\t") + ("" + "a")) {
        case "\n\t":
            var v17 uint64 = 356
            _ = v17
        case "hello":
            switch -267 {
            case -2147483648:
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 bool = false
                    v19 = (!true)
                    _ = fi18
                    _ = v19
                }
            case 4294967295:
                fmt.Println(fi14, p0)
            default:
                var v20 int16 = -892
                _ = v20
            }
        default:
            v21 := GenS1{F0: 1}
            _ = v21
        }
        var v22 uint8 = 2
        _ = fi14
        _ = v22
    }
    v23 := GenS1{F0: 2}
    _ = p0
    _ = v13
    _ = v23
    return -128
}

func main() {
    v1 := []string{"hello", "", ("foo bar" + "foo bar")}
    var v2 uint64 = 905
    v3 := []int{}
    var v4 int = -1
    for fi24 := 0; fi24 < 4; fi24++ {
        v4 += 4294967295
        _ = fi24
    }
    if false {
        fmt.Println(v2, v4)
    }
    v3 = append(v3, (^(v4)))
    v25 := GenS1{F0: 520}
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 uint32 = 482
        _ = fi26
        _ = v27
    }
    v28 := GenS2{F0: 4294967295, F1: 2, F2: 2}
    v2 -= 18446744073709551615
    v1 = append(v1, ((("hello" + "foo bar") + "foo bar") + ("\n\t" + "hello")))
    var v29 uint8 = 128
    var v30 uint16 = 0
    v2 = v2
    v31 := GenS2{F0: v4, F1: v30, F2: 903}
    var v32 uint32 = 0
    switch ("" + "hello") {
    case "\n\t":
        v30 = v30
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = 127
        }
    default:
        if len(v1) > 0 {
            v1[0] = ""
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
