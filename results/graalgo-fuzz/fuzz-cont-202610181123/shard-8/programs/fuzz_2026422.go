package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint32) int8 {
    v4 := []int{-207}
    fmt.Println(p0)
    p0 -= p0
    _ = p0
    _ = v4
    return 42
}
func genFn1() int {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{}
        var v7 int32 = -559
        var v8 int16 = -32768
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    if ((0.1 > 3.14159) && (1 <= 2)) {
        var v9 float32 = 0.0
        fmt.Println(v9)
        _ = v9
    }
    return -2147483648
}

func main() {
    var v1 int16 = 32767
    var v2 int8 = 102
    var v3 int64 = 42
    fmt.Println(v2)
    fmt.Println(v2, v1)
    v1 = int16(v3)
    v10 := GenS1{F0: 0}
    v11 := []int{}
    var v12 bool = ((!true) || false)
    var v13 int64 = (^(v3))
    for fi14 := 0; fi14 < 15; fi14++ {
        switch fi14 {
        case 42:
            v11 = append(v11, int(v1))
        case 32768:
            _ = len(v11)
        case 127:
            v15 := GenS1{F0: 2}
            _ = v15
        default:
            fmt.Println(v1, v12)
        }
        switch "hello" {
        case "foo bar":
            if len(v11) > 0 {
                v11[0] = 42
            }
        case "hello":
            v12 = false
        case "":
            _ = len(v11)
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            if len(v11) > 0 {
                v11[0] = fi14
            }
            if true {
                var v17 int8 = (-(v2))
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = v17
            } else {
                if false {
                    fmt.Println(v12)
                    v1 += (^(v1))
                }
                var v18 uint16 = 2
                fmt.Println(v12)
                _ = v18
            }
            _ = fi16
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
