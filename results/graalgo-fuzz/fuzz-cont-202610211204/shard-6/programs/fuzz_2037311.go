package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 272
    var v3 uint8 = 0
    var v4 int8 = 127
    var v5 float32 = 1e10
    var v6 uint16 = 2
    v1["\n\t"] = -398
    v3 -= (v3 ^ (v3 + 127))
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{}
        var v9 int = 2147483647
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint16 = (v6 ^ 609)
    for fi11 := 0; fi11 < 17; fi11++ {
        if true {
            v2 -= (v2 << ((529) & 63))
            fmt.Println(v10, v5)
            if ((v6 * v10) < (v6 % ((v6) | 1))) {
                v12 := v1[("" + ("hello" + ("foo bar" + "\n\t")))]
                _ = v12
            }
        }
        _ = fi11
    }
    var v13 int16 = -32768
    v14 := []int{4294967295, 128, 1}
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 uint32 = 2147483648
        _ = fi15
        _ = v16
    }
    v13 -= -668
    var v17 int8 = (v4 - -1)
    var v18 bool = (false || (true && ((!(true || true)) && true)))
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v6, v18)
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 int64 = 1
            v6 = 604
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    fmt.Println(v2, v10, v13)
    v6 += v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
