package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int64) uint {
    switch "" {
    case "foo bar":
        var v7 uint8 = 0
        _ = v7
    case "":
        switch "hello" {
        case "hello":
            var v8 uint16 = 32767
            _ = v8
        case "\n\t":
            v9 := GenS1{F0: -1, F1: (("a" + ((("hello" + "") + "") + "\n\t")) + "\n\t")}
            _ = v9
        case "a":
            v10 := []int{748, 32768}
            _ = v10
        }
    default:
        v11 := []int{0}
        _ = v11
    }
    var v12 uint16 = 0
    _ = p0
    _ = v12
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 1
    v3 := make(map[string]bool)
    var v4 int32 = -2147483648
    var v5 int32 = v4
    var v6 int8 = 127
    v13 := GenS2{F0: 243}
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS1{F0: v6, F1: "hello"}
        fmt.Println(fi14)
        _ = fi14
        _ = v15
    }
    var v16 uint = 1
    fmt.Println(v2, v6)
    v5 = (v5 - (^(v4)))
    var v17 uint64 = 0
    var v18 int64 = -1
    var v19 int8 = v6
    if false {
        v4 -= v4
        v20 := []int{-210, 32767, 127}
        v21 := GenS2{F0: 68}
        _ = v20
        _ = v21
    }
    var v22 int8 = v19
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 14; fi24++ {
            v25 := GenS1{F0: (v6 % (((v22 * v19)) | 1)), F1: ("foo bar" + "a")}
            var v26 bool = (190 <= -101)
            _ = fi24
            _ = v25
            _ = v26
        }
        if true {
            v3["\n\t"] = (!((v16 << ((v16) & 63)) <= (v16 << ((255) & 63))))
            var v27 float32 = 1.5
            _ = v27
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 int16 = 622
            _ = fi28
            _ = v29
        }
        _ = fi23
    }
    v6 = 1
    for fi30 := 0; fi30 < 4; fi30++ {
        v4 -= v5
        v31 := []int{(fi30 << ((genFn0(-1)) & 63))}
        _ = fi30
        _ = v31
    }
    v32 := []int{127, 2147483647}
    if len(v32) > 0 {
        v32[0] = int(v17)
    }
    if len(v32) > 0 {
        v32[0] = -1
    }
    var v33 int8 = (^(v22))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
