package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8) int16 {
    var v7 bool = false
    var v8 uint64 = 1
    p0 += p0
    v9 := GenS1{F0: 2}
    v7 = (!(("\n\t" < "hello") || (!(2147483647 != 128))))
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 911
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 uint8 = 127
    var v4 int = 32768
    var v5 int8 = 0
    var v6 uint32 = 830
    v10 := GenS1{F0: 2}
    var v11 uint8 = (^(v3))
    var v12 string = "\n\t"
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := v1[v12]
        if (9223372036854775807 < 454) {
            var v15 uint64 = 914
            v3 -= v11
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    if true {
        var v16 int16 = genFn0((v5 >> ((uint(v6)) & 63)))
        _ = v16
    }
    v6 = v6
    v17 := GenS2{F0: v3}
    switch "a" {
    case "foo bar":
        delete(v1, v12)
    case "hello":
        v11 = 71
    default:
        v18 := []int{-124, (v4 - (v4 % ((32767) | 1))), 42, v4}
        _ = v18
    }
    fmt.Println(v12, v11, v5)
    var v19 int8 = v5
    for fi20 := 0; fi20 < 3; fi20++ {
        if (!((!true) || (1 != -1))) {
            var v21 uint64 = 1
            v11 = (v11 << ((255) & 63))
            var v22 int64 = 42
            _ = v21
            _ = v22
        }
        var v23 uint16 = 42
        if false {
            v12 = ""
        }
        _ = fi20
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
