package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint, p2 bool) int64 {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint8 = 30
        _ = fi7
        _ = v8
    }
    var v9 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{}
    var v3 int8 = 1
    var v4 int64 = 42
    var v5 int8 = 100
    var v6 uint32 = 1
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v6)
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: 2}
            _ = fi11
            _ = v12
        }
        fmt.Println(v6, v3, v4)
        _ = fi10
    }
    v13 := GenS2{F0: 0}
    switch 22 {
    case 4294967295:
        _ = len(v2)
    }
    var v14 int32 = -2147483648
    v4 = (v4 + 1)
    if (("hello" + "foo bar") == "a") {
        var v15 bool = true
        _ = v15
    } else {
        v14 = (v14 % (((v14 | v14)) | 1))
    }
    switch "a" {
    case "":
        fmt.Println(v3)
    case "\n\t":
        switch ("\n\t" + "foo bar") {
        case "a":
            v4 += v4
        case "hello":
            v5 = v5
        default:
            switch 644 {
            case 421:
                for fi16 := 0; fi16 < 8; fi16++ {
                    fmt.Println(fi16, v5)
                    var v17 int8 = v5
                    _ = fi16
                    _ = v17
                }
            case 557:
                fmt.Println(v4, v6, v3)
            case -9223372036854775808:
                var v18 int = 1
                _ = v18
            }
        }
    default:
        var v19 uint = 255
        _ = v19
    }
    var v20 uint64 = 1
    v21 := GenS1{F0: 106}
    v14 = 0
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int8 = 7
        for fi24 := 0; fi24 < 8; fi24++ {
            v3 = v5
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
