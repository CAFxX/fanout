package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int8; F1 bool }
func genFn0(p0 uint16) int8 {
    if (!(854 == 758)) {
        if true {
            p0 = (p0 / ((p0) | 1))
            var v4 int16 = 1000
            _ = v4
        }
    } else {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 string = ("a" + "foo bar")
            _ = fi5
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 0, F1: 1}
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 120
    _ = p0
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := []string{"\n\t", "foo bar"}
    var v10 int = 42
    if false {
        var v11 int = (v10 ^ (v10 >> ((598) & 63)))
        _ = v11
    }
    var v12 int16 = -207
    v12 += 1000
    if true {
        v13 := GenS1{F0: -128, F1: -1}
        _ = v13
    } else {
        v14 := GenS2{F0: 127, F1: (true || true)}
        fmt.Println(v12, v12)
        _ = v14
    }
    switch (v10 / ((v10) | 1)) {
    case 1:
        var v15 int8 = 89
        _ = v15
    default:
        switch "hello" {
        case "a":
            switch v10 {
            case 1:
                if false {
                    var v16 int = v10
                    v17 := GenS1{F0: -31, F1: 233}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v16
                    _ = v17
                }
            case -809:
                var v18 float32 = 1.5
                _ = v18
            case 128:
                var v19 int8 = genFn0(1)
                _ = v19
            }
        default:
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 uint64 = 42
                v21 = v21
                _ = fi20
                _ = v21
            }
        }
    }
    if true {
        var v22 int8 = -1
        var v23 float64 = -2.25
        _ = v22
        _ = v23
    } else {
        v10 = (v10 >> ((65535) & 63))
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        delete(v1, "a")
        v25 := GenS2{F0: -84, F1: false}
        var v26 uint32 = 2147483647
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
}
