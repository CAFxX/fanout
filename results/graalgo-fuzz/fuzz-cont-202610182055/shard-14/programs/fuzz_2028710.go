package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint; F2 uint }
func genFn0() int {
    if true {
        v5 := GenS1{F0: 42, F1: 0}
        var v6 int = -9223372036854775808
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := GenS1{F0: 42, F1: 182}
            v6 = 875
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS1{F0: 1, F1: 2}
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        _ = v5
        _ = v6
    } else {
        if true {
            var v11 int16 = -1
            v11 = (^(v11))
            _ = v11
        } else {
            var v12 uint16 = 376
            _ = v12
        }
    }
    v13 := []int{2147483647, 19, -137, 32767}
    _ = v13
    return 1
}

func main() {
    v1 := []string{}
    var v2 int64 = 735
    var v3 string = (((("" + (("a" + (("hello" + "") + ("a" + ("foo bar" + ("" + "\n\t"))))) + "a")) + ("" + "hello")) + "a") + ("foo bar" + "hello"))
    v4 := make(map[uint32]string)
    var v14 int16 = -1
    if (0.0 != (((0.0 * 1e10) + -2.25) - (1.5 - (float64(3.14159) + -2.25)))) {
        v3 = "\n\t"
        v3 = v3
        var v15 bool = false
        _ = v15
    } else {
        var v16 int32 = 40
        v3 = (("hello" + ("foo bar" + "\n\t")) + v3)
        switch 876 {
        case 32767:
            v17 := GenS1{F0: 155, F1: 788}
            _ = v17
        }
        _ = v16
    }
    var v18 int32 = 1
    if false {
        switch 1 {
        case -1:
            if (1 != 2) {
                var v19 uint32 = 28
                for fi20 := 0; fi20 < 4; fi20++ {
                    fmt.Println(v3, v3, v2)
                    if false {
                        for fi21 := 0; fi21 < 14; fi21++ {
                            _ = len(v4)
                            _ = fi21
                        }
                    } else {
                        v1 = append(v1, "foo bar")
                        var v22 uint64 = 1
                        _ = v22
                    }
                    _ = fi20
                }
                v14 = (^(v14))
                _ = v19
            }
        case 4294967295:
            switch genFn0() {
            case -478:
                v23 := GenS2{F0: v18, F1: 18446744073709551615, F2: 1}
                _ = v23
            case 355:
                v24 := GenS1{F0: 0, F1: 65535}
                _ = v24
            }
        case 0:
            v25 := GenS3{F0: 33, F1: 2, F2: 707}
            _ = v25
        }
        if (!true) {
            switch 680 {
            case 127:
                var v26 int8 = 127
                _ = v26
            case 9223372036854775807:
                v4[2147483648] = ("hello" + v3)
            }
            if (0 == 2147483648) {
                v27 := GenS1{F0: 2, F1: 42}
                for fi28 := 0; fi28 < 4; fi28++ {
                    v29 := GenS1{F0: 4294967295, F1: 65535}
                    for fi30 := 0; fi30 < 12; fi30++ {
                        fmt.Println(v3)
                        _ = fi30
                    }
                    _ = fi28
                    _ = v29
                }
                switch -9223372036854775808 {
                case -424:
                    v18 = v18
                }
                _ = v27
            }
        }
    } else {
        v18 += v18
        switch v3 {
        case "a":
            v18 = (v18 + v18)
        }
    }
    fmt.Println(v2, v3, v3)
    switch -1 {
    case 2147483647:
        var v31 int8 = -128
        _ = v31
    case -95:
        for fi32 := 0; fi32 < 20; fi32++ {
            v2 += 9223372036854775807
            for fi33 := 0; fi33 < 3; fi33++ {
                _ = len(v4)
                _ = fi33
            }
            fmt.Println(v3, v2, v18)
            _ = fi32
        }
    }
    if true {
        for fi34 := 0; fi34 < 5; fi34++ {
            v35 := GenS2{F0: 42, F1: 377, F2: 2}
            v14 -= (v14 + (v14 ^ -32768))
            delete(v4, 0)
            _ = fi34
            _ = v35
        }
    }
    v36 := GenS1{F0: 159, F1: 2}
    _ = len(v1)
    fmt.Println(v14, v18)
    v37 := GenS1{F0: 18446744073709551615, F1: 631}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
