package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint64) int {
    v7 := GenS1{F0: 1}
    var v8 float32 = ((((1e10 / 0.1) * (float32((0.1 + 0.1)) + -2.25)) + 0.1) * 1e10)
    p0 = (p0 & (p0 << ((42) & 63)))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32768
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    var v3 int64 = -1
    v4 := []uint32{}
    var v5 uint32 = 0
    v6 := []string{"foo bar", "hello", ("a" + (("foo bar" + "") + "hello"))}
    v9 := GenS2{F0: "", F1: (^(v5)), F2: uint64(v5)}
    v5 -= 113
    if false {
        v3 = v3
    } else {
        switch ((("\n\t" + "hello") + ("\n\t" + "")) + "\n\t") {
        case "hello":
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 bool = false
                v11 = true
                if false {
                    v11 = v11
                    var v12 int16 = -853
                    if false {
                        v12 += (v12 / ((v12) | 1))
                    } else {
                        var v13 float64 = -2.25
                        _ = len(v2)
                        for fi14 := 0; fi14 < 20; fi14++ {
                            switch fi10 {
                            case 2147483647:
                                v15 := GenS2{F0: "\n\t", F1: (^(v5)), F2: 1}
                                _ = v15
                            default:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            v16 := v2[(fi10 * fi10)]
                            switch 9223372036854775807 {
                            case 0:
                                var v17 int16 = v12
                                _ = v17
                            case -351:
                                if v11 {
                                    v18 := v1[""]
                                    delete(v2, fi10)
                                    _ = v18
                                } else {
                                    if false {
                                        var v19 float64 = v13
                                        switch 42 {
                                        case 2147483647:
                                            var v20 uint64 = 4294967295
                                            _ = v20
                                        case 4294967295:
                                            v4 = append(v4, 65535)
                                        }
                                        v21 := GenS2{F0: ((((("foo bar" + "hello") + "hello") + "foo bar") + "\n\t") + "\n\t"), F1: (^(v5)), F2: 1}
                                        _ = v19
                                        _ = v21
                                    }
                                    v11 = (0 == 65535)
                                }
                            case -28:
                                if v11 {
                                    v16 += v16
                                    if v11 {
                                        var v22 int64 = v3
                                        _ = v22
                                    } else {
                                        if (v11 || v11) {
                                            fmt.Println(v11, v16)
                                        }
                                    }
                                } else {
                                    var v23 uint64 = 125
                                    _ = v23
                                }
                            default:
                                var v24 int32 = 0
                                _ = v24
                            }
                            _ = fi14
                            _ = v16
                        }
                        _ = v13
                    }
                    _ = v12
                } else {
                    v25 := GenS2{F0: "foo bar", F1: (v5 & (v5 - v5)), F2: 42}
                    var v26 string = "\n\t"
                    _ = v25
                    _ = v26
                }
                _ = fi10
                _ = v11
            }
        case "\n\t":
            fmt.Println(v3)
        case "foo bar":
            v1["foo bar"] = ""
        default:
            v3 -= v3
        }
        var v27 string = "a"
        if true {
            v28 := GenS3{F0: 42}
            _ = v28
        }
        _ = v27
    }
    v5 -= v5
    if true {
        fmt.Println(v3)
    }
    v1["\n\t"] = "foo bar"
    if (!((4294967295 > 0) || (!(!(2 > uint16(v5)))))) {
        v3 -= 945
    }
    if true {
        if false {
            v5 = (^(v5))
            var v29 string = ""
            v30 := GenS1{F0: 1}
            _ = v29
            _ = v30
        } else {
            v2[-226] = 1
        }
        v5 -= (^(v5))
    }
    var v31 int16 = -32768
    v5 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
