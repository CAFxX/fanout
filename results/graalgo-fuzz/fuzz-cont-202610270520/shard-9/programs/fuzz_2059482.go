package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 int64, p1 uint64, p2 uint8) int16 {
    switch -2147483648 {
    case 4294967295:
        if ((0.0 <= float64(float32((0.1 * 0.0)))) && false) {
            if ((p1 & p1) <= p1) {
                switch "hello" {
                case "foo bar":
                    var v5 int8 = 127
                    _ = v5
                }
                var v6 int8 = -1
                var v7 int64 = (p0 / (((p0 | (p0 << ((65535) & 63)))) | 1))
                _ = v6
                _ = v7
            }
            p1 = (p1 | (p1 + (p1 / (((p1 + 394)) | 1))))
        } else {
            p2 = p2
            if false {
                if true {
                    v8 := []int{4294967295, -366}
                    var v9 uint64 = (^(p1))
                    _ = v8
                    _ = v9
                } else {
                    v10 := GenS1{F0: 1, F1: (p1 * (p1 - (^(p1)))), F2: 815}
                    p1 = 0
                    for fi11 := 0; fi11 < 12; fi11++ {
                        var v12 uint32 = 65535
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                }
                var v13 uint8 = uint8(p1)
                _ = v13
            } else {
                p2 -= (p2 * (p2 - 1))
                p0 = (p0 - p0)
            }
            p2 += p2
        }
    default:
        v14 := GenS2{F0: ""}
        _ = v14
    }
    p0 = -1
    _ = p0
    _ = p1
    _ = p2
    return int16(p1)
}

func main() {
    var v1 string = "hello"
    var v2 float64 = (1.5 - ((0.1 * -2.25) * 0.1))
    var v3 float64 = v2
    v4 := make(map[string]string)
    if (v2 < v2) {
        fmt.Println(v1)
    }
    v2 = float64((1e10 - 0.1))
    v15 := GenS3{F0: "hello", F1: v1}
    var v16 int32 = 0
    v2 += (v3 * (((v3 / 3.14159) - 0.0) + v3))
    v16 += (v16 * 42)
    v17 := []int{-688}
    var v18 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
