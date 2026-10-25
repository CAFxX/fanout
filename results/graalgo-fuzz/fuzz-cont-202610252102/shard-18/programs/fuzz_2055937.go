package main

import (
    "fmt"
)


func main() {
    v1 := []string{("foo bar" + "foo bar"), ("hello" + "a")}
    var v2 float32 = ((1.5 - 1e10) / 1.0)
    var v3 bool = (953 >= 42)
    var v4 uint = 65535
    var v5 float32 = v2
    v2 -= v5
    if (839 >= (^(v4))) {
        fmt.Println(v3, v3, v3)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        v5 = (v2 + (-2.25 * v5))
        _ = fi6
    }
    if ((0.0 / v2) < (v2 / float32(1e10))) {
        fmt.Println(v3, v2, v4)
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        v5 = (((v5 - (3.14159 * (v2 + 1.5))) - v5) / v2)
        if v3 {
            var v8 int16 = 758
            var v9 uint32 = 4294967295
            var v10 int64 = 9223372036854775807
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    var v11 uint64 = 146
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
