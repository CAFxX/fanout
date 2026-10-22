package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int64 }
func genFn0() uint {
    v6 := []int{}
    if (false || (false || false)) {
        switch ("" + ("hello" + "hello")) {
        case "":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        v7 := GenS1{F0: 113, F1: 2147483647, F2: -9223372036854775808}
        var v8 uint32 = 2147483648
        _ = v7
        _ = v8
    }
    var v9 bool = true
    _ = v6
    _ = v9
    return 471
}

func main() {
    var v1 uint8 = 255
    var v2 float64 = 0.1
    v3 := make(map[string]int)
    var v4 int = 764
    var v5 int8 = 0
    var v10 int = (v4 ^ v4)
    v3[""] = v4
    fmt.Println(v5)
    v3[("" + "")] = v4
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: (v5 >> ((4) & 63)), F1: 1, F2: -917}
        if false {
            var v13 uint32 = 42
            fmt.Println(v10, v5)
            _ = v13
        } else {
            v14 := GenS1{F0: -128, F1: 1, F2: 42}
            v1 = (v1 << ((230) & 63))
            _ = v14
        }
        v15 := GenS1{F0: v5, F1: 0, F2: 9223372036854775807}
        _ = fi11
        _ = v12
        _ = v15
    }
    v5 = (v5 & v5)
    var v16 int = v4
    v10 += v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
