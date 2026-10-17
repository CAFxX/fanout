package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }
func genFn0() int32 {
    if ((92 < 18446744073709551615) && (!(-501 < 9223372036854775807))) {
        var v5 int32 = -257
        if false {
            var v6 int32 = (^(v5))
            _ = v6
        } else {
            var v7 float64 = 0.0
            _ = v7
        }
        _ = v5
    } else {
        var v8 int16 = 0
        v9 := GenS1{F0: ("" + "a")}
        switch 1 {
        case -1:
            v8 = 32767
        case 2147483647:
            v10 := []int{128}
            _ = v10
        }
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 uint32 = 2
            v13 = v13
            _ = fi12
            _ = v13
        }
        v14 := GenS2{F0: 0, F1: 4294967295}
        var v15 int32 = 0
        _ = fi11
        _ = v14
        _ = v15
    }
    if false {
        var v16 uint8 = 236
        v16 -= v16
        _ = v16
    } else {
        v17 := GenS2{F0: 42, F1: 1}
        _ = v17
    }
    var v18 int8 = 100
    v19 := GenS3{F0: 255}
    _ = v18
    _ = v19
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -128
    var v3 uint = 172
    v4 := make(map[string]uint32)
    v20 := GenS2{F0: 1000, F1: (v3 * v3)}
    v2 += v2
    for fi21 := 0; fi21 < 6; fi21++ {
        fmt.Println(fi21, v3, v2)
        v22 := GenS4{F0: 42, F1: 826, F2: 44}
        _ = fi21
        _ = v22
    }
    fmt.Println(v3, v2)
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 int32 = -101
        var v25 uint32 = 2147483647
        for fi26 := 0; fi26 < 14; fi26++ {
            v2 = (v2 + (v2 + -90))
            _ = fi26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    v27 := []int{}
    switch ("foo bar" + "\n\t") {
    case "":
        var v28 int32 = 1
        _ = v28
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        delete(v4, "foo bar")
        _ = fi29
    }
    var v30 uint16 = 419
    v31 := GenS1{F0: (("hello" + "a") + ((("" + "foo bar") + "\n\t") + "hello"))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
