package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 string = ("foo bar" + "hello")
    v2 := make(map[string]int)
    v3 := []int{-509, 4294967295}
    var v4 uint16 = 255
    var v5 uint8 = 42
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            v1 = "foo bar"
            _ = fi7
        }
        var v8 int8 = 127
        _ = fi6
        _ = v8
    }
    v1 = "\n\t"
    v4 -= v4
    var v9 int32 = 2147483647
    v10 := GenS1{F0: false}
    if false {
        _ = len(v3)
        v11 := v2[v1]
        _ = v11
    }
    switch v1 {
    case "":
        var v12 int16 = -512
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        if (false || false) {
            var v14 float64 = (((0.0 / 3.14159) * 1e10) + 0.1)
            _ = v14
        }
        _ = fi13
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        delete(v2, "")
        if true {
            v16 := GenS1{F0: (!(!true))}
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 bool = ((-2.25 - 1.5) <= 1e10)
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        _ = fi15
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        fmt.Println(v9, fi19)
        switch v1 {
        case "":
            var v20 float64 = -2.25
            _ = v20
        case "foo bar":
            var v21 string = (v1 + v1)
            _ = v21
        case "hello":
            v22 := v2[""]
            _ = v22
        default:
            v3 = append(v3, (fi19 >> ((255) & 63)))
        }
        if (0 >= 100) {
            v23 := GenS1{F0: ((0.0 > (1.5 / float64((-2.25 + 1.5)))) || true)}
            _ = v23
        }
        _ = fi19
    }
    v5 += 120
    var v24 int16 = -91
    v25 := GenS1{F0: (v4 <= v4)}
    v26 := GenS1{F0: (false || (2147483648 >= uint32(v5)))}
    var v27 uint64 = 2
    var v28 int32 = (v9 | v9)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
