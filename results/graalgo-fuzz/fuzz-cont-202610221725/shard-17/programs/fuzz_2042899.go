package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int; F2 int8 }
type GenS5 struct { F0 int8; F1 int64; F2 int32 }
type GenS6 struct { F0 bool; F1 int32; F2 string }
func genFn0(p0 int, p1 int) int16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        if true {
            for fi6 := 0; fi6 < 18; fi6++ {
                v7 := GenS1{F0: ""}
                v8 := []int{p0, 42, fi5}
                var v9 int64 = -9223372036854775808
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            v10 := GenS2{F0: 565}
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := []int{}
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        v13 := []int{404}
        _ = fi5
        _ = v13
    }
    v14 := []int{p0}
    v14 = append(v14, (p1 * (p0 & -23)))
    _ = p0
    _ = p1
    _ = v14
    return 1000
}
func genFn1(p0 uint) uint64 {
    fmt.Println(p0)
    p0 = 18446744073709551615
    _ = p0
    return 2
}

func main() {
    var v1 bool = true
    var v2 float32 = 3.14159
    v3 := make(map[uint32]int)
    v4 := []int{42}
    if v1 {
        if v1 {
            var v15 uint16 = 844
            v16 := GenS3{F0: "foo bar"}
            _ = v15
            _ = v16
        } else {
            var v17 string = ("foo bar" + "a")
            _ = v17
        }
        delete(v3, 760)
    } else {
        var v18 bool = false
        _ = v18
    }
    v19 := GenS3{F0: (("foo bar" + ("\n\t" + ("a" + ("\n\t" + "")))) + "")}
    v2 += 1e10
    fmt.Println(v1)
    switch ("foo bar" + "hello") {
    case "":
        v2 = ((v2 + ((float32((1e10 + 1.5)) + 1e10) * 3.14159)) * ((-2.25 - (v2 - (float32(-2.25) - 0.0))) / v2))
    case "hello":
        v20 := GenS3{F0: "\n\t"}
        _ = v20
    case "\n\t":
        v21 := GenS1{F0: "\n\t"}
        _ = v21
    }
    fmt.Println(v2, v1)
    switch 52 {
    case 32768:
        if v1 {
            fmt.Println(v1)
        }
    case -893:
        var v22 float64 = 3.14159
        _ = v22
    default:
        switch "hello" {
        case "hello":
            fmt.Println(v2, v2)
        case "foo bar":
            v23 := GenS1{F0: "foo bar"}
            _ = v23
        }
    }
    switch 1 {
    case -1:
        var v24 bool = v1
        _ = v24
    case 127:
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 bool = v1
            fmt.Println(v1)
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 uint16 = 0
                if true {
                    var v29 int64 = 887
                    var v30 int64 = v29
                    _ = v29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
            _ = fi25
            _ = v26
        }
    case 2147483647:
        switch "" {
        case "hello":
            v1 = (-385 <= -775)
        }
    default:
        switch 127 {
        case 127:
            var v31 int8 = -1
            _ = v31
        case 183:
            if v1 {
                var v32 int8 = 127
                _ = v32
            }
        case 128:
            if ("" < "a") {
                for fi33 := 0; fi33 < 8; fi33++ {
                    _ = len(v4)
                    for fi34 := 0; fi34 < 5; fi34++ {
                        v2 += v2
                        _ = fi34
                    }
                    v35 := GenS2{F0: 0}
                    _ = fi33
                    _ = v35
                }
                for fi36 := 0; fi36 < 7; fi36++ {
                    var v37 uint8 = 128
                    if v1 {
                        v1 = (578 <= 42)
                    } else {
                        v38 := GenS4{F0: 0, F1: (fi36 / ((9223372036854775807) | 1)), F2: -128}
                        _ = v38
                    }
                    _ = fi36
                    _ = v37
                }
            }
        default:
            if (397 == 42) {
                for fi39 := 0; fi39 < 18; fi39++ {
                    for fi40 := 0; fi40 < 12; fi40++ {
                        if v1 {
                            v2 -= (1e10 * ((float32((1.5 - (float64((v2 + ((((float32(1e10) / v2) / 3.14159) * v2) * (v2 * 1e10)))) - 3.14159))) / 1.5) * (1e10 - v2)))
                            if ((1e10 * float64(float32(3.14159))) <= float64(v2)) {
                                var v41 int32 = 360
                                v1 = v1
                                fmt.Println(v2)
                                _ = v41
                            } else {
                                v1 = v1
                                if true {
                                    var v42 uint64 = 42
                                    _ = v42
                                }
                                v4 = append(v4, fi40)
                            }
                            var v43 int16 = -932
                            _ = v43
                        }
                        var v44 int32 = 1
                        var v45 int32 = v44
                        _ = fi40
                        _ = v44
                        _ = v45
                    }
                    fmt.Println(v2, fi39, v1)
                    var v46 float32 = (1.5 + (3.14159 * 0.0))
                    _ = fi39
                    _ = v46
                }
            } else {
                var v47 uint = 983
                _ = v47
            }
        }
    }
    v2 -= v2
    for fi48 := 0; fi48 < 7; fi48++ {
        v49 := GenS5{F0: 127, F1: -964, F2: 42}
        v2 -= (0.1 / (v2 / v2))
        v50 := GenS6{F0: true, F1: 0, F2: ""}
        _ = fi48
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 1; fi51++ {
        for fi52 := 0; fi52 < 11; fi52++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v53 uint16 = 1
            _ = fi52
            _ = v53
        }
        switch fi51 {
        case 32768:
            var v54 uint16 = 1
            _ = v54
        case -587:
            v2 = v2
        case -2147483648:
            fmt.Println(fi51, v1)
        }
        fmt.Println(v1)
        _ = fi51
    }
    v2 = float32((-2.25 - ((1.5 - ((1e10 - ((((float64(v2) / -2.25) - float64((v2 - v2))) * (0.1 + 3.14159)) / float64(((v2 - v2) / (v2 * 3.14159))))) / ((float64(float32(-2.25)) / ((((-2.25 - -2.25) / 1.5) + 1.5) / (-2.25 + (0.0 + 3.14159)))) + (float64((3.14159 / -2.25)) / 1.0)))) * 1.5)))
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
