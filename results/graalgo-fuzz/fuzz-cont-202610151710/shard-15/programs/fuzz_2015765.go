package main

import (
    "fmt"
)

func genFn0(p0 uint) uint16 {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = 1
    }
    var v8 bool = false
    _ = p0
    _ = v7
    _ = v8
    return 129
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]bool)
    v3 := []string{v1}
    var v4 bool = (!true)
    v5 := make(map[string]int)
    var v6 float32 = -2.25
    fmt.Println(v6)
    if v4 {
        v4 = (!(0 != 18446744073709551615))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v9 := v5[((((v1 + "a") + (("" + "a") + "hello")) + v1) + v1)]
        _ = v9
    }
    var v10 uint16 = genFn0(689)
    v11 := v5[v1]
    var v12 int64 = -1
    if true {
        v11 -= (v11 / (((v11 / ((v11) | 1))) | 1))
        fmt.Println(v12, v4)
    }
    v10 -= 105
    fmt.Println(v1, v6, v12)
    v3 = append(v3, ((v1 + (v1 + ("foo bar" + "\n\t"))) + "\n\t"))
    var v13 int8 = -22
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
