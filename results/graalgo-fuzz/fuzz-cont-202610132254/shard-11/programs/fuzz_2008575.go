package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }

func main() {
    v1 := []uint32{603, 0, 2147483648}
    v2 := make(map[int]string)
    var v3 float64 = 3.14159
    v4 := make(map[uint32]string)
    var v5 int64 = 494
    var v6 bool = (32767 >= -797)
    var v7 string = "foo bar"
    v8 := GenS1{F0: false, F1: 2, F2: 0}
    switch 9223372036854775807 {
    case 112:
        if (v6 && (100 > -7)) {
            if (!(v6 && v6)) {
                var v9 uint32 = 2147483648
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v3, v3)
                case "a":
                    var v10 uint = 4294967295
                    _ = v10
                case "":
                    if v6 {
                        for fi11 := 0; fi11 < 11; fi11++ {
                            v3 = v3
                            v12 := GenS1{F0: v6, F1: 0, F2: 127}
                            v3 = (v3 * v3)
                            _ = fi11
                            _ = v12
                        }
                    } else {
                        var v13 float32 = 3.14159
                        delete(v4, v9)
                        _ = v13
                    }
                }
                _ = v9
            }
            var v14 uint8 = 231
            _ = v14
        } else {
            v5 = 9223372036854775807
            var v15 uint16 = 1
            var v16 uint32 = 42
            _ = v15
            _ = v16
        }
    case 0:
        if (!true) {
            var v17 int8 = -1
            var v18 int8 = (^(v17))
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        if true {
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 bool = v6
                _ = fi20
                _ = v21
            }
        } else {
            var v22 uint32 = 953
            for fi23 := 0; fi23 < 5; fi23++ {
                switch (("a" + v7) + v7) {
                case "hello":
                    v7 = v7
                case "a":
                    v6 = (-73 > 127)
                default:
                    var v24 int64 = -748
                    _ = v24
                }
                _ = fi23
            }
            var v25 uint16 = 466
            _ = v22
            _ = v25
        }
        v26 := GenS1{F0: (92 > -4), F1: 2, F2: 56}
        var v27 int32 = -2147483648
        _ = fi19
        _ = v26
        _ = v27
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!("" != (v7 + "a"))) {
        v28 := GenS1{F0: false, F1: 0, F2: 147}
        _ = v28
    }
    v29 := GenS1{F0: v6, F1: 42, F2: 127}
    v30 := v2[4294967295]
    fmt.Println(v6, v5)
    var v31 int64 = v5
    for fi32 := 0; fi32 < 16; fi32++ {
        fmt.Println(v31)
        _ = fi32
    }
    fmt.Println(v3, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
