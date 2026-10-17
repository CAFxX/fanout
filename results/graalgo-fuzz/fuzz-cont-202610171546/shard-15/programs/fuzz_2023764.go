package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 string; F1 int8 }
func genFn0(p0 uint, p1 int16) bool {
    v6 := GenS1{F0: "\n\t", F1: (p1 | -1), F2: 144}
    if (true && true) {
        p1 -= p1
        var v7 int16 = p1
        p1 = (p1 + -914)
        _ = v7
    } else {
        var v8 string = "foo bar"
        p1 = -32768
        var v9 int64 = -9223372036854775808
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: ("\n\t" + ("hello" + "\n\t")), F1: 2}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return (!false)
}

func main() {
    var v1 uint32 = 2
    v2 := []uint32{v1}
    var v3 string = ""
    var v4 uint64 = 25
    var v5 uint32 = v1
    if len(v2) > 0 {
        v2[0] = (v5 * (v1 * v1))
    }
    var v11 uint8 = 2
    var v12 uint = 4294967295
    if len(v2) > 0 {
        v2[0] = (v1 >> ((uint(v4)) & 63))
    }
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    v1 = v1
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS3{F0: (v3 + v3), F1: 16}
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint = 211
            if len(v2) > 0 {
                v2[0] = v1
            }
            v11 += v11
            v1 += (v1 - (v1 % (((v5 % ((v5) | 1))) | 1)))
            _ = fi15
            _ = v16
        }
        var v17 float64 = (0.0 * 1e10)
        _ = fi13
        _ = v14
        _ = v17
    }
    var v18 string = v3
    fmt.Println(v18)
    var v19 uint16 = 1
    fmt.Println(v11)
    v5 = uint32(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
