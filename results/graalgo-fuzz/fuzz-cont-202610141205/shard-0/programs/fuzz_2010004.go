package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float64 = ((1e10 * (float64(1.5) * 3.14159)) * -2.25)
    v2 := make(map[string]string)
    var v3 int64 = 1
    v4 := []int{42, -488, 0, -2147483648}
    v5 := GenS1{F0: 1, F1: 793, F2: false}
    if (!(0 == -2147483648)) {
        v6 := GenS2{F0: 42}
        _ = v6
    }
    v2[(("a" + ("a" + "a")) + "\n\t")] = "\n\t"
    if true {
        fmt.Println(v3, v3)
    }
    v2["a"] = "a"
    v7 := GenS3{F0: 1}
    switch 9223372036854775807 {
    case 9223372036854775807:
        for fi8 := 0; fi8 < 5; fi8++ {
            v2["foo bar"] = "hello"
            v3 = (v3 + v3)
            _ = fi8
        }
    }
    _ = len(v2)
    var v9 int32 = 457
    v3 += v3
    v3 -= 419
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := v2[("" + "a")]
        v2[(("" + "foo bar") + ("foo bar" + v11))] = (v11 + "")
        _ = fi10
        _ = v11
    }
    _ = len(v4)
    var v12 uint16 = 65535
    v12 += (v12 >> ((456) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
