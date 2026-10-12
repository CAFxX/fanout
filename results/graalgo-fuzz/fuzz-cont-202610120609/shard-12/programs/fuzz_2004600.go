package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int16, p1 uint32) bool {
    v6 := []int{4294967295, 0}
    p0 -= p0
    v7 := GenS1{F0: 32767}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int = 42
    v3 := []int{}
    var v4 uint32 = (v1 * v1)
    v5 := make(map[string]int)
    var v8 string = (("hello" + "\n\t") + "foo bar")
    v9 := GenS1{F0: 1}
    var v10 uint8 = 1
    v10 = (v10 << ((65535) & 63))
    var v11 bool = true
    _ = len(v5)
    _ = len(v5)
    v12 := GenS1{F0: (v2 << ((494) & 63))}
    v1 += (v1 >> ((615) & 63))
    var v13 string = (v8 + "hello")
    for fi14 := 0; fi14 < 13; fi14++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi14
    }
    var v15 uint8 = v10
    if ((!false) && (0 >= -1)) {
        v11 = (false && (1 < -2147483648))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
