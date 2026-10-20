package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint8 }
func genFn0(p0 int64, p1 int16) int8 {
    switch 9223372036854775807 {
    case 32768:
        switch 804 {
        case 32768:
            fmt.Println(p0)
        case 32767:
            fmt.Println(p1, p0)
        case 1:
            var v7 int32 = 2147483647
            _ = v7
        }
    case -2147483648:
        v8 := GenS1{F0: true}
        _ = v8
    case 124:
        if false {
            v9 := []int{-343, 9223372036854775807}
            v10 := GenS2{F0: 0, F1: 42}
            for fi11 := 0; fi11 < 5; fi11++ {
                if len(v9) > 0 {
                    v9[0] = fi11
                }
                _ = fi11
            }
            _ = v10
            _ = v9
        }
    default:
        switch "foo bar" {
        case "\n\t":
            switch ("a" + "") {
            case "hello":
                var v12 uint32 = 985
                _ = v12
            case "":
                fmt.Println(p1)
            case "\n\t":
                switch 128 {
                case 331:
                    switch "" {
                    case "":
                        v13 := GenS2{F0: 1, F1: 128}
                        _ = v13
                    case "\n\t":
                        p0 = p0
                    case "hello":
                        fmt.Println(p1, p0)
                    default:
                        v14 := []int{4294967295}
                        _ = v14
                    }
                }
            default:
                var v15 uint64 = 151
                _ = v15
            }
        case "a":
            p0 -= 808
        case "foo bar":
            var v16 float64 = (0.1 - ((1.5 + -2.25) + 0.1))
            _ = v16
        default:
            p1 += p1
        }
    }
    if false {
        var v17 float32 = 3.14159
        _ = v17
    } else {
        fmt.Println(p1, p0)
        var v18 int64 = p0
        var v19 float64 = -2.25
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    return 100
}
func genFn1(p0 int8, p1 bool) int8 {
    for fi20 := 0; fi20 < 9; fi20++ {
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 float32 = (1.5 / (0.1 * 0.1))
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    p0 = (p0 << ((65535) & 63))
    fmt.Println(p1, p1)
    if (127 != 255) {
        v23 := GenS2{F0: -9223372036854775808, F1: 255}
        switch -620 {
        case -716:
            fmt.Println(p1, p1)
        default:
            p1 = (1 != 0)
        }
        _ = v23
    } else {
        if p1 {
            var v24 uint32 = 65535
            fmt.Println(p1)
            v25 := []int{-765, 128, 744}
            _ = v24
            _ = v25
        }
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    var v5 uint32 = 178
    v6 := make(map[int]bool)
    v5 += (v5 - 2147483647)
    switch 42 {
    case 128:
        if (42 > 65535) {
            var v26 int16 = 42
            _ = v26
        } else {
            v27 := GenS3{F0: 417, F1: genFn0(v1, 346), F2: (v5 + 65535)}
            switch ("a" + "\n\t") {
            case "hello":
                var v28 uint16 = 351
                _ = v28
            case "foo bar":
                v1 -= (-(v1))
            case "\n\t":
                if false {
                    fmt.Println(v1, v1)
                } else {
                    fmt.Println(v1, v1)
                }
            default:
                var v29 float32 = float32(0.1)
                _ = v29
            }
            var v30 string = ("a" + "a")
            _ = v27
            _ = v30
        }
    case -9223372036854775808:
        v1 -= v1
    case 1:
        delete(v6, 2147483647)
    }
    var v31 int32 = 745
    v32 := GenS1{F0: (639 >= -91)}
    for fi33 := 0; fi33 < 10; fi33++ {
        v1 = 42
        v5 = (v5 & 454)
        fmt.Println(v1, fi33)
        _ = fi33
    }
    var v34 int32 = -1
    v35 := GenS4{F0: v5, F1: 127, F2: 255}
    fmt.Println(v5, v1, v1)
    var v36 int32 = 1
    fmt.Println(v31, v31, v31)
    v31 = v31
    v37 := GenS3{F0: 2, F1: -1, F2: v5}
    if false {
        var v38 int16 = 0
        _ = v38
    }
    var v39 string = ((("foo bar" + "hello") + ("\n\t" + "foo bar")) + "hello")
    switch v39 {
    case "":
        v40 := []int{128, -2147483648}
        _ = v40
    default:
        v3[(v39 + "foo bar")] = ""
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
