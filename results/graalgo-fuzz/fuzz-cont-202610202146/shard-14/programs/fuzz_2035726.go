package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
func genFn0() int16 {
    if false {
        var v7 int16 = -32768
        fmt.Println(v7)
        var v8 float32 = ((3.14159 * 1.5) + (1.5 / ((((float32(0.1) / 1.0) - 0.1) - 3.14159) * float32((1.5 + -2.25)))))
        _ = v7
        _ = v8
    } else {
        var v9 uint32 = 1
        fmt.Println(v9)
        v9 = (v9 << ((1) & 63))
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint64 = 775
        _ = fi10
        _ = v11
    }
    var v12 bool = (255 < 255)
    var v13 bool = v12
    _ = v12
    _ = v13
    return 32767
}
func genFn1(p0 uint64, p1 bool, p2 uint32) uint32 {
    fmt.Println(p1, p2)
    v14 := GenS1{F0: -982}
    var v15 uint8 = 57
    v16 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return p2
}

func main() {
    v1 := []int{4294967295, -1}
    v2 := []int{103}
    var v3 int8 = -77
    var v4 uint16 = 255
    var v5 uint16 = v4
    v6 := []uint32{4294967295, 2147483647}
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            v5 -= (v5 / ((v4) | 1))
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v19 uint64 = 1
            var v20 int64 = 0
            if (977 == -475) {
                var v21 int32 = 451
                v20 -= 0
                _ = v21
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    switch "a" {
    case "":
        fmt.Println(v3)
    default:
        var v22 string = "a"
        _ = v22
    }
    if (false && (32767 > 1)) {
        fmt.Println(v4, v3, v5)
        var v23 string = "\n\t"
        _ = v23
    } else {
        fmt.Println(v5, v4)
        fmt.Println(v3)
    }
    fmt.Println(v4, v3)
    if len(v2) > 0 {
        v2[0] = 404
    }
    v3 = (v3 ^ (v3 / (((v3 / (((v3 * (v3 / ((-128) | 1)))) | 1))) | 1)))
    fmt.Println(v4)
    var v24 int = -468
    fmt.Println(v24, v5, v24)
    for fi25 := 0; fi25 < 12; fi25++ {
        fmt.Println(v3, v4, v5)
        if (!(uint64(v5) < 2)) {
            var v26 int8 = v3
            var v27 float32 = 1e10
            v27 = 0.0
            _ = v26
            _ = v27
        } else {
            v6 = append(v6, 42)
        }
        _ = fi25
    }
    var v28 int64 = 42
    v29 := GenS2{F0: (v24 % ((128) | 1))}
    var v30 uint32 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
