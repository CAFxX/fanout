package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint }
func genFn0() uint64 {
    var v7 int16 = -799
    v8 := []int{}
    v7 = v7
    var v9 uint16 = 332
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 179
    var v2 uint8 = 2
    var v3 uint32 = 239
    v4 := []string{"", ("\n\t" + (("\n\t" + "\n\t") + "")), (("\n\t" + "hello") + ("\n\t" + ("foo bar" + (("hello" + "hello") + "foo bar"))))}
    var v5 float64 = 0.1
    v6 := make(map[string]bool)
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS1{F0: -2147483648, F1: (2147483647 <= -2147483648), F2: 1}
        v6["foo bar"] = false
        var v12 uint32 = v1
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int = 127
    switch (v13 * v13) {
    case -1:
        for fi14 := 0; fi14 < 4; fi14++ {
            if false {
                var v15 float64 = v5
                _ = v15
            } else {
                v16 := GenS1{F0: 614, F1: false, F2: 42}
                _ = v16
            }
            var v17 uint16 = 544
            _ = fi14
            _ = v17
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v18 := GenS1{F0: -2147483648, F1: true, F2: 4294967295}
    var v19 float32 = 0.1
    v2 = (v2 * (v2 ^ v2))
    if (v1 <= (v3 - 422)) {
        v20 := GenS1{F0: 0, F1: ((!false) && false), F2: 0}
        _ = len(v4)
        for fi21 := 0; fi21 < 18; fi21++ {
            v19 = v19
            _ = fi21
        }
        _ = v20
    } else {
        var v22 int32 = -136
        fmt.Println(v3, v2)
        _ = v22
    }
    fmt.Println(v13, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
