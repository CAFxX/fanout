package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int32 }
func genFn0(p0 int32) uint16 {
    fmt.Println(p0)
    var v7 uint16 = 2
    p0 = 42
    _ = p0
    _ = v7
    return (v7 * v7)
}
func genFn1() int64 {
    var v8 int8 = 0
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := []int{}
        for fi11 := 0; fi11 < 2; fi11++ {
            switch (fi11 >> ((395) & 63)) {
            case 42:
                if (false && false) {
                    fmt.Println(fi11)
                }
            case 175:
                v12 := GenS1{F0: v8}
                _ = v12
            default:
                v13 := GenS2{F0: 32767, F1: genFn0(0), F2: 127}
                _ = v13
            }
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    v8 -= (v8 & (v8 << ((65535) & 63)))
    v14 := []int{4294967295, 4294967295, 4294967295}
    _ = v14
    _ = v8
    return 0
}

func main() {
    var v1 float32 = ((((0.1 / 1.0) / float32(-2.25)) - 1.5) - float32(0.0))
    var v2 uint8 = 127
    var v3 int64 = 1
    var v4 uint = 450
    var v5 uint = 65535
    v6 := make(map[string]uint32)
    switch 796 {
    case 1:
        var v15 uint32 = 0
        _ = v15
    case 32768:
        v3 -= v3
    default:
        var v16 int16 = -968
        _ = v16
    }
    if (1 != -9223372036854775808) {
        var v17 int8 = 1
        var v18 float64 = (3.14159 + 0.0)
        fmt.Println(v5, v18, v18)
        _ = v17
        _ = v18
    }
    v3 = v3
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 int16 = 791
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := GenS1{F0: 127}
            if true {
                v23 := v6[("" + "foo bar")]
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    var v24 uint = (v5 * v4)
    var v25 uint = 4294967295
    for fi26 := 0; fi26 < 7; fi26++ {
        for fi27 := 0; fi27 < 13; fi27++ {
            v28 := GenS3{F0: genFn0(42), F1: 42}
            fmt.Println(v4, v4)
            _ = fi27
            _ = v28
        }
        v29 := []int{fi26, (fi26 ^ 9223372036854775807), fi26, (fi26 >> ((205) & 63))}
        _ = fi26
        _ = v29
    }
    v30 := []int{-541, 32767}
    fmt.Println(v24, v1)
    for fi31 := 0; fi31 < 7; fi31++ {
        if true {
            v25 = (v25 & 144)
        }
        var v32 uint = (^(v5))
        _ = fi31
        _ = v32
    }
    var v33 int8 = -121
    if (!true) {
        var v34 bool = false
        for fi35 := 0; fi35 < 2; fi35++ {
            if v34 {
                v24 = 2
                if v34 {
                    switch 2147483647 {
                    case 9223372036854775807:
                        var v36 bool = v34
                        _ = v36
                    }
                    if true {
                        for fi37 := 0; fi37 < 17; fi37++ {
                            if len(v30) >= 2 {
                                v30 = v30[1:len(v30)]
                            }
                            v38 := GenS1{F0: (v33 / ((78) | 1))}
                            _ = fi37
                            _ = v38
                        }
                    } else {
                        v2 += (v2 >> (((^(v4))) & 63))
                    }
                }
            } else {
                for fi39 := 0; fi39 < 3; fi39++ {
                    var v40 uint64 = 18446744073709551615
                    fmt.Println(fi39, v1, fi39)
                    _ = fi39
                    _ = v40
                }
                v41 := GenS4{F0: 909, F1: 42}
                _ = v41
            }
            v42 := GenS1{F0: (v33 & (v33 + v33))}
            _ = fi35
            _ = v42
        }
        fmt.Println(v34, v5)
        _ = v34
    }
    for fi43 := 0; fi43 < 18; fi43++ {
        if (false && true) {
            v44 := GenS4{F0: 255, F1: 2147483647}
            for fi45 := 0; fi45 < 9; fi45++ {
                fmt.Println(v2, v33)
                _ = fi45
            }
            fmt.Println(v1, v1, v4)
            _ = v44
        }
        if len(v30) > 0 {
            v30[0] = (-(fi43))
        }
        var v46 int32 = 2147483647
        _ = fi43
        _ = v46
    }
    v2 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
