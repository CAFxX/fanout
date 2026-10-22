package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint; F1 int64; F2 string }
type GenS3 struct { F0 int8; F1 uint64; F2 uint }
type GenS4 struct { F0 bool }
func genFn0(p0 int32, p1 uint8, p2 int8) uint32 {
    var v6 uint16 = 255
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 = 1
        if (!false) {
            v6 += (v6 | 65535)
        }
        fmt.Println(v6, v6, p0)
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := []int{127, (fi8 << ((2) & 63)), fi8, fi8}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi8
        _ = v9
    }
    if false {
        v10 := GenS1{F0: 65535, F1: 2147483647, F2: p1}
        _ = v10
    } else {
        v11 := []int{-744, 0, -1, 2147483647}
        _ = v11
    }
    switch ("" + "a") {
    case "a":
        v6 = v6
    default:
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(v6, fi12, p2)
            for fi13 := 0; fi13 < 5; fi13++ {
                if false {
                    v14 := []int{(fi13 & 2147483647), 32768}
                    var v15 float32 = (0.1 + 3.14159)
                    p1 += (p1 & 127)
                    _ = v14
                    _ = v15
                }
                if (!(true || (false && true))) {
                    var v16 float64 = 0.1
                    v17 := []int{fi13, fi12, fi13}
                    _ = v16
                    _ = v17
                } else {
                    v6 -= (v6 + 906)
                    v18 := []int{(fi12 + 1), (fi13 + 127), fi12}
                    var v19 float32 = 1e10
                    _ = v18
                    _ = v19
                }
                _ = fi13
            }
            _ = fi12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 int16 = -32768
    var v2 bool = false
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 bool = v2
    switch ("hello" + "hello") {
    case "\n\t":
        switch int(v1) {
        case -1:
            var v20 int8 = 42
            _ = v20
        case 127:
            var v21 int16 = (v1 | v1)
            _ = v21
        case 339:
            v1 -= v1
        default:
            v22 := GenS2{F0: 0, F1: 9223372036854775807, F2: "a"}
            _ = v22
        }
    case "a":
        var v23 uint64 = 96
        _ = v23
    }
    v24 := GenS3{F0: 127, F1: 1, F2: 18446744073709551615}
    if (!(float64(0.1) > (float64(-2.25) * (3.14159 + ((0.1 / 1.0) / float64(1.5)))))) {
        var v25 int8 = 0
        fmt.Println(v1, v5, v2)
        var v26 bool = (!((("" + ("foo bar" + "")) + "\n\t") != ""))
        _ = v25
        _ = v26
    } else {
        if true {
            v2 = false
            v27 := []int{32767, -1}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 15; fi28++ {
                v2 = v5
                var v29 int8 = 127
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 2; fi30++ {
                v2 = v5
                _ = fi30
            }
            var v31 uint = 255
            _ = v31
        }
        fmt.Println(v5)
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 float64 = 1.5
        var v34 int64 = -1
        var v35 int64 = v34
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    v1 += v1
    v36 := GenS4{F0: (0 >= 119)}
    v5 = v2
    var v37 float64 = (((0.1 - 1e10) - float64(1e10)) / 0.1)
    var v38 uint64 = 964
    fmt.Println(v5, v37)
    if ((v37 / v37) < v37) {
        v3[0] = (-1 > 100)
    } else {
        for fi39 := 0; fi39 < 15; fi39++ {
            v40 := []int{}
            _ = fi39
            _ = v40
        }
        var v41 float64 = (v37 / v37)
        _ = v41
    }
    v5 = ((false || (v1 < (v1 + (v1 / ((1) | 1))))) && ((v37 / (0.1 - (3.14159 * (1.5 + (3.14159 - v37))))) >= v37))
    v1 += v1
    if (("hello" + "a") < (("a" + ("" + ("foo bar" + ("foo bar" + "foo bar")))) + (("\n\t" + "\n\t") + "hello"))) {
        v42 := GenS4{F0: (v38 > 1)}
        _ = v42
    }
    v43 := GenS1{F0: 65535, F1: -2147483648, F2: 115}
    fmt.Println(v38)
    v44 := GenS3{F0: 1, F1: 4294967295, F2: 213}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
