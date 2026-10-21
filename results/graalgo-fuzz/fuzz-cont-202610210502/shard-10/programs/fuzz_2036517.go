package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 2
    var v3 float64 = (float64(-2.25) / 1.0)
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: 42, F1: (fi7 ^ fi7)}
        _ = fi7
        _ = v8
    }
    if true {
        v2 += (v2 * (v2 | (v2 + v2)))
        v9 := []int{-90}
        _ = v9
    } else {
        v10 := []int{42, 868, 32767, -2147483648}
        var v11 float64 = v3
        _ = v10
        _ = v11
    }
    var v12 uint32 = 2147483648
    v12 += v12
    v13 := GenS1{F0: 1, F1: 42}
    fmt.Println(v2, v12, v3)
    for fi14 := 0; fi14 < 16; fi14++ {
        if (2147483647 != 2147483647) {
            switch "foo bar" {
            case "foo bar":
                v15 := []int{1, fi14}
                _ = v15
            case "hello":
                fmt.Println(v2)
            default:
                delete(v6, 706)
            }
            var v16 float32 = -2.25
            _ = v16
        } else {
            v3 += v3
            v17 := GenS1{F0: 2147483647, F1: 9223372036854775807}
            var v18 int = fi14
            _ = v17
            _ = v18
        }
        if (true && (!true)) {
            fmt.Println(v3, v2)
        }
        _ = fi14
    }
    v2 = v2
    v3 = v3
    var v19 uint = uint(v2)
    v2 -= v2
    switch 1 {
    case -2147483648:
        v20 := GenS2{F0: 32767, F1: 89}
        _ = v20
    case 127:
        for fi21 := 0; fi21 < 14; fi21++ {
            if (true || (true && true)) {
                var v22 float32 = 0.0
                var v23 uint8 = 1
                fmt.Println(v3, fi21)
                _ = v22
                _ = v23
            }
            var v24 string = "hello"
            _ = fi21
            _ = v24
        }
    case 128:
        for fi25 := 0; fi25 < 16; fi25++ {
            for fi26 := 0; fi26 < 18; fi26++ {
                if true {
                    var v27 int = (-(fi25))
                    switch "" {
                    case "a":
                        var v28 int64 = 0
                        _ = v28
                    case "":
                        v19 += (^(v19))
                    }
                    v2 += v2
                    _ = v27
                } else {
                    var v29 uint8 = 1
                    _ = v29
                }
                var v30 bool = (!(-32768 != -303))
                _ = fi26
                _ = v30
            }
            var v31 uint16 = v2
            var v32 int64 = 9223372036854775807
            _ = fi25
            _ = v31
            _ = v32
        }
    default:
        v33 := GenS2{F0: 156, F1: 1}
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
