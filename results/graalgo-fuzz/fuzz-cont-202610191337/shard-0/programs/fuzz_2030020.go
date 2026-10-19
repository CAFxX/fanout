package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint64 }
func genFn0(p0 string, p1 uint16, p2 bool) uint64 {
    fmt.Println(p2, p0)
    v6 := GenS1{F0: 1, F1: 123, F2: 4294967295}
    var v7 uint64 = 598
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (v7 / ((v7) | 1))
}

func main() {
    v1 := []string{"foo bar", "a"}
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]int)
    var v5 int32 = -1
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 float64 = 1e10
        _ = fi8
        _ = v9
    }
    v5 -= -1
    v4[("foo bar" + "foo bar")] = 2147483647
    fmt.Println(v3)
    v5 = -1
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint32 = 42
        var v12 bool = true
        _ = fi10
        _ = v11
        _ = v12
    }
    v5 = (v5 << ((42) & 63))
    var v13 string = "foo bar"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 = v5
    if false {
        switch (v13 + "foo bar") {
        case "":
            var v14 uint32 = 2147483648
            _ = v14
        }
        v15 := v2[0]
        _ = v15
    } else {
        for fi16 := 0; fi16 < 17; fi16++ {
            switch (-(fi16)) {
            case 926:
                if (!false) {
                    var v17 string = v13
                    _ = v17
                }
            case -2147483648:
                var v18 uint64 = 145
                _ = v18
            case 1:
                switch "foo bar" {
                case "a":
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if ((true || false) || (true && (fi19 > fi16))) {
                            var v20 uint16 = 42
                            var v21 int32 = (v5 - v5)
                            v22 := v2[185]
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        _ = fi19
                    }
                case "hello":
                    var v23 int16 = -419
                    _ = v23
                case "":
                    for fi24 := 0; fi24 < 8; fi24++ {
                        if true {
                            delete(v2, 42)
                            v25 := v2[781]
                            _ = v25
                        } else {
                            if (32767 != 858) {
                                delete(v4, v13)
                                var v26 float32 = (0.0 * 3.14159)
                                switch -742 {
                                case -2147483648:
                                    var v27 float64 = (0.0 * (3.14159 + float64(v26)))
                                    _ = v27
                                case -1:
                                    var v28 int64 = v3
                                    _ = v28
                                case 127:
                                    if (false || (fi16 > (fi16 << ((255) & 63)))) {
                                        var v29 uint32 = 4294967295
                                        _ = v29
                                    }
                                }
                                v26 = ((((0.0 * 1.5) / v26) - v26) + float32(3.14159))
                                if false {
                                    _ = len(v2)
                                    switch fi16 {
                                    case 4294967295:
                                        _ = len(v1)
                                    }
                                    v4[(v13 + ((v13 + "") + "\n\t"))] = fi24
                                }
                                v13 = v13
                                fmt.Println(v3, v3)
                                _ = v26
                            }
                        }
                        fmt.Println(fi16)
                        switch fi16 {
                        case 0:
                            switch 1 {
                            case 713:
                                v3 -= v3
                            default:
                                v3 -= -9223372036854775808
                                v2[2] = 1
                                fmt.Println(fi16, fi16)
                            }
                        default:
                            var v30 uint = 736
                            _ = v30
                        }
                        v3 += int64(v5)
                        _ = fi24
                    }
                default:
                    var v31 float32 = -2.25
                    _ = v31
                }
            }
            _ = fi16
        }
        fmt.Println(v13, v3)
        for fi32 := 0; fi32 < 10; fi32++ {
            var v33 uint32 = 261
            v34 := GenS1{F0: 42, F1: v33, F2: genFn0(v13, 1, true)}
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    var v35 uint8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
