package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 string; F1 uint8; F2 uint }
func genFn0() uint32 {
    v4 := []int{128, -9223372036854775808, -9223372036854775808, 460}
    for fi5 := 0; fi5 < 18; fi5++ {
        if (false && true) {
            var v6 string = "a"
            _ = v6
        } else {
            var v7 string = ""
            _ = v7
        }
        v8 := GenS1{F0: 9223372036854775807, F1: 42, F2: 0}
        var v9 int32 = 2147483647
        _ = fi5
        _ = v8
        _ = v9
    }
    _ = v4
    return 472
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    switch ("foo bar" + "") {
    case "a":
        switch "a" {
        case "foo bar":
            if (147 != 836) {
                switch 4294967295 {
                case -2147483648:
                    v10 := []int{}
                    _ = v10
                case -1:
                    var v11 uint8 = 50
                    _ = v11
                case 32768:
                    v12 := GenS1{F0: 1, F1: 2147483648, F2: 42}
                    _ = v12
                default:
                    v13 := GenS2{F0: "\n\t", F1: 231, F2: 345}
                    _ = v13
                }
            }
        default:
            var v14 int8 = 42
            _ = v14
        }
    case "foo bar":
        var v15 float64 = 1.5
        _ = v15
    default:
        switch 128 {
        case -294:
            if (0 > 578) {
                switch 9223372036854775807 {
                case -2147483648:
                    var v16 int64 = -9223372036854775808
                    _ = v16
                case 42:
                    v17 := []int{}
                    _ = v17
                case 32767:
                    switch 4294967295 {
                    case -175:
                        var v18 uint32 = 0
                        _ = v18
                    default:
                        switch 9223372036854775807 {
                        case 32767:
                            var v19 uint64 = 919
                            _ = v19
                        case -334:
                            if true {
                                var v20 int32 = -1
                                if true {
                                    var v21 uint32 = 2
                                    _ = v21
                                } else {
                                    fmt.Println(v20)
                                    for fi22 := 0; fi22 < 17; fi22++ {
                                        fmt.Println(v20)
                                        _ = fi22
                                    }
                                    var v23 int16 = 1
                                    _ = v23
                                }
                                v20 = v20
                                _ = v20
                            } else {
                                var v24 int16 = 963
                                var v25 uint64 = 18446744073709551615
                                _ = v24
                                _ = v25
                            }
                        }
                    }
                }
            } else {
                var v26 string = ""
                v27 := GenS1{F0: -873, F1: 1, F2: 0}
                fmt.Println(v26)
                _ = v26
                _ = v27
            }
        case 32767:
            v28 := GenS1{F0: 0, F1: 2147483648, F2: 65535}
            _ = v28
        case -2147483648:
            var v29 uint8 = 0
            _ = v29
        }
    }
    var v30 uint = 819
    v31 := GenS2{F0: "a", F1: 1, F2: (v30 ^ (^(v30)))}
    if true {
        fmt.Println(v30)
    }
    _ = len(v2)
    v32 := v2[943]
    var v33 uint32 = 1
    v32 = genFn0()
    if false {
        switch 204 {
        case 42:
            v34 := []int{}
            _ = v34
        case -1:
            v35 := []int{9223372036854775807}
            _ = v35
        }
    } else {
        v32 = 42
        v36 := GenS2{F0: "hello", F1: 2, F2: (v30 / ((v30) | 1))}
        v33 = (v33 % (((^(v32))) | 1))
        _ = v36
    }
    v37 := v2[v33]
    if true {
        fmt.Println(v30, v30, v32)
    } else {
        v38 := GenS2{F0: "foo bar", F1: uint8(v30), F2: (^(v30))}
        _ = v38
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        fmt.Println(v33)
        _ = fi39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v37
}
