package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int16) uint64 {
    p0 += (p0 << ((42) & 63))
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int32 = -1
        var v6 float64 = (0.1 - (1.5 - 0.1))
        _ = fi4
        _ = v5
        _ = v6
    }
    _ = p0
    return 512
}
func genFn1() int64 {
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int32 = -28
            var v9 uint8 = 216
            _ = fi7
            _ = v8
            _ = v9
        }
        if false {
            v10 := GenS1{F0: 9223372036854775807}
            var v11 uint64 = 0
            _ = v10
            _ = v11
        } else {
            var v12 uint8 = 1
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int16 = 42
                _ = fi13
                _ = v14
            }
            _ = v12
        }
    } else {
        var v15 uint32 = 65535
        var v16 uint16 = 32767
        _ = v15
        _ = v16
    }
    var v17 int32 = -1
    v17 = v17
    var v18 int16 = 1000
    v19 := []int{0, 960}
    _ = v17
    _ = v18
    _ = v19
    return 1
}

func main() {
    var v1 int32 = 42
    var v2 int64 = 9223372036854775807
    var v3 string = (("a" + (("a" + "hello") + "a")) + "")
    switch -9223372036854775808 {
    case 32767:
        var v20 uint32 = 0
        _ = v20
    case 475:
        v3 = "a"
    case 977:
        v3 = v3
    default:
        v2 += v2
    }
    var v21 uint32 = 486
    switch "foo bar" {
    case "":
        v3 = (v3 + ("a" + v3))
    case "a":
        var v22 int8 = 100
        _ = v22
    case "foo bar":
        var v23 int8 = 127
        _ = v23
    }
    var v24 uint64 = 547
    var v25 int16 = 215
    var v26 int16 = 594
    var v27 int = 1
    if (!false) {
        var v28 int = (^(v27))
        _ = v28
    } else {
        v3 = v3
        fmt.Println(v26)
    }
    v25 -= 1000
    switch "\n\t" {
    case "hello":
        var v29 uint8 = 127
        _ = v29
    case "a":
        v1 += (v1 * v1)
    case "foo bar":
        for fi30 := 0; fi30 < 9; fi30++ {
            var v31 int64 = (-(v2))
            var v32 int16 = 0
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    var v33 uint8 = 145
    v34 := []int{600, v27}
    if (!false) {
        for fi35 := 0; fi35 < 14; fi35++ {
            _ = len(v34)
            _ = fi35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
}
