package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint64; F1 int16 }

func main() {
    var v1 float64 = (3.14159 * 1.5)
    var v2 uint16 = 2
    var v3 uint16 = v2
    var v4 int16 = -1
    v5 := []int{}
    var v6 uint64 = 18446744073709551615
    var v7 uint16 = (v2 & v2)
    for fi8 := 0; fi8 < 5; fi8++ {
        v2 -= (v3 & 540)
        _ = fi8
    }
    _ = len(v5)
    var v9 uint16 = v7
    var v10 int64 = -203
    if (-2147483648 < -1) {
        v2 -= uint16(v6)
        v11 := GenS1{F0: 4294967295}
        _ = v11
    } else {
        var v12 string = "foo bar"
        _ = v12
    }
    if ((3.14159 * 3.14159) == 3.14159) {
        v13 := GenS1{F0: 94}
        _ = len(v5)
        _ = v13
    }
    var v14 int8 = int8(v10)
    for fi15 := 0; fi15 < 1; fi15++ {
        if true {
            for fi16 := 0; fi16 < 19; fi16++ {
                switch (fi15 & -1) {
                case 128:
                    if (2 > 1) {
                        var v17 bool = (0 >= 2)
                        var v18 uint16 = v3
                        _ = v17
                        _ = v18
                    }
                default:
                    v14 += v14
                }
                v19 := GenS2{F0: 2147483647, F1: 86}
                if true {
                    v4 -= v4
                    var v20 int = fi15
                    fmt.Println(v10, fi16, fi15)
                    _ = v20
                }
                _ = fi16
                _ = v19
            }
        }
        v14 = v14
        _ = fi15
    }
    v1 = v1
    if true {
        switch "" {
        case "":
            if true {
                fmt.Println(v2)
            }
        case "hello":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "a":
            for fi21 := 0; fi21 < 11; fi21++ {
                _ = len(v5)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi21
            }
        default:
            var v22 uint8 = uint8(v6)
            _ = v22
        }
        if len(v5) > 0 {
            v5[0] = -543
        }
        v23 := GenS3{F0: (v6 ^ v6), F1: (^(v4))}
        _ = v23
    }
    v6 -= 0
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
