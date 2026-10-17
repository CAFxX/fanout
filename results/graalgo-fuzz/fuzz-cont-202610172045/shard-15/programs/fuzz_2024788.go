package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
func genFn0() int32 {
    v4 := []int{-1, 2147483647}
    var v5 uint16 = 32767
    _ = v4
    _ = v5
    return -1
}
func genFn1() string {
    if true {
        v6 := []int{-840, 128}
        v7 := GenS1{F0: "hello", F1: 294, F2: 128}
        var v8 int32 = 1
        _ = v6
        _ = v7
        _ = v8
    } else {
        switch "hello" {
        case "hello":
            var v9 uint16 = 504
            _ = v9
        case "":
            var v10 int32 = genFn0()
            _ = v10
        case "\n\t":
            var v11 int64 = 9223372036854775807
            _ = v11
        default:
            switch "foo bar" {
            case "hello":
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 string = "\n\t"
                    _ = fi12
                    _ = v13
                }
            }
        }
        var v14 uint8 = 0
        switch "a" {
        case "foo bar":
            fmt.Println(v14)
        case "\n\t":
            fmt.Println(v14)
        case "":
            v14 -= (v14 & v14)
        default:
            v14 = v14
        }
        _ = v14
    }
    v15 := GenS1{F0: "foo bar", F1: -510, F2: 128}
    var v16 uint32 = 2147483648
    switch 0 {
    case -9223372036854775808:
        v17 := []int{42, 913, 552}
        _ = v17
    case -157:
        v16 = v16
    case -969:
        var v18 float32 = 0.1
        _ = v18
    }
    v19 := []int{2147483647, 32768, 218, 9223372036854775807}
    _ = v15
    _ = v16
    _ = v19
    return "hello"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 2
    var v3 int8 = -25
    v3 -= (v3 + (v3 | v3))
    v2 += (v2 >> ((255) & 63))
    fmt.Println(v3)
    for fi20 := 0; fi20 < 6; fi20++ {
        if (((("hello" + (("foo bar" + ("\n\t" + "\n\t")) + "a")) + ("\n\t" + "\n\t")) == "hello") || (18446744073709551615 < 42)) {
            if false {
                delete(v1, "a")
                v3 -= (v3 | (v3 - v3))
            }
            v1[("\n\t" + "")] = v2
        } else {
            fmt.Println(fi20, v3)
            var v21 int64 = 9223372036854775807
            v2 += (v2 << ((0) & 63))
            _ = v21
        }
        _ = fi20
    }
    v3 += v3
    for fi22 := 0; fi22 < 14; fi22++ {
        delete(v1, "")
        v2 += v2
        _ = fi22
    }
    v23 := []int{128, 9223372036854775807}
    switch 32768 {
    case -503:
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
    default:
        var v24 int16 = 561
        _ = v24
    }
    fmt.Println(v2, v3)
    _ = len(v1)
    var v25 float64 = 3.14159
    for fi26 := 0; fi26 < 18; fi26++ {
        switch fi26 {
        case 127:
            var v27 uint16 = 32767
            _ = v27
        case -2147483648:
            var v28 int = 489
            _ = v28
        case -770:
            v23 = append(v23, 0)
        }
        if true {
            v25 -= v25
            for fi29 := 0; fi29 < 20; fi29++ {
                v30 := v1[((("\n\t" + "a") + "") + ("hello" + ((("" + (("a" + ("" + ("foo bar" + "a"))) + "a")) + "hello") + ((("\n\t" + "a") + "foo bar") + "hello"))))]
                v31 := GenS1{F0: "hello", F1: (fi29 >> ((4294967295) & 63)), F2: 3}
                v2 = (v2 / (((^(v2))) | 1))
                _ = fi29
                _ = v30
                _ = v31
            }
        } else {
            for fi32 := 0; fi32 < 5; fi32++ {
                v33 := GenS1{F0: "foo bar", F1: 4294967295, F2: 2}
                if len(v23) > 0 {
                    v23[0] = -697
                }
                v2 = 78
                _ = fi32
                _ = v33
            }
            var v34 int = 42
            _ = v34
        }
        fmt.Println(v2, fi26)
        _ = fi26
    }
    v25 = v25
    for fi35 := 0; fi35 < 4; fi35++ {
        for fi36 := 0; fi36 < 15; fi36++ {
            fmt.Println(v3)
            v25 = (float64(1e10) + v25)
            switch (fi35 + 4294967295) {
            case 2147483647:
                delete(v1, "foo bar")
            case -440:
                _ = len(v1)
            case -2147483648:
                var v37 uint = 255
                _ = v37
            }
            _ = fi36
        }
        v25 = 0.1
        _ = fi35
    }
    fmt.Println(v2, v25, v3)
    fmt.Println(v3)
    _ = len(v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
}
