package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int64; F2 int16 }
func genFn0(p0 string) uint64 {
    v5 := []int{294}
    var v6 uint32 = 65535
    v7 := GenS1{F0: 32767, F1: 42}
    for fi8 := 0; fi8 < 16; fi8++ {
        switch fi8 {
        case 945:
            var v9 int32 = 42
            _ = v9
        case 156:
            var v10 uint = 474
            _ = v10
        default:
            fmt.Println(fi8, p0, v6)
        }
        _ = fi8
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 bool = (!((false || false) || ((((("hello" + ("\n\t" + "\n\t")) + ("foo bar" + ("hello" + ""))) + "") + "") == "hello")))
    var v2 float32 = 0.0
    v3 := make(map[uint32]int)
    var v4 uint8 = 174
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: 0, F1: 1}
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 988, F1: 100}
    v2 = (v2 - (float32(((0.1 / 1e10) / 0.1)) * 0.0))
    fmt.Println(v1, v2, v2)
    var v14 float64 = -2.25
    if v1 {
        if v1 {
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := GenS2{F0: 32767}
                var v17 int32 = 349
                v18 := v3[807]
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        v19 := GenS1{F0: 849, F1: -106}
        _ = v19
    } else {
        fmt.Println(v4)
        var v20 float32 = v2
        var v21 float64 = v14
        _ = v20
        _ = v21
    }
    fmt.Println(v1, v4, v2)
    var v22 int16 = -32768
    v1 = v1
    var v23 string = ""
    var v24 float64 = (v14 * (-2.25 + v14))
    v22 = 32767
    fmt.Println(v24)
    v22 -= v22
    fmt.Println(v22, v4)
    if v1 {
        var v25 float32 = 1e10
        v26 := v3[42]
        v22 += v22
        _ = v25
        _ = v26
    } else {
        v27 := GenS3{F0: v4, F1: 1, F2: 32767}
        v22 -= v22
        _ = v27
    }
    v22 -= 1000
    for fi28 := 0; fi28 < 10; fi28++ {
        if (!v1) {
            if (2147483647 < 0) {
                v22 += v22
                v29 := GenS3{F0: (v4 << ((2) & 63)), F1: -1, F2: (v22 << ((255) & 63))}
                _ = v29
            } else {
                for fi30 := 0; fi30 < 2; fi30++ {
                    var v31 float32 = 3.14159
                    v32 := GenS1{F0: (v22 >> ((2) & 63)), F1: 1}
                    delete(v3, 65535)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        }
        v1 = (!(!v1))
        var v33 float32 = -2.25
        _ = fi28
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
