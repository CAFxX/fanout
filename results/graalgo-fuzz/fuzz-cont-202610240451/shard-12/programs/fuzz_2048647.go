package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int32; F1 int64; F2 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 uint; F1 uint8; F2 uint64 }
func genFn0() int16 {
    if (!((!false) && false)) {
        var v5 uint = 4294967295
        v5 = 410
        if (!false) {
            v6 := GenS1{F0: true, F1: 196}
            var v7 int32 = 1
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 2; fi8++ {
                v5 = (v5 % ((42) | 1))
                v9 := []int{fi8, fi8, (-(fi8))}
                _ = fi8
                _ = v9
            }
        }
        _ = v5
    } else {
        if true {
            v10 := []int{128, 128, 4294967295, -9223372036854775808}
            _ = v10
        }
        var v11 float64 = (((-2.25 - (float64(1e10) - (0.0 - float64(1e10)))) - 1.5) + 1.5)
        _ = v11
    }
    var v12 bool = (0 != 32767)
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: (!v12), F1: 32767}
        _ = fi13
        _ = v14
    }
    _ = v12
    return 907
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int16 = 0
    var v3 int16 = (^(v2))
    var v4 bool = false
    v15 := []int{}
    switch (("" + ("\n\t" + "")) + "foo bar") {
    case "":
        for fi16 := 0; fi16 < 4; fi16++ {
            if (!(255 < 194)) {
                var v17 float64 = 0.0
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = v17
            } else {
                fmt.Println(v3, fi16)
                fmt.Println(v4, v1, v4)
                v15 = append(v15, (v1 ^ (fi16 / ((v1) | 1))))
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 int8 = 41
                v20 := GenS1{F0: (-1 > 646), F1: 42}
                var v21 int = fi18
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            v4 = v4
            _ = fi16
        }
    case "hello":
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS2{F0: 1}
            var v24 int32 = 0
            v1 = (fi22 * 9223372036854775807)
            _ = fi22
            _ = v23
            _ = v24
        }
    case "a":
        v1 += v1
    default:
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    }
    v25 := GenS2{F0: 9223372036854775807}
    switch ("hello" + (("\n\t" + "a") + "")) {
    case "\n\t":
        for fi26 := 0; fi26 < 13; fi26++ {
            if v4 {
                v27 := GenS3{F0: 128, F1: ("\n\t" + (("a" + ("" + "\n\t")) + "foo bar"))}
                var v28 uint16 = 42
                var v29 uint64 = 18446744073709551615
                _ = v27
                _ = v28
                _ = v29
            } else {
                fmt.Println(v4)
                v30 := GenS4{F0: 1, F1: 1, F2: 1}
                for fi31 := 0; fi31 < 8; fi31++ {
                    switch 725 {
                    case -1:
                        v1 = 128
                    case 237:
                        var v32 int8 = -128
                        _ = v32
                    }
                    v3 -= (v3 + (-(v3)))
                    v4 = (!false)
                    _ = fi31
                }
                _ = v30
            }
            if len(v15) > 0 {
                v15[0] = -120
            }
            var v33 bool = (61 == 127)
            _ = fi26
            _ = v33
        }
    case "":
        fmt.Println(v2)
    case "hello":
        fmt.Println(v4)
    }
    v34 := GenS5{F0: 255}
    var v35 float64 = float64(-2.25)
    switch "" {
    case "foo bar":
        v36 := GenS5{F0: 2}
        _ = v36
    default:
        var v37 int16 = v2
        _ = v37
    }
    var v38 bool = v4
    if len(v15) > 0 {
        v15[0] = v1
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        v40 := GenS3{F0: 1, F1: (("foo bar" + ((("\n\t" + "hello") + "foo bar") + "hello")) + ("" + ("a" + "hello")))}
        for fi41 := 0; fi41 < 10; fi41++ {
            switch "a" {
            case "a":
                v2 = (v3 >> ((255) & 63))
            case "hello":
                if (0.1 < 1e10) {
                    var v42 uint32 = 860
                    var v43 int64 = -224
                    _ = v42
                    _ = v43
                } else {
                    v44 := GenS6{F0: 0, F1: 111}
                    if ((fi39 | (v1 ^ 128)) != (fi39 + 507)) {
                        v45 := GenS7{F0: 18446744073709551615, F1: 128, F2: 18446744073709551615}
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = v45
                    }
                    _ = v44
                }
            }
            _ = fi41
        }
        v1 -= v1
        _ = fi39
        _ = v40
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        var v47 float64 = 3.14159
        _ = fi46
        _ = v47
    }
    v35 = (v35 - v35)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v4
}
