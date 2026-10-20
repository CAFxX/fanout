package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint16) uint {
    v5 := []int{-522, 32767, 9223372036854775807}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v5
    return 65535
}
func genFn1(p0 uint, p1 uint32) uint32 {
    var v6 uint16 = 350
    v6 -= 419
    var v7 uint16 = (v6 % ((v6) | 1))
    v6 += (^(v7))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 258
}

func main() {
    var v1 int16 = 1000
    var v2 int32 = 262
    v3 := []uint32{0, 387, 65535, 2}
    var v4 int = 32767
    if false {
        var v8 int = -932
        _ = v8
    } else {
        v4 = (v4 / ((v4) | 1))
        if (true || (v1 != 0)) {
            v1 = v1
            v9 := GenS1{F0: 587}
            _ = v9
        } else {
            var v10 uint16 = 0
            if len(v3) > 0 {
                v3[0] = 227
            }
            _ = v10
        }
        v3 = append(v3, 2147483647)
    }
    if (v2 != 2147483647) {
        var v11 int16 = v1
        _ = v11
    }
    fmt.Println(v1, v1)
    fmt.Println(v2, v4)
    for fi12 := 0; fi12 < 12; fi12++ {
        v1 += (^(v1))
        v2 = (v2 >> ((281) & 63))
        _ = fi12
    }
    switch v4 {
    case 4294967295:
        var v13 uint8 = 2
        _ = v13
    case 127:
        v4 = v4
    }
    switch "foo bar" {
    case "":
        v14 := GenS1{F0: 973}
        _ = v14
    case "hello":
        var v15 int8 = 42
        _ = v15
    case "foo bar":
        v4 -= (v4 >> ((42) & 63))
    default:
        _ = len(v3)
    }
    var v16 int32 = v2
    var v17 int16 = (v1 % ((-32768) | 1))
    if (255 == 32767) {
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 int = fi18
            _ = fi18
            _ = v19
        }
        v17 = v17
        var v20 float64 = 1.5
        _ = v20
    } else {
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
