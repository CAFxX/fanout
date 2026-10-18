package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 1
    var v3 float32 = (float32(((1.5 / float64((1.5 - 3.14159))) / 1e10)) / (-2.25 + 1.5))
    var v4 bool = true
    fmt.Println(v3, v3)
    switch "" {
    case "\n\t":
        v2 -= (v2 * v2)
    case "hello":
        v5 := []int{}
        _ = v5
    }
    var v6 string = "foo bar"
    if (v4 || v4) {
        var v7 uint8 = 0
        v8 := []int{411, 32768}
        _ = v7
        _ = v8
    }
    var v9 uint64 = 4294967295
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if (-1 != 32767) {
                v12 := GenS1{F0: v9, F1: 2}
                fmt.Println(fi11, fi10)
                v13 := GenS2{F0: -32768, F1: (v2 % (((v2 & v2)) | 1)), F2: 42}
                _ = v12
                _ = v13
            }
            _ = len(v1)
            var v14 int32 = 0
            _ = fi11
            _ = v14
        }
        _ = fi10
    }
    if (485 < 524) {
        if v4 {
            var v15 uint16 = 0
            fmt.Println(v2, v2, v3)
            _ = v15
        }
        v16 := GenS3{F0: (!v4)}
        v3 = v3
        _ = v16
    } else {
        v17 := GenS3{F0: true}
        v18 := GenS3{F0: true}
        var v19 int16 = -751
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println(v4)
    fmt.Println(v6, v3)
    v20 := GenS1{F0: v9, F1: 32767}
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := GenS3{F0: (!false)}
        _ = fi21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
