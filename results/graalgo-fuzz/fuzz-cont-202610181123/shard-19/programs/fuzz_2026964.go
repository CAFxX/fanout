package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32) string {
    for fi5 := 0; fi5 < 12; fi5++ {
        if (127 <= 42) {
            v6 := GenS1{F0: ("foo bar" + "hello"), F1: -67, F2: 813}
            _ = v6
        }
        _ = fi5
    }
    v7 := []int{9223372036854775807, 32767, 32767}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return ((("a" + "") + "foo bar") + "foo bar")
}

func main() {
    var v1 int8 = -18
    var v2 int = -9223372036854775808
    var v3 string = ""
    var v4 float64 = 1e10
    v8 := []int{127, v2}
    if true {
        fmt.Println(v3)
        var v9 uint = 530
        switch (^(v2)) {
        case -791:
            switch 32768 {
            case 42:
                fmt.Println(v1, v2)
            case 128:
                v10 := GenS1{F0: v3, F1: -1, F2: 65535}
                _ = v10
            }
        case 2147483647:
            v9 = (v9 ^ (v9 + (v9 / (((^(v9))) | 1))))
        case 127:
            for fi11 := 0; fi11 < 13; fi11++ {
                v12 := GenS1{F0: v3, F1: v1, F2: (^(v9))}
                _ = fi11
                _ = v12
            }
        default:
            if true {
                fmt.Println(v1, v3)
                v8 = append(v8, v2)
                if len(v8) > 0 {
                    v8[0] = (v2 - (v2 / (((v2 - v2)) | 1)))
                }
            }
        }
        _ = v9
    } else {
        var v13 bool = (0 > -94)
        _ = v13
    }
    v1 += 127
    var v14 string = (v3 + ((v3 + v3) + v3))
    var v15 float64 = 3.14159
    v16 := GenS2{F0: 2147483647}
    v17 := GenS1{F0: v14, F1: (v1 % ((-1) | 1)), F2: 4294967295}
    var v18 uint = 2
    if false {
        var v19 int32 = 2147483647
        var v20 string = "hello"
        _ = v19
        _ = v20
    }
    var v21 uint64 = 0
    for fi22 := 0; fi22 < 6; fi22++ {
        if len(v8) > 0 {
            v8[0] = (fi22 % ((fi22) | 1))
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            v8 = append(v8, (v2 + (fi23 >> (((v18 | v18)) & 63))))
            switch 631 {
            case -1:
                var v24 float32 = 0.0
                _ = v24
            case -9223372036854775808:
                switch fi23 {
                case -2147483648:
                    switch (v2 / (((v2 % (((fi22 + (v2 << (((v18 - v18)) & 63)))) | 1))) | 1)) {
                    case 0:
                        v15 += ((v4 * v4) - ((0.0 - v4) / ((v4 * 1.5) - (v4 + v15))))
                    case 943:
                        var v25 int64 = 42
                        _ = v25
                    }
                default:
                    v14 = v3
                }
            }
            _ = fi23
        }
        _ = fi22
    }
    if false {
        v14 = v14
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    } else {
        if (true || false) {
            fmt.Println(v4, v2, v1)
        }
        if (false && (-1 >= (v2 + (-(v2))))) {
            if false {
                v8 = append(v8, (v2 << (((^(v18))) & 63)))
                var v26 bool = (!(0 == 32767))
                _ = v26
            }
            fmt.Println(v14)
            var v27 float64 = 0.0
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v8
}
