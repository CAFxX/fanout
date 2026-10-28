package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 bool) int64 {
    fmt.Println(p1, p0)
    fmt.Println(p1, p1)
    v7 := GenS1{F0: "\n\t"}
    var v8 float64 = (1.5 * 0.1)
    switch "\n\t" {
    case "hello":
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 float64 = (-2.25 + -2.25)
            _ = fi9
            _ = v10
        }
    case "\n\t":
        v11 := []int{9223372036854775807, -834, 127}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1() bool {
    var v12 int = 674
    switch (^(v12)) {
    case 772:
        v12 = 698
    case 1:
        fmt.Println(v12)
    case 2147483647:
        var v13 bool = true
        _ = v13
    default:
        fmt.Println(v12)
    }
    v14 := []int{(v12 ^ (-(v12))), (v12 & (v12 + 4294967295))}
    v15 := GenS2{F0: 42}
    v12 = (-(v12))
    _ = v12
    _ = v14
    _ = v15
    return true
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint32 = 2147483648
    var v3 int32 = -603
    v4 := make(map[string]int)
    var v5 uint16 = 545
    var v6 int32 = v1
    var v16 uint64 = 0
    if true {
        fmt.Println(v5, v1, v6)
    }
    v2 += (v2 % ((v2) | 1))
    if false {
        v6 = (v1 + 2147483647)
    } else {
        switch 557 {
        case -12:
            v17 := []int{-2147483648, 9223372036854775807, -344, 297}
            _ = v17
        }
        switch ("a" + "") {
        case "foo bar":
            if (false && true) {
                var v18 uint16 = (^(v5))
                _ = v18
            } else {
                v19 := GenS2{F0: v5}
                _ = v19
            }
        case "a":
            var v20 uint64 = (v16 % ((2) | 1))
            _ = v20
        default:
            v4[""] = -683
        }
    }
    v21 := v4[""]
    fmt.Println(v6, v5)
    v22 := GenS2{F0: v5}
    fmt.Println(v21)
    var v23 int32 = 190
    if (((681 >= 782) || false) && false) {
        fmt.Println(v23, v21)
        switch ("hello" + ("\n\t" + "a")) {
        case "foo bar":
            var v24 uint8 = 150
            _ = v24
        default:
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 uint8 = 44
                _ = fi25
                _ = v26
            }
        }
    }
    v27 := GenS3{F0: "foo bar"}
    v6 += 2147483647
    var v28 int64 = -687
    v29 := GenS3{F0: "foo bar"}
    if (("a" + ("hello" + ("\n\t" + ("" + "a")))) == ("hello" + "\n\t")) {
        var v30 int64 = v28
        switch (v21 * -2147483648) {
        case 42:
            var v31 uint8 = 139
            _ = v31
        case 599:
            fmt.Println(v6, v28, v28)
        case 128:
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 uint8 = 42
                _ = fi32
                _ = v33
            }
        }
        var v34 int64 = (-(v30))
        _ = v30
        _ = v34
    }
    v35 := []int{(v21 | (^(v21)))}
    var v36 uint64 = 2
    for fi37 := 0; fi37 < 12; fi37++ {
        if false {
            var v38 int32 = (v1 / (((v23 * -1)) | 1))
            var v39 int16 = 1
            for fi40 := 0; fi40 < 20; fi40++ {
                var v41 uint = 0
                var v42 bool = false
                _ = len(v35)
                v21 -= (fi40 % ((4294967295) | 1))
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = v38
            _ = v39
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
