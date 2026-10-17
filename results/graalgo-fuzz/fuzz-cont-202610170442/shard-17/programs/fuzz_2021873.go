package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 string }
type GenS2 struct { F0 int64; F1 int32; F2 uint }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 int8, p1 int64) int32 {
    for fi5 := 0; fi5 < 11; fi5++ {
        if false {
            p1 -= p1
            v6 := []int{524, fi5, 601}
            _ = v6
        }
        var v7 bool = true
        _ = fi5
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        switch fi8 {
        case 1:
            var v9 float32 = 1e10
            _ = v9
        case 4294967295:
            fmt.Println(p1, p1, fi8)
        }
        if false {
            fmt.Println(p0, p0, p1)
            var v10 uint32 = 2
            _ = v10
        }
        v11 := []int{fi8, fi8, (^(fi8)), (fi8 | (fi8 % ((fi8) | 1)))}
        _ = fi8
        _ = v11
    }
    v12 := GenS1{F0: 65535, F1: 1, F2: (("a" + "") + ("" + "\n\t"))}
    _ = p0
    _ = p1
    _ = v12
    return -2147483648
}

func main() {
    var v1 uint8 = 42
    var v2 bool = (((0.1 * (0.1 - 0.1)) <= (0.1 / (((float32(-2.25) / 1e10) + -2.25) * -2.25))) || true)
    var v3 uint = 0
    var v4 uint = v3
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int32 = -1
        _ = fi13
        _ = v14
    }
    fmt.Println(v3, v1, v4)
    var v15 int64 = 42
    switch ("a" + "hello") {
    case "\n\t":
        v15 = v15
    default:
        v16 := GenS2{F0: 0, F1: -329, F2: uint(v1)}
        _ = v16
    }
    v4 = (v4 + (v4 / ((v3) | 1)))
    v17 := GenS1{F0: v4, F1: (^(v3)), F2: "hello"}
    switch -939 {
    case 84:
        if (!v2) {
            for fi18 := 0; fi18 < 6; fi18++ {
                v3 -= (^(v4))
                var v19 uint32 = 559
                _ = fi18
                _ = v19
            }
            fmt.Println(v1, v2, v1)
            if false {
                var v20 int16 = -1
                fmt.Println(v3, v2)
                _ = v20
            } else {
                v1 += (v1 * v1)
                fmt.Println(v3, v1, v2)
                v21 := GenS2{F0: 9223372036854775807, F1: -2147483648, F2: (^(v4))}
                _ = v21
            }
        } else {
            v4 -= v4
            v1 -= 255
            v15 = v15
        }
    case 128:
        v22 := []int{}
        _ = v22
    }
    var v23 uint8 = (v1 & (^(v1)))
    v24 := GenS3{F0: -32768, F1: -128}
    v3 = v4
    var v25 uint = (^(v4))
    fmt.Println(v1)
    v25 -= 65535
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
