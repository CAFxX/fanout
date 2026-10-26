package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 string; F2 int64 }
func genFn0(p0 uint8, p1 uint32, p2 uint64) int8 {
    p1 = p1
    fmt.Println(p0)
    v4 := []int{}
    v5 := GenS1{F0: (p2 | uint64(p0))}
    var v6 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int16, p1 string, p2 uint16) int16 {
    switch (p1 + "foo bar") {
    case "\n\t":
        switch "a" {
        case "":
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := GenS1{F0: 119}
                _ = fi7
                _ = v8
            }
        default:
            v9 := GenS1{F0: 0}
            _ = v9
        }
    case "a":
        var v10 int16 = (p0 >> ((712) & 63))
        _ = v10
    case "hello":
        v11 := []int{}
        _ = v11
    }
    p2 = p2
    switch (("\n\t" + "hello") + p1) {
    case "\n\t":
        for fi12 := 0; fi12 < 11; fi12++ {
            if ((fi12 / ((fi12) | 1)) > (fi12 | 9223372036854775807)) {
                var v13 int16 = (-(p0))
                var v14 int64 = 0
                _ = v13
                _ = v14
            }
            var v15 uint32 = 2147483647
            p1 = p1
            _ = fi12
            _ = v15
        }
    }
    if (genFn0(uint8(p2), 2147483647, 802) >= -1) {
        v16 := []int{4294967295, 9223372036854775807, 1, -9223372036854775808}
        var v17 float32 = ((1.5 / -2.25) * float32(0.0))
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        _ = v16
        _ = v17
    } else {
        if true {
            var v18 uint8 = 104
            p0 = p0
            _ = v18
        } else {
            if (true && true) {
                var v19 int32 = 2147483647
                switch "foo bar" {
                case "foo bar":
                    v20 := []int{0}
                    _ = v20
                case "":
                    var v21 uint16 = p2
                    _ = v21
                case "\n\t":
                    var v22 uint64 = 42
                    _ = v22
                default:
                    p1 = ""
                }
                var v23 int16 = (p0 + (p0 % (((-(p0))) | 1)))
                _ = v19
                _ = v23
            }
            switch 32768 {
            case -2147483648:
                fmt.Println(p1, p1)
            case 9223372036854775807:
                var v24 string = "a"
                _ = v24
            }
        }
    }
    v25 := []int{508}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 0
}

func main() {
    var v1 string = ""
    var v2 int = 127
    var v3 uint32 = 283
    v26 := []int{}
    var v27 float64 = -2.25
    switch (v2 | v2) {
    case -743:
        if (1 <= 911) {
            if true {
                var v28 int16 = -32768
                var v29 uint8 = 42
                if (1 == v29) {
                    switch (v2 + v2) {
                    case 9223372036854775807:
                        fmt.Println(v28, v2)
                    default:
                        v29 -= v29
                    }
                    if false {
                        switch 9223372036854775807 {
                        case -442:
                            if true {
                                if len(v26) > 0 {
                                    v26[0] = (v2 & v2)
                                }
                                fmt.Println(v28)
                                for fi30 := 0; fi30 < 6; fi30++ {
                                    for fi31 := 0; fi31 < 10; fi31++ {
                                        v26 = append(v26, (fi30 % ((fi31) | 1)))
                                        v26 = append(v26, (fi31 & v2))
                                        _ = fi31
                                    }
                                    switch fi30 {
                                    case 42:
                                        if len(v26) > 0 {
                                            v26[0] = 1
                                        }
                                    case -2147483648:
                                        var v32 uint16 = 255
                                        _ = v32
                                    case 0:
                                        v26 = append(v26, (v2 >> ((218) & 63)))
                                    default:
                                        fmt.Println(v29, fi30, v2)
                                    }
                                    _ = fi30
                                }
                            } else {
                                switch (v1 + v1) {
                                case "hello":
                                    v3 = 415
                                case "foo bar":
                                    v33 := GenS1{F0: 111}
                                    _ = v33
                                case "\n\t":
                                    var v34 uint = 42
                                    _ = v34
                                }
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                if len(v26) > 0 {
                                    v26[0] = (v2 << ((2) & 63))
                                }
                            }
                        case -1:
                            v1 = (v1 + v1)
                        default:
                            v27 = 0.1
                        }
                        fmt.Println(v2, v27)
                    } else {
                        switch v1 {
                        case "foo bar":
                            switch ("" + (v1 + (v1 + "hello"))) {
                            case "hello":
                                for fi35 := 0; fi35 < 7; fi35++ {
                                    v28 = (-(v28))
                                    v27 = 1e10
                                    for fi36 := 0; fi36 < 19; fi36++ {
                                        if (!(9223372036854775807 == -1)) {
                                            var v37 bool = false
                                            _ = v37
                                        }
                                        if (!true) {
                                            v38 := GenS1{F0: 688}
                                            var v39 bool = false
                                            _ = v38
                                            _ = v39
                                        } else {
                                            v2 = fi35
                                            v26 = append(v26, (fi36 * fi35))
                                        }
                                        _ = fi36
                                    }
                                    _ = fi35
                                }
                            case "\n\t":
                                for fi40 := 0; fi40 < 20; fi40++ {
                                    var v41 uint64 = 0
                                    _ = fi40
                                    _ = v41
                                }
                            }
                        default:
                            v3 = (v3 / ((738) | 1))
                        }
                        v26 = append(v26, v2)
                        v42 := GenS1{F0: 826}
                        _ = v42
                    }
                }
                _ = v28
                _ = v29
            }
            v2 -= (v2 << ((0) & 63))
            var v43 bool = (!true)
            _ = v43
        } else {
            switch v2 {
            case -826:
                var v44 int32 = 2147483647
                _ = v44
            case 127:
                v2 += v2
            default:
                if len(v26) > 0 {
                    v26[0] = 42
                }
            }
            v27 = (((v27 - float64((1.5 - -2.25))) * v27) / float64((0.1 * 1e10)))
        }
    }
    var v45 int8 = 0
    v2 = -729
    v46 := GenS1{F0: 0}
    switch ((v1 + (("hello" + "") + v1)) + "a") {
    case "hello":
        for fi47 := 0; fi47 < 20; fi47++ {
            fmt.Println(v2, v3)
            v45 = 42
            _ = fi47
        }
    case "":
        switch v1 {
        case "a":
            switch (v2 ^ v2) {
            case 32768:
                fmt.Println(v45, v27, v45)
            case -2147483648:
                v48 := GenS1{F0: 466}
                _ = v48
            }
        case "hello":
            for fi49 := 0; fi49 < 15; fi49++ {
                fmt.Println(v3, v27)
                var v50 uint64 = 2
                v3 += v3
                _ = fi49
                _ = v50
            }
        case "\n\t":
            var v51 int = v2
            _ = v51
        default:
            fmt.Println(v1, v2)
        }
    default:
        fmt.Println(v3)
    }
    v52 := GenS2{F0: -531, F1: v1, F2: -777}
    fmt.Println(v2, v1, v2)
    var v53 string = (v1 + ("" + ((v1 + (v1 + v1)) + v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v45
    _ = v46
    _ = v52
    _ = v53
}
