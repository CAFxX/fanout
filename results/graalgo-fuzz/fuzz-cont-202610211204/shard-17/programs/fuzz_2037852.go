package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 string; F2 int64 }
func genFn0(p0 int8, p1 int, p2 uint32) uint {
    v6 := GenS1{F0: "a", F1: false}
    var v7 uint64 = 742
    fmt.Println(p0)
    switch p1 {
    case 1:
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint = 42
    v2 := []string{}
    v3 := make(map[string]int)
    v4 := []string{("" + ("\n\t" + "\n\t")), (("\n\t" + "\n\t") + "hello")}
    var v5 uint64 = 18446744073709551615
    if false {
        if (false && true) {
            v8 := GenS2{F0: false}
            v9 := GenS1{F0: ("\n\t" + "\n\t"), F1: true}
            _ = v8
            _ = v9
        } else {
            fmt.Println(v1)
            v5 -= 230
            switch "\n\t" {
            case "hello":
                v1 = uint(v5)
            default:
                if (-635 == 639) {
                    fmt.Println(v1, v5)
                }
            }
        }
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int8 = 0
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v1, v5)
    v5 = 420
    if (!((v1 - (v1 >> ((genFn0(42, 354, 715)) & 63))) > v1)) {
        var v12 int32 = 0
        fmt.Println(v12, v5)
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 int16 = 0
            _ = fi13
            _ = v14
        }
        _ = v12
    } else {
        v5 -= (v5 - (v5 % ((v5) | 1)))
    }
    if false {
        switch "hello" {
        case "":
            var v15 bool = ((true || true) || ((v5 * (v5 ^ v5)) == v5))
            _ = v15
        }
        if (!(((true || ((3.14159 >= -2.25) || (-82 != 100))) || false) && true)) {
            v5 = (v5 ^ (v5 % ((v5) | 1)))
            for fi16 := 0; fi16 < 14; fi16++ {
                v1 -= genFn0(-1, fi16, 715)
                _ = fi16
            }
        }
    } else {
        var v17 uint8 = 255
        _ = v17
    }
    var v18 uint8 = 2
    switch -939 {
    case -9223372036854775808:
        switch 0 {
        case 1:
            fmt.Println(v5, v5, v18)
        }
    default:
        v18 = v18
    }
    var v19 int8 = -128
    if false {
        v1 = v1
    }
    var v20 string = "a"
    for fi21 := 0; fi21 < 12; fi21++ {
        if ((fi21 % ((fi21) | 1)) == 2147483647) {
            switch v20 {
            case "\n\t":
                fmt.Println(fi21)
            case "foo bar":
                v22 := GenS3{F0: (v18 - (v18 + v18))}
                _ = v22
            case "hello":
                v23 := GenS4{F0: true}
                _ = v23
            }
            if false {
                var v24 uint64 = v5
                if false {
                    var v25 int8 = (-(v19))
                    v19 -= -43
                    _ = v25
                } else {
                    var v26 bool = (v5 < v5)
                    var v27 float64 = (1.5 / -2.25)
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 16; fi28++ {
                    var v29 uint32 = 799
                    _ = len(v3)
                    var v30 uint64 = 865
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = v24
            }
            _ = len(v2)
        } else {
            if (("foo bar" + ("foo bar" + "\n\t")) > ("\n\t" + "hello")) {
                v18 = (v18 / ((v18) | 1))
                var v31 uint32 = 1
                var v32 int8 = (-(v19))
                _ = v31
                _ = v32
            } else {
                v33 := GenS1{F0: ("foo bar" + (v20 + "foo bar")), F1: (false && (true || (v1 > (v1 | (v1 + 568)))))}
                v5 = v5
                _ = v33
            }
        }
        v34 := GenS5{F0: 42, F1: v20, F2: -9223372036854775808}
        v3[(((("" + (v20 + v20)) + ("\n\t" + v20)) + v20) + v20)] = 128
        _ = fi21
        _ = v34
    }
    v5 += v5
    v35 := GenS4{F0: true}
    for fi36 := 0; fi36 < 10; fi36++ {
        v19 += v19
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
