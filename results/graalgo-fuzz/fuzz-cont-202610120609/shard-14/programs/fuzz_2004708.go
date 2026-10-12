package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int8 = -7
    v2 := []int{0, 32767, 864}
    var v3 int32 = 2147483647
    v4 := make(map[string]string)
    v5 := []uint32{202, 4294967295}
    v6 := GenS1{F0: 2}
    var v7 float32 = 1e10
    if ((0.0 > ((0.0 + 1e10) - 0.0)) || (!((!true) || false))) {
        fmt.Println(v3, v7, v1)
        v8 := GenS1{F0: 2}
        var v9 int64 = -344
        _ = v8
        _ = v9
    } else {
        switch ("foo bar" + "foo bar") {
        case "":
            _ = len(v5)
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    v1 += v1
    fmt.Println(v7)
    v5 = append(v5, 414)
    if true {
        var v10 bool = false
        _ = v10
    }
    v11 := GenS1{F0: 994}
    if (("\n\t" + ((("" + "hello") + "a") + "")) != ("" + "foo bar")) {
        _ = len(v2)
        v1 = 0
        for fi12 := 0; fi12 < 6; fi12++ {
            v3 -= int32(fi12)
            _ = fi12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
