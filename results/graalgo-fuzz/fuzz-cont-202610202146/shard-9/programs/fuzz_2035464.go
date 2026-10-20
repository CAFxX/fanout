package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }

func main() {
    v1 := []uint32{981, 4294967295}
    v2 := []int{2147483647}
    var v3 uint16 = 67
    var v4 int64 = 0
    v5 := make(map[string]int)
    switch 2147483647 {
    case 2147483647:
        if true {
            var v6 int8 = -1
            _ = v6
        }
    case 128:
        v7 := GenS1{F0: 65535, F1: 804, F2: 4294967295}
        _ = v7
    default:
        switch ("foo bar" + "hello") {
        case "":
            var v8 uint16 = (^(v3))
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := GenS1{F0: 0, F1: 135, F2: 998}
                switch -9223372036854775808 {
                case 127:
                    v11 := v5["hello"]
                    _ = v11
                case 42:
                    if (!((true && (("foo bar" + "a") <= "")) || (true && (255 >= 127)))) {
                        v3 -= (v3 * (^(v3)))
                        fmt.Println(v3, v4, v4)
                        var v12 uint32 = uint32(v3)
                        _ = v12
                    }
                default:
                    v3 -= v3
                }
                _ = fi9
                _ = v10
            }
        case "hello":
            v3 -= (^(v3))
        default:
            v3 += v3
        }
    }
    if ((0.1 - (3.14159 - -2.25)) <= float64((1.5 * 0.0))) {
        var v13 uint16 = v3
        switch 0 {
        case -9223372036854775808:
            switch -188 {
            case -666:
                v13 = v3
            default:
                for fi14 := 0; fi14 < 18; fi14++ {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v16 := v5[((("\n\t" + "a") + "") + "")]
                        var v17 uint16 = v3
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 16; fi18++ {
                        switch fi14 {
                        case -9223372036854775808:
                            fmt.Println(v4)
                        default:
                            var v19 uint16 = v3
                            _ = v19
                        }
                        _ = fi18
                    }
                    _ = fi14
                }
            }
        case 4294967295:
            v20 := GenS2{F0: ("" == ("" + "\n\t"))}
            _ = v20
        case 127:
            v3 += v3
        default:
            if (1.5 <= float32(1e10)) {
                var v21 int8 = 116
                _ = len(v1)
                _ = v21
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v13 = (v3 | (v3 * (^(v13))))
                for fi22 := 0; fi22 < 13; fi22++ {
                    fmt.Println(v4, v13)
                    if (0 != 0) {
                        fmt.Println(fi22)
                        switch "foo bar" {
                        case "\n\t":
                            v23 := GenS2{F0: true}
                            _ = v23
                        case "foo bar":
                            v3 += v13
                        }
                    }
                    if false {
                        if (161 < 255) {
                            v24 := GenS2{F0: true}
                            switch (^(fi22)) {
                            case 128:
                                v13 += v13
                            case -2147483648:
                                v25 := GenS3{F0: 65535}
                                _ = v25
                            case -1:
                                for fi26 := 0; fi26 < 11; fi26++ {
                                    v4 += v4
                                    v4 = (v4 << ((255) & 63))
                                    var v27 bool = true
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v28 float32 = -2.25
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                            }
                            _ = v24
                        }
                        v29 := GenS3{F0: 255}
                        v30 := GenS2{F0: false}
                        _ = v29
                        _ = v30
                    }
                    _ = fi22
                }
            }
        }
        _ = v13
    }
    var v31 int32 = -2147483648
    v32 := GenS2{F0: true}
    v31 = (v31 << ((4294967295) & 63))
    if false {
        v31 += v31
        for fi33 := 0; fi33 < 7; fi33++ {
            var v34 float64 = (1e10 / 1.0)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        v36 := GenS1{F0: 42, F1: uint(v3), F2: 922}
        v37 := v5[("" + "")]
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 int8 = 0
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v4
    _ = v5
}
