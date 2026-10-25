package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int16 }
func genFn0() uint {
    switch "\n\t" {
    case "":
        var v7 int16 = 0
        _ = v7
    case "hello":
        v8 := []int{-9223372036854775808, 4294967295}
        _ = v8
    }
    v9 := []int{}
    var v10 float32 = -2.25
    v10 += v10
    var v11 int8 = 126
    _ = v10
    _ = v11
    _ = v9
    return 604
}
func genFn1(p0 string, p1 int8, p2 uint8) uint16 {
    if (p2 != 0) {
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 bool = (4294967295 < 18446744073709551615)
                _ = fi13
                _ = v14
            }
            switch fi12 {
            case 0:
                var v15 uint = 675
                _ = v15
            case 32768:
                p1 = p1
            }
            _ = fi12
        }
        p1 = 85
        v16 := []int{400}
        _ = v16
    }
    var v17 float32 = (3.14159 * 0.1)
    if (!(-2147483648 <= 365)) {
        if false {
            v18 := []int{0, 32767, -669, -359}
            if false {
                v19 := GenS1{F0: 32767, F1: p2}
                _ = v19
            }
            var v20 string = p0
            _ = v18
            _ = v20
        }
    } else {
        switch 4294967295 {
        case -6:
            v21 := []int{}
            _ = v21
        case 4294967295:
            v22 := []int{2147483647, 127, 9223372036854775807, 2}
            _ = v22
        case 127:
            v17 += (v17 * (v17 * v17))
        default:
            switch 0 {
            case -671:
                for fi23 := 0; fi23 < 3; fi23++ {
                    v24 := GenS2{F0: "hello", F1: uint16(p2), F2: uint64(p2)}
                    _ = fi23
                    _ = v24
                }
            case 4294967295:
                v25 := GenS1{F0: 52, F1: p2}
                _ = v25
            case 1:
                v26 := GenS3{F0: -196}
                _ = v26
            default:
                var v27 uint16 = 234
                _ = v27
            }
        }
        v28 := GenS3{F0: 0}
        _ = v28
    }
    var v29 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v29
    return 2
}

func main() {
    v1 := []string{}
    v2 := []int{-446, 4294967295, 42}
    var v3 int = -2147483648
    var v4 string = ("hello" + (("hello" + "\n\t") + ((("a" + ("" + "\n\t")) + ("\n\t" + (("foo bar" + ("foo bar" + "hello")) + ((("a" + "\n\t") + ("a" + "a")) + "\n\t")))) + "a")))
    var v5 int = v3
    v6 := make(map[int]uint32)
    fmt.Println(v5, v4)
    v30 := GenS3{F0: 172}
    var v31 int32 = 42
    var v32 float64 = float64(0.0)
    var v33 int32 = (v31 / ((v31) | 1))
    fmt.Println(v3)
    var v34 float64 = 0.1
    v6[v5] = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
