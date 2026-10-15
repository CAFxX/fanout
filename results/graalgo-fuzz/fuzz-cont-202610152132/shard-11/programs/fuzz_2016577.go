package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int8 }
type GenS2 struct { F0 uint; F1 string; F2 uint16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 bool; F1 int16; F2 uint16 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 uint = 0
    var v2 string = ""
    var v3 uint32 = 2
    v4 := make(map[string]int)
    var v5 uint8 = 128
    v6 := []int{-809}
    v7 := GenS1{F0: "\n\t", F1: -934, F2: 0}
    v2 = ""
    if (true && true) {
        if false {
            var v8 int8 = 42
            if true {
                var v9 bool = (3.14159 < (0.1 / 0.1))
                _ = v9
            }
            v10 := GenS2{F0: (^(v1)), F1: v2, F2: 32767}
            _ = v10
            _ = v8
        } else {
            switch -2147483648 {
            case 4294967295:
                var v11 int = -653
                _ = v11
            case -281:
                v12 := GenS3{F0: 2147483647, F1: "a"}
                _ = v12
            default:
                v2 = v2
            }
            var v13 bool = (!false)
            if (v13 && (((v2 + "hello") > v2) || v13)) {
                if (!v13) {
                    if (9223372036854775807 == 322) {
                        if (!((^(v3)) <= 347)) {
                            v2 = v2
                            fmt.Println(v1, v5)
                        }
                        v13 = v13
                        var v14 uint16 = 0
                        _ = v14
                    }
                }
                v15 := GenS4{F0: ((!(100 >= 60)) || v13), F1: 311, F2: 0}
                _ = v15
            } else {
                if v13 {
                    var v16 uint16 = uint16(v5)
                    _ = v16
                } else {
                    var v17 uint32 = v3
                    _ = v17
                }
                var v18 string = "foo bar"
                _ = len(v4)
                _ = v18
            }
            _ = v13
        }
        fmt.Println(v5)
    }
    var v19 int = 2147483647
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v1)
    switch -661 {
    case -550:
        if (!((v3 > v3) && false)) {
            var v20 int16 = 1
            _ = v20
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 float64 = (0.0 / (1.5 + -2.25))
                    fmt.Println(v5)
                    v24 := GenS5{F0: (v2 + ("a" + ((("\n\t" + "a") + "a") + ("foo bar" + ((("" + (v2 + (v2 + v2))) + v2) + v2))))), F1: v2}
                    var v25 int64 = 42
                    if (v25 > v25) {
                        var v26 uint32 = (v3 * 692)
                        switch v2 {
                        case "foo bar":
                            delete(v4, "hello")
                            _ = len(v6)
                            fmt.Println(v1)
                        case "\n\t":
                            v4["foo bar"] = (fi22 / (((fi21 << ((v1) & 63))) | 1))
                        }
                        v1 -= v1
                        _ = v26
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi21
            }
        }
    case 1:
        var v27 int8 = 42
        _ = v27
    default:
        for fi28 := 0; fi28 < 7; fi28++ {
            v29 := GenS1{F0: "a", F1: (fi28 >> ((v1) & 63)), F2: -68}
            delete(v4, v2)
            _ = fi28
            _ = v29
        }
    }
    var v30 uint16 = 0
    fmt.Println(v2, v1)
    v2 = v2
    v31 := v4[v2]
    v32 := GenS6{F0: 574}
    for fi33 := 0; fi33 < 15; fi33++ {
        v1 = v1
        v6 = append(v6, (v19 | v31))
        _ = len(v6)
        _ = fi33
    }
    var v34 uint = uint(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
