package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int }
type GenS2 struct { F0 uint32 }
func genFn0() bool {
    v6 := GenS1{F0: 218, F1: 0, F2: 2147483647}
    var v7 uint32 = 0
    if (true && (!(-32768 >= 32767))) {
        var v8 string = ("\n\t" + "")
        _ = v8
    } else {
        var v9 int16 = -1
        _ = v9
    }
    _ = v6
    _ = v7
    return (318 < 4294967295)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"foo bar", "", "", (("foo bar" + ("" + "foo bar")) + "\n\t")}
    var v3 int16 = 32767
    v4 := []int{483, -656, 2147483647}
    v5 := make(map[uint32]string)
    var v10 int32 = 1
    if ((0 != 2) && (true || (163 <= 0))) {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int16 = (v3 ^ (v3 / ((v3) | 1)))
            _ = fi11
            _ = v12
        }
        if (false || ((v3 / (((v3 - v3)) | 1)) >= v3)) {
            v13 := GenS1{F0: 127, F1: 75, F2: 128}
            v10 = v10
            _ = v13
        }
        v14 := GenS1{F0: 219, F1: 255, F2: 148}
        _ = v14
    }
    if (true && true) {
        var v15 uint32 = 42
        for fi16 := 0; fi16 < 3; fi16++ {
            v4 = append(v4, (fi16 | fi16))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi16
        }
        _ = v15
    }
    var v17 uint32 = 75
    var v18 float32 = (1.5 + float32((((0.1 + (((-2.25 + 1.5) + 1.5) - 1e10)) / 3.14159) / (0.1 - 1e10))))
    switch "a" {
    case "\n\t":
        var v19 uint64 = 42
        _ = v19
    case "":
        var v20 uint = 574
        _ = v20
    case "foo bar":
        var v21 uint16 = 2
        _ = v21
    default:
        v22 := v5[v17]
        _ = v22
    }
    var v23 int64 = 0
    if (42 != -122) {
        var v24 int8 = -83
        var v25 int64 = v23
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: 1, F1: 255, F2: 0}
    v10 = v10
    var v27 uint32 = v17
    var v28 uint16 = 72
    switch -217 {
    case -18:
        for fi29 := 0; fi29 < 19; fi29++ {
            v18 = (-2.25 - v18)
            var v30 uint32 = (v27 << ((817) & 63))
            v31 := GenS2{F0: v27}
            _ = fi29
            _ = v30
            _ = v31
        }
    case 1:
        v18 -= v18
    case 562:
        var v32 int64 = (v23 % (((v23 << ((2) & 63))) | 1))
        _ = v32
    default:
        var v33 uint16 = v28
        _ = v33
    }
    fmt.Println(v28)
    v17 -= (v17 - 42)
    v17 -= v27
    var v34 float32 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
