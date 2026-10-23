package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 uint64 = 2
    v4 := make(map[int]int)
    var v5 uint16 = 42
    v5 += 1
    v6 := GenS1{F0: "hello"}
    var v7 uint32 = 482
    if false {
        v3 += (v3 & (v3 & 1))
    }
    v8 := GenS1{F0: "\n\t"}
    var v9 int16 = 32767
    v5 -= v5
    v10 := GenS1{F0: "hello"}
    var v11 float32 = float32(float64(0.1))
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(v9)
        v13 := GenS1{F0: "hello"}
        for fi14 := 0; fi14 < 20; fi14++ {
            if (201 < 127) {
                switch "\n\t" {
                case "hello":
                    v15 := GenS2{F0: true}
                    _ = v15
                case "":
                    var v16 float64 = 1.5
                    _ = v16
                }
                fmt.Println(v11, fi12)
            }
            var v17 int = fi12
            _ = fi14
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    var v18 int32 = 0
    var v19 int64 = -416
    for fi20 := 0; fi20 < 1; fi20++ {
        if (2 != 4294967295) {
            fmt.Println(fi20, fi20, v5)
            var v21 int8 = 100
            v22 := GenS2{F0: true}
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 int32 = v18
                fmt.Println(v3)
                var v25 int16 = (v9 % (((v9 + int16(v24))) | 1))
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
