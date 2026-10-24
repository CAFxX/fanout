package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int64 }
func genFn0(p0 uint, p1 uint64) string {
    var v6 uint16 = 255
    fmt.Println(p1, v6)
    _ = p0
    _ = p1
    _ = v6
    return ""
}
func genFn1(p0 int16, p1 int8) uint64 {
    if false {
        var v7 uint16 = 42
        fmt.Println(v7, v7)
        _ = v7
    }
    p0 += (-(p0))
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint8 = 128
    var v2 uint32 = 0
    var v3 int8 = 11
    var v4 int64 = 42
    var v5 string = "a"
    v5 = ((v5 + v5) + "")
    v1 = v1
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: 305, F1: "hello"}
        var v10 int32 = 7
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 1, F1: v5}
    switch -975 {
    case -2147483648:
        switch v5 {
        case "\n\t":
            switch v5 {
            case "a":
                v12 := GenS1{F0: 1, F1: ""}
                _ = v12
            case "\n\t":
                v3 += v3
            default:
                v1 = (v1 * 213)
            }
        case "foo bar":
            if false {
                if (0 > 42) {
                    if false {
                        v13 := GenS1{F0: 65535, F1: v5}
                        _ = v13
                    }
                    if false {
                        if (!true) {
                            fmt.Println(v4, v4)
                            v3 = (v3 | (-(v3)))
                            v14 := []int{-1, 42, 32768, 128}
                            _ = v14
                        }
                    }
                } else {
                    fmt.Println(v5)
                    v15 := GenS2{F0: 0}
                    for fi16 := 0; fi16 < 7; fi16++ {
                        fmt.Println(v5, v1, v2)
                        v4 -= -1
                        v17 := GenS2{F0: 580}
                        _ = fi16
                        _ = v17
                    }
                    _ = v15
                }
            }
        case "a":
            if false {
                v18 := GenS2{F0: 449}
                fmt.Println(v2, v4, v3)
                _ = v18
            }
        default:
            var v19 float64 = -2.25
            _ = v19
        }
    case 695:
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 string = v5
            fmt.Println(v4)
            _ = fi20
            _ = v21
        }
    case 769:
        v22 := GenS1{F0: 1, F1: v5}
        _ = v22
    default:
        var v23 uint8 = v1
        _ = v23
    }
    if false {
        var v24 string = ("" + "")
        for fi25 := 0; fi25 < 19; fi25++ {
            var v26 int8 = 78
            _ = fi25
            _ = v26
        }
        _ = v24
    } else {
        var v27 uint = 4294967295
        switch (v5 + v5) {
        case "hello":
            v28 := []int{}
            _ = v28
        case "\n\t":
            var v29 int32 = -888
            _ = v29
        case "":
            v3 += (^(v3))
        default:
            var v30 uint16 = 255
            _ = v30
        }
        _ = v27
    }
    switch (v5 + ("a" + "hello")) {
    case "hello":
        for fi31 := 0; fi31 < 15; fi31++ {
            var v32 uint16 = 32767
            v1 += (v1 - v1)
            _ = fi31
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        v2 = v2
        fmt.Println(v3, v1)
        v34 := GenS3{F0: (v4 | v4), F1: (v4 | v4)}
        _ = fi33
        _ = v34
    }
    var v35 bool = true
    var v36 uint8 = (v1 | (v1 | v1))
    fmt.Println(v36)
    if v35 {
        v37 := GenS3{F0: (-(v4)), F1: v4}
        switch (v5 + "foo bar") {
        case "a":
            var v38 float64 = 0.1
            _ = v38
        case "\n\t":
            fmt.Println(v35, v35)
        }
        v39 := []int{}
        _ = v37
        _ = v39
    }
    var v40 bool = false
    v41 := GenS3{F0: v4, F1: v4}
    v42 := GenS1{F0: 515, F1: v5}
    v36 += (v1 % (((v1 % ((42) | 1))) | 1))
    switch v5 {
    case "hello":
        var v43 uint = 1
        _ = v43
    default:
        fmt.Println(v2, v36)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
