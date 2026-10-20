package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint32; F1 uint; F2 uint64 }
func genFn0(p0 int8, p1 int64, p2 uint32) int {
    p0 += p0
    var v6 uint = 1
    switch 771 {
    case 4294967295:
        p2 = 2
    }
    if true {
        v7 := []int{42}
        var v8 bool = false
        _ = v7
        _ = v8
    } else {
        v9 := []int{127, 675, 32768, 42}
        if len(v9) > 0 {
            v9[0] = 42
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 359
}

func main() {
    var v1 uint = 0
    var v2 float64 = 0.1
    var v3 int16 = 0
    v4 := []string{"foo bar"}
    v5 := make(map[int]string)
    var v10 int = genFn0(127, -9223372036854775808, 65535)
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            v1 = (v1 / ((v1) | 1))
            var v13 uint64 = 0
            if (65535 == 961) {
                switch ("hello" + "foo bar") {
                case "foo bar":
                    var v14 int8 = 42
                    _ = v14
                case "hello":
                    v15 := GenS1{F0: fi11, F1: (v1 + (v1 / (((v1 >> ((2) & 63))) | 1)))}
                    _ = v15
                default:
                    var v16 int = 9223372036854775807
                    _ = v16
                }
                var v17 bool = false
                v18 := GenS1{F0: v10, F1: (v1 & v1)}
                v1 += v1
                var v19 uint64 = 42
                _ = v17
                _ = v18
                _ = v19
            } else {
                fmt.Println(v3, v13)
                var v20 int32 = 0
                v1 = uint(v13)
                _ = v20
            }
            var v21 int16 = (v3 / (((v3 ^ (^(v3)))) | 1))
            _ = fi12
            _ = v13
            _ = v21
        }
        _ = fi11
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        switch ((("hello" + "a") + ("foo bar" + "\n\t")) + (("foo bar" + ("\n\t" + "a")) + "a")) {
        case "hello":
            v23 := GenS2{F0: 0, F1: v1, F2: 384}
            _ = v23
        default:
            if ((!(42 < 42)) && (4294967295 >= 16)) {
                switch fi22 {
                case 0:
                    var v24 int16 = v3
                    _ = v24
                case 2147483647:
                    if len(v4) > 0 {
                        v4[0] = ("" + ("foo bar" + "\n\t"))
                    }
                case -2147483648:
                    v1 -= (v1 - v1)
                }
                v10 -= genFn0(42, -1, 65535)
            } else {
                var v25 int64 = -9223372036854775808
                v2 = v2
                _ = v25
            }
        }
        v4 = append(v4, (("hello" + "hello") + ""))
        _ = fi22
    }
    var v26 int = (v10 | genFn0(4, int64(v1), 42))
    fmt.Println(v1, v26)
    v26 -= genFn0(int8(v26), 115, 910)
    if (0 < 125) {
        var v27 uint64 = 42
        v1 -= v1
        for fi28 := 0; fi28 < 7; fi28++ {
            v26 = (v10 & (v10 + 456))
            v29 := GenS2{F0: 525, F1: 42, F2: (^(v27))}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        delete(v5, v10)
        var v30 string = ("hello" + "a")
        _ = v30
    }
    v31 := GenS2{F0: 2, F1: (v1 - 0), F2: 1}
    for fi32 := 0; fi32 < 8; fi32++ {
        v10 = fi32
        v33 := v5[0]
        fmt.Println(fi32, v3, v26)
        _ = fi32
        _ = v33
    }
    if (!(!(1 != 1))) {
        fmt.Println(v1)
    }
    if true {
        var v34 int8 = 100
        fmt.Println(v2)
        var v35 float32 = ((float32((float64(1e10) / v2)) - (float32((float64(((float32((float64(3.14159) + (0.1 - v2))) + (-2.25 / float32((v2 - (v2 - 3.14159))))) / 3.14159)) * 0.0)) / 3.14159)) / 1.5)
        _ = v34
        _ = v35
    }
    switch v26 {
    case -9223372036854775808:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 0:
        switch (v10 << (((v1 | v1)) & 63)) {
        case 0:
            v26 = v10
        default:
            v2 += v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
