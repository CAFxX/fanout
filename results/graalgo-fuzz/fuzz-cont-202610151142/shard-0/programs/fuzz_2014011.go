package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }
func genFn0() int32 {
    var v4 uint = 65535
    if ((0 < -850) && (false && (-36 < 127))) {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 uint32 = 0
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{(fi7 * fi7), 2147483647, (-(fi7)), (fi7 >> (((v4 + v4)) & 63))}
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: "\n\t", F1: 4294967295, F2: v4}
        _ = v9
    } else {
        v10 := []int{1, 4294967295, 9223372036854775807, -782}
        fmt.Println(v4)
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 int8 = -1
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    _ = v4
    return 42
}

func main() {
    var v1 int16 = -353
    v2 := []string{}
    var v3 uint32 = 0
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v1, v3, v1)
        fmt.Println(fi13)
        _ = fi13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 float64 = ((float64(1e10) * 0.0) * 0.0)
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: "foo bar", F1: 462, F2: 18446744073709551615}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v17 float64 = -2.25
    v2 = append(v2, "")
    if true {
        v18 := GenS1{F0: "\n\t", F1: -516, F2: 255}
        var v19 int = 42
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        if true {
            v1 -= v1
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 string = (("foo bar" + "hello") + "")
                _ = fi21
                _ = v22
            }
            var v23 float32 = float32(float64(float32(float64(0.0))))
            _ = v23
        }
        var v24 int32 = -308
        _ = fi20
        _ = v24
    }
    var v25 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
}
