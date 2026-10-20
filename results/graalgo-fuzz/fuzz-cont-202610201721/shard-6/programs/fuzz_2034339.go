package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS4 struct { F0 string }
func genFn0() uint8 {
    var v6 uint = 18446744073709551615
    fmt.Println(v6)
    var v7 float64 = float64(1.5)
    v8 := GenS1{F0: (!false), F1: (v6 % ((v6) | 1))}
    var v9 uint = v6
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{1}
    var v3 int16 = 1
    v4 := make(map[int]string)
    var v5 int32 = -953
    v10 := GenS1{F0: false, F1: 1}
    v11 := v4[32768]
    v12 := GenS1{F0: true, F1: 42}
    var v13 int16 = (v3 ^ v3)
    fmt.Println(v3)
    if (v11 < v11) {
        v5 = (^(v5))
        var v14 int8 = 127
        if false {
            fmt.Println(v13, v11, v3)
            if false {
                switch (((v11 + "\n\t") + ("" + v11)) + v11) {
                case "foo bar":
                    _ = len(v2)
                default:
                    v15 := GenS2{F0: (!true), F1: 437}
                    _ = v15
                }
            } else {
                _ = len(v2)
                if (!(!false)) {
                    if false {
                        v14 = v14
                        v13 = v3
                    } else {
                        fmt.Println(v11, v3, v11)
                        switch "a" {
                        case "":
                            v16 := GenS2{F0: (!(42 >= 91)), F1: 725}
                            _ = v16
                        case "hello":
                            v17 := GenS3{F0: 127, F1: 1, F2: 65535}
                            _ = v17
                        }
                        if true {
                            v14 += -111
                            v18 := v1["a"]
                            _ = v18
                        }
                    }
                    switch v11 {
                    case "a":
                        delete(v4, 878)
                    }
                }
                var v19 int32 = (v5 * (v5 / ((v5) | 1)))
                _ = v19
            }
        } else {
            if (((((v11 + v11) + v11) + "hello") + ((v11 + (v11 + (((v11 + "a") + v11) + "foo bar"))) + v11)) != "hello") {
                v14 += (^(v14))
                v20 := GenS2{F0: ((false || (1 < -4)) || (true || (v3 < v13))), F1: 127}
                _ = v20
            }
            var v21 int32 = (^(v5))
            for fi22 := 0; fi22 < 16; fi22++ {
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 int64 = -615
                    var v25 int = 42
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                if true {
                    v5 = (v5 % ((v5) | 1))
                    v13 -= v3
                }
                v26 := GenS4{F0: v11}
                _ = fi22
                _ = v26
            }
            _ = v21
        }
        _ = v14
    } else {
        var v27 uint32 = 65535
        if (127 <= -91) {
            for fi28 := 0; fi28 < 18; fi28++ {
                v29 := v4[(fi28 - (fi28 | (fi28 ^ (fi28 >> ((42) & 63)))))]
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi28
                _ = v29
            }
        }
        _ = v27
    }
    var v30 string = v11
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v11 = "\n\t"
    var v31 bool = true
    v32 := GenS4{F0: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
