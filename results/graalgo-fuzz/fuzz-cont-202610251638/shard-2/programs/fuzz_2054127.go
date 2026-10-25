package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    var v3 uint8 = 255
    v4 := make(map[uint32]string)
    var v5 float32 = float32(3.14159)
    if false {
        v6 := GenS1{F0: 112}
        switch 128 {
        case -9223372036854775808:
            var v7 bool = (32767 != -383)
            _ = v7
        case 205:
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS2{F0: 2147483648, F1: "hello"}
                v2[194] = ("\n\t" + "hello")
                v10 := GenS3{F0: -1, F1: 42}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = v6
    } else {
        fmt.Println(v3, v5)
        var v11 string = "hello"
        _ = v11
    }
    if true {
        v12 := v4[42]
        _ = v12
    }
    var v13 uint32 = 957
    if ((!false) && false) {
        v14 := GenS1{F0: 0}
        fmt.Println(v13)
        _ = v14
    }
    v3 = (v3 + (v3 % (((v3 * v3)) | 1)))
    var v15 uint32 = (v13 & v13)
    v3 = v3
    v13 = (v13 >> ((255) & 63))
    v5 += 3.14159
    fmt.Println(v13, v13, v15)
    v16 := GenS1{F0: 59}
    var v17 bool = (((^(v3)) > (v3 + v3)) || (955 != 1))
    v18 := GenS4{F0: 1000}
    var v19 int16 = 32767
    for fi20 := 0; fi20 < 11; fi20++ {
        v5 = (1.5 + -2.25)
        v5 -= 3.14159
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
