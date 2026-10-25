package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int }

func main() {
    v1 := []string{"\n\t", ("" + "\n\t"), ("foo bar" + (((("a" + "\n\t") + "\n\t") + ("" + "\n\t")) + ("foo bar" + "")))}
    var v2 uint16 = 42
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    var v5 int8 = -128
    fmt.Println(v5, v2)
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 12; fi7++ {
        _ = len(v1)
        fmt.Println(v2, v2)
        _ = fi7
    }
    v8 := GenS1{F0: v5}
    v9 := GenS1{F0: (v5 >> ((42) & 63))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int = (fi10 ^ fi10)
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 string = "a"
            switch fi12 {
            case 4294967295:
                if true {
                    if true {
                        v6 = (^(v6))
                        var v14 int32 = 42
                        _ = v14
                    } else {
                        var v15 bool = false
                        _ = v15
                    }
                    var v16 int = fi12
                    v17 := GenS2{F0: ("foo bar" + v13), F1: (^(fi12))}
                    v5 = 90
                    _ = v16
                    _ = v17
                } else {
                    var v18 int64 = 1
                    var v19 int64 = (v18 - (v18 % ((0) | 1)))
                    if (true && (2147483647 < 2147483647)) {
                        if true {
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = v13
                                }
                            }
                            var v20 int64 = v19
                            _ = v20
                        } else {
                            v1 = append(v1, ("" + ("hello" + (v13 + v13))))
                            if true {
                                var v21 uint8 = v6
                                _ = v21
                            }
                            if (1e10 != (1e10 - 0.1)) {
                                if false {
                                    v18 -= (v18 * (v19 ^ 1))
                                    var v22 string = ""
                                    _ = v22
                                }
                                var v23 uint64 = 2
                                _ = v23
                            } else {
                                v13 = ((("a" + v13) + (v13 + "a")) + (v13 + (v13 + "hello")))
                                var v24 float32 = -2.25
                                if true {
                                    var v25 float32 = v24
                                    v26 := v3[669]
                                    v27 := GenS1{F0: (-(v5))}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                _ = v24
                            }
                        }
                    }
                    _ = v18
                    _ = v19
                }
            default:
                if false {
                    v13 = v13
                    fmt.Println(v5, v11, v11)
                }
            }
            _ = fi12
            _ = v13
        }
        var v28 uint32 = 2
        _ = fi10
        _ = v11
        _ = v28
    }
    switch "\n\t" {
    case "\n\t":
        var v29 bool = true
        _ = v29
    }
    if false {
        fmt.Println(v6)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
