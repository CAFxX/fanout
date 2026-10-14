package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0() int16 {
    var v6 bool = true
    fmt.Println(v6)
    for fi7 := 0; fi7 < 10; fi7++ {
        v6 = (-9223372036854775808 == 465)
        var v8 int32 = -912
        _ = fi7
        _ = v8
    }
    _ = v6
    return -1
}
func genFn1() int8 {
    var v9 int16 = -1
    switch 127 {
    case 0:
        var v10 int16 = (v9 - genFn0())
        _ = v10
    default:
        v11 := []int{489, 127, 0, -9223372036854775808}
        _ = v11
    }
    var v12 uint64 = 4294967295
    var v13 int64 = 9223372036854775807
    if (true && (2147483647 != -697)) {
        v14 := []int{-1}
        v15 := GenS1{F0: 260, F1: 4294967295}
        v16 := GenS1{F0: 0, F1: 2}
        _ = v14
        _ = v15
        _ = v16
    } else {
        v13 = v13
        switch "hello" {
        case "hello":
            v17 := GenS2{F0: 65535}
            _ = v17
        }
        var v18 float64 = (0.1 * -2.25)
        _ = v18
    }
    _ = v12
    _ = v13
    _ = v9
    return 102
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 32767
    v3 := []string{}
    v4 := []string{"hello", "\n\t", ""}
    v5 := make(map[int]string)
    switch 738 {
    case 2147483647:
        fmt.Println(v2)
    case 32768:
        v19 := GenS1{F0: v2, F1: 4294967295}
        _ = v19
    case -162:
        if len(v4) > 0 {
            v4[0] = (("\n\t" + "a") + "foo bar")
        }
    }
    v2 += v2
    if false {
        if true {
            v20 := GenS2{F0: 497}
            v21 := GenS2{F0: (v2 - (v2 >> ((993) & 63)))}
            v22 := GenS1{F0: v2, F1: 919}
            _ = v20
            _ = v21
            _ = v22
        }
        v23 := GenS1{F0: 32767, F1: 1}
        v2 -= (v2 >> ((67) & 63))
        _ = v23
    } else {
        var v24 int = 1
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 int8 = 0
            _ = fi25
            _ = v26
        }
        switch ("hello" + "hello") {
        case "":
            var v27 bool = (51 == 42)
            _ = v27
        case "hello":
            fmt.Println(v24, v2)
        }
        _ = v24
    }
    v28 := v1[("\n\t" + "a")]
    v29 := GenS2{F0: (v2 | v2)}
    var v30 uint64 = 0
    v28 += (v28 / ((v28) | 1))
    v28 -= (v28 + (v28 * (v28 >> ((65535) & 63))))
    for fi31 := 0; fi31 < 8; fi31++ {
        if (true && false) {
            v30 += (v30 | (v30 | (v30 / ((442) | 1))))
        }
        v4 = append(v4, "\n\t")
        fmt.Println(fi31, fi31)
        _ = fi31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
