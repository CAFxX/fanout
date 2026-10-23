package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 uint32) uint16 {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := []int{0}
        _ = fi4
        _ = v5
    }
    if true {
        var v6 uint8 = 1
        v7 := []int{-9223372036854775808, 127}
        _ = v6
        _ = v7
    }
    switch "\n\t" {
    case "\n\t":
        v8 := GenS1{F0: 1000, F1: p0}
        _ = v8
    case "foo bar":
        v9 := GenS2{F0: -9223372036854775808, F1: -742}
        _ = v9
    case "hello":
        var v10 uint8 = 1
        _ = v10
    }
    _ = p0
    return 794
}
func genFn1(p0 uint8) string {
    var v11 int16 = 42
    switch 127 {
    case 0:
        if (2147483647 == 2147483647) {
            v12 := GenS2{F0: 90, F1: -32768}
            if false {
                v11 = v11
                v11 -= -1
            }
            if false {
                v13 := []int{int(v11), 2147483647, 1}
                var v14 int16 = v11
                _ = v13
                _ = v14
            }
            _ = v12
        } else {
            p0 -= (p0 - 128)
            v15 := []int{-403, 42}
            _ = v15
        }
    case 4294967295:
        var v16 uint16 = 42
        _ = v16
    }
    var v17 string = ("a" + ("hello" + ("hello" + "\n\t")))
    _ = p0
    _ = v11
    _ = v17
    return "\n\t"
}

func main() {
    v1 := []uint32{2147483647, 2147483648, 26, 40}
    var v2 string = "hello"
    v3 := make(map[uint32]int)
    v2 = "a"
    switch 4294967295 {
    case 127:
        var v18 float32 = (-2.25 - 0.0)
        _ = v18
    default:
        fmt.Println(v2)
    }
    var v19 int64 = 238
    if false {
        fmt.Println(v2, v19)
    } else {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        var v20 float32 = 0.0
        if true {
            switch "\n\t" {
            case "":
                if (245 != 326) {
                    var v21 uint8 = 255
                    var v22 uint = 255
                    _ = v21
                    _ = v22
                } else {
                    var v23 int16 = -633
                    _ = v23
                }
            }
        } else {
            switch (v2 + (v2 + ((v2 + v2) + (("foo bar" + "\n\t") + (v2 + v2))))) {
            case "hello":
                var v24 uint64 = 42
                _ = v24
            case "foo bar":
                var v25 uint = 0
                _ = v25
            case "a":
                switch v2 {
                case "foo bar":
                    var v26 float64 = (0.0 + (1.5 / 1.0))
                    _ = v26
                case "\n\t":
                    switch "a" {
                    case "foo bar":
                        var v27 int32 = 1
                        _ = v27
                    default:
                        v20 += (3.14159 / v20)
                    }
                default:
                    v3[1] = 781
                }
            default:
                v2 = "a"
            }
        }
        _ = v20
    }
    if (true || true) {
        v28 := GenS1{F0: -1, F1: 2147483648}
        _ = v28
    } else {
        v3[2147483648] = 759
    }
    var v29 string = ""
    v29 = ("a" + ((v29 + "foo bar") + "hello"))
    v30 := GenS1{F0: -473, F1: 1}
    v31 := GenS2{F0: 42, F1: 274}
    switch -2147483648 {
    case 32767:
        v19 = v19
    }
    v32 := GenS2{F0: (v19 / ((v19) | 1)), F1: 485}
    delete(v3, 1)
    var v33 int64 = -70
    delete(v3, 274)
    delete(v3, 62)
    v29 = (((v29 + v2) + (v2 + ("hello" + v29))) + v29)
    var v34 uint32 = 299
    for fi35 := 0; fi35 < 4; fi35++ {
        var v36 int64 = v33
        v37 := GenS2{F0: v19, F1: -32768}
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
