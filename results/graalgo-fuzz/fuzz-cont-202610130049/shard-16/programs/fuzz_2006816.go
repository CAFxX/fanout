package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int16 = 1000
    var v2 uint8 = 128
    v3 := make(map[int]string)
    var v4 uint8 = 147
    v4 += v2
    v5 := GenS1{F0: 4294967295, F1: 32768}
    v6 := GenS2{F0: int16(v4), F1: 4294967295, F2: 4294967295}
    v7 := []int{1}
    var v8 int16 = (v1 + (v1 << ((1) & 63)))
    var v9 int32 = 42
    var v10 uint32 = 2147483648
    fmt.Println(v2, v9)
    fmt.Println(v4, v2)
    if true {
        v2 += v4
        v11 := GenS3{F0: (v4 >> ((65535) & 63))}
        _ = v11
    }
    if (!(!(664 > 4294967295))) {
        v12 := GenS1{F0: (v10 + 42), F1: -9223372036854775808}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v13 uint64 = 938
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: int16(v2), F1: 42, F2: v10}
    var v15 uint16 = 255
    if (!(0 == 127)) {
        for fi16 := 0; fi16 < 11; fi16++ {
            v2 = (v4 & 255)
            var v17 bool = true
            _ = fi16
            _ = v17
        }
        var v18 int32 = (-(v9))
        if false {
            v2 = 255
            var v19 bool = (!(788 < 32767))
            v20 := GenS3{F0: 2}
            _ = v19
            _ = v20
        }
        _ = v18
    }
    switch "" {
    case "a":
        v1 = v8
    case "":
        v9 -= (^(v9))
    case "\n\t":
        if len(v7) > 0 {
            v7[0] = 32768
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
