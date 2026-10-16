package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
type GenS3 struct { F0 int64; F1 int; F2 int64 }
type GenS4 struct { F0 int64; F1 string; F2 uint }

func main() {
    var v1 float32 = 0.1
    var v2 uint8 = 252
    v3 := []string{("\n\t" + ""), "", "hello"}
    v4 := []int{0, 172}
    fmt.Println(v2, v2)
    fmt.Println(v2, v1)
    for fi5 := 0; fi5 < 17; fi5++ {
        if (uint64(v2) != 1) {
            for fi6 := 0; fi6 < 18; fi6++ {
                v7 := GenS1{F0: 0, F1: false}
                v8 := GenS1{F0: 0, F1: true}
                _ = fi6
                _ = v7
                _ = v8
            }
            v9 := GenS2{F0: 2, F1: -1, F2: 42}
            _ = v9
        } else {
            var v10 string = "\n\t"
            var v11 int32 = -2147483648
            if len(v4) > 0 {
                v4[0] = fi5
            }
            _ = v10
            _ = v11
        }
        v1 = v1
        var v12 string = ("\n\t" + ("\n\t" + "foo bar"))
        _ = fi5
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        if (!(!true)) {
            if (554 == 2) {
                if len(v4) > 0 {
                    v4[0] = (fi13 << ((4294967295) & 63))
                }
                v1 += v1
            }
            _ = len(v3)
        } else {
            if len(v4) > 0 {
                v4[0] = fi13
            }
        }
        fmt.Println(fi13)
        if (false && (0 >= -639)) {
            var v14 uint8 = v2
            v14 = (v2 << ((65535) & 63))
            _ = v14
        } else {
            v3 = append(v3, "foo bar")
        }
        _ = fi13
    }
    var v15 uint = 0
    for fi16 := 0; fi16 < 1; fi16++ {
        if false {
            v4 = append(v4, fi16)
        }
        if len(v3) > 0 {
            v3[0] = ("hello" + "a")
        }
        _ = fi16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint = (^(v15))
        fmt.Println(v18, fi17, v18)
        if (int64(v2) >= -779) {
            v3 = append(v3, "hello")
            var v19 uint8 = v2
            for fi20 := 0; fi20 < 1; fi20++ {
                fmt.Println(v19)
                _ = fi20
            }
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    if (197 == 128) {
        fmt.Println(v1, v15, v1)
        v3 = append(v3, ("a" + "hello"))
    }
    fmt.Println(v1)
    v15 += 18446744073709551615
    var v21 uint = 697
    if false {
        if (!(!true)) {
            var v22 float32 = (float32(((float64(v1) - float64((-2.25 - float32(1.5)))) / 0.1)) - 1.5)
            _ = v22
        }
        v23 := GenS2{F0: 54, F1: 291, F2: 100}
        for fi24 := 0; fi24 < 7; fi24++ {
            if ((-1 < -1) && ((v2 + 40) > v2)) {
                v1 += (3.14159 - (float32((-2.25 / float64(3.14159))) * ((v1 + (((v1 + v1) / float32(float64(v1))) * -2.25)) / v1)))
                v15 = v21
            } else {
                fmt.Println(v21, v21)
                v25 := GenS2{F0: 1, F1: 9223372036854775807, F2: 19}
                _ = v25
            }
            if true {
                if (true || false) {
                    fmt.Println(fi24, v2)
                } else {
                    for fi26 := 0; fi26 < 15; fi26++ {
                        switch 339 {
                        case 127:
                            switch (fi24 - int(v21)) {
                            case 4294967295:
                                v1 -= (v1 / (v1 * (1e10 / (v1 / float32(-2.25)))))
                            case -2147483648:
                                if false {
                                    v15 -= (v15 << (((v15 ^ 65535)) & 63))
                                    v27 := GenS2{F0: v15, F1: 1, F2: 0}
                                    if false {
                                        v28 := GenS3{F0: 881, F1: fi24, F2: -9223372036854775808}
                                        switch fi26 {
                                        case 0:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        default:
                                            v1 -= (float32(0.1) - 1e10)
                                        }
                                        _ = v28
                                    }
                                    _ = v27
                                }
                            case -963:
                                switch ("a" + "hello") {
                                case "foo bar":
                                    v15 -= (v15 & 4294967295)
                                case "":
                                    fmt.Println(v21)
                                default:
                                    var v29 int64 = 9223372036854775807
                                    _ = v29
                                }
                            }
                        case -2147483648:
                            var v30 float32 = -2.25
                            _ = v30
                        }
                        fmt.Println(v2)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi26
                    }
                    v1 = 1.5
                }
            } else {
                v15 = (v21 + (v21 - v15))
                for fi31 := 0; fi31 < 17; fi31++ {
                    v32 := GenS2{F0: (v21 >> (((v21 | (^(v21)))) & 63)), F1: -48, F2: -1}
                    _ = fi31
                    _ = v32
                }
                v21 = 4294967295
            }
            v2 += uint8(v21)
            _ = fi24
        }
        _ = v23
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "\n\t" {
        case "\n\t":
            var v33 int8 = -84
            _ = v33
        case "foo bar":
            switch "\n\t" {
            case "":
                var v34 uint16 = 1
                _ = v34
            default:
                var v35 int8 = 1
                _ = v35
            }
        default:
            fmt.Println(v1, v1, v2)
        }
        v36 := GenS4{F0: 579, F1: "a", F2: v15}
        _ = v36
    }
    v1 = v1
    var v37 uint32 = uint32(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v4
}
