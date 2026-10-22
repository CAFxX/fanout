package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint64) uint64 {
    var v4 uint32 = 65535
    for fi5 := 0; fi5 < 11; fi5++ {
        switch "foo bar" {
        case "hello":
            switch -81 {
            case 0:
                fmt.Println(v4)
            }
        }
        _ = fi5
    }
    switch 232 {
    case 2147483647:
        if ((p0 & (p0 & p0)) != (^(p0))) {
            if true {
                v6 := GenS1{F0: 1}
                fmt.Println(v4, v4)
                _ = v6
            } else {
                var v7 string = "a"
                _ = v7
            }
        }
    }
    _ = p0
    _ = v4
    return p0
}

func main() {
    var v1 int = 32767
    v2 := []int{-1, v1, 374, v1}
    var v3 uint8 = 77
    fmt.Println(v1, v3)
    var v8 int = v1
    for fi9 := 0; fi9 < 5; fi9++ {
        v3 = v3
        _ = fi9
    }
    if true {
        var v10 int16 = -32768
        _ = v10
    } else {
        v1 += (v8 | 0)
        for fi11 := 0; fi11 < 18; fi11++ {
            if ((true && false) || (-2147483648 == 2147483647)) {
                var v12 float32 = (((float32(3.14159) / ((0.0 + ((0.1 * 0.0) + 0.1)) - 1e10)) + (float32(-2.25) / 1.5)) / 1.5)
                v1 = v8
                _ = v12
            } else {
                var v13 uint64 = 12
                _ = v13
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi11
        }
        var v14 string = "\n\t"
        _ = v14
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v15 := GenS1{F0: 446}
    var v16 int16 = 42
    fmt.Println(v16)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v17 uint32 = 698
    fmt.Println(v17, v3, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
