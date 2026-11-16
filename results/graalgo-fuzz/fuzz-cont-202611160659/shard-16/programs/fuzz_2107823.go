package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint; F1 bool; F2 bool }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 string }
func genFn0(p0 uint32) uint8 {
    p0 += p0
    v4 := []int{-2147483648, 2147483647, -2147483648, 32767}
    _ = p0
    _ = v4
    return 120
}
func genFn1(p0 int8, p1 int8, p2 int32) bool {
    if (false || true) {
        var v5 uint16 = 0
        _ = v5
    }
    if false {
        var v6 float64 = 1.5
        if (!false) {
            v7 := []int{}
            p0 += (p1 | (p0 % (((p1 >> ((65535) & 63))) | 1)))
            _ = v7
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(v6, v6)
                p0 = p1
                _ = fi8
            }
            v9 := GenS1{F0: p0, F1: 18446744073709551615, F2: 669}
            _ = v9
        }
        switch "\n\t" {
        case "foo bar":
            if true {
                v10 := GenS2{F0: (!true), F1: int16(p1)}
                _ = v10
            } else {
                v11 := GenS3{F0: 128, F1: -512, F2: 74}
                for fi12 := 0; fi12 < 20; fi12++ {
                    v13 := GenS3{F0: (fi12 / ((fi12) | 1)), F1: 42, F2: 1}
                    fmt.Println(p0)
                    v14 := GenS2{F0: false, F1: -637}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = v11
            }
        }
        _ = v6
    }
    p1 -= p0
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return (1 == 9223372036854775807)
}

func main() {
    var v1 string = "\n\t"
    v2 := []int{}
    var v3 float32 = (((1e10 - (-2.25 + 1e10)) / (0.1 + ((1e10 + 1e10) * 1.5))) / 3.14159)
    switch 445 {
    case 42:
        if false {
            if (42 >= 127) {
                var v15 int64 = 0
                v2 = append(v2, 0)
                _ = v15
            }
        } else {
            fmt.Println(v3)
            var v16 bool = (!(18446744073709551615 == 1))
            if v16 {
                var v17 float32 = v3
                fmt.Println(v1, v3, v16)
                _ = v17
            } else {
                if true {
                    v3 += v3
                    v18 := GenS4{F0: 18446744073709551615, F1: v16, F2: ((42 >= 42) || true)}
                    _ = v18
                }
            }
            _ = v16
        }
    case 4294967295:
        if true {
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS3{F0: fi19, F1: -620, F2: 0}
                if (510 > 255) {
                    fmt.Println(fi19, v3)
                    var v21 int64 = 0
                    for fi22 := 0; fi22 < 4; fi22++ {
                        switch (v1 + "") {
                        case "a":
                            v3 += v3
                        }
                        var v23 int8 = 0
                        _ = fi22
                        _ = v23
                    }
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
        } else {
            if false {
                v24 := GenS1{F0: -101, F1: 4294967295, F2: 0}
                _ = v24
            }
        }
    case 32767:
        v1 = (("\n\t" + (("" + "hello") + (v1 + (("foo bar" + v1) + v1)))) + ((((v1 + v1) + v1) + v1) + (v1 + v1)))
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v3, v1)
    var v25 uint8 = 1
    v1 = "foo bar"
    if (!false) {
        if true {
            if true {
                v1 = v1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        v26 := GenS5{F0: 42, F1: v25}
        _ = v26
    } else {
        if false {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        } else {
            var v27 uint16 = 65535
            switch -501 {
            case -274:
                var v28 uint8 = (v25 << ((741) & 63))
                _ = v28
            case 716:
                v27 = 42
            case 204:
                for fi29 := 0; fi29 < 18; fi29++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v30 string = "foo bar"
                    _ = fi29
                    _ = v30
                }
            }
            _ = v27
        }
    }
    if false {
        if len(v2) > 0 {
            v2[0] = -865
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            for fi32 := 0; fi32 < 15; fi32++ {
                _ = len(v2)
                var v33 uint = 18446744073709551615
                _ = fi32
                _ = v33
            }
            v34 := GenS6{F0: 2147483647, F1: v1}
            if true {
                v35 := GenS7{F0: "foo bar"}
                _ = v35
            } else {
                var v36 uint32 = 2
                var v37 uint64 = 18446744073709551615
                _ = v36
                _ = v37
            }
            _ = fi31
            _ = v34
        }
    } else {
        v3 = (v3 / v3)
    }
    v2 = append(v2, 4294967295)
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
}
