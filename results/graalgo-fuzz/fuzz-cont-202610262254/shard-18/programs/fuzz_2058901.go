package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 int64) int16 {
    p0 = (^(p0))
    fmt.Println(p0)
    v7 := []int{-9223372036854775808}
    switch "" {
    case "foo bar":
        v8 := GenS1{F0: 1, F1: 99, F2: (p0 | 1)}
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        v7 = append(v7, fi9)
        p0 -= (p0 >> ((18446744073709551615) & 63))
        fmt.Println(p0)
        _ = fi9
    }
    _ = p0
    _ = v7
    return -369
}
func genFn1(p0 int16, p1 int8) int16 {
    if false {
        v10 := []int{}
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 int32 = 42
            fmt.Println(p0)
            switch (fi11 * fi11) {
            case 32767:
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 int64 = -1
                    var v15 uint16 = 32767
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    switch "foo bar" {
    case "hello":
        v16 := []int{32768, 385}
        _ = v16
    case "":
        p1 = 127
    }
    if false {
        var v17 uint32 = 65535
        var v18 int32 = -735
        v19 := GenS1{F0: (p0 - -1), F1: (p1 >> ((0) & 63)), F2: -9223372036854775808}
        _ = v17
        _ = v18
        _ = v19
    } else {
        switch "" {
        case "hello":
            var v20 int = -586
            _ = v20
        case "":
            p1 = p1
        case "a":
            var v21 int32 = -2147483648
            _ = v21
        default:
            v22 := []int{32767, 457, 2147483647}
            _ = v22
        }
    }
    v23 := []int{4294967295, -960, 32768}
    for fi24 := 0; fi24 < 15; fi24++ {
        if true {
            v25 := GenS2{F0: 983}
            _ = v25
        }
        var v26 uint32 = 2
        p1 = 42
        _ = fi24
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v23
    return (p0 * -1)
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint8 = 203
    v3 := make(map[string]string)
    var v4 int32 = -1
    v5 := make(map[uint32]bool)
    var v6 uint32 = (^(v1))
    for fi27 := 0; fi27 < 1; fi27++ {
        fmt.Println(v4, v2)
        var v28 int8 = 100
        for fi29 := 0; fi29 < 4; fi29++ {
            v30 := GenS3{F0: v28, F1: 42}
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    v2 -= 77
    switch "hello" {
    case "hello":
        if ((0.1 - -2.25) <= float64((((-2.25 + 1.5) + -2.25) + 0.1))) {
            v4 -= (-(v4))
        } else {
            var v31 uint32 = v6
            if true {
                v4 = (v4 | -1)
            }
            fmt.Println(v31, v2, v1)
            _ = v31
        }
    case "a":
        v32 := []int{42, 1, 2147483647}
        _ = v32
    case "\n\t":
        for fi33 := 0; fi33 < 10; fi33++ {
            v34 := GenS3{F0: -1, F1: -1}
            _ = fi33
            _ = v34
        }
    }
    fmt.Println(v4, v4)
    var v35 uint8 = v2
    fmt.Println(v1, v4)
    for fi36 := 0; fi36 < 15; fi36++ {
        switch "a" {
        case "foo bar":
            var v37 int64 = -1
            _ = v37
        case "a":
            for fi38 := 0; fi38 < 8; fi38++ {
                v2 = 127
                switch 128 {
                case -337:
                    v1 += (v1 << ((42) & 63))
                default:
                    v5[(v6 - (v1 & v1))] = (false && false)
                }
                delete(v3, (("" + ("foo bar" + "")) + ("a" + (("" + "\n\t") + "\n\t"))))
                v39 := GenS4{F0: "foo bar"}
                _ = fi38
                _ = v39
            }
        default:
            v6 += (v6 & v6)
        }
        _ = fi36
    }
    _ = len(v5)
    v40 := GenS4{F0: "hello"}
    var v41 int = 127
    var v42 int = v41
    if (32768 <= (^(v41))) {
        var v43 int = (-(v42))
        _ = v43
    }
    var v44 uint = uint(v6)
    v45 := []int{(v41 % ((int(v4)) | 1)), (-(v41)), -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
