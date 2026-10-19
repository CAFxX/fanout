package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint16; F1 int16; F2 int32 }
type GenS5 struct { F0 int; F1 string }
func genFn0(p0 uint8, p1 int16, p2 int8) int {
    if (!true) {
        p1 = p1
        var v6 int8 = p2
        _ = v6
    }
    var v7 int16 = (p1 | (p1 / (((p1 * p1)) | 1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 477
}
func genFn1(p0 uint8) uint {
    p0 -= 1
    fmt.Println(p0)
    v8 := []int{-2147483648, 2147483647, -2147483648}
    v9 := GenS1{F0: 100, F1: -9223372036854775808, F2: 42}
    _ = p0
    _ = v8
    _ = v9
    return 536
}

func main() {
    var v1 float32 = (((((1.5 * -2.25) * 1e10) - 0.0) - 0.0) * 0.1)
    var v2 string = "foo bar"
    var v3 uint16 = 42
    v4 := make(map[string]int)
    var v5 uint8 = 128
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int = fi10
        v2 = "hello"
        v11 = (v11 + -228)
        _ = fi10
        _ = v11
    }
    if false {
        fmt.Println(v3, v1, v1)
        for fi12 := 0; fi12 < 9; fi12++ {
            fmt.Println(fi12, fi12)
            v13 := GenS2{F0: v3}
            if false {
                var v14 int = (-(fi12))
                fmt.Println(v3, v5, fi12)
                var v15 float32 = (0.1 + float32(float64(v1)))
                _ = v14
                _ = v15
            } else {
                var v16 uint32 = 0
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    }
    fmt.Println(v1, v3, v3)
    if false {
        for fi17 := 0; fi17 < 17; fi17++ {
            for fi18 := 0; fi18 < 5; fi18++ {
                _ = len(v4)
                var v19 int32 = -583
                v20 := GenS1{F0: -128, F1: 9223372036854775807, F2: -1}
                v1 = v1
                _ = fi18
                _ = v19
                _ = v20
            }
            fmt.Println(v5, v1, fi17)
            _ = fi17
        }
        v2 = v2
    }
    switch 776 {
    case -2147483648:
        if (false && (!false)) {
            v2 = "hello"
        }
    case 32768:
        var v21 bool = ((true && (2147483647 > 1)) || false)
        _ = v21
    default:
        if (!false) {
            v3 = (v3 / ((945) | 1))
            if true {
                var v22 float64 = (0.0 / 1.5)
                _ = v22
            }
        }
    }
    v23 := GenS3{F0: 0, F1: 76}
    if ((3.14159 / float64(v1)) == 3.14159) {
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS4{F0: v3, F1: 32767, F2: 0}
            v26 := GenS4{F0: v3, F1: -1, F2: -519}
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    v2 = ((v2 + ("a" + "foo bar")) + "foo bar")
    switch ("a" + v2) {
    case "":
        switch 2147483647 {
        case 774:
            for fi27 := 0; fi27 < 14; fi27++ {
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 bool = true
                    v4[v2] = (fi28 % (((fi28 ^ fi27)) | 1))
                    v3 = (v3 ^ (v3 * v3))
                    _ = fi28
                    _ = v29
                }
                v1 = 3.14159
                _ = fi27
            }
        default:
            v30 := GenS2{F0: (v3 & 2)}
            _ = v30
        }
    case "foo bar":
        var v31 bool = (!false)
        _ = v31
    default:
        v32 := []int{42}
        _ = v32
    }
    fmt.Println(v5, v2, v3)
    v33 := GenS5{F0: -2147483648, F1: "\n\t"}
    v34 := v4[v2]
    v35 := v4[v2]
    v1 = (float32(1.5) * v1)
    fmt.Println(v5)
    v5 -= (^(v5))
    switch "a" {
    case "a":
        var v36 bool = (!(v2 >= (v2 + v2)))
        _ = v36
    case "":
        switch v34 {
        case 42:
            delete(v4, v2)
        }
    case "hello":
        var v37 uint8 = v5
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
