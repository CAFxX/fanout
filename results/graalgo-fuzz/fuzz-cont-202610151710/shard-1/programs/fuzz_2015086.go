package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 bool; F1 uint64; F2 int }

func main() {
    var v1 float64 = 1e10
    var v2 bool = (-705 >= 797)
    var v3 float32 = (((((0.1 / 1.5) / -2.25) * ((0.1 / 1.0) - (((0.0 * 1e10) * 0.0) * -2.25))) - 1.5) - 3.14159)
    v4 := make(map[uint32]string)
    var v5 float64 = (v1 + (1.5 - -2.25))
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int32 = -1
        var v8 int64 = 289
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint = 18446744073709551615
        v3 = float32(3.14159)
        v11 := GenS1{F0: 42, F1: "hello"}
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v5, v5, v2)
    v5 = 1.5
    for fi12 := 0; fi12 < 4; fi12++ {
        if false {
            v2 = false
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 int = -977
                var v15 uint = 2
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            var v16 int32 = 646
            var v17 uint32 = 2147483648
            switch (fi12 ^ 414) {
            case 2147483647:
                v3 += v3
            case 0:
                if v2 {
                    switch (fi12 & 2147483647) {
                    case 198:
                        v18 := GenS1{F0: 245, F1: ("foo bar" + "\n\t")}
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 20; fi19++ {
                            var v20 float64 = v1
                            _ = fi19
                            _ = v20
                        }
                    }
                }
            }
            _ = v16
            _ = v17
        }
        v4[4294967295] = "\n\t"
        v21 := GenS2{F0: 42, F1: 317}
        _ = fi12
        _ = v21
    }
    var v22 string = "\n\t"
    for fi23 := 0; fi23 < 5; fi23++ {
        fmt.Println(v5, v3, v2)
        v24 := GenS1{F0: 705, F1: ""}
        _ = fi23
        _ = v24
    }
    switch ((v22 + "\n\t") + "") {
    case "\n\t":
        fmt.Println(v5)
    case "":
        switch "\n\t" {
        case "foo bar":
            var v25 uint = 857
            _ = v25
        default:
            v26 := GenS3{F0: v2, F1: 0, F2: 1}
            _ = v26
        }
    }
    var v27 bool = false
    switch -2147483648 {
    case 42:
        var v28 int32 = -2147483648
        _ = v28
    case 2147483647:
        v22 = ((v22 + v22) + v22)
    case 507:
        var v29 int32 = 0
        _ = v29
    default:
        v2 = true
    }
    var v30 bool = true
    var v31 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
