package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint16 }
func genFn0(p0 int32) string {
    var v7 uint = 255
    for fi8 := 0; fi8 < 4; fi8++ {
        if false {
            v9 := GenS1{F0: "", F1: v7, F2: 1}
            _ = v9
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return ("" + "hello")
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    var v3 uint = 255
    v4 := []string{}
    v5 := make(map[int]uint32)
    v6 := []uint32{65535, 65535}
    v3 += v3
    var v10 uint16 = 0
    for fi11 := 0; fi11 < 16; fi11++ {
        v10 = v10
        _ = fi11
    }
    if ((1.5 == (0.0 * (0.1 - 1e10))) && false) {
        var v12 uint8 = 142
        v3 += (v3 & 595)
        var v13 string = "\n\t"
        _ = v12
        _ = v13
    }
    if false {
        _ = len(v5)
    }
    var v14 float32 = 0.1
    for fi15 := 0; fi15 < 18; fi15++ {
        switch "\n\t" {
        case "":
            v1[-93] = "a"
        default:
            v2 = append(v2, "a")
        }
        v10 += (v10 >> (((v3 & v3)) & 63))
        _ = fi15
    }
    _ = len(v5)
    switch 4294967295 {
    case 802:
        v14 += ((float32(0.0) - (3.14159 - 1e10)) / 3.14159)
    case 127:
        _ = len(v6)
    default:
        v5[-570] = 42
    }
    if len(v6) > 0 {
        v6[0] = 65535
    }
    v1[-385] = ""
    var v16 int64 = 9223372036854775807
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(fi17, v16, v14)
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
