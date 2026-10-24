package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 bool }
func genFn0(p0 int32) string {
    if true {
        for fi4 := 0; fi4 < 18; fi4++ {
            v5 := []int{}
            fmt.Println(fi4, p0)
            _ = fi4
            _ = v5
        }
    } else {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := []int{}
            fmt.Println(p0, fi6)
            p0 -= (p0 % ((p0) | 1))
            _ = fi6
            _ = v7
        }
    }
    var v8 int64 = -1
    _ = p0
    _ = v8
    return ("a" + "\n\t")
}
func genFn1(p0 int8) int {
    if ((!true) && (!(!true))) {
        p0 = p0
    }
    if false {
        v9 := []int{9223372036854775807}
        _ = v9
    }
    var v10 string = "a"
    v11 := GenS1{F0: 255, F1: -1}
    v12 := GenS1{F0: 1, F1: 992}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return -2147483648
}

func main() {
    var v1 float64 = ((float64(float32((1e10 + (float64((1.5 / 1.5)) + float64(-2.25))))) - 1e10) - -2.25)
    v2 := []string{"a", "a", "hello", (("\n\t" + ("\n\t" + "")) + "a")}
    var v3 string = (("\n\t" + "") + (("" + "\n\t") + "\n\t"))
    var v13 uint = 425
    switch 256 {
    case -957:
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int = genFn1(100)
            for fi16 := 0; fi16 < 16; fi16++ {
                fmt.Println(v3)
                v17 := GenS1{F0: 65535, F1: 32767}
                var v18 int = (fi14 << (((v13 >> (((v13 + v13)) & 63))) & 63))
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 uint64 = 2
            _ = fi14
            _ = v15
            _ = v19
        }
    case 42:
        var v20 uint8 = 123
        _ = v20
    default:
        switch 9223372036854775807 {
        case 1:
            var v21 float32 = (1.5 - float32(v1))
            _ = v21
        case 127:
            var v22 float64 = (v1 / v1)
            _ = v22
        }
    }
    v1 += v1
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 uint = (v13 / ((v13) | 1))
        fmt.Println(fi23)
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 int64 = -9223372036854775808
            fmt.Println(v3)
            var v27 uint64 = 2
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    var v28 uint32 = 2
    for fi29 := 0; fi29 < 11; fi29++ {
        v3 = ("foo bar" + v3)
        switch (v3 + v3) {
        case "\n\t":
            fmt.Println(v13, v1)
        default:
            fmt.Println(fi29, fi29)
        }
        _ = fi29
    }
    switch v3 {
    case "\n\t":
        v30 := GenS2{F0: 391, F1: 113, F2: uint32(v13)}
        _ = v30
    case "hello":
        v13 -= v13
    }
    fmt.Println(v3)
    var v31 uint64 = 994
    switch "foo bar" {
    case "\n\t":
        switch 197 {
        case 1:
            switch "\n\t" {
            case "hello":
                v28 = (v28 % ((v28) | 1))
            case "":
                if true {
                    v2 = append(v2, v3)
                    v28 = 2
                    v28 += (v28 ^ v28)
                }
            case "a":
                for fi32 := 0; fi32 < 16; fi32++ {
                    v1 = (v1 - 3.14159)
                    var v33 uint = 583
                    var v34 string = v3
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            default:
                for fi35 := 0; fi35 < 12; fi35++ {
                    var v36 float32 = ((1e10 - 1e10) + ((3.14159 + (1.5 / (3.14159 / 3.14159))) + 1.5))
                    _ = fi35
                    _ = v36
                }
            }
        case -743:
            v37 := GenS2{F0: 9223372036854775807, F1: 2, F2: v28}
            _ = v37
        }
    default:
        var v38 int64 = -479
        _ = v38
    }
    if true {
        for fi39 := 0; fi39 < 1; fi39++ {
            v31 += (v31 & v31)
            v3 = (v3 + "\n\t")
            _ = fi39
        }
        var v40 float32 = 0.0
        _ = v40
    } else {
        if false {
            v3 = "\n\t"
        } else {
            v3 = "\n\t"
            fmt.Println(v1, v3, v3)
            switch v3 {
            case "foo bar":
                v31 = v31
            case "a":
                v31 -= v31
            case "":
                v41 := GenS3{F0: 1, F1: v28, F2: true}
                _ = v41
            default:
                fmt.Println(v31, v31)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v31
}
