package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint64, p1 string, p2 uint8) int16 {
    fmt.Println(p0)
    v6 := []int{0, 32768, 2147483647, 3}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}

func main() {
    var v1 bool = false
    var v2 float64 = 3.14159
    var v3 int16 = 0
    v4 := make(map[string]string)
    var v5 uint32 = 2147483647
    switch -9223372036854775808 {
    case -408:
        v1 = true
    case 128:
        var v7 float32 = (((1e10 - float32(v2)) * 0.1) - -2.25)
        _ = v7
    case -1:
        v8 := []int{}
        _ = v8
    default:
        fmt.Println(v5)
    }
    v5 -= v5
    _ = len(v4)
    if (42 > 4294967295) {
        v2 = ((v2 * v2) / ((3.14159 / 1.0) * (v2 * (v2 + v2))))
        var v9 bool = v1
        v10 := GenS1{F0: 42, F1: 42, F2: "\n\t"}
        _ = v10
        _ = v9
    }
    switch "\n\t" {
    case "":
        v11 := []int{0, 32767}
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 string = "a"
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := v4[""]
            fmt.Println(v5, v13)
            v16 := []int{fi12, (^(fi14)), -2147483648}
            _ = fi14
            _ = v15
            _ = v16
        }
        if v1 {
            if (!v1) {
                var v17 int32 = 2147483647
                v5 = (v5 % (((v5 * 4294967295)) | 1))
                _ = v17
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                v5 = v5
                v19 := []int{}
                delete(v4, ((v13 + "") + "hello"))
                _ = fi18
                _ = v19
            }
            v3 += (v3 * genFn0(773, v13, 99))
        }
        _ = fi12
        _ = v13
    }
    v20 := GenS2{F0: 255, F1: -9223372036854775808}
    switch 32767 {
    case -305:
        var v21 float64 = v2
        _ = v21
    default:
        v1 = (float64(1e10) >= float64(((-2.25 / 0.1) / ((3.14159 / 1e10) * -2.25))))
    }
    if true {
        var v22 uint = 129
        v23 := GenS1{F0: 1, F1: 4294967295, F2: ("\n\t" + "a")}
        v24 := GenS2{F0: 42, F1: 1}
        _ = v22
        _ = v23
        _ = v24
    } else {
        v2 = -2.25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
