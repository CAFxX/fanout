package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int32; F2 bool }
func genFn0(p0 uint, p1 int8) uint8 {
    if (2147483648 >= 359) {
        v4 := GenS1{F0: ((!((("\n\t" + "foo bar") + (("foo bar" + ("\n\t" + "\n\t")) + "")) < "hello")) || (!(55 < -9223372036854775808)))}
        _ = v4
    } else {
        p0 += 117
        switch 128 {
        case 4294967295:
            v5 := []int{-9223372036854775808}
            _ = v5
        case -224:
            v6 := GenS2{F0: 127, F1: -2147483648, F2: true}
            _ = v6
        }
    }
    var v7 int8 = -30
    var v8 int = 2147483647
    v9 := GenS2{F0: (v8 * 4294967295), F1: 42, F2: false}
    fmt.Println(p0, v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 146
}

func main() {
    var v1 uint = 0
    var v2 int32 = 42
    v3 := make(map[int]int)
    v1 = v1
    v2 += (^(v2))
    v10 := GenS2{F0: 128, F1: v2, F2: true}
    var v11 uint32 = 984
    if (1 <= 0) {
        v12 := []int{0, int(v2)}
        v13 := v3[-172]
        _ = v12
        _ = v13
    }
    delete(v3, 128)
    v14 := []int{}
    v11 = 1
    v15 := GenS1{F0: true}
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int8 = 1
        fmt.Println(v17)
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v14 = append(v14, 217)
        v19 := GenS2{F0: (fi18 + (^(fi18))), F1: v2, F2: false}
        _ = fi18
        _ = v19
    }
    var v20 uint16 = 0
    var v21 uint16 = v20
    v11 = (^(v11))
    fmt.Println(v2)
    switch -159 {
    case 736:
        v21 -= (v20 - v20)
    case -2147483648:
        v14 = append(v14, 32767)
    case 247:
        var v22 float32 = (1.5 / -2.25)
        _ = v22
    default:
        fmt.Println(v2, v21, v20)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
