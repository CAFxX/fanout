package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int16 }
func genFn0(p0 uint32) int32 {
    p0 = p0
    var v5 float32 = 0.1
    if ((564 > 951) && true) {
        v6 := GenS1{F0: 950, F1: true, F2: 0}
        var v7 float32 = -2.25
        _ = v6
        _ = v7
    }
    var v8 int8 = 1
    _ = p0
    _ = v5
    _ = v8
    return -15
}
func genFn1(p0 int8, p1 uint64) uint8 {
    var v9 int8 = p0
    switch ("\n\t" + "\n\t") {
    case "":
        v9 += (-(p0))
    case "hello":
        var v10 int = 1
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint64 = (p1 * (p1 & 987))
            _ = fi11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    v1 := []string{(("foo bar" + "a") + "foo bar")}
    v2 := make(map[uint32]string)
    v3 := []uint32{65535, 880, 2, 861}
    v4 := make(map[string]int)
    var v13 uint16 = 255
    var v14 uint64 = 18446744073709551615
    v15 := GenS2{F0: 65535}
    if false {
        if len(v3) > 0 {
            v3[0] = 65535
        }
        v16 := GenS3{F0: 2147483647}
        fmt.Println(v13)
        _ = v16
    } else {
        var v17 uint16 = 32767
        fmt.Println(v14, v13, v13)
        _ = v17
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 uint16 = (v13 * 843)
        _ = len(v1)
        _ = fi18
        _ = v19
    }
    var v20 int16 = 32767
    v13 = 0
    switch -1 {
    case 1:
        v21 := v2[928]
        _ = v21
    default:
        v22 := GenS4{F0: 42, F1: v13, F2: -32768}
        _ = v22
    }
    fmt.Println(v14, v20)
    v20 += -1
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 uint16 = v13
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
