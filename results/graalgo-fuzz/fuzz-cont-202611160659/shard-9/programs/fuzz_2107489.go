package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 string }

func main() {
    var v1 int16 = 270
    var v2 float64 = (((1e10 + -2.25) * -2.25) / 1.0)
    v3 := make(map[int]bool)
    var v4 float64 = v2
    var v5 int16 = (-(v1))
    var v6 uint8 = 128
    if false {
        var v7 uint32 = 0
        v8 := GenS1{F0: v7}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            v1 += (^(v5))
            fmt.Println(v1, v4)
            _ = fi9
        }
        v3[-926] = false
        var v10 int64 = 0
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v2 = 3.14159
        _ = fi11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            v6 = v6
            switch "a" {
            case "":
                var v14 int8 = -1
                _ = v14
            case "\n\t":
                v15 := GenS2{F0: fi13, F1: "a"}
                _ = v15
            }
            _ = fi13
        }
        _ = fi12
    }
    if (!(!(!true))) {
        var v16 bool = false
        for fi17 := 0; fi17 < 13; fi17++ {
            v18 := []int{(fi17 / ((32767) | 1)), (fi17 - (fi17 + fi17)), (-(fi17))}
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 uint = 921
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                _ = fi19
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v5 -= v1
    if (1 < 12) {
        fmt.Println(v5, v1)
        v4 -= v2
    }
    v4 += float64((3.14159 - 0.1))
    delete(v3, 127)
    v21 := GenS1{F0: 1}
    v22 := GenS1{F0: 2}
    v23 := []int{-1, 9223372036854775807, -9223372036854775808, 4294967295}
    fmt.Println(v6, v1)
    switch -579 {
    case -148:
        if (true || (1.5 == (3.14159 + (float32(v4) - 1e10)))) {
            v24 := GenS1{F0: 2}
            v25 := GenS2{F0: 682, F1: "a"}
            _ = v24
            _ = v25
        }
    case 270:
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
    case -9223372036854775808:
        v26 := GenS1{F0: 2147483647}
        _ = v26
    }
    var v27 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
