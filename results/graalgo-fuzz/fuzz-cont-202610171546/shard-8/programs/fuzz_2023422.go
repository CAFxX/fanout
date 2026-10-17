package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint) int16 {
    if false {
        for fi5 := 0; fi5 < 17; fi5++ {
            fmt.Println(fi5, p0)
            p0 += p0
            fmt.Println(p0, p1)
            _ = fi5
        }
        var v6 uint64 = 305
        var v7 int16 = -393
        _ = v6
        _ = v7
    } else {
        var v8 int8 = 0
        _ = v8
    }
    p1 = (p1 << (((p1 % ((p1) | 1))) & 63))
    fmt.Println(p1, p1)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return -32768
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 1
    var v3 float64 = 3.14159
    v4 := make(map[string]uint32)
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int16 = -32768
            var v12 int64 = 42
            _ = fi10
            _ = v11
            _ = v12
        }
        if (-528 <= 42) {
            var v13 int8 = 100
            _ = v13
        }
        v14 := GenS1{F0: 2, F1: 128}
        _ = fi9
        _ = v14
    }
    v15 := GenS2{F0: 0, F1: 1000, F2: 927}
    var v16 uint = 1
    for fi17 := 0; fi17 < 2; fi17++ {
        switch "foo bar" {
        case "a":
            fmt.Println(v16)
        }
        var v18 int16 = 1000
        v18 += (v18 | v18)
        _ = fi17
        _ = v18
    }
    v16 = v16
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 float32 = (3.14159 + float32(v3))
        _ = fi19
        _ = v20
    }
    v16 = v16
    v2 = (v2 | (v2 * v2))
    var v21 float32 = 0.1
    if false {
        delete(v4, ("hello" + ("" + "\n\t")))
        fmt.Println(v3, v3)
    } else {
        var v22 bool = (false || false)
        _ = v22
    }
    v23 := GenS3{F0: 1}
    var v24 uint32 = uint32(v16)
    for fi25 := 0; fi25 < 2; fi25++ {
        v3 += v3
        if (!(42 == 767)) {
            _ = len(v4)
        }
        v3 -= 1.5
        _ = fi25
    }
    v26 := GenS2{F0: 2147483647, F1: -32768, F2: v24}
    v27 := []int{343}
    var v28 float32 = v21
    for fi29 := 0; fi29 < 3; fi29++ {
        switch fi29 {
        case -1:
            var v30 uint = v16
            _ = v30
        case 9223372036854775807:
            v28 = 0.0
        default:
            fmt.Println(v28)
        }
        switch "\n\t" {
        case "":
            v3 = (v3 * 1.5)
        case "foo bar":
            var v31 uint64 = 4294967295
            _ = v31
        case "\n\t":
            v3 = v3
        default:
            v21 = 1.5
        }
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
