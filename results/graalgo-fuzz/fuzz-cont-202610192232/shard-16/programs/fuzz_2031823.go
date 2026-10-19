package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int, p1 int16) uint32 {
    p1 = (p1 ^ 1)
    p1 = p1
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "\n\t"
    var v3 string = ((("a" + v2) + "\n\t") + v2)
    var v4 uint8 = 42
    v5 := make(map[int]bool)
    v6 := make(map[int]int)
    switch v3 {
    case "a":
        for fi7 := 0; fi7 < 5; fi7++ {
            fmt.Println(v3)
            var v8 uint64 = 18446744073709551615
            _ = fi7
            _ = v8
        }
    case "foo bar":
        fmt.Println(v4, v3, v4)
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint = 18446744073709551615
        for fi11 := 0; fi11 < 9; fi11++ {
            if false {
                var v12 string = v2
                var v13 int = 42
                fmt.Println(v3, fi11)
                _ = v12
                _ = v13
            }
            v14 := GenS1{F0: 4294967295}
            _ = fi11
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    v3 = v3
    fmt.Println(v2, v3, v3)
    v15 := GenS1{F0: 2147483648}
    var v16 int16 = 810
    v17 := GenS1{F0: 0}
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 int32 = 2147483647
        v3 = ((v2 + v2) + v3)
        var v20 uint = 529
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
