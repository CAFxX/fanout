package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int32, p1 string, p2 uint8) uint32 {
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 float32 = ((0.1 - 1e10) * float32(0.1))
        if false {
            p1 = p1
            p1 = p1
            switch (fi5 - -176) {
            case 236:
                var v7 int8 = 0
                _ = v7
            case 2147483647:
                v8 := []int{fi5, (^(fi5)), 0, (fi5 / ((fi5) | 1))}
                _ = v8
            }
        }
        var v9 int32 = (p0 >> ((616) & 63))
        _ = fi5
        _ = v6
        _ = v9
    }
    var v10 int8 = 24
    v11 := []int{}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 65535
}

func main() {
    var v1 uint64 = 264
    v2 := []int{128, -598}
    v3 := []uint32{}
    var v4 uint8 = 1
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(fi12, v1, v1)
        _ = fi12
    }
    if len(v2) > 0 {
        v2[0] = 128
    }
    fmt.Println(v1, v4)
    v13 := GenS1{F0: 75, F1: 704, F2: 1}
    fmt.Println(v1)
    var v14 int64 = 1
    if (0 <= 32767) {
        v1 = uint64(v4)
        v1 += v1
        var v15 int16 = 1
        _ = v15
    } else {
        if (!true) {
            var v16 uint = 4294967295
            var v17 uint8 = v4
            switch "" {
            case "foo bar":
                _ = len(v2)
            case "\n\t":
                v18 := GenS1{F0: -1, F1: 627, F2: -550}
                _ = v18
            case "hello":
                fmt.Println(v1, v16)
            default:
                v16 = v16
            }
            _ = v16
            _ = v17
        } else {
            v14 += (v14 / ((v14) | 1))
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v19 uint8 = 42
    switch 21 {
    case 959:
        v20 := GenS2{F0: v14}
        _ = v20
    case 128:
        v4 = (v4 >> ((255) & 63))
    case 32768:
        if len(v3) > 0 {
            v3[0] = 65535
        }
    default:
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 uint8 = v19
            _ = fi21
            _ = v22
        }
    }
    v4 -= (v19 * v19)
    if false {
        v3 = append(v3, 352)
        var v23 uint64 = 18446744073709551615
        _ = v23
    } else {
        switch ("a" + "a") {
        case "a":
            var v24 uint64 = v1
            _ = v24
        }
        fmt.Println(v19, v4, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
