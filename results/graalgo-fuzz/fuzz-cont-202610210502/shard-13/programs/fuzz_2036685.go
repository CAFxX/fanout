package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0() int16 {
    var v6 uint8 = 0
    fmt.Println(v6)
    var v7 int8 = 100
    fmt.Println(v6, v6)
    _ = v6
    _ = v7
    return 1000
}
func genFn1(p0 uint8, p1 uint16) uint32 {
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint = 255
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 282}
    v11 := []int{-9223372036854775808, 611}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 924
}

func main() {
    var v1 uint16 = 536
    var v2 uint8 = 42
    var v3 int16 = -32768
    v4 := make(map[uint32]string)
    var v5 int = 42
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        v2 -= 75
    }
    switch "" {
    case "foo bar":
        v12 := GenS1{F0: v5}
        _ = v12
    }
    v5 = (-(v5))
    _ = len(v4)
    var v13 uint8 = 169
    var v14 float32 = -2.25
    var v15 uint = 953
    var v16 uint = v15
    if (((float64(1e10) / 1.0) <= (0.0 * 3.14159)) || true) {
        v17 := GenS2{F0: (^(v13)), F1: v15, F2: ""}
        v18 := GenS1{F0: (v5 * 32767)}
        delete(v4, 2147483648)
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: (v16 - (v15 | (v16 - (v16 % ((v15) | 1))))), F1: 940}
    v13 = v13
    v20 := v4[genFn1(v2, 0)]
    if (((-1 >= -9223372036854775808) && (true && false)) || (463 == -1)) {
        var v21 string = v20
        _ = v21
    }
    v20 = v20
    v22 := []int{v5}
    _ = len(v22)
    switch v20 {
    case "hello":
        v1 = 270
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
