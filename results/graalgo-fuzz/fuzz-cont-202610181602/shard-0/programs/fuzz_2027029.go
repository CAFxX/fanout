package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
func genFn0(p0 uint32) string {
    v7 := GenS1{F0: 2, F1: 63, F2: 2}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int32 = 688
        fmt.Println(v9, p0, v9)
        v10 := []int{4294967295}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = "\n\t"
    v12 := GenS1{F0: 42, F1: 32767, F2: 22}
    _ = p0
    _ = v11
    _ = v12
    _ = v7
    return "a"
}
func genFn1() bool {
    var v13 uint32 = 820
    if true {
        var v14 uint32 = (v13 + (v13 + 2))
        _ = v14
    }
    if true {
        if false {
            if true {
                v13 += (v13 / (((v13 << ((2) & 63))) | 1))
            }
            var v15 uint32 = (v13 + v13)
            _ = v15
        }
        fmt.Println(v13)
    }
    _ = v13
    return true
}

func main() {
    v1 := []int{32767}
    var v2 bool = (!(-1 > 164))
    var v3 int64 = 0
    v4 := make(map[string]uint32)
    var v5 uint16 = 42
    v6 := make(map[uint32]uint32)
    switch 127 {
    case 4294967295:
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint16 = v5
            if (v2 && (2 != 128)) {
                if ((!false) || (true && (!v2))) {
                    _ = len(v4)
                    v2 = v2
                    var v18 string = ("hello" + "")
                    _ = v18
                }
                var v19 uint = 4294967295
                _ = v19
            } else {
                var v20 int64 = (v3 & v3)
                _ = v20
            }
            var v21 int32 = 140
            _ = fi16
            _ = v17
            _ = v21
        }
    case -173:
        var v22 int16 = -32768
        _ = v22
    default:
        switch "" {
        case "":
            fmt.Println(v2, v3)
        case "hello":
            var v23 int64 = v3
            _ = v23
        case "a":
            v24 := GenS1{F0: 909, F1: v5, F2: 669}
            _ = v24
        }
    }
    if false {
        var v25 int = 1
        v2 = v2
        v26 := GenS1{F0: 0, F1: (v5 - (^(v5))), F2: (v5 << ((282) & 63))}
        _ = v25
        _ = v26
    }
    var v27 uint = 0
    v3 -= int64(v27)
    for fi28 := 0; fi28 < 2; fi28++ {
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 int16 = 0
            _ = fi29
            _ = v30
        }
        v3 -= (v3 ^ (v3 >> ((v27) & 63)))
        _ = fi28
    }
    v31 := GenS1{F0: 4294967295, F1: 823, F2: 42}
    for fi32 := 0; fi32 < 4; fi32++ {
        for fi33 := 0; fi33 < 2; fi33++ {
            if len(v1) > 0 {
                v1[0] = 128
            }
            v1 = append(v1, (^(fi32)))
            _ = fi33
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi32
    }
    v34 := GenS1{F0: 1, F1: (v5 ^ v5), F2: v5}
    delete(v6, 4294967295)
    if v2 {
        var v35 int16 = -407
        v36 := GenS1{F0: 0, F1: v5, F2: 255}
        v27 += v27
        _ = v35
        _ = v36
    }
    var v37 int = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
