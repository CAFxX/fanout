package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int8 }

func main() {
    var v1 float32 = (0.1 + ((float32(3.14159) - float32((1.5 + 0.0))) + 0.0))
    var v2 int16 = 95
    var v3 uint16 = 255
    var v4 float32 = (v1 - 0.0)
    v4 -= v1
    for fi5 := 0; fi5 < 3; fi5++ {
        v2 = v2
        for fi6 := 0; fi6 < 9; fi6++ {
            if (1 > 100) {
                var v7 uint8 = 42
                switch fi5 {
                case 0:
                    v8 := GenS1{F0: int8(v3)}
                    _ = v8
                case -2147483648:
                    v3 = v3
                default:
                    v3 = v3
                }
                _ = v7
            }
            v9 := []int{(fi5 | fi6), fi5, fi5}
            _ = fi6
            _ = v9
        }
        _ = fi5
    }
    v10 := []int{32768, 1, 573, 4294967295}
    _ = len(v10)
    v4 = 0.1
    switch "foo bar" {
    case "hello":
        fmt.Println(v1)
    case "\n\t":
        var v11 int32 = -1
        _ = v11
    default:
        v3 -= v3
    }
    if ((2 >= 4294967295) || (766 >= 4294967295)) {
        var v12 int32 = 890
        var v13 int32 = 8
        _ = v12
        _ = v13
    } else {
        if len(v10) > 0 {
            v10[0] = 128
        }
    }
    if false {
        v2 = v2
        for fi14 := 0; fi14 < 6; fi14++ {
            v1 -= float32(((3.14159 / (-2.25 + (1e10 * ((0.0 / (3.14159 + 0.0)) + (0.1 + (1.5 * (1e10 / 1e10))))))) / 1e10))
            if true {
                switch ("hello" + "foo bar") {
                case "\n\t":
                    switch 128 {
                    case 9223372036854775807:
                        switch (fi14 << ((18446744073709551615) & 63)) {
                        case 2147483647:
                            _ = len(v10)
                        }
                    default:
                        for fi15 := 0; fi15 < 14; fi15++ {
                            v16 := GenS2{F0: ("a" + "\n\t")}
                            _ = fi15
                            _ = v16
                        }
                    }
                default:
                    var v17 uint8 = 42
                    _ = v17
                }
                for fi18 := 0; fi18 < 15; fi18++ {
                    if true {
                        var v19 uint32 = 2147483648
                        v20 := GenS2{F0: "hello"}
                        v1 = (1.5 * v1)
                        _ = v19
                        _ = v20
                    }
                    _ = fi18
                }
                if len(v10) > 0 {
                    v10[0] = fi14
                }
            } else {
                switch (fi14 >> ((255) & 63)) {
                case 127:
                    switch fi14 {
                    case 42:
                        for fi21 := 0; fi21 < 14; fi21++ {
                            fmt.Println(v3)
                            _ = fi21
                        }
                    case -9223372036854775808:
                        if len(v10) > 0 {
                            v10[0] = (-(fi14))
                        }
                    }
                case -2147483648:
                    var v22 int64 = 171
                    _ = v22
                }
            }
            _ = fi14
        }
    }
    var v23 uint32 = 2147483648
    var v24 bool = true
    switch 42 {
    case 9223372036854775807:
        v2 = (v2 ^ 1)
    case 1:
        if (298 > 292) {
            var v25 float32 = (1e10 / ((v4 + (3.14159 - ((1.5 + 0.0) / 1e10))) * ((v4 * (((v4 * 0.0) * (0.1 * v4)) - (v4 * float32(1e10)))) + v4)))
            var v26 int16 = v2
            switch "" {
            case "":
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            _ = v25
            _ = v26
        }
    case -784:
        v24 = ((((1.5 / 1.0) - float64(1e10)) + ((0.1 - 3.14159) - 0.1)) == -2.25)
    default:
        var v27 int = 4294967295
        _ = v27
    }
    fmt.Println(v4, v1)
    fmt.Println(v2, v3)
    if (v2 == (v2 >> ((779) & 63))) {
        v4 = v4
        v28 := GenS2{F0: "hello"}
        _ = v28
    } else {
        fmt.Println(v1, v24, v23)
        switch 9223372036854775807 {
        case -9223372036854775808:
            v29 := GenS3{F0: (v2 | 32767), F1: 1}
            _ = v29
        }
        v2 = v2
    }
    var v30 uint64 = 4294967295
    if v24 {
        for fi31 := 0; fi31 < 13; fi31++ {
            v3 += 2
            _ = fi31
        }
        for fi32 := 0; fi32 < 17; fi32++ {
            for fi33 := 0; fi33 < 7; fi33++ {
                v10 = append(v10, (fi33 + 32767))
                v34 := GenS1{F0: 0}
                _ = fi33
                _ = v34
            }
            v10 = append(v10, 32768)
            for fi35 := 0; fi35 < 18; fi35++ {
                var v36 int8 = 0
                _ = fi35
                _ = v36
            }
            _ = fi32
        }
    }
    var v37 bool = (42 > -1)
    var v38 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
}
