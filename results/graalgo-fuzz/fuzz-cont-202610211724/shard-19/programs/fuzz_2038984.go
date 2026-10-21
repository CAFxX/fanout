package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 string; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }

func main() {
    var v1 bool = (1000 > 1)
    var v2 uint8 = 128
    v3 := make(map[string]int)
    var v4 float32 = ((3.14159 / 1.5) / -2.25)
    var v5 int32 = 658
    var v6 uint64 = 4294967295
    switch -303 {
    case -480:
        v4 += float32((0.1 * ((-2.25 * (3.14159 / float64(((v4 - v4) - (1.5 - v4))))) - -2.25)))
    }
    v4 = v4
    if true {
        v6 -= v6
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int = fi7
            v6 = 42
            _ = fi7
            _ = v8
        }
    } else {
        if (349 != 2147483648) {
            var v9 uint32 = 2
            if (!(v2 <= v2)) {
                if false {
                    var v10 uint16 = 1
                    v9 = uint32(v6)
                    _ = v10
                } else {
                    v3["hello"] = 32768
                }
                for fi11 := 0; fi11 < 2; fi11++ {
                    v12 := GenS1{F0: v6}
                    _ = len(v3)
                    v13 := GenS2{F0: 42, F1: ("foo bar" + "hello"), F2: 2147483648}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            _ = v9
        } else {
            fmt.Println(v5)
            v14 := GenS3{F0: (135 >= 0)}
            _ = v14
        }
        v15 := GenS4{F0: 255}
        var v16 uint = 345
        _ = v15
        _ = v16
    }
    fmt.Println(v1, v6)
    var v17 uint = 42
    fmt.Println(v2, v5)
    v18 := []int{32768}
    v2 = v2
    v6 += 2
    v1 = (!(!false))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
