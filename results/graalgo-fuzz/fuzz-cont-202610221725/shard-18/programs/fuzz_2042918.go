package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 string; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint }
func genFn0() int {
    var v7 uint64 = 42
    v8 := GenS1{F0: -50, F1: 0}
    v9 := GenS1{F0: -1, F1: 9223372036854775807}
    v10 := []int{-9223372036854775808, -1, -2147483648}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int64 = 0
    var v2 int32 = 1
    var v3 float32 = 1e10
    var v4 uint64 = 1
    var v5 uint16 = 42
    var v6 int64 = v1
    v11 := []int{32768, 24}
    var v12 int32 = (v2 & 2147483647)
    _ = len(v11)
    v13 := GenS1{F0: -1, F1: 42}
    v14 := GenS2{F0: 42, F1: -32768}
    var v15 uint8 = 2
    _ = len(v11)
    var v16 uint = 684
    var v17 int = -9223372036854775808
    for fi18 := 0; fi18 < 13; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            if true {
                if (true && (uint16(v15) != 32767)) {
                    if true {
                        if true {
                            if len(v11) > 0 {
                                v11[0] = genFn0()
                            }
                            var v20 uint32 = 42
                            v3 -= (((v3 + ((3.14159 / v3) * (3.14159 + v3))) * v3) - (v3 / v3))
                            v21 := GenS2{F0: uint64(v16), F1: -1}
                            var v22 uint16 = (^(v5))
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            if false {
                                switch "a" {
                                case "a":
                                    _ = len(v11)
                                    _ = len(v11)
                                    var v23 bool = (42 <= v12)
                                    var v24 uint = v16
                                    switch ("" + "a") {
                                    case "foo bar":
                                        v11 = append(v11, v17)
                                    case "hello":
                                        v25 := GenS1{F0: 122, F1: (fi18 & int(v6))}
                                        _ = v25
                                    default:
                                        var v26 float64 = 1e10
                                        _ = v26
                                    }
                                    _ = v23
                                    _ = v24
                                }
                                v12 -= (v12 * v2)
                                var v27 uint = v16
                                var v28 int64 = v6
                                v5 += 65
                                _ = v27
                                _ = v28
                            }
                            fmt.Println(v3, v16, v17)
                        }
                        if len(v11) > 0 {
                            v11[0] = (fi18 >> (((v16 << ((0) & 63))) & 63))
                        }
                        var v29 uint64 = v4
                        _ = v29
                    }
                } else {
                    fmt.Println(v12, v6)
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        v11 = append(v11, 2147483647)
                    case "\n\t":
                        v4 -= 2
                    default:
                        var v30 int64 = v6
                        _ = v30
                    }
                    switch "foo bar" {
                    case "\n\t":
                        v31 := GenS1{F0: 0, F1: fi19}
                        _ = v31
                    case "a":
                        switch v17 {
                        case -604:
                            var v32 float32 = (v3 + (v3 - 1.5))
                            _ = v32
                        default:
                            v33 := GenS3{F0: "\n\t", F1: (v6 / (((v6 << ((v16) & 63))) | 1)), F2: 2}
                            fmt.Println(v12, v12)
                            if false {
                                v34 := GenS2{F0: (v4 ^ (v4 & v4)), F1: -567}
                                var v35 int = (fi18 % ((v17) | 1))
                                _ = v34
                                _ = v35
                            } else {
                                fmt.Println(fi19, fi18)
                                v1 = -9223372036854775808
                                var v36 uint32 = 4294967295
                                _ = v36
                            }
                            _ = v33
                        }
                        var v37 int8 = -128
                        var v38 bool = true
                        _ = v37
                        _ = v38
                    }
                }
                v39 := GenS3{F0: ("hello" + "a"), F1: v1, F2: 32767}
                _ = v39
            }
            var v40 uint8 = 0
            v41 := GenS1{F0: 127, F1: (fi18 / (((fi18 + (-(fi19)))) | 1))}
            _ = fi19
            _ = v40
            _ = v41
        }
        fmt.Println(v16, v15)
        switch genFn0() {
        case -121:
            for fi42 := 0; fi42 < 2; fi42++ {
                var v43 int8 = 1
                _ = fi42
                _ = v43
            }
        default:
            var v44 uint = 688
            _ = v44
        }
        _ = fi18
    }
    for fi45 := 0; fi45 < 11; fi45++ {
        v46 := GenS4{F0: v16}
        _ = fi45
        _ = v46
    }
    for fi47 := 0; fi47 < 3; fi47++ {
        v12 = v12
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
