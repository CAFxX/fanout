package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 int }
type GenS3 struct { F0 int16; F1 bool; F2 int32 }
func genFn0(p0 int16) uint16 {
    var v6 int = 128
    var v7 uint64 = 4294967295
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: 127, F1: v7}
        _ = fi8
        _ = v9
    }
    v6 = -395
    _ = p0
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 int64) int16 {
    if (false || (((p0 % ((p0) | 1)) >= (p0 - p0)) && false)) {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int64 = p0
            if true {
                v11 += (v11 - (v11 / ((1) | 1)))
                v12 := GenS2{F0: 0, F1: (false && (genFn0(506) <= 42)), F2: (fi10 << ((42) & 63))}
                var v13 uint16 = 32767
                _ = v12
                _ = v13
            } else {
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 uint32 = 4294967295
                    fmt.Println(fi10, fi14)
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 16; fi16++ {
                    fmt.Println(v11, v11, p0)
                    v11 += (-(p0))
                    _ = fi16
                }
                if false {
                    v11 += (-(p0))
                }
            }
            _ = fi10
            _ = v11
        }
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 uint8 = 42
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    fmt.Println(p0)
    if false {
        p0 -= -9223372036854775808
        v20 := []int{-1, 912, -640, 128}
        _ = v20
    } else {
        var v21 int8 = -41
        v22 := GenS2{F0: (v21 << ((176) & 63)), F1: false, F2: int(p0)}
        switch ("foo bar" + "") {
        case "a":
            v21 = -1
        case "foo bar":
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 float32 = float32(3.14159)
                _ = fi23
                _ = v24
            }
        case "":
            v25 := GenS1{F0: (-(v21)), F1: 1}
            _ = v25
        default:
            var v26 uint8 = 127
            _ = v26
        }
        _ = v21
        _ = v22
    }
    p0 = p0
    _ = p0
    return 32767
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    v4 := []int{128, 683, 127, 127}
    v5 := make(map[int]uint32)
    for fi27 := 0; fi27 < 18; fi27++ {
        switch 2147483647 {
        case 9223372036854775807:
            v1 -= genFn0(1)
        case 2147483647:
            switch fi27 {
            case 127:
                for fi28 := 0; fi28 < 14; fi28++ {
                    if true {
                        v29 := GenS3{F0: 32767, F1: false, F2: 0}
                        v30 := GenS1{F0: 127, F1: 834}
                        var v31 int8 = 127
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    var v32 int16 = 1
                    _ = fi28
                    _ = v32
                }
            default:
                fmt.Println(v1, fi27)
            }
        default:
            delete(v5, (fi27 % (((-(fi27))) | 1)))
        }
        v33 := GenS1{F0: 100, F1: 0}
        _ = fi27
        _ = v33
    }
    var v34 uint32 = 4294967295
    v35 := v3[-32]
    fmt.Println(v35)
    fmt.Println(v35, v34)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v36 uint8 = 42
        _ = v36
    } else {
        v37 := GenS1{F0: -128, F1: 1}
        for fi38 := 0; fi38 < 9; fi38++ {
            if false {
                for fi39 := 0; fi39 < 20; fi39++ {
                    if false {
                        if (2147483647 < 416) {
                            v35 = (v35 + (v34 << ((1) & 63)))
                            var v40 uint16 = genFn0(1)
                            _ = v40
                        } else {
                            if false {
                                v41 := GenS2{F0: 0, F1: ((((!(!false)) || false) && (0 == 18446744073709551615)) && (1 < (v35 ^ (v35 + uint32(v1))))), F2: 771}
                                _ = v41
                            }
                        }
                        fmt.Println(v35)
                    }
                    switch "foo bar" {
                    case "\n\t":
                        delete(v5, int(v35))
                    }
                    var v42 int16 = 32767
                    _ = fi39
                    _ = v42
                }
                var v43 int32 = 42
                var v44 uint64 = 42
                _ = v43
                _ = v44
            }
            _ = fi38
        }
        switch -652 {
        case 32768:
            var v45 int16 = -1
            _ = v45
        }
        _ = v37
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        fmt.Println(fi46, v1, v34)
        var v47 uint64 = 18446744073709551615
        v2["a"] = 1
        _ = fi46
        _ = v47
    }
    if ((-2.25 + 0.0) != (1.5 + 1e10)) {
        v35 = v35
        for fi48 := 0; fi48 < 18; fi48++ {
            switch "a" {
            case "\n\t":
                fmt.Println(v35, v1)
            }
            var v49 string = ("" + ("foo bar" + "hello"))
            _ = fi48
            _ = v49
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
