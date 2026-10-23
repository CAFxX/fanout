package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 string = (("" + "hello") + "")
    var v3 int16 = -32768
    var v4 int8 = 0
    v2 = v2
    var v5 int32 = -1
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int16 = (v3 * (v3 % ((v3) | 1)))
            var v9 uint8 = 2
            var v10 int16 = 42
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        v2 = v2
        _ = fi6
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := []int{fi11, fi11, 2147483647, fi12}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 int64 = 1
    v15 := GenS1{F0: 2147483648}
    v4 -= 127
    v16 := GenS1{F0: 1}
    var v17 int64 = (v14 * (v14 << ((0) & 63)))
    var v18 uint = 522
    if (!(v2 < v2)) {
        v19 := GenS1{F0: 2147483648}
        fmt.Println(v4, v4, v14)
        _ = v19
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
