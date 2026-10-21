package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
func genFn0(p0 uint) string {
    v6 := []int{-979, int(p0)}
    _ = len(v6)
    _ = p0
    _ = v6
    return ""
}
func genFn1(p0 uint32) uint16 {
    var v7 int16 = 32767
    fmt.Println(v7)
    p0 = (^(p0))
    var v8 int64 = -1
    _ = p0
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 42
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 42
    v5 := make(map[string]string)
    for fi9 := 0; fi9 < 16; fi9++ {
        v2 += v2
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 uint8 = 255
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := []int{781}
    if true {
        v2 = (v2 & (-(v2)))
        v13 := GenS1{F0: true, F1: 1, F2: 42}
        var v14 int = 128
        _ = v13
        _ = v14
    }
    var v15 int8 = -112
    v1["foo bar"] = 2147483647
    switch -9223372036854775808 {
    case -1:
        var v16 float32 = 0.0
        _ = v16
    case 9223372036854775807:
        var v17 float64 = -2.25
        _ = v17
    }
    if len(v12) > 0 {
        v12[0] = -9223372036854775808
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            if true {
                v20 := GenS1{F0: (!false), F1: 986, F2: 2}
                fmt.Println(fi18, v15, v15)
                if ((-(v15)) >= (v15 | int8(v2))) {
                    fmt.Println(v15, fi19)
                    switch "" {
                    case "foo bar":
                        v15 = (-(v15))
                    case "hello":
                        if false {
                            v15 = (v15 >> ((0) & 63))
                            var v21 uint = 255
                            fmt.Println(v21)
                            _ = v21
                        }
                    case "\n\t":
                        var v22 int = (fi19 << ((0) & 63))
                        v23 := v5[("foo bar" + "foo bar")]
                        v24 := GenS1{F0: false, F1: 4294967295, F2: 67}
                        _ = v22
                        _ = v23
                        _ = v24
                    default:
                        var v25 uint16 = 2
                        _ = v25
                    }
                    v26 := GenS1{F0: true, F1: 42, F2: 80}
                    _ = v26
                }
                _ = v20
            } else {
                var v27 uint64 = 18446744073709551615
                _ = v27
            }
            v15 += (v15 - (v15 - v15))
            _ = fi19
        }
        v12 = append(v12, (fi18 & fi18))
        _ = fi18
    }
    var v28 int8 = (v15 * (v15 % (((^(v15))) | 1)))
    if false {
        v4 -= genFn1(1)
        if false {
            var v29 float32 = 0.0
            switch "foo bar" {
            case "\n\t":
                v15 += (v28 - v15)
            case "a":
                switch "hello" {
                case "\n\t":
                    fmt.Println(v15, v4, v28)
                case "a":
                    for fi30 := 0; fi30 < 7; fi30++ {
                        if true {
                            for fi31 := 0; fi31 < 5; fi31++ {
                                var v32 int = fi30
                                if ((!true) || false) {
                                    v33 := GenS2{F0: 924, F1: 65535, F2: (v4 % (((^(v4))) | 1))}
                                    _ = v33
                                }
                                _ = fi31
                                _ = v32
                            }
                        } else {
                            if false {
                                var v34 uint8 = 42
                                _ = v34
                            }
                        }
                        var v35 uint64 = 2
                        _ = fi30
                        _ = v35
                    }
                case "":
                    var v36 int16 = -32768
                    _ = v36
                }
            }
            _ = v29
        }
        switch int(v2) {
        case 4294967295:
            var v37 uint32 = 2147483647
            _ = v37
        default:
            v1["foo bar"] = 201
        }
    } else {
        fmt.Println(v15, v2)
        v28 += -97
        v4 += (^(v4))
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    if (!((v2 - v2) < (v2 - (^(v2))))) {
        var v38 int8 = 127
        for fi39 := 0; fi39 < 15; fi39++ {
            v4 += v4
            _ = fi39
        }
        v28 = v28
        _ = v38
    }
    _ = len(v5)
    fmt.Println(v28, v15, v2)
    var v40 float32 = 1.5
    var v41 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
