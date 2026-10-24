package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }
type GenS2 struct { F0 string; F1 int16; F2 bool }
type GenS3 struct { F0 uint64 }
func genFn0() uint8 {
    v7 := []int{4294967295, -9223372036854775808, 42}
    if len(v7) > 0 {
        v7[0] = -9223372036854775808
    }
    var v8 int32 = -186
    fmt.Println(v8)
    for fi9 := 0; fi9 < 15; fi9++ {
        v8 -= v8
        _ = fi9
    }
    _ = v7
    _ = v8
    return 176
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 float64 = 0.0
    var v5 bool = ((130 > 127) || (3.14159 >= ((float32(float64(float32(1e10))) + 1.5) - -2.25)))
    var v6 bool = v5
    var v10 uint16 = 32767
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            fmt.Println(fi12)
            v2[(fi11 & (fi11 - 652))] = "\n\t"
            _ = fi12
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            if false {
                v14 := GenS1{F0: "foo bar", F1: 1, F2: 1000}
                v5 = (-1 >= -632)
                if true {
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v6 = false
                        v15 := GenS2{F0: "foo bar", F1: -1, F2: (18446744073709551615 >= 0)}
                        if (!(uint32(v10) == 63)) {
                            v4 -= v4
                        } else {
                            v10 = (^(v10))
                        }
                        _ = v15
                    case "":
                        v5 = (42 >= 42)
                    }
                    if (fi11 < fi13) {
                        v4 = (1.5 * ((0.0 + (v4 - (v4 - v4))) + 1e10))
                        v16 := GenS2{F0: ("foo bar" + "\n\t"), F1: 410, F2: v5}
                        switch fi11 {
                        case -660:
                            v17 := GenS2{F0: ("a" + ("\n\t" + "")), F1: 581, F2: (v4 >= ((0.0 - v4) / (3.14159 / (v4 / (1e10 - (3.14159 - v4))))))}
                            _ = v17
                        case -9223372036854775808:
                            var v18 string = "hello"
                            _ = v18
                        case 705:
                            v6 = v5
                            if true {
                                switch fi11 {
                                case 9223372036854775807:
                                    switch "" {
                                    case "foo bar":
                                        var v19 uint16 = v10
                                        _ = v19
                                    }
                                case 127:
                                    var v20 int32 = 0
                                    if len(v1) > 0 {
                                        v1[0] = 65535
                                    }
                                    var v21 int64 = -1
                                    _ = v20
                                    _ = v21
                                case 2147483647:
                                    v22 := GenS2{F0: ("foo bar" + "hello"), F1: -32768, F2: (v10 > 65535)}
                                    _ = v22
                                default:
                                    v4 = v4
                                }
                                v5 = v6
                            }
                        }
                        if v6 {
                            switch (fi11 << ((481) & 63)) {
                            case -2147483648:
                                var v23 int32 = 0
                                _ = v23
                            case 1:
                                v4 -= v4
                            }
                            var v24 int64 = -953
                            _ = v24
                        }
                        var v25 int8 = 42
                        _ = v16
                        _ = v25
                    }
                    var v26 bool = (42 > 2147483648)
                    _ = v26
                } else {
                    var v27 int32 = 85
                    _ = v27
                }
                _ = v14
            } else {
                v6 = v6
                var v28 string = "hello"
                _ = v28
            }
            _ = fi13
        }
        for fi29 := 0; fi29 < 3; fi29++ {
            var v30 int16 = 0
            _ = fi29
            _ = v30
        }
        _ = fi11
    }
    v31 := GenS1{F0: "a", F1: 1, F2: 42}
    var v32 uint32 = 42
    var v33 int16 = 32767
    if true {
        if (!(!(65535 > 255))) {
            v34 := GenS2{F0: "hello", F1: (v33 / ((v33) | 1)), F2: ((!(v5 && true)) && (true || false))}
            _ = v34
        } else {
            if len(v1) > 0 {
                v1[0] = (v32 + v32)
            }
            v10 -= (^(v10))
            var v35 uint64 = 0
            _ = v35
        }
        switch "\n\t" {
        case "":
            fmt.Println(v32)
        case "foo bar":
            v36 := v2[2147483647]
            _ = v36
        case "hello":
            v10 = (v10 / ((42) | 1))
        default:
            fmt.Println(v32)
        }
    } else {
        for fi37 := 0; fi37 < 18; fi37++ {
            if v6 {
                v4 = ((v4 / (v4 + v4)) * ((((v4 - v4) / -2.25) - (v4 + 3.14159)) * v4))
                var v38 float32 = float32(v4)
                v32 += 42
                _ = v38
            }
            delete(v3, fi37)
            v5 = false
            _ = fi37
        }
        var v39 bool = (18446744073709551615 == 42)
        _ = v39
    }
    fmt.Println(v4, v10, v33)
    for fi40 := 0; fi40 < 9; fi40++ {
        if (false && ((v4 - v4) < v4)) {
            fmt.Println(v32, fi40, fi40)
        }
        fmt.Println(v4)
        _ = fi40
    }
    var v41 uint = 1
    v42 := GenS2{F0: "foo bar", F1: v33, F2: (255 != 0)}
    for fi43 := 0; fi43 < 2; fi43++ {
        v41 = (^(v41))
        v41 = 255
        switch "foo bar" {
        case "a":
            for fi44 := 0; fi44 < 10; fi44++ {
                v32 = v32
                _ = fi44
            }
        case "":
            if (v5 || (-30 < -1)) {
                var v45 int32 = 1
                v41 -= (v41 >> ((4294967295) & 63))
                switch "hello" {
                case "\n\t":
                    fmt.Println(v41, v33, v45)
                case "foo bar":
                    if ((!((((("foo bar" + "hello") == "foo bar") || (((fi43 >> (((v41 % (((v41 * v41)) | 1))) & 63)) == fi43) || (float64(((float32(v4) * 1.5) * 3.14159)) < v4))) && true) || (!true))) || v5) {
                        v46 := GenS2{F0: "a", F1: (v33 & v33), F2: v5}
                        v33 = v33
                        _ = v46
                    } else {
                        var v47 uint32 = v32
                        _ = len(v1)
                        v48 := GenS3{F0: 1}
                        _ = v47
                        _ = v48
                    }
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = v45
            } else {
                var v49 uint64 = 2
                _ = v49
            }
        case "foo bar":
            if (false || v6) {
                if ((((v4 * float64(float32(v4))) + v4) * (-2.25 + v4)) < (0.1 / 1.5)) {
                    _ = len(v3)
                    v6 = (fi43 >= fi43)
                    v4 += ((float64(0.0) + 1.5) - float64(0.0))
                } else {
                    var v50 float64 = 1.5
                    _ = v50
                }
            } else {
                var v51 int16 = (v33 % ((v33) | 1))
                v10 += 0
                _ = v51
            }
        }
        _ = fi43
    }
    _ = len(v3)
    fmt.Println(v33, v4, v41)
    v5 = (!(!v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
