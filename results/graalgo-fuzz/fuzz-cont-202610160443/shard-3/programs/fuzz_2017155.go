package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
func genFn0(p0 uint16, p1 uint8, p2 bool) uint32 {
    p2 = true
    fmt.Println(p0, p2, p0)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 985
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 19
    var v3 string = "a"
    var v4 uint32 = (v2 >> ((791) & 63))
    v5 := []int{-1, 32768, -410, -573}
    var v6 int32 = 1
    var v8 int32 = (v6 * -759)
    if false {
        var v9 uint = 1
        _ = v9
    } else {
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v10 uint8 = 0
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int32 = 0
        var v13 uint16 = 255
        var v14 int32 = v12
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS1{F0: 127, F1: (fi15 ^ -227), F2: -44}
        var v17 uint64 = 482
        if len(v5) > 0 {
            v5[0] = fi15
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v4)
    v3 = ((v3 + v3) + "\n\t")
    v8 -= -2147483648
    v18 := GenS1{F0: -1, F1: -197, F2: 0}
    var v19 float64 = 3.14159
    var v20 uint32 = genFn0(792, 128, (4294967295 > 1))
    v20 -= (^(v20))
    var v21 string = ("hello" + ("" + ""))
    var v22 int32 = v6
    v23 := v1[403]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
