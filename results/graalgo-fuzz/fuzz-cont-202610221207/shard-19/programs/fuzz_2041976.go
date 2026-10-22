package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint64; F2 int8 }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint64; F1 int8; F2 uint8 }

func main() {
    var v1 int32 = -1
    var v2 uint64 = 42
    var v3 string = "hello"
    var v4 float32 = 3.14159
    v5 := []int{128, 32767}
    var v6 bool = (2 == 931)
    if (0 > (^(v2))) {
        switch 739 {
        case -8:
            fmt.Println(v1)
        }
        for fi7 := 0; fi7 < 13; fi7++ {
            switch 128 {
            case 42:
                var v8 float32 = v4
                _ = v8
            case -143:
                var v9 float64 = ((0.0 + (((float64(1.5) + -2.25) + (1.5 / (1e10 * (float64((v4 + (v4 / v4))) + float64(v4))))) - 1.5)) - 3.14159)
                _ = v9
            }
            v3 = "foo bar"
            v2 += (v2 | 0)
            _ = fi7
        }
        switch (("hello" + v3) + v3) {
        case "foo bar":
            fmt.Println(v3)
        case "hello":
            if false {
                if (255 < 897) {
                    if true {
                        for fi10 := 0; fi10 < 14; fi10++ {
                            var v11 float32 = 3.14159
                            if true {
                                v2 -= (v2 % (((v2 * 749)) | 1))
                            } else {
                                _ = len(v5)
                                v11 += v4
                                v12 := GenS1{F0: 1, F1: 1, F2: 65535}
                                _ = v12
                            }
                            _ = fi10
                            _ = v11
                        }
                        v4 += v4
                    }
                } else {
                    var v13 int8 = 49
                    _ = v13
                }
                v6 = false
            }
        case "a":
            v14 := GenS2{F0: 1000, F1: (^(v2)), F2: -1}
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 2; fi15++ {
            fmt.Println(v6)
            _ = fi15
        }
    }
    if v6 {
        v1 = 42
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int16 = 1000
            _ = fi16
            _ = v17
        }
    }
    v5 = append(v5, -9223372036854775808)
    for fi18 := 0; fi18 < 15; fi18++ {
        fmt.Println(v4, v4, v6)
        if false {
            var v19 int = fi18
            v3 = v3
            if v6 {
                v5 = append(v5, fi18)
            } else {
                if (!v6) {
                    var v20 uint = 0
                    v21 := GenS3{F0: (v1 - 0), F1: 2147483647, F2: 1}
                    fmt.Println(fi18)
                    _ = v20
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 5; fi22++ {
                        var v23 uint = 4294967295
                        if (!v6) {
                            fmt.Println(fi18, v2)
                        }
                        _ = fi22
                        _ = v23
                    }
                    if (!v6) {
                        v19 = -1
                        var v24 uint64 = (v2 + (v2 + (v2 | (v2 | 1))))
                        v6 = true
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            v26 := GenS4{F0: 178, F1: 997}
                            v27 := GenS3{F0: v1, F1: v1, F2: -128}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                }
                v28 := GenS5{F0: 0, F1: v19}
                _ = len(v5)
                _ = v28
            }
            _ = v19
        }
        _ = fi18
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        _ = len(v5)
        var v30 bool = (v6 || ((1e10 < ((-2.25 / float64(float32(1e10))) / (3.14159 / (0.1 + (0.1 * 1e10))))) || false))
        v2 = v2
        _ = fi29
        _ = v30
    }
    v1 = v1
    switch v3 {
    case "a":
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 float32 = 0.0
            var v33 int = 491
            _ = fi31
            _ = v32
            _ = v33
        }
    case "hello":
        for fi34 := 0; fi34 < 4; fi34++ {
            for fi35 := 0; fi35 < 1; fi35++ {
                fmt.Println(fi35, v3)
                v3 = v3
                _ = fi35
            }
            for fi36 := 0; fi36 < 16; fi36++ {
                if (42 < fi36) {
                    v6 = v6
                    fmt.Println(v1, v4)
                }
                if (!false) {
                    if (2147483647 == v1) {
                        var v37 int = (fi36 - (fi34 << ((255) & 63)))
                        v1 += (-(v1))
                        _ = v37
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v38 uint8 = 255
                        v2 += 201
                        var v39 int64 = 9223372036854775807
                        fmt.Println(v6, v4)
                        _ = v38
                        _ = v39
                    }
                    var v40 uint16 = 42
                    var v41 int8 = 127
                    _ = v40
                    _ = v41
                }
                var v42 float64 = (-2.25 - 1.5)
                _ = fi36
                _ = v42
            }
            switch (^(fi34)) {
            case 9223372036854775807:
                v2 = 1
            default:
                for fi43 := 0; fi43 < 9; fi43++ {
                    _ = len(v5)
                    v6 = v6
                    _ = fi43
                }
            }
            _ = fi34
        }
    default:
        var v44 int8 = 3
        _ = v44
    }
    if (42 <= -1) {
        for fi45 := 0; fi45 < 18; fi45++ {
            v6 = (18446744073709551615 <= 931)
            if (117 > 1) {
                fmt.Println(v1, v4)
                if (true || v6) {
                    if true {
                        _ = len(v5)
                    }
                    for fi46 := 0; fi46 < 10; fi46++ {
                        _ = len(v5)
                        _ = fi46
                    }
                }
                v2 -= (v2 ^ (v2 | v2))
            } else {
                switch v3 {
                case "":
                    fmt.Println(v2, fi45, v4)
                case "\n\t":
                    fmt.Println(v4)
                case "a":
                    var v47 uint16 = 1
                    _ = v47
                }
                var v48 int8 = 42
                v4 -= (float32(-2.25) - ((v4 * (0.1 * v4)) + float32(((1.5 / (1e10 / ((-2.25 + (3.14159 * (1e10 * -2.25))) - -2.25))) + (float64(v4) * float64(float32(((((3.14159 + 0.1) * float64(v4)) - -2.25) - 3.14159))))))))
                _ = v48
            }
            _ = fi45
        }
        var v49 uint16 = 393
        switch "a" {
        case "hello":
            var v50 uint16 = (v49 % (((^(v49))) | 1))
            _ = v50
        case "a":
            v51 := GenS1{F0: -84, F1: 346, F2: 405}
            _ = v51
        }
        _ = v49
    } else {
        fmt.Println(v4, v3)
        var v52 float64 = (1e10 * float64((1.5 / (float32(float64((float32(0.0) / -2.25))) / (0.0 / (3.14159 / (v4 + v4)))))))
        switch (v3 + (v3 + "")) {
        case "\n\t":
            for fi53 := 0; fi53 < 1; fi53++ {
                var v54 int64 = -38
                v5 = append(v5, 1)
                _ = fi53
                _ = v54
            }
        }
        _ = v52
    }
    v55 := GenS6{F0: v2, F1: -98, F2: 127}
    fmt.Println(v4, v3, v3)
    var v56 uint = 1
    if (12 >= 42) {
        v3 = "a"
        var v57 uint = v56
        _ = v57
    }
    v6 = (!(v3 >= "foo bar"))
    if len(v5) > 0 {
        v5[0] = 703
    }
    var v58 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v58
    _ = v6
}
