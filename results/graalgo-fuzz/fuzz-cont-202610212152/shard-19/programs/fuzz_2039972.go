package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int, p1 int8, p2 int) int8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        if true {
            p1 = p1
        }
        _ = fi7
    }
    if true {
        var v8 uint32 = 2147483648
        _ = v8
    }
    v9 := GenS1{F0: 119}
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            if (-632 == 2147483647) {
                var v12 int64 = -9223372036854775808
                v13 := []int{}
                _ = v12
                _ = v13
            }
            p2 = 9223372036854775807
            _ = fi11
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p1
}

func main() {
    var v1 float32 = 3.14159
    var v2 int = -9223372036854775808
    v3 := make(map[uint32]uint32)
    var v4 string = ("\n\t" + "")
    var v5 uint16 = 348
    var v6 int32 = 42
    v14 := GenS1{F0: 2}
    var v15 int8 = 42
    fmt.Println(v2)
    v16 := v3[958]
    v4 = ((v4 + (v4 + ("\n\t" + "hello"))) + "\n\t")
    var v17 uint32 = 644
    var v18 uint16 = v5
    v19 := GenS1{F0: 238}
    var v20 int8 = (v15 & (v15 + 83))
    var v21 int = -9223372036854775808
    v15 = v20
    if false {
        v4 = v4
        v21 = (^(v2))
        v18 += 0
    }
    v22 := []int{v21, v2, (^(v2))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
