package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int; F1 string; F2 int64 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint }
func genFn0(p0 int, p1 bool, p2 int8) bool {
    v5 := GenS1{F0: 65535}
    p1 = (1 == 32767)
    var v6 uint16 = 2
    var v7 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 float64 = -2.25
    var v2 int64 = 1
    v3 := make(map[uint32]bool)
    var v4 uint64 = 4294967295
    switch "hello" {
    case "a":
        v8 := []int{32767, -438, 128, 931}
        _ = v8
    case "\n\t":
        fmt.Println(v4)
    }
    v9 := GenS2{F0: v4, F1: 1000}
    v10 := v3[2147483647]
    switch 17 {
    case -9223372036854775808:
        v1 = v1
    case -508:
        v11 := GenS3{F0: 2147483647, F1: "hello", F2: (^(v2))}
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        if (-128 < 49) {
            v10 = (!false)
            fmt.Println(v2, v1, v2)
        } else {
            v13 := GenS4{F0: v10, F1: 8}
            v14 := GenS5{F0: -1}
            v10 = (v10 || (float64(0.0) == v1))
            _ = v13
            _ = v14
        }
        var v15 float64 = v1
        if v10 {
            var v16 int = -865
            _ = v16
        }
        _ = fi12
        _ = v15
    }
    if (v10 || (true || ((false && (1 < 1)) || true))) {
        v10 = v10
    } else {
        v17 := GenS1{F0: 0}
        var v18 string = ("hello" + ("" + "\n\t"))
        _ = v17
        _ = v18
    }
    v2 += v2
    for fi19 := 0; fi19 < 5; fi19++ {
        if v10 {
            fmt.Println(v2)
        }
        var v20 int16 = 829
        _ = fi19
        _ = v20
    }
    if (2 < 128) {
        if (v4 != (v4 | 42)) {
            v21 := GenS6{F0: 1}
            v3[42] = (2 < 42)
            v10 = (!(0.1 <= (float32((float64((3.14159 / 1.5)) * (((0.0 * (1.5 / v1)) + ((0.1 / (3.14159 - v1)) + 3.14159)) / v1))) - (1.5 * -2.25))))
            _ = v21
        } else {
            if v10 {
                var v22 uint64 = (v4 << ((42) & 63))
                _ = v22
            } else {
                var v23 uint8 = 127
                v24 := []int{0, -2147483648}
                v25 := GenS4{F0: ((v23 <= (v23 - (v23 >> ((42) & 63)))) || v10), F1: 1}
                _ = v23
                _ = v24
                _ = v25
            }
        }
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        if true {
            if (!v10) {
                v3[2] = ((-128 < 83) && false)
                switch fi26 {
                case -442:
                    var v27 uint32 = 42
                    _ = v27
                case 127:
                    if v10 {
                        fmt.Println(v2, v4, v2)
                        var v28 int = fi26
                        v29 := GenS5{F0: 100}
                        _ = v28
                        _ = v29
                    }
                case -1:
                    v30 := []int{(fi26 | 42), (fi26 >> ((255) & 63)), fi26}
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 5; fi31++ {
                        var v32 uint64 = 42
                        _ = fi31
                        _ = v32
                    }
                }
                switch fi26 {
                case 9223372036854775807:
                    v33 := GenS5{F0: -1}
                    _ = v33
                case 127:
                    v2 = (v2 % ((0) | 1))
                case -1:
                    v3[42] = true
                }
            }
            fmt.Println(v4)
        }
        fmt.Println(v1, v2, v1)
        v34 := GenS7{F0: 255}
        _ = fi26
        _ = v34
    }
    var v35 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v9
}
