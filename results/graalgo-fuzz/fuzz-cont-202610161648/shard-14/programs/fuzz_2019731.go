package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = -306
    var v3 uint = 224
    v4 := []int{}
    var v5 float64 = 3.14159
    v6 := make(map[int]uint32)
    var v7 uint16 = 65535
    v8 := GenS1{F0: 18446744073709551615, F1: v3, F2: (v3 | v3)}
    v5 = 3.14159
    if false {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int64 = 0
            var v11 uint64 = 18446744073709551615
            _ = fi9
            _ = v10
            _ = v11
        }
        if false {
            fmt.Println(v7, v7)
            v12 := GenS1{F0: 4294967295, F1: v3, F2: 2}
            delete(v1, 4294967295)
            _ = v12
        } else {
            var v13 uint8 = 2
            _ = len(v1)
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(v5, v2, v2)
        var v15 int16 = v2
        _ = fi14
        _ = v15
    }
    var v16 int64 = -1
    v7 = uint16(v3)
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 uint8 = 42
        var v19 int8 = -128
        _ = fi17
        _ = v18
        _ = v19
    }
    switch ("a" + "") {
    case "foo bar":
        switch "" {
        case "foo bar":
            var v20 uint32 = 0
            _ = v20
        case "\n\t":
            v21 := GenS1{F0: 0, F1: v3, F2: v3}
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 bool = true
            var v25 int8 = 127
            _ = fi23
            _ = v24
            _ = v25
        }
        if false {
            if (!true) {
                v26 := GenS1{F0: 1, F1: (v3 - (v3 + v3)), F2: v3}
                _ = v26
            }
            switch "foo bar" {
            case "":
                for fi27 := 0; fi27 < 8; fi27++ {
                    var v28 int8 = 127
                    var v29 uint32 = 342
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            case "a":
                v16 = (v16 % (((v16 + (-(v16)))) | 1))
            case "foo bar":
                v6[fi22] = 65535
            }
            fmt.Println(v16, v16)
        }
        if true {
            var v30 int32 = 42
            fmt.Println(v5, fi22, v2)
            v31 := v1[77]
            _ = v30
            _ = v31
        }
        _ = fi22
    }
    var v32 int = 329
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
