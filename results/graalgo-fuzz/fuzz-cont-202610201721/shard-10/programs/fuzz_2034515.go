package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 float32 = float32(-2.25)
    var v2 uint16 = 32767
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 10; fi4++ {
        delete(v3, "\n\t")
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 float32 = (v1 / 1.0)
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    if true {
        var v7 uint8 = 8
        _ = v7
    } else {
        var v8 int8 = 42
        _ = v8
    }
    var v9 uint8 = 172
    v9 -= (v9 & (v9 << ((255) & 63)))
    var v10 uint16 = (v2 >> ((4294967295) & 63))
    if false {
        var v11 bool = ((false || false) && true)
        v12 := []int{4294967295, -543, 127, -9223372036854775808}
        _ = v11
        _ = v12
    } else {
        fmt.Println(v2, v9, v9)
        var v13 string = "a"
        _ = v13
    }
    v2 += 896
    v14 := GenS1{F0: v9}
    v15 := v3[("" + "")]
    for fi16 := 0; fi16 < 17; fi16++ {
        v15 += fi16
        if true {
            var v17 string = ("\n\t" + ("hello" + "\n\t"))
            _ = v17
        } else {
            _ = len(v3)
            var v18 float32 = float32((1e10 - -2.25))
            _ = v18
        }
        _ = fi16
    }
    v19 := []int{}
    v20 := GenS2{F0: 769}
    switch "\n\t" {
    case "":
        _ = len(v19)
    }
    v21 := GenS2{F0: 1}
    v9 = v9
    switch ("" + ("hello" + "hello")) {
    case "\n\t":
        var v22 bool = false
        _ = v22
    case "a":
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(v15, v15)
            fmt.Println(v2, v10, v15)
            for fi24 := 0; fi24 < 4; fi24++ {
                v2 += v10
                var v25 int = 192
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
    }
    if true {
        for fi26 := 0; fi26 < 14; fi26++ {
            v10 += (v2 + (v2 * v10))
            if true {
                var v27 uint64 = 0
                _ = v27
            }
            _ = fi26
        }
    } else {
        var v28 uint8 = uint8(v10)
        var v29 uint32 = 441
        v2 -= (v10 | (v10 % (((v10 & (^(v10)))) | 1)))
        _ = v28
        _ = v29
    }
    if false {
        var v30 int = -1
        var v31 uint = 133
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v9
}
