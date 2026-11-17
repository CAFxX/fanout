package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint32) uint32 {
    v6 := []int{}
    if (true && false) {
        if false {
            _ = len(v6)
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := GenS1{F0: 38}
                _ = fi7
                _ = v8
            }
        }
        var v9 int16 = -1
        var v10 uint32 = (p1 / ((4294967295) | 1))
        _ = v10
        _ = v9
    } else {
        var v11 uint64 = 4294967295
        _ = v11
    }
    v12 := GenS2{F0: (("foo bar" + "foo bar") + "\n\t"), F1: 42, F2: -681}
    var v13 float32 = float32(1.5)
    v14 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v6
    return 245
}
func genFn1() string {
    var v15 int16 = 641
    v15 = (v15 * v15)
    switch 4294967295 {
    case -1:
        var v16 int = 128
        _ = v16
    case 4294967295:
        if false {
            fmt.Println(v15)
        }
    }
    var v17 uint16 = 633
    _ = v15
    _ = v17
    return "a"
}

func main() {
    v1 := []string{"a", ((("\n\t" + "a") + ((("hello" + "\n\t") + ("hello" + "foo bar")) + "a")) + "a"), "foo bar"}
    v2 := make(map[string]bool)
    var v3 float64 = 1e10
    var v4 uint8 = 249
    v5 := []int{1, -2147483648, 1, -2147483648}
    if true {
        var v18 float64 = v3
        v4 = v4
        _ = v18
    }
    if (!true) {
        switch ("\n\t" + "\n\t") {
        case "":
            fmt.Println(v4, v3)
        }
        v3 -= (0.1 / v3)
        v4 += 47
    }
    var v19 uint8 = 30
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int32 = 0
        switch fi20 {
        case 2147483647:
            _ = len(v2)
        case -2147483648:
            switch (^(fi20)) {
            case -755:
                var v22 float64 = v3
                _ = v22
            default:
                var v23 int32 = 42
                _ = v23
            }
        case 1:
            v24 := GenS3{F0: 796}
            _ = v24
        default:
            var v25 uint16 = 65535
            _ = v25
        }
        _ = fi20
        _ = v21
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 int32 = 0
        for fi28 := 0; fi28 < 14; fi28++ {
            switch (fi28 % ((fi26) | 1)) {
            case 42:
                if (32768 == fi28) {
                    fmt.Println(fi28, v27, v3)
                } else {
                    v29 := GenS2{F0: (("hello" + "\n\t") + ("\n\t" + "foo bar")), F1: (v4 - v4), F2: -2147483648}
                    _ = v29
                }
            default:
                v4 -= 4
            }
            _ = fi28
        }
        v30 := GenS3{F0: 461}
        _ = fi26
        _ = v27
        _ = v30
    }
    v4 += 0
    v31 := GenS2{F0: "foo bar", F1: v4, F2: 1}
    v4 = (v19 | (v4 << ((255) & 63)))
    var v32 uint8 = (v4 & v19)
    switch -211 {
    case 9223372036854775807:
        var v33 uint16 = 42
        _ = v33
    case 127:
        if true {
            if (!false) {
                _ = len(v1)
            }
        }
    case 0:
        var v34 uint = 696
        _ = v34
    }
    if false {
        var v35 uint64 = 18446744073709551615
        var v36 int16 = 42
        fmt.Println(v32)
        _ = v35
        _ = v36
    } else {
        var v37 uint64 = 0
        _ = v37
    }
    var v38 float64 = 1.5
    v19 += 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v4
    _ = v5
}
