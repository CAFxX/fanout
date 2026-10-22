package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int16) int64 {
    v6 := []int{}
    v7 := GenS1{F0: 65535, F1: 2, F2: 32767}
    switch "\n\t" {
    case "":
        v8 := GenS1{F0: 65535, F1: 1, F2: -1}
        _ = v8
    default:
        var v9 uint8 = 1
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 331
}
func genFn1(p0 int64, p1 int32, p2 uint8) int8 {
    fmt.Println(p2)
    var v10 int64 = genFn0(32767)
    if false {
        var v11 int64 = v10
        var v12 float64 = (0.0 * -2.25)
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 9; fi13++ {
            p1 -= p1
            fmt.Println(fi13, p2, p2)
            var v14 int64 = (^(p0))
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS1{F0: 42, F1: 1, F2: -523}
        v17 := []int{fi15}
        var v18 int32 = (p1 % ((p1) | 1))
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return int8(p0)
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]int)
    var v3 float64 = ((1.5 + 0.1) * 3.14159)
    v4 := []int{2147483647}
    v5 := []int{}
    for fi19 := 0; fi19 < 12; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 int64 = 1
            var v22 string = "a"
            if (2147483647 == -1) {
                if false {
                    v21 += (v21 % ((v21) | 1))
                    fmt.Println(v1)
                    if (42 <= 42) {
                        v21 = (v21 << ((189) & 63))
                    } else {
                        var v23 int32 = 1
                        _ = v23
                    }
                } else {
                    v24 := GenS1{F0: 2, F1: 18446744073709551615, F2: (fi20 % ((fi20) | 1))}
                    v1 = 116
                    _ = v24
                }
                v22 = (v22 + v22)
                v1 = (v1 * v1)
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(v3, fi19, v1)
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    if false {
        v3 += -2.25
        switch (("hello" + "a") + "") {
        case "hello":
            v25 := v2["a"]
            _ = v25
        case "foo bar":
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 float32 = 0.0
                v28 := GenS1{F0: 255, F1: 2, F2: (fi26 - 0)}
                v29 := GenS2{F0: 32767}
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
        default:
            if (false && (-2.25 != -2.25)) {
                v30 := GenS3{F0: -457, F1: (v3 != (float64(3.14159) / 1.5))}
                if false {
                    v3 += v3
                    var v31 uint32 = 42
                    _ = v31
                }
                _ = v30
            }
        }
        if true {
            if (false || false) {
                v2["a"] = 9223372036854775807
            }
            v32 := GenS4{F0: 0}
            v3 -= (0.0 + 1.5)
            _ = v32
        } else {
            for fi33 := 0; fi33 < 18; fi33++ {
                v1 -= v1
                _ = fi33
            }
            for fi34 := 0; fi34 < 5; fi34++ {
                var v35 float64 = v3
                _ = fi34
                _ = v35
            }
            var v36 int64 = 0
            _ = v36
        }
    } else {
        v37 := v2[(("\n\t" + ("foo bar" + ("" + "hello"))) + "")]
        v37 = v37
        var v38 int32 = -2147483648
        _ = v37
        _ = v38
    }
    v3 = v3
    for fi39 := 0; fi39 < 15; fi39++ {
        var v40 int16 = 32767
        if len(v4) > 0 {
            v4[0] = fi39
        }
        _ = fi39
        _ = v40
    }
    v3 -= v3
    for fi41 := 0; fi41 < 7; fi41++ {
        v3 = v3
        for fi42 := 0; fi42 < 15; fi42++ {
            var v43 uint16 = 255
            _ = fi42
            _ = v43
        }
        v3 = float64(1.5)
        _ = fi41
    }
    var v44 uint = 694
    v45 := GenS5{F0: 653}
    delete(v2, "hello")
    var v46 uint64 = 4294967295
    var v47 uint64 = (^(v46))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v48 := GenS2{F0: 32767}
    if true {
        v46 = v46
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        v50 := GenS3{F0: (fi49 + fi49), F1: true}
        for fi51 := 0; fi51 < 2; fi51++ {
            v1 = 67
            var v52 uint32 = 2147483647
            _ = fi51
            _ = v52
        }
        v47 = 42
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}
