package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    v2 := []uint32{2147483648}
    v3 := make(map[uint32]string)
    var v4 int = 4294967295
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    var v5 float32 = 1.5
    v1[v4] = ("foo bar" + "a")
    v4 += v4
    v4 += v4
    fmt.Println(v4)
    for fi6 := 0; fi6 < 2; fi6++ {
        if false {
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 bool = true
                var v9 int8 = 42
                if (3.14159 < 3.14159) {
                    var v10 uint = 192
                    v8 = (!(!((3.14159 * ((0.1 + 1e10) * float64(v5))) != ((1.5 * 0.0) - float64(v5)))))
                    _ = v10
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            if true {
                v4 += fi6
            }
            var v11 uint64 = 892
            _ = v11
        } else {
            v4 = v4
        }
        fmt.Println(v5, fi6, v4)
        v5 = (((3.14159 * v5) / v5) - float32((-2.25 - 1e10)))
        _ = fi6
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
