package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 string) int16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        p0 = p0
        fmt.Println(p0, fi5)
        _ = fi5
    }
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return 1000
}

func main() {
    v1 := []string{"a", ""}
    var v2 string = "foo bar"
    var v3 uint16 = 1
    v4 := make(map[int]int)
    for fi6 := 0; fi6 < 16; fi6++ {
        switch fi6 {
        case 822:
            for fi7 := 0; fi7 < 3; fi7++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v2, fi7, v2)
                _ = fi7
            }
        case 42:
            fmt.Println(fi6, v3, v3)
        default:
            var v8 int16 = -32768
            _ = v8
        }
        _ = fi6
    }
    switch v2 {
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "\n\t":
        v4[4] = 2147483647
    case "a":
        v9 := GenS1{F0: 830}
        _ = v9
    }
    var v10 uint = 949
    var v11 int32 = 0
    v11 += 0
    var v12 uint32 = 2147483647
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v12, v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
