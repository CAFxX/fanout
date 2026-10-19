package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint; F1 int32; F2 uint64 }
type GenS3 struct { F0 int32; F1 bool; F2 int32 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 bool, p1 int64, p2 int8) int16 {
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(p0)
        switch "hello" {
        case "foo bar":
            fmt.Println(p2, p0)
        case "hello":
            var v7 uint = 42
            _ = v7
        default:
            if p0 {
                v8 := []int{}
                fmt.Println(p2, p1, fi6)
                var v9 int64 = 630
                _ = v8
                _ = v9
            }
        }
        v10 := GenS1{F0: 42, F1: 18446744073709551615}
        _ = fi6
        _ = v10
    }
    fmt.Println(p2, p0)
    v11 := []int{636, 127, 9223372036854775807, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}
func genFn1() uint64 {
    var v12 string = ("foo bar" + ("foo bar" + "hello"))
    var v13 int = 32768
    v14 := []int{}
    v13 -= 9223372036854775807
    var v15 int64 = 9223372036854775807
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 601
}

func main() {
    var v1 uint64 = 258
    var v2 bool = false
    var v3 int32 = -2147483648
    v4 := make(map[string]int)
    var v5 uint64 = v1
    v16 := GenS1{F0: 29, F1: 2}
    v5 = v5
    v17 := GenS2{F0: 42, F1: v3, F2: 0}
    if (!(("hello" + "\n\t") > "a")) {
        v18 := GenS3{F0: (v3 | (^(v3))), F1: v2, F2: v3}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 uint32 = 2147483648
            _ = fi19
            _ = v20
        }
        if v2 {
            switch 4294967295 {
            case 127:
                switch ("" + "\n\t") {
                case "":
                    _ = len(v4)
                case "\n\t":
                    var v21 uint32 = 1
                    _ = v21
                }
            case 42:
                v5 -= v5
            default:
                var v22 float64 = 3.14159
                _ = v22
            }
        }
        var v23 float32 = 0.1
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            switch fi24 {
            case 128:
                if (42 == int8(fi24)) {
                    v26 := []int{(fi24 | fi25), (fi25 ^ 530)}
                    if (v2 || (127 > 100)) {
                        fmt.Println(fi25, fi25, fi24)
                        var v27 int8 = 0
                        _ = v27
                    } else {
                        v28 := GenS2{F0: 0, F1: 2147483647, F2: (v5 >> ((1) & 63))}
                        v2 = false
                        var v29 int64 = -67
                        _ = v28
                        _ = v29
                    }
                    _ = v26
                } else {
                    var v30 float64 = 3.14159
                    v31 := GenS4{F0: v5}
                    v32 := []int{(fi25 / ((fi25) | 1)), fi24, fi24}
                    _ = v30
                    _ = v31
                    _ = v32
                }
            case -9223372036854775808:
                fmt.Println(v3, v5, fi24)
            }
            _ = fi25
        }
        _ = fi24
    }
    var v33 bool = (!false)
    v5 = v1
    v5 = (v1 >> ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
