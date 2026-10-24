package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 bool }
type GenS2 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 uint16 = 163
    v4 := make(map[string]int)
    v5 := []string{("a" + ((("\n\t" + ("\n\t" + "hello")) + ("" + "a")) + "")), "foo bar", "\n\t"}
    v3 = v3
    if (!(!(32767 > 42))) {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: -797, F1: 262, F2: true}
            var v8 int = 4294967295
            _ = fi6
            _ = v7
            _ = v8
        }
        fmt.Println(v3)
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint8 = 255
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 string = (("a" + ("\n\t" + "a")) + "hello")
        _ = fi11
        _ = v12
    }
    if (((537 >= 32768) || (!false)) || (v3 > v3)) {
        var v13 float32 = (1e10 / (1.5 * 1e10))
        if false {
            var v14 bool = false
            v13 += float32((1e10 / -2.25))
            _ = v14
        } else {
            switch "foo bar" {
            case "":
                var v15 int16 = 19
                _ = v15
            case "\n\t":
                v13 = v13
            case "a":
                fmt.Println(v13, v3)
            default:
                if (!(((1.5 / 3.14159) == 1e10) || ((!(true && ((!true) || (!(201 <= 245))))) && false))) {
                    v3 = 44
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 uint8 = 1
                        switch "foo bar" {
                        case "hello":
                            var v18 int64 = -52
                            _ = v18
                        case "foo bar":
                            v17 += 255
                        }
                        var v19 int16 = 1
                        _ = fi16
                        _ = v17
                        _ = v19
                    }
                } else {
                    var v20 float64 = 1.5
                    _ = len(v5)
                    v5 = append(v5, ("hello" + ""))
                    _ = v20
                }
            }
        }
        var v21 uint = 926
        _ = v13
        _ = v21
    } else {
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 uint64 = 2
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 float32 = 1.5
            v3 = v3
            v26 := v4["foo bar"]
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    if false {
        v3 = v3
        if (false || true) {
            _ = len(v1)
        }
        var v27 float64 = 1e10
        _ = v27
    }
    fmt.Println(v3)
    var v28 int8 = 1
    v29 := GenS2{F0: 2, F1: 306}
    var v30 float32 = 0.0
    v31 := v4["hello"]
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 uint = 1
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
