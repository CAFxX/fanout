package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint }

func main() {
    var v1 int = -991
    v2 := make(map[int]string)
    var v3 uint8 = 0
    var v4 uint32 = 2
    v5 := []int{v1, v1}
    _ = len(v2)
    for fi6 := 0; fi6 < 6; fi6++ {
        if (!(!false)) {
            v7 := GenS1{F0: 50, F1: v4}
            _ = v7
        }
        v8 := GenS1{F0: 1, F1: (v4 ^ 59)}
        var v9 string = ""
        _ = fi6
        _ = v8
        _ = v9
    }
    var v10 int64 = -9223372036854775808
    fmt.Println(v1)
    if true {
        v4 -= v4
        v2[v1] = (("" + "hello") + (("foo bar" + "hello") + "foo bar"))
        var v11 uint = 656
        _ = v11
    }
    v3 = v3
    switch "hello" {
    case "\n\t":
        v3 = 0
    case "":
        v12 := GenS1{F0: 1, F1: v4}
        _ = v12
    default:
        v13 := GenS1{F0: 1, F1: 42}
        _ = v13
    }
    if true {
        if true {
            var v14 int = v1
            if false {
                for fi15 := 0; fi15 < 3; fi15++ {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        fmt.Println(v14)
                        if ((^(v14)) == v14) {
                            v10 = (-(v10))
                            var v17 float64 = (0.0 * -2.25)
                            _ = len(v2)
                            if true {
                                delete(v2, v1)
                                if true {
                                    switch v1 {
                                    case 1:
                                        v17 -= float64(0.0)
                                    default:
                                        fmt.Println(v3, v1)
                                    }
                                    v18 := GenS1{F0: 548, F1: (v4 * v4)}
                                    switch v1 {
                                    case 128:
                                        v19 := GenS1{F0: 0, F1: (v4 % ((v4) | 1))}
                                        _ = v19
                                    case 4294967295:
                                        v20 := GenS2{F0: v3, F1: 100, F2: 42}
                                        _ = v20
                                    }
                                    fmt.Println(v17, v3, fi16)
                                    delete(v2, (fi16 << ((42) & 63)))
                                    _ = v18
                                }
                            } else {
                                if false {
                                    fmt.Println(v14, fi15)
                                    var v21 int16 = -32768
                                    _ = v21
                                } else {
                                    v14 = (-(fi15))
                                }
                            }
                            _ = v17
                        } else {
                            if (85 == 255) {
                                delete(v2, (v14 - (v1 << ((42) & 63))))
                            }
                        }
                        v14 = (fi15 ^ (v14 * v14))
                        _ = fi16
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi15
                }
            }
            var v22 uint32 = 452
            _ = v14
            _ = v22
        } else {
            switch ((((("a" + "\n\t") + (("\n\t" + "") + (("foo bar" + "") + "a"))) + "a") + "a") + "hello") {
            case "foo bar":
                v4 = v4
            case "":
                for fi23 := 0; fi23 < 7; fi23++ {
                    var v24 int32 = 0
                    v24 = v24
                    _ = fi23
                    _ = v24
                }
            default:
                var v25 int8 = -19
                _ = v25
            }
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(fi26, v4, v1)
                _ = fi26
            }
        }
    }
    var v27 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
