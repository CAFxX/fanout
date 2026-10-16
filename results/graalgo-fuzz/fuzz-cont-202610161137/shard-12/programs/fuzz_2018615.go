package main

import (
    "fmt"
)


func main() {
    var v1 bool = (100 <= 1)
    var v2 uint32 = 1
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    v5 := []int{32767}
    if v1 {
        var v6 int32 = 529
        v2 += (^(v2))
        _ = v6
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v2, v1)
    }
    var v7 uint16 = 42
    v1 = v1
    fmt.Println(v7, v2)
    var v8 bool = v1
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(v8)
        _ = fi10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int16 = -3
        if (!v8) {
            if v8 {
                if len(v5) > 0 {
                    v5[0] = 249
                }
                v12 += 32767
            }
        } else {
            switch (fi11 % ((2147483647) | 1)) {
            case 9223372036854775807:
                for fi13 := 0; fi13 < 1; fi13++ {
                    v9 = v9
                    var v14 uint8 = 10
                    _ = fi13
                    _ = v14
                }
            }
            if ((v12 ^ (v12 << ((83) & 63))) == int16(fi11)) {
                var v15 string = v9
                var v16 int32 = 42
                _ = v15
                _ = v16
            } else {
                v5 = append(v5, fi11)
            }
            v12 += v12
        }
        var v17 int8 = 1
        _ = fi11
        _ = v12
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
