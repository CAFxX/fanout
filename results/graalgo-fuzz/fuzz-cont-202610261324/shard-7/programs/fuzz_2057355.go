package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string }
func genFn0(p0 int, p1 uint, p2 uint32) uint32 {
    v5 := GenS1{F0: 4294967295, F1: 9223372036854775807}
    if true {
        v6 := GenS1{F0: 4294967295, F1: 9223372036854775807}
        switch "foo bar" {
        case "foo bar":
            var v7 int = p0
            _ = v7
        }
        _ = v6
    } else {
        p1 = (p1 << (((^(p1))) & 63))
        var v8 int = 128
        p0 -= 80
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint8 = 127
        if (p1 >= (p1 ^ p1)) {
            fmt.Println(p0, fi9)
        } else {
            v11 := GenS1{F0: 274, F1: 0}
            v12 := GenS2{F0: 659, F1: 0}
            _ = v11
            _ = v12
        }
        if false {
            v13 := GenS3{F0: 9223372036854775807}
            v14 := GenS3{F0: -9223372036854775808}
            _ = v13
            _ = v14
        } else {
            fmt.Println(p0, p0, fi9)
        }
        _ = fi9
        _ = v10
    }
    p0 = p0
    var v15 float64 = (1.5 - 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v5
    return (p2 | p2)
}
func genFn1() int8 {
    var v16 uint64 = 42
    v16 += (v16 ^ v16)
    v16 = v16
    v16 -= (^(v16))
    _ = v16
    return 42
}

func main() {
    var v1 int8 = 127
    v2 := make(map[string]string)
    var v3 int8 = v1
    v4 := make(map[int]int)
    v17 := GenS4{F0: (!true)}
    fmt.Println(v3, v3)
    v18 := v4[2147483647]
    switch "foo bar" {
    case "\n\t":
        var v19 string = ("a" + ("hello" + "a"))
        _ = v19
    case "":
        for fi20 := 0; fi20 < 14; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                v18 += (fi20 | 32768)
                var v22 int64 = -9223372036854775808
                v18 += v18
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    case "a":
        var v23 float32 = 3.14159
        _ = v23
    default:
        v18 = v18
    }
    v18 = v18
    var v24 bool = false
    var v25 uint16 = 65535
    fmt.Println(v18, v24, v24)
    v24 = true
    v1 += v1
    v1 -= (v1 ^ v3)
    var v26 int = (^(v18))
    var v27 string = (("hello" + "\n\t") + "foo bar")
    if v24 {
        v28 := GenS3{F0: -1}
        _ = v28
    } else {
        if v24 {
            var v29 int16 = 1
            v24 = (v27 >= v27)
            _ = v29
        } else {
            var v30 int = -910
            v24 = (127 < (v1 << ((42) & 63)))
            v30 -= v26
            _ = v30
        }
        if ((0.0 / 0.1) < 0.0) {
            if (((v24 && v24) && (2 > v25)) || (!(255 != 808))) {
                v24 = (!false)
            } else {
                switch 988 {
                case 32767:
                    v3 += (-(v3))
                default:
                    v27 = "\n\t"
                }
                var v31 int64 = 250
                _ = v31
            }
        } else {
            v24 = (0 > 83)
            switch v27 {
            case "foo bar":
                var v32 uint64 = 4294967295
                _ = v32
            case "a":
                v33 := GenS4{F0: (-2147483648 >= 151)}
                _ = v33
            }
        }
    }
    v34 := GenS5{F0: v27}
    switch (v26 % ((v18) | 1)) {
    case 0:
        fmt.Println(v26, v25)
    case -2147483648:
        v18 += (v18 * 128)
    case 9223372036854775807:
        var v35 bool = v24
        _ = v35
    default:
        delete(v4, 42)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
}
