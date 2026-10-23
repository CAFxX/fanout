package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 int16 = 398
    v1 += 42
    if (683 == -511) {
        v4 -= (v4 | int16(v1))
        fmt.Println(v4, v1)
    } else {
        var v5 int64 = -1
        v4 = 1000
        v1 += (v1 ^ (v1 / ((65535) | 1)))
        _ = v5
    }
    v6 := []int{-826, -730, 12}
    var v7 int = -2147483648
    _ = len(v6)
    if ((!(!(false || (!true)))) || (true || (!false))) {
        v4 = (v4 | (v4 << ((4294967295) & 63)))
    } else {
        switch "hello" {
        case "":
            v4 += 32767
        case "\n\t":
            var v8 uint = 255
            _ = v8
        default:
            var v9 uint = 42
            _ = v9
        }
        if (1 > 1) {
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := GenS1{F0: 1}
                _ = fi10
                _ = v11
            }
            var v12 uint16 = 824
            _ = v12
        } else {
            var v13 bool = ((-2.25 < float64((0.0 * -2.25))) && false)
            _ = v13
        }
    }
    var v14 int16 = 42
    var v15 float64 = ((((1.5 / 1.5) * float64(float32(1.5))) + -2.25) * (1e10 * (3.14159 / 1.0)))
    if true {
        if false {
            var v16 int64 = -1
            _ = v16
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 uint32 = 2
            _ = fi17
            _ = v18
        }
        if true {
            _ = len(v2)
            var v19 bool = (!true)
            var v20 float32 = -2.25
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 float32 = (3.14159 / 1.0)
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 2}
    v6 = append(v6, (v7 * (v7 | 2147483647)))
    var v24 string = "a"
    fmt.Println(v1)
    v25 := GenS1{F0: 18446744073709551615}
    fmt.Println(v1, v24, v15)
    for fi26 := 0; fi26 < 1; fi26++ {
        _ = len(v2)
        switch v7 {
        case 417:
            var v27 int32 = 42
            _ = v27
        case -652:
            var v28 bool = false
            _ = v28
        default:
            for fi29 := 0; fi29 < 5; fi29++ {
                v4 += v14
                var v30 int32 = 505
                _ = fi29
                _ = v30
            }
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
