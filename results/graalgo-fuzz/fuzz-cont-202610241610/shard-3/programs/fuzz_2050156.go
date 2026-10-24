package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint) uint8 {
    p0 += p0
    if (497 == -32768) {
        fmt.Println(p0)
        v6 := GenS1{F0: 42, F1: 1}
        _ = v6
    }
    p0 += p0
    var v7 bool = true
    var v8 int8 = 127
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = float32((3.14159 / (0.1 / 0.1)))
    var v3 uint8 = 128
    v4 := make(map[int]uint32)
    var v5 int64 = 9223372036854775807
    v9 := GenS1{F0: 29, F1: 4294967295}
    switch 127 {
    case 32768:
        v2 -= v2
    default:
        v10 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615}
        _ = v10
    }
    delete(v1, "\n\t")
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(v2, v2)
        var v12 string = ("hello" + "\n\t")
        for fi13 := 0; fi13 < 3; fi13++ {
            if false {
                var v14 uint64 = 4294967295
                _ = v14
            }
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    var v15 int32 = -523
    v16 := v4[32768]
    v4[-2147483648] = 2147483648
    v3 = 53
    v17 := GenS1{F0: 0, F1: 330}
    v18 := GenS1{F0: 2, F1: 42}
    v19 := GenS1{F0: 4294967295, F1: 196}
    var v20 int32 = (v15 << ((820) & 63))
    fmt.Println(v2, v15)
    switch ("a" + "a") {
    case "a":
        v1["\n\t"] = "hello"
    case "hello":
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 string = "hello"
            _ = fi21
            _ = v22
        }
    }
    v23 := GenS1{F0: 1, F1: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
