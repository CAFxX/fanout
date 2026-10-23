package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint32; F1 int16; F2 int16 }
func genFn0() uint8 {
    v7 := GenS1{F0: 0}
    v8 := []int{-50, 128}
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 int32 = 42
    var v2 bool = false
    var v3 int64 = -9223372036854775808
    v4 := []int{17, 399, 9223372036854775807, 128}
    var v5 uint32 = 65535
    v6 := []int{245, 32767}
    if ((float64(float32(1e10)) == (float64(0.1) - (1e10 + (1.5 * (3.14159 * (1.5 * (-2.25 / 1.0))))))) || v2) {
        v2 = v2
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS2{F0: 42, F1: ("\n\t" + "foo bar")}
        v3 = v3
        switch 32768 {
        case 614:
            if (166 < 198) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if ((v2 && v2) && (940 <= 2)) {
                    v2 = (0.0 != (3.14159 / 1.0))
                } else {
                    var v11 int64 = (^(v3))
                    v5 = 223
                    _ = v11
                }
            } else {
                if (v2 && (1 == 2)) {
                    var v12 int16 = int16(v1)
                    v13 := GenS2{F0: 801, F1: "foo bar"}
                    _ = v12
                    _ = v13
                } else {
                    var v14 int64 = (v3 >> ((65535) & 63))
                    if (false && v2) {
                        for fi15 := 0; fi15 < 16; fi15++ {
                            var v16 int32 = (v1 + -290)
                            var v17 int8 = -85
                            _ = len(v6)
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        v3 = v14
                    } else {
                        var v18 uint16 = 255
                        v1 -= (v1 ^ (-(v1)))
                        v3 = v14
                        _ = v18
                    }
                    _ = len(v4)
                    _ = v14
                }
                var v19 uint16 = 255
                v19 += (v19 % ((65535) | 1))
                _ = v19
            }
        case 127:
            var v20 uint = 0
            _ = v20
        case -9223372036854775808:
            if (2 != 65535) {
                v2 = (v2 && v2)
                if v2 {
                    _ = len(v4)
                    for fi21 := 0; fi21 < 2; fi21++ {
                        v1 -= 1
                        _ = fi21
                    }
                }
            }
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi9
        _ = v10
    }
    var v22 int = -1
    if (v5 > v5) {
        for fi23 := 0; fi23 < 7; fi23++ {
            v22 -= fi23
            v2 = ("\n\t" < "a")
            _ = fi23
        }
        var v24 uint64 = 112
        if v2 {
            if (v2 && ((0 > uint(v24)) || v2)) {
                _ = len(v4)
            }
        } else {
            var v25 bool = v2
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 uint16 = 65535
                for fi28 := 0; fi28 < 9; fi28++ {
                    if (!v2) {
                        _ = len(v6)
                        v3 = 42
                        var v29 uint32 = v5
                        _ = v29
                    }
                    _ = fi28
                }
                _ = fi26
                _ = v27
            }
            if ((!true) && true) {
                if ((((v1 % ((v1) | 1)) >= v1) && v2) && v2) {
                    v3 -= (-(v3))
                }
                var v30 uint64 = (^(v24))
                _ = v30
            } else {
                v1 += v1
                var v31 int64 = (v3 / ((v3) | 1))
                if v2 {
                    fmt.Println(v1)
                    if v25 {
                        fmt.Println(v31)
                        var v32 float64 = ((-2.25 / 1.0) - 1.5)
                        _ = v32
                    } else {
                        fmt.Println(v24, v24, v3)
                        switch "\n\t" {
                        case "hello":
                            v25 = v25
                        }
                        fmt.Println(v2, v25, v22)
                    }
                    _ = len(v4)
                } else {
                    switch "" {
                    case "":
                        v24 = v24
                    default:
                        switch "\n\t" {
                        case "\n\t":
                            var v33 bool = v25
                            _ = v33
                        default:
                            switch (v22 >> ((0) & 63)) {
                            case 0:
                                _ = len(v4)
                            case 1:
                                v34 := GenS1{F0: (v24 >> ((11) & 63))}
                                _ = v34
                            default:
                                var v35 int64 = v31
                                _ = v35
                            }
                        }
                    }
                    var v36 float64 = (1.5 - 1.5)
                    _ = v36
                }
                _ = v31
            }
            _ = v25
        }
        _ = v24
    } else {
        switch "\n\t" {
        case "foo bar":
            v37 := GenS1{F0: 595}
            _ = v37
        }
        var v38 uint = 650
        _ = v38
    }
    v1 = v1
    v3 -= v3
    switch (v22 - 850) {
    case 128:
        if len(v6) > 0 {
            v6[0] = v22
        }
    case 366:
        if len(v6) > 0 {
            v6[0] = (v22 & 296)
        }
    }
    var v39 int64 = (v3 ^ (v3 >> ((18446744073709551615) & 63)))
    var v40 uint16 = 42
    v5 = v5
    v3 += (v3 & (v3 - int64(v1)))
    for fi41 := 0; fi41 < 5; fi41++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v42 := GenS3{F0: v5, F1: -736, F2: 42}
        if len(v4) > 0 {
            v4[0] = -896
        }
        _ = fi41
        _ = v42
    }
    if (float64(float32(3.14159)) >= (0.0 - (0.0 + 1.5))) {
        fmt.Println(v40)
        var v43 int8 = 127
        _ = v43
    } else {
        for fi44 := 0; fi44 < 11; fi44++ {
            var v45 int = 32768
            fmt.Println(v2)
            _ = fi44
            _ = v45
        }
        var v46 bool = ((2147483647 <= v5) && (65535 != 4294967295))
        _ = v46
    }
    switch -2147483648 {
    case 4294967295:
        switch int(v1) {
        case 127:
            v4 = append(v4, (v22 + 1))
        case 128:
            _ = len(v4)
        }
    case 32768:
        var v47 int16 = 684
        _ = v47
    default:
        switch "hello" {
        case "a":
            v22 += -273
        case "\n\t":
            v2 = (true || v2)
        }
    }
    switch ("foo bar" + "a") {
    case "a":
        switch (v22 | (v22 / ((32768) | 1))) {
        case 761:
            var v48 uint16 = (v40 | v40)
            _ = v48
        default:
            v4 = append(v4, (v22 >> ((18446744073709551615) & 63)))
        }
    case "hello":
        fmt.Println(v1, v1, v2)
    case "foo bar":
        var v49 string = ("a" + "foo bar")
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
