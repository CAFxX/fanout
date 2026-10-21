package main

import (
    "fmt"
)

func genFn0(p0 int8, p1 int, p2 int) uint8 {
    fmt.Println(p1)
    p1 = p2
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1() string {
    var v7 uint64 = 4294967295
    var v8 uint = 1
    _ = v7
    _ = v8
    return ("hello" + ("hello" + "hello"))
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{}
    var v3 uint8 = 255
    var v4 uint64 = 18446744073709551615
    var v5 uint16 = 181
    var v6 string = "\n\t"
    switch v6 {
    case "":
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 bool = (!(-911 != -9223372036854775808))
            _ = fi9
            _ = v10
        }
    }
    v6 = v6
    var v11 bool = true
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v12 uint8 = (v3 * v3)
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v4, v3, v11)
        _ = fi13
    }
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    v12 = v3
    var v14 int64 = 9223372036854775807
    var v15 uint8 = (^(v12))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
