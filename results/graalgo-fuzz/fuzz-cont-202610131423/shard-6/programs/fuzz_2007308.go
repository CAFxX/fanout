package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint; F2 int8 }

func main() {
    v1 := []string{"a", ("hello" + "hello"), ""}
    v2 := make(map[int]int)
    v3 := []string{}
    v4 := make(map[int]int)
    var v5 string = "\n\t"
    v5 = v5
    var v6 bool = (false || (v5 > "foo bar"))
    if false {
        fmt.Println(v6, v6)
        switch -487 {
        case -1:
            v5 = "hello"
        default:
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 int = fi7
                delete(v2, 647)
                _ = fi7
                _ = v8
            }
        }
        if v6 {
            switch "foo bar" {
            case "a":
                for fi9 := 0; fi9 < 9; fi9++ {
                    v5 = v5
                    v6 = (0.0 != float64(-2.25))
                    _ = fi9
                }
            case "\n\t":
                v10 := GenS1{F0: 2, F1: -2147483648, F2: 2}
                _ = v10
            }
        } else {
            v6 = (0.1 != 1.5)
        }
    }
    var v11 string = "foo bar"
    fmt.Println(v5)
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 float64 = 0.0
        _ = fi12
        _ = v13
    }
    if v6 {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v11 = v5
        if v6 {
            switch v5 {
            case "":
                var v14 uint64 = 18446744073709551615
                _ = v14
            case "foo bar":
                v15 := GenS2{F0: v6, F1: 65535, F2: -42}
                _ = v15
            case "\n\t":
                var v16 float64 = 3.14159
                _ = v16
            default:
                var v17 int64 = 1
                _ = v17
            }
            switch "" {
            case "":
                fmt.Println(v6)
            case "a":
                if (false && true) {
                    var v18 string = (v5 + (v11 + v5))
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 13; fi19++ {
                switch (fi19 & (fi19 / ((2147483647) | 1))) {
                case -797:
                    var v20 int16 = 32767
                    _ = v20
                default:
                    var v21 uint = 706
                    _ = v21
                }
                fmt.Println(v11, v6, v11)
                _ = fi19
            }
        }
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 13; fi23++ {
            fmt.Println(fi22, fi22)
            _ = fi23
        }
        var v24 int = -1
        v6 = v6
        _ = fi22
        _ = v24
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = v11
    }
    v6 = v6
    var v25 uint16 = 445
    var v26 int8 = -14
    var v27 int16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
