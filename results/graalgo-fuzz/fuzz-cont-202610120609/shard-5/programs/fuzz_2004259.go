package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int; F1 int; F2 int16 }
func genFn0() int8 {
    if true {
        if (55 < 42) {
            if true {
                var v5 int16 = -1
                _ = v5
            }
            var v6 int32 = 2147483647
            v6 = 42
            _ = v6
        } else {
            if true {
                v7 := []int{2147483647, -248, -309, 254}
                var v8 uint = 1
                _ = v7
                _ = v8
            } else {
                v9 := GenS1{F0: -192, F1: 2147483647}
                for fi10 := 0; fi10 < 1; fi10++ {
                    if (2 <= 0) {
                        var v11 int64 = -9223372036854775808
                        _ = v11
                    }
                    var v12 uint8 = 70
                    _ = fi10
                    _ = v12
                }
                var v13 float64 = 1.5
                _ = v13
                _ = v9
            }
            var v14 bool = true
            for fi15 := 0; fi15 < 4; fi15++ {
                v16 := GenS1{F0: -9223372036854775808, F1: 42}
                if false {
                    var v17 uint32 = 42
                    var v18 uint64 = 42
                    var v19 int = (fi15 - (fi15 - (-(fi15))))
                    _ = v17
                    _ = v18
                    _ = v19
                }
                for fi20 := 0; fi20 < 18; fi20++ {
                    v21 := GenS2{F0: 470, F1: fi15, F2: -802}
                    _ = fi20
                    _ = v21
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        var v22 int64 = 129
        _ = v22
    } else {
        if (0 < -2147483648) {
            v23 := []int{0, -9223372036854775808, 128}
            _ = v23
        }
        v24 := GenS2{F0: -9223372036854775808, F1: -1, F2: 0}
        var v25 float32 = 3.14159
        _ = v24
        _ = v25
    }
    v26 := []int{32768, 32768}
    _ = v26
    return -1
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[int]bool)
    var v3 uint16 = 331
    v4 := make(map[int]uint32)
    var v27 string = "hello"
    fmt.Println(v3)
    v28 := []int{-39, 128, -803}
    switch (v27 + "") {
    case "hello":
        v1 = 0.1
    default:
        if false {
            for fi29 := 0; fi29 < 11; fi29++ {
                fmt.Println(v1)
                _ = fi29
            }
            switch 127 {
            case -846:
                var v30 float64 = v1
                _ = v30
            case 2147483647:
                switch v27 {
                case "":
                    v31 := v2[-832]
                    _ = v31
                }
            default:
                v3 -= 65535
            }
            v28 = append(v28, 0)
        } else {
            var v32 uint16 = 511
            v33 := GenS1{F0: 9223372036854775807, F1: -1}
            _ = v32
            _ = v33
        }
    }
    if false {
        var v34 string = v27
        var v35 float64 = -2.25
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        v3 = v3
        _ = fi36
    }
    switch v27 {
    case "\n\t":
        var v37 uint8 = 1
        _ = v37
    case "":
        if ((0 == 42) || true) {
            for fi38 := 0; fi38 < 18; fi38++ {
                v39 := v4[fi38]
                var v40 float64 = v1
                delete(v4, fi38)
                _ = fi38
                _ = v39
                _ = v40
            }
            fmt.Println(v1)
        }
    case "foo bar":
        var v41 int = 362
        _ = v41
    default:
        v27 = v27
    }
    var v42 int = -2147483648
    var v43 int8 = -45
    v3 += v3
    if (false && false) {
        v42 += (-(v42))
        var v44 float64 = v1
        v45 := v2[(v42 | v42)]
        _ = v44
        _ = v45
    }
    v43 += 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
