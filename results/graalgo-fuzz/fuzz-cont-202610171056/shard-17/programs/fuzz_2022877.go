package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 int64; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }

func main() {
    var v1 float32 = 1.5
    var v2 int = 4294967295
    var v3 int16 = 158
    v3 = v3
    v4 := GenS1{F0: -764, F1: v2, F2: 42}
    fmt.Println(v3)
    var v5 bool = (-441 == -537)
    var v6 int = 9223372036854775807
    if v5 {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := []int{750, -1, -1, v2}
            v8 = append(v8, v6)
            _ = fi7
            _ = v8
        }
        v9 := GenS2{F0: 4294967295, F1: 1, F2: (!true)}
        _ = v9
    } else {
        v10 := GenS3{F0: 666}
        _ = v10
    }
    v1 = v1
    v11 := GenS3{F0: 576}
    fmt.Println(v2)
    var v12 int8 = 42
    v13 := []int{(v2 + v6)}
    for fi14 := 0; fi14 < 16; fi14++ {
        fmt.Println(v2, v1, fi14)
        _ = fi14
    }
    fmt.Println(v3, v6, v6)
    switch ("\n\t" + (("" + ((("\n\t" + "foo bar") + "a") + "\n\t")) + ("a" + (("hello" + "hello") + "")))) {
    case "\n\t":
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 13; fi16++ {
                v5 = (v5 && v5)
                _ = len(v13)
                v1 += (0.1 * v1)
                _ = fi16
            }
            switch (-(fi15)) {
            case -581:
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 float32 = -2.25
                    v18 += 3.14159
                    _ = fi17
                    _ = v18
                }
            case 0:
                if (!v5) {
                    if true {
                        if len(v13) > 0 {
                            v13[0] = fi15
                        }
                        var v19 int32 = 42
                        _ = v19
                    }
                    var v20 float64 = 0.1
                    _ = v20
                }
            default:
                if (0.1 > 0.0) {
                    v12 -= 1
                    if (v5 && false) {
                        for fi21 := 0; fi21 < 14; fi21++ {
                            fmt.Println(v2, v5, v6)
                            if v5 {
                                fmt.Println(v1)
                            }
                            _ = fi21
                        }
                    }
                } else {
                    fmt.Println(v2, v12)
                }
            }
            _ = fi15
        }
    case "foo bar":
        var v22 int64 = 1
        _ = v22
    default:
        var v23 bool = (!(977 != 1))
        _ = v23
    }
    switch "hello" {
    case "hello":
        if false {
            var v24 float32 = -2.25
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            if len(v13) > 0 {
                v13[0] = (-(v2))
            }
            _ = v24
        } else {
            for fi25 := 0; fi25 < 3; fi25++ {
                for fi26 := 0; fi26 < 16; fi26++ {
                    v3 -= (v3 >> ((4294967295) & 63))
                    v27 := GenS4{F0: -2147483648}
                    if len(v13) > 0 {
                        v13[0] = -527
                    }
                    _ = fi26
                    _ = v27
                }
                for fi28 := 0; fi28 < 11; fi28++ {
                    if len(v13) > 0 {
                        v13[0] = fi28
                    }
                    switch (v6 << ((1) & 63)) {
                    case -2147483648:
                        var v29 float32 = float32((((1e10 + ((((0.0 + float64(1e10)) + -2.25) + 0.1) / (0.1 + (((float64((v1 / float32(0.1))) - 0.0) / 3.14159) + -2.25)))) - 3.14159) * -2.25))
                        _ = v29
                    }
                    v3 = v3
                    var v30 bool = v5
                    _ = fi28
                    _ = v30
                }
                _ = fi25
            }
            v31 := GenS5{F0: false}
            var v32 int64 = 42
            _ = v31
            _ = v32
        }
    case "":
        v1 = (0.0 * -2.25)
    case "a":
        v13 = append(v13, (v2 * v6))
    default:
        var v33 int32 = 995
        _ = v33
    }
    v3 = (^(v3))
    fmt.Println(v1, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
