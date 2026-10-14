package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int32; F2 int16 }
type GenS4 struct { F0 bool }
func genFn0() int64 {
    switch "foo bar" {
    case "hello":
        var v7 int = -396
        _ = v7
    case "foo bar":
        var v8 uint64 = 892
        _ = v8
    case "\n\t":
        var v9 int = -501
        _ = v9
    }
    var v10 int8 = -24
    v10 = v10
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 int16 = -32768
        _ = fi11
        _ = v12
    }
    var v13 float64 = 3.14159
    _ = v10
    _ = v13
    return 0
}
func genFn1(p0 int16, p1 int32, p2 uint64) int8 {
    p1 = 2147483647
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 string = "a"
        _ = fi14
        _ = v15
    }
    var v16 float32 = 1.5
    switch 127 {
    case 128:
        var v17 int32 = (p1 >> ((uint(p2)) & 63))
        _ = v17
    }
    var v18 int32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v18
    return 1
}

func main() {
    v1 := []int{32767, 42}
    var v2 int8 = 1
    var v3 uint8 = 108
    var v4 float32 = 1.5
    var v5 string = ""
    var v6 uint64 = 0
    v3 -= v3
    v6 = 765
    for fi19 := 0; fi19 < 9; fi19++ {
        v4 -= (v4 - v4)
        _ = fi19
    }
    v20 := GenS1{F0: 65535, F1: (v2 % ((82) | 1))}
    switch -392 {
    case 128:
        if true {
            if len(v1) > 0 {
                v1[0] = -9223372036854775808
            }
            if (!(v6 != (v6 ^ v6))) {
                v6 += (v6 | (v6 ^ v6))
                for fi21 := 0; fi21 < 11; fi21++ {
                    var v22 int8 = 127
                    for fi23 := 0; fi23 < 20; fi23++ {
                        _ = len(v1)
                        var v24 uint8 = (v3 << ((808) & 63))
                        var v25 float32 = v4
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi21
                    _ = v22
                }
            }
        }
    }
    switch -734 {
    case -579:
        for fi26 := 0; fi26 < 19; fi26++ {
            v2 -= v2
            _ = fi26
        }
    case 0:
        v27 := GenS1{F0: 255, F1: (v2 * (v2 - genFn1(42, -1, v6)))}
        _ = v27
    default:
        switch 929 {
        case 9223372036854775807:
            if true {
                for fi28 := 0; fi28 < 9; fi28++ {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        v6 += (^(v6))
                        v30 := GenS1{F0: uint16(v3), F1: (^(v2))}
                        var v31 int32 = 2147483647
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    v3 = (v3 | uint8(v6))
                    _ = fi28
                }
            } else {
                var v32 float64 = -2.25
                if true {
                    v33 := GenS2{F0: int64(v2)}
                    _ = v33
                } else {
                    v34 := GenS2{F0: 42}
                    switch 4294967295 {
                    case 96:
                        var v35 float32 = -2.25
                        _ = v35
                    case 0:
                        var v36 uint = 18446744073709551615
                        _ = v36
                    }
                    _ = v34
                }
                _ = v32
            }
        default:
            for fi37 := 0; fi37 < 3; fi37++ {
                var v38 float32 = 0.1
                _ = fi37
                _ = v38
            }
        }
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 int64 = 42
        _ = fi39
        _ = v40
    }
    v4 = v4
    var v41 uint = 2
    if ((((1e10 / (1e10 / (v4 / float32(0.1)))) * v4) - (v4 - (-2.25 + float32(((-2.25 + 1e10) / 1e10))))) >= 0.0) {
        if len(v1) > 0 {
            v1[0] = -1
        }
    }
    v42 := GenS3{F0: (v5 + (v5 + v5)), F1: 2147483647, F2: -1}
    v43 := GenS4{F0: true}
    v4 += v4
    var v44 int16 = int16(v41)
    v5 = ("\n\t" + (v5 + ("hello" + ((((v5 + ((v5 + (((v5 + v5) + (v5 + (v5 + v5))) + ("foo bar" + v5))) + v5)) + "") + (("hello" + ((v5 + ("" + v5)) + v5)) + v5)) + (v5 + (v5 + ((v5 + (v5 + "foo bar")) + (v5 + v5))))))))
    if ((^(v44)) == v44) {
        if true {
            var v45 uint64 = 4294967295
            var v46 int = 207
            v47 := GenS2{F0: genFn0()}
            _ = v45
            _ = v46
            _ = v47
        }
        fmt.Println(v3, v4)
        var v48 int64 = -767
        _ = v48
    }
    v44 -= v44
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
