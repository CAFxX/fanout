package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 int32 = 0
    v1 = append(v1, -9223372036854775808)
    switch 2147483647 {
    case 9223372036854775807:
        v5 := GenS1{F0: 694}
        _ = v5
    default:
        fmt.Println(v4)
    }
    switch (("a" + ("a" + "a")) + ("" + ("a" + "a"))) {
    case "a":
        if (!(!(570 >= 793))) {
            var v6 uint16 = 42
            v6 = (v6 * v6)
            _ = len(v1)
            _ = v6
        } else {
            v4 -= v4
            v4 = v4
            var v7 float32 = 1.5
            _ = v7
        }
    case "":
        delete(v3, "\n\t")
    }
    var v8 int32 = (v4 >> ((2) & 63))
    v9 := GenS1{F0: 1}
    if (!((v4 / (((v4 % ((-2147483648) | 1))) | 1)) > (v8 ^ 565))) {
        var v10 int16 = -898
        for fi11 := 0; fi11 < 3; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int8 = -82
                var v14 int64 = 9223372036854775807
                v14 = v14
                var v15 uint32 = 605
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            switch fi11 {
            case 9223372036854775807:
                v16 := v3["\n\t"]
                _ = v16
            }
            var v17 int64 = 9223372036854775807
            _ = fi11
            _ = v17
        }
        _ = v10
    } else {
        v8 -= v4
    }
    v8 = (v4 << ((1) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
