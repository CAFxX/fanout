package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint8 }
func genFn0(p0 int8) uint32 {
    var v4 int64 = -1
    v5 := GenS1{F0: 1, F1: -208}
    if true {
        v6 := []int{}
        _ = v6
    } else {
        if (true || (2147483647 >= 1)) {
            if (!(42 < 323)) {
                p0 = p0
            } else {
                var v7 float64 = (float64(1.5) - (((3.14159 * 1.5) + 1.5) - 3.14159))
                v8 := GenS2{F0: 128, F1: 0, F2: 592}
                _ = v7
                _ = v8
            }
            v9 := GenS3{F0: 18446744073709551615}
            _ = v9
        }
        var v10 int = -9223372036854775808
        _ = v10
    }
    _ = p0
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint) string {
    v11 := GenS1{F0: 658, F1: -486}
    p0 += 255
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            p0 = p0
            _ = fi13
        }
        _ = fi12
    }
    fmt.Println(p0)
    _ = p0
    _ = v11
    return ("" + ("a" + ((("hello" + ("a" + ("\n\t" + "\n\t"))) + "") + "foo bar")))
}

func main() {
    var v1 float64 = 1e10
    var v2 int64 = 9223372036854775807
    var v3 string = ("foo bar" + "")
    if ((-2.25 * 1e10) < 0.0) {
        v14 := GenS2{F0: 128, F1: 255, F2: 67}
        switch 874 {
        case 269:
            v3 = v3
        case 128:
            var v15 string = ((v3 + (v3 + "a")) + "hello")
            _ = v15
        default:
            v16 := GenS4{F0: 494, F1: 899, F2: 2}
            _ = v16
        }
        var v17 uint = 4294967295
        _ = v14
        _ = v17
    }
    v18 := GenS1{F0: 90, F1: 0}
    var v19 int32 = 0
    v20 := []int{}
    v20 = append(v20, 699)
    v3 = "hello"
    switch "hello" {
    case "hello":
        for fi21 := 0; fi21 < 11; fi21++ {
            fmt.Println(v1, v3)
            var v22 float64 = v1
            _ = fi21
            _ = v22
        }
    case "":
        var v23 int8 = 100
        _ = v23
    case "foo bar":
        _ = len(v20)
    }
    var v24 int32 = v19
    switch (("a" + "\n\t") + (v3 + "")) {
    case "":
        v25 := GenS3{F0: 2}
        _ = v25
    case "\n\t":
        v26 := GenS3{F0: 679}
        _ = v26
    default:
        for fi27 := 0; fi27 < 4; fi27++ {
            if false {
                var v28 bool = true
                _ = v28
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                var v30 uint8 = 2
                _ = fi29
                _ = v30
            }
            _ = fi27
        }
    }
    v2 -= -9223372036854775808
    v3 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
}
