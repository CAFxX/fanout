package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint; F1 uint; F2 int8 }

func main() {
    v1 := []int{797, -837, -9223372036854775808, 128}
    v2 := []string{"a"}
    var v3 uint64 = 18446744073709551615
    var v4 string = "hello"
    v5 := []uint32{772}
    var v6 int = 32768
    if (false && (2 >= 571)) {
        var v7 float32 = float32(-2.25)
        if false {
            v3 -= (v3 + v3)
        }
        _ = v7
    } else {
        var v8 uint16 = 0
        v9 := GenS1{F0: -865, F1: (v8 >> ((4294967295) & 63)), F2: (v3 & 55)}
        _ = v8
        _ = v9
    }
    if (!false) {
        v10 := GenS1{F0: 42, F1: 65535, F2: (v3 / ((v3) | 1))}
        _ = v10
    }
    v11 := GenS2{F0: 2}
    if false {
        v6 -= v6
        fmt.Println(v4, v6)
    } else {
        var v12 int32 = -357
        v12 -= v12
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        switch (((v4 + v4) + "\n\t") + (v4 + "")) {
        case "a":
            v14 := GenS2{F0: 106}
            _ = v14
        case "":
            switch "\n\t" {
            case "\n\t":
                var v15 int = (v6 * (-(v6)))
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 bool = (!true)
            _ = fi16
            _ = v17
        }
        var v18 int = (fi13 ^ 127)
        _ = fi13
        _ = v18
    }
    v19 := GenS3{F0: 32767, F1: 4294967295, F2: v3}
    v20 := GenS2{F0: 156}
    if len(v5) > 0 {
        v5[0] = 42
    }
    var v21 int64 = -9223372036854775808
    var v22 int32 = 0
    if (!(true && (!(-1 < -89)))) {
        v23 := GenS4{F0: 723, F1: 32767, F2: 2}
        var v24 float32 = -2.25
        _ = v23
        _ = v24
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (true || false) {
        v22 += v22
        for fi25 := 0; fi25 < 7; fi25++ {
            if (v22 != (^(v22))) {
                if (!((false || false) && true)) {
                    for fi26 := 0; fi26 < 19; fi26++ {
                        if (!(0 >= 456)) {
                            v3 = v3
                            if false {
                                v21 = v21
                            } else {
                                var v27 uint32 = 2147483648
                                v27 = 2147483648
                                var v28 float32 = (1.5 - ((0.0 / float32(-2.25)) / 0.1))
                                v21 -= -9223372036854775808
                                if ((v21 / ((v21) | 1)) > (^(v21))) {
                                    fmt.Println(v21)
                                }
                                if true {
                                    var v29 bool = true
                                    v30 := GenS5{F0: 65535, F1: 0, F2: -1}
                                    v22 -= -2147483648
                                    _ = v29
                                    _ = v30
                                } else {
                                    v22 = (v22 & (v22 ^ v22))
                                }
                                _ = v27
                                _ = v28
                            }
                        } else {
                            v22 = v22
                            var v31 int64 = (v21 & (v21 * (v21 & v21)))
                            switch ((v4 + (v4 + (v4 + (v4 + v4)))) + "\n\t") {
                            case "hello":
                                switch (((v4 + "foo bar") + "") + v4) {
                                case "\n\t":
                                    v31 -= (v31 + (v31 << ((0) & 63)))
                                case "hello":
                                    _ = len(v5)
                                default:
                                    var v32 int16 = 32767
                                    _ = v32
                                }
                            }
                            _ = v31
                        }
                        var v33 float64 = 1.5
                        _ = fi26
                        _ = v33
                    }
                    v4 = (v4 + v4)
                    if (v4 <= v4) {
                        v6 = v6
                        v6 = (fi25 & 0)
                        var v34 uint16 = 524
                        _ = v34
                    } else {
                        fmt.Println(v3, v21)
                        v2 = append(v2, v4)
                    }
                } else {
                    v3 = v3
                    for fi35 := 0; fi35 < 13; fi35++ {
                        v21 = 1
                        switch ("foo bar" + ("" + ("" + "\n\t"))) {
                        case "":
                            fmt.Println(v3, v6)
                        }
                        _ = fi35
                    }
                }
                v36 := GenS1{F0: (v22 / (((v22 >> ((2) & 63))) | 1)), F1: 2, F2: v3}
                _ = v36
            } else {
                var v37 uint32 = 0
                _ = v37
            }
            if false {
                v22 = (v22 - (v22 % ((v22) | 1)))
            } else {
                switch v4 {
                case "a":
                    var v38 uint64 = v3
                    _ = v38
                default:
                    var v39 int = (fi25 * -2147483648)
                    _ = v39
                }
            }
            _ = fi25
        }
        var v40 int8 = int8(v21)
        _ = v40
    }
    v21 += (v21 << ((364) & 63))
    fmt.Println(v4)
    switch (v4 + ("a" + v4)) {
    case "hello":
        if (false && (v3 > v3)) {
            for fi41 := 0; fi41 < 20; fi41++ {
                v22 += (v22 * int32(fi41))
                v22 += v22
                if (!true) {
                    v42 := GenS2{F0: 18446744073709551615}
                    v1 = append(v1, (v6 & (v6 & (-(fi41)))))
                    for fi43 := 0; fi43 < 3; fi43++ {
                        var v44 bool = false
                        v44 = v44
                        var v45 uint8 = 2
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    _ = v42
                }
                _ = fi41
            }
        }
    case "\n\t":
        for fi46 := 0; fi46 < 19; fi46++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v6)
            v4 = ("" + "")
            _ = fi46
        }
    default:
        fmt.Println(v3)
    }
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
