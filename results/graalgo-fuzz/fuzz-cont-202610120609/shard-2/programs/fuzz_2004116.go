package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0() int {
    v7 := []int{-2147483648}
    var v8 string = "a"
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int) uint {
    p0 = p0
    p0 -= (-(p0))
    if (42 >= 4294967295) {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int8 = 127
            _ = fi9
            _ = v10
        }
        var v11 uint16 = 57
        v12 := []int{(p0 | p0), 1, 2147483647}
        _ = v11
        _ = v12
    } else {
        v13 := GenS1{F0: 1, F1: 1, F2: 2}
        if true {
            if (255 == 65535) {
                var v14 uint32 = 2
                _ = v14
            }
            p0 = (p0 + genFn0())
        }
        _ = v13
    }
    switch p0 {
    case -705:
        fmt.Println(p0)
    case 32767:
        switch genFn0() {
        case 1:
            p0 -= 105
        case -2147483648:
            if true {
                p0 += p0
                var v15 int64 = -9223372036854775808
                _ = v15
            } else {
                for fi16 := 0; fi16 < 7; fi16++ {
                    fmt.Println(fi16, p0)
                    _ = fi16
                }
            }
        case 232:
            v17 := []int{(p0 % (((-(p0))) | 1)), p0, p0}
            _ = v17
        }
    case -9223372036854775808:
        for fi18 := 0; fi18 < 8; fi18++ {
            fmt.Println(p0, p0)
            var v19 int64 = -1
            _ = fi18
            _ = v19
        }
    default:
        if false {
            v20 := []int{}
            if (127 != -1) {
                var v21 int64 = -608
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                var v22 float64 = 0.0
                _ = v21
                _ = v22
            } else {
                v23 := GenS2{F0: 18446744073709551615, F1: 0}
                switch p0 {
                case 4294967295:
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                case -9223372036854775808:
                    var v24 int32 = -1
                    _ = v24
                }
                p0 = p0
                _ = v23
            }
            _ = v20
        }
    }
    switch "" {
    case "":
        if false {
            if false {
                for fi25 := 0; fi25 < 12; fi25++ {
                    fmt.Println(fi25, p0)
                    var v26 int64 = 42
                    switch 4294967295 {
                    case 0:
                        fmt.Println(fi25)
                    case 2147483647:
                        v27 := GenS2{F0: 18446744073709551615, F1: 255}
                        _ = v27
                    case 127:
                        p0 -= genFn0()
                    }
                    _ = fi25
                    _ = v26
                }
            } else {
                p0 -= 986
            }
            if true {
                switch "foo bar" {
                case "":
                    p0 = genFn0()
                default:
                    p0 -= p0
                }
                v28 := GenS1{F0: 18446744073709551615, F1: 1, F2: 0}
                v29 := GenS1{F0: 0, F1: 0, F2: 206}
                _ = v28
                _ = v29
            } else {
                v30 := []int{p0, (p0 << ((82) & 63)), p0}
                _ = v30
            }
        }
    case "hello":
        v31 := []int{p0, p0}
        _ = v31
    case "\n\t":
        v32 := GenS3{F0: 18446744073709551615}
        _ = v32
    }
    _ = p0
    return 377
}

func main() {
    var v1 string = ("hello" + ("" + ""))
    var v2 int = 59
    v3 := make(map[string]int)
    var v4 string = ((v1 + "foo bar") + "hello")
    var v5 float64 = (1.5 / -2.25)
    v6 := []string{"\n\t", v1, v4}
    var v33 int = v2
    var v34 string = "hello"
    var v35 uint32 = 804
    var v36 uint16 = 1
    if false {
        var v37 string = (v4 + "")
        var v38 uint16 = v36
        _ = v37
        _ = v38
    }
    v39 := GenS4{F0: 1, F1: 18446744073709551615}
    switch v33 {
    case 1:
        v40 := GenS4{F0: 42, F1: 42}
        _ = v40
    }
    v41 := GenS1{F0: 1, F1: v36, F2: 52}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
