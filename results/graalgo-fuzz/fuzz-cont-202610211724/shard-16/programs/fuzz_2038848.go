package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 570
    var v3 bool = true
    var v4 uint32 = 2147483648
    var v5 int32 = 1
    var v6 bool = ((v5 > (v5 - -917)) || (42 != (v5 * (v5 ^ (v5 | int32(v4))))))
    v3 = true
    v3 = v6
    var v7 int64 = -85
    var v8 bool = (!false)
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 127}
    if false {
        if ((v3 || ((!v3) || v8)) && v8) {
            v5 -= v5
        } else {
            var v12 float32 = float32((-2.25 / (float64((1e10 + ((3.14159 * (0.1 - -2.25)) - 1e10))) + ((3.14159 - 3.14159) + float64(0.0)))))
            v7 += (v7 | 1)
            _ = v12
        }
    } else {
        v13 := []int{}
        _ = v13
    }
    fmt.Println(v2)
    if (209 == int(v7)) {
        v6 = true
        v14 := []int{1, -1, 4294967295, 32767}
        _ = v14
    }
    v15 := GenS1{F0: 42}
    fmt.Println(v4)
    v16 := []int{127}
    v17 := GenS1{F0: -36}
    v18 := GenS2{F0: (v2 % ((v2) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
