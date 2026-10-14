package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    v2 := []int{128, -9223372036854775808, 345}
    v3 := make(map[uint32]int)
    var v4 string = "hello"
    v5 := []uint32{2, 827, 394, 0}
    var v6 int32 = -2147483648
    fmt.Println(v4)
    var v7 uint64 = 18446744073709551615
    switch v4 {
    case "\n\t":
        delete(v3, 42)
    }
    delete(v3, 938)
    switch ("hello" + ((("foo bar" + v4) + v4) + (v4 + ""))) {
    case "hello":
        var v8 uint32 = 326
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                if false {
                    v4 = v4
                    v6 += v6
                } else {
                    v7 = 2
                }
                var v11 uint16 = 1
                fmt.Println(v7)
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    var v12 float32 = (3.14159 / (1e10 / 3.14159))
    if (!(!false)) {
        v4 = v4
        _ = len(v5)
    }
    v7 += v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
