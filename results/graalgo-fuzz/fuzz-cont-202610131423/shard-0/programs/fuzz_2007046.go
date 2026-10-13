package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0(p0 int32) int {
    v6 := []int{}
    for fi7 := 0; fi7 < 10; fi7++ {
        if len(v6) > 0 {
            v6[0] = -2147483648
        }
        var v8 int32 = p0
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 2147483647
}
func genFn1() int16 {
    var v9 int32 = 42
    if false {
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 13; fi11++ {
                v9 = 1
                var v12 uint64 = 0
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 16; fi13++ {
                v9 = (v9 ^ 42)
                _ = fi13
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                fmt.Println(fi14, fi14, fi14)
                v15 := []int{}
                _ = fi14
                _ = v15
            }
            _ = fi10
        }
        fmt.Println(v9)
        var v16 uint8 = 128
        _ = v16
    } else {
        var v17 uint8 = 42
        v17 = v17
        v17 = (v17 + 10)
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 uint32 = 158
            v21 := []int{(fi18 / (((^(fi19))) | 1))}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    var v22 int16 = 32767
    _ = v22
    _ = v9
    return 860
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    var v3 int16 = 42
    var v4 int8 = 0
    var v5 string = ("\n\t" + "")
    var v23 float64 = -2.25
    fmt.Println(v4)
    var v24 uint8 = 255
    if (!false) {
        if ((v4 << ((1) & 63)) >= v4) {
            var v25 uint16 = 453
            v26 := v2[v5]
            switch "\n\t" {
            case "a":
                var v27 float32 = 1e10
                _ = v27
            }
            _ = v25
            _ = v26
        }
        v23 = float64(1.5)
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        fmt.Println(v24, v23)
        fmt.Println(v23)
        v23 = (3.14159 + v23)
        _ = fi28
    }
    if true {
        var v29 int32 = 42
        _ = v29
    }
    var v30 bool = (!false)
    v31 := GenS1{F0: 1, F1: 0}
    v32 := []int{127, -1}
    fmt.Println(v5, v4, v23)
    var v33 uint8 = (v24 | (v24 << ((65535) & 63)))
    v33 += v33
    switch (v5 + v5) {
    case "":
        _ = len(v1)
    default:
        var v34 int = 32768
        _ = v34
    }
    fmt.Println(v23, v23)
    v35 := GenS1{F0: 32767, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
}
