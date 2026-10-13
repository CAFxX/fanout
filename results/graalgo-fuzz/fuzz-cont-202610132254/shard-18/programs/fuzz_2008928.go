package main

import (
    "fmt"
)

func genFn0(p0 int32, p1 int8, p2 int32) int32 {
    switch 9223372036854775807 {
    case 32768:
        if (!false) {
            var v7 uint16 = 20
            _ = v7
        } else {
            fmt.Println(p0)
            p1 = p1
            p2 += -813
        }
    case -801:
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int64 = 0
            p2 -= -1
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        p0 += p2
        var v11 string = (("foo bar" + "foo bar") + "a")
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 uint32 = 2147483648
        for fi14 := 0; fi14 < 17; fi14++ {
            if (661 == 1) {
                p0 = 2147483647
                var v15 int = fi12
                var v16 int32 = (p2 ^ 0)
                _ = v15
                _ = v16
            }
            p1 += 127
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    v17 := []int{42, 32767, 384}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{4294967295, -1, 4294967295}
    v3 := make(map[int]int)
    var v4 uint8 = 255
    var v5 uint32 = 2147483647
    v6 := make(map[string]uint32)
    var v18 uint8 = (v4 / (((v4 >> ((18446744073709551615) & 63))) | 1))
    var v19 int64 = 148
    if (0 == genFn0(-865, 88, 1)) {
        var v20 int16 = 1000
        var v21 bool = (((1 == -64) || ((!((69 != 255) || (42 != 0))) || true)) && true)
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 uint = 164
            _ = fi22
            _ = v23
        }
        switch "" {
        case "hello":
            switch "hello" {
            case "foo bar":
                v5 += 1
            case "":
                var v24 uint = 469
                _ = v24
            default:
                var v25 uint = 340
                _ = v25
            }
        case "\n\t":
            v19 += 1
        default:
            fmt.Println(v4, v19)
        }
        fmt.Println(v5)
    }
    if false {
        var v26 float32 = ((0.0 * (-2.25 * 3.14159)) / 1.0)
        fmt.Println(v26, v19, v19)
        _ = v26
    }
    fmt.Println(v18, v19)
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v19 = (v19 | (-(v19)))
    v5 = v5
    v18 += v4
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    v18 = (v18 + (v18 * v18))
    v19 += v19
    for fi27 := 0; fi27 < 14; fi27++ {
        if false {
            var v28 float64 = 1e10
            v28 += -2.25
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v28
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
