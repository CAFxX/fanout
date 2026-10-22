package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint32, p1 bool) string {
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := []int{1, (fi4 >> ((4294967295) & 63)), -491}
        p1 = (1 != 0)
        _ = fi4
        _ = v5
    }
    var v6 float64 = (1e10 * 0.1)
    var v7 string = "hello"
    v8 := GenS1{F0: p1, F1: 255}
    v9 := GenS2{F0: 81, F1: p0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v7 + v7)
}
func genFn1() uint8 {
    var v10 uint = 255
    var v11 uint8 = 2
    fmt.Println(v11, v11)
    v10 += (v10 - (v10 << ((v10) & 63)))
    v12 := []int{-1, 0, int(v11)}
    _ = v10
    _ = v11
    _ = v12
    return v11
}

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 2
    var v3 int16 = 32767
    if (false && false) {
        v13 := []int{-9223372036854775808}
        _ = v13
    }
    var v14 uint = 0
    var v15 uint16 = 125
    v16 := GenS3{F0: uint16(v2)}
    v15 += (v15 + (v15 + (v15 & (v15 ^ (v15 & v15)))))
    v2 -= v2
    fmt.Println(v3, v1, v14)
    switch 0 {
    case -811:
        v14 = (v14 ^ 42)
    case 9223372036854775807:
        v17 := []int{127}
        _ = v17
    default:
        var v18 int8 = 6
        _ = v18
    }
    var v19 float64 = (1e10 / 1.0)
    v20 := []int{32767, 239, 32767}
    if (float64(3.14159) > float64(0.0)) {
        v15 -= 255
        var v21 string = v1
        _ = v21
    } else {
        if len(v20) > 0 {
            v20[0] = 32767
        }
    }
    if len(v20) > 0 {
        v20[0] = 600
    }
    fmt.Println(v14, v3)
    v15 += 821
    var v22 bool = ((1e10 / 3.14159) <= 1e10)
    if ((!v22) && (!(v3 >= 346))) {
        var v23 uint32 = 65535
        _ = v23
    }
    v15 += v15
    if len(v20) > 0 {
        v20[0] = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
}
