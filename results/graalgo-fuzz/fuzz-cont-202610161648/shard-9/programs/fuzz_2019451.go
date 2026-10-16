package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int16; F2 string }
type GenS5 struct { F0 int32; F1 string; F2 int64 }
func genFn0(p0 int, p1 int64) uint8 {
    v5 := []int{-702, (^(p0)), p0}
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: (p0 >> ((71) & 63)), F1: 2, F2: (0.0 <= ((0.1 + 1.5) / 1.5))}
        if false {
            var v8 uint8 = 128
            _ = v8
        } else {
            if false {
                _ = len(v5)
            }
        }
        _ = fi6
        _ = v7
    }
    v9 := GenS1{F0: (p0 + p0), F1: 42, F2: (1e10 <= 0.0)}
    if (true || false) {
        v10 := GenS1{F0: p0, F1: 0, F2: false}
        v11 := GenS1{F0: p0, F1: 65535, F2: (-2147483648 >= p0)}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 0
}
func genFn1(p0 uint32, p1 uint) int16 {
    if false {
        var v12 bool = (p0 != 2147483647)
        v13 := GenS2{F0: v12}
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 9223372036854775807, F1: 0, F2: false}
        v15 := GenS1{F0: 97, F1: 32767, F2: (32768 == 32767)}
        fmt.Println(p0)
        _ = v14
        _ = v15
    }
    var v16 uint32 = p0
    v17 := []int{int(p1), 4294967295, 4294967295}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 1000
}

func main() {
    var v1 uint32 = 0
    var v2 uint32 = v1
    v3 := make(map[string]int)
    var v4 uint64 = 632
    for fi18 := 0; fi18 < 2; fi18++ {
        v19 := GenS3{F0: 18446744073709551615}
        for fi20 := 0; fi20 < 16; fi20++ {
            v21 := []int{}
            var v22 uint64 = 879
            var v23 float32 = (3.14159 / (1e10 + (0.1 / 1e10)))
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        var v24 int64 = 134
        _ = fi18
        _ = v19
        _ = v24
    }
    v4 -= v4
    v1 += 695
    fmt.Println(v2, v4)
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 uint64 = 4294967295
        if false {
            v27 := []int{fi25, 32768}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 18; fi28++ {
                v29 := []int{}
                _ = len(v3)
                _ = fi28
                _ = v29
            }
            var v30 float64 = (((((float64(-2.25) - ((0.1 + 1.5) + float64(3.14159))) - 0.1) * 1e10) / ((0.1 - 0.0) * 1.5)) / float64(-2.25))
            v31 := []int{(fi25 % ((fi25) | 1))}
            _ = v30
            _ = v31
        }
        _ = fi25
        _ = v26
    }
    var v32 bool = true
    if v32 {
        for fi33 := 0; fi33 < 8; fi33++ {
            v2 = v1
            _ = fi33
        }
        v34 := GenS3{F0: uint(v2)}
        var v35 uint = uint(v4)
        _ = v34
        _ = v35
    } else {
        var v36 uint64 = 306
        v36 = (v36 ^ v4)
        _ = v36
    }
    fmt.Println(v2, v2)
    for fi37 := 0; fi37 < 14; fi37++ {
        v38 := GenS4{F0: -32768, F1: -78, F2: ""}
        fmt.Println(fi37, v2, v2)
        switch "" {
        case "hello":
            for fi39 := 0; fi39 < 16; fi39++ {
                v40 := GenS1{F0: (-(fi39)), F1: 566, F2: v32}
                v4 = 2
                v1 = 824
                _ = fi39
                _ = v40
            }
        case "foo bar":
            v32 = (fi37 < 42)
        }
        _ = fi37
        _ = v38
    }
    if true {
        if (((v2 == v2) && false) || v32) {
            v2 += v2
        }
    }
    var v41 int64 = -536
    var v42 uint8 = 2
    if (v42 > v42) {
        if v32 {
            v43 := GenS5{F0: -1, F1: ("" + "a"), F2: int64(v42)}
            _ = v43
        }
        v44 := []int{-612, int(v4), 127}
        _ = v44
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        var v46 uint64 = v4
        var v47 int = fi45
        _ = fi45
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v41
    _ = v42
}
