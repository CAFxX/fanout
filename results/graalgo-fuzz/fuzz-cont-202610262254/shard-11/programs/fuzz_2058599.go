package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    var v3 bool = (-1 >= 1000)
    v4 := make(map[uint32]uint32)
    var v5 int32 = 2147483647
    v6 := []int{594}
    v3 = v3
    var v7 uint32 = 863
    switch int(v5) {
    case -1:
        v5 -= v5
    default:
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int64 = -9223372036854775808
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int64 = v9
                var v12 uint64 = 632
                var v13 string = "hello"
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            fmt.Println(v9, v7)
            _ = fi8
            _ = v9
        }
    }
    var v14 int16 = 42
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 int32 = (v5 % ((-254) | 1))
            _ = fi16
            _ = v17
        }
        v7 -= v7
        if true {
            v7 -= v7
        }
        _ = fi15
    }
    if v3 {
        var v18 uint8 = 0
        _ = len(v4)
        _ = v18
    }
    v1[(v7 / ((0) | 1))] = v7
    if false {
        v19 := GenS1{F0: 2, F1: 148, F2: 420}
        v20 := GenS1{F0: 127, F1: 255, F2: 18446744073709551615}
        v14 += (v14 - (-(v14)))
        _ = v19
        _ = v20
    }
    v5 = v5
    v21 := GenS1{F0: 255, F1: 248, F2: 362}
    fmt.Println(v14, v5, v14)
    v22 := GenS1{F0: 0, F1: 0, F2: 65535}
    if (191 > 854) {
        v14 += v14
        var v23 int8 = 127
        _ = v23
    } else {
        for fi24 := 0; fi24 < 18; fi24++ {
            if (!v3) {
                v25 := GenS1{F0: 0, F1: 2, F2: uint(v7)}
                var v26 string = "\n\t"
                v3 = (!(255 <= 1))
                _ = v25
                _ = v26
            }
            _ = fi24
        }
        switch 4294967295 {
        case -1:
            for fi27 := 0; fi27 < 2; fi27++ {
                switch 1 {
                case 9223372036854775807:
                    v7 = 42
                case 32767:
                    v28 := GenS1{F0: 2, F1: 255, F2: 65535}
                    _ = v28
                default:
                    v2["foo bar"] = "a"
                }
                _ = fi27
            }
        case 9223372036854775807:
            if v3 {
                if true {
                    fmt.Println(v3)
                }
                switch "hello" {
                case "a":
                    v29 := GenS2{F0: v7}
                    _ = v29
                case "":
                    v30 := GenS2{F0: 378}
                    _ = v30
                case "hello":
                    for fi31 := 0; fi31 < 1; fi31++ {
                        var v32 int = fi31
                        for fi33 := 0; fi33 < 15; fi33++ {
                            switch (fi31 ^ v32) {
                            case 4294967295:
                                v7 += (v7 / ((v7) | 1))
                            default:
                                v5 = (v5 | v5)
                            }
                            if (v3 || (2 != 740)) {
                                v3 = ("hello" >= "")
                                v14 -= -32768
                                v3 = v3
                            } else {
                                var v34 uint8 = 127
                                v5 = (v5 & (v5 / ((42) | 1)))
                                _ = v34
                            }
                            _ = fi33
                        }
                        var v35 uint = 424
                        _ = fi31
                        _ = v32
                        _ = v35
                    }
                }
                fmt.Println(v3, v3)
            } else {
                v7 = 4294967295
                v36 := GenS2{F0: v7}
                _ = v36
            }
        default:
            v1[v7] = (v7 | v7)
        }
    }
    var v37 uint16 = 65535
    if (false || (!(v3 && false))) {
        if len(v6) > 0 {
            v6[0] = 424
        }
        var v38 int32 = v5
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
