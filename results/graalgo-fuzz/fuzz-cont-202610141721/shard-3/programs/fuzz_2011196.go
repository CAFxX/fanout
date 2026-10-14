package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() int8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{fi7, (fi7 / ((fi7) | 1)), fi7, fi7}
        var v9 string = ("\n\t" + "\n\t")
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 1
    v10 -= (v10 ^ 1)
    var v11 string = ((("\n\t" + "") + "\n\t") + "a")
    v12 := GenS1{F0: (16 == 65535)}
    _ = v10
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 bool = (0.1 != 0.0)
    var v2 int16 = 42
    v3 := make(map[string]string)
    var v4 uint32 = 2147483648
    var v5 int64 = 0
    var v6 int16 = (v2 * v2)
    var v13 uint64 = 1
    fmt.Println(v4, v4)
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := []int{}
        _ = len(v15)
        _ = fi14
        _ = v15
    }
    var v16 string = ""
    v17 := v3[v16]
    var v18 int16 = (v2 % ((v6) | 1))
    v6 += (v18 << ((18446744073709551615) & 63))
    if v1 {
        v19 := v3[v16]
        var v20 uint32 = (v4 % ((2) | 1))
        _ = v19
        _ = v20
    }
    fmt.Println(v17)
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
