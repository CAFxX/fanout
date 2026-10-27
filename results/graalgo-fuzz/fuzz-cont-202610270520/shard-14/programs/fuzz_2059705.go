package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0(p0 int) uint32 {
    v6 := []int{430, p0, -2147483648, p0}
    if (!(4294967295 < 0)) {
        var v7 int = -1
        v7 += p0
        _ = v7
    } else {
        var v8 uint64 = 120
        _ = v8
    }
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 bool = false
    v2 := []uint32{1, 2147483648}
    v3 := []string{(("a" + "") + "\n\t"), ("a" + (("\n\t" + "a") + ("" + ("" + "a")))), ("a" + "")}
    var v4 int16 = -891
    var v5 string = "foo bar"
    fmt.Println(v4, v4)
    var v9 uint32 = 703
    var v10 bool = (false || v1)
    v4 += v4
    v3 = append(v3, (v5 + "a"))
    var v11 float32 = (1.5 - (3.14159 * (0.0 + 3.14159)))
    var v12 int64 = 42
    if len(v2) > 0 {
        v2[0] = genFn0(9223372036854775807)
    }
    v13 := GenS1{F0: 1, F1: (v5 + "foo bar"), F2: 255}
    _ = len(v3)
    switch -649 {
    case 128:
        for fi14 := 0; fi14 < 7; fi14++ {
            switch "a" {
            case "a":
                fmt.Println(v12, v4, v4)
            default:
                switch fi14 {
                case -902:
                    var v15 int8 = 127
                    _ = v15
                case -717:
                    switch fi14 {
                    case 9223372036854775807:
                        var v16 int64 = -765
                        _ = v16
                    case 4294967295:
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                    default:
                        v4 += (v4 - 1000)
                    }
                case -1:
                    var v17 uint64 = 40
                    _ = v17
                }
            }
            var v18 float32 = v11
            v9 += v9
            _ = fi14
            _ = v18
        }
    case 4294967295:
        var v19 int8 = -20
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 string = (v5 + ("\n\t" + ("\n\t" + "a")))
        fmt.Println(v21, v12, v21)
        switch v21 {
        case "":
            switch (-(fi20)) {
            case -9223372036854775808:
                for fi22 := 0; fi22 < 18; fi22++ {
                    v11 = v11
                    v23 := GenS1{F0: (fi22 ^ -58), F1: v5, F2: 1}
                    _ = fi22
                    _ = v23
                }
            }
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "a":
            if len(v3) > 0 {
                v3[0] = v5
            }
        default:
            for fi24 := 0; fi24 < 5; fi24++ {
                if v10 {
                    fmt.Println(v10, v12)
                    if v10 {
                        v25 := GenS2{F0: 32767, F1: v4}
                        _ = v25
                    } else {
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                    }
                }
                _ = len(v2)
                v1 = (!false)
                _ = fi24
            }
        }
        _ = fi20
        _ = v21
    }
    var v26 uint64 = 4294967295
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
