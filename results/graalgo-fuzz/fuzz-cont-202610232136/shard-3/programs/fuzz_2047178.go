package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint64 }

func main() {
    var v1 uint16 = 495
    v2 := []string{"a", "a", "\n\t"}
    var v3 string = ("" + (("a" + "") + "a"))
    v4 := make(map[string]int)
    v5 := []string{v3, v3}
    v6 := []int{}
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 string = v3
            _ = fi9
            _ = v10
        }
        switch "foo bar" {
        case "hello":
            fmt.Println(v3, v1, v3)
        case "foo bar":
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 float64 = (0.1 * float64((1.5 + ((float32((v7 * float64(-2.25))) + float32(0.1)) / (1e10 + -2.25)))))
                _ = len(v4)
                var v13 int16 = 1
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = fi8
    }
    fmt.Println(v7, v3)
    fmt.Println(v3)
    var v14 uint8 = 42
    fmt.Println(v1)
    var v15 int16 = -1
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(v3, v7, v1)
            var v18 uint64 = 420
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    var v19 uint8 = (v14 >> ((151) & 63))
    v20 := GenS1{F0: 313, F1: -1, F2: 4294967295}
    v21 := GenS1{F0: 660, F1: 100, F2: 18446744073709551615}
    if (false && true) {
        v15 += (v15 & v15)
        var v22 string = (v3 + "")
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
