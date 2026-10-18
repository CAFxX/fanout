package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint32 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint16, p2 int64) uint8 {
    v6 := GenS1{F0: 1, F1: (!(1 > 75)), F2: 2}
    v7 := GenS2{F0: "foo bar", F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int8 = 109
    var v2 float32 = (1e10 * (1.5 + (1e10 / 0.1)))
    v3 := make(map[uint32]bool)
    var v4 int32 = 0
    v5 := make(map[int]uint32)
    if (775 == 32767) {
        fmt.Println(v2)
        v8 := GenS2{F0: "a", F1: 100}
        v4 = v4
        _ = v8
    }
    if (!(!false)) {
        if false {
            v9 := GenS3{F0: 32767}
            v10 := GenS4{F0: 32767}
            switch 520 {
            case 128:
                v4 += 672
            case 42:
                var v11 uint16 = 0
                _ = v11
            default:
                v12 := []int{2147483647, 9223372036854775807, 2147483647, 42}
                _ = v12
            }
            _ = v10
            _ = v9
        } else {
            if true {
                if (!(!true)) {
                    if true {
                        var v13 uint = 0
                        _ = v13
                    }
                    if true {
                        fmt.Println(v2)
                        v4 -= (^(v4))
                        v14 := GenS1{F0: (v1 | (v1 + v1)), F1: (0 >= 460), F2: 4294967295}
                        _ = v14
                    }
                }
                switch (("hello" + "\n\t") + "hello") {
                case "\n\t":
                    var v15 float32 = 0.1
                    _ = v15
                case "":
                    switch "foo bar" {
                    case "a":
                        v1 = v1
                    case "foo bar":
                        var v16 int16 = 1
                        _ = v16
                    case "hello":
                        var v17 bool = (0.1 >= 0.0)
                        _ = v17
                    }
                case "a":
                    v4 -= (v4 + -1)
                }
            }
        }
        v18 := GenS3{F0: -323}
        _ = v18
    } else {
        var v19 int16 = 81
        if true {
            fmt.Println(v4, v2, v1)
            v4 = v4
            fmt.Println(v4)
        }
        switch "foo bar" {
        case "foo bar":
            v20 := []int{32768, -9223372036854775808}
            _ = v20
        default:
            v1 += v1
        }
        _ = v19
    }
    switch 766 {
    case 9223372036854775807:
        fmt.Println(v2)
    case 1:
        if ("\n\t" > "foo bar") {
            _ = len(v5)
        } else {
            v1 += (v1 + v1)
            v5[42] = 1
            v21 := []int{128}
            _ = v21
        }
    case -1:
        var v22 string = ""
        _ = v22
    default:
        v2 += 1e10
    }
    if false {
        fmt.Println(v1, v4, v2)
        switch "a" {
        case "a":
            var v23 int8 = v1
            _ = v23
        default:
            fmt.Println(v2, v4, v2)
        }
    }
    fmt.Println(v1)
    v24 := GenS3{F0: 0}
    v25 := []int{32768}
    var v26 uint64 = 398
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
