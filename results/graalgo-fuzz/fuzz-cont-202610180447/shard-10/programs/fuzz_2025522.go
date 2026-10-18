package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 0.0
    var v3 int64 = 1
    var v4 uint32 = 1
    var v5 string = (("hello" + ((("hello" + (("a" + ((("a" + "hello") + ("" + "")) + (("\n\t" + "foo bar") + "a"))) + "\n\t")) + ("hello" + (("a" + "hello") + ("foo bar" + "hello")))) + "foo bar")) + ("foo bar" + ""))
    var v6 int32 = -1
    fmt.Println(v6, v2)
    v3 -= (v3 << ((787) & 63))
    var v7 uint8 = 42
    var v8 uint64 = 845
    if (v4 > uint32(v8)) {
        fmt.Println(v4, v5)
    } else {
        v9 := GenS1{F0: -1, F1: 113, F2: v6}
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(v4, fi10, v3)
            if false {
                for fi11 := 0; fi11 < 19; fi11++ {
                    fmt.Println(fi11, v7, v4)
                    fmt.Println(fi11, v7)
                    _ = fi11
                }
            }
            _ = fi10
        }
        _ = v9
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(v3, v6, fi12)
        for fi13 := 0; fi13 < 1; fi13++ {
            delete(v1, ("\n\t" + v5))
            v14 := []int{(^(fi13)), fi13}
            _ = fi13
            _ = v14
        }
        v2 = v2
        _ = fi12
    }
    v15 := GenS1{F0: -120, F1: 128, F2: 0}
    if ((1e10 * 0.0) >= 0.1) {
        v7 += v7
    }
    v1[("hello" + v5)] = true
    v3 = -9223372036854775808
    v16 := GenS1{F0: 60, F1: 9223372036854775807, F2: 254}
    var v17 int32 = (v6 + v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
