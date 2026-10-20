package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int8 }
func genFn0() uint {
    var v6 uint32 = 28
    var v7 string = "\n\t"
    var v8 uint64 = 104
    if (2147483647 >= 0) {
        var v9 int8 = 1
        _ = v9
    } else {
        var v10 string = v7
        if (false || ((^(v8)) == (^(v8)))) {
            fmt.Println(v6, v6, v7)
            fmt.Println(v6)
            v11 := []int{127, -1}
            _ = v11
        }
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return 334
}
func genFn1() string {
    v12 := []int{0, 583, 830, 9223372036854775807}
    var v13 int = 32767
    if true {
        v13 = v13
        v12 = append(v12, (v13 * (v13 % ((0) | 1))))
        var v14 int8 = 57
        _ = v14
    }
    if (true && true) {
        v13 = v13
        v13 += v13
        fmt.Println(v13)
    } else {
        var v15 int32 = 494
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 int = (v13 - (v13 >> ((42) & 63)))
            v13 = (fi16 & -2147483648)
            var v18 uint8 = 79
            _ = fi16
            _ = v17
            _ = v18
        }
        switch (("\n\t" + "\n\t") + ("foo bar" + "foo bar")) {
        case "a":
            if false {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                var v19 bool = true
                v15 += (v15 << ((322) & 63))
                _ = v19
            }
        case "hello":
            var v20 uint8 = 42
            _ = v20
        default:
            var v21 float32 = 1.5
            _ = v21
        }
        _ = v15
    }
    if (-128 == 127) {
        v22 := GenS1{F0: true, F1: 271}
        fmt.Println(v13)
        _ = v22
    } else {
        for fi23 := 0; fi23 < 12; fi23++ {
            switch (("hello" + (("" + ("\n\t" + "foo bar")) + "")) + "foo bar") {
            case "\n\t":
                v13 = fi23
            default:
                v24 := GenS2{F0: 192, F1: 18446744073709551615, F2: -44}
                _ = v24
            }
            var v25 bool = (415 < 4294967295)
            _ = fi23
            _ = v25
        }
        var v26 int = (v13 * 9223372036854775807)
        _ = v26
    }
    _ = v12
    _ = v13
    return "foo bar"
}

func main() {
    v1 := []string{"a"}
    var v2 int = -2147483648
    v3 := make(map[int]bool)
    var v4 int = (v2 - v2)
    v5 := make(map[int]string)
    fmt.Println(v4, v4)
    var v27 int64 = -688
    for fi28 := 0; fi28 < 2; fi28++ {
        fmt.Println(v27, v4)
        _ = fi28
    }
    fmt.Println(v4, v4)
    var v29 int16 = 32767
    for fi30 := 0; fi30 < 20; fi30++ {
        switch "hello" {
        case "":
            if false {
                v2 = (v2 & (v2 + (fi30 >> ((426) & 63))))
                fmt.Println(v29)
                for fi31 := 0; fi31 < 13; fi31++ {
                    v27 = v27
                    fmt.Println(v2)
                    v29 -= (-(v29))
                    _ = fi31
                }
            }
        case "foo bar":
            v4 = v4
        }
        v32 := GenS2{F0: 42, F1: 16, F2: 37}
        _ = fi30
        _ = v32
    }
    switch v2 {
    case 317:
        v33 := GenS1{F0: false, F1: 255}
        _ = v33
    case 419:
        v34 := GenS1{F0: (v27 < -941), F1: 42}
        _ = v34
    case 128:
        v29 -= 256
    }
    v29 = v29
    v27 -= (-(v27))
    v35 := GenS1{F0: (32767 < 0), F1: 1}
    var v36 int16 = v29
    var v37 float64 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
