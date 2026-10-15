package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint }
func genFn0(p0 bool, p1 uint16) bool {
    v7 := []int{2147483647, 4294967295}
    var v8 uint16 = 82
    if (!false) {
        switch "hello" {
        case "\n\t":
            var v9 uint = 42
            _ = v9
        case "":
            var v10 uint8 = 42
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (((4294967295 > 73) && false) || p0)
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[uint32]string)
    var v3 uint32 = 2
    var v4 int32 = 42
    var v5 int8 = 0
    v6 := []int{32768, -734, 4294967295}
    v5 += 0
    for fi11 := 0; fi11 < 9; fi11++ {
        v3 = v3
        v12 := GenS1{F0: "", F1: (v4 * int32(v3))}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi11
        _ = v12
    }
    v4 = v4
    v5 = -1
    v4 += (v4 & v4)
    v4 -= 42
    if false {
        if true {
            var v13 uint32 = (v3 ^ v3)
            var v14 uint16 = 451
            if len(v6) > 0 {
                v6[0] = 128
            }
            _ = v13
            _ = v14
        }
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    }
    if (!(false || (!true))) {
        v15 := GenS1{F0: "\n\t", F1: (v4 << ((4294967295) & 63))}
        _ = v15
    }
    _ = len(v1)
    v5 = v5
    v16 := GenS2{F0: 4294967295, F1: -931, F2: 309}
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v5, fi17, v3)
        v18 := GenS3{F0: (v3 * 1), F1: (v3 & v3)}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint = 65535
        _ = fi19
        _ = v20
    }
    v21 := GenS4{F0: -2147483648, F1: (v5 >> ((2) & 63))}
    if true {
        for fi22 := 0; fi22 < 17; fi22++ {
            delete(v2, 4294967295)
            _ = fi22
        }
        v23 := GenS5{F0: 127, F1: 30, F2: 495}
        _ = v23
    } else {
        var v24 int16 = -1
        _ = v24
    }
    var v25 float32 = 1e10
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS1{F0: ((("\n\t" + "") + "\n\t") + ""), F1: -1}
        var v28 float64 = float64(1.5)
        v4 = 0
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
