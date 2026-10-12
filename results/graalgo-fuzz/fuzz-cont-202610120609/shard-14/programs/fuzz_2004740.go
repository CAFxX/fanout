package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int; F2 uint }
type GenS3 struct { F0 uint16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int32 = 1
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int = fi5
            fmt.Println(fi5, fi5, v8)
            switch fi7 {
            case 1:
                var v9 string = ""
                _ = v9
            case -2147483648:
                var v10 uint32 = 2147483648
                var v11 uint32 = v10
                v12 := GenS1{F0: uint(v11)}
                v13 := GenS2{F0: v6, F1: fi7, F2: 42}
                _ = v10
                _ = v11
                _ = v12
                _ = v13
            case -1:
                if true {
                    fmt.Println(v6)
                    if true {
                        var v14 bool = false
                        v6 += (v6 / ((42) | 1))
                        _ = v14
                    }
                    v15 := GenS1{F0: 255}
                    _ = v15
                }
            default:
                var v16 uint16 = 65535
                _ = v16
            }
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 string = (("a" + "foo bar") + "\n\t")
        var v19 bool = (v18 == "\n\t")
        _ = fi17
        _ = v18
        _ = v19
    }
    switch -9223372036854775808 {
    case -2147483648:
        var v20 uint = 2
        _ = v20
    case 4294967295:
        switch 127 {
        case 1:
            var v21 float64 = (((((1e10 + -2.25) + float64((0.1 - 1e10))) / float64((3.14159 / 0.1))) + 1.5) / 1.0)
            _ = v21
        case 32768:
            var v22 int16 = -8
            _ = v22
        case 127:
            if true {
                var v23 uint = 18446744073709551615
                v24 := []int{-2147483648, int(v23)}
                switch "" {
                case "":
                    fmt.Println(v23)
                case "\n\t":
                    v23 = 1
                case "foo bar":
                    var v25 uint32 = 1
                    _ = v25
                }
                _ = v23
                _ = v24
            }
        default:
            v26 := []int{32768, -9223372036854775808}
            _ = v26
        }
    }
    v27 := []int{}
    _ = v27
    return -403
}
func genFn1(p0 uint8, p1 int, p2 uint8) int16 {
    if true {
        v28 := []int{}
        v29 := GenS1{F0: 4294967295}
        _ = len(v28)
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        var v31 uint16 = 2
        _ = fi30
        _ = v31
    }
    _ = p0
    _ = p1
    _ = p2
    return genFn0()
}

func main() {
    v1 := []string{("a" + ""), ("hello" + "hello")}
    v2 := make(map[string]bool)
    var v3 uint32 = 181
    v4 := make(map[uint32]bool)
    var v32 int8 = 127
    var v33 float64 = (3.14159 / 0.1)
    switch 0 {
    case 4294967295:
        var v34 uint16 = 42
        _ = v34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        var v36 uint = 2
        var v37 uint64 = 4294967295
        _ = fi35
        _ = v36
        _ = v37
    }
    v32 += (v32 / (((v32 % (((v32 << ((18446744073709551615) & 63))) | 1))) | 1))
    switch (("foo bar" + "a") + "hello") {
    case "a":
        v38 := GenS2{F0: 42, F1: 519, F2: 0}
        _ = v38
    case "":
        fmt.Println(v3, v32)
    case "foo bar":
        if true {
            v33 -= 1e10
        } else {
            var v39 uint16 = 2
            _ = v39
        }
    }
    var v40 float64 = v33
    var v41 uint16 = 165
    v3 += (v3 / ((v3) | 1))
    _ = len(v4)
    var v42 int64 = 198
    var v43 int16 = -1
    var v44 uint32 = v3
    for fi45 := 0; fi45 < 7; fi45++ {
        fmt.Println(v43)
        v4[(v44 * v3)] = true
        var v46 uint8 = 2
        _ = fi45
        _ = v46
    }
    switch 42 {
    case -36:
        v40 -= float64(1.5)
    case -1:
        v44 = v44
    }
    if (!false) {
        v47 := GenS3{F0: (^(v41))}
        _ = v47
    } else {
        var v48 uint64 = 4294967295
        _ = len(v4)
        _ = v48
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
