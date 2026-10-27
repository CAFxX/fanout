package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 uint8) uint8 {
    p0 -= p0
    if false {
        p0 -= (p0 / (((p0 % ((p0) | 1))) | 1))
    }
    v6 := []int{-22}
    _ = p0
    _ = v6
    return p0
}

func main() {
    var v1 uint16 = 65535
    var v2 bool = false
    v3 := make(map[int]string)
    v4 := []int{-888, -9223372036854775808, -9223372036854775808, 748}
    v5 := make(map[int]bool)
    fmt.Println(v1, v1)
    var v7 int = -2147483648
    var v8 string = ("foo bar" + "a")
    var v9 uint32 = 725
    _ = len(v5)
    v9 = v9
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v10 int = v7
    v8 = "hello"
    if v2 {
        v11 := GenS1{F0: 155}
        _ = v11
    } else {
        v12 := GenS1{F0: 44}
        var v13 float64 = 0.1
        if (-34 >= 0) {
            for fi14 := 0; fi14 < 6; fi14++ {
                v9 += (^(v9))
                _ = fi14
            }
            for fi15 := 0; fi15 < 3; fi15++ {
                if true {
                    v8 = v8
                    var v16 int64 = 9223372036854775807
                    _ = v16
                } else {
                    v17 := GenS2{F0: false, F1: 1000}
                    v18 := GenS2{F0: (!(744 >= 1)), F1: 1}
                    _ = v17
                    _ = v18
                }
                v19 := GenS3{F0: v7, F1: 128}
                _ = fi15
                _ = v19
            }
            var v20 float64 = (v13 / -2.25)
            _ = v20
        }
        _ = v12
        _ = v13
    }
    var v21 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
