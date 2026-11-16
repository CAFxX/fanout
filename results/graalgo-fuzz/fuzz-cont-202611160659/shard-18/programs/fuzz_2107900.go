package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 int32, p1 uint64, p2 uint64) int8 {
    p2 = p2
    if false {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := GenS1{F0: "hello"}
            _ = fi6
            _ = v7
        }
        var v8 int16 = 42
        for fi9 := 0; fi9 < 19; fi9++ {
            p2 += (p2 + (p2 * p1))
            v8 = v8
            if ((!(false && true)) && ("\n\t" < "hello")) {
                if ((v8 << ((462) & 63)) == v8) {
                    if true {
                        var v10 bool = ((881 == 2147483648) || (!false))
                        _ = v10
                    }
                } else {
                    p1 += (p2 * p2)
                }
                var v11 float32 = 1.5
                _ = v11
            }
            _ = fi9
        }
        _ = v8
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        p0 = 349
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}
func genFn1(p0 int, p1 string) bool {
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS2{F0: genFn0(-2147483648, 270, 0), F1: 78}
        p1 = p1
        fmt.Println(p0, fi13)
        _ = fi13
        _ = v14
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 uint32 = 306
    var v2 uint = 119
    var v3 uint32 = (v1 + v1)
    v4 := make(map[uint32]bool)
    v5 := make(map[string]uint32)
    switch "a" {
    case "hello":
        v1 += v3
    case "\n\t":
        var v15 float64 = float64(3.14159)
        _ = v15
    case "foo bar":
        v3 = (v3 & v1)
    default:
        var v16 int8 = -111
        _ = v16
    }
    var v17 float64 = float64(1e10)
    fmt.Println(v3)
    for fi18 := 0; fi18 < 7; fi18++ {
        v2 -= 587
        fmt.Println(v3, v17)
        delete(v5, "")
        _ = fi18
    }
    v19 := []int{32768}
    v2 += v2
    var v20 int8 = int8(v1)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println(v1)
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 uint32 = 282
        var v23 int16 = -942
        switch ("a" + ("hello" + ((("\n\t" + "foo bar") + "hello") + (("foo bar" + "hello") + "hello")))) {
        case "":
            for fi24 := 0; fi24 < 15; fi24++ {
                v25 := GenS2{F0: (v20 | -1), F1: v20}
                var v26 string = ("hello" + "a")
                v27 := GenS1{F0: v26}
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
        case "\n\t":
            v28 := GenS1{F0: ("" + "hello")}
            _ = v28
        case "hello":
            delete(v4, 65535)
        default:
            switch "a" {
            case "foo bar":
                v20 = 1
            case "a":
                switch (("\n\t" + "hello") + "hello") {
                case "\n\t":
                    if (107 >= 4294967295) {
                        v2 = v2
                        v23 = 107
                    }
                case "":
                    var v29 int8 = (v20 * 127)
                    _ = v29
                case "hello":
                    if len(v19) > 0 {
                        v19[0] = 128
                    }
                }
            case "hello":
                v23 -= (v23 / ((v23) | 1))
            }
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    var v30 uint16 = 804
    if (v1 < 0) {
        fmt.Println(v20, v30, v2)
    } else {
        v19 = append(v19, 619)
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        v17 = v17
        v3 = v3
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
