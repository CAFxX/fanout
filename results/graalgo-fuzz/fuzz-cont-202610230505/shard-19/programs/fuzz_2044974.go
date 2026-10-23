package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 int32; F2 int16 }
func genFn0() string {
    var v4 uint64 = 18446744073709551615
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(fi5, v4)
        _ = fi5
    }
    var v6 uint = 255
    for fi7 := 0; fi7 < 7; fi7++ {
        v6 = 18446744073709551615
        v8 := GenS1{F0: 1, F1: 127}
        _ = fi7
        _ = v8
    }
    v4 = v4
    _ = v4
    _ = v6
    return ("foo bar" + "a")
}

func main() {
    var v1 float64 = 1e10
    v2 := []uint32{}
    var v3 int8 = -128
    if (100 != v3) {
        var v9 uint16 = 49
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        switch ("foo bar" + "a") {
        case "hello":
            fmt.Println(v3, v1)
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: 441, F1: -128}
            _ = fi11
            _ = v12
        }
        v1 += v1
        _ = fi10
    }
    v13 := GenS2{F0: 0, F1: 1, F2: 324}
    for fi14 := 0; fi14 < 2; fi14++ {
        switch fi14 {
        case 4294967295:
            var v15 uint = 65535
            _ = v15
        }
        var v16 uint = 405
        fmt.Println(fi14)
        _ = fi14
        _ = v16
    }
    var v17 int = -669
    for fi18 := 0; fi18 < 11; fi18++ {
        switch 127 {
        case 42:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 127:
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := GenS2{F0: 1, F1: 1, F2: 950}
                if false {
                    var v21 int64 = 1
                    if (true || true) {
                        if false {
                            v17 -= fi19
                            fmt.Println(v3, v21)
                        } else {
                            var v22 int16 = 1
                            if len(v2) > 0 {
                                v2[0] = 2147483647
                            }
                            v1 += v1
                            _ = v22
                        }
                        fmt.Println(v3, v21, v1)
                        var v23 uint32 = 488
                        var v24 string = ""
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v3, v21, fi19)
                    if (42 <= -2147483648) {
                        if ((v17 + 1) > -443) {
                            fmt.Println(fi18, fi18)
                            v21 = 0
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            var v25 uint16 = 65535
                            v17 = (^(v17))
                            _ = v25
                        }
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                        var v26 uint64 = 60
                        _ = v26
                    }
                    if false {
                        if false {
                            switch "\n\t" {
                            case "foo bar":
                                v2 = append(v2, 4294967295)
                            case "\n\t":
                                var v27 uint64 = 0
                                v21 = v21
                                var v28 int16 = -850
                                _ = v27
                                _ = v28
                            case "a":
                                fmt.Println(fi19, fi18)
                            }
                        }
                        fmt.Println(fi19, fi19)
                        var v29 int16 = 1000
                        _ = v29
                    }
                    _ = v21
                } else {
                    if len(v2) > 0 {
                        v2[0] = 247
                    }
                }
                _ = fi19
                _ = v20
            }
        case 32767:
            var v30 int = (v17 | 9223372036854775807)
            _ = v30
        default:
            for fi31 := 0; fi31 < 14; fi31++ {
                if len(v2) > 0 {
                    v2[0] = 2
                }
                if false {
                    var v32 int16 = -6
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    _ = v32
                } else {
                    if ((!((v1 + 0.1) != (((1e10 * -2.25) / 0.1) + v1))) && (0 > 0)) {
                        var v33 uint16 = 32767
                        var v34 string = ""
                        _ = v33
                        _ = v34
                    } else {
                        fmt.Println(v17)
                        var v35 uint = 46
                        if true {
                            var v36 int32 = 306
                            v37 := GenS3{F0: 1}
                            _ = v36
                            _ = v37
                        } else {
                            if false {
                                var v38 uint32 = 4294967295
                                _ = v38
                            }
                            fmt.Println(fi18)
                        }
                        _ = v35
                    }
                    switch "" {
                    case "":
                        var v39 uint = 110
                        _ = v39
                    case "foo bar":
                        v17 += fi31
                    default:
                        if true {
                            v3 += (-(v3))
                            v3 = (v3 << ((376) & 63))
                        }
                    }
                    var v40 uint32 = 2147483647
                    _ = v40
                }
                _ = fi31
            }
        }
        var v41 bool = true
        _ = fi18
        _ = v41
    }
    var v42 int16 = 457
    var v43 uint8 = 42
    v2 = append(v2, 2)
    _ = len(v2)
    fmt.Println(v17)
    v44 := GenS4{F0: v43, F1: -2147483648, F2: v42}
    fmt.Println(v1, v1)
    v1 -= (v1 / (0.1 / 0.1))
    fmt.Println(v43)
    fmt.Println(v42, v17, v43)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
}
