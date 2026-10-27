package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
type GenS2 struct { F0 int }
func genFn0(p0 int16, p1 int32) int8 {
    v4 := []int{}
    var v5 float32 = ((1.5 / (1e10 - (0.1 * -2.25))) * (float32(0.1) + 3.14159))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 uint8 = 77
    v2 := make(map[uint32]bool)
    v3 := []int{1, -2147483648}
    var v6 uint16 = 65535
    for fi7 := 0; fi7 < 9; fi7++ {
        switch (-(fi7)) {
        case 0:
            if true {
                var v8 uint32 = 513
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 int32 = -645
                    var v11 int16 = 0
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = v8
            }
        }
        var v12 uint16 = (v6 | v6)
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = fi7
        _ = v12
    }
    fmt.Println(v6)
    if (!(!true)) {
        var v13 string = ("" + ("foo bar" + "hello"))
        v1 += v1
        _ = v13
    } else {
        v14 := GenS1{F0: 0, F1: 0, F2: -9223372036854775808}
        switch "\n\t" {
        case "":
            v15 := GenS2{F0: 610}
            _ = v15
        case "foo bar":
            v2[65535] = true
        case "a":
            v1 = v1
        }
        var v16 string = "a"
        _ = v14
        _ = v16
    }
    v17 := GenS2{F0: 616}
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 13; fi19++ {
            v1 -= (v1 + (^(v1)))
            _ = fi19
        }
        v6 += 32767
        _ = fi18
    }
    v6 = (^(v6))
    var v20 int8 = 111
    if ((-2.25 / 0.1) >= ((((0.1 + float64(-2.25)) * (0.1 - 1.5)) * 1e10) + ((1.5 - -2.25) / 1e10))) {
        v21 := GenS2{F0: 32767}
        _ = v21
    } else {
        v1 -= 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v6
}
