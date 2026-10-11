package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint {
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
        _ = fi5
    }
    _ = v4
    return 2
}
func genFn1(p0 int32, p1 int, p2 uint16) string {
    v7 := []int{25}
    if true {
        var v8 int16 = 824
        _ = v8
    } else {
        var v9 uint8 = 2
        var v10 uint = 628
        _ = v10
        _ = v9
    }
    var v11 float64 = (((0.0 - 1e10) / 1.5) * float64(0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return ("foo bar" + "\n\t")
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    fmt.Println(v1)
    switch -2147483648 {
    case 9223372036854775807:
        fmt.Println(v1)
    default:
        v12 := GenS1{F0: -12}
        _ = v12
    }
    var v13 int16 = 32767
    v14 := v3["\n\t"]
    var v15 uint32 = 2147483648
    fmt.Println(v1, v14)
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 uint8 = 2
        _ = fi16
        _ = v17
    }
    v1 += float64((-2.25 + 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
