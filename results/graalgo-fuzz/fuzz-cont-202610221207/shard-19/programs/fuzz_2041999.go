package main

import (
    "fmt"
)

func genFn0(p0 uint64, p1 int8) uint16 {
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            p1 = p1
        } else {
            var v8 int = (fi7 >> ((65535) & 63))
            v9 := []int{127, fi7, v8}
            _ = len(v9)
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    p1 += (p1 ^ -1)
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 uint64 = 926
    v2 := make(map[int]string)
    var v3 int = 128
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    var v6 int16 = 1
    v3 += v3
    var v10 uint16 = 42
    var v11 int32 = -2147483648
    switch "hello" {
    case "hello":
        v1 -= (v1 & uint64(v10))
    }
    var v12 int8 = -13
    v11 = int32(v12)
    var v13 int16 = v6
    fmt.Println(v13)
    v6 = v13
    fmt.Println(v10, v12, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
