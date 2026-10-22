package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 bool }
func genFn0(p0 int32, p1 int64) int64 {
    var v7 bool = (-2.25 < 0.0)
    var v8 uint32 = 4294967295
    v8 = (v8 * (v8 / (((v8 & 4294967295)) | 1)))
    if false {
        if false {
            var v9 int8 = 1
            _ = v9
        }
        v10 := GenS1{F0: -546, F1: 42}
        _ = v10
    } else {
        v11 := GenS1{F0: 42, F1: p1}
        v8 = 1
        var v12 float32 = (-2.25 * (3.14159 + 3.14159))
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 uint32) uint16 {
    if true {
        v13 := []int{1}
        v14 := GenS2{F0: 128}
        _ = v13
        _ = v14
    }
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "\n\t"
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    var v5 uint64 = 42
    v6 := make(map[int]bool)
    v5 += 0
    fmt.Println(v2, v2)
    var v15 string = v2
    if false {
        v16 := GenS3{F0: 736}
        switch "" {
        case "\n\t":
            v2 = (v2 + ("a" + (v2 + v2)))
        default:
            v17 := GenS1{F0: 1, F1: 9223372036854775807}
            _ = v17
        }
        var v18 uint16 = genFn1(197)
        _ = v16
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        v2 = "\n\t"
        _ = len(v4)
        v2 = v15
        _ = fi19
    }
    fmt.Println(v15)
    var v20 string = (v15 + ("foo bar" + (v15 + "a")))
    v21 := GenS4{F0: -1, F1: (true && true)}
    var v22 uint16 = 65535
    switch ("\n\t" + v15) {
    case "\n\t":
        var v23 uint64 = v5
        _ = v23
    case "hello":
        switch int(v5) {
        case 42:
            v24 := []int{1}
            _ = v24
        case 127:
            v6[415] = (((1e10 - 1.5) <= 1e10) || true)
        case 32768:
            if (42 < 4294967295) {
                var v25 bool = (2 != 2147483648)
                var v26 int64 = 1
                _ = v25
                _ = v26
            }
        }
    }
    var v27 int = 9223372036854775807
    var v28 int = (^(v27))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
