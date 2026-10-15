package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    var v1 float64 = ((1.5 / 0.1) - (3.14159 + 1.5))
    var v2 float32 = 0.0
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    var v5 int32 = -39
    v6 := make(map[int]uint32)
    switch 9223372036854775807 {
    case 0:
        var v7 uint32 = 2
        _ = v7
    case 42:
        if (false && ((true && (255 <= 848)) || true)) {
            var v8 uint64 = 1
            var v9 bool = false
            _ = v8
            _ = v9
        } else {
            var v10 int64 = -111
            _ = len(v4)
            _ = v10
        }
    case 32767:
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                v5 = -1
                var v13 int64 = 822
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    default:
        v14 := GenS1{F0: (((("hello" + "a") + ("hello" + "a")) + "foo bar") + ((("a" + ("hello" + "a")) + "") + "foo bar")), F1: 1}
        _ = v14
    }
    if true {
        var v15 uint = 863
        fmt.Println(v2, v5, v1)
        _ = v15
    }
    var v16 float64 = (v1 / (v1 + 1e10))
    fmt.Println(v16)
    if (0 >= 65535) {
        v17 := GenS1{F0: "\n\t", F1: 0}
        v18 := GenS1{F0: "", F1: 447}
        _ = v17
        _ = v18
    } else {
        v19 := GenS1{F0: (("hello" + "") + "\n\t"), F1: 1}
        v20 := []int{4294967295}
        _ = v19
        _ = v20
    }
    v5 = (v5 + v5)
    var v21 uint64 = 42
    fmt.Println(v5, v16, v21)
    delete(v4, 1)
    fmt.Println(v5)
    fmt.Println(v16, v1)
    if true {
        v5 = v5
    }
    fmt.Println(v5, v1)
    switch 725 {
    case 32768:
        fmt.Println(v16, v21)
    case -9223372036854775808:
        switch -576 {
        case 42:
            v1 -= 1.5
        default:
            if false {
                v2 = v2
                for fi22 := 0; fi22 < 3; fi22++ {
                    v2 += (v2 / ((((v2 + v2) - v2) + v2) - ((v2 - float32(v1)) / v2)))
                    var v23 uint = 4294967295
                    _ = fi22
                    _ = v23
                }
            }
        }
    default:
        v2 += (v2 - ((float32(float64(0.1)) + (v2 + -2.25)) - (1.5 * v2)))
    }
    var v24 uint8 = 1
    v5 -= v5
    fmt.Println(v24, v1)
    if (0 > -986) {
        fmt.Println(v5)
    } else {
        if false {
            v21 = v21
            var v25 uint = uint(v24)
            var v26 uint = v25
            _ = v25
            _ = v26
        } else {
            v16 = 3.14159
        }
        switch 752 {
        case 903:
            if true {
                var v27 uint32 = uint32(v21)
                _ = v27
            }
        case -9223372036854775808:
            fmt.Println(v2, v1, v24)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
