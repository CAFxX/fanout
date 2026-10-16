package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0(p0 bool, p1 uint16, p2 uint32) int8 {
    v6 := GenS1{F0: 104, F1: 90}
    switch -9223372036854775808 {
    case 32768:
        var v7 uint8 = 127
        _ = v7
    }
    if (!true) {
        var v8 bool = p0
        _ = v8
    }
    if (!(42 > 4294967295)) {
        var v9 int16 = -405
        _ = v9
    } else {
        v10 := []int{-560}
        _ = len(v10)
        _ = v10
    }
    switch ("hello" + "foo bar") {
    case "":
        switch 347 {
        case -166:
            v11 := GenS1{F0: 1, F1: -126}
            _ = v11
        case 127:
            p2 += 72
        case -564:
            var v12 uint32 = (p2 + p2)
            _ = v12
        default:
            v13 := GenS1{F0: -2147483648, F1: -91}
            _ = v13
        }
    case "foo bar":
        v14 := GenS1{F0: 34, F1: 1}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 int64) uint32 {
    if (false || false) {
        switch 127 {
        case 32768:
            v15 := []int{42, -2147483648, -9223372036854775808, 4294967295}
            _ = v15
        default:
            for fi16 := 0; fi16 < 17; fi16++ {
                p0 = p0
                var v17 uint32 = 65535
                _ = fi16
                _ = v17
            }
        }
    } else {
        if (127 != 2) {
            var v18 int = 32768
            var v19 uint64 = 4294967295
            _ = v18
            _ = v19
        }
    }
    p0 += p0
    _ = p0
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "a"
    var v3 float32 = (0.0 / 1.5)
    var v4 float64 = 3.14159
    var v5 uint8 = 255
    var v20 float32 = 1.5
    var v21 uint32 = 0
    var v22 float64 = ((0.1 * v4) + -2.25)
    switch -1 {
    case 9223372036854775807:
        var v23 int16 = 32767
        _ = v23
    case 32768:
        if (0 != v5) {
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := []int{fi24, -9223372036854775808, (fi24 ^ -1), fi24}
                if (-128 <= genFn0(false, 806, genFn1(0))) {
                    _ = len(v1)
                }
                var v26 uint64 = 18446744073709551615
                _ = fi24
                _ = v25
                _ = v26
            }
            fmt.Println(v22)
        } else {
            v3 -= (v3 / 3.14159)
            v20 += v20
            if (v2 == "hello") {
                var v27 uint16 = 2
                if true {
                    fmt.Println(v3)
                    var v28 int16 = -32768
                    for fi29 := 0; fi29 < 16; fi29++ {
                        var v30 float64 = (3.14159 / 1.5)
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
                _ = v27
            }
        }
    }
    if false {
        for fi31 := 0; fi31 < 3; fi31++ {
            v20 = ((v20 / v3) - (v20 + (v20 - (3.14159 / 1.0))))
            v32 := []int{(fi31 % (((^(fi31))) | 1)), fi31}
            var v33 int = 127
            _ = fi31
            _ = v32
            _ = v33
        }
        v2 = v2
    }
    fmt.Println(v4, v2, v4)
    switch 32768 {
    case -1:
        v5 = v5
    case -2147483648:
        v34 := GenS1{F0: 0, F1: 42}
        _ = v34
    }
    var v35 uint64 = 2
    if false {
        var v36 uint32 = (v21 ^ (v21 % ((v21) | 1)))
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
