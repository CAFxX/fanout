package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int64, p1 uint64, p2 int8) string {
    var v6 bool = (1000 != 42)
    fmt.Println(p0, v6)
    v6 = (918 >= 0)
    v7 := GenS1{F0: p1}
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ("" + "a")
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 2
    var v3 int32 = 2147483647
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 string = "hello"
        _ = fi8
        _ = v9
    }
    if true {
        if true {
            v3 = (v3 / ((v3) | 1))
        } else {
            v10 := []int{1, -406}
            v11 := GenS2{F0: 731}
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint8 = 215
                _ = fi12
                _ = v13
            }
            _ = v10
            _ = v11
        }
        if false {
            v14 := GenS3{F0: 255}
            if true {
                v15 := []int{1, -572, 32768}
                var v16 int16 = -542
                _ = v15
                _ = v16
            } else {
                fmt.Println(v3, v3)
                v2 -= v2
                for fi17 := 0; fi17 < 12; fi17++ {
                    v2 = (v2 + v2)
                    _ = fi17
                }
            }
            _ = v14
        }
    }
    var v18 float32 = 1.5
    v3 = (v3 >> ((65535) & 63))
    v19 := GenS1{F0: 627}
    v20 := GenS2{F0: -9223372036854775808}
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 float32 = 0.1
        var v23 uint32 = (v2 ^ v2)
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 int32 = v3
    v25 := GenS2{F0: -1}
    var v26 uint = 241
    var v27 string = "\n\t"
    var v28 int = -257
    var v29 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
