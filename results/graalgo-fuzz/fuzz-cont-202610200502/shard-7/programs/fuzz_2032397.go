package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int16; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
func genFn0() uint8 {
    var v5 int64 = 42
    v6 := GenS1{F0: 1000}
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 4294967295
    var v3 int16 = 0
    v4 := make(map[string]uint32)
    fmt.Println(v2, v3)
    v3 = 652
    v7 := GenS1{F0: v3}
    v1 -= v1
    var v8 int64 = -424
    var v9 string = "foo bar"
    if ((v2 << ((42) & 63)) != (^(v2))) {
        var v10 int32 = 0
        for fi11 := 0; fi11 < 20; fi11++ {
            switch "foo bar" {
            case "foo bar":
                v12 := GenS2{F0: (v3 + v3), F1: (v3 / ((v3) | 1)), F2: fi11}
                _ = v12
            default:
                v13 := GenS1{F0: 0}
                _ = v13
            }
            if (!true) {
                switch 127 {
                case -9223372036854775808:
                    v14 := v4[(v9 + v9)]
                    _ = v14
                default:
                    var v15 int64 = v8
                    _ = v15
                }
                v9 = ((("a" + "a") + "hello") + "\n\t")
                v16 := []int{32767}
                _ = v16
            } else {
                v17 := []int{(^(fi11))}
                _ = len(v17)
                _ = v17
            }
            fmt.Println(v1, v9, v9)
            _ = fi11
        }
        _ = v10
    }
    v9 = v9
    v18 := GenS3{F0: -1}
    switch "foo bar" {
    case "a":
        if true {
            if false {
                switch v9 {
                case "a":
                    v8 = v8
                case "":
                    for fi19 := 0; fi19 < 8; fi19++ {
                        v20 := GenS4{F0: (v8 + v8)}
                        v21 := GenS4{F0: -9223372036854775808}
                        for fi22 := 0; fi22 < 20; fi22++ {
                            v23 := GenS4{F0: (v8 - v8)}
                            v24 := []int{(fi19 >> ((1) & 63)), fi22, 4294967295, fi19}
                            var v25 int8 = v1
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                case "\n\t":
                    fmt.Println(v3, v2, v1)
                }
                for fi26 := 0; fi26 < 6; fi26++ {
                    v3 = 789
                    if false {
                        if false {
                            v9 = v9
                            v2 = v2
                        }
                        v27 := v4["foo bar"]
                        for fi28 := 0; fi28 < 1; fi28++ {
                            var v29 float32 = 1e10
                            fmt.Println(fi28, v29, v3)
                            _ = fi28
                            _ = v29
                        }
                        _ = v27
                    }
                    _ = fi26
                }
            }
        } else {
            var v30 int16 = v3
            _ = v30
        }
    case "foo bar":
        var v31 uint16 = 255
        _ = v31
    case "":
        var v32 uint16 = 2
        _ = v32
    }
    v33 := GenS4{F0: v8}
    var v34 float64 = 1e10
    if (!true) {
        v2 += v2
        fmt.Println(v3, v3)
    }
    var v35 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
