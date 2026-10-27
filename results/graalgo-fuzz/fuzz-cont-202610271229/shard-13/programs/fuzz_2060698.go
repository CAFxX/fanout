package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 string }
func genFn0(p0 int32) int32 {
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := GenS1{F0: 821, F1: 2147483647}
        if false {
            p0 += p0
            p0 += -1
            switch fi4 {
            case -9223372036854775808:
                for fi6 := 0; fi6 < 7; fi6++ {
                    p0 += p0
                    v7 := []int{}
                    p0 -= p0
                    _ = fi6
                    _ = v7
                }
            default:
                v8 := GenS1{F0: 255, F1: 2}
                _ = v8
            }
        }
        var v9 uint8 = 125
        _ = fi4
        _ = v5
        _ = v9
    }
    var v10 string = "a"
    _ = p0
    _ = v10
    return (p0 | (p0 / (((p0 / ((p0) | 1))) | 1)))
}
func genFn1(p0 uint16) int64 {
    var v11 int32 = 2147483647
    fmt.Println(p0, v11)
    var v12 int64 = 1
    _ = p0
    _ = v11
    _ = v12
    return v12
}

func main() {
    var v1 float32 = 0.1
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]bool)
    v3[2147483648] = true
    if (677 >= 637) {
        fmt.Println(v1, v1)
        _ = len(v3)
    } else {
        v2 -= v2
    }
    switch 127 {
    case 928:
        var v13 uint32 = 190
        _ = v13
    case -232:
        var v14 bool = false
        _ = v14
    }
    switch 42 {
    case 42:
        var v15 float64 = -2.25
        _ = v15
    case 128:
        for fi16 := 0; fi16 < 17; fi16++ {
            v1 = (v1 * v1)
            _ = fi16
        }
    case 2147483647:
        fmt.Println(v2)
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int32 = -1
        v18 += (^(v18))
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 uint64 = v2
            v18 = v18
            switch "foo bar" {
            case "hello":
                v2 += 1
                v21 := GenS1{F0: 500, F1: 42}
                fmt.Println(v1, fi19)
                _ = v21
            case "\n\t":
                fmt.Println(v18, v20, v1)
            default:
                v22 := GenS2{F0: 4294967295, F1: v18}
                _ = v22
            }
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    v23 := GenS3{F0: 1, F1: 0, F2: 766}
    for fi24 := 0; fi24 < 4; fi24++ {
        _ = len(v3)
        _ = fi24
    }
    v25 := GenS2{F0: 1, F1: -656}
    var v26 bool = false
    if (!((((v1 - (((v1 / v1) * v1) - v1)) < 0.0) && (!v26)) && ((v2 - 0) <= (v2 * (v2 << ((855) & 63)))))) {
        var v27 int = 0
        v26 = v26
        _ = v27
    } else {
        var v28 uint64 = 18446744073709551615
        if (2147483647 == int32(v2)) {
            fmt.Println(v2, v28)
            var v29 uint32 = 4294967295
            v30 := GenS4{F0: 520}
            _ = v29
            _ = v30
        }
        _ = v28
    }
    v31 := []int{32767, 42}
    fmt.Println(v1)
    v32 := v3[0]
    v26 = (!(!(2 >= v2)))
    for fi33 := 0; fi33 < 14; fi33++ {
        for fi34 := 0; fi34 < 16; fi34++ {
            switch ("a" + (("" + "\n\t") + "")) {
            case "foo bar":
                var v35 int = fi33
                _ = v35
            case "hello":
                var v36 uint16 = 2
                _ = v36
            case "":
                var v37 int32 = -57
                _ = v37
            }
            v38 := GenS2{F0: 2147483647, F1: genFn0(-596)}
            _ = fi34
            _ = v38
        }
        fmt.Println(v2)
        v39 := GenS5{F0: 18446744073709551615}
        _ = fi33
        _ = v39
    }
    v31 = append(v31, -2147483648)
    if (genFn0(42) >= genFn0(513)) {
        v40 := v3[42]
        if len(v31) > 0 {
            v31[0] = 32767
        }
        v41 := GenS6{F0: 2, F1: (("" + "a") + (((("" + "foo bar") + ("\n\t" + "foo bar")) + "\n\t") + (("hello" + ("a" + "a")) + "a")))}
        _ = v40
        _ = v41
    } else {
        var v42 float32 = ((3.14159 - v1) + v1)
        switch 32768 {
        case 4294967295:
            v31 = append(v31, -736)
        default:
            v43 := GenS1{F0: 307, F1: 1}
            _ = v43
        }
        v42 += (1e10 + v1)
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
}
