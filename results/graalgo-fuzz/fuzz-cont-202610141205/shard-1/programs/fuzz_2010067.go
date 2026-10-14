package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 int64; F1 int8; F2 int32 }
func genFn0() int8 {
    var v7 int = 0
    fmt.Println(v7)
    _ = v7
    return 1
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]int)
    v3 := []int{128, 1, -2147483648}
    v4 := make(map[uint32]bool)
    v5 := []int{32768, 2147483647, 563}
    var v6 int32 = -1
    v8 := GenS1{F0: ((774 > uint32(v1)) && (-1 > -550)), F1: 2, F2: 128}
    var v9 int16 = -323
    if true {
        switch "hello" {
        case "a":
            v1 += 1
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            v10 := GenS1{F0: true, F1: 32767, F2: -9223372036854775808}
            _ = v10
        }
        v3 = append(v3, 4294967295)
    } else {
        fmt.Println(v1)
        v6 += (v6 | -2147483648)
    }
    v9 = v9
    var v11 int32 = (v6 * (v6 ^ v6))
    v9 -= -738
    v11 -= 37
    if (false || false) {
        var v12 int32 = (v11 + (v6 + (v11 / (((v11 * v11)) | 1))))
        fmt.Println(v11, v6)
        _ = v12
    }
    var v13 int16 = v9
    fmt.Println(v13, v1)
    v9 = v9
    var v14 float32 = 0.1
    v15 := GenS2{F0: -159, F1: -1, F2: (v6 ^ v11)}
    var v16 int8 = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
