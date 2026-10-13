package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool }
func genFn0(p0 int16, p1 uint8, p2 int32) uint8 {
    var v4 int16 = 847
    v5 := []int{32767}
    for fi6 := 0; fi6 < 9; fi6++ {
        switch (fi6 | 0) {
        case -2147483648:
            for fi7 := 0; fi7 < 17; fi7++ {
                if true {
                    var v8 uint16 = 42
                    _ = v8
                } else {
                    p1 -= p1
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                v9 := GenS1{F0: fi7, F1: ((((("\n\t" + "foo bar") + "") + ("\n\t" + "hello")) + "\n\t") + "\n\t")}
                switch "" {
                case "a":
                    fmt.Println(p0)
                default:
                    p2 += 0
                }
                _ = fi7
                _ = v9
            }
        }
        var v10 float64 = float64(3.14159)
        _ = len(v5)
        _ = fi6
        _ = v10
    }
    switch -2147483648 {
    case 1:
        p1 -= 254
    case 861:
        var v11 float64 = 1e10
        _ = v11
    case 32767:
        var v12 uint = 285
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p1
}

func main() {
    var v1 uint16 = 826
    v2 := make(map[uint32]uint32)
    var v3 int = 9223372036854775807
    fmt.Println(v1, v3)
    v13 := GenS1{F0: v3, F1: ("foo bar" + "\n\t")}
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS1{F0: fi14, F1: ""}
        v16 := GenS1{F0: fi14, F1: ("" + "a")}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int = v3
    v18 := GenS1{F0: (v17 - v17), F1: ("foo bar" + "")}
    v17 = v17
    v19 := v2[252]
    v20 := GenS1{F0: -2147483648, F1: "\n\t"}
    var v21 uint32 = (v19 >> ((163) & 63))
    switch "hello" {
    case "hello":
        v3 += v3
    case "\n\t":
        var v22 int8 = 0
        _ = v22
    case "":
        fmt.Println(v19)
    default:
        if ((v19 > 2147483648) && true) {
            for fi23 := 0; fi23 < 10; fi23++ {
                v3 = v17
                if (((("\n\t" + ("a" + "hello")) + "hello") + "\n\t") <= ("" + ("foo bar" + "a"))) {
                    switch "hello" {
                    case "foo bar":
                        v24 := GenS2{F0: (!(!true))}
                        _ = v24
                    default:
                        var v25 uint16 = (v1 | 42)
                        _ = v25
                    }
                    var v26 uint32 = (v19 * (^(v19)))
                    v27 := v2[v26]
                    _ = v26
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 7; fi28++ {
                        switch (v3 << ((920) & 63)) {
                        case 855:
                            v29 := []int{}
                            _ = v29
                        case -1:
                            v30 := []int{(fi28 << ((65535) & 63)), v3, 127}
                            _ = v30
                        case 127:
                            delete(v2, v21)
                        }
                        var v31 int = 9223372036854775807
                        _ = fi28
                        _ = v31
                    }
                }
                v19 += v21
                _ = fi23
            }
            var v32 uint8 = 26
            _ = v32
        }
    }
    var v33 int8 = -85
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 float32 = -2.25
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
}
