package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 42
    var v4 int32 = v1
    if true {
        fmt.Println(v4)
        fmt.Println(v1, v4, v1)
    }
    fmt.Println(v3, v4)
    fmt.Println(v3)
    var v5 int8 = -127
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 1, F1: 472}
        _ = fi6
        _ = v7
    }
    var v8 int32 = v4
    var v9 string = "foo bar"
    var v10 uint = 18446744073709551615
    if (!(42 != 9223372036854775807)) {
        for fi11 := 0; fi11 < 16; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                if false {
                    var v13 uint = (v10 & (v10 - v10))
                    v3 += uint32(v13)
                    v2[(v3 >> ((1) & 63))] = 229
                    v9 = v9
                    _ = v13
                } else {
                    fmt.Println(v10, v1)
                    var v14 int16 = 32767
                    fmt.Println(v1, fi12, v14)
                    _ = v14
                }
                _ = fi12
            }
            v15 := []int{480, -2147483648, fi11, 4294967295}
            _ = fi11
            _ = v15
        }
        var v16 string = v9
        _ = v16
    } else {
        v17 := GenS2{F0: v10, F1: v9}
        _ = v17
    }
    if true {
        v18 := GenS2{F0: (v10 | v10), F1: ((v9 + "a") + ((v9 + "\n\t") + ((v9 + "foo bar") + "\n\t")))}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS3{F0: true, F1: 2147483647}
            if (v5 != v5) {
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 uint32 = (v3 >> (((v10 - 0)) & 63))
                    v23 := GenS1{F0: (v5 + (v5 >> ((255) & 63))), F1: 1000}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                var v24 int64 = 662
                _ = len(v2)
                _ = v24
            } else {
                v3 = v3
            }
            _ = fi19
            _ = v20
        }
        v3 = v3
        var v25 float32 = 3.14159
        _ = v25
    }
    v26 := []int{149, -635}
    v8 -= -408
    if false {
        _ = len(v2)
    }
    if (!(float64(0.1) > 3.14159)) {
        _ = len(v2)
        v5 = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
