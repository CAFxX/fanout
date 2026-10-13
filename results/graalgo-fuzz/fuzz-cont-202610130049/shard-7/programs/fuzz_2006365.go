package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 int; F2 bool }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int16, p1 int) int16 {
    if (2 < 125) {
        var v5 int = p1
        _ = v5
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int32 = 1
            fmt.Println(v8, p1)
            switch "" {
            case "\n\t":
                switch "foo bar" {
                case "a":
                    switch fi6 {
                    case 0:
                        var v9 int32 = v8
                        fmt.Println(v9)
                        v9 -= (v8 ^ (v9 >> ((2) & 63)))
                        _ = v9
                    case 739:
                        p1 = fi6
                    default:
                        if false {
                            if true {
                                p0 = (p0 ^ (p0 + p0))
                            }
                        }
                    }
                case "hello":
                    v10 := []int{(-(fi6)), fi7, 32768}
                    _ = v10
                }
            case "hello":
                var v11 uint16 = 110
                _ = v11
            case "":
                v12 := []int{fi7, (fi7 - fi7)}
                _ = v12
            }
            _ = fi7
            _ = v8
        }
        fmt.Println(p0, p1, fi6)
        switch 42 {
        case 563:
            v13 := GenS1{F0: (fi6 ^ p1), F1: 1}
            _ = v13
        case -569:
            var v14 string = "\n\t"
            _ = v14
        case -9223372036854775808:
            p1 = -2147483648
        default:
            switch 127 {
            case 128:
                for fi15 := 0; fi15 < 4; fi15++ {
                    switch ((("hello" + "") + "\n\t") + "a") {
                    case "":
                        fmt.Println(fi6)
                    case "hello":
                        switch (p1 << ((600) & 63)) {
                        case 42:
                            v16 := []int{9223372036854775807}
                            _ = v16
                        case 127:
                            v17 := []int{p1, (-(fi15)), 128, fi15}
                            _ = v17
                        }
                    default:
                        v18 := GenS2{F0: 326, F1: true}
                        _ = v18
                    }
                    var v19 uint32 = 1
                    _ = fi15
                    _ = v19
                }
            case -553:
                var v20 int = (fi6 % (((p1 >> ((961) & 63))) | 1))
                _ = v20
            case 127:
                p1 -= (p1 + 32767)
            }
        }
        _ = fi6
    }
    if true {
        p0 -= p0
    }
    v21 := GenS3{F0: 42}
    _ = p0
    _ = p1
    _ = v21
    return (p0 | p0)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    if (!true) {
        if (false && false) {
            var v22 float64 = float64((0.0 - 0.0))
            switch -9223372036854775808 {
            case 42:
                v23 := GenS2{F0: 2, F1: (!true)}
                _ = v23
            case 205:
                fmt.Println(v22)
            }
            v24 := GenS4{F0: 1}
            _ = v22
            _ = v24
        }
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := GenS5{F0: -123, F1: -152, F2: (!false)}
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 3; fi27++ {
            v3[fi27] = 2
            fmt.Println(fi27)
            _ = fi27
        }
    }
    v28 := GenS6{F0: 4294967295}
    var v29 uint = 65535
    fmt.Println(v29)
    v29 += v29
    v30 := v1[1]
    var v31 uint16 = 32767
    v29 += v29
    fmt.Println(v29, v29)
    delete(v3, 165)
    v29 = 18446744073709551615
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 bool = (!true)
        _ = fi32
        _ = v33
    }
    v34 := GenS4{F0: 65535}
    if true {
        for fi35 := 0; fi35 < 3; fi35++ {
            var v36 uint32 = 956
            _ = fi35
            _ = v36
        }
        v37 := []int{128, 1}
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
}
