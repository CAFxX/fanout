package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
func genFn0() bool {
    switch "foo bar" {
    case "a":
        for fi4 := 0; fi4 < 14; fi4++ {
            v5 := []int{(fi4 * (fi4 / ((fi4) | 1))), fi4, (-(fi4)), fi4}
            switch ("foo bar" + "a") {
            case "\n\t":
                var v6 uint64 = 553
                _ = v6
            default:
                var v7 int32 = 2147483647
                _ = v7
            }
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 string = "foo bar"
                if false {
                    fmt.Println(v9, fi8)
                    switch ("\n\t" + v9) {
                    case "a":
                        var v10 string = (v9 + "\n\t")
                        _ = v10
                    case "foo bar":
                        var v11 int32 = -1
                        _ = v11
                    case "hello":
                        v12 := GenS1{F0: 78, F1: "foo bar"}
                        _ = v12
                    }
                    if (0 > 1) {
                        if len(v5) > 0 {
                            v5[0] = 630
                        }
                        var v13 uint8 = 2
                        v9 = ((((v9 + "a") + "hello") + v9) + v9)
                        _ = v13
                    } else {
                        if ((-1 > -32768) || (true && true)) {
                            if (9223372036854775807 < -852) {
                                fmt.Println(fi8, v9)
                                fmt.Println(fi8, v9, fi4)
                                v14 := GenS1{F0: 2, F1: ""}
                                _ = v14
                            } else {
                                v15 := GenS1{F0: 290, F1: (v9 + "")}
                                _ = v15
                            }
                        }
                    }
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi8
                _ = v9
            }
            _ = fi4
            _ = v5
        }
    }
    var v16 int16 = 971
    var v17 int16 = v16
    var v18 bool = false
    _ = v16
    _ = v17
    _ = v18
    return v18
}

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 1
    var v3 int16 = -32768
    var v19 uint16 = 0
    v2 -= (v19 | 42)
    switch "\n\t" {
    case "":
        v20 := []int{2147483647, -9, -125}
        _ = v20
    case "a":
        var v21 float32 = 1.5
        _ = v21
    }
    if false {
        fmt.Println(v3)
        var v22 uint = 65535
        _ = v22
    }
    fmt.Println(v1, v3, v2)
    switch 9223372036854775807 {
    case 4294967295:
        v23 := []int{32767}
        _ = v23
    case 829:
        if (((3.14159 + 1.5) / (v1 + 1e10)) >= ((1.5 * 0.1) - -2.25)) {
            v3 = 32767
        }
    case -64:
        var v24 bool = false
        _ = v24
    }
    v25 := []int{}
    fmt.Println(v3, v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
}
