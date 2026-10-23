package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 uint, p1 int32, p2 string) string {
    v5 := GenS1{F0: (p2 + p2), F1: 4294967295}
    if (!(-689 < 32767)) {
        p2 = "foo bar"
        p1 = p1
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        p0 += p0
        for fi7 := 0; fi7 < 19; fi7++ {
            if false {
                v8 := []int{fi6, fi7, 9223372036854775807, fi6}
                _ = v8
            }
            fmt.Println(p0, fi7)
            _ = fi7
        }
        _ = fi6
    }
    v9 := GenS1{F0: (p2 + ("\n\t" + p2)), F1: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return ""
}
func genFn1(p0 int) uint8 {
    v10 := GenS1{F0: "", F1: 4294967295}
    p0 += 1
    v11 := []int{}
    v12 := GenS1{F0: ("\n\t" + "\n\t"), F1: 263}
    if false {
        for fi13 := 0; fi13 < 15; fi13++ {
            fmt.Println(p0, fi13)
            p0 = (^(fi13))
            _ = fi13
        }
        var v14 float32 = 3.14159
        var v15 bool = (!(true || ((true && (!false)) && true)))
        _ = v14
        _ = v15
    } else {
        var v16 uint16 = 1
        switch p0 {
        case -30:
            if true {
                var v17 uint16 = 0
                _ = v17
            }
        case 42:
            var v18 uint = 537
            _ = v18
        case 851:
            p0 = p0
        default:
            if (358 == 2) {
                for fi19 := 0; fi19 < 14; fi19++ {
                    fmt.Println(v16)
                    var v20 int16 = 1000
                    v20 = (v20 >> ((uint(v16)) & 63))
                    _ = fi19
                    _ = v20
                }
                var v21 int = p0
                _ = v21
            }
        }
        _ = v16
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"a", "hello", "hello"}
    var v3 float32 = (1.5 + float32(1.5))
    var v4 uint8 = 29
    for fi22 := 0; fi22 < 13; fi22++ {
        switch "foo bar" {
        case "hello":
            fmt.Println(v4)
        default:
            var v23 uint16 = 255
            _ = v23
        }
        v24 := GenS1{F0: "foo bar", F1: 781}
        _ = fi22
        _ = v24
    }
    v1["\n\t"] = "hello"
    var v25 float32 = v3
    v25 = (1.5 - v25)
    var v26 int = 42
    switch (v26 ^ 0) {
    case -1:
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v3)
        case "hello":
            var v27 int32 = 1
            _ = v27
        case "foo bar":
            var v28 uint8 = (v4 << ((18446744073709551615) & 63))
            _ = v28
        }
    case 9223372036854775807:
        for fi29 := 0; fi29 < 19; fi29++ {
            switch "\n\t" {
            case "a":
                if true {
                    v26 -= (fi29 << ((0) & 63))
                } else {
                    var v30 float32 = ((1.5 + (float32(1e10) + (0.1 - v3))) / (0.0 * v25))
                    v25 -= (v3 + 0.1)
                    var v31 int = v26
                    _ = v30
                    _ = v31
                }
            default:
                _ = len(v1)
            }
            var v32 float64 = -2.25
            var v33 float64 = v32
            _ = fi29
            _ = v32
            _ = v33
        }
    }
    var v34 uint = 1
    for fi35 := 0; fi35 < 16; fi35++ {
        var v36 int8 = 28
        var v37 uint16 = uint16(v4)
        switch v26 {
        case 127:
            var v38 float64 = (1.5 * 0.0)
            _ = v38
        case 1:
            var v39 uint32 = uint32(v34)
            _ = v39
        case 109:
            if (true || true) {
                var v40 uint32 = 0
                _ = v40
            }
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println(v26)
    var v41 int = -116
    fmt.Println(v34)
    switch "" {
    case "hello":
        v34 = v34
    default:
        v42 := GenS2{F0: (v34 / ((v34) | 1)), F1: 42}
        _ = v42
    }
    var v43 bool = (false || false)
    fmt.Println(v26)
    var v44 float64 = 0.0
    var v45 string = ("a" + "\n\t")
    v46 := GenS3{F0: v4, F1: v41}
    var v47 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
