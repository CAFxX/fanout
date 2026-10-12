package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int16; F2 string }
type GenS3 struct { F0 string }
func genFn0(p0 uint, p1 int16, p2 uint32) int64 {
    var v7 uint8 = 127
    p0 = p0
    var v8 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 uint64 = 594
    var v2 int8 = 42
    v3 := make(map[string]int)
    var v4 int8 = 1
    var v5 uint16 = 65535
    var v6 int32 = 42
    var v9 int16 = -1
    fmt.Println(v4, v2, v9)
    v10 := []int{32768}
    v10 = append(v10, 32768)
    v6 += int32(v5)
    _ = len(v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: 0}
        var v13 int32 = v6
        v14 := v3[""]
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    switch "" {
    case "hello":
        if len(v10) > 0 {
            v10[0] = 128
        }
    case "":
        if (-1 <= 456) {
            v15 := GenS1{F0: 1}
            var v16 uint = 874
            if false {
                v1 += 1
                fmt.Println(v5)
            }
            _ = v15
            _ = v16
        }
    }
    if true {
        fmt.Println(v5, v4, v2)
        if true {
            for fi17 := 0; fi17 < 9; fi17++ {
                fmt.Println(v2, v1, v4)
                v2 = -128
                _ = fi17
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                switch (("a" + "foo bar") + (("a" + "foo bar") + "a")) {
                case "":
                    v19 := GenS1{F0: 42}
                    _ = v19
                case "hello":
                    fmt.Println(v1, v4, v2)
                default:
                    switch "foo bar" {
                    case "hello":
                        if true {
                            if (false || true) {
                                var v20 float64 = ((-2.25 - -2.25) - ((1.5 + 1.5) * 0.0))
                                for fi21 := 0; fi21 < 16; fi21++ {
                                    switch 9223372036854775807 {
                                    case 32768:
                                        var v22 uint32 = 42
                                        _ = v22
                                    }
                                    v23 := GenS2{F0: "", F1: -1, F2: ""}
                                    v24 := GenS2{F0: "hello", F1: (v9 % ((v9) | 1)), F2: "\n\t"}
                                    _ = fi21
                                    _ = v23
                                    _ = v24
                                }
                                _ = v20
                            }
                            v2 += v4
                        } else {
                            var v25 int8 = (^(v2))
                            _ = v25
                        }
                    case "foo bar":
                        if (float64(0.0) >= ((1e10 + 0.0) * 0.1)) {
                            v26 := GenS3{F0: ""}
                            var v27 uint64 = v1
                            _ = v26
                            _ = v27
                        } else {
                            v1 = (v1 & v1)
                            for fi28 := 0; fi28 < 20; fi28++ {
                                v5 -= (v5 ^ (^(v5)))
                                _ = fi28
                            }
                            v9 -= -32768
                        }
                    case "\n\t":
                        v6 = (-(v6))
                    default:
                        var v29 uint16 = 65535
                        _ = v29
                    }
                }
                v30 := v3[(("" + "foo bar") + "\n\t")]
                _ = fi18
                _ = v30
            }
            delete(v3, "hello")
        }
    } else {
        fmt.Println(v1)
    }
    var v31 uint64 = 2
    if true {
        v3[(("foo bar" + ("" + "foo bar")) + ("\n\t" + "a"))] = 2147483647
        for fi32 := 0; fi32 < 12; fi32++ {
            fmt.Println(v4, v5, v1)
            var v33 int8 = (v2 / (((v4 * v4)) | 1))
            _ = fi32
            _ = v33
        }
        if true {
            v10 = append(v10, 2147483647)
        } else {
            var v34 int16 = (v9 & v9)
            v34 = v34
            _ = v34
        }
    }
    v6 = (v6 << ((0) & 63))
    if (4294967295 == (^(v1))) {
        v6 -= v6
    } else {
        v4 += (v4 - (v2 << ((898) & 63)))
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 uint8 = 2
            _ = len(v10)
            v37 := GenS1{F0: 391}
            _ = fi35
            _ = v36
            _ = v37
        }
        v1 = (v1 + v31)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
