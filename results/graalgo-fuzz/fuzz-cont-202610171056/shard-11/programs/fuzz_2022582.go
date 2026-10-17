package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 bool, p1 uint16) int32 {
    fmt.Println(p1, p1)
    p1 = (p1 + p1)
    _ = p0
    _ = p1
    return -696
}

func main() {
    var v1 uint8 = 128
    var v2 string = "\n\t"
    var v3 uint16 = 5
    var v4 float32 = ((float32(1e10) / 0.1) - 1e10)
    v5 := make(map[string]int)
    var v6 int8 = 62
    switch -1 {
    case 42:
        v7 := []int{-9223372036854775808}
        _ = v7
    case 333:
        fmt.Println(v4)
    case 32768:
        v2 = (v2 + "a")
    default:
        if (true && (32768 != -756)) {
            var v8 bool = false
            var v9 bool = v8
            v10 := GenS1{F0: 2}
            _ = v10
            _ = v8
            _ = v9
        }
    }
    var v11 uint32 = 34
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS2{F0: "", F1: v3}
        v5["foo bar"] = 127
        _ = fi12
        _ = v13
    }
    var v14 string = v2
    var v15 bool = true
    v2 = v2
    var v16 int64 = -226
    var v17 uint = 994
    fmt.Println(v1, v17)
    var v18 uint8 = v1
    _ = len(v5)
    var v19 uint16 = 65535
    if (((v14 + "foo bar") + v2) <= (((v14 + "hello") + "foo bar") + v2)) {
        var v20 string = "a"
        _ = v20
    }
    switch -628 {
    case -753:
        var v21 int64 = -1
        _ = v21
    case 32767:
        v22 := []int{}
        _ = v22
    case -1:
        if v15 {
            v23 := v5[((v2 + v2) + (v2 + (v14 + "")))]
            if (v15 || (v23 > (v23 >> (((v17 >> (((v17 * (v17 ^ v17))) & 63))) & 63)))) {
                var v24 bool = (1e10 < (((0.0 * float32((float64(v4) * float64((((v4 + v4) * v4) * 0.0))))) - 0.1) - (v4 / v4)))
                v25 := GenS1{F0: (v11 ^ (v11 | uint32(v17)))}
                _ = v24
                _ = v25
            }
            _ = v23
        }
    default:
        var v26 float64 = 3.14159
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        v2 = v14
        _ = fi27
    }
    delete(v5, v14)
    var v28 int32 = genFn0(v15, (^(v3)))
    if (2147483648 != (v11 | v11)) {
        v14 = (v2 + (((v14 + v14) + "\n\t") + "\n\t"))
        fmt.Println(v6)
        for fi29 := 0; fi29 < 8; fi29++ {
            fmt.Println(v4, v2, v1)
            _ = fi29
        }
    } else {
        v28 += genFn0(v15, (v19 ^ (v3 - 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
