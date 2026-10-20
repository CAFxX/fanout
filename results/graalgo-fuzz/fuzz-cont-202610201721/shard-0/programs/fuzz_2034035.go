package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 4294967295
    v3 := []string{"foo bar", "foo bar", ("hello" + "")}
    var v4 uint32 = 2
    var v5 int32 = 1
    var v6 uint8 = 1
    var v7 int = 2147483647
    fmt.Println(v4)
    v3 = append(v3, "")
    var v8 int = 981
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v9 int64 = 9223372036854775807
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(v5, v7, v9)
        v5 -= 474
        v9 -= (^(v9))
        _ = fi10
    }
    delete(v1, ("" + (("" + ("hello" + "foo bar")) + ("\n\t" + "\n\t"))))
    v4 -= 1
    var v11 float32 = 1e10
    fmt.Println(v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
