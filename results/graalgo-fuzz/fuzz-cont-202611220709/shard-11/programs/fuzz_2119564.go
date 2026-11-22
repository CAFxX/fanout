package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int8 }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0(p0 int32) uint8 {
    p0 = (-(p0))
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: (((0.1 - (1.5 * 0.1)) != 3.14159) && (fi6 >= (fi6 << ((1) & 63)))), F1: 127, F2: -1}
        _ = fi6
        _ = v7
    }
    p0 -= p0
    fmt.Println(p0)
    if (3.14159 != (1e10 / float32(float64((3.14159 / 1.0))))) {
        v8 := []int{}
        var v9 int8 = 1
        v9 -= 100
        _ = v8
        _ = v9
    } else {
        fmt.Println(p0)
        var v10 float32 = (3.14159 + (1.5 / 1.0))
        _ = v10
    }
    _ = p0
    return 42
}
func genFn1(p0 int64) uint16 {
    var v11 float64 = 1.5
    v12 := []int{517}
    switch "" {
    case "a":
        p0 += (-(p0))
    case "foo bar":
        if (!(false || (1000 != -136))) {
            var v13 bool = false
            switch ("a" + "a") {
            case "\n\t":
                _ = len(v12)
            }
            switch ("hello" + "a") {
            case "hello":
                p0 += 696
            case "a":
                switch 42 {
                case 128:
                    var v14 float64 = (v11 - (v11 - v11))
                    _ = v14
                case 0:
                    v15 := GenS1{F0: (((((v13 || (100 >= 42)) || v13) || ((!true) && (v13 && (2147483648 < 1)))) || v13) && v13), F1: 32768, F2: 1}
                    _ = v15
                default:
                    var v16 int8 = 127
                    _ = v16
                }
            }
            _ = v13
        }
    case "\n\t":
        p0 += (p0 * p0)
    }
    if len(v12) > 0 {
        v12[0] = 127
    }
    _ = p0
    _ = v11
    _ = v12
    return 697
}

func main() {
    v1 := []uint32{}
    v2 := []string{"a", "a"}
    v3 := []uint32{2147483647, 65535, 65535, 2}
    var v4 int32 = -493
    v5 := make(map[string]int)
    var v17 bool = true
    fmt.Println(v4)
    switch ("foo bar" + ("" + (("a" + ("\n\t" + "\n\t")) + ""))) {
    case "\n\t":
        if true {
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 float32 = 1.5
                var v20 bool = (v4 != -1)
                fmt.Println(v20)
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = len(v1)
        }
    case "foo bar":
        var v21 int32 = (v4 - v4)
        _ = v21
    default:
        v17 = false
    }
    switch (("foo bar" + "\n\t") + "") {
    case "foo bar":
        v17 = (2 != 255)
    }
    switch 128 {
    case -2147483648:
        var v22 bool = v17
        _ = v22
    case 9223372036854775807:
        v4 -= (v4 % ((-1) | 1))
    default:
        var v23 int64 = 0
        _ = v23
    }
    if ("hello" >= "hello") {
        if v17 {
            v17 = v17
        } else {
            var v24 uint = 2
            if false {
                v25 := GenS2{F0: 2, F1: 0}
                v4 = v4
                _ = v25
            } else {
                switch "\n\t" {
                case "a":
                    v26 := GenS1{F0: true, F1: -590, F2: 126}
                    _ = v26
                }
                if false {
                    fmt.Println(v4, v17, v24)
                } else {
                    v4 = 211
                    v4 = 0
                }
            }
            _ = v24
        }
    }
    v27 := v5["foo bar"]
    v28 := GenS1{F0: v17, F1: (v27 - 32768), F2: 88}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
