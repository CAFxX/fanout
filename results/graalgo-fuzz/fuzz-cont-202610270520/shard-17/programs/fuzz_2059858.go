package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
func genFn0(p0 uint8) int8 {
    fmt.Println(p0)
    var v7 uint = 42
    _ = p0
    _ = v7
    return 127
}
func genFn1(p0 uint32, p1 string) uint32 {
    v8 := []int{}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if false {
        var v9 uint = 1
        v8 = append(v8, -1)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v8
    return p0
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{127, 859, -1}
    v3 := []uint32{0}
    var v4 uint16 = 32767
    var v5 int32 = 2147483647
    v6 := make(map[int]int)
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int16 = 23
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: (-1 == 0), F1: 1}
    v5 += -202
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v4 = (v4 ^ (v4 - v4))
        _ = fi13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 int8 = -105
            switch fi15 {
            case 633:
                var v17 float64 = 3.14159
                _ = v17
            case 460:
                v5 -= v5
            }
            _ = fi15
            _ = v16
        }
        var v18 int64 = 742
        _ = fi14
        _ = v18
    }
    var v19 int16 = -32768
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 int32 = (v5 >> ((0) & 63))
        v21 -= (v5 >> ((18446744073709551615) & 63))
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 int32 = (v21 - v21)
            switch ("foo bar" + "a") {
            case "":
                if (("" >= (("a" + ("\n\t" + "a")) + (("hello" + ("\n\t" + ((("\n\t" + ("" + ("foo bar" + ("\n\t" + ("hello" + ("hello" + "foo bar")))))) + "\n\t") + ("foo bar" + "")))) + "hello"))) || (fi20 <= fi20)) {
                    v2 = append(v2, fi20)
                } else {
                    v1["a"] = "\n\t"
                    var v24 uint32 = genFn1(2147483648, "a")
                    _ = v24
                }
            }
            fmt.Println(v19, v4)
            if (1e10 >= (0.1 * (float64(0.1) + (float64(1.5) + 1e10)))) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                var v25 float32 = 1.5
                _ = v25
            }
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    switch "foo bar" {
    case "\n\t":
        var v26 string = "a"
        _ = v26
    case "":
        v27 := GenS2{F0: (v5 - (v5 & (v5 & -1))), F1: 1, F2: 4294967295}
        _ = v27
    case "foo bar":
        v4 += v4
    }
    var v28 string = ("hello" + (("" + "hello") + ("hello" + "hello")))
    fmt.Println(v4)
    for fi29 := 0; fi29 < 4; fi29++ {
        v30 := v1["foo bar"]
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: true, F1: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
