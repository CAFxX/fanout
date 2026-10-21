package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }
type GenS2 struct { F0 uint64; F1 string; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 3.14159
    var v3 int = 128
    v4 := make(map[uint32]uint32)
    var v5 int16 = -1
    v6 := []int{v3, v3, v3, v3}
    if false {
        v7 := GenS1{F0: (v5 > v5), F1: -93, F2: 4294967295}
        var v8 int32 = 356
        var v9 float64 = v2
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 string = "a"
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int64 = 42
        var v13 int32 = 333
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint = 18446744073709551615
    var v15 uint32 = 2147483647
    v16 := GenS1{F0: false, F1: -94, F2: v14}
    v17 := GenS1{F0: (-356 > 0), F1: 100, F2: (v14 << ((v14) & 63))}
    v18 := GenS1{F0: false, F1: -128, F2: v14}
    v14 -= (v14 ^ (v14 >> (((v14 << (((v14 & v14)) & 63))) & 63)))
    var v19 uint32 = (^(v15))
    if true {
        if (false && (!(-207 >= 1))) {
            if false {
                for fi20 := 0; fi20 < 16; fi20++ {
                    switch ("\n\t" + (("hello" + "") + ("\n\t" + (((("a" + (((("" + "a") + "a") + "hello") + (((("hello" + ("foo bar" + ("foo bar" + ("hello" + ("\n\t" + ("a" + "foo bar")))))) + "hello") + "") + "\n\t"))) + "foo bar") + "") + "a")))) {
                    case "hello":
                        v6 = append(v6, 370)
                    }
                    _ = len(v6)
                    _ = fi20
                }
            }
            fmt.Println(v19, v5, v19)
            v21 := v4[v19]
            _ = v21
        } else {
            var v22 bool = true
            v23 := GenS1{F0: false, F1: 0, F2: v14}
            _ = v22
            _ = v23
        }
    }
    v14 = (^(v14))
    for fi24 := 0; fi24 < 12; fi24++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if false {
            fmt.Println(v14)
            for fi25 := 0; fi25 < 14; fi25++ {
                v6 = append(v6, (fi24 >> ((607) & 63)))
                var v26 uint64 = 18446744073709551615
                _ = fi25
                _ = v26
            }
            fmt.Println(v3)
        }
        _ = fi24
    }
    _ = len(v1)
    var v27 uint16 = 255
    if (851 == 355) {
        _ = len(v1)
        for fi28 := 0; fi28 < 11; fi28++ {
            fmt.Println(v5, v3, v3)
            _ = fi28
        }
        fmt.Println(v14, v19)
    }
    v29 := GenS2{F0: 94, F1: ("\n\t" + "a"), F2: 42}
    v14 = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
