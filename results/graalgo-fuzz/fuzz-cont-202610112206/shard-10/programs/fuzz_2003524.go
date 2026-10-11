package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 string; F1 bool; F2 string }
type GenS3 struct { F0 bool; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0(p0 int32, p1 int16, p2 uint32) uint8 {
    v4 := []int{658}
    var v5 string = "\n\t"
    var v6 uint = 1
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            fmt.Println(v5)
            _ = fi8
        }
        var v9 uint8 = 255
        v10 := GenS1{F0: (fi7 | fi7), F1: p0}
        _ = fi7
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int32 = 274
    v3 := []string{("\n\t" + "\n\t"), "\n\t"}
    switch (("" + "\n\t") + (("" + "foo bar") + "foo bar")) {
    case "":
        v11 := GenS2{F0: "a", F1: (-1 <= int16(v1)), F2: "a"}
        _ = v11
    default:
        var v12 uint32 = 2147483647
        _ = v12
    }
    var v13 int = -498
    if false {
        if false {
            v14 := GenS3{F0: false, F1: -1, F2: 2}
            v15 := GenS3{F0: ((-2.25 + (-2.25 * 1.5)) > (1e10 + 1e10)), F1: 0, F2: 1}
            _ = v14
            _ = v15
        }
        v16 := GenS4{F0: 42, F1: -32768}
        _ = v16
    } else {
        var v17 int32 = v2
        v3 = append(v3, "hello")
        _ = v17
    }
    fmt.Println(v2, v1)
    v18 := GenS3{F0: false, F1: (v1 & 0), F2: 1}
    var v19 bool = false
    if (!v19) {
        v20 := GenS4{F0: 255, F1: 785}
        _ = v20
    }
    if v19 {
        v19 = true
    }
    switch v13 {
    case -2147483648:
        switch v13 {
        case 0:
            for fi21 := 0; fi21 < 5; fi21++ {
                for fi22 := 0; fi22 < 14; fi22++ {
                    if (!(!(-1 < -1))) {
                        var v23 string = "hello"
                        v24 := GenS2{F0: v23, F1: (42 != 1), F2: "foo bar"}
                        fmt.Println(v2, v2, fi21)
                        v23 = "foo bar"
                        fmt.Println(v19, v2)
                        _ = v23
                        _ = v24
                    }
                    var v25 int16 = -1
                    v3 = append(v3, "")
                    _ = fi22
                    _ = v25
                }
                _ = fi21
            }
        case 489:
            var v26 float64 = 0.1
            _ = v26
        default:
            v27 := GenS4{F0: 196, F1: 1000}
            _ = v27
        }
    case 2147483647:
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 int = fi28
            v19 = v19
            var v30 int = v13
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    v19 = v19
    v19 = v19
    v3 = append(v3, "hello")
    var v31 int = v13
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v32 uint8 = 229
    fmt.Println(v31, v31, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
}
