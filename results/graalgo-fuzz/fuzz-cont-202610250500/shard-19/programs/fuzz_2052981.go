package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0() string {
    var v6 uint = 65535
    v7 := GenS1{F0: -940, F1: (v6 | (v6 - 481)), F2: 2147483647}
    for fi8 := 0; fi8 < 15; fi8++ {
        v6 += 577
        _ = fi8
    }
    v9 := []int{0, 32768}
    v10 := GenS2{F0: 0}
    _ = v10
    _ = v6
    _ = v7
    _ = v9
    return "a"
}

func main() {
    v1 := []string{"foo bar"}
    var v2 string = "foo bar"
    var v3 uint32 = 675
    v4 := []int{128, 222, 128, 4294967295}
    var v5 uint64 = 2
    v5 += (v5 | (v5 >> ((42) & 63)))
    var v11 bool = (1 <= 112)
    var v12 uint = 18446744073709551615
    v2 = v2
    v12 = v12
    v1 = append(v1, (v2 + (v2 + "hello")))
    v11 = false
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v13 int32 = 1
    if (42 != (v12 << ((v12) & 63))) {
        v11 = (!(128 == 42))
    } else {
        fmt.Println(v2, v13, v11)
        fmt.Println(v3, v5)
        fmt.Println(v13, v5, v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
