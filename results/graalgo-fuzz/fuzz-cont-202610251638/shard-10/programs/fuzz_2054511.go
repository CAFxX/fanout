package main

import (
    "fmt"
)


func main() {
    var v1 uint = 255
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 3
    v4 := []int{}
    v3 = 255
    var v5 string = "hello"
    if (!(!false)) {
        var v6 bool = (!(-32768 != int16(v2)))
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            v2 += -445
            var v9 string = "\n\t"
            var v10 uint64 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v5 = ((("\n\t" + ((v5 + "hello") + "hello")) + (v5 + v5)) + v5)
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    v4 = append(v4, -219)
    if false {
        v3 -= v3
    } else {
        v1 += (v1 & 2)
        fmt.Println(v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
