package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int }
func genFn0(p0 int64) int32 {
    if (true || false) {
        v4 := []int{919, 4294967295}
        var v5 uint32 = 657
        _ = v4
        _ = v5
    } else {
        p0 += 42
        if false {
            switch 42 {
            case -1:
                switch -9223372036854775808 {
                case 4294967295:
                    p0 = (p0 + (p0 | (p0 >> ((452) & 63))))
                }
            default:
                if (true && true) {
                    if true {
                        var v6 int8 = -42
                        _ = v6
                    } else {
                        var v7 uint8 = 255
                        v8 := []int{2147483647, 129, -822}
                        for fi9 := 0; fi9 < 7; fi9++ {
                            v7 += 0
                            _ = fi9
                        }
                        _ = v7
                        _ = v8
                    }
                    p0 = p0
                    if true {
                        p0 += (^(p0))
                    }
                } else {
                    v10 := []int{}
                    p0 = (p0 << ((805) & 63))
                    if false {
                        for fi11 := 0; fi11 < 19; fi11++ {
                            for fi12 := 0; fi12 < 2; fi12++ {
                                if len(v10) > 0 {
                                    v10[0] = (^(fi12))
                                }
                                switch ("" + "foo bar") {
                                case "":
                                    var v13 uint16 = 656
                                    if (175 == 204) {
                                        var v14 int8 = 1
                                        if (2 != 0) {
                                            _ = len(v10)
                                        }
                                        var v15 uint = 65535
                                        _ = v14
                                        _ = v15
                                    }
                                    _ = v13
                                case "a":
                                    p0 = -501
                                case "foo bar":
                                    v16 := GenS1{F0: 279}
                                    _ = v16
                                default:
                                    switch ("\n\t" + ("a" + "foo bar")) {
                                    case "\n\t":
                                        v17 := GenS1{F0: 1}
                                        _ = v17
                                    case "":
                                        p0 += (p0 >> ((1) & 63))
                                    case "foo bar":
                                        var v18 uint32 = 2
                                        _ = v18
                                    }
                                }
                                _ = fi12
                            }
                            v19 := GenS1{F0: -781}
                            _ = fi11
                            _ = v19
                        }
                    } else {
                        p0 -= (p0 & (p0 - (p0 - (p0 & 538))))
                    }
                    _ = v10
                }
            }
            v20 := []int{9223372036854775807, 0, -2147483648}
            p0 = 0
            _ = v20
        }
        if (!((18446744073709551615 <= 142) && (true || (!false)))) {
            var v21 uint8 = 128
            _ = v21
        } else {
            p0 = (p0 | p0)
            v22 := GenS1{F0: 1}
            _ = v22
        }
    }
    v23 := []int{4294967295, 9223372036854775807, -717}
    fmt.Println(p0)
    p0 += -9223372036854775808
    _ = p0
    _ = v23
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{4294967295, 0, 0, 65535}
    v3 := []int{}
    if false {
        if (!(341 <= 762)) {
            delete(v1, 992)
            for fi24 := 0; fi24 < 10; fi24++ {
                for fi25 := 0; fi25 < 11; fi25++ {
                    v26 := GenS2{F0: 127, F1: 1, F2: 8}
                    _ = fi25
                    _ = v26
                }
                _ = len(v1)
                _ = fi24
            }
            for fi27 := 0; fi27 < 2; fi27++ {
                v28 := GenS1{F0: 42}
                _ = fi27
                _ = v28
            }
        } else {
            _ = len(v1)
            var v29 uint32 = 65535
            fmt.Println(v29)
            _ = v29
        }
        var v30 uint64 = 658
        fmt.Println(v30)
        _ = v30
    }
    var v31 int32 = 2147483647
    v31 = (v31 | v31)
    fmt.Println(v31)
    v3 = append(v3, 4294967295)
    v32 := GenS3{F0: 0}
    _ = len(v1)
    var v33 uint8 = 127
    v34 := GenS4{F0: 65535, F1: -1}
    v33 = (v33 >> ((1) & 63))
    var v35 uint16 = 542
    switch "" {
    case "":
        for fi36 := 0; fi36 < 15; fi36++ {
            var v37 int = 742
            _ = fi36
            _ = v37
        }
    case "a":
        for fi38 := 0; fi38 < 9; fi38++ {
            v39 := GenS2{F0: v33, F1: uint(v35), F2: 127}
            v33 += v33
            if true {
                for fi40 := 0; fi40 < 18; fi40++ {
                    _ = len(v2)
                    v35 += 65535
                    _ = fi40
                }
            }
            _ = fi38
            _ = v39
        }
    default:
        _ = len(v3)
    }
    v31 = genFn0(42)
    var v41 uint32 = 2
    v42 := GenS5{F0: 2147483647}
    v35 += (v35 & v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v41
    _ = v42
}
