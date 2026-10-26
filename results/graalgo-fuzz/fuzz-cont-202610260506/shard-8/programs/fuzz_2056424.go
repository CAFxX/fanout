package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int32, p1 uint64, p2 uint32) int {
    var v5 int32 = p0
    var v6 uint32 = 2147483648
    p1 -= 1
    var v7 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 593
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]string)
    var v3 uint16 = 475
    var v4 uint16 = v3
    v8 := v2["a"]
    for fi9 := 0; fi9 < 13; fi9++ {
        v8 = v8
        v10 := []int{genFn0(-2147483648, 1, 129)}
        _ = fi9
        _ = v10
    }
    v11 := []int{529, -2147483648, genFn0(323, 829, 4294967295), 32768}
    for fi12 := 0; fi12 < 1; fi12++ {
        v8 = v8
        _ = fi12
    }
    var v13 uint8 = uint8(v4)
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS1{F0: 0}
        _ = fi14
        _ = v15
    }
    fmt.Println(v1)
    v2[v8] = ((("hello" + (v8 + v8)) + v8) + v8)
    if (1 > 18446744073709551615) {
        v16 := v2[""]
        fmt.Println(v13)
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        v11 = append(v11, genFn0((v1 - v1), 0, 2147483648))
        for fi18 := 0; fi18 < 13; fi18++ {
            v1 = v1
            v19 := GenS1{F0: 0}
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
