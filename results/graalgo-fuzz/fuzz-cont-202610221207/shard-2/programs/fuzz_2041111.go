package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int8; F2 uint16 }
type GenS3 struct { F0 int8; F1 string }

func main() {
    var v1 uint16 = 942
    v2 := []int{-2147483648, 128, 217}
    var v3 uint8 = 255
    v4 := make(map[string]string)
    v5 := []int{42, -2147483648, -884}
    var v6 uint32 = 408
    fmt.Println(v3)
    v7 := v4[("a" + "hello")]
    switch 128 {
    case 1:
        fmt.Println(v1, v1, v1)
    default:
        if false {
            if (!true) {
                var v8 bool = (0.0 >= 0.0)
                var v9 int16 = -1
                var v10 uint32 = 65535
                _ = v10
                _ = v8
                _ = v9
            } else {
                v6 -= v6
                v6 = (v6 << ((65535) & 63))
            }
        }
    }
    v11 := v4[(("foo bar" + (v7 + ("" + ("\n\t" + v7)))) + v7)]
    delete(v4, v7)
    fmt.Println(v6, v6)
    v1 = v1
    var v12 uint32 = (v6 - 4294967295)
    v13 := GenS1{F0: -59}
    fmt.Println(v1, v12)
    for fi14 := 0; fi14 < 4; fi14++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            if false {
                fmt.Println(v3, fi15, v11)
                v6 += v6
                v16 := v4[("a" + "foo bar")]
                _ = v16
            }
            fmt.Println(v3, v6)
            v12 += (v6 << ((0) & 63))
            _ = fi15
        }
        switch fi14 {
        case 0:
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 uint64 = 640
                if false {
                    v6 += v12
                    var v19 int32 = 2147483647
                    switch (-(fi17)) {
                    case -2147483648:
                        var v20 float64 = (-2.25 * ((-2.25 * (1e10 - (-2.25 / 1.5))) - float64(1.5)))
                        _ = v20
                    case 42:
                        v1 = 1
                    }
                    _ = v19
                } else {
                    v5 = append(v5, (^(fi14)))
                    var v21 int = fi14
                    if false {
                        fmt.Println(v21)
                        v6 += 142
                        var v22 string = "foo bar"
                        _ = v22
                    } else {
                        v1 = v1
                        fmt.Println(v7)
                        fmt.Println(v12)
                    }
                    _ = v21
                }
                _ = fi17
                _ = v18
            }
        case 127:
            var v23 int = -242
            _ = v23
        case -2147483648:
            for fi24 := 0; fi24 < 11; fi24++ {
                if len(v5) > 0 {
                    v5[0] = (^(fi14))
                }
                switch (fi14 % (((fi14 ^ (fi14 >> ((0) & 63)))) | 1)) {
                case 1:
                    var v25 string = v11
                    if false {
                        fmt.Println(v12)
                        v6 = v12
                        v7 = ("" + v25)
                        if ((v3 << ((130) & 63)) == 0) {
                            switch fi24 {
                            case 128:
                                switch fi24 {
                                case -224:
                                    v7 = ((v7 + (((v25 + ("foo bar" + v25)) + v7) + "\n\t")) + (("a" + (v25 + "a")) + v7))
                                }
                            case 127:
                                fmt.Println(v12)
                            case 42:
                                switch fi24 {
                                case -1:
                                    v26 := GenS1{F0: fi14}
                                    _ = v26
                                case 9223372036854775807:
                                    v27 := GenS1{F0: fi24}
                                    _ = v27
                                }
                            default:
                                fmt.Println(v12, fi14)
                            }
                            v1 -= v1
                            var v28 int16 = int16(fi24)
                            _ = v28
                        }
                        v29 := GenS1{F0: -104}
                        _ = v29
                    } else {
                        fmt.Println(v11, v6, fi24)
                    }
                    _ = v25
                case 0:
                    v12 -= v12
                }
                _ = fi24
            }
        }
        _ = fi14
    }
    if false {
        v30 := GenS2{F0: ((((v7 + (v7 + v11)) + "foo bar") + v7) + v11), F1: -1, F2: v1}
        v31 := GenS2{F0: v11, F1: 0, F2: v1}
        v11 = v11
        _ = v30
        _ = v31
    } else {
        if (4294967295 <= v6) {
            v32 := GenS3{F0: 112, F1: (v11 + v11)}
            switch v7 {
            case "\n\t":
                v33 := GenS1{F0: -1}
                _ = v33
            case "a":
                if (735 > -26) {
                    for fi34 := 0; fi34 < 14; fi34++ {
                        var v35 float32 = 1.5
                        _ = fi34
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 16; fi36++ {
                        _ = len(v4)
                        for fi37 := 0; fi37 < 11; fi37++ {
                            v38 := v4[v11]
                            fmt.Println(fi37, fi36, v7)
                            _ = len(v5)
                            var v39 uint16 = v1
                            if len(v5) > 0 {
                                v5[0] = 1
                            }
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        _ = fi36
                    }
                    if (1e10 < float64((((0.1 * -2.25) + 3.14159) - (1.5 * 0.1)))) {
                        v11 = ""
                    }
                } else {
                    if (false || (!(156 < 0))) {
                        v3 -= v3
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi40 := 0; fi40 < 14; fi40++ {
                            v41 := v4["a"]
                            _ = fi40
                            _ = v41
                        }
                    }
                    v42 := GenS2{F0: (((v11 + v11) + v11) + v11), F1: 42, F2: (v1 | (v1 >> ((65535) & 63)))}
                    _ = v42
                }
            }
            _ = v32
        } else {
            v3 -= (v3 % ((v3) | 1))
            v12 = 1
            v43 := v4[("a" + v11)]
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
