package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() uint16 {
    var v7 int16 = -1
    v8 := GenS1{F0: "hello"}
    var v9 uint16 = 32767
    var v10 bool = ((true || false) && false)
    var v11 int32 = 2147483647
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 uint8, p1 int32) int {
    v12 := []int{-466, 0, 656, 32768}
    if len(v12) > 0 {
        v12[0] = -1
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    _ = p0
    _ = p1
    _ = v12
    return 9223372036854775807
}

func main() {
    var v1 string = "a"
    var v2 int8 = -1
    v3 := []string{(v1 + ((v1 + (((v1 + (("\n\t" + (v1 + ((((("hello" + ("a" + "foo bar")) + (v1 + v1)) + "hello") + "") + ((((v1 + v1) + (v1 + "hello")) + ("hello" + v1)) + ("\n\t" + (v1 + v1)))))) + "hello")) + v1) + ("" + v1))) + "hello")), v1, v1, ("foo bar" + v1)}
    v4 := make(map[uint32]string)
    var v5 uint64 = 18446744073709551615
    var v6 uint64 = (v5 - v5)
    _ = len(v3)
    switch v1 {
    case "foo bar":
        fmt.Println(v6)
    case "":
        v13 := GenS1{F0: "foo bar"}
        _ = v13
    case "hello":
        var v14 string = "\n\t"
        _ = v14
    default:
        var v15 int = 2147483647
        _ = v15
    }
    v2 -= 101
    delete(v4, 492)
    var v16 string = ""
    fmt.Println(v6)
    v5 += 315
    if (9223372036854775807 == 377) {
        fmt.Println(v16, v16, v5)
        fmt.Println(v16, v1)
        var v17 int = 2147483647
        _ = v17
    } else {
        if len(v3) > 0 {
            v3[0] = (v1 + (v1 + v16))
        }
    }
    var v18 int64 = -906
    v19 := GenS1{F0: (v1 + (v1 + ("hello" + v16)))}
    var v20 int64 = (v18 & v18)
    delete(v4, uint32(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
