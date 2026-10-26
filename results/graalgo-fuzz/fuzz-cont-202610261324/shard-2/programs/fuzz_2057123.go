package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
func genFn0(p0 int) int32 {
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 uint64 = 2
            v9 := []int{32767, -2147483648, (p0 >> ((4294967295) & 63)), (p0 / (((fi7 | -550)) | 1))}
            fmt.Println(v8)
            _ = fi7
            _ = v8
            _ = v9
        }
        switch 42 {
        case -2147483648:
            fmt.Println(fi6, p0)
        case 204:
            var v10 int64 = 0
            _ = v10
        }
        p0 = (p0 ^ fi6)
        _ = fi6
    }
    var v11 uint8 = 0
    _ = p0
    _ = v11
    return 1
}
func genFn1(p0 bool, p1 bool) uint32 {
    var v12 int16 = 593
    v13 := []int{602, -2147483648}
    for fi14 := 0; fi14 < 11; fi14++ {
        _ = len(v13)
        fmt.Println(fi14, p1, p0)
        _ = len(v13)
        _ = fi14
    }
    v15 := GenS1{F0: "\n\t", F1: 695, F2: 246}
    for fi16 := 0; fi16 < 13; fi16++ {
        switch ("" + ("hello" + "")) {
        case "":
            p1 = false
        case "\n\t":
            v17 := GenS1{F0: (("a" + "foo bar") + "foo bar"), F1: 1, F2: 255}
            _ = v17
        case "hello":
            var v18 bool = (197 > 127)
            _ = v18
        default:
            if true {
                _ = len(v13)
            } else {
                switch fi16 {
                case 127:
                    p0 = p0
                case -535:
                    var v19 int = fi16
                    _ = v19
                }
            }
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v15
    return 2
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]uint32)
    v3 := []int{4294967295, 32767, -2147483648}
    var v4 int = 4294967295
    v5 := []uint32{}
    var v20 int32 = 0
    if (42 <= 1) {
        if (129 == v1) {
            v21 := v2[(("hello" + "a") + ("" + ""))]
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 string = ""
                _ = fi22
                _ = v23
            }
            _ = v21
        } else {
            v4 = -9223372036854775808
        }
    }
    v24 := GenS2{F0: (^(v1))}
    var v25 string = "a"
    var v26 int = (v4 - (^(v4)))
    for fi27 := 0; fi27 < 2; fi27++ {
        v1 = 117
        v28 := GenS3{F0: 32767}
        _ = fi27
        _ = v28
    }
    v1 = (v1 << ((65535) & 63))
    v20 = (v20 >> ((255) & 63))
    if (42 > -9223372036854775808) {
        fmt.Println(v1, v25)
    } else {
        fmt.Println(v20, v1, v20)
        var v29 int8 = 85
        var v30 int64 = 42
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 int32 = v20
            v1 = (^(v1))
            _ = fi32
            _ = v33
        }
        v20 -= (v20 / ((1) | 1))
        _ = fi31
    }
    if true {
        var v34 int16 = 0
        switch "" {
        case "\n\t":
            switch "" {
            case "hello":
                v26 = v4
            case "\n\t":
                v35 := GenS4{F0: v26}
                _ = v35
            }
        case "a":
            v36 := GenS1{F0: (v25 + v25), F1: 18446744073709551615, F2: (v1 / ((128) | 1))}
            _ = v36
        case "foo bar":
            v2[(v25 + ("\n\t" + v25))] = genFn1(false, true)
        }
        v34 = 1
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
