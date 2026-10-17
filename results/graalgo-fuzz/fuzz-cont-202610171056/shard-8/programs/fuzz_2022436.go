package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint8 }
func genFn0(p0 int16, p1 int64) int8 {
    v5 := []int{-139}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        var v6 int8 = 42
        var v7 string = "hello"
        var v8 uint8 = 1
        _ = v6
        _ = v7
        _ = v8
    } else {
        if (int(p0) <= -1) {
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 int8 = -108
                v11 := GenS1{F0: 615, F1: 683, F2: (p1 & 35)}
                var v12 uint = 65535
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            p1 -= 0
        } else {
            var v13 uint8 = 224
            _ = v13
        }
        switch 128 {
        case 1:
            var v14 float64 = 1.5
            _ = v14
        case -2147483648:
            v15 := GenS2{F0: 100, F1: 143}
            _ = v15
        case -400:
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                if true {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        p0 = 696
                        _ = fi16
                    }
                    var v17 uint = 802
                    _ = v17
                } else {
                    var v18 int8 = 0
                    _ = len(v5)
                    _ = v18
                }
            }
        default:
            if true {
                if len(v5) > 0 {
                    v5[0] = 32767
                }
                var v19 int16 = (p0 << ((18446744073709551615) & 63))
                fmt.Println(p1, p1)
                _ = v19
            } else {
                if len(v5) > 0 {
                    v5[0] = 69
                }
                switch "a" {
                case "hello":
                    p1 = (p1 & -811)
                case "a":
                    if true {
                        v20 := GenS1{F0: 2, F1: 881, F2: 340}
                        _ = v20
                    } else {
                        v21 := GenS3{F0: 0}
                        p0 += p0
                        _ = v21
                    }
                case "":
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 int8 = 0
                        _ = fi22
                        _ = v23
                    }
                }
            }
        }
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 uint16 = 0
            if len(v5) > 0 {
                v5[0] = fi24
            }
            _ = fi24
            _ = v25
        }
    }
    v5 = append(v5, 42)
    _ = p0
    _ = p1
    _ = v5
    return -1
}
func genFn1(p0 int16) int16 {
    var v26 uint64 = 4294967295
    var v27 int32 = 2147483647
    for fi28 := 0; fi28 < 19; fi28++ {
        v27 += v27
        var v29 float32 = 3.14159
        _ = fi28
        _ = v29
    }
    _ = p0
    _ = v26
    _ = v27
    return (^(p0))
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    var v3 bool = (1 < 0)
    var v4 int16 = 0
    var v30 string = ("a" + "hello")
    v31 := []int{1, 32767, -2147483648, 2147483647}
    if len(v31) > 0 {
        v31[0] = 1
    }
    if true {
        v1[128] = (v30 + v30)
    } else {
        if (v4 <= (v4 & v4)) {
            var v32 int32 = -335
            v33 := GenS1{F0: 42, F1: 586, F2: 9223372036854775807}
            for fi34 := 0; fi34 < 10; fi34++ {
                v35 := GenS2{F0: 127, F1: 178}
                v32 = (v32 / (((^(v32))) | 1))
                _ = fi34
                _ = v35
            }
            _ = v32
            _ = v33
        }
        if false {
            var v36 bool = (499 <= 18446744073709551615)
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        fmt.Println(v4, fi37, v4)
        var v38 bool = true
        _ = fi37
        _ = v38
    }
    switch -145 {
    case 32767:
        v39 := GenS4{F0: 42, F1: 127, F2: 127}
        _ = v39
    case 2147483647:
        fmt.Println(v3, v3)
    }
    v4 = v4
    v3 = false
    var v40 uint8 = 128
    v4 -= -32768
    var v41 int8 = genFn0(v4, 42)
    v30 = ((v30 + ((((v30 + v30) + "foo bar") + v30) + v30)) + ("a" + (v30 + v30)))
    v31 = append(v31, 997)
    fmt.Println(v3, v40)
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
    _ = v41
}
