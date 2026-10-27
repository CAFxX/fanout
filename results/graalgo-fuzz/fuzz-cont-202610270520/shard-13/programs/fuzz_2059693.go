package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int64; F2 uint8 }
func genFn0(p0 int64, p1 uint16, p2 int32) uint64 {
    var v7 int8 = 102
    var v8 uint16 = 32767
    var v9 string = "foo bar"
    var v10 bool = (4294967295 < -261)
    p1 += 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() uint {
    var v11 int = 463
    if false {
        v11 = (v11 * v11)
        switch "foo bar" {
        case "hello":
            v12 := GenS1{F0: "hello"}
            _ = v12
        case "foo bar":
            var v13 uint8 = 255
            _ = v13
        }
        v11 = (-(v11))
    }
    _ = v11
    return 2
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]uint32)
    var v3 int8 = 1
    v4 := make(map[int]uint32)
    var v5 uint8 = 0
    var v6 uint = 721
    if (false && false) {
        v14 := []int{32768, 93}
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := GenS2{F0: int(v5), F1: 65535}
            _ = fi15
            _ = v16
        }
        v17 := GenS3{F0: -251}
        _ = v14
        _ = v17
    }
    var v18 string = ("foo bar" + "foo bar")
    switch v18 {
    case "":
        v5 = 172
    case "a":
        var v19 float32 = 3.14159
        _ = v19
    case "foo bar":
        if ((0.0 > 3.14159) || ((1 == 42) && false)) {
            v20 := GenS4{F0: false, F1: 42, F2: uint8(v6)}
            v21 := GenS4{F0: (false && ((v1 * v1) != v1)), F1: 9, F2: 101}
            for fi22 := 0; fi22 < 16; fi22++ {
                v23 := GenS4{F0: true, F1: 819, F2: v5}
                _ = fi22
                _ = v23
            }
            _ = v20
            _ = v21
        }
    }
    fmt.Println(v18)
    v3 = (v3 / (((^(v3))) | 1))
    var v24 int = 4294967295
    for fi25 := 0; fi25 < 14; fi25++ {
        v18 = "foo bar"
        var v26 uint16 = 255
        _ = fi25
        _ = v26
    }
    var v27 uint64 = genFn0(685, 1, 1)
    var v28 uint = v6
    var v29 float64 = -2.25
    for fi30 := 0; fi30 < 4; fi30++ {
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 uint64 = v27
            v33 := v4[1]
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
