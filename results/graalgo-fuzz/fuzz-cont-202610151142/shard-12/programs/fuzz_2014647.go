package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int16; F1 int64; F2 uint16 }
type GenS3 struct { F0 int64; F1 bool; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 string; F1 uint8; F2 int64 }
func genFn0(p0 int32, p1 uint8) uint8 {
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            if true {
                v8 := GenS1{F0: 622, F1: p1}
                fmt.Println(fi6, fi7, p1)
                _ = v8
            }
            var v9 string = ("a" + "\n\t")
            _ = fi7
            _ = v9
        }
        var v10 int8 = -1
        p0 -= p0
        _ = fi6
        _ = v10
    }
    v11 := GenS2{F0: -875, F1: -673, F2: 568}
    _ = p0
    _ = p1
    _ = v11
    return p1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = -2.25
    var v3 bool = true
    var v4 uint16 = 776
    v5 := []uint32{2147483648, 308, 65535}
    _ = len(v5)
    var v12 uint64 = 914
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v15 uint64 = 1
            var v16 uint8 = 128
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    var v17 bool = true
    for fi18 := 0; fi18 < 1; fi18++ {
        if (!true) {
            var v19 int = (fi18 | fi18)
            v12 -= (v12 >> ((0) & 63))
            _ = v19
        } else {
            v20 := GenS3{F0: -1, F1: (v3 && false), F2: 2147483648}
            var v21 string = (("" + (("\n\t" + "\n\t") + "\n\t")) + ("hello" + "a"))
            var v22 uint64 = (v12 - v12)
            _ = v20
            _ = v21
            _ = v22
        }
        fmt.Println(v12)
        _ = fi18
    }
    v1[2147483647] = 4294967295
    v4 -= (v4 - 232)
    switch "hello" {
    case "foo bar":
        v2 = -2.25
    default:
        v12 = 0
    }
    switch -2147483648 {
    case 146:
        v23 := v1[2147483647]
        _ = v23
    case 128:
        v4 = (v4 & v4)
    case 1:
        v3 = (v17 || (v17 || false))
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v12 = (v12 - v12)
    v24 := GenS1{F0: 65535, F1: 0}
    v25 := GenS4{F0: (true || (v17 && (v2 <= 0.1))), F1: 979}
    var v26 int32 = 0
    switch 9223372036854775807 {
    case -776:
        v27 := GenS5{F0: "\n\t", F1: 255, F2: -9223372036854775808}
        _ = v27
    case 127:
        v1[2147483648] = 710
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
