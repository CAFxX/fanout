package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0() string {
    var v7 uint = 931
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint8 = 127
        _ = fi8
        _ = v9
    }
    _ = v7
    return "hello"
}

func main() {
    var v1 uint16 = 55
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 int64 = 0
    var v5 uint64 = 18446744073709551615
    v6 := []uint32{uint32(v1), 2147483647, 4294967295}
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int = fi10
        _ = len(v6)
        v12 := GenS1{F0: ("foo bar" + "hello"), F1: -128}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch 4294967295 {
    case 32767:
        _ = len(v3)
    }
    var v13 uint16 = (^(v1))
    v14 := GenS1{F0: "", F1: -1}
    v15 := GenS1{F0: ("a" + "hello"), F1: 127}
    for fi16 := 0; fi16 < 3; fi16++ {
        if (!true) {
            fmt.Println(v1, v4)
            switch "a" {
            case "a":
                if false {
                    var v17 int16 = 32767
                    var v18 int16 = v17
                    _ = v17
                    _ = v18
                } else {
                    if ((!(18446744073709551615 == uint(v1))) && (true && true)) {
                        var v19 bool = ((42 >= 221) && ((v5 / ((v5) | 1)) < (v5 + v5)))
                        _ = v19
                    } else {
                        var v20 int8 = -128
                        switch "hello" {
                        case "":
                            for fi21 := 0; fi21 < 16; fi21++ {
                                v22 := GenS1{F0: "foo bar", F1: v20}
                                v23 := v2[494]
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        case "foo bar":
                            v13 -= (v1 / (((v1 << ((2) & 63))) | 1))
                        default:
                            var v24 int32 = -2147483648
                            _ = v24
                        }
                        var v25 uint16 = v13
                        _ = v20
                        _ = v25
                    }
                }
            default:
                switch fi16 {
                case -1:
                    v26 := GenS2{F0: (v13 & (v1 >> ((255) & 63)))}
                    _ = v26
                case 128:
                    if false {
                        fmt.Println(v4)
                    } else {
                        var v27 int = (fi16 - -115)
                        for fi28 := 0; fi28 < 15; fi28++ {
                            var v29 bool = true
                            if ((fi28 | -854) <= v27) {
                                v4 += v4
                                _ = len(v3)
                                v6 = append(v6, 2)
                                fmt.Println(v27, fi16)
                            }
                            var v30 uint16 = 2
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        _ = v27
                    }
                case 9223372036854775807:
                    var v31 bool = ((1 > 42) || true)
                    _ = v31
                default:
                    v6 = append(v6, 643)
                }
            }
            var v32 uint8 = 42
            _ = v32
        } else {
            var v33 uint8 = 235
            if true {
                var v34 uint64 = (v5 << ((65535) & 63))
                var v35 int16 = 32767
                var v36 bool = ((v1 + (^(v1))) >= v13)
                _ = v34
                _ = v35
                _ = v36
            } else {
                v37 := GenS2{F0: (v1 ^ 255)}
                _ = v37
            }
            var v38 uint = 2
            _ = v33
            _ = v38
        }
        for fi39 := 0; fi39 < 16; fi39++ {
            var v40 int8 = 1
            _ = fi39
            _ = v40
        }
        _ = fi16
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        switch fi41 {
        case 128:
            v4 = (v4 / ((v4) | 1))
        default:
            switch "foo bar" {
            case "a":
                var v42 int32 = -2147483648
                _ = v42
            }
        }
        v4 = v4
        var v43 int32 = 42
        _ = fi41
        _ = v43
    }
    var v44 uint = 964
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
