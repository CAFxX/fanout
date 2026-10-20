package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int32; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint; F2 string }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 float32 = ((((0.0 - float32(1e10)) + float32(1e10)) / -2.25) - 1.5)
    var v2 int = 32768
    var v3 int8 = 42
    var v4 uint8 = 255
    var v5 int8 = (v3 + v3)
    v5 -= 42
    v2 += v2
    v6 := GenS1{F0: (708 < 504)}
    var v7 uint = 4294967295
    if (-889 > -898) {
        v5 -= (v5 % (((v3 >> ((v7) & 63))) | 1))
    } else {
        fmt.Println(v1, v2, v3)
        v8 := []int{}
        _ = len(v8)
        _ = v8
    }
    if (!true) {
        v9 := GenS1{F0: true}
        for fi10 := 0; fi10 < 3; fi10++ {
            if true {
                v11 := []int{(-(v2)), -9223372036854775808}
                v12 := GenS1{F0: true}
                _ = v11
                _ = v12
            } else {
                fmt.Println(v2, v4)
                for fi13 := 0; fi13 < 7; fi13++ {
                    v5 = v3
                    _ = fi13
                }
            }
            _ = fi10
        }
        v3 += (-(v5))
        _ = v9
    }
    if (3.14159 >= float64(1.5)) {
        v2 -= (v2 + (-(v2)))
        v14 := GenS1{F0: (false && true)}
        v15 := GenS1{F0: (0 < -1)}
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 794, F1: 263, F2: -82}
    v4 = v4
    for fi17 := 0; fi17 < 11; fi17++ {
        v1 = v1
        if (0 <= 0) {
            switch (v2 | (v2 + fi17)) {
            case 1:
                for fi18 := 0; fi18 < 3; fi18++ {
                    fmt.Println(v7)
                    if false {
                        var v19 int64 = 0
                        v20 := GenS2{F0: -32768, F1: 1, F2: (v19 + (v19 << ((v7) & 63)))}
                        _ = v19
                        _ = v20
                    }
                    _ = fi18
                }
            }
            v21 := []int{}
            if (195 == 4294967295) {
                if (9223372036854775807 >= -1) {
                    v22 := GenS3{F0: 9223372036854775807}
                    _ = v22
                } else {
                    v2 = v2
                    if (false || (false || (!true))) {
                        for fi23 := 0; fi23 < 11; fi23++ {
                            var v24 uint8 = (^(v4))
                            if true {
                                v5 -= (^(v3))
                                v3 -= v5
                                if true {
                                    v25 := GenS1{F0: true}
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    if ((^(v24)) >= 128) {
                                        if true {
                                            var v26 int8 = v5
                                            if true {
                                                var v27 int16 = 1000
                                                v3 -= (v26 % ((127) | 1))
                                                switch "hello" {
                                                case "a":
                                                    var v28 float32 = (v1 + (v1 / v1))
                                                    _ = v28
                                                case "hello":
                                                    var v29 uint = 520
                                                    _ = v29
                                                case "\n\t":
                                                    var v30 float32 = 1.5
                                                    _ = v30
                                                default:
                                                    var v31 int16 = (v27 / (((v27 + v27)) | 1))
                                                    var v32 float64 = 0.0
                                                    if len(v21) >= 2 {
                                                        v21 = v21[1:len(v21)]
                                                    }
                                                    _ = v31
                                                    _ = v32
                                                }
                                                _ = v27
                                            } else {
                                                v21 = append(v21, 362)
                                            }
                                            _ = v26
                                        } else {
                                            v33 := GenS2{F0: 0, F1: 42, F2: 9223372036854775807}
                                            var v34 uint8 = (v24 / ((v24) | 1))
                                            _ = v33
                                            _ = v34
                                        }
                                        var v35 float32 = v1
                                        v1 -= ((float32((float64((((0.0 - (-2.25 * -2.25)) + float32(0.0)) / (v1 - v1))) / -2.25)) - (1e10 + v1)) / v1)
                                        _ = v35
                                    } else {
                                        switch "foo bar" {
                                        case "foo bar":
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v36 int64 = 0
                                                _ = v36
                                            }
                                        case "":
                                            v2 = 32768
                                        case "a":
                                            fmt.Println(v4, v24, v5)
                                        }
                                        var v37 bool = true
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                        _ = v37
                                    }
                                    _ = v25
                                } else {
                                    fmt.Println(v24, v4, fi23)
                                    if false {
                                        switch v2 {
                                        case 785:
                                            v3 = -31
                                        case -19:
                                            var v38 uint16 = 929
                                            _ = v38
                                        case 32767:
                                            fmt.Println(fi23)
                                        }
                                    }
                                }
                            } else {
                                var v39 uint32 = 42
                                _ = v39
                            }
                            var v40 string = "hello"
                            v41 := GenS2{F0: -303, F1: -2147483648, F2: -9223372036854775808}
                            _ = fi23
                            _ = v24
                            _ = v40
                            _ = v41
                        }
                        v4 = (v4 + v4)
                        var v42 int16 = 1000
                        _ = v42
                    }
                    switch v2 {
                    case -1:
                        var v43 int8 = (v3 ^ (-(v5)))
                        _ = v43
                    default:
                        var v44 int = v2
                        _ = v44
                    }
                }
                var v45 float32 = (0.1 - v1)
                for fi46 := 0; fi46 < 11; fi46++ {
                    _ = len(v21)
                    if (!(((("foo bar" + "\n\t") + (("" + ("a" + "a")) + "\n\t")) + ("foo bar" + "hello")) <= "hello")) {
                        v21 = append(v21, (fi17 - (fi46 >> (((v7 / ((v7) | 1))) & 63))))
                        if false {
                            _ = len(v21)
                            v45 += v45
                        } else {
                            v2 -= 42
                        }
                        v47 := GenS3{F0: v2}
                        v3 += 42
                        var v48 float32 = v1
                        var v49 float32 = (3.14159 / ((-2.25 - 1.5) / 1.5))
                        var v50 uint = v7
                        v7 = v50
                        v7 = (v50 / ((39) | 1))
                        var v51 float32 = (float32(float64(v48)) * (float32(0.1) + 0.0))
                        var v52 int = int(v4)
                        _ = v47
                        _ = v48
                        _ = v49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    v53 := GenS2{F0: int16(fi46), F1: -2147483648, F2: 1}
                    if true {
                        if false {
                            v54 := GenS4{F0: (fi17 >> ((4294967295) & 63)), F1: 42, F2: ("hello" + "foo bar")}
                            v2 -= fi17
                            _ = v54
                        }
                        v2 += (v2 | (fi46 & (v2 % ((fi17) | 1))))
                        v55 := GenS5{F0: 65535}
                        _ = v55
                    } else {
                        v2 = (fi46 - -2)
                        fmt.Println(fi17, v1, v45)
                    }
                    _ = fi46
                    _ = v53
                }
                _ = v45
            }
            _ = v21
        } else {
            var v56 int16 = 775
            _ = v56
        }
        _ = fi17
    }
    var v57 int8 = v5
    fmt.Println(v3, v1)
    v7 += v7
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v7
}
