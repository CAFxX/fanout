package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint32) int32 {
    v5 := []int{127, -2147483648, 0, -2147483648}
    switch "a" {
    case "":
        var v6 int32 = -2147483648
        _ = v6
    case "hello":
        p0 -= p0
    case "foo bar":
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
    }
    if false {
        p0 = 65535
    }
    if (1e10 == float32(((0.0 * 0.0) / 3.14159))) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v7 bool = (false || true)
        _ = v7
    } else {
        var v8 uint64 = 98
        v5 = append(v5, -852)
        _ = v8
    }
    _ = p0
    _ = v5
    return -2147483648
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 128
    var v3 int32 = -629
    var v4 float64 = 0.1
    switch 563 {
    case 4294967295:
        var v9 int64 = -9223372036854775808
        _ = v9
    case -2147483648:
        v10 := []int{0, 990}
        _ = v10
    default:
        var v11 bool = true
        _ = v11
    }
    v12 := GenS1{F0: 33}
    v13 := GenS1{F0: 127}
    var v14 uint16 = 42
    var v15 int16 = 32767
    switch v1 {
    case "hello":
        if true {
            if (!(0 != -268)) {
                var v16 float64 = 1e10
                _ = v16
            } else {
                v17 := []int{1, int(v2), -735, 913}
                _ = v17
            }
        } else {
            v15 = (v15 & v15)
        }
    }
    fmt.Println(v15)
    switch 127 {
    case -502:
        if false {
            v1 = ((("\n\t" + (v1 + "a")) + v1) + v1)
            var v18 bool = (200 == 4294967295)
            switch ("a" + (v1 + "")) {
            case "foo bar":
                var v19 uint = 0
                _ = v19
            case "":
                var v20 int = 994
                _ = v20
            case "hello":
                var v21 int32 = v3
                _ = v21
            }
            _ = v18
        } else {
            if (true && false) {
                v15 += (v15 - v15)
                fmt.Println(v4, v2)
                v2 = (^(v2))
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 uint32 = 42
                _ = fi22
                _ = v23
            }
            v15 += (v15 % ((v15) | 1))
        }
    case 267:
        v24 := GenS2{F0: 849}
        _ = v24
    case 4294967295:
        v25 := []int{32768, 4294967295, 42, 42}
        _ = v25
    default:
        if false {
            var v26 uint64 = 18446744073709551615
            if (((!false) || (!false)) || true) {
                v27 := GenS1{F0: 0}
                v28 := []int{-1, int(v14)}
                for fi29 := 0; fi29 < 15; fi29++ {
                    v28 = append(v28, fi29)
                    var v30 int32 = v3
                    _ = fi29
                    _ = v30
                }
                _ = v27
                _ = v28
            }
            _ = v26
        } else {
            v2 += uint8(v14)
            var v31 bool = (3.14159 < 1e10)
            fmt.Println(v1)
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
