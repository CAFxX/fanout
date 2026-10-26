package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint }
type GenS2 struct { F0 int32; F1 uint16; F2 int }

func main() {
    var v1 int = 899
    var v2 bool = (668 >= -527)
    var v3 uint32 = 1
    v4 := make(map[int]bool)
    v5 := []string{}
    var v6 int32 = -338
    var v7 float32 = float32(1.5)
    v8 := GenS1{F0: 1, F1: 32767, F2: 42}
    if false {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int16 = 1
            v7 = (float32(1.5) - v7)
            _ = fi9
            _ = v10
        }
    }
    if true {
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint64 = 2
            var v13 int32 = 1
            _ = fi11
            _ = v12
            _ = v13
        }
        if v2 {
            if v2 {
                v3 = v3
            } else {
                var v14 int8 = -128
                _ = v14
            }
            for fi15 := 0; fi15 < 14; fi15++ {
                if v2 {
                    var v16 string = "a"
                    _ = v16
                } else {
                    var v17 int = v1
                    v18 := GenS1{F0: 715, F1: 42, F2: 18446744073709551615}
                    fmt.Println(v7, v3)
                    _ = v17
                    _ = v18
                }
                fmt.Println(v2, v6)
                _ = fi15
            }
        } else {
            v2 = v2
            v7 = 3.14159
        }
        fmt.Println(v7)
    } else {
        v7 = (3.14159 + ((3.14159 + v7) - v7))
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 uint16 = 2
        _ = fi19
        _ = v20
    }
    switch ("hello" + ("" + ("\n\t" + "hello"))) {
    case "hello":
        v21 := GenS1{F0: 554, F1: -740, F2: 1}
        _ = v21
    case "":
        var v22 int8 = 55
        _ = v22
    default:
        if v2 {
            for fi23 := 0; fi23 < 19; fi23++ {
                v2 = false
                var v24 bool = v2
                _ = fi23
                _ = v24
            }
        } else {
            var v25 uint8 = 42
            var v26 int64 = -1
            _ = v25
            _ = v26
        }
    }
    v1 = int(v3)
    for fi27 := 0; fi27 < 12; fi27++ {
        fmt.Println(v6, v7)
        _ = fi27
    }
    var v28 bool = true
    delete(v4, v1)
    var v29 int8 = 104
    if (347 >= 4294967295) {
        var v30 int8 = v29
        v7 = (float32(-2.25) * (1e10 + v7))
        v31 := GenS1{F0: -1, F1: 1000, F2: 985}
        _ = v30
        _ = v31
    } else {
        v32 := GenS2{F0: v6, F1: 32767, F2: 2147483647}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v28)
        _ = v32
    }
    var v33 uint16 = 63
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
