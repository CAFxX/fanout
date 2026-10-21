package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int8; F1 int64; F2 int64 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int }
type GenS5 struct { F0 int32 }

func main() {
    var v1 string = ("" + "hello")
    v2 := []int{4294967295, -9223372036854775808}
    var v3 bool = (2 != 4294967295)
    var v4 string = ""
    var v5 uint8 = 1
    for fi6 := 0; fi6 < 1; fi6++ {
        switch v4 {
        case "\n\t":
            var v7 uint16 = 1
            _ = v7
        }
        _ = fi6
    }
    v8 := GenS1{F0: 0, F1: -9223372036854775808}
    var v9 int = 4294967295
    var v10 int32 = 42
    v11 := GenS2{F0: 44, F1: 0, F2: 1}
    v12 := GenS2{F0: int8(v9), F1: 945, F2: 1}
    switch v1 {
    case "\n\t":
        v4 = "hello"
    case "foo bar":
        if true {
            v10 = 2147483647
            var v13 float32 = 0.0
            switch v1 {
            case "":
                v1 = (v1 + ((v4 + v4) + v1))
            case "foo bar":
                var v14 bool = true
                _ = v14
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v15 int = v9
                _ = v15
            }
            _ = v13
        } else {
            for fi16 := 0; fi16 < 14; fi16++ {
                if len(v2) > 0 {
                    v2[0] = (fi16 & fi16)
                }
                v10 += v10
                _ = fi16
            }
            if v3 {
                var v17 bool = ((v1 + (v4 + (v1 + "a"))) != "hello")
                var v18 int64 = 339
                for fi19 := 0; fi19 < 3; fi19++ {
                    fmt.Println(v3, v10, v9)
                    if len(v2) > 0 {
                        v2[0] = fi19
                    }
                    if v3 {
                        fmt.Println(v1, fi19)
                        v20 := GenS3{F0: 4294967295, F1: 0}
                        v2 = append(v2, fi19)
                        _ = v20
                    }
                    _ = fi19
                }
                _ = v17
                _ = v18
            }
        }
    }
    if ((1 != 0) || ((v3 || true) || true)) {
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := GenS1{F0: 65535, F1: 536}
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 uint8 = v5
                switch v9 {
                case 127:
                    v25 := GenS1{F0: 1, F1: 9223372036854775807}
                    _ = v25
                case 741:
                    var v26 string = "a"
                    _ = v26
                }
                fmt.Println(v1)
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        v9 += -696
        if ((500 < 0) && ((v10 - (v10 << ((4294967295) & 63))) != v10)) {
            var v27 uint16 = 42
            _ = v27
        } else {
            switch v9 {
            case 9223372036854775807:
                if (v4 < v4) {
                    for fi28 := 0; fi28 < 5; fi28++ {
                        if (!(42 < 0)) {
                            var v29 string = v4
                            switch v9 {
                            case -69:
                                for fi30 := 0; fi30 < 2; fi30++ {
                                    v10 -= (v10 / (((-(v10))) | 1))
                                    _ = fi30
                                }
                            }
                            v5 = v5
                            _ = v29
                        } else {
                            for fi31 := 0; fi31 < 12; fi31++ {
                                if (v3 && false) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                v32 := GenS1{F0: 4294967295, F1: -846}
                                if len(v2) > 0 {
                                    v2[0] = (fi28 - v9)
                                }
                                _ = fi31
                                _ = v32
                            }
                            var v33 int = v9
                            _ = v33
                        }
                        v9 += (v9 & (v9 * 13))
                        _ = fi28
                    }
                    switch 127 {
                    case -291:
                        var v34 int8 = 42
                        _ = v34
                    case -2147483648:
                        if true {
                            for fi35 := 0; fi35 < 15; fi35++ {
                                for fi36 := 0; fi36 < 6; fi36++ {
                                    switch v9 {
                                    case -2147483648:
                                        var v37 uint16 = 255
                                        _ = v37
                                    case 127:
                                        fmt.Println(v5, v4, fi35)
                                    default:
                                        v38 := GenS4{F0: 2, F1: 566, F2: (fi36 >> ((18446744073709551615) & 63))}
                                        _ = v38
                                    }
                                    _ = len(v2)
                                    _ = fi36
                                }
                                switch (v9 - 848) {
                                case 639:
                                    if ((42 > 42) && (9223372036854775807 <= -9223372036854775808)) {
                                        if len(v2) > 0 {
                                            v2[0] = fi35
                                        }
                                    }
                                case -2147483648:
                                    var v39 float32 = (-2.25 + ((1e10 - 1.5) / 0.1))
                                    _ = v39
                                case 9223372036854775807:
                                    if ((v4 + "foo bar") == "a") {
                                        v40 := GenS5{F0: v10}
                                        _ = v40
                                    }
                                }
                                if len(v2) > 0 {
                                    v2[0] = 32768
                                }
                                _ = fi35
                            }
                        } else {
                            v10 = 1
                        }
                    case 2147483647:
                        var v41 uint64 = 1
                        _ = v41
                    }
                } else {
                    if false {
                        var v42 int8 = 1
                        _ = v42
                    }
                    fmt.Println(v3, v10, v4)
                    if (-1 == 17) {
                        v10 = (v10 + (-(v10)))
                        fmt.Println(v4, v5, v1)
                        v10 -= v10
                    }
                }
            case -365:
                var v43 int32 = -400
                _ = v43
            case -21:
                var v44 uint32 = 903
                _ = v44
            default:
                var v45 int32 = -1
                _ = v45
            }
            var v46 int8 = -1
            _ = v46
        }
    } else {
        var v47 bool = false
        var v48 uint64 = 0
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v47
        _ = v48
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
