package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint; F2 int32 }
type GenS3 struct { F0 int64 }
func genFn0() int {
    var v6 int64 = 0
    if false {
        var v7 uint64 = 311
        var v8 int = -2147483648
        v9 := []int{v8, (v8 << ((279) & 63)), v8}
        _ = v7
        _ = v8
        _ = v9
    }
    _ = v6
    return 42
}

func main() {
    var v1 uint16 = 255
    var v2 uint32 = 42
    var v3 int8 = 78
    v4 := []uint32{(v2 ^ v2), v2, v2}
    v5 := make(map[uint32]bool)
    if false {
        var v10 uint16 = 2
        _ = v10
    } else {
        v11 := GenS1{F0: 1000, F1: 2147483647}
        var v12 int = 971
        v3 += v3
        _ = v11
        _ = v12
    }
    if (false && false) {
        var v13 uint64 = 18446744073709551615
        var v14 int = -86
        _ = v13
        _ = v14
    }
    var v15 int = 0
    switch genFn0() {
    case 127:
        for fi16 := 0; fi16 < 20; fi16++ {
            v5[(^(v2))] = (!((^(v1)) == v1))
            switch ((("" + "a") + "") + "foo bar") {
            case "foo bar":
                v17 := GenS1{F0: 0, F1: (v2 >> ((570) & 63))}
                _ = v17
            case "hello":
                v18 := GenS1{F0: -32768, F1: v2}
                _ = v18
            case "":
                fmt.Println(fi16)
            default:
                _ = len(v5)
            }
            _ = fi16
        }
    default:
        v19 := GenS1{F0: 1, F1: v2}
        _ = v19
    }
    _ = len(v4)
    var v20 uint8 = 255
    if (!((-(v15)) > (-(v15)))) {
        fmt.Println(v20, v20)
    } else {
        fmt.Println(v15, v3, v3)
        v1 += v1
        fmt.Println(v2)
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        if false {
            var v22 int64 = -1
            v15 = (fi21 * fi21)
            _ = v22
        } else {
            var v23 uint64 = 4294967295
            var v24 uint64 = (v23 * v23)
            if true {
                for fi25 := 0; fi25 < 13; fi25++ {
                    v26 := GenS2{F0: uint32(v20), F1: 18446744073709551615, F2: 42}
                    v20 = (v20 | v20)
                    var v27 float32 = -2.25
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
            _ = v23
            _ = v24
        }
        switch fi21 {
        case 32768:
            var v28 int = (v15 << ((1) & 63))
            _ = v28
        case 2147483647:
            v29 := GenS3{F0: -256}
            _ = v29
        case 9223372036854775807:
            var v30 int = -1
            _ = v30
        default:
            for fi31 := 0; fi31 < 5; fi31++ {
                fmt.Println(v2, v20, v15)
                _ = fi31
            }
        }
        _ = fi21
    }
    var v32 int16 = -32768
    var v33 uint16 = v1
    v33 -= v1
    switch "a" {
    case "foo bar":
        for fi34 := 0; fi34 < 7; fi34++ {
            for fi35 := 0; fi35 < 9; fi35++ {
                if len(v4) > 0 {
                    v4[0] = v2
                }
                var v36 uint = 1
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    }
    v5[(v2 * 42)] = true
    for fi37 := 0; fi37 < 5; fi37++ {
        var v38 bool = (false || true)
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
