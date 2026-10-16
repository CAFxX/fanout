package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int32; F1 uint16; F2 string }
func genFn0() uint8 {
    var v7 uint = 65535
    var v8 uint32 = 2
    _ = v7
    _ = v8
    return 143
}
func genFn1() int64 {
    if (0 == 2147483647) {
        var v9 int32 = 661
        for fi10 := 0; fi10 < 4; fi10++ {
            v9 = v9
            v11 := []int{(fi10 - fi10), fi10, (-(fi10))}
            _ = fi10
            _ = v11
        }
        var v12 string = "a"
        _ = v12
        _ = v9
    } else {
        if (32767 < 0) {
            var v13 int = -9223372036854775808
            if ((-2.25 - 1e10) <= 1e10) {
                v14 := GenS1{F0: v13, F1: -9223372036854775808, F2: (!(2 <= 916))}
                v15 := []int{(v13 + v13)}
                var v16 bool = true
                _ = v14
                _ = v15
                _ = v16
            } else {
                v17 := []int{(v13 * v13)}
                var v18 string = "hello"
                _ = v17
                _ = v18
            }
            _ = v13
        }
        var v19 float64 = -2.25
        _ = v19
    }
    switch -631 {
    case 0:
        var v20 uint64 = 558
        _ = v20
    default:
        var v21 int16 = 42
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS2{F0: "\n\t"}
        var v24 uint = 18446744073709551615
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS3{F0: false, F1: (!(17 != 2))}
    v26 := GenS1{F0: 2147483647, F1: 9223372036854775807, F2: false}
    _ = v25
    _ = v26
    return 0
}

func main() {
    v1 := []string{"foo bar", "a", "\n\t"}
    v2 := make(map[uint32]bool)
    var v3 int64 = 42
    var v4 int8 = -113
    v5 := make(map[string]bool)
    v6 := make(map[int]string)
    fmt.Println(v4)
    v27 := GenS1{F0: 4294967295, F1: 1, F2: (true || false)}
    fmt.Println(v4)
    delete(v5, ("foo bar" + "\n\t"))
    if false {
        var v28 int32 = 0
        var v29 uint8 = 127
        fmt.Println(v4, v4)
        _ = v28
        _ = v29
    }
    v30 := GenS4{F0: -2147483648, F1: 42, F2: ""}
    var v31 int16 = -30
    var v32 float64 = (0.0 / 1.5)
    v33 := GenS3{F0: (869 == 519), F1: true}
    if ((false && false) && (true && false)) {
        v3 -= v3
        for fi34 := 0; fi34 < 8; fi34++ {
            v1 = append(v1, ("\n\t" + "\n\t"))
            fmt.Println(v31, v4)
            _ = fi34
        }
    }
    var v35 float32 = 3.14159
    v36 := v2[1]
    fmt.Println(v36)
    v37 := GenS1{F0: 259, F1: -1, F2: (genFn1() > (^(v3)))}
    if (!(-273 == 32767)) {
        if true {
            v1 = append(v1, "hello")
        } else {
            var v38 uint64 = 0
            var v39 uint32 = 2147483648
            var v40 bool = v36
            _ = v38
            _ = v39
            _ = v40
        }
        fmt.Println(v4, v35)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
