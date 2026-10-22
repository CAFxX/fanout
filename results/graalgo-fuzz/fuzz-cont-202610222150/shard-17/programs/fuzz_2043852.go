package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }

func main() {
    v1 := []string{"foo bar", (("a" + "foo bar") + "foo bar"), "hello"}
    v2 := []uint32{2, 4294967295, 0, 731}
    v3 := make(map[string]bool)
    if true {
        var v4 float32 = 1.5
        _ = v4
    } else {
        if true {
            if false {
                for fi5 := 0; fi5 < 1; fi5++ {
                    v6 := GenS1{F0: 10, F1: 1}
                    delete(v3, "\n\t")
                    for fi7 := 0; fi7 < 20; fi7++ {
                        var v8 string = (("" + "") + ("" + ("hello" + "foo bar")))
                        v2 = append(v2, 321)
                        _ = fi7
                        _ = v8
                    }
                    _ = fi5
                    _ = v6
                }
            } else {
                var v9 bool = (!(150 > 42))
                _ = v9
            }
        }
        var v10 uint64 = 285
        v10 = 377
        _ = v10
    }
    var v11 bool = (false && true)
    v11 = (62 != -128)
    switch 7 {
    case -9223372036854775808:
        var v12 int = 1
        _ = v12
    case 258:
        fmt.Println(v11)
    case 1:
        var v13 float32 = 0.1
        _ = v13
    }
    var v14 uint8 = 1
    delete(v3, "")
    if v11 {
        fmt.Println(v11, v11)
        v11 = (!v11)
    }
    var v15 uint8 = (v14 % ((231) | 1))
    v16 := GenS1{F0: 100, F1: 2}
    var v17 uint64 = 1
    v18 := GenS2{F0: 0, F1: 0, F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
