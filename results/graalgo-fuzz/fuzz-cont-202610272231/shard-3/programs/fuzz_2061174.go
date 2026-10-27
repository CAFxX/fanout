package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int32 }
type GenS2 struct { F0 int8; F1 int8; F2 string }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint16) int8 {
    v6 := GenS1{F0: -762, F1: false, F2: -720}
    v7 := GenS1{F0: 1000, F1: false, F2: (p0 * 1)}
    fmt.Println(p1, p1)
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float64 = (0.0 + float64(1e10))
        fmt.Println(v9)
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int, p1 int16, p2 uint32) uint64 {
    v11 := []int{int(p2), (p0 + (p0 % ((p0) | 1)))}
    var v12 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 861
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 333
    v3 := []uint32{42}
    var v4 int = -290
    v5 := []int{v4, (v4 & v4), v4, (^(v4))}
    if false {
        var v13 uint32 = 2
        if len(v5) > 0 {
            v5[0] = v4
        }
        _ = v13
    }
    if (768 > 69) {
        v4 = v4
    } else {
        v14 := GenS2{F0: genFn0(-806, 32767), F1: 0, F2: "hello"}
        _ = v14
    }
    switch "foo bar" {
    case "":
        switch 127 {
        case -9223372036854775808:
            fmt.Println(v4, v2)
        }
    case "foo bar":
        fmt.Println(v4)
    }
    v15 := GenS3{F0: 28, F1: 42}
    if true {
        if (false || (2147483647 >= 2147483647)) {
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 int16 = (^(v2))
                _ = fi16
                _ = v17
            }
            v2 = v2
            v4 = v4
        } else {
            v4 = v4
            fmt.Println(v2)
        }
    } else {
        var v18 string = ((("\n\t" + "foo bar") + "\n\t") + "hello")
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int16 = 32767
        switch (v4 - (v4 - (^(v4)))) {
        case -9223372036854775808:
            v5 = append(v5, fi19)
        case 32768:
            v21 := GenS4{F0: 42}
            _ = v21
        default:
            v22 := GenS3{F0: 399, F1: 4294967295}
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    delete(v1, (v4 & 4294967295))
    fmt.Println(v2, v4)
    switch ("hello" + (("\n\t" + ("foo bar" + ("a" + "\n\t"))) + "hello")) {
    case "hello":
        v2 += (v2 * 32767)
    case "":
        var v23 int16 = v2
        _ = v23
    }
    fmt.Println(v2, v2)
    fmt.Println(v4)
    v24 := GenS1{F0: 32767, F1: false, F2: 799}
    for fi25 := 0; fi25 < 17; fi25++ {
        if false {
            v26 := GenS4{F0: 0}
            var v27 int32 = 2147483647
            v4 = (^(v4))
            _ = v26
            _ = v27
        } else {
            fmt.Println(fi25)
        }
        if true {
            fmt.Println(v2, v4, fi25)
            v2 = v2
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 float64 = (3.14159 * ((float64((float32(-2.25) * 3.14159)) + 3.14159) / float64(-2.25)))
                _ = fi28
                _ = v29
            }
        }
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
