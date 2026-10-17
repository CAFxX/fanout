package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 uint16, p2 int) int8 {
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(p2)
        _ = fi5
    }
    p1 = 108
    _ = p0
    _ = p1
    _ = p2
    return 66
}
func genFn1(p0 int16) int16 {
    v6 := GenS1{F0: 9223372036854775807}
    v7 := GenS1{F0: -68}
    _ = p0
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 uint = 0
    var v2 uint8 = 222
    var v3 uint32 = 1
    var v4 uint8 = (v2 << (((v1 / ((v1) | 1))) & 63))
    var v8 string = ("hello" + "a")
    fmt.Println(v2, v4)
    var v9 string = ("hello" + v8)
    var v10 int16 = 55
    v11 := []int{144, 0, -154, -914}
    var v12 uint16 = 255
    v12 = (v12 | v12)
    var v13 uint32 = 2
    v14 := GenS2{F0: (v10 + int16(v4)), F1: 65535}
    if ((0 == 2147483647) && true) {
        var v15 uint16 = v12
        if (((0.1 + 3.14159) * (1e10 / float32(3.14159))) == 3.14159) {
            v9 = v8
            v3 -= (v13 + v3)
        } else {
            var v16 int32 = 1
            _ = v16
        }
        if (v4 >= 127) {
            v1 = (v1 & (v1 ^ (^(v1))))
        }
        _ = v15
    }
    if ((v8 + ("hello" + v9)) >= (((v8 + v9) + "hello") + "hello")) {
        var v17 int16 = v10
        var v18 int = 32767
        _ = v17
        _ = v18
    }
    v12 -= (v12 | 250)
    v9 = "\n\t"
    v13 -= (v13 * v3)
    var v19 uint16 = v12
    var v20 int = 731
    switch v20 {
    case -1:
        var v21 int32 = 1
        _ = v21
    default:
        if ((-74 == 100) && (!false)) {
            var v22 string = (v8 + v9)
            var v23 int = (v20 * v20)
            _ = v22
            _ = v23
        } else {
            if ((v13 / ((v3) | 1)) < (^(v3))) {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
