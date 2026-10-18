package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 int32 }
func genFn0(p0 bool) uint {
    v6 := []int{32767, -9223372036854775808, 9223372036854775807}
    for fi7 := 0; fi7 < 9; fi7++ {
        if (false || false) {
            switch (("\n\t" + "foo bar") + ("\n\t" + "")) {
            case "a":
                p0 = (((!(4294967295 > 2147483648)) && p0) || p0)
            case "":
                var v8 int = fi7
                _ = v8
            case "hello":
                v9 := GenS1{F0: 65535, F1: -1}
                _ = v9
            }
            var v10 int8 = -128
            _ = v10
        } else {
            var v11 bool = p0
            fmt.Println(fi7)
            var v12 int64 = -730
            _ = v11
            _ = v12
        }
        p0 = false
        _ = fi7
    }
    _ = p0
    _ = v6
    return 2
}
func genFn1() uint8 {
    for fi13 := 0; fi13 < 12; fi13++ {
        if (0 == 42) {
            if ((fi13 >> ((57) & 63)) < 127) {
                var v14 bool = (-292 < 9223372036854775807)
                _ = v14
            } else {
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 float32 = 1e10
                    _ = fi15
                    _ = v16
                }
                switch 702 {
                case -9223372036854775808:
                    if (!(!true)) {
                        if (true && true) {
                            var v17 float32 = (1e10 + float32((3.14159 / (1e10 * -2.25))))
                            v17 = v17
                            v18 := []int{(fi13 - (fi13 | fi13)), fi13, 1, (-(fi13))}
                            _ = v17
                            _ = v18
                        }
                    } else {
                        v19 := GenS2{F0: -1}
                        _ = v19
                    }
                case -308:
                    var v20 uint64 = 2
                    _ = v20
                default:
                    if false {
                        v21 := GenS1{F0: 18446744073709551615, F1: -1}
                        _ = v21
                    } else {
                        v22 := []int{(fi13 | fi13)}
                        fmt.Println(fi13)
                        switch ("foo bar" + (("foo bar" + ("hello" + "hello")) + "foo bar")) {
                        case "a":
                            v23 := GenS3{F0: 1, F1: 1}
                            _ = v23
                        case "\n\t":
                            v24 := GenS2{F0: 32767}
                            _ = v24
                        case "hello":
                            if false {
                                fmt.Println(fi13)
                                v25 := GenS4{F0: 255}
                                _ = v25
                            }
                        default:
                            if (-1 != 0) {
                                var v26 int64 = 42
                                _ = v26
                            }
                        }
                        _ = v22
                    }
                }
            }
            switch fi13 {
            case 2147483647:
                var v27 bool = (true || (true && false))
                _ = v27
            }
        } else {
            var v28 int8 = -48
            fmt.Println(fi13, v28)
            _ = v28
        }
        if (0.1 >= 1.5) {
            v29 := []int{(^(fi13)), (-(fi13)), 1, fi13}
            _ = v29
        }
        _ = fi13
    }
    var v30 float64 = float64(0.1)
    v30 += (v30 - (v30 + v30))
    fmt.Println(v30)
    var v31 int32 = 0
    _ = v30
    _ = v31
    return 127
}

func main() {
    v1 := []int{1, 127, 32767}
    var v2 uint16 = 1
    var v3 float64 = 1.5
    v4 := make(map[string]int)
    var v5 uint64 = 18446744073709551615
    if (0.1 >= (float32((v3 * (1e10 - 1e10))) + 0.0)) {
        if true {
            fmt.Println(v3)
            var v32 uint16 = 65535
            _ = v32
        } else {
            v33 := GenS2{F0: 1000}
            v34 := GenS3{F0: 42, F1: -128}
            for fi35 := 0; fi35 < 17; fi35++ {
                v3 -= v3
                v3 += ((v3 * 3.14159) - v3)
                for fi36 := 0; fi36 < 4; fi36++ {
                    var v37 uint64 = 783
                    _ = fi36
                    _ = v37
                }
                _ = fi35
            }
            _ = v33
            _ = v34
        }
        switch 499 {
        case 9223372036854775807:
            v5 += uint64(v2)
        }
    }
    _ = len(v4)
    v2 += (v2 - (v2 | (v2 - (v2 | v2))))
    v38 := GenS2{F0: 32767}
    switch 127 {
    case 32768:
        switch "hello" {
        case "\n\t":
            var v39 int16 = 32767
            _ = v39
        }
    case -245:
        var v40 uint64 = 42
        _ = v40
    case 127:
        for fi41 := 0; fi41 < 15; fi41++ {
            v42 := GenS5{F0: ("\n\t" + "hello"), F1: 4294967295}
            fmt.Println(fi41, v3, v5)
            if false {
                v43 := GenS5{F0: ("foo bar" + ""), F1: 27}
                _ = v43
            }
            _ = fi41
            _ = v42
        }
    default:
        v3 = v3
    }
    switch 42 {
    case 1:
        var v44 float64 = 3.14159
        _ = v44
    case -9223372036854775808:
        v45 := GenS2{F0: 0}
        _ = v45
    case -2147483648:
        v5 -= (v5 * v5)
    default:
        v1 = append(v1, 42)
    }
    _ = len(v4)
    v46 := GenS6{F0: -2147483648}
    fmt.Println(v3, v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v46
    _ = v5
}
