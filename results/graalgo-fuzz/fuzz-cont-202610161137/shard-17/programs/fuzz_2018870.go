package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0() int32 {
    if (!(1 == 697)) {
        var v7 uint32 = 188
        fmt.Println(v7)
        _ = v7
    } else {
        v8 := []int{4294967295, -633, 32767}
        _ = v8
    }
    var v9 float64 = 3.14159
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{531, 0}
        var v12 int = fi10
        if len(v11) > 0 {
            v11[0] = (fi10 / ((fi10) | 1))
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v9
    return 1
}

func main() {
    var v1 int = 51
    v2 := make(map[uint32]string)
    var v3 int8 = -94
    v4 := make(map[uint32]uint32)
    var v5 bool = true
    v6 := make(map[uint32]bool)
    var v13 uint16 = 2
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 uint32 = 649
        var v16 uint16 = (v13 | v13)
        v17 := []int{v1, v1, v1}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    if (!v5) {
        var v18 bool = (true && v5)
        v19 := GenS1{F0: (v13 >> ((580) & 63))}
        v5 = v18
        _ = v18
        _ = v19
    } else {
        switch "foo bar" {
        case "":
            var v20 int8 = 127
            _ = v20
        default:
            v21 := GenS1{F0: v13}
            _ = v21
        }
        v13 -= v13
        switch "\n\t" {
        case "a":
            fmt.Println(v1, v13, v13)
        case "":
            var v22 uint16 = v13
            _ = v22
        }
    }
    if true {
        v1 += v1
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        v1 += (v1 >> ((18446744073709551615) & 63))
        var v24 int64 = 0
        var v25 uint8 = 42
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := []int{}
    v27 := GenS1{F0: (v13 >> ((288) & 63))}
    v5 = v5
    fmt.Println(v5, v5)
    fmt.Println(v13)
    for fi28 := 0; fi28 < 1; fi28++ {
        v29 := GenS2{F0: 1000, F1: uint8(v13)}
        _ = fi28
        _ = v29
    }
    fmt.Println(v5)
    var v30 bool = v5
    var v31 uint = 1
    var v32 string = ("a" + "")
    var v33 uint16 = 255
    if true {
        if v5 {
            switch (v1 << ((v31) & 63)) {
            case 32767:
                fmt.Println(v31, v32, v32)
            case -9223372036854775808:
                if v5 {
                    v34 := GenS1{F0: v33}
                    _ = v34
                }
            }
        } else {
            v35 := GenS2{F0: 32767, F1: 0}
            _ = v35
        }
        v36 := GenS2{F0: -32768, F1: 231}
        _ = v36
    } else {
        v3 += 127
    }
    v13 += v33
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
