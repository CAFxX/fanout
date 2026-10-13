package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint, p1 string, p2 uint64) int16 {
    v5 := GenS1{F0: 42}
    v6 := []int{-1, 127, 32767}
    v7 := GenS1{F0: 42}
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 32767
}
func genFn1() int8 {
    var v8 uint16 = 32767
    var v9 uint = 4294967295
    switch 1 {
    case -9223372036854775808:
        v10 := GenS2{F0: -2147483648}
        _ = v10
    case 95:
        v9 = 867
    }
    var v11 uint64 = 717
    v11 = uint64(v9)
    _ = v11
    _ = v8
    _ = v9
    return -103
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 698
    var v3 string = ((("hello" + "hello") + "a") + "\n\t")
    var v4 int = 349
    if false {
        if (v4 <= 32767) {
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 int = (fi12 ^ v4)
                _ = fi12
                _ = v13
            }
        } else {
            v1 -= (v1 * v1)
            if true {
                var v14 float32 = 1e10
                _ = v14
            }
        }
    }
    v3 = v3
    if (int64(v4) < 793) {
        v15 := GenS1{F0: v1}
        fmt.Println(v1, v2, v2)
        _ = v15
    }
    v2 -= genFn0(2, (v3 + v3), 55)
    v16 := []int{-9223372036854775808}
    _ = len(v16)
    v4 = (^(v4))
    var v17 float64 = 3.14159
    var v18 uint32 = 110
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
