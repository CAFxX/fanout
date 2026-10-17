package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int64 }
func genFn0(p0 int) uint32 {
    p0 = (p0 - 9223372036854775807)
    switch -521 {
    case 2147483647:
        if false {
            var v6 uint = 2
            _ = v6
        } else {
            fmt.Println(p0)
        }
    case -298:
        if true {
            var v7 uint = 424
            var v8 int8 = 56
            v9 := []int{p0, (p0 + p0)}
            _ = v7
            _ = v8
            _ = v9
        }
    case 128:
        v10 := GenS1{F0: "\n\t", F1: 127, F2: 9223372036854775807}
        _ = v10
    default:
        var v11 uint16 = 0
        _ = v11
    }
    switch "\n\t" {
    case "hello":
        p0 += p0
    case "\n\t":
        var v12 uint32 = 0
        _ = v12
    }
    _ = p0
    return 133
}

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    var v3 uint64 = 137
    var v4 int8 = 1
    var v5 string = "foo bar"
    switch 369 {
    case 287:
        v13 := GenS1{F0: (("foo bar" + (v5 + v5)) + v5), F1: -22, F2: 1}
        _ = v13
    case -511:
        for fi14 := 0; fi14 < 7; fi14++ {
            _ = len(v2)
            _ = fi14
        }
    case 2147483647:
        var v15 float32 = 0.1
        _ = v15
    }
    switch -940 {
    case 0:
        v16 := GenS1{F0: v5, F1: v4, F2: 9223372036854775807}
        _ = v16
    default:
        v5 = v5
    }
    _ = len(v2)
    var v17 int8 = 42
    v1 = true
    _ = len(v2)
    for fi18 := 0; fi18 < 8; fi18++ {
        v3 += v3
        for fi19 := 0; fi19 < 14; fi19++ {
            switch "a" {
            case "":
                var v20 bool = ((v17 == v17) && false)
                _ = v20
            case "hello":
                var v21 uint8 = 128
                _ = v21
            }
            v22 := GenS1{F0: v5, F1: v17, F2: -9223372036854775808}
            if (875 <= 0) {
                var v23 float32 = 1e10
                v24 := GenS1{F0: ((("hello" + v5) + "") + v5), F1: (v17 & 0), F2: 9223372036854775807}
                var v25 int = (fi19 - fi19)
                switch "hello" {
                case "hello":
                    v25 -= fi18
                case "\n\t":
                    switch (fi19 % ((fi19) | 1)) {
                    case 32767:
                        if (v1 || false) {
                            var v26 uint32 = 4294967295
                            var v27 uint64 = 4294967295
                            fmt.Println(fi18, fi18, v3)
                            _ = v26
                            _ = v27
                        }
                    case -287:
                        if v1 {
                            if (0 < -240) {
                                v25 = fi18
                            }
                            var v28 int16 = 171
                            var v29 float32 = (v23 / v23)
                            _ = v28
                            _ = v29
                        }
                    }
                case "foo bar":
                    v5 = v5
                default:
                    v30 := []int{}
                    _ = v30
                }
                var v31 int32 = 0
                _ = v23
                _ = v24
                _ = v25
                _ = v31
            } else {
                v17 = (v17 & (v17 >> ((255) & 63)))
            }
            _ = fi19
            _ = v22
        }
        _ = fi18
    }
    delete(v2, "a")
    var v32 int32 = -2147483648
    if (42 >= 0) {
        fmt.Println(v32)
    }
    var v33 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
