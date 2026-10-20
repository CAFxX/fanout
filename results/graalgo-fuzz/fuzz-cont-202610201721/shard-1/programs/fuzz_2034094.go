package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 string) uint32 {
    p0 = ("a" + p0)
    fmt.Println(p0)
    _ = p0
    return 869
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 0
    v3 := []int{}
    var v4 uint16 = 2
    var v5 float64 = 1e10
    v6 := v1["a"]
    v2 = v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v7 := GenS1{F0: -1}
    fmt.Println(v6)
    if false {
        var v8 int32 = 38
        _ = v8
    } else {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint16 = v4
            var v11 int32 = 2147483647
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int8 = 0
        _ = v12
    }
    var v13 uint64 = 1
    var v14 uint64 = (v13 >> ((0) & 63))
    var v15 string = ("a" + (("hello" + ((("" + "hello") + "\n\t") + ("" + ("foo bar" + "\n\t")))) + ("a" + (((("foo bar" + "") + "a") + (("a" + ("" + "foo bar")) + ("" + "hello"))) + ("" + ("a" + "hello"))))))
    v1[v15] = 944
    var v16 int32 = -1
    v3 = append(v3, int(v6))
    var v17 int8 = 100
    delete(v1, (v15 + v15))
    v13 = 945
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
