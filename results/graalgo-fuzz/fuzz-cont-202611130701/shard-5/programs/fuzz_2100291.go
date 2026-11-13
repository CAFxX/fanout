package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint32 }
func genFn0(p0 int64) uint32 {
    v6 := []int{}
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(p0)
        _ = fi7
    }
    _ = p0
    _ = v6
    return 648
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    v3 := []uint32{42, 4294967295}
    v4 := []uint32{115, 42}
    var v5 float32 = ((1.5 + (1e10 - 1e10)) * (((0.1 * 0.1) / 0.1) + 0.1))
    if true {
        v8 := GenS1{F0: (4294967295 != 0), F1: 18446744073709551615}
        v1 = "\n\t"
        v1 = v1
        _ = v8
    }
    switch 1 {
    case 127:
        v9 := GenS1{F0: (!(true || (32767 > 0))), F1: 629}
        _ = v9
    default:
        var v10 int64 = 264
        _ = v10
    }
    v5 = v5
    var v11 uint = 786
    v11 = v11
    var v12 uint32 = genFn0(9223372036854775807)
    for fi13 := 0; fi13 < 4; fi13++ {
        v4 = append(v4, v12)
        switch fi13 {
        case 0:
            if false {
                if (((76 == 2) || false) && (fi13 == -262)) {
                    v14 := v2[v12]
                    v1 = v1
                    if true {
                        v15 := GenS1{F0: false, F1: 445}
                        _ = v15
                    } else {
                        v16 := GenS2{F0: genFn0(-537)}
                        var v17 int32 = 58
                        _ = v16
                        _ = v17
                    }
                    _ = v14
                }
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi13
    }
    if false {
        var v18 uint = (^(v11))
        _ = v18
    }
    switch v1 {
    case "":
        var v19 int64 = -9223372036854775808
        _ = v19
    }
    _ = len(v3)
    if true {
        _ = len(v2)
        v3 = append(v3, (v12 | (v12 * v12)))
        if false {
            v1 = v1
            var v20 bool = (42 >= 4294967295)
            v21 := GenS2{F0: v12}
            _ = v20
            _ = v21
        } else {
            var v22 uint16 = 42
            _ = v22
        }
    }
    v23 := GenS3{F0: 127, F1: 18446744073709551615}
    for fi24 := 0; fi24 < 15; fi24++ {
        fmt.Println(v5, fi24)
        _ = fi24
    }
    v25 := GenS4{F0: -9223372036854775808, F1: genFn0(9223372036854775807)}
    _ = len(v2)
    if (!true) {
        var v26 bool = (2 != 529)
        _ = v26
    } else {
        var v27 bool = true
        for fi28 := 0; fi28 < 16; fi28++ {
            v1 = "foo bar"
            var v29 uint8 = 2
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    var v30 float32 = (v5 * v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
