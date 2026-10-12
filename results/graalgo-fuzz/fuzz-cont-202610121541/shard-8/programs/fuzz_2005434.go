package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint16, p1 int8) int64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        p1 = (^(p1))
        _ = fi6
    }
    var v7 uint32 = 2147483647
    v8 := []int{}
    var v9 int = 284
    fmt.Println(v9, p0, v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}
func genFn1() string {
    var v10 int16 = -32768
    v10 += v10
    v11 := []int{}
    var v12 uint16 = 554
    var v13 int8 = int8(v10)
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return "\n\t"
}

func main() {
    var v1 float64 = (0.0 / ((1.5 / (1.5 + ((float64(((-2.25 * (1.5 + (1.5 / 0.1))) * -2.25)) + 1.5) * (1e10 * 1.5)))) + 1.5))
    v2 := make(map[string]uint32)
    var v3 uint16 = 2
    v4 := []string{}
    var v5 float64 = (1.5 - (v1 * float64(1e10)))
    if (-2.25 <= (v5 / (1e10 - float64(0.0)))) {
        v5 = (-2.25 / 1.5)
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 uint8 = 128
            v4 = append(v4, "")
            _ = fi14
            _ = v15
        }
    } else {
        var v16 int16 = 1
        _ = v16
    }
    v2["hello"] = 2147483648
    if (((1 > -1) && true) || true) {
        var v17 float32 = (1.5 - -2.25)
        v4 = append(v4, ("a" + "foo bar"))
        var v18 uint32 = 2147483647
        _ = v17
        _ = v18
    }
    var v19 float32 = 0.1
    var v20 uint8 = 2
    v5 = float64(v19)
    for fi21 := 0; fi21 < 14; fi21++ {
        fmt.Println(v19)
        _ = fi21
    }
    fmt.Println(v20, v19, v1)
    v22 := GenS1{F0: -439}
    if true {
        var v23 int8 = 0
        for fi24 := 0; fi24 < 13; fi24++ {
            v1 -= (0.1 / 1.0)
            _ = fi24
        }
        v23 = 100
        _ = v23
    }
    var v25 int = 1
    v26 := GenS1{F0: (v25 - -277)}
    switch "hello" {
    case "foo bar":
        var v27 bool = (!(0.0 > (v1 + v5)))
        _ = v27
    default:
        var v28 int16 = -602
        _ = v28
    }
    v4 = append(v4, (("\n\t" + (("" + "\n\t") + "\n\t")) + "a"))
    fmt.Println(v5)
    for fi29 := 0; fi29 < 19; fi29++ {
        switch (fi29 + v25) {
        case -2147483648:
            v4 = append(v4, "foo bar")
        }
        var v30 uint16 = (^(v3))
        var v31 uint64 = 956
        _ = fi29
        _ = v30
        _ = v31
    }
    v5 += 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
