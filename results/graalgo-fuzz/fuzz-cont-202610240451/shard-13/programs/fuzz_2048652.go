package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() uint32 {
    v5 := []int{42, 9223372036854775807, 4294967295, 1}
    var v6 uint32 = 0
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 984
    var v3 float32 = (0.0 - 0.1)
    v4 := make(map[uint32]string)
    var v7 uint = 810
    switch 42 {
    case -2147483648:
        v8 := []int{2147483647, 32767, 4294967295}
        _ = v8
    case 127:
        if true {
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 uint16 = 0
                _ = fi9
                _ = v10
            }
        } else {
            v11 := []int{42}
            _ = v11
        }
    }
    switch "foo bar" {
    case "a":
        v3 -= float32((0.1 + 1.5))
    case "foo bar":
        var v12 uint32 = v2
        _ = v12
    case "":
        v3 = (float32(float64(v3)) - (0.0 - float32(((0.0 - (1e10 - (0.0 + 1e10))) - (1e10 / 1.0)))))
    default:
        fmt.Println(v3)
    }
    v2 = (v2 % (((v2 ^ 2)) | 1))
    var v13 uint16 = 0
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int64 = 417
        _ = fi14
        _ = v15
    }
    var v16 float32 = ((-2.25 + (1e10 * ((((v3 + 0.0) - v3) - v3) / 1.5))) * v3)
    var v17 bool = false
    for fi18 := 0; fi18 < 7; fi18++ {
        fmt.Println(fi18, v17)
        v19 := GenS1{F0: 1}
        if false {
            v20 := GenS1{F0: -2147483648}
            var v21 int8 = -91
            _ = v20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    v2 = v2
    fmt.Println(v2, v2)
    v17 = v17
    var v22 bool = (!v17)
    var v23 bool = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
}
