package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint64; F2 string }
type GenS3 struct { F0 string; F1 uint8; F2 int }

func main() {
    var v1 uint8 = 128
    var v2 int = 42
    var v3 uint8 = (^(v1))
    var v4 float32 = ((-2.25 - (1e10 - 0.1)) + 3.14159)
    var v5 int64 = 1
    switch 4294967295 {
    case 2147483647:
        v6 := GenS1{F0: 152}
        _ = v6
    case 4294967295:
        for fi7 := 0; fi7 < 11; fi7++ {
            switch "" {
            case "foo bar":
                var v8 uint = 18446744073709551615
                _ = v8
            case "hello":
                v9 := GenS2{F0: v2, F1: 471, F2: "hello"}
                _ = v9
            default:
                v10 := []int{}
                _ = v10
            }
            _ = fi7
        }
    case 128:
        if (true || (1 < v5)) {
            if (-203 < -2147483648) {
                var v11 uint32 = 373
                _ = v11
            }
            fmt.Println(v5)
        } else {
            if (42 != 76) {
                switch "" {
                case "foo bar":
                    var v12 int16 = -1
                    _ = v12
                default:
                    v13 := GenS2{F0: v2, F1: 741, F2: ""}
                    _ = v13
                }
                v14 := []int{(v2 ^ 32768), (v2 - int(v5)), (-(v2)), v2}
                if true {
                    var v15 uint16 = 32767
                    v16 := GenS1{F0: 255}
                    _ = len(v14)
                    _ = v15
                    _ = v16
                } else {
                    if ((v4 < v4) || (false && (65535 != 618))) {
                        var v17 uint64 = 2
                        fmt.Println(v17, v4, v3)
                        v17 -= (v17 >> ((4294967295) & 63))
                        _ = v17
                    }
                    v18 := GenS2{F0: (v2 << ((683) & 63)), F1: 42, F2: "\n\t"}
                    if false {
                        var v19 int64 = v5
                        v19 += (v5 % ((v19) | 1))
                        _ = v19
                    }
                    _ = v18
                }
                _ = v14
            } else {
                switch "a" {
                case "hello":
                    v20 := []int{(v2 >> ((111) & 63)), v2, v2}
                    _ = v20
                default:
                    var v21 bool = (((3.14159 * v4) >= float32(3.14159)) && (18446744073709551615 == 227))
                    _ = v21
                }
            }
            fmt.Println(v1, v3)
        }
    default:
        var v22 uint8 = (^(v1))
        _ = v22
    }
    v23 := []int{(-(v2)), (v2 << ((324) & 63))}
    if len(v23) > 0 {
        v23[0] = (v2 << ((615) & 63))
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        v4 = (v4 / v4)
        _ = len(v23)
        v3 += v1
        _ = fi24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := GenS3{F0: "\n\t", F1: (v3 - (v3 + 128)), F2: 32767}
        for fi27 := 0; fi27 < 12; fi27++ {
            v3 -= 128
            if ("\n\t" < "hello") {
                v5 -= v5
                fmt.Println(fi27, v1, v5)
            } else {
                var v28 int64 = v5
                if len(v23) > 0 {
                    v23[0] = (-(v2))
                }
                _ = v28
            }
            _ = fi27
        }
        if true {
            fmt.Println(v2)
            switch (fi25 % ((v2) | 1)) {
            case 1:
                var v29 string = "hello"
                _ = v29
            case -1:
                for fi30 := 0; fi30 < 2; fi30++ {
                    switch (("a" + "\n\t") + ("foo bar" + "\n\t")) {
                    case "":
                        fmt.Println(v3, v3, fi25)
                    }
                    _ = fi30
                }
            default:
                fmt.Println(v2, v1)
            }
            for fi31 := 0; fi31 < 4; fi31++ {
                if false {
                    if true {
                        fmt.Println(fi25, v3)
                    } else {
                        switch "\n\t" {
                        case "a":
                            var v32 float64 = 0.1
                            _ = v32
                        case "foo bar":
                            v2 = (^(v2))
                        case "":
                            fmt.Println(v5, v5)
                        }
                        if false {
                            fmt.Println(fi31, fi25, v5)
                        } else {
                            v2 = fi31
                            _ = len(v23)
                        }
                        if true {
                            if false {
                                if true {
                                    v3 = v1
                                } else {
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    v3 = v1
                                    var v33 uint64 = 18446744073709551615
                                    v23 = append(v23, int(v3))
                                    _ = len(v23)
                                    _ = v33
                                }
                                if (-2147483648 > 762) {
                                    if len(v23) > 0 {
                                        v23[0] = -578
                                    }
                                }
                                var v34 int = (fi31 ^ 127)
                                _ = v34
                            } else {
                                var v35 int = (^(fi31))
                                var v36 float64 = 1.5
                                switch (fi31 & fi25) {
                                case -2147483648:
                                    var v37 int32 = 42
                                    _ = v37
                                case -1:
                                    v1 = v3
                                case 917:
                                    v38 := GenS2{F0: fi25, F1: 0, F2: ""}
                                    _ = v38
                                }
                                _ = v35
                                _ = v36
                            }
                            v5 = int64(v1)
                        } else {
                            v3 = 128
                            var v39 int = 9223372036854775807
                            v23 = append(v23, (fi25 % (((v39 + (v39 >> ((1) & 63)))) | 1)))
                            _ = v39
                        }
                    }
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v3)
                    }
                } else {
                    switch v2 {
                    case 4294967295:
                        var v40 uint32 = 0
                        _ = v40
                    case -980:
                        var v41 bool = (42 == 315)
                        _ = v41
                    case 94:
                        var v42 float64 = ((1e10 + 1e10) * (1e10 + 0.1))
                        _ = v42
                    default:
                        fmt.Println(v5, v1, v2)
                    }
                }
                v43 := GenS2{F0: v2, F1: 492, F2: ("\n\t" + "a")}
                _ = fi31
                _ = v43
            }
        }
        _ = fi25
        _ = v26
    }
    v44 := GenS1{F0: 966}
    if ((339 == 32767) && (0.1 > float32(((-2.25 * -2.25) * 1.5)))) {
        var v45 uint64 = 374
        _ = v45
    } else {
        var v46 uint8 = v1
        v2 = v2
        if true {
            _ = len(v23)
            v1 = (^(v3))
        }
        _ = v46
    }
    v1 -= 127
    v5 = v5
    v1 -= v1
    v1 = v3
    v4 -= ((v4 / 1.5) / v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
