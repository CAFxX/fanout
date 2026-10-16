package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint16 }
func genFn0() bool {
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: "a", F1: 0, F2: 0}
        var v7 float32 = 3.14159
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: "hello", F1: 84, F2: 42}
    var v9 int32 = -1
    var v10 float64 = (-2.25 * -2.25)
    _ = v10
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]uint32)
    var v3 int64 = -1
    var v4 float32 = float32(1e10)
    v4 += ((v4 / v4) - v4)
    var v11 int16 = 32767
    v3 = (v3 << ((506) & 63))
    v12 := GenS1{F0: "", F1: -1, F2: 828}
    v13 := GenS1{F0: ("foo bar" + "\n\t"), F1: -21, F2: 2}
    v11 += (v11 << ((255) & 63))
    var v14 uint16 = 65535
    var v15 int32 = 42
    for fi16 := 0; fi16 < 17; fi16++ {
        delete(v2, fi16)
        _ = fi16
    }
    switch -2 {
    case 936:
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := []int{int(v15), 9223372036854775807}
            var v19 float64 = float64(1.5)
            var v20 uint64 = v1
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
    case 33:
        v21 := []int{}
        _ = v21
    default:
        var v22 uint32 = 65535
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
