package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint; F1 string; F2 uint16 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint8; F2 int }
func genFn0(p0 int32, p1 int16) uint32 {
    v7 := GenS1{F0: ""}
    fmt.Println(p1, p0)
    v8 := []int{}
    v9 := GenS1{F0: "hello"}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int16 = 334
    var v2 int = 140
    var v3 float32 = -2.25
    var v4 int8 = -128
    var v5 int16 = v1
    v6 := []int{1, (v2 + 32767)}
    switch "foo bar" {
    case "":
        v6 = append(v6, (^(v2)))
    }
    if true {
        var v10 uint64 = 588
        _ = v10
    }
    if false {
        if (0 == 1) {
            if ((0 != genFn0(-222, v5)) || (-553 == 0)) {
                v11 := GenS1{F0: (("foo bar" + ("hello" + "")) + ("a" + ("\n\t" + "\n\t")))}
                _ = len(v6)
                for fi12 := 0; fi12 < 15; fi12++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = len(v6)
                    _ = fi12
                }
                _ = v11
            } else {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v6 = append(v6, fi13)
                    _ = fi13
                }
                v14 := GenS2{F0: false}
                _ = v14
            }
            v4 = v4
        }
        v6 = append(v6, (v2 - 9))
        var v15 uint32 = 670
        _ = v15
    }
    v5 = int16(v2)
    v1 = v1
    if (4294967295 >= 313) {
        var v16 int32 = 177
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v16
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch -61 {
        case 4294967295:
            var v17 uint32 = 42
            _ = v17
        case 32768:
            if (-9223372036854775808 > -787) {
                for fi18 := 0; fi18 < 13; fi18++ {
                    if (((v5 + (v5 & (v5 / ((v5) | 1)))) >= (v1 * -56)) && true) {
                        v19 := GenS1{F0: "a"}
                        var v20 uint = 4294967295
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(fi18, fi18, v4)
                    }
                    _ = fi18
                }
                v5 -= (v5 << ((255) & 63))
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        var v21 int = (-(v2))
        _ = v21
    case "":
        v22 := GenS3{F0: 9223372036854775807, F1: v4}
        _ = v22
    case "hello":
        v6 = append(v6, (v2 - v2))
    }
    if (v4 != 100) {
        v23 := GenS2{F0: (423 != 1)}
        _ = v23
    } else {
        var v24 uint16 = 32767
        v25 := GenS4{F0: 1, F1: "foo bar", F2: 702}
        if true {
            if (false && (1 > int32(v24))) {
                v1 -= v5
                if (!false) {
                    if (true || false) {
                        if (v5 > (-(v1))) {
                            var v26 uint32 = 208
                            _ = v26
                        }
                    }
                    var v27 float64 = ((1.5 * (3.14159 - -2.25)) + (3.14159 / 1e10))
                    if (2147483647 == 2147483647) {
                        if ((-(v5)) >= v1) {
                            fmt.Println(v3, v24)
                        }
                        v28 := GenS1{F0: "foo bar"}
                        _ = v28
                    }
                    _ = v27
                }
                v29 := GenS4{F0: 2, F1: "hello", F2: v24}
                _ = v29
            }
        }
        _ = v24
        _ = v25
    }
    v30 := GenS5{F0: v2, F1: 277}
    for fi31 := 0; fi31 < 20; fi31++ {
        v32 := GenS6{F0: 0, F1: 2, F2: -373}
        _ = fi31
        _ = v32
    }
    v1 = (v1 + (v1 / ((v5) | 1)))
    fmt.Println(v1)
    v6 = append(v6, (-(v2)))
    var v33 int16 = (^(v1))
    switch (v2 / ((v2) | 1)) {
    case -2147483648:
        for fi34 := 0; fi34 < 2; fi34++ {
            v35 := GenS3{F0: -1, F1: (v4 | v4)}
            _ = fi34
            _ = v35
        }
    case 1:
        var v36 bool = (("a" > "\n\t") && (!true))
        _ = v36
    case 2147483647:
        var v37 uint = 18446744073709551615
        _ = v37
    default:
        var v38 float64 = (float64(v3) / 1.0)
        _ = v38
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        for fi40 := 0; fi40 < 1; fi40++ {
            _ = len(v6)
            var v41 bool = false
            _ = len(v6)
            _ = fi40
            _ = v41
        }
        v42 := GenS4{F0: 4294967295, F1: "a", F2: 42}
        _ = fi39
        _ = v42
    }
    v43 := GenS6{F0: -1, F1: 127, F2: (v2 + (v2 & 128))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
