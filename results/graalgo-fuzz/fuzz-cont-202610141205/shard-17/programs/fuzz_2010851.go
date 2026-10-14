package main

import (
    "fmt"
)

func genFn0(p0 uint16, p1 uint32) int64 {
    fmt.Println(p1, p0)
    var v7 uint32 = p1
    _ = p0
    _ = p1
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int8 = 127
    var v2 int32 = 522
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    var v5 uint16 = 42
    v6 := []string{"", ("a" + ""), ((("" + "hello") + ("foo bar" + (("a" + "") + ("\n\t" + "foo bar")))) + ("\n\t" + "a"))}
    v5 = (v5 + v5)
    _ = len(v6)
    if ((v5 % (((v5 - (v5 ^ v5))) | 1)) <= 0) {
        var v8 int16 = 1
        _ = v8
    }
    var v9 int32 = (v2 ^ v2)
    v9 -= (v2 + (v2 << ((42) & 63)))
    v2 = v9
    v5 += v5
    var v10 float64 = ((((3.14159 + 0.1) / (1.5 - 0.1)) / 0.1) + -2.25)
    v9 -= (v2 + (v2 << ((167) & 63)))
    fmt.Println(v1, v1)
    var v11 uint = 42
    v12 := v3[0]
    var v13 int8 = 0
    v2 += (v2 / (((v9 - (v2 ^ 1))) | 1))
    v9 += 0
    v14 := v3[65535]
    var v15 int64 = -1
    v3[65535] = ((1 > -9223372036854775808) || (true && true))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
