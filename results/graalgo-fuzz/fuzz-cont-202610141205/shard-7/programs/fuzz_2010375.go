package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int16 }

func main() {
    var v1 float32 = ((3.14159 + 3.14159) / 1.0)
    v2 := []int{32768, 32768}
    v3 := []string{"", "foo bar", ("foo bar" + "hello")}
    var v4 int32 = -2147483648
    var v5 float32 = (1.5 + v1)
    v6 := GenS1{F0: 123}
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int16 = -1
        v3 = append(v3, "foo bar")
        for fi9 := 0; fi9 < 3; fi9++ {
            if (42 == 42) {
                if len(v3) > 0 {
                    v3[0] = ("hello" + "a")
                }
            } else {
                var v10 int64 = 42
                var v11 int = (fi9 - (fi7 | -2147483648))
                if false {
                    v4 -= v4
                    if (-1 == 42) {
                        if false {
                            fmt.Println(v1)
                        } else {
                            fmt.Println(fi7, v8)
                        }
                        fmt.Println(v11)
                        switch v11 {
                        case 0:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 127:
                            v5 -= v5
                        case -553:
                            switch "foo bar" {
                            case "a":
                                if true {
                                    fmt.Println(v1, v10)
                                    v1 = v5
                                    var v12 float64 = ((((3.14159 / (1e10 - ((float64(v1) / -2.25) / 3.14159))) * 0.1) + 0.1) + (float64(-2.25) - 0.1))
                                    if true {
                                        _ = len(v3)
                                        var v13 uint32 = 387
                                        if ((v13 == (v13 & v13)) || true) {
                                            v3 = append(v3, "a")
                                            v12 -= (-2.25 + (v12 * v12))
                                            var v14 int16 = v8
                                            _ = v14
                                        }
                                        _ = v13
                                    }
                                    _ = v12
                                }
                            case "hello":
                                if (1e10 >= float64(v1)) {
                                    v10 -= -774
                                }
                            }
                        default:
                            if ((-(v4)) >= (v4 + -2147483648)) {
                                v15 := GenS2{F0: 859, F1: 383}
                                switch int(v10) {
                                case 765:
                                    _ = len(v3)
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                    switch (fi9 << ((963) & 63)) {
                                    case 2147483647:
                                        var v16 bool = (int(v4) < 401)
                                        _ = v16
                                    case 128:
                                        if (!false) {
                                            if (255 == 2) {
                                                v5 += 3.14159
                                                v17 := GenS3{F0: 42, F1: ""}
                                                _ = v17
                                            }
                                        } else {
                                            var v18 float64 = 1.5
                                            _ = v18
                                        }
                                    case 4294967295:
                                        switch (v11 - fi7) {
                                        case 4294967295:
                                            v8 += 1
                                        case 2147483647:
                                            var v19 uint8 = 1
                                            _ = v19
                                        default:
                                            v5 = ((v1 + v5) / 0.1)
                                        }
                                    }
                                }
                                v10 -= (v10 ^ v10)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v15
                            } else {
                                v5 = v5
                            }
                        }
                    }
                } else {
                    _ = len(v3)
                }
                _ = v10
                _ = v11
            }
            if (((false && ((fi9 == fi7) || ((2 <= 1) || true))) || (fi7 != -614)) && false) {
                if false {
                    if false {
                        if true {
                            v2 = append(v2, fi7)
                            if len(v2) > 0 {
                                v2[0] = (fi9 & (fi7 + fi7))
                            }
                            var v20 uint8 = 0
                            var v21 float32 = float32(((float64((-2.25 + v1)) * 1e10) + (1.5 * (0.1 * (-2.25 * -2.25)))))
                            _ = v20
                            _ = v21
                        } else {
                            fmt.Println(v1)
                            _ = len(v3)
                            v22 := GenS4{F0: (v8 << ((429) & 63))}
                            _ = v22
                        }
                    } else {
                        v23 := GenS1{F0: 0}
                        _ = v23
                    }
                    v4 = v4
                } else {
                    var v24 string = ""
                    fmt.Println(v24)
                    _ = v24
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = ""
                }
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    if true {
        if true {
            for fi25 := 0; fi25 < 18; fi25++ {
                for fi26 := 0; fi26 < 19; fi26++ {
                    v4 = (v4 | -2147483648)
                    if (true && true) {
                        v5 = -2.25
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = len(v3)
                        v4 = v4
                        v3 = append(v3, "a")
                    }
                    v1 -= (1.5 * v5)
                    _ = fi26
                }
                _ = fi25
            }
        } else {
            var v27 float64 = 0.1
            _ = v27
        }
        var v28 uint32 = 0
        _ = v28
    }
    var v29 int8 = int8(v4)
    switch 313 {
    case 127:
        v1 = float32(float64(-2.25))
    case 4294967295:
        if (230 >= 127) {
            v30 := GenS1{F0: (v29 ^ (v29 % ((v29) | 1)))}
            var v31 uint64 = 0
            fmt.Println(v31, v5)
            _ = v30
            _ = v31
        }
    }
    var v32 float32 = v5
    var v33 int16 = -299
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v34 := GenS5{F0: 0, F1: (v33 % ((v33) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
