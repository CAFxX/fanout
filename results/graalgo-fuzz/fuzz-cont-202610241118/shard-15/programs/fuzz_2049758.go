package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]int)
    v2 := []string{"a", "foo bar", "hello"}
    var v3 uint64 = 0
    v4 := []int{-862, -235, 0, 127}
    var v5 int64 = 282
    var v6 int8 = 127
    _ = len(v1)
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int64 = int64(fi7)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v6 += 1
            var v10 uint8 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 string = ("" + "")
        _ = fi7
        _ = v11
    }
    var v12 uint64 = v3
    for fi13 := 0; fi13 < 3; fi13++ {
        v5 += (v5 >> ((42) & 63))
        var v14 float32 = 1e10
        _ = fi13
        _ = v14
    }
    delete(v1, "")
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 int16 = -1
        v2 = append(v2, "\n\t")
        _ = fi15
        _ = v16
    }
    var v17 uint32 = 0
    if true {
        var v18 int32 = 1
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 uint8 = 2
            for fi21 := 0; fi21 < 1; fi21++ {
                fmt.Println(v6, v18)
                v17 = uint32(v20)
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    v5 += (^(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
