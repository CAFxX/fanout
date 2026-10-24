package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint16; F2 int16 }

func main() {
    v1 := []int{}
    var v2 float32 = (0.1 / 1.5)
    v3 := []uint32{1}
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    var v6 uint32 = 1
    fmt.Println(v2, v6)
    v3 = append(v3, 2147483647)
    for fi7 := 0; fi7 < 5; fi7++ {
        if (!(true || (65535 != 2))) {
            v8 := GenS1{F0: 1000}
            v9 := GenS2{F0: (42 == 1), F1: 42, F2: 42}
            _ = v8
            _ = v9
        } else {
            if false {
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch (fi10 >> ((118) & 63)) {
                    case 32767:
                        v6 -= v6
                    case -9223372036854775808:
                        fmt.Println(v6, v2)
                    }
                    _ = fi10
                }
                v11 := GenS2{F0: (0 >= 128), F1: 563, F2: -1}
                _ = v11
            } else {
                var v12 float32 = v2
                var v13 bool = (!true)
                _ = v12
                _ = v13
            }
            fmt.Println(fi7, fi7, v6)
        }
        var v14 string = "\n\t"
        _ = fi7
        _ = v14
    }
    var v15 int32 = 0
    var v16 uint64 = 1
    v17 := GenS2{F0: true, F1: 1, F2: -32768}
    var v18 uint8 = 0
    fmt.Println(v2, v2)
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 int32 = 42
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 int8 = -1
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    v16 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
