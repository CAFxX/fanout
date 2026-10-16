package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int }
func genFn0() int {
    v7 := GenS1{F0: -1}
    v8 := []int{}
    var v9 uint8 = 125
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int8 = -85
    var v2 uint8 = 10
    var v3 int64 = -9223372036854775808
    var v4 int64 = 1
    var v5 int8 = 33
    v6 := []string{("foo bar" + "hello"), "a", "foo bar"}
    var v10 uint64 = 781
    v2 = (^(v2))
    var v11 uint16 = 32767
    v11 = v11
    _ = len(v6)
    v4 = 42
    v11 -= v11
    var v12 int16 = -1
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v13 int = -1
        _ = v13
    }
    _ = len(v6)
    var v14 uint32 = 2147483647
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v6 = append(v6, "")
    v12 -= v12
    fmt.Println(v4)
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 20; fi16++ {
            v5 -= v1
            _ = fi16
        }
        if false {
            v17 := GenS1{F0: -469}
            _ = v17
        } else {
            var v18 bool = false
            fmt.Println(v10)
            if false {
                for fi19 := 0; fi19 < 20; fi19++ {
                    v20 := GenS2{F0: (^(v1)), F1: 506}
                    if false {
                        var v21 float64 = (float64(float32(1e10)) + 0.1)
                        _ = v21
                    } else {
                        if (3.14159 >= float32(1.5)) {
                            v6 = append(v6, "\n\t")
                            var v22 uint8 = 86
                            _ = v22
                        }
                        _ = len(v6)
                    }
                    var v23 int16 = v12
                    _ = fi19
                    _ = v20
                    _ = v23
                }
                for fi24 := 0; fi24 < 15; fi24++ {
                    v10 -= v10
                    _ = fi24
                }
                v5 = (v5 | 1)
            }
            _ = v18
        }
        var v25 uint16 = (v11 / ((v11) | 1))
        _ = fi15
        _ = v25
    }
    if false {
        var v26 uint16 = v11
        for fi27 := 0; fi27 < 11; fi27++ {
            v12 += 32767
            if false {
                if false {
                    var v28 uint16 = (v26 + 666)
                    switch "\n\t" {
                    case "foo bar":
                        _ = len(v6)
                    case "\n\t":
                        fmt.Println(v4, v4, v28)
                    case "":
                        var v29 uint8 = v2
                        _ = v29
                    default:
                        v30 := GenS3{F0: fi27}
                        _ = v30
                    }
                    _ = v28
                } else {
                    var v31 float32 = 0.1
                    _ = v31
                }
                for fi32 := 0; fi32 < 14; fi32++ {
                    fmt.Println(fi32, v14, v2)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi32
                }
            }
            v5 = (-(v5))
            _ = fi27
        }
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
