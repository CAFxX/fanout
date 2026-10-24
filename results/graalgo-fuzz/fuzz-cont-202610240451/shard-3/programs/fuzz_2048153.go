package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 string }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint64; F1 bool }
func genFn0() int8 {
    if (2147483647 >= -2147483648) {
        if true {
            switch "a" {
            case "foo bar":
                var v7 int8 = -128
                _ = v7
            case "\n\t":
                var v8 bool = (1 == 2147483648)
                _ = v8
            }
        } else {
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 float32 = 0.1
                v11 := []int{1, (fi9 / ((fi9) | 1)), fi9}
                _ = fi9
                _ = v10
                _ = v11
            }
            switch "" {
            case "a":
                var v12 int64 = -9223372036854775808
                _ = v12
            case "hello":
                var v13 int8 = -128
                _ = v13
            default:
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 int32 = -790
                    _ = fi14
                    _ = v15
                }
            }
            v16 := []int{32768}
            _ = v16
        }
        v17 := GenS1{F0: 0, F1: 127, F2: "hello"}
        _ = v17
    }
    var v18 float32 = (-2.25 + (0.0 + 1e10))
    _ = v18
    return 67
}

func main() {
    var v1 bool = false
    var v2 bool = false
    var v3 uint16 = 2
    v4 := make(map[string]int)
    v5 := []string{"", (("a" + "foo bar") + "foo bar"), "a", "a"}
    v6 := make(map[int]uint32)
    v1 = (v2 || (964 != -9223372036854775808))
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    v19 := GenS2{F0: 0, F1: -949}
    if (!v1) {
        v20 := GenS3{F0: -882, F1: 1, F2: v3}
        var v21 int32 = -2147483648
        _ = v20
        _ = v21
    } else {
        var v22 float64 = 3.14159
        fmt.Println(v1)
        v1 = (!v2)
        _ = v22
    }
    var v23 int16 = 42
    if (127 > 1) {
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS4{F0: (fi24 | (fi24 + -1)), F1: 90, F2: 0}
            fmt.Println(v2, v1)
            _ = fi24
            _ = v25
        }
    }
    v2 = (true && v1)
    switch -323 {
    case 22:
        var v26 uint8 = 0
        _ = v26
    case 676:
        for fi27 := 0; fi27 < 19; fi27++ {
            var v28 uint16 = v3
            fmt.Println(v1, fi27)
            fmt.Println(v28)
            _ = fi27
            _ = v28
        }
    case 1:
        v3 += (v3 >> ((4294967295) & 63))
    default:
        fmt.Println(v1)
    }
    if v2 {
        var v29 string = "\n\t"
        v30 := GenS5{F0: 0, F1: v2}
        v31 := GenS5{F0: 0, F1: (4294967295 >= 2)}
        _ = v29
        _ = v30
        _ = v31
    } else {
        v3 -= v3
    }
    v32 := GenS4{F0: 0, F1: 42, F2: 0}
    v3 = 453
    switch "" {
    case "":
        v3 -= 431
    case "hello":
        switch "" {
        case "foo bar":
            v2 = true
        case "a":
            var v33 int16 = -87
            _ = v33
        case "hello":
            v3 = 884
        }
    case "\n\t":
        _ = len(v6)
    default:
        v5 = append(v5, "foo bar")
    }
    switch int(v23) {
    case 128:
        var v34 string = (("\n\t" + ((("hello" + ("hello" + ("" + ("a" + ("" + "\n\t"))))) + "a") + "")) + "\n\t")
        _ = v34
    }
    var v35 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
