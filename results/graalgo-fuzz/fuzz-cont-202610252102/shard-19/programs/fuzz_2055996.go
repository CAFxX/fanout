package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    var v3 int8 = 100
    var v4 int32 = -2147483648
    var v5 float32 = float32(3.14159)
    var v6 uint32 = 332
    fmt.Println(v1)
    var v7 uint16 = 628
    v7 += (v7 % (((v7 << ((4294967295) & 63))) | 1))
    switch "\n\t" {
    case "\n\t":
        var v8 uint64 = 0
        _ = v8
    case "hello":
        switch -2147483648 {
        case 32767:
            v9 := GenS1{F0: 2}
            _ = v9
        case 4294967295:
            v10 := v2[0]
            _ = v10
        case 0:
            var v11 string = ("a" + (("a" + "") + "hello"))
            _ = v11
        default:
            v6 = v6
        }
    default:
        var v12 bool = true
        _ = v12
    }
    if (!(false || v1)) {
        for fi13 := 0; fi13 < 13; fi13++ {
            v6 += 42
            if true {
                v1 = v1
            } else {
                v14 := []int{}
                _ = v14
            }
            if (1.5 < 1.5) {
                fmt.Println(v4, v6, v3)
                for fi15 := 0; fi15 < 9; fi15++ {
                    v6 -= (v6 - (^(v6)))
                    var v16 uint8 = 246
                    var v17 int8 = (^(v3))
                    v7 = v7
                    fmt.Println(v1)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            } else {
                v2[(v6 - 2)] = ("foo bar" + "")
            }
            _ = fi13
        }
    } else {
        fmt.Println(v6)
        var v18 int32 = -2147483648
        v7 = (v7 >> ((18446744073709551615) & 63))
        _ = v18
    }
    var v19 bool = false
    v3 = (v3 * int8(v6))
    var v20 uint = 255
    var v21 uint = v20
    v22 := GenS1{F0: 1}
    for fi23 := 0; fi23 < 20; fi23++ {
        _ = len(v2)
        fmt.Println(v6, v19, v21)
        _ = fi23
    }
    v24 := GenS1{F0: 42}
    v25 := GenS2{F0: v21}
    var v26 uint32 = (v6 ^ 42)
    fmt.Println(v1, v19)
    v27 := GenS2{F0: v21}
    v5 = v5
    v28 := GenS3{F0: (v4 | (v4 % (((v4 % ((v4) | 1))) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
