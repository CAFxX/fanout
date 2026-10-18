package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int8; F2 int32 }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0() int16 {
    var v6 string = "\n\t"
    switch v6 {
    case "\n\t":
        if (false && (9223372036854775807 != -2147483648)) {
            fmt.Println(v6)
            v6 = ((v6 + v6) + v6)
            v6 = v6
        } else {
            v7 := GenS1{F0: (127 != 454), F1: -76, F2: 9223372036854775807}
            v8 := GenS2{F0: 42}
            _ = v7
            _ = v8
        }
    case "foo bar":
        var v9 bool = (!(true && (((0.0 * 1e10) / 1.5) != float64(3.14159))))
        _ = v9
    case "hello":
        var v10 uint32 = 2147483647
        _ = v10
    }
    if true {
        if (!(32767 >= 42)) {
            v11 := []int{-9223372036854775808, -689, 9223372036854775807, -987}
            var v12 string = ((v6 + "hello") + ("foo bar" + v6))
            _ = v11
            _ = v12
        } else {
            if false {
                var v13 float64 = 3.14159
                switch 381 {
                case -2147483648:
                    v13 += v13
                }
                if true {
                    var v14 uint16 = 69
                    _ = v14
                } else {
                    v15 := GenS3{F0: 42, F1: 42, F2: -1}
                    _ = v15
                }
                _ = v13
            } else {
                switch 1 {
                case 128:
                    var v16 float64 = (float64(-2.25) - float64(((-2.25 - -2.25) * 1.5)))
                    _ = v16
                case -226:
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 float32 = -2.25
                        _ = fi17
                        _ = v18
                    }
                case 2147483647:
                    fmt.Println(v6)
                }
                v6 = "a"
                v19 := []int{-1}
                _ = v19
            }
            var v20 uint8 = 64
            var v21 float32 = (1.5 * 1.5)
            _ = v20
            _ = v21
        }
    } else {
        v6 = "foo bar"
        var v22 uint32 = 153
        _ = v22
    }
    if true {
        v6 = v6
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := []int{(-(fi23)), (^(fi23)), fi23, (fi23 % (((^(fi23))) | 1))}
            v6 = v6
            switch 32768 {
            case 1:
                var v25 float32 = (0.0 - (3.14159 + 1e10))
                _ = v25
            case -728:
                var v26 uint32 = 2147483648
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            for fi28 := 0; fi28 < 2; fi28++ {
                v6 = (v6 + (v6 + v6))
                var v29 uint8 = 2
                _ = fi28
                _ = v29
            }
            v30 := []int{(fi27 * fi27), fi27}
            _ = fi27
            _ = v30
        }
    }
    _ = v6
    return 943
}
func genFn1(p0 int64) string {
    v31 := []int{4294967295, 837}
    if len(v31) > 0 {
        v31[0] = -1
    }
    var v32 uint8 = 128
    _ = p0
    _ = v31
    _ = v32
    return ("\n\t" + "hello")
}

func main() {
    var v1 uint16 = 1
    var v2 int16 = -564
    var v3 int64 = 0
    v4 := make(map[string]bool)
    var v5 bool = (117 >= 2147483647)
    if (18446744073709551615 >= 1) {
        fmt.Println(v1)
        v33 := GenS3{F0: 1, F1: 117, F2: -2147483648}
        _ = v33
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        fmt.Println(v3)
        _ = fi34
    }
    v3 -= 0
    v35 := []int{341}
    var v36 int16 = 708
    for fi37 := 0; fi37 < 4; fi37++ {
        if (v5 || (!v5)) {
            v2 = (v36 - genFn0())
        }
        _ = fi37
    }
    v3 -= (v3 % ((0) | 1))
    for fi38 := 0; fi38 < 7; fi38++ {
        delete(v4, ("a" + "foo bar"))
        v3 = (v3 | (v3 & (^(v3))))
        _ = fi38
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 bool = true
        v5 = v40
        _ = fi39
        _ = v40
    }
    var v41 uint = 0
    for fi42 := 0; fi42 < 15; fi42++ {
        v5 = true
        _ = fi42
    }
    switch -1 {
    case 1:
        v43 := GenS4{F0: 32767, F1: 1}
        _ = v43
    case -9223372036854775808:
        var v44 string = "hello"
        _ = v44
    }
    _ = len(v4)
    for fi45 := 0; fi45 < 12; fi45++ {
        var v46 int16 = v2
        v46 = (v46 >> (((v41 >> ((v41) & 63))) & 63))
        _ = fi45
        _ = v46
    }
    var v47 bool = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v47
    _ = v5
}
