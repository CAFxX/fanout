package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
func genFn0() int8 {
    var v4 int16 = 32767
    switch 0 {
    case 32768:
        var v5 uint16 = 88
        _ = v5
    case 42:
        v6 := []int{600}
        _ = v6
    }
    var v7 float64 = 1.5
    _ = v4
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 bool, p2 int64) uint {
    var v8 int8 = -128
    p0 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p0
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int16 = 688
    var v3 int8 = -128
    switch (("\n\t" + "foo bar") + "foo bar") {
    case "hello":
        v1 = (^(v1))
    case "":
        var v9 uint64 = 42
        _ = v9
    }
    v10 := []int{}
    if len(v10) > 0 {
        v10[0] = 982
    }
    v2 = -32768
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 uint8 = 255
            fmt.Println(fi11, v1, v1)
            _ = fi12
            _ = v13
        }
        if false {
            var v14 float32 = 1e10
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 int32 = 0
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        _ = fi11
    }
    if true {
        switch -925 {
        case 1:
            var v17 uint8 = 42
            _ = v17
        case 128:
            var v18 int = 190
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 uint8 = 141
        _ = fi19
        _ = v20
    }
    switch ("\n\t" + "a") {
    case "":
        if true {
            v10 = append(v10, 935)
            fmt.Println(v3, v3)
        } else {
            v3 = -14
            v21 := GenS1{F0: 2147483648, F1: false}
            _ = v21
        }
    }
    fmt.Println(v1, v3)
    fmt.Println(v1, v3, v3)
    v10 = append(v10, 9223372036854775807)
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
