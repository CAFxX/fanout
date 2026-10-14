package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[string]int)
    var v3 string = "a"
    v4 := make(map[int]string)
    var v5 int64 = 42
    v6 := []uint32{65535, 817}
    var v7 uint32 = 212
    if false {
        fmt.Println(v5, v3)
    } else {
        var v8 float32 = (3.14159 - 0.0)
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(v8, v7)
            _ = fi9
        }
        _ = v8
    }
    if false {
        if len(v6) > 0 {
            v6[0] = v7
        }
        var v10 int16 = 1
        _ = v10
    }
    fmt.Println(v5, v5, v5)
    v4[v1] = ("\n\t" + v3)
    v11 := GenS1{F0: 1, F1: (1e10 < (1e10 + 0.1)), F2: 42}
    if true {
        for fi12 := 0; fi12 < 11; fi12++ {
            switch v3 {
            case "":
                fmt.Println(v3, v7)
            case "foo bar":
                v3 = v3
            }
            v13 := v4[(-(fi12))]
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 13; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                fmt.Println(fi15)
                if (v3 < v3) {
                    v3 = v3
                }
                v16 := v2["hello"]
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    v17 := GenS2{F0: (-(v5))}
    _ = len(v6)
    v6 = append(v6, (v7 << ((42) & 63)))
    fmt.Println(v1, v3)
    var v18 string = (((v3 + v3) + v3) + "a")
    var v19 int16 = -1
    switch (v18 + "a") {
    case "a":
        v20 := GenS2{F0: v5}
        _ = v20
    case "\n\t":
        fmt.Println(v5)
    default:
        v7 += v7
    }
    switch (v1 & v1) {
    case 40:
        var v21 int64 = v5
        _ = v21
    case 32767:
        v7 = (v7 << ((2) & 63))
    default:
        if (false && false) {
            if true {
                for fi22 := 0; fi22 < 6; fi22++ {
                    if (4294967295 == 4294967295) {
                        if false {
                            v23 := GenS1{F0: 42, F1: false, F2: 244}
                            var v24 int = 4294967295
                            _ = v23
                            _ = v24
                        } else {
                            v3 = v3
                            switch ((v18 + ("\n\t" + "hello")) + (v18 + v18)) {
                            case "a":
                                fmt.Println(v1)
                            case "foo bar":
                                switch v18 {
                                case "foo bar":
                                    var v25 float64 = (((0.0 / -2.25) * 1e10) * (0.0 - (1e10 - 1.5)))
                                    _ = v25
                                case "":
                                    for fi26 := 0; fi26 < 11; fi26++ {
                                        var v27 int32 = 899
                                        var v28 int32 = v27
                                        _ = fi26
                                        _ = v27
                                        _ = v28
                                    }
                                default:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                            default:
                                switch (v1 / ((fi22) | 1)) {
                                case -626:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                default:
                                    v5 += (^(v5))
                                }
                            }
                            if len(v6) > 0 {
                                v6[0] = 674
                            }
                        }
                    } else {
                        _ = len(v4)
                        var v29 int8 = 0
                        fmt.Println(v3, v29)
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 7; fi30++ {
                        var v31 uint64 = 971
                        v32 := GenS1{F0: 1, F1: (179 <= 1), F2: 168}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS1{F0: -2147483648, F1: false, F2: 128}
                    _ = fi22
                    _ = v33
                }
                var v34 uint32 = 0
                fmt.Println(v1)
                _ = v34
            }
            switch -805 {
            case -2147483648:
                v18 = v3
            case 128:
                if (746 == 32767) {
                    fmt.Println(v5)
                    for fi35 := 0; fi35 < 7; fi35++ {
                        var v36 string = ((v18 + v18) + v18)
                        _ = fi35
                        _ = v36
                    }
                } else {
                    var v37 int = 128
                    fmt.Println(v18, v3, v3)
                    _ = v37
                }
            }
            for fi38 := 0; fi38 < 18; fi38++ {
                if true {
                    v1 = 127
                    if (-9223372036854775808 != (fi38 % ((v1) | 1))) {
                        var v39 string = ("a" + v3)
                        switch (fi38 | (v1 - (fi38 | 32768))) {
                        case -1:
                            if true {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                switch ("foo bar" + v18) {
                                case "\n\t":
                                    fmt.Println(v3, v39, v18)
                                case "a":
                                    for fi40 := 0; fi40 < 14; fi40++ {
                                        v41 := GenS3{F0: int32(v19)}
                                        if len(v6) > 0 {
                                            v6[0] = (^(v7))
                                        }
                                        _ = fi40
                                        _ = v41
                                    }
                                default:
                                    fmt.Println(v1, fi38)
                                }
                                var v42 string = v18
                                _ = v42
                            } else {
                                var v43 int8 = -1
                                fmt.Println(v18, v39)
                                if len(v6) > 0 {
                                    v6[0] = 4294967295
                                }
                                _ = v43
                            }
                        case 128:
                            for fi44 := 0; fi44 < 19; fi44++ {
                                if false {
                                    v39 = v39
                                    var v45 uint32 = v7
                                    _ = v45
                                }
                                var v46 uint32 = (v7 / (((v7 * 2)) | 1))
                                _ = fi44
                                _ = v46
                            }
                        }
                        _ = v39
                    }
                    var v47 uint32 = 0
                    _ = v47
                } else {
                    var v48 uint8 = 0
                    _ = v48
                }
                fmt.Println(v19)
                var v49 uint = 65535
                _ = fi38
                _ = v49
            }
        }
    }
    var v50 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
