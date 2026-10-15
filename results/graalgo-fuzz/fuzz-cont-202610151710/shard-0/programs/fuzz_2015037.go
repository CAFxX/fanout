package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -5
    v3 := []string{("hello" + "hello"), "", ("\n\t" + "")}
    v4 := make(map[int]int)
    var v5 int16 = -32768
    switch -1 {
    case -525:
        v2 = v2
    default:
        v2 = (v2 & 42)
    }
    v6 := GenS1{F0: (v2 - v2), F1: 45}
    var v7 int32 = 325
    v7 = (v7 + (^(v7)))
    delete(v1, 4294967295)
    v7 -= v7
    var v8 int64 = -526
    var v9 int64 = (v8 / (((v8 * v8)) | 1))
    for fi10 := 0; fi10 < 7; fi10++ {
        if ((-(v2)) != int8(fi10)) {
            if (0.1 <= (1e10 + 0.1)) {
                var v11 int32 = (v7 % ((v7) | 1))
                for fi12 := 0; fi12 < 7; fi12++ {
                    v9 = (v9 & v8)
                    _ = fi12
                }
                v13 := GenS2{F0: -89, F1: 48, F2: 676}
                _ = v11
                _ = v13
            }
        }
        var v14 int = fi10
        _ = fi10
        _ = v14
    }
    v15 := GenS1{F0: v2, F1: 814}
    for fi16 := 0; fi16 < 20; fi16++ {
        v2 += v2
        _ = fi16
    }
    v7 += 0
    var v17 int16 = v5
    v9 = (v8 & (v8 + v9))
    v17 = v17
    if true {
        switch -2147483648 {
        case 4294967295:
            v18 := GenS3{F0: "a"}
            _ = v18
        default:
            switch 591 {
            case 127:
                v7 -= int32(v9)
            case -320:
                for fi19 := 0; fi19 < 1; fi19++ {
                    fmt.Println(v2)
                    v5 += (v5 % ((v5) | 1))
                    _ = fi19
                }
            }
        }
        if (!(0 > 0)) {
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 bool = false
                v8 = (v8 - v9)
                _ = len(v1)
                _ = fi20
                _ = v21
            }
            var v22 uint = 4294967295
            fmt.Println(v17, v17, v5)
            _ = v22
        } else {
            var v23 uint = 79
            _ = v23
        }
    }
    v2 += (v2 + 100)
    fmt.Println(v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
