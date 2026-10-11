package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool, p1 string) bool {
    var v4 uint64 = 42
    v5 := GenS1{F0: true, F1: 251, F2: 65535}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (766 != 2)
}

func main() {
    var v1 uint = 2
    v2 := []int{-382, 412, -9223372036854775808, int(v1)}
    v3 := make(map[int]bool)
    v2 = append(v2, 127)
    v6 := GenS2{F0: 255}
    delete(v3, 42)
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            if len(v2) > 0 {
                v2[0] = (fi8 ^ (fi7 | fi8))
            }
            var v9 int8 = 127
            _ = fi8
            _ = v9
        }
        v3[(fi7 / ((14) | 1))] = ((550 > -725) && (!false))
        v1 = v1
        _ = fi7
    }
    var v10 float64 = 0.1
    if false {
        fmt.Println(v10, v10)
        for fi11 := 0; fi11 < 16; fi11++ {
            if true {
                _ = len(v2)
                var v12 string = "foo bar"
                _ = v12
            } else {
                var v13 int8 = 1
                var v14 string = "hello"
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    } else {
        for fi15 := 0; fi15 < 7; fi15++ {
            if false {
                if ("" != "\n\t") {
                    fmt.Println(v10)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    v16 := GenS1{F0: (4294967295 <= 42), F1: 160, F2: 65535}
                    var v17 bool = true
                    for fi18 := 0; fi18 < 14; fi18++ {
                        delete(v3, (fi18 * fi18))
                        _ = fi18
                    }
                    _ = v16
                    _ = v17
                }
                _ = len(v2)
            }
            if true {
                v19 := GenS1{F0: (-1 == 32767), F1: 938, F2: 50}
                if false {
                    switch ((("\n\t" + "foo bar") + "a") + ("foo bar" + "")) {
                    case "":
                        v20 := GenS1{F0: (((("a" + "") + "hello") >= "\n\t") || true), F1: 1, F2: 4294967295}
                        _ = v20
                    }
                    var v21 int8 = 0
                    _ = v21
                }
                for fi22 := 0; fi22 < 6; fi22++ {
                    _ = len(v2)
                    fmt.Println(fi22, fi15, fi22)
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (v1 <= v1) {
                            if false {
                                v1 += 276
                                v10 -= v10
                                var v23 float32 = 3.14159
                                _ = v23
                            } else {
                                var v24 int32 = -2147483648
                                var v25 uint8 = 2
                                _ = len(v2)
                                if len(v2) > 0 {
                                    v2[0] = 9223372036854775807
                                }
                                _ = v24
                                _ = v25
                            }
                        } else {
                            var v26 uint8 = 127
                            var v27 float32 = 1e10
                            fmt.Println(v27, fi15, fi22)
                            v28 := GenS3{F0: 965}
                            var v29 uint32 = 2147483647
                            _ = v26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    } else {
                        v1 += 382
                        var v30 uint16 = 1
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v30
                    }
                    _ = fi22
                }
                _ = v19
            }
            var v31 uint = v1
            _ = fi15
            _ = v31
        }
        if true {
            _ = len(v2)
            if false {
                delete(v3, -9223372036854775808)
            }
        }
    }
    var v32 string = "\n\t"
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v6
}
