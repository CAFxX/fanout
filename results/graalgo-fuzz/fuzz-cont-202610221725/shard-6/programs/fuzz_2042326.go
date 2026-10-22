package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int64) int {
    var v4 uint = 255
    var v5 int8 = -17
    p0 = (^(p0))
    _ = p0
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    v1 := []int{}
    v2 := make(map[string]bool)
    var v3 int16 = 0
    v6 := GenS1{F0: 79}
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            var v8 int8 = 127
            _ = len(v1)
            for fi9 := 0; fi9 < 9; fi9++ {
                _ = len(v1)
                v8 = 1
                _ = fi9
            }
            _ = v8
        }
        if true {
            var v10 uint8 = 46
            _ = v10
        } else {
            var v11 uint16 = 80
            v11 = (v11 - v11)
            v3 += v3
            _ = v11
        }
        switch fi7 {
        case 584:
            v3 += (v3 + 42)
        case 32767:
            v1 = append(v1, fi7)
        case 987:
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS1{F0: 2}
                var v14 int32 = -2147483648
                var v15 uint16 = 0
                if ((-2147483648 != fi7) || (!false)) {
                    v15 = v15
                    delete(v2, ("foo bar" + "hello"))
                }
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        default:
            if true {
                delete(v2, (("\n\t" + ("\n\t" + ("" + ("foo bar" + "a")))) + ("" + "a")))
                if true {
                    fmt.Println(fi7, v3)
                    v16 := GenS2{F0: 930, F1: 0}
                    _ = v16
                }
                for fi17 := 0; fi17 < 18; fi17++ {
                    if false {
                        v3 = v3
                        v18 := v2["foo bar"]
                        _ = v18
                    } else {
                        var v19 int64 = 0
                        _ = v19
                    }
                    _ = fi17
                }
            }
        }
        _ = fi7
    }
    if false {
        _ = len(v2)
        v3 -= (v3 & -32768)
    } else {
        switch 0 {
        case 400:
            for fi20 := 0; fi20 < 1; fi20++ {
                fmt.Println(fi20, fi20)
                var v21 uint = 1
                _ = fi20
                _ = v21
            }
        }
        v3 = (v3 + (v3 - (v3 / (((v3 | (-(v3)))) | 1))))
    }
    fmt.Println(v3)
    var v22 int64 = 9223372036854775807
    v3 = v3
    v22 = 42
    var v23 uint = 65535
    var v24 uint = (v23 - (v23 >> (((^(v23))) & 63)))
    if false {
        var v25 string = "\n\t"
        for fi26 := 0; fi26 < 16; fi26++ {
            _ = len(v2)
            v3 -= v3
            if true {
                switch ("\n\t" + (v25 + v25)) {
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "\n\t":
                    for fi27 := 0; fi27 < 19; fi27++ {
                        switch v25 {
                        case "\n\t":
                            if true {
                                fmt.Println(v22, v24)
                            } else {
                                var v28 uint64 = 4294967295
                                _ = v28
                            }
                        case "a":
                            var v29 int8 = 41
                            _ = v29
                        case "":
                            var v30 uint64 = 4294967295
                            _ = v30
                        default:
                            if ((true || true) && true) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if true {
                                    if true {
                                        fmt.Println(v3)
                                        var v31 float64 = 0.1
                                        v23 = (v23 & (^(v24)))
                                        _ = v31
                                    }
                                }
                                var v32 string = "\n\t"
                                v1 = append(v1, fi27)
                                _ = v32
                            }
                        }
                        v33 := v2["hello"]
                        var v34 int16 = v3
                        _ = fi27
                        _ = v33
                        _ = v34
                    }
                default:
                    v23 = (v23 << ((v23) & 63))
                }
            }
            _ = fi26
        }
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v6
}
