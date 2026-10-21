package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int64, p1 uint16) int64 {
    var v6 uint32 = 2147483648
    v6 += v6
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 uint8 = 139
    var v2 uint8 = (v1 * v1)
    var v3 bool = false
    v4 := make(map[string]string)
    var v5 string = ""
    _ = len(v4)
    v2 += v2
    v4[((v5 + "foo bar") + "\n\t")] = v5
    v7 := GenS1{F0: 0}
    var v8 uint8 = 255
    fmt.Println(v1, v3, v3)
    v9 := []int{4294967295, -1, 9223372036854775807}
    var v10 string = "foo bar"
    v11 := GenS1{F0: -210}
    if (!((v3 && false) || v3)) {
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 int32 = 84
            _ = fi12
            _ = v13
        }
        _ = len(v9)
    }
    v8 -= (^(v2))
    v5 = ""
    var v14 uint32 = 238
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
