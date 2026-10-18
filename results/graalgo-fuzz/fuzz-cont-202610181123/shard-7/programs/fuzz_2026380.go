package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint16; F1 int8 }
func genFn0() uint8 {
    v4 := []int{273, -563, 1, 4294967295}
    var v5 float32 = float32(-2.25)
    var v6 uint8 = 1
    _ = v4
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 uint64, p1 int8, p2 int8) uint64 {
    p0 = p0
    v7 := GenS1{F0: 42, F1: 169, F2: 1}
    p2 += (p1 % ((-1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 ^ p0)
}

func main() {
    var v1 int = 70
    var v2 uint = 18446744073709551615
    var v3 int8 = -85
    var v8 uint64 = 629
    switch v1 {
    case -1:
        if true {
            v9 := []int{}
            var v10 int = v1
            _ = len(v9)
            _ = v10
            _ = v9
        }
    case 90:
        v11 := GenS2{F0: genFn1((^(v8)), v3, (v3 + 1))}
        _ = v11
    }
    v12 := GenS3{F0: 1}
    var v13 int32 = 0
    v1 += -720
    fmt.Println(v8, v3)
    v8 += (v8 + v8)
    switch -577 {
    case 32767:
        v14 := GenS4{F0: "a", F1: 42}
        _ = v14
    case -250:
        var v15 string = ""
        _ = v15
    }
    if (true || true) {
        var v16 uint16 = 255
        _ = v16
    }
    if false {
        var v17 int16 = int16(v8)
        var v18 int = 9223372036854775807
        _ = v17
        _ = v18
    }
    var v19 bool = (!(70 <= 255))
    switch (v1 / ((v1) | 1)) {
    case -9223372036854775808:
        if (v3 >= (v3 - 100)) {
            v20 := []int{}
            var v21 int64 = -497
            switch v1 {
            case 32767:
                fmt.Println(v8)
            case -58:
                var v22 uint8 = 42
                _ = v22
            case 840:
                for fi23 := 0; fi23 < 5; fi23++ {
                    var v24 uint64 = (v8 - genFn1((v8 ^ (^(v8))), v3, 121))
                    _ = fi23
                    _ = v24
                }
            }
            _ = v20
            _ = v21
        } else {
            v25 := []int{}
            var v26 int = (v1 >> ((323) & 63))
            _ = v25
            _ = v26
        }
    default:
        var v27 float32 = 0.0
        _ = v27
    }
    switch "foo bar" {
    case "\n\t":
        var v28 int = v1
        _ = v28
    case "hello":
        v8 += 2
    case "foo bar":
        v8 += v8
    }
    v29 := GenS5{F0: 32767, F1: (v3 - v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v8
}
