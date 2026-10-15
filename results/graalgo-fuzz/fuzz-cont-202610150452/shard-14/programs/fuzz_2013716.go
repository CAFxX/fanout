package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []uint32{42}
    v2 := make(map[uint32]int)
    var v3 int = -2147483648
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    v6 := []uint32{2147483648, 2147483648, 4294967295}
    if len(v6) > 0 {
        v6[0] = 0
    }
    v3 = (v3 << ((18446744073709551615) & 63))
    v3 = v3
    v5[(v3 / ((-613) | 1))] = "a"
    if true {
        fmt.Println(v3)
        var v7 uint32 = 2147483647
        _ = v7
    } else {
        var v8 int64 = -973
        _ = v8
    }
    switch 127 {
    case 141:
        fmt.Println(v3)
    case 1:
        for fi9 := 0; fi9 < 6; fi9++ {
            switch "\n\t" {
            case "a":
                for fi10 := 0; fi10 < 19; fi10++ {
                    v11 := GenS1{F0: 874, F1: (("hello" + "hello") + "hello"), F2: 42}
                    v3 = (-(v3))
                    _ = fi10
                    _ = v11
                }
            case "":
                if (849 != 42) {
                    v12 := GenS2{F0: true, F1: 65535}
                    for fi13 := 0; fi13 < 3; fi13++ {
                        v3 = (fi13 << ((255) & 63))
                        _ = fi13
                    }
                    v3 = v3
                    _ = v12
                }
            }
            delete(v5, (-(v3)))
            _ = fi9
        }
    case 385:
        var v14 uint = 65535
        _ = v14
    }
    v15 := GenS3{F0: 127}
    v3 = v3
    if (4294967295 == 129) {
        switch "\n\t" {
        case "hello":
            var v16 uint16 = 32767
            _ = v16
        case "\n\t":
            var v17 string = ("\n\t" + "foo bar")
            _ = v17
        case "":
            if false {
                v3 -= 4294967295
                var v18 int = (v3 * (^(v3)))
                fmt.Println(v18, v3)
                _ = v18
            } else {
                var v19 uint64 = 752
                _ = v19
            }
        default:
            v1 = append(v1, 670)
        }
        if (-1 > -59) {
            for fi20 := 0; fi20 < 15; fi20++ {
                v3 = 127
                var v21 int8 = 1
                v21 -= v21
                _ = fi20
                _ = v21
            }
            var v22 uint32 = 2147483647
            _ = v22
        }
    } else {
        v23 := GenS1{F0: 4294967295, F1: ("foo bar" + "foo bar"), F2: 42}
        fmt.Println(v3)
        switch -712 {
        case 1:
            if len(v6) > 0 {
                v6[0] = 126
            }
        case -350:
            v3 += 811
        }
        _ = v23
    }
    if true {
        for fi24 := 0; fi24 < 19; fi24++ {
            for fi25 := 0; fi25 < 12; fi25++ {
                var v26 int64 = 557
                if true {
                    fmt.Println(v3, fi25, fi24)
                } else {
                    v27 := GenS4{F0: -173}
                    var v28 float64 = (1.5 + 1e10)
                    _ = v27
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
            var v29 int8 = -78
            switch (v3 | -481) {
            case 4294967295:
                v29 = (^(v29))
            }
            _ = fi24
            _ = v29
        }
        v30 := GenS1{F0: 2, F1: "foo bar", F2: 134}
        _ = v30
    }
    delete(v4, (((("a" + ("\n\t" + ("foo bar" + "\n\t"))) + ("foo bar" + ("a" + ("foo bar" + ("a" + "\n\t"))))) + "a") + "a"))
    var v31 float64 = 1e10
    v32 := GenS1{F0: 325, F1: (("foo bar" + "a") + "foo bar"), F2: 1}
    v31 += v31
    var v33 float32 = (0.1 * float32(((v31 + 1e10) / v31)))
    var v34 int8 = 36
    v34 -= (-(v34))
    var v35 uint = 247
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
