package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int32; F1 int32 }

func main() {
    var v1 int64 = 594
    var v2 uint32 = 65535
    var v3 bool = (-1 == int16(v1))
    v4 := []uint32{941, (^(v2))}
    var v5 int8 = -128
    v6 := make(map[uint32]int)
    if true {
        var v7 uint16 = 242
        _ = v7
    } else {
        var v8 int8 = (-(v5))
        var v9 float64 = float64(0.0)
        _ = v8
        _ = v9
    }
    switch "" {
    case "\n\t":
        v2 = 0
    case "hello":
        var v10 string = (("" + "hello") + "hello")
        _ = v10
    default:
        for fi11 := 0; fi11 < 7; fi11++ {
            switch (fi11 >> ((0) & 63)) {
            case 127:
                fmt.Println(v1, v1, v5)
            case -9223372036854775808:
                v1 += (v1 - v1)
            case 32767:
                var v12 int64 = v1
                _ = v12
            default:
                v5 = v5
            }
            if (!false) {
                var v13 int64 = -9223372036854775808
                _ = v13
            } else {
                if v3 {
                    fmt.Println(v5, v2)
                } else {
                    v1 = (-(v1))
                }
                var v14 uint32 = (v2 / ((2) | 1))
                _ = v14
            }
            var v15 uint = 659
            _ = fi11
            _ = v15
        }
    }
    if len(v4) > 0 {
        v4[0] = v2
    }
    var v16 float32 = 0.0
    v17 := v6[(v2 >> ((4294967295) & 63))]
    var v18 uint32 = v2
    v3 = false
    v4 = append(v4, (v18 >> ((138) & 63)))
    var v19 uint = 965
    v18 = (^(v2))
    for fi20 := 0; fi20 < 20; fi20++ {
        if v3 {
            v2 += 42
        } else {
            for fi21 := 0; fi21 < 18; fi21++ {
                v22 := GenS1{F0: 1, F1: (v1 % (((^(v1))) | 1))}
                var v23 int8 = 6
                v4 = append(v4, (v18 % ((v18) | 1)))
                var v24 uint64 = 701
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            var v25 bool = v3
            var v26 uint16 = 0
            _ = v25
            _ = v26
        }
        _ = fi20
    }
    if (3.14159 <= ((0.1 + 0.0) - -2.25)) {
        for fi27 := 0; fi27 < 17; fi27++ {
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 string = (("a" + ("foo bar" + "a")) + ((("" + "a") + ("foo bar" + (("foo bar" + "") + ""))) + "foo bar"))
                switch fi27 {
                case 127:
                    v6[(v18 / ((2) | 1))] = (^(fi27))
                }
                switch (((((v29 + v29) + (v29 + (("foo bar" + "a") + "\n\t"))) + "\n\t") + (v29 + v29)) + "hello") {
                case "a":
                    v2 = (v18 << ((v19) & 63))
                case "hello":
                    if (!((v1 ^ (-(v1))) >= (v1 * v1))) {
                        var v30 float64 = float64(((((v16 + (1e10 + v16)) + 1.5) + v16) / (v16 * float32((3.14159 - (float64((v16 * ((1.5 - v16) * v16))) * -2.25))))))
                        v31 := GenS1{F0: v18, F1: v1}
                        v32 := GenS2{F0: 646}
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                }
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
    }
    v6[2147483647] = (v17 + 784)
    var v33 uint64 = 42
    switch v17 {
    case -2147483648:
        v34 := GenS3{F0: (v19 ^ 4294967295), F1: 2147483647}
        _ = v34
    case 525:
        v35 := GenS2{F0: v19}
        _ = v35
    case 2147483647:
        if true {
            fmt.Println(v3, v5, v33)
        }
    default:
        v5 -= 25
    }
    switch (^(v17)) {
    case 1:
        for fi36 := 0; fi36 < 5; fi36++ {
            switch "" {
            case "":
                var v37 uint8 = 51
                _ = v37
            case "a":
                v3 = v3
            case "hello":
                if ((v1 / (((v1 + (v1 | (v1 & 42)))) | 1)) == (v1 << ((924) & 63))) {
                    if v3 {
                        v18 += 2147483648
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v18)
                    } else {
                        v3 = v3
                        if v3 {
                            v17 += int(v5)
                            v38 := GenS4{F0: (!v3), F1: 0}
                            var v39 uint16 = 2
                            _ = v38
                            _ = v39
                        } else {
                            var v40 int8 = (v5 * 42)
                            v41 := GenS1{F0: v18, F1: (-(v1))}
                            switch "hello" {
                            case "hello":
                                fmt.Println(v18, fi36, v2)
                            case "a":
                                var v42 uint = (v19 - v19)
                                _ = v42
                            default:
                                v43 := GenS5{F0: 1, F1: -17}
                                _ = v43
                            }
                            _ = v40
                            _ = v41
                        }
                    }
                } else {
                    _ = len(v6)
                    if (1 > 553) {
                        _ = len(v4)
                        var v44 uint8 = 128
                        var v45 bool = false
                        _ = v44
                        _ = v45
                    }
                }
            }
            _ = fi36
        }
    case -1:
        var v46 int32 = 1
        _ = v46
    case 42:
        _ = len(v4)
    default:
        v47 := GenS2{F0: v19}
        _ = v47
    }
    fmt.Println(v3, v5)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
