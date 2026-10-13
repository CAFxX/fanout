package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int16; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 int64; F1 uint; F2 bool }
type GenS8 struct { F0 uint32 }
func genFn0(p0 bool, p1 string) int64 {
    v5 := GenS1{F0: 235}
    var v6 uint8 = 1
    v7 := []int{32768}
    if (((^(v6)) <= (v6 - (v6 ^ (v6 / ((127) | 1))))) && p0) {
        v6 = (v6 - (v6 + v6))
        fmt.Println(p0, p1)
        var v8 float64 = 0.0
        _ = v8
    }
    v7 = append(v7, 9223372036854775807)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 88
}

func main() {
    var v1 uint16 = 32767
    var v2 float32 = (-2.25 - 0.0)
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint32 = 911
            v1 -= 32767
            if true {
                v12 := GenS1{F0: v11}
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        v1 = (v1 << ((1) & 63))
        var v13 float32 = (v2 / v2)
        _ = fi9
        _ = v13
    }
    v14 := GenS1{F0: 0}
    v2 = v2
    if false {
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := v3[1]
            if (112 != 42) {
                if true {
                    var v17 uint64 = 1
                    v18 := []int{(fi15 * v16), (^(v16))}
                    v19 := GenS2{F0: (v17 * (v17 | v17)), F1: 0, F2: 0}
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v20 := GenS2{F0: 1, F1: 42, F2: -32768}
                v16 -= 42
                _ = v20
            }
            _ = fi15
            _ = v16
        }
    }
    v21 := GenS2{F0: 4294967295, F1: 42, F2: -1}
    if true {
        if (-826 >= 838) {
            v22 := []int{}
            var v23 int64 = 42
            var v24 bool = true
            _ = v22
            _ = v23
            _ = v24
        } else {
            if false {
                var v25 int8 = 42
                _ = v25
            }
        }
        var v26 uint = 42
        _ = v26
    }
    v27 := []int{-765}
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := GenS3{F0: 190}
        for fi30 := 0; fi30 < 1; fi30++ {
            if false {
                if ((true || false) && true) {
                    v1 = v1
                    var v31 uint32 = 2
                    v31 = v31
                    var v32 string = ""
                    _ = v31
                    _ = v32
                } else {
                    v33 := GenS4{F0: 0, F1: 757}
                    fmt.Println(v2, fi28, fi30)
                    _ = v33
                }
            }
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        fmt.Println(fi34)
        _ = fi34
    }
    v2 = (float32(((0.0 + 3.14159) * -2.25)) + v2)
    switch ("" + "") {
    case "foo bar":
        v1 = (v1 / ((v1) | 1))
    case "":
        var v35 string = "hello"
        _ = v35
    }
    v36 := v4[9223372036854775807]
    v36 += v36
    var v37 int = 32768
    if false {
        _ = len(v4)
        v38 := v4[9223372036854775807]
        for fi39 := 0; fi39 < 12; fi39++ {
            fmt.Println(fi39)
            if (!(895 >= 42)) {
                v40 := GenS5{F0: 4294967295, F1: 2}
                _ = v40
            } else {
                if false {
                    if false {
                        v41 := GenS6{F0: 18446744073709551615, F1: v37}
                        _ = v41
                    }
                    var v42 int8 = -19
                    _ = v42
                } else {
                    var v43 int32 = -1
                    v43 -= (v43 + 2147483647)
                    switch (fi39 ^ v37) {
                    case 9223372036854775807:
                        fmt.Println(v38)
                    default:
                        for fi44 := 0; fi44 < 2; fi44++ {
                            v38 += v38
                            _ = fi44
                        }
                    }
                    _ = v43
                }
                v27 = append(v27, v37)
            }
            for fi45 := 0; fi45 < 16; fi45++ {
                var v46 uint = 2
                switch "\n\t" {
                case "\n\t":
                    v36 = (v38 + (v38 & v38))
                default:
                    v46 -= v46
                }
                _ = fi45
                _ = v46
            }
            _ = fi39
        }
        _ = v38
    }
    for fi47 := 0; fi47 < 8; fi47++ {
        var v48 uint64 = 18446744073709551615
        v49 := GenS7{F0: genFn0((411 < 1), ("a" + ("" + "hello"))), F1: 65535, F2: ((!(!false)) && (((false && false) || ((!true) && (v37 >= int(v1)))) && true))}
        v50 := GenS8{F0: 170}
        _ = fi47
        _ = v48
        _ = v49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
