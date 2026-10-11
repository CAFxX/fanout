package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []uint32{274}
    v2 := []int{487, -6, -2147483648, -2147483648}
    var v3 int = 649
    v4 := []string{"a", ("" + "foo bar"), (("" + "\n\t") + "a")}
    v5 := make(map[string]uint32)
    v6 := v5[(("a" + ("foo bar" + "\n\t")) + ("a" + "hello"))]
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 70, F1: 209}
        for fi9 := 0; fi9 < 16; fi9++ {
            _ = len(v5)
            if ((fi9 % ((32768) | 1)) != (^(fi7))) {
                if false {
                    v6 -= v6
                } else {
                    var v10 uint32 = v6
                    _ = v10
                }
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        switch (^(fi11)) {
        case -9223372036854775808:
            if (!false) {
                if (uint16(v6) < 32767) {
                    v6 = (v6 / (((v6 & v6)) | 1))
                    switch fi11 {
                    case 627:
                        if false {
                            var v12 uint16 = 585
                            var v13 uint = 4294967295
                            for fi14 := 0; fi14 < 5; fi14++ {
                                fmt.Println(v12, v3)
                                _ = fi14
                            }
                            _ = v12
                            _ = v13
                        } else {
                            for fi15 := 0; fi15 < 13; fi15++ {
                                v16 := GenS1{F0: 0, F1: 206}
                                switch "a" {
                                case "foo bar":
                                    var v17 bool = (42 > 286)
                                    switch "" {
                                    case "a":
                                        fmt.Println(v6, v3)
                                    case "foo bar":
                                        var v18 bool = v17
                                        _ = v18
                                    default:
                                        v17 = v17
                                    }
                                    var v19 float32 = 0.0
                                    var v20 uint = 568
                                    switch -274 {
                                    case 858:
                                        switch "\n\t" {
                                        case "a":
                                            fmt.Println(fi15, v19, v17)
                                        case "":
                                            v20 += v20
                                        case "foo bar":
                                            var v21 uint16 = 65535
                                            _ = v21
                                        default:
                                            v3 -= fi11
                                        }
                                    case 1:
                                        v19 = 0.0
                                    default:
                                        v20 += (^(v20))
                                    }
                                    _ = v17
                                    _ = v19
                                    _ = v20
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                v6 += (v6 / ((42) | 1))
                                _ = fi15
                                _ = v16
                            }
                            v6 += v6
                        }
                    case 127:
                        switch v3 {
                        case 4294967295:
                            v3 = 9223372036854775807
                        case 32768:
                            var v22 float64 = 1.5
                            _ = v22
                        default:
                            fmt.Println(v6)
                        }
                    default:
                        var v23 uint8 = 242
                        _ = v23
                    }
                }
            } else {
                var v24 int = -9223372036854775808
                switch (-(fi11)) {
                case 127:
                    for fi25 := 0; fi25 < 11; fi25++ {
                        if false {
                            switch "" {
                            case "foo bar":
                                v26 := GenS2{F0: -629}
                                _ = v26
                            case "\n\t":
                                fmt.Println(fi25)
                            }
                            if (18446744073709551615 < 255) {
                                var v27 float64 = 0.0
                                var v28 int = (fi11 * fi11)
                                fmt.Println(v28, v24)
                                _ = v27
                                _ = v28
                            }
                        }
                        fmt.Println(v3, v6)
                        _ = fi25
                    }
                case 32767:
                    var v29 int16 = 1000
                    _ = v29
                case -1:
                    var v30 int64 = 1
                    _ = v30
                }
                _ = v24
            }
        case 32768:
            fmt.Println(fi11)
        default:
            var v31 uint32 = 172
            _ = v31
        }
        v32 := GenS1{F0: 471, F1: 65535}
        var v33 float64 = 1e10
        _ = fi11
        _ = v32
        _ = v33
    }
    v34 := GenS1{F0: 648, F1: 0}
    var v35 bool = false
    var v36 int8 = 68
    for fi37 := 0; fi37 < 3; fi37++ {
        v36 += (v36 >> ((65535) & 63))
        if len(v2) > 0 {
            v2[0] = v3
        }
        v35 = false
        _ = fi37
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        v5["hello"] = (v6 << ((65535) & 63))
        var v39 bool = ((42 >= 471) && (true && (65535 > 65535)))
        var v40 uint32 = (^(v6))
        _ = fi38
        _ = v39
        _ = v40
    }
    _ = len(v4)
    if (v35 || (4294967295 != v6)) {
        v6 += v6
        var v41 bool = (((v36 - (v36 ^ v36)) != 100) || (true && (-687 == 0)))
        _ = v41
    }
    if (0 < 103) {
        fmt.Println(v3, v6, v36)
        var v42 bool = false
        v35 = (25 > uint8(v6))
        _ = v42
    }
    v6 = (^(v6))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
