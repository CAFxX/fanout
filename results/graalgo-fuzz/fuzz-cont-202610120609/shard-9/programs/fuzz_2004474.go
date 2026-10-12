package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 int64 }
func genFn0() uint8 {
    switch "\n\t" {
    case "":
        var v4 int32 = 42
        _ = v4
    default:
        var v5 int16 = 32767
        _ = v5
    }
    if false {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 bool = ((3.14159 / 0.1) != 1e10)
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 6; fi8++ {
            if true {
                var v9 uint16 = 185
                v9 += v9
                _ = v9
            }
            var v10 int16 = -876
            _ = fi8
            _ = v10
        }
    } else {
        var v11 bool = (!true)
        v12 := []int{-905, -9223372036854775808, -2147483648, 372}
        _ = len(v12)
        _ = v11
        _ = v12
    }
    var v13 uint = 18446744073709551615
    v14 := GenS1{F0: 0, F1: 18446744073709551615, F2: 42}
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]bool)
    v3 := []int{574}
    var v15 int32 = 1
    var v16 float32 = 0.1
    v15 = (v15 + v15)
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS2{F0: 109, F1: -1, F2: -9223372036854775808}
        var v19 int64 = 1
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 bool = false
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 int64 = 9223372036854775807
        fmt.Println(v1)
        v15 = v15
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v16 = 0.1
        _ = fi23
    }
    v24 := GenS1{F0: 32767, F1: 0, F2: -32768}
    v15 = (v15 ^ v15)
    var v25 bool = (929 < 65535)
    if (false && ((float32(0.0) != ((((float32((float64(v16) / 1.0)) / (v16 - float32(1.5))) - v16) * -2.25) * (v16 / v16))) || (42 != 100))) {
        for fi26 := 0; fi26 < 17; fi26++ {
            v27 := GenS1{F0: -32768, F1: v1, F2: -32768}
            var v28 float32 = 0.0
            var v29 float64 = 3.14159
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        var v30 float32 = v16
        if ((^(v15)) != -745) {
            v31 := GenS1{F0: 1, F1: (^(v1)), F2: -599}
            _ = v31
        }
        _ = v30
    } else {
        _ = len(v3)
        fmt.Println(v15, v1)
        var v32 float32 = v16
        _ = v32
    }
    fmt.Println(v1)
    switch ("\n\t" + (("a" + "") + ((("hello" + "hello") + ("\n\t" + "a")) + "hello"))) {
    case "foo bar":
        if v25 {
            switch 1 {
            case 42:
                switch (("a" + "foo bar") + "hello") {
                case "foo bar":
                    var v33 uint16 = 0
                    _ = v33
                }
            case 38:
                v34 := GenS1{F0: 0, F1: (^(v1)), F2: 42}
                _ = v34
            case 32768:
                fmt.Println(v16, v20, v16)
            default:
                var v35 uint32 = 2
                _ = v35
            }
            if v20 {
                v16 += (v16 * float32((0.1 + -2.25)))
            }
        }
    case "a":
        v2[""] = v20
    }
    switch -9223372036854775808 {
    case 2147483647:
        fmt.Println(v1)
    default:
        _ = len(v3)
    }
    v15 = v15
    v25 = true
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
}
