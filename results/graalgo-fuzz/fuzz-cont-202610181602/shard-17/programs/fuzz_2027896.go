package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int32 = 42
    var v2 bool = (((451 <= 32767) || false) || true)
    var v3 bool = (4294967295 < -189)
    var v4 uint16 = 65535
    var v5 float64 = 1.5
    var v6 float32 = (((float32((0.1 / v5)) * 3.14159) + 1.5) - (0.1 + (1.5 - (((-2.25 + 1.5) * ((3.14159 / -2.25) + (-2.25 + 0.1))) / float32(1.5)))))
    v4 = 507
    v7 := []int{}
    var v8 int64 = 9223372036854775807
    switch -2147483648 {
    case -1:
        v6 = float32((v5 * (v5 / v5)))
    case 0:
        if len(v7) > 0 {
            v7[0] = -9223372036854775808
        }
    default:
        if v2 {
            fmt.Println(v3)
        }
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        switch "\n\t" {
        case "hello":
            var v10 int16 = 1
            _ = v10
        default:
            v3 = (42 < v1)
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 int64 = (v8 % ((v8) | 1))
            var v13 uint32 = 536
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 int32 = v1
            fmt.Println(v2, v15)
            _ = fi14
            _ = v15
        }
        _ = fi9
    }
    v16 := GenS1{F0: -10}
    v3 = v2
    var v17 int16 = -401
    v18 := GenS1{F0: 100}
    v2 = (v8 > v8)
    fmt.Println(v17, v6, v3)
    switch 1 {
    case 32767:
        for fi19 := 0; fi19 < 18; fi19++ {
            v6 -= 0.0
            fmt.Println(v4, v1, v5)
            v3 = (v4 == v4)
            _ = fi19
        }
    case -1:
        switch -474 {
        case -2147483648:
            if v2 {
                v1 += (v1 & v1)
                v20 := GenS1{F0: 123}
                if (!v2) {
                    if ((float64(0.0) - v5) > v5) {
                        for fi21 := 0; fi21 < 13; fi21++ {
                            v22 := GenS2{F0: 1, F1: (v1 + (^(v1)))}
                            if v3 {
                                var v23 int = fi21
                                switch (("hello" + "foo bar") + "\n\t") {
                                case "a":
                                    v6 = v6
                                }
                                _ = v23
                            }
                            _ = fi21
                            _ = v22
                        }
                    } else {
                        var v24 bool = true
                        switch ("" + "\n\t") {
                        case "foo bar":
                            fmt.Println(v6, v8)
                        case "":
                            var v25 uint16 = v4
                            _ = v25
                        default:
                            for fi26 := 0; fi26 < 12; fi26++ {
                                v6 = (1.5 / v6)
                                for fi27 := 0; fi27 < 3; fi27++ {
                                    if true {
                                        fmt.Println(v17)
                                    } else {
                                        var v28 uint = 932
                                        v2 = v24
                                        _ = len(v7)
                                        if ((!true) && v24) {
                                            if (!v3) {
                                                v24 = (((!(-1 == v17)) || true) || true)
                                                var v29 uint32 = 4294967295
                                                if v3 {
                                                    v7 = append(v7, fi27)
                                                } else {
                                                    v30 := GenS3{F0: (v17 + 0)}
                                                    var v31 int8 = int8(v28)
                                                    _ = v30
                                                    _ = v31
                                                }
                                                _ = v29
                                            }
                                            fmt.Println(v28)
                                        }
                                        _ = v28
                                    }
                                    _ = fi27
                                }
                                _ = fi26
                            }
                        }
                        if len(v7) > 0 {
                            v7[0] = 127
                        }
                        _ = v24
                    }
                    if v2 {
                        var v32 uint8 = 2
                        var v33 bool = v2
                        _ = v32
                        _ = v33
                    } else {
                        var v34 int8 = -68
                        _ = v34
                    }
                    v5 += 3.14159
                }
                _ = v20
            }
        default:
            var v35 int16 = v17
            _ = v35
        }
    default:
        if false {
            switch -263 {
            case 32768:
                v36 := GenS3{F0: v17}
                _ = v36
            case 9223372036854775807:
                for fi37 := 0; fi37 < 4; fi37++ {
                    for fi38 := 0; fi38 < 5; fi38++ {
                        var v39 float32 = 1.5
                        fmt.Println(v2, v17, v39)
                        _ = fi38
                        _ = v39
                    }
                    v4 = v4
                    _ = fi37
                }
            case 4294967295:
                for fi40 := 0; fi40 < 4; fi40++ {
                    _ = len(v7)
                    var v41 uint = 94
                    if false {
                        fmt.Println(v4)
                        switch "" {
                        case "foo bar":
                            if (!true) {
                                var v42 int32 = (v1 << ((v41) & 63))
                                _ = v42
                            } else {
                                switch "" {
                                case "\n\t":
                                    fmt.Println(v41, v5)
                                default:
                                    v43 := GenS3{F0: (^(v17))}
                                    _ = v43
                                }
                            }
                        case "":
                            fmt.Println(v1, v6, v41)
                        case "hello":
                            fmt.Println(fi40)
                        default:
                            if true {
                                var v44 uint8 = 0
                                _ = v44
                            } else {
                                var v45 int64 = int64(v1)
                                _ = v45
                            }
                        }
                    }
                    _ = fi40
                    _ = v41
                }
            }
        } else {
            if false {
                var v46 uint8 = 188
                _ = v46
            }
            v1 -= 852
            var v47 bool = (v3 && (v3 && v3))
            _ = v47
        }
    }
    v48 := GenS3{F0: v17}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
