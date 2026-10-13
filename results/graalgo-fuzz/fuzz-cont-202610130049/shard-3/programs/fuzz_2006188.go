package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
func genFn0() int16 {
    var v6 int = 9223372036854775807
    var v7 float64 = -2.25
    if (-1 > 62) {
        fmt.Println(v7)
    }
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 int32 = 0
    var v4 uint8 = 127
    v5 := make(map[uint32]bool)
    var v8 uint32 = 2147483647
    fmt.Println(v4, v8)
    var v9 int8 = 100
    fmt.Println(v9)
    var v10 uint64 = 42
    v3 += (^(v3))
    if (-2.25 != (3.14159 * float64(0.1))) {
        v5[98] = (uint(v10) >= 4294967295)
    } else {
        v9 = v9
    }
    if false {
        switch "a" {
        case "foo bar":
            for fi11 := 0; fi11 < 12; fi11++ {
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 int8 = v9
                    var v14 uint8 = v4
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                if true {
                    v9 -= 1
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 bool = (-2.25 < 1e10)
                        switch "hello" {
                        case "hello":
                            fmt.Println(v9)
                        case "foo bar":
                            v9 = (v9 * 107)
                        case "\n\t":
                            v17 := GenS1{F0: 9223372036854775807}
                            _ = v17
                        default:
                            var v18 int64 = 9223372036854775807
                            _ = v18
                        }
                        _ = fi15
                        _ = v16
                    }
                } else {
                    fmt.Println(fi11, v3)
                }
                var v19 float32 = 0.0
                _ = fi11
                _ = v19
            }
        case "hello":
            var v20 uint16 = 1
            _ = v20
        }
    }
    fmt.Println(v10)
    var v21 bool = (0 > (v9 & int8(v10)))
    v9 = (v9 | (v9 * 0))
    fmt.Println(v8)
    switch -909 {
    case 0:
        var v22 uint32 = 4294967295
        _ = v22
    case -450:
        var v23 uint32 = 2147483647
        _ = v23
    default:
        v24 := []int{32767, -9223372036854775808}
        _ = v24
    }
    switch 0 {
    case -980:
        v25 := GenS2{F0: 397}
        _ = v25
    case 855:
        var v26 int8 = v9
        _ = v26
    default:
        v27 := GenS2{F0: -56}
        _ = v27
    }
    var v28 int16 = 1
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 uint64 = uint64(v4)
        _ = fi29
        _ = v30
    }
    if v21 {
        if (0.1 >= -2.25) {
            v31 := GenS2{F0: -1}
            _ = v31
        } else {
            for fi32 := 0; fi32 < 7; fi32++ {
                v4 = (v4 / ((128) | 1))
                _ = fi32
            }
        }
        switch "a" {
        case "a":
            var v33 int32 = v3
            _ = v33
        case "":
            fmt.Println(v4, v21, v9)
        case "\n\t":
            v34 := []int{4294967295, 128}
            _ = v34
        }
        if (-1 <= 42) {
            var v35 int = -2147483648
            fmt.Println(v9)
            _ = v35
        }
    } else {
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 int64 = 42
            _ = fi36
            _ = v37
        }
        v28 = -1
        fmt.Println(v28)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
