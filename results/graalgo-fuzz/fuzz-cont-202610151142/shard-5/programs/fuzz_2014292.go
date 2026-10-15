package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{}
    var v3 uint16 = 515
    v4 := make(map[string]int)
    var v5 int = -976
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int8 = -1
        if false {
            var v8 uint32 = 2147483648
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    var v9 string = "foo bar"
    v2 = append(v2, -2147483648)
    switch (v5 << ((4294967295) & 63)) {
    case 0:
        var v10 float64 = 3.14159
        _ = v10
    case 331:
        _ = len(v4)
    case 515:
        v1 = v1
    default:
        v9 = ((("foo bar" + v9) + v9) + (v9 + ((v9 + ("foo bar" + v9)) + v9)))
    }
    if (!(-9223372036854775808 != int64(v3))) {
        v11 := GenS1{F0: 255}
        _ = v11
    } else {
        if true {
            v4[v9] = (v5 - (v5 - (v5 & (v5 * (^(v5))))))
            if (false || false) {
                _ = len(v2)
                var v12 int32 = 1
                _ = v12
            } else {
                var v13 int64 = 482
                _ = v13
            }
        }
    }
    if len(v2) > 0 {
        v2[0] = v5
    }
    v5 = (v5 + (-(v5)))
    var v14 uint = 255
    var v15 string = v9
    var v16 int32 = 1
    var v17 float64 = 1.5
    var v18 uint32 = 233
    var v19 string = v15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
