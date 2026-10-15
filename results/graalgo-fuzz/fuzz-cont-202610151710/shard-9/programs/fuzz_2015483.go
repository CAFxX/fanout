package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint8 }

func main() {
    var v1 float32 = float32(1e10)
    var v2 bool = ((!(true && false)) || (!((782 <= -324) || true)))
    v3 := []int{624}
    if true {
        var v4 float32 = (v1 - v1)
        v5 := GenS1{F0: 255, F1: 42, F2: 241}
        v4 = v1
        _ = v4
        _ = v5
    } else {
        _ = len(v3)
        v6 := GenS1{F0: 0, F1: 42, F2: 2}
        if (4294967295 < 635) {
            _ = len(v3)
            v7 := GenS1{F0: 255, F1: 0, F2: 255}
            _ = v7
        } else {
            var v8 int16 = -1
            _ = v8
        }
        _ = v6
    }
    fmt.Println(v2, v1)
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            var v10 string = (("foo bar" + "") + (("\n\t" + "hello") + ("foo bar" + "a")))
            if len(v3) > 0 {
                v3[0] = 32768
            }
            switch 692 {
            case 1:
                v10 = "a"
            case 127:
                v1 = v1
            }
            _ = v10
        } else {
            if (true && v2) {
                var v11 uint = 18446744073709551615
                var v12 int64 = 42
                _ = v11
                _ = v12
            }
        }
        _ = len(v3)
        _ = fi9
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint8 = 2
        _ = fi13
        _ = v14
    }
    v2 = (!(2 <= 172))
    var v15 float32 = v1
    var v16 int64 = 9223372036854775807
    fmt.Println(v2, v15, v2)
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 float64 = 0.1
        var v19 bool = v2
        switch -2147483648 {
        case -9223372036854775808:
            fmt.Println(v16)
        default:
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 int8 = 100
                if (v2 || true) {
                    fmt.Println(v16)
                    v21 -= v21
                }
                v2 = v2
                _ = fi20
                _ = v21
            }
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        fmt.Println(v1, v2)
        _ = fi22
    }
    if (719 < -32768) {
        fmt.Println(v15, v2)
        v3 = append(v3, -677)
        switch 626 {
        case -669:
            v2 = true
        default:
            var v23 int = 2147483647
            _ = v23
        }
    }
    var v24 float32 = (v1 + 1e10)
    v1 -= (v1 + 0.1)
    if true {
        v15 += (((v1 - v1) / (float32((1.5 * 1.5)) + 3.14159)) + v24)
    }
    var v25 uint8 = 42
    var v26 uint64 = 42
    v27 := GenS1{F0: 2, F1: 0, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
