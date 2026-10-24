package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint16, p1 uint, p2 uint) string {
    v4 := GenS1{F0: 2, F1: 649}
    v5 := GenS2{F0: 42, F1: ("a" + ("a" + ("hello" + ("" + ("\n\t" + "a")))))}
    v6 := GenS3{F0: -1}
    if ((p0 & (p0 + (p0 | 42))) < 163) {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 uint32 = 2
            _ = fi7
            _ = v8
        }
        if true {
            v9 := []int{128}
            fmt.Println(p2)
            var v10 uint16 = (^(p0))
            _ = v10
            _ = v9
        }
    }
    switch "hello" {
    case "a":
        for fi11 := 0; fi11 < 16; fi11++ {
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 uint32 = 2147483647
                p0 = (^(p0))
                if false {
                    v14 := []int{}
                    p1 = 701
                    var v15 bool = true
                    _ = v14
                    _ = v15
                } else {
                    v13 += v13
                }
                _ = fi12
                _ = v13
            }
            fmt.Println(p2, fi11, p2)
            _ = fi11
        }
    case "foo bar":
        var v16 uint16 = 2
        _ = v16
    case "hello":
        if (0 > 1) {
            p2 = 65535
            v17 := []int{83}
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    var v3 string = "\n\t"
    for fi18 := 0; fi18 < 18; fi18++ {
        fmt.Println(v2, v3)
        switch 1 {
        case 1:
            v2 = v2
        default:
            v19 := v1[fi18]
            _ = v19
        }
        _ = fi18
    }
    v20 := GenS2{F0: 100, F1: "hello"}
    for fi21 := 0; fi21 < 19; fi21++ {
        _ = len(v1)
        v22 := GenS4{F0: 477}
        _ = fi21
        _ = v22
    }
    var v23 int32 = 704
    for fi24 := 0; fi24 < 18; fi24++ {
        v2 = (false && false)
        _ = fi24
    }
    v2 = false
    _ = len(v1)
    if (32767 < -23) {
        v23 = -2147483648
        var v25 uint8 = 179
        _ = v25
    } else {
        switch "\n\t" {
        case "foo bar":
            var v26 int8 = -128
            _ = v26
        case "\n\t":
            var v27 bool = v2
            _ = v27
        case "hello":
            v3 = v3
        default:
            var v28 uint16 = 42
            _ = v28
        }
        var v29 uint8 = 42
        _ = v29
    }
    var v30 uint32 = 688
    v31 := []int{32768}
    var v32 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
