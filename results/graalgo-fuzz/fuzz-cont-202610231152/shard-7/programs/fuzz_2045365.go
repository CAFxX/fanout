package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 int16, p1 int32, p2 string) uint8 {
    fmt.Println(p0, p1)
    v7 := GenS1{F0: -1, F1: p1, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}

func main() {
    var v1 int16 = 712
    var v2 string = ("" + "a")
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    v5 := []string{((("foo bar" + v2) + (v2 + (v2 + v2))) + v2), v2}
    var v6 int8 = 127
    var v8 bool = false
    if v8 {
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                switch (v2 + v2) {
                case "hello":
                    v2 = ""
                default:
                    var v11 uint8 = 232
                    _ = v11
                }
                v8 = ((-(v1)) != v1)
                if false {
                    v6 += (v6 << ((541) & 63))
                } else {
                    v1 += (v1 + (v1 - -32768))
                }
                _ = fi10
            }
            var v12 int32 = int32(v6)
            fmt.Println(v8)
            _ = fi9
            _ = v12
        }
        var v13 int16 = v1
        var v14 int8 = v6
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 4; fi15++ {
            for fi16 := 0; fi16 < 5; fi16++ {
                if (2 >= 65535) {
                    var v17 float64 = 1.5
                    if ((255 < 255) || v8) {
                        v18 := GenS1{F0: 42, F1: -66, F2: 768}
                        _ = v18
                    } else {
                        var v19 uint = 1
                        if (!(728 >= 18446744073709551615)) {
                            if true {
                                var v20 float64 = (((((v17 * ((float64(3.14159) / 3.14159) - 1.5)) * v17) + (float64(0.0) - 3.14159)) / (v17 * float64(-2.25))) * (v17 * 0.0))
                                _ = v20
                            }
                        } else {
                            if v8 {
                                fmt.Println(v6, v19, fi16)
                                var v21 bool = (((v2 + v2) + (v2 + v2)) != ("foo bar" + (v2 + "hello")))
                                v19 = v19
                                _ = v21
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v19 += v19
                        }
                        _ = v19
                    }
                    v17 = float64(1.5)
                    _ = v17
                } else {
                    var v22 int = (fi15 * 2147483647)
                    if true {
                        if false {
                            if v8 {
                                v3[(fi16 >> ((42) & 63))] = 2147483647
                            } else {
                                fmt.Println(fi16)
                                fmt.Println(v2)
                                if true {
                                    v23 := GenS2{F0: 9223372036854775807, F1: 2147483648}
                                    _ = v23
                                }
                            }
                            var v24 uint = 18446744073709551615
                            _ = v24
                        }
                        v25 := v4[v2]
                        _ = v25
                    } else {
                        var v26 uint32 = 1
                        v26 -= (v26 >> ((1) & 63))
                        _ = v26
                    }
                    v27 := GenS2{F0: 1, F1: 0}
                    var v28 int64 = 501
                    _ = v22
                    _ = v27
                    _ = v28
                }
                _ = fi16
            }
            if true {
                var v29 uint8 = 74
                v8 = true
                _ = len(v5)
                _ = v29
            }
            v2 = (v2 + v2)
            _ = fi15
        }
        v30 := GenS1{F0: 1, F1: -1, F2: 41}
        _ = v30
    }
    var v31 float32 = 1e10
    var v32 float32 = 0.1
    v33 := GenS1{F0: -32768, F1: 2147483647, F2: 794}
    var v34 string = v2
    for fi35 := 0; fi35 < 20; fi35++ {
        var v36 uint32 = 1
        _ = fi35
        _ = v36
    }
    v1 = v1
    switch 0 {
    case 1:
        if (0 >= 18446744073709551615) {
            switch "hello" {
            case "\n\t":
                for fi37 := 0; fi37 < 3; fi37++ {
                    for fi38 := 0; fi38 < 19; fi38++ {
                        v8 = (true && v8)
                        v39 := GenS2{F0: 107, F1: 2}
                        v40 := v4[(v2 + (v34 + ""))]
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    v34 = ("foo bar" + "hello")
                    _ = fi37
                }
            case "":
                var v41 float32 = (0.1 + ((((((v31 + v31) - -2.25) + (1.5 + v31)) - 1e10) / (((v31 * 3.14159) - (1e10 / -2.25)) * 3.14159)) * v31))
                _ = v41
            default:
                var v42 float32 = ((v32 / float32(1.5)) * v32)
                _ = v42
            }
        } else {
            var v43 int8 = v6
            _ = v43
        }
    }
    if false {
        var v44 int = -198
        _ = v44
    }
    var v45 int32 = -1
    fmt.Println(v6, v2, v1)
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v8
}
