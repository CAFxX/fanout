package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint, p1 bool) int8 {
    if (0 > -208) {
        if false {
            p1 = true
            var v7 uint8 = 234
            fmt.Println(p1, p1, p0)
            _ = v7
        }
    }
    if true {
        v8 := []int{856, 1, 2147483647}
        p1 = (!p1)
        switch 42 {
        case -224:
            switch "" {
            case "":
                var v9 int8 = -126
                _ = v9
            }
        default:
            fmt.Println(p0)
        }
        _ = v8
    }
    v10 := []int{}
    var v11 uint = (p0 & (^(p0)))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -128
}

func main() {
    var v1 float64 = (1e10 / 1e10)
    v2 := make(map[string]uint32)
    var v3 uint16 = 565
    var v4 float32 = 0.1
    var v5 bool = false
    var v6 int32 = -1
    var v12 uint8 = 0
    if ((false || v5) && true) {
        v6 = (v6 * v6)
        fmt.Println(v4, v1, v12)
    }
    v13 := GenS1{F0: 1, F1: "\n\t"}
    switch -994 {
    case 42:
        v4 = v4
    case 127:
        v14 := GenS1{F0: 0, F1: ""}
        _ = v14
    default:
        fmt.Println(v1, v5, v4)
    }
    v3 -= v3
    if (v3 == 0) {
        var v15 int16 = -1
        switch (("" + "") + (("hello" + ("a" + "a")) + ("a" + ((("hello" + "") + (("" + "hello") + "\n\t")) + "a")))) {
        case "a":
            if ((v12 - v12) == (^(v12))) {
                v16 := []int{2147483647, 4294967295, int(v6)}
                var v17 float64 = v1
                _ = v16
                _ = v17
            }
        case "":
            if v5 {
                var v18 int8 = 42
                var v19 int = 1
                _ = v18
                _ = v19
            }
        case "hello":
            var v20 uint64 = 51
            _ = v20
        }
        v5 = (505 <= 4294967295)
        _ = v15
    }
    var v21 string = "a"
    v12 += v12
    switch -1 {
    case 4294967295:
        fmt.Println(v21, v1, v1)
    case -2147483648:
        if v5 {
            if (!(128 <= uint(v12))) {
                fmt.Println(v6, v3)
            }
        }
    case 313:
        v22 := GenS1{F0: 716, F1: v21}
        _ = v22
    }
    fmt.Println(v3)
    v23 := GenS2{F0: (v12 << ((42) & 63)), F1: 32767}
    fmt.Println(v6, v1)
    v24 := GenS1{F0: 0, F1: (("foo bar" + v21) + v21)}
    switch v21 {
    case "hello":
        var v25 uint8 = v12
        _ = v25
    }
    v3 -= v3
    var v26 int32 = 2147483647
    if (!(v6 == (v26 - (v26 | v6)))) {
        v27 := []int{-2147483648, 0, 2147483647, -607}
        switch v21 {
        case "hello":
            if v5 {
                if true {
                    for fi28 := 0; fi28 < 15; fi28++ {
                        for fi29 := 0; fi29 < 19; fi29++ {
                            switch (^(fi28)) {
                            case -268:
                                v1 = (v1 - (v1 - float64(3.14159)))
                            case 470:
                                v26 -= v26
                                var v30 int = fi28
                                _ = v30
                            case 0:
                                v31 := GenS2{F0: (^(v12)), F1: 1000}
                                _ = v31
                            default:
                                var v32 bool = true
                                _ = v32
                            }
                            v12 = (v12 + (v12 - (v12 / (((v12 * v12)) | 1))))
                            fmt.Println(v12, v26, v1)
                            _ = fi29
                        }
                        v33 := GenS1{F0: 402, F1: v21}
                        _ = fi28
                        _ = v33
                    }
                }
                if (!v5) {
                    var v34 uint32 = 2147483647
                    _ = v34
                }
                v1 = v1
            }
        default:
            var v35 int64 = -9223372036854775808
            _ = v35
        }
        v1 = v1
        _ = v27
    } else {
        v26 = (v26 * -2147483648)
    }
    v36 := GenS1{F0: 4294967295, F1: v21}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
