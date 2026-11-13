package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() uint32 {
    v5 := []int{42, 2147483647, 90}
    var v6 float32 = (1.5 * 0.1)
    var v7 int = -187
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 42
    var v3 uint16 = 255
    v4 := []uint32{65535}
    if (33 != -1) {
        var v8 uint16 = (v3 % ((v3) | 1))
        _ = len(v4)
        v1[("\n\t" + ((("" + ("a" + "foo bar")) + (("hello" + ((((("" + "") + "\n\t") + "a") + "") + "a")) + "")) + (((("hello" + "a") + "\n\t") + ((("\n\t" + (("" + ("" + "a")) + ("" + ""))) + "") + "")) + "a")))] = 42
        _ = v8
    }
    v9 := GenS1{F0: 128}
    v2 -= v2
    v2 += 1
    var v10 string = ("" + (("hello" + ("hello" + "")) + ""))
    v10 = (v10 + ("foo bar" + "\n\t"))
    var v11 bool = false
    v2 = 995
    v3 = (v3 * (v3 + (v3 ^ 42)))
    v10 = v10
    fmt.Println(v11, v10)
    v12 := GenS1{F0: 52}
    if v11 {
        v11 = ((-75 > -8) || v11)
        v11 = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
