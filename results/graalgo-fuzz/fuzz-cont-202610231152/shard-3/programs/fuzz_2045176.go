package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int64; F1 int32 }
func genFn0(p0 uint16) int64 {
    p0 -= 65535
    switch 368 {
    case 42:
        p0 += p0
    case 32767:
        if false {
            v7 := []int{32767, -114, 32768}
            fmt.Println(p0)
            p0 = p0
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 float64 = ((1.5 * 0.0) - float64((((float32((0.0 * 0.1)) * float32(0.0)) - 3.14159) - 1e10)))
            p0 = (p0 + p0)
            switch "foo bar" {
            case "foo bar":
                v10 := GenS1{F0: 1, F1: 0}
                _ = v10
            case "\n\t":
                p0 -= (p0 / ((p0) | 1))
            }
            _ = fi8
            _ = v9
        }
    }
    if false {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 string = (("hello" + "hello") + "hello")
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 1, F1: 18446744073709551615}
        switch ("\n\t" + ((("" + "") + "a") + "a")) {
        case "foo bar":
            var v14 int8 = 100
            _ = v14
        case "":
            if false {
                v15 := []int{}
                _ = v15
            } else {
                if true {
                    v16 := []int{-2147483648, -1}
                    p0 = (p0 - 32767)
                    _ = v16
                }
                switch -788 {
                case 128:
                    p0 -= (^(p0))
                default:
                    switch 0 {
                    case -9223372036854775808:
                        v17 := []int{-2147483648, 32768, -293}
                        _ = v17
                    case 128:
                        var v18 int = 1
                        _ = v18
                    case 9223372036854775807:
                        var v19 int8 = 1
                        _ = v19
                    }
                }
                if false {
                    fmt.Println(p0)
                } else {
                    fmt.Println(p0)
                }
            }
        default:
            v20 := []int{-567, -584}
            _ = v20
        }
        _ = v13
    } else {
        v21 := GenS2{F0: 255, F1: -2147483648}
        if (-110 != -128) {
            fmt.Println(p0)
            p0 = (p0 / ((p0) | 1))
        }
        _ = v21
    }
    p0 = (^(p0))
    switch ("a" + "a") {
    case "foo bar":
        switch (("foo bar" + "\n\t") + "foo bar") {
        case "":
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 float32 = 3.14159
                _ = fi22
                _ = v23
            }
        case "hello":
            if true {
                fmt.Println(p0)
            } else {
                for fi24 := 0; fi24 < 14; fi24++ {
                    var v25 int64 = -9223372036854775808
                    _ = fi24
                    _ = v25
                }
                v26 := []int{1, 128, 9223372036854775807, 0}
                fmt.Println(p0)
                _ = v26
            }
        case "foo bar":
            v27 := GenS3{F0: 255, F1: p0, F2: 2}
            _ = v27
        }
    default:
        v28 := []int{774, 32767, 1}
        _ = v28
    }
    _ = p0
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{(("\n\t" + "a") + "\n\t"), ("\n\t" + (("foo bar" + "\n\t") + "foo bar")), "foo bar"}
    var v3 uint64 = 2
    var v4 int = -2147483648
    var v5 uint64 = (v3 | v3)
    v6 := []uint32{881, 2147483647, 1}
    fmt.Println(v5, v5)
    var v29 uint8 = 2
    var v30 string = "hello"
    fmt.Println(v30, v3)
    for fi31 := 0; fi31 < 3; fi31++ {
        v5 = v3
        _ = fi31
    }
    fmt.Println(v29, v30)
    v30 = v30
    if true {
        var v32 string = ((("" + v30) + "foo bar") + "\n\t")
        fmt.Println(v5, v4)
        _ = v32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        v30 = v30
        v34 := GenS4{F0: 1, F1: 649}
        _ = fi33
        _ = v34
    }
    var v35 uint16 = 0
    var v36 uint = 571
    fmt.Println(v4, v29)
    v37 := GenS2{F0: (v36 ^ (v36 * 858)), F1: 470}
    if len(v6) > 0 {
        v6[0] = 1
    }
    if false {
        fmt.Println(v5, v3, v29)
    } else {
        var v38 float32 = -2.25
        _ = v38
    }
    v4 = (v4 - (v4 / ((int(v35)) | 1)))
    v39 := GenS1{F0: uint32(v29), F1: v5}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
