package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint16 = 0
    var v2 float32 = 3.14159
    var v3 int16 = 1
    v2 = v2
    if (-1 == 172) {
        for fi4 := 0; fi4 < 19; fi4++ {
            switch ("a" + "hello") {
            case "foo bar":
                v1 += 2
            default:
                fmt.Println(v1)
            }
            v2 -= v2
            var v5 uint16 = (v1 / (((^(v1))) | 1))
            _ = fi4
            _ = v5
        }
    } else {
        var v6 uint8 = 128
        if true {
            fmt.Println(v6)
            var v7 uint16 = v1
            _ = v7
        }
        var v8 int8 = 0
        _ = v6
        _ = v8
    }
    switch 858 {
    case 0:
        if false {
            if (((v1 - (^(v1))) < (^(v1))) && (false && false)) {
                var v9 float64 = 3.14159
                v3 = -760
                _ = v9
            }
            switch 4294967295 {
            case -9223372036854775808:
                v1 += v1
            }
            v10 := []int{32767, 128, -911, 128}
            _ = v10
        } else {
            var v11 string = ("foo bar" + "")
            _ = v11
        }
    case 32768:
        v12 := GenS1{F0: 622}
        _ = v12
    case 127:
        v1 += (v1 ^ 255)
    }
    v1 = (v1 + v1)
    for fi13 := 0; fi13 < 15; fi13++ {
        if false {
            v2 -= 1e10
            var v14 float64 = (float64((0.1 + v2)) * ((0.1 - -2.25) + ((3.14159 * (float64(float32(0.0)) - (0.1 / 1e10))) - 0.1)))
            _ = v14
        } else {
            var v15 int64 = 0
            _ = v15
        }
        var v16 int32 = 2147483647
        _ = fi13
        _ = v16
    }
    switch 9223372036854775807 {
    case -107:
        v17 := GenS1{F0: 42}
        _ = v17
    case 65:
        var v18 string = "hello"
        _ = v18
    default:
        if (("a" < "") && (((true || (1e10 == ((float64(float32(1e10)) + ((0.1 * 0.1) - -2.25)) - 1e10))) || (((42 == 2) && (!true)) || true)) || true)) {
            var v19 float64 = -2.25
            var v20 uint64 = 42
            _ = v19
            _ = v20
        }
    }
    var v21 bool = true
    v22 := []int{}
    var v23 int32 = 1
    var v24 uint32 = 4294967295
    v2 = v2
    if (("foo bar" + "a") != ("hello" + ("foo bar" + "hello"))) {
        var v25 string = "hello"
        _ = v25
    }
    fmt.Println(v1, v23)
    var v26 uint16 = v1
    var v27 int = 32768
    var v28 int = v27
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
