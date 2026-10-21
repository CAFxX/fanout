package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 bool }
type GenS2 struct { F0 uint16; F1 int; F2 uint }
func genFn0() int64 {
    var v7 string = "a"
    v8 := GenS1{F0: -47, F1: 96, F2: true}
    var v9 float64 = ((float64((1e10 + 1.5)) * (-2.25 * 1.5)) * 0.0)
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 bool = true
        fmt.Println(v9, fi10)
        if (!v11) {
            v12 := GenS1{F0: 32767, F1: 0, F2: v11}
            _ = v12
        } else {
            fmt.Println(v9, v9, v9)
            v13 := []int{32767, (fi10 / ((4294967295) | 1)), (fi10 | 4294967295), (fi10 % (((fi10 % ((32767) | 1))) | 1))}
            v9 = (float64(3.14159) * (v9 + (v9 + float64((1.5 * -2.25)))))
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 uint32 = 65535
    _ = v14
    _ = v7
    _ = v8
    _ = v9
    return 306
}
func genFn1(p0 int8, p1 int8, p2 uint) uint8 {
    fmt.Println(p2, p2)
    v15 := GenS1{F0: 1000, F1: 193, F2: ((p2 * p2) < p2)}
    var v16 float64 = (float64(3.14159) / 3.14159)
    v16 -= 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 128
}

func main() {
    var v1 uint8 = 255
    var v2 uint = 42
    v3 := []int{32768, -1, 190, 42}
    var v4 string = ""
    var v5 string = v4
    var v6 uint16 = 393
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint16 = 65535
        _ = fi17
        _ = v18
    }
    fmt.Println(v5, v5, v1)
    var v19 bool = (!true)
    fmt.Println(v6)
    if false {
        v6 = 0
        v1 += v1
        if true {
            if len(v3) > 0 {
                v3[0] = 32767
            }
            switch v4 {
            case "\n\t":
                v6 = uint16(v1)
            case "a":
                v6 = v6
            case "":
                var v20 int64 = -491
                _ = v20
            default:
                var v21 int32 = 2147483647
                _ = v21
            }
            var v22 string = (v4 + "hello")
            _ = v22
        }
    }
    _ = len(v3)
    fmt.Println(v5)
    var v23 uint32 = 65535
    v24 := GenS1{F0: 291, F1: v1, F2: true}
    switch 127 {
    case -712:
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    case 127:
        var v25 float64 = 0.1
        _ = v25
    }
    var v26 uint64 = 449
    v27 := GenS2{F0: (v6 >> ((v2) & 63)), F1: 9223372036854775807, F2: v2}
    var v28 uint16 = (v6 * (v6 / ((v6) | 1)))
    v29 := GenS1{F0: -32768, F1: v1, F2: v19}
    v23 = (v23 | v23)
    v30 := GenS2{F0: 32767, F1: 32767, F2: v2}
    for fi31 := 0; fi31 < 14; fi31++ {
        if ((uint32(v1) > 1) || false) {
            var v32 int16 = int16(v1)
            for fi33 := 0; fi33 < 5; fi33++ {
                var v34 int64 = genFn0()
                _ = fi33
                _ = v34
            }
            _ = v32
        }
        v1 += 0
        var v35 float32 = (1e10 - 1e10)
        _ = fi31
        _ = v35
    }
    for fi36 := 0; fi36 < 17; fi36++ {
        var v37 uint = 42
        v19 = v19
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
