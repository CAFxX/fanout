package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 string; F1 int; F2 string }
func genFn0(p0 int) bool {
    var v7 uint16 = 255
    v8 := GenS1{F0: 2}
    var v9 int16 = 32767
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := []uint32{}
    var v2 int64 = 172
    var v3 int16 = -1
    v4 := make(map[uint32]bool)
    v5 := []string{"\n\t", "hello", "foo bar", "hello"}
    v6 := []int{int(v2), 128, 42, -1}
    if false {
        if false {
            fmt.Println(v3, v2)
        }
        v10 := GenS2{F0: (-102 < 47)}
        _ = v10
    } else {
        if (!(!true)) {
            v2 += 0
        }
        if false {
            fmt.Println(v2, v2)
        }
        v3 = (v3 / ((355) | 1))
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v11 := v4[955]
        _ = v11
    default:
        v4[416] = false
    }
    var v12 uint8 = 118
    v2 -= (v2 - -879)
    if false {
        v3 = (v3 & (v3 % ((v3) | 1)))
        fmt.Println(v3, v2)
        fmt.Println(v3)
    }
    var v13 int8 = 127
    _ = len(v5)
    v14 := GenS2{F0: (!(true && true))}
    var v15 int32 = 2147483647
    v16 := GenS3{F0: (v15 ^ (-(v15)))}
    var v17 int16 = 1
    switch "foo bar" {
    case "foo bar":
        v1 = append(v1, 442)
    case "a":
        v18 := GenS2{F0: true}
        _ = v18
    case "":
        v4[42] = true
    }
    v19 := GenS4{F0: (v2 << ((329) & 63)), F1: (v12 % (((v12 - v12)) | 1)), F2: 1}
    switch ("a" + "hello") {
    case "":
        if len(v5) > 0 {
            v5[0] = "a"
        }
    case "hello":
        switch 707 {
        case 128:
            switch "a" {
            case "foo bar":
                for fi20 := 0; fi20 < 12; fi20++ {
                    fmt.Println(fi20)
                    var v21 int16 = (v3 / ((-292) | 1))
                    v22 := GenS5{F0: 65535, F1: 1}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
        default:
            var v23 uint16 = 897
            _ = v23
        }
    default:
        if false {
            if (388 == 18446744073709551615) {
                v3 -= v3
            } else {
                for fi24 := 0; fi24 < 9; fi24++ {
                    v12 += (v12 & (v12 + (v12 / ((v12) | 1))))
                    v25 := GenS3{F0: 2147483647}
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v15)
                    }
                    _ = fi24
                    _ = v25
                }
                v3 += (-(v17))
            }
            switch "\n\t" {
            case "a":
                switch 2147483647 {
                case 740:
                    v26 := GenS6{F0: -2147483648, F1: 4294967295}
                    _ = v26
                case -782:
                    v27 := GenS5{F0: 65535, F1: 255}
                    _ = v27
                case 182:
                    v28 := GenS6{F0: -1, F1: 2}
                    _ = v28
                }
            case "":
                var v29 float32 = 1.5
                _ = v29
            case "\n\t":
                var v30 int32 = (-(v15))
                _ = v30
            default:
                v2 = -415
            }
            v31 := GenS7{F0: ("\n\t" + (((("a" + "foo bar") + "\n\t") + "\n\t") + ("" + ""))), F1: 32767, F2: ("hello" + "\n\t")}
            _ = v31
        } else {
            delete(v4, 944)
            v32 := GenS3{F0: int32(v13)}
            var v33 bool = (2147483647 != 482)
            _ = v32
            _ = v33
        }
    }
    var v34 int64 = v2
    var v35 int = 32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
