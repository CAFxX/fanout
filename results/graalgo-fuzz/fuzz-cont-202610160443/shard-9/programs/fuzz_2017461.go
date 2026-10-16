package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint32) uint {
    v7 := []int{-9223372036854775808, 210, 32767}
    switch "" {
    case "\n\t":
        var v8 uint16 = 255
        _ = v8
    default:
        var v9 int8 = 1
        _ = v9
    }
    var v10 uint16 = 0
    _ = p0
    _ = v10
    _ = v7
    return 0
}

func main() {
    var v1 uint8 = 255
    var v2 string = "a"
    v3 := make(map[int]int)
    var v4 float32 = (1.5 + 0.0)
    var v5 int16 = 683
    var v6 bool = (-22 <= 42)
    v4 -= 1.5
    var v11 int32 = -1
    var v12 int32 = v11
    switch "a" {
    case "a":
        switch v2 {
        case "foo bar":
            v4 = float32(float64(0.0))
        case "hello":
            var v13 uint16 = 65535
            _ = v13
        case "\n\t":
            v2 = v2
        }
    case "hello":
        var v14 int16 = (v5 * 42)
        _ = v14
    case "\n\t":
        v1 -= 127
    default:
        var v15 uint8 = 255
        _ = v15
    }
    delete(v3, 0)
    var v16 int32 = v11
    v1 += (v1 << ((818) & 63))
    if false {
        v1 += (v1 >> ((1) & 63))
        v11 += v16
    } else {
        fmt.Println(v5, v11, v5)
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := GenS1{F0: 880}
            v19 := []int{(fi17 & -1), fi17}
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    fmt.Println(v1, v12, v4)
    var v20 uint16 = 2
    if v6 {
        for fi21 := 0; fi21 < 14; fi21++ {
            v22 := GenS1{F0: 4294967295}
            switch int(v20) {
            case 9223372036854775807:
                v23 := []int{(fi21 - (fi21 + fi21)), (^(fi21)), fi21}
                _ = v23
            }
            _ = fi21
            _ = v22
        }
    }
    var v24 uint16 = 32767
    v25 := GenS1{F0: 0}
    var v26 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
