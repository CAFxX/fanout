package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int64 }
func genFn0() int32 {
    v5 := []int{0}
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint8 = 2
        fmt.Println(v7)
        _ = fi6
        _ = v7
    }
    var v8 uint = 602
    _ = v5
    _ = v8
    return -125
}

func main() {
    v1 := []string{"a", "\n\t"}
    var v2 float64 = 3.14159
    v3 := []uint32{2, 4294967295}
    var v4 uint16 = 197
    v9 := GenS1{F0: false, F1: "", F2: -9223372036854775808}
    switch 4294967295 {
    case -706:
        for fi10 := 0; fi10 < 6; fi10++ {
            if len(v3) > 0 {
                v3[0] = 2
            }
            _ = fi10
        }
    case 127:
        v1 = append(v1, "hello")
    default:
        v11 := GenS1{F0: true, F1: "a", F2: 1}
        _ = v11
    }
    var v12 int = 127
    v2 += 1e10
    v3 = append(v3, uint32(v4))
    var v13 int32 = genFn0()
    v4 = (v4 - v4)
    switch "" {
    case "a":
        fmt.Println(v13, v13)
    }
    var v14 string = "a"
    var v15 uint16 = v4
    var v16 uint32 = 2147483647
    v17 := GenS1{F0: false, F1: (v14 + ("\n\t" + v14)), F2: -1}
    if len(v1) > 0 {
        v1[0] = (v14 + v14)
    }
    v2 -= 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
