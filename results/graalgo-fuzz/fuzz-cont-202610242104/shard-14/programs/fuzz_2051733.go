package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 uint32 = 2147483647
    var v2 int = -9223372036854775808
    v3 := []uint32{v1, v1, v1}
    v4 := []int{(v2 + v2)}
    v5 := []uint32{8}
    v2 = v2
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 100, F1: -2147483648, F2: -14}
        _ = fi6
        _ = v7
    }
    v2 -= 83
    v2 = (v2 % ((1) | 1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: -1, F1: 2147483647, F2: -128}
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 2
    if false {
        switch 4294967295 {
        case 4294967295:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v2 -= -828
    } else {
        fmt.Println(v2, v1, v10)
        for fi11 := 0; fi11 < 14; fi11++ {
            switch "a" {
            case "a":
                var v12 string = "foo bar"
                _ = v12
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = v2
                }
            }
            _ = fi11
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                if ((!false) && (v10 > (v1 & 1))) {
                    if false {
                        switch (-(fi13)) {
                        case 156:
                            fmt.Println(v1, fi13, v10)
                        default:
                            switch ("foo bar" + "\n\t") {
                            case "a":
                                switch ("\n\t" + "a") {
                                case "a":
                                    var v15 float64 = 1e10
                                    _ = v15
                                case "":
                                    fmt.Println(v10, v10)
                                default:
                                    if true {
                                        var v16 uint = 65535
                                        _ = v16
                                    }
                                }
                            case "foo bar":
                                v17 := GenS2{F0: 18446744073709551615}
                                _ = v17
                            default:
                                _ = len(v5)
                            }
                            if false {
                                if true {
                                    if false {
                                        v18 := GenS2{F0: 873}
                                        _ = v18
                                    }
                                    switch "" {
                                    case "":
                                        _ = len(v4)
                                    case "\n\t":
                                        var v19 uint64 = 18446744073709551615
                                        _ = v19
                                    case "a":
                                        var v20 float32 = 1.5
                                        var v21 uint32 = (v10 * (v10 << ((65535) & 63)))
                                        _ = v20
                                        _ = v21
                                    }
                                } else {
                                    if false {
                                        switch 1 {
                                        case 4294967295:
                                            var v22 int32 = 0
                                            _ = v22
                                        case -2147483648:
                                            switch ((("\n\t" + "\n\t") + "foo bar") + ("foo bar" + "foo bar")) {
                                            case "\n\t":
                                                v5 = append(v5, (v1 % ((v1) | 1)))
                                            case "":
                                                v23 := GenS2{F0: 4294967295}
                                                v24 := GenS3{F0: -1, F1: 2, F2: 42}
                                                if (!(false || true)) {
                                                    fmt.Println(v10)
                                                }
                                                _ = v23
                                                _ = v24
                                            case "a":
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            default:
                                                if len(v3) > 0 {
                                                    v3[0] = (^(v10))
                                                }
                                            }
                                        default:
                                            v1 += (v1 % ((0) | 1))
                                            var v25 uint = 255
                                            v26 := GenS2{F0: 4294967295}
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = len(v5)
                                    }
                                    v10 -= 2147483648
                                    v1 += (v1 | (v1 * 1))
                                }
                                var v27 uint16 = 1
                                v28 := GenS4{F0: 2, F1: 4294967295}
                                var v29 int64 = 9223372036854775807
                                v30 := GenS3{F0: 42, F1: v27, F2: int64(v10)}
                                var v31 int64 = -881
                                _ = v27
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            var v32 float32 = (0.1 / 1.0)
                            var v33 string = "foo bar"
                            fmt.Println(v2, v33, v2)
                            _ = v32
                            _ = v33
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        v2 = 2147483647
                    }
                    if ((v10 << ((943) & 63)) < 1) {
                        if len(v3) > 0 {
                            v3[0] = (^(v10))
                        }
                    } else {
                        var v34 uint8 = 255
                        _ = v34
                    }
                    v4 = append(v4, v2)
                } else {
                    if len(v5) > 0 {
                        v5[0] = v10
                    }
                }
                _ = fi14
            }
            for fi35 := 0; fi35 < 11; fi35++ {
                if (255 <= -1) {
                    v2 += fi13
                }
                var v36 float32 = -2.25
                _ = fi35
                _ = v36
            }
            switch ("" + "") {
            case "foo bar":
                _ = len(v5)
            }
            _ = fi13
        }
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        v38 := GenS2{F0: 4294967295}
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 10; fi39++ {
        var v40 int64 = -1
        v10 = v1
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
