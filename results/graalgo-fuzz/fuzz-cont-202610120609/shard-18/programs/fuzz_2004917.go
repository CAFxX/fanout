package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 int = 128
    var v3 uint = 751
    v4 := []string{""}
    v5 := make(map[int]bool)
    v6 := make(map[string]string)
    switch "" {
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "hello":
        v4 = append(v4, ("a" + ""))
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int32 = -470
        v8 = v8
        v3 -= 4294967295
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 42
    v9 = v9
    fmt.Println(v9, v9)
    v3 = (v3 * v3)
    switch -375 {
    case 23:
        v9 = 42
    case 4294967295:
        for fi10 := 0; fi10 < 10; fi10++ {
            if false {
                fmt.Println(v9)
                var v11 int16 = 1
                _ = v11
            }
            var v12 int64 = 0
            _ = fi10
            _ = v12
        }
    case 32767:
        v2 = (v2 & (v2 - 2147483647))
    }
    switch "\n\t" {
    case "hello":
        var v13 int64 = 1
        _ = v13
    case "\n\t":
        var v14 uint32 = 887
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
