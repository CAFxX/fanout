package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32) uint16 {
    if ((895 <= 1) && true) {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint64 = 1
            var v8 uint = 335
            _ = fi6
            _ = v7
            _ = v8
        }
        fmt.Println(p0)
    } else {
        v9 := []int{32768, 1, 1}
        _ = v9
    }
    var v10 int64 = 437
    _ = p0
    _ = v10
    return 32767
}
func genFn1() int8 {
    if false {
        v11 := GenS1{F0: 42}
        var v12 uint32 = 4294967295
        v13 := []int{160}
        _ = v11
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 134}
        var v15 uint32 = 104
        v15 -= 65535
        _ = v14
        _ = v15
    }
    var v16 uint32 = 0
    fmt.Println(v16)
    _ = v16
    return 1
}

func main() {
    var v1 int8 = 42
    var v2 int64 = 42
    v3 := []int{}
    var v4 int16 = 42
    v5 := make(map[string]string)
    switch "hello" {
    case "hello":
        v17 := GenS1{F0: (v4 - v4)}
        _ = v17
    case "\n\t":
        v2 += v2
    default:
        if (false || false) {
            v2 -= 0
            switch "" {
            case "\n\t":
                var v18 uint16 = 164
                _ = v18
            case "foo bar":
                v19 := GenS2{F0: 69}
                _ = v19
            case "":
                if false {
                    v20 := GenS2{F0: genFn1()}
                    _ = v20
                }
            }
        }
    }
    fmt.Println(v2, v2, v2)
    for fi21 := 0; fi21 < 11; fi21++ {
        if (42 != 42) {
            var v22 uint16 = 65535
            _ = v22
        }
        fmt.Println(v4, v2)
        v4 += -32768
        _ = fi21
    }
    v4 -= 1
    v23 := GenS1{F0: v4}
    var v24 uint64 = 2
    v25 := GenS2{F0: (v1 + (v1 << ((255) & 63)))}
    v24 -= v24
    v1 = -128
    v2 = 9223372036854775807
    var v26 int16 = 42
    switch "\n\t" {
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v5[""] = "foo bar"
    v2 = 0
    delete(v5, "foo bar")
    if len(v3) > 0 {
        v3[0] = 48
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        v1 += 1
        v28 := GenS1{F0: v26}
        fmt.Println(v2)
        _ = fi27
        _ = v28
    }
    v26 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
