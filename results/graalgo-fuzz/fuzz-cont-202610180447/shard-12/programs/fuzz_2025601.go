package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 0
    var v2 uint16 = 0
    var v3 int64 = 0
    v4 := []string{"a", (("hello" + "\n\t") + "\n\t"), "foo bar"}
    var v5 float64 = 0.1
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint = 774
        var v9 uint64 = v1
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = len(v4)
    var v10 int32 = 1
    if true {
        fmt.Println(v2, v1)
    }
    var v11 string = "a"
    v1 -= v1
    if ((v3 >> ((2) & 63)) > v3) {
        if true {
            if false {
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 float32 = 0.0
                    _ = fi12
                    _ = v13
                }
                fmt.Println(v5, v1, v10)
            }
            if (255 > 1) {
                switch int(v2) {
                case 9223372036854775807:
                    v6 -= (v6 >> ((0) & 63))
                case 371:
                    var v14 int64 = -319
                    _ = v14
                }
            } else {
                var v15 uint32 = 990
                _ = v15
            }
        }
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
