package main

import (
    "fmt"
)

func genFn0() uint32 {
    var v6 int8 = -128
    var v7 bool = ((!(-103 != 32768)) || ((-2147483648 == 42) || (false && false)))
    v8 := []int{9223372036854775807}
    fmt.Println(v7, v7)
    _ = v6
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    v1 := []string{((("a" + ("foo bar" + "a")) + (("a" + "hello") + "foo bar")) + "a")}
    var v2 uint = 265
    var v3 int16 = 1000
    v4 := make(map[uint32]int)
    v5 := []string{"a", "\n\t"}
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(v2, fi9)
        _ = fi9
    }
    fmt.Println(v2, v2)
    fmt.Println(v3, v3)
    var v10 int64 = 201
    fmt.Println(v10, v2, v3)
    switch "\n\t" {
    case "\n\t":
        v11 := v4[uint32(v2)]
        _ = v11
    case "a":
        switch -675 {
        case -521:
            _ = len(v4)
        }
    case "foo bar":
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 int64 = 42
            v2 = v2
            _ = fi12
            _ = v13
        }
    }
    fmt.Println(v2)
    var v14 uint64 = 544
    var v15 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
