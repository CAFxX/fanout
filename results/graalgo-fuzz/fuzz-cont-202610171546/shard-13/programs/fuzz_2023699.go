package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    v5 := []string{"a", "", "hello"}
    _ = len(v1)
    var v6 int = -1
    if len(v5) > 0 {
        v5[0] = ("" + ("a" + "hello"))
    }
    var v7 string = ""
    if len(v5) > 0 {
        v5[0] = (v7 + v7)
    }
    fmt.Println(v6, v6)
    v8 := GenS1{F0: 149, F1: "\n\t", F2: 42}
    for fi9 := 0; fi9 < 2; fi9++ {
        if false {
            var v10 float64 = -2.25
            _ = len(v5)
            switch v7 {
            case "foo bar":
                if false {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        if ((1 > -1) || (515 <= 18446744073709551615)) {
                            fmt.Println(v7)
                        } else {
                            v6 += fi11
                        }
                        v12 := GenS1{F0: 42, F1: "a", F2: 0}
                        var v13 uint8 = 117
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
            case "\n\t":
                var v14 int32 = 1
                _ = v14
            }
            _ = v10
        }
        _ = fi9
    }
    if (!(float64(3.14159) <= 1e10)) {
        v7 = v7
        switch "\n\t" {
        case "a":
            var v15 uint64 = 342
            _ = v15
        default:
            v7 = ""
        }
        _ = len(v5)
    }
    v16 := GenS1{F0: 2, F1: v7, F2: 370}
    var v17 float64 = 1.5
    v18 := v4[1]
    v19 := GenS2{F0: false, F1: 42, F2: 39}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
