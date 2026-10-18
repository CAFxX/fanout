package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 string; F1 int; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 uint; F1 uint32 }
func genFn0(p0 int64) uint16 {
    var v7 int32 = 42
    v8 := []int{0, 42, -48, 32768}
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int16 = -1
    v2 := []uint32{}
    var v3 uint16 = 65535
    v4 := make(map[string]int)
    var v5 uint64 = 0
    v6 := make(map[string]int)
    fmt.Println(v3, v5, v5)
    var v9 int64 = 9223372036854775807
    if true {
        v5 = (^(v5))
    } else {
        fmt.Println(v1, v1)
        v5 += (v5 << ((0) & 63))
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        v4["a"] = (-(fi10))
        _ = len(v2)
        var v11 uint8 = 127
        _ = fi10
        _ = v11
    }
    v9 = v9
    if false {
        switch -1 {
        case 0:
            switch 688 {
            case 1:
                v1 = (v1 * (v1 * v1))
            case 0:
                switch "a" {
                case "foo bar":
                    fmt.Println(v1, v3)
                case "a":
                    if false {
                        for fi12 := 0; fi12 < 19; fi12++ {
                            switch "" {
                            case "foo bar":
                                v5 = v5
                            default:
                                for fi13 := 0; fi13 < 3; fi13++ {
                                    var v14 int = fi12
                                    var v15 int16 = (v1 * (^(v1)))
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            }
                            var v16 int16 = int16(fi12)
                            _ = fi12
                            _ = v16
                        }
                        var v17 uint64 = v5
                        _ = len(v6)
                        _ = v17
                    }
                }
            }
        case -2147483648:
            if len(v2) > 0 {
                v2[0] = 843
            }
        case 9223372036854775807:
            for fi18 := 0; fi18 < 11; fi18++ {
                v19 := GenS1{F0: (v3 * (v3 & v3)), F1: v9}
                var v20 int32 = 0
                var v21 uint64 = (v5 * v5)
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
        fmt.Println(v3)
    }
    switch (("\n\t" + ((("\n\t" + "\n\t") + "") + "hello")) + "foo bar") {
    case "":
        if true {
            for fi22 := 0; fi22 < 7; fi22++ {
                v23 := GenS1{F0: genFn0((v9 >> ((4294967295) & 63))), F1: v9}
                v24 := GenS2{F0: "", F1: (fi22 | fi22), F2: 65535}
                var v25 int8 = 1
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            v6["\n\t"] = 671
            v26 := GenS3{F0: 1, F1: -2147483648, F2: 947}
            _ = v26
        }
    case "hello":
        for fi27 := 0; fi27 < 17; fi27++ {
            for fi28 := 0; fi28 < 8; fi28++ {
                switch "" {
                case "foo bar":
                    v1 += -135
                case "\n\t":
                    var v29 float32 = float32(1e10)
                    _ = v29
                default:
                    v30 := GenS4{F0: 0}
                    _ = v30
                }
                v31 := GenS5{F0: -1, F1: (-(fi28))}
                if false {
                    v3 += 32767
                    _ = len(v4)
                } else {
                    if (2 <= 255) {
                        var v32 int = (fi27 % ((fi28) | 1))
                        _ = v32
                    }
                }
                _ = fi28
                _ = v31
            }
            v3 -= (^(v3))
            fmt.Println(fi27, v9)
            _ = fi27
        }
    case "a":
        v2 = append(v2, 4294967295)
    default:
        v33 := GenS4{F0: 18446744073709551615}
        _ = v33
    }
    v34 := v6[((("foo bar" + "\n\t") + "") + ("" + ("foo bar" + "\n\t")))]
    if false {
        for fi35 := 0; fi35 < 8; fi35++ {
            for fi36 := 0; fi36 < 12; fi36++ {
                switch "a" {
                case "hello":
                    v37 := v4[("a" + "hello")]
                    _ = v37
                case "":
                    v1 = v1
                default:
                    var v38 uint32 = 360
                    _ = v38
                }
                fmt.Println(v34)
                _ = fi36
            }
            for fi39 := 0; fi39 < 11; fi39++ {
                v40 := v6["\n\t"]
                _ = fi39
                _ = v40
            }
            v41 := GenS6{F0: 929, F1: 596}
            _ = fi35
            _ = v41
        }
    } else {
        var v42 uint16 = v3
        var v43 int16 = v1
        var v44 int8 = 0
        _ = v42
        _ = v43
        _ = v44
    }
    var v45 float64 = ((0.1 - -2.25) + float64(1e10))
    v5 -= (v5 | v5)
    v2 = append(v2, 4294967295)
    fmt.Println(v1)
    var v46 uint64 = v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v9
}
