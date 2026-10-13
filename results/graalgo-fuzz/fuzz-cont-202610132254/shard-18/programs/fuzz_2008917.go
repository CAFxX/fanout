package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint8) int {
    if (1 == (p0 - (p0 & (p0 + 472)))) {
        var v7 string = ("hello" + "")
        switch 128 {
        case 32768:
            var v8 string = v7
            _ = v8
        case 42:
            p0 = (p0 + (^(p0)))
        case 1:
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 uint8 = (^(p1))
                var v11 uint16 = 65535
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = v7
    }
    var v12 uint8 = 42
    p1 = p1
    fmt.Println(p1, p0, v12)
    _ = p0
    _ = p1
    _ = v12
    return -497
}

func main() {
    var v1 int64 = -724
    v2 := []int{-1, -2147483648, -114, -925}
    var v3 uint64 = 4294967295
    var v4 float32 = 3.14159
    var v5 int = 1
    var v6 int64 = (v1 | v1)
    if false {
        fmt.Println(v5, v5)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v13 float32 = v4
    var v14 uint32 = 2
    _ = len(v2)
    v14 = 1
    if false {
        v15 := GenS1{F0: v14, F1: v6, F2: -41}
        v16 := GenS1{F0: v14, F1: (v1 % ((v6) | 1)), F2: 42}
        v2 = append(v2, (v5 / ((genFn0(v3, 127)) | 1)))
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := GenS2{F0: 42, F1: uint8(v14)}
            _ = fi17
            _ = v18
        }
        v19 := GenS3{F0: (v3 * (v3 % ((v3) | 1)))}
        _ = v19
    }
    var v20 string = ("\n\t" + "")
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
