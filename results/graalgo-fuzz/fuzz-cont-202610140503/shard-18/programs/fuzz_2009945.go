package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int16 }
func genFn0(p0 int32, p1 int, p2 bool) uint {
    var v7 int16 = 32767
    switch p1 {
    case 527:
        v8 := GenS1{F0: 1}
        _ = v8
    case 127:
        if (float32(0.0) >= float32((-2.25 * 1e10))) {
            if ((v7 << ((2) & 63)) != v7) {
                var v9 int8 = 0
                fmt.Println(p1, p0)
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 16; fi10++ {
                fmt.Println(p1, fi10)
                _ = fi10
            }
            var v11 uint8 = 42
            _ = v11
        }
    default:
        var v12 int8 = -1
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    var v1 uint64 = 2
    v2 := []string{"foo bar", "", (("" + "a") + ("" + "a"))}
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 uint64 = (v1 >> ((44) & 63))
    var v6 int16 = 1000
    var v13 uint16 = 1
    _ = len(v4)
    var v14 uint32 = 2147483648
    v15 := GenS2{F0: 208}
    fmt.Println(v14, v14)
    v4["hello"] = 0
    v16 := GenS3{F0: 1}
    var v17 uint16 = 255
    if (v6 == v6) {
        var v18 float64 = 0.0
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS2{F0: 127}
            _ = fi19
            _ = v20
        }
        var v21 float64 = 0.0
        _ = v18
        _ = v21
    }
    var v22 int8 = -20
    if true {
        switch "" {
        case "hello":
            var v23 uint8 = 42
            _ = v23
        }
        v5 = v5
        fmt.Println(v14, v13, v5)
    } else {
        v13 = v13
        v22 += v22
        var v24 uint16 = v17
        _ = v24
    }
    switch 288 {
    case 2147483647:
        if (((^(v1)) == 646) && true) {
            var v25 int16 = (v6 >> ((genFn0(-2147483648, 128, false)) & 63))
            var v26 uint8 = 127
            fmt.Println(v17)
            _ = v25
            _ = v26
        }
    case -464:
        if true {
            v27 := GenS4{F0: (467 == (v1 + v5)), F1: 42}
            _ = v27
        }
    case 32768:
        if len(v2) > 0 {
            v2[0] = ("hello" + "\n\t")
        }
    default:
        v6 = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
