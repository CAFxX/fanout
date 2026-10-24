package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{1}
    var v2 bool = false
    var v3 bool = (0 < 502)
    v4 := []int{-1}
    v5 := make(map[string]int)
    if (v2 || (255 != 18446744073709551615)) {
        for fi6 := 0; fi6 < 2; fi6++ {
            v2 = (v3 && v3)
            v2 = (-890 != 0)
            for fi7 := 0; fi7 < 13; fi7++ {
                v2 = (1 != 1)
                _ = fi7
            }
            _ = fi6
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v3 = v3
    var v8 float32 = 1e10
    var v9 int64 = -9223372036854775808
    var v10 float64 = 1e10
    var v11 bool = false
    fmt.Println(v2, v9, v8)
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 uint64 = 995
        delete(v5, ("" + "foo bar"))
        if v3 {
            var v14 int16 = 113
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    v2 = (-1 != 9223372036854775807)
    v15 := v5[("foo bar" + "\n\t")]
    var v16 uint8 = 226
    v8 = -2.25
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
