package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
func genFn0(p0 string) bool {
    v7 := []int{4294967295, 9223372036854775807, -395}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int64 = -5
        v10 := GenS1{F0: 42, F1: v9, F2: ("" + p0)}
        if (3.14159 < (1e10 + 3.14159)) {
            v9 -= 42
            var v11 float64 = 0.0
            if len(v7) > 0 {
                v7[0] = 9223372036854775807
            }
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    p0 = p0
    p0 = ((("a" + (("a" + "a") + p0)) + p0) + p0)
    _ = p0
    _ = v7
    return true
}
func genFn1(p0 int) int32 {
    p0 = (-(p0))
    var v12 uint16 = 1
    _ = p0
    _ = v12
    return 2147483647
}

func main() {
    var v1 int16 = 1000
    var v2 uint8 = 127
    v3 := []string{("" + ""), "\n\t"}
    var v4 string = (("foo bar" + ("hello" + (("hello" + "") + "foo bar"))) + "\n\t")
    v5 := make(map[uint32]int)
    var v6 uint32 = 4294967295
    var v13 bool = (false && ((v6 << ((2) & 63)) != (^(v6))))
    var v14 int8 = 49
    var v15 string = (v4 + (("\n\t" + (("foo bar" + "\n\t") + "\n\t")) + v4))
    fmt.Println(v2, v14)
    _ = len(v3)
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS1{F0: v6, F1: -1, F2: v4}
        var v18 uint64 = 591
        _ = fi16
        _ = v17
        _ = v18
    }
    if len(v3) > 0 {
        v3[0] = ("a" + v4)
    }
    v19 := GenS1{F0: (^(v6)), F1: -9223372036854775808, F2: "hello"}
    if true {
        v6 += v6
        var v20 float64 = 0.1
        var v21 int64 = 0
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 15; fi22++ {
            v13 = ("hello" > v4)
            _ = fi22
        }
        if true {
            var v23 float64 = (0.1 * -2.25)
            var v24 float32 = -2.25
            var v25 int16 = (v1 << ((2) & 63))
            _ = v23
            _ = v24
            _ = v25
        } else {
            v26 := GenS1{F0: (v6 / ((v6) | 1)), F1: 257, F2: ((((v4 + ("" + "")) + "foo bar") + v15) + v4)}
            _ = v26
        }
        var v27 float64 = (((0.0 / (0.0 - (((3.14159 + 0.0) * (0.1 - 1.5)) / 1.5))) - (((0.1 - 0.1) / 1.0) - 1e10)) - (0.0 * (1.5 / 3.14159)))
        _ = v27
    }
    switch 955 {
    case 128:
        if (!(!v13)) {
            var v28 int8 = v14
            switch v15 {
            case "":
                var v29 uint8 = (v2 ^ (v2 - v2))
                _ = v29
            }
            v30 := v5[v6]
            _ = v28
            _ = v30
        }
    default:
        var v31 int64 = -1
        _ = v31
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        _ = len(v5)
        _ = len(v3)
        var v33 uint8 = v2
        _ = fi32
        _ = v33
    }
    var v34 uint32 = 41
    v14 = (v14 * v14)
    for fi35 := 0; fi35 < 10; fi35++ {
        fmt.Println(v2, fi35, v1)
        var v36 uint8 = 254
        var v37 uint = 544
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
