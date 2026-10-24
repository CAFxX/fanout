package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint8, p1 uint) int16 {
    var v5 uint32 = 445
    v5 = (v5 - v5)
    v6 := []int{1, 4294967295}
    v6 = append(v6, 0)
    p1 = 65535
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 248
}

func main() {
    var v1 string = (("a" + "\n\t") + "")
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    var v7 int64 = 0
    v7 += v7
    v8 := []int{}
    for fi9 := 0; fi9 < 19; fi9++ {
        if (v7 == v7) {
            switch ("foo bar" + (("\n\t" + v1) + v1)) {
            case "a":
                v7 -= v7
            }
            delete(v2, (fi9 / ((fi9) | 1)))
            var v10 int64 = -9223372036854775808
            _ = v10
        } else {
            _ = len(v8)
            v11 := GenS1{F0: ""}
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v1 = v1
        _ = fi12
    }
    var v13 uint64 = 1
    delete(v4, 895)
    for fi14 := 0; fi14 < 9; fi14++ {
        v7 = (v7 >> ((65535) & 63))
        v13 -= v13
        v13 += 42
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
