package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int64; F1 int; F2 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 810
    var v3 uint = 65535
    var v4 int = -2147483648
    v5 := make(map[string]bool)
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint32 = 65535
        var v9 int8 = -28
        var v10 uint8 = 42
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 float32 = 1e10
    v12 := []int{}
    if (43 >= 1) {
        if (uint(v2) >= (v3 + (v3 - (v3 >> ((4294967295) & 63))))) {
            switch ("foo bar" + "a") {
            case "a":
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 float64 = (-2.25 + -2.25)
                    fmt.Println(v2, v14)
                    _ = fi13
                    _ = v14
                }
            case "hello":
                var v15 int32 = -2147483648
                _ = v15
            default:
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 uint32 = v2
                    for fi18 := 0; fi18 < 14; fi18++ {
                        switch "" {
                        case "":
                            var v19 uint = v3
                            _ = v19
                        case "hello":
                            v12 = append(v12, (-(fi18)))
                        case "\n\t":
                            v20 := v6[(v2 & (v17 | 2))]
                            _ = v20
                        default:
                            fmt.Println(v11, v17, v3)
                        }
                        var v21 int16 = 1
                        v22 := v6[2]
                        v23 := v6[v22]
                        _ = fi18
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 14; fi24++ {
                        v2 = uint32(v3)
                        if false {
                            v2 += v17
                            var v25 int16 = -1
                            switch (fi24 >> ((v3) & 63)) {
                            case 802:
                                v17 -= v2
                            case 127:
                                v4 -= (fi24 - fi16)
                            default:
                                if false {
                                    fmt.Println(v25)
                                } else {
                                    switch (v4 / (((-(fi16))) | 1)) {
                                    case 2147483647:
                                        delete(v6, 11)
                                    case 32768:
                                        switch ("" + "") {
                                        case "a":
                                            fmt.Println(v17, v2)
                                        }
                                    default:
                                        switch (fi16 * (-(fi16))) {
                                        case -2147483648:
                                            var v26 int32 = -2147483648
                                            _ = v26
                                        case 127:
                                            v4 += fi16
                                        default:
                                            fmt.Println(v17)
                                        }
                                    }
                                    switch "" {
                                    case "a":
                                        var v27 uint8 = 2
                                        _ = v27
                                    case "":
                                        v25 = v25
                                    default:
                                        var v28 uint16 = 2
                                        _ = v28
                                    }
                                    v3 = (v3 % ((v3) | 1))
                                }
                            }
                            _ = len(v5)
                            _ = v25
                        }
                        if true {
                            _ = len(v12)
                            switch (v4 >> (((v3 >> ((uint(v17)) & 63))) & 63)) {
                            case 32768:
                                v11 -= ((v11 * 1.5) - v11)
                            case -56:
                                if (((((0.0 * -2.25) * (3.14159 - -2.25)) * (-2.25 + 0.1)) > 0.1) || (0 != 22)) {
                                    fmt.Println(v2)
                                    _ = len(v6)
                                } else {
                                    v29 := GenS1{F0: 9223372036854775807}
                                    var v30 bool = (608 != 600)
                                    _ = v29
                                    _ = v30
                                }
                            case -650:
                                v31 := GenS2{F0: (v17 & (v2 - v17)), F1: 0}
                                _ = v31
                            default:
                                v32 := GenS2{F0: 42, F1: -1}
                                _ = v32
                            }
                            delete(v5, "\n\t")
                        } else {
                            switch (("\n\t" + "a") + ("\n\t" + (("\n\t" + ("" + ("foo bar" + "a"))) + "hello"))) {
                            case "foo bar":
                                var v33 uint8 = uint8(v3)
                                _ = v33
                            }
                            v34 := GenS1{F0: -9223372036854775808}
                            var v35 string = "foo bar"
                            _ = v34
                            _ = v35
                        }
                        _ = fi24
                    }
                    _ = fi16
                    _ = v17
                }
            }
            if true {
                if true {
                    v36 := GenS3{F0: -9223372036854775808, F1: -9223372036854775808, F2: 24}
                    var v37 int16 = 1
                    _ = v36
                    _ = v37
                }
            }
        } else {
            var v38 uint32 = (^(v2))
            _ = v38
        }
        var v39 int16 = -138
        _ = v39
    }
    var v40 int = v4
    var v41 int8 = 42
    v41 = (v41 & 0)
    switch ((((("" + "") + "a") + "") + "") + "a") {
    case "a":
        v12 = append(v12, (v4 * int(v2)))
    default:
        delete(v1, -708)
    }
    v42 := GenS3{F0: 42, F1: (^(v40)), F2: v41}
    for fi43 := 0; fi43 < 19; fi43++ {
        for fi44 := 0; fi44 < 7; fi44++ {
            var v45 float64 = (float64(v11) + (0.1 / (0.1 + float64(((v11 / v11) + v11)))))
            var v46 uint32 = v2
            if true {
                if (-128 != v41) {
                    fmt.Println(v41, v46, v46)
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                } else {
                    fmt.Println(v2, v45)
                    if (!true) {
                        v40 += 1
                        v47 := GenS1{F0: 9223372036854775807}
                        v12 = append(v12, -644)
                        _ = v47
                    }
                    v48 := GenS3{F0: 42, F1: (v40 >> ((2) & 63)), F2: (v41 | v41)}
                    _ = v48
                }
                if (uint64(v3) <= 417) {
                    var v49 uint = v3
                    _ = v49
                }
                v11 += v11
            }
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = fi43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
