package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint32 = 0
    var v2 int64 = 1
    v3 := []int{4294967295, 9223372036854775807, 355, -2147483648}
    v4 := []uint32{v1, 838, (v1 - v1), 0}
    var v5 uint64 = 0
    var v6 uint16 = uint16(v1)
    fmt.Println(v2, v1, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 9223372036854775807 {
    case 42:
        var v7 int8 = 100
        _ = v7
    case -847:
        v8 := GenS1{F0: (!(!(v1 <= v1))), F1: v1}
        _ = v8
    case 127:
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 uint64 = v5
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    var v12 float32 = (3.14159 - (0.1 - (3.14159 - (0.0 * 0.0))))
    var v13 int = 743
    var v14 uint16 = v6
    var v15 uint8 = 2
    if false {
        var v16 int32 = 42
        _ = v16
    } else {
        v1 = uint32(v5)
    }
    v2 = v2
    _ = len(v4)
    if len(v4) > 0 {
        v4[0] = v1
    }
    v2 = 42
    if (int32(v15) == 1) {
        var v17 int = 0
        _ = v17
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            if (18446744073709551615 <= 87) {
                fmt.Println(fi18, v15)
                v19 := GenS2{F0: "a"}
                switch (v13 + 128) {
                case 0:
                    var v20 uint64 = v5
                    _ = v20
                case 127:
                    var v21 uint8 = 0
                    _ = v21
                }
                _ = v19
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 float32 = 0.1
                v2 = v2
                v24 := GenS3{F0: 0, F1: 42}
                _ = fi22
                _ = v23
                _ = v24
            }
            switch -233 {
            case 47:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case -114:
                v4 = append(v4, 2147483647)
            }
            _ = fi18
        }
        var v25 uint16 = 638
        _ = v25
    }
    var v26 string = "foo bar"
    var v27 int8 = -128
    v27 = (v27 | (v27 / (((^(v27))) | 1)))
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
