package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 string; F1 bool; F2 uint }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint32 }
func genFn0(p0 bool, p1 uint, p2 int64) int8 {
    for fi4 := 0; fi4 < 13; fi4++ {
        fmt.Println(p2)
        _ = fi4
    }
    switch 469 {
    case -1:
        p2 = (p2 << (((^(p1))) & 63))
    case 774:
        if (33 >= 0) {
            var v5 float64 = -2.25
            for fi6 := 0; fi6 < 8; fi6++ {
                v7 := []int{fi6}
                var v8 uint16 = 255
                _ = len(v7)
                _ = fi6
                _ = v7
                _ = v8
            }
            p0 = true
            _ = v5
        } else {
            if false {
                p0 = ((((((3.14159 - 0.1) - 1.5) + 0.1) / 3.14159) >= -2.25) && p0)
                for fi9 := 0; fi9 < 6; fi9++ {
                    var v10 float32 = 3.14159
                    v11 := GenS1{F0: (fi9 ^ 1), F1: 1, F2: -656}
                    var v12 int32 = 2147483647
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                switch int(p2) {
                case 4294967295:
                    var v13 uint64 = 1
                    _ = v13
                case 0:
                    for fi14 := 0; fi14 < 13; fi14++ {
                        for fi15 := 0; fi15 < 15; fi15++ {
                            var v16 string = "foo bar"
                            v16 = v16
                            switch (v16 + v16) {
                            case "\n\t":
                                fmt.Println(v16)
                            default:
                                var v17 uint64 = 0
                                _ = v17
                            }
                            _ = fi15
                            _ = v16
                        }
                        for fi18 := 0; fi18 < 16; fi18++ {
                            if p0 {
                                if false {
                                    var v19 int8 = int8(p1)
                                    _ = v19
                                }
                                v20 := []int{128, 1}
                                var v21 int8 = 127
                                _ = v20
                                _ = v21
                            }
                            _ = fi18
                        }
                        _ = fi14
                    }
                case 32768:
                    v22 := GenS1{F0: -9223372036854775808, F1: 40, F2: -2147483648}
                    _ = v22
                }
            } else {
                for fi23 := 0; fi23 < 14; fi23++ {
                    v24 := GenS2{F0: "hello", F1: false, F2: (p1 - (p1 >> ((1) & 63)))}
                    _ = fi23
                    _ = v24
                }
            }
            var v25 uint8 = 127
            var v26 int64 = p2
            _ = v25
            _ = v26
        }
    default:
        var v27 int32 = -234
        _ = v27
    }
    switch (("\n\t" + "\n\t") + "a") {
    case "foo bar":
        var v28 int16 = 32767
        _ = v28
    }
    p0 = (("\n\t" + "a") < "")
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 bool = true
    var v2 float64 = ((((0.0 - 0.1) + 1e10) / 1.0) / ((0.0 / (1e10 / float64((0.0 + -2.25)))) - (0.0 + float64(3.14159))))
    var v3 float32 = float32((float64(1e10) / v2))
    switch (("hello" + "") + "") {
    case "foo bar":
        if (!v1) {
            switch 128 {
            case -601:
                v2 -= v2
            }
            v29 := GenS3{F0: 0, F1: -128}
            v1 = true
            _ = v29
        } else {
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 uint32 = 42
                v3 += v3
                var v32 int8 = 1
                _ = fi30
                _ = v31
                _ = v32
            }
            var v33 int16 = 42
            v34 := GenS4{F0: 389, F1: -9223372036854775808}
            _ = v33
            _ = v34
        }
    case "\n\t":
        fmt.Println(v2, v3, v2)
    case "":
        v35 := []int{32767}
        _ = v35
    }
    v36 := []int{-2147483648}
    var v37 bool = (58 >= -1)
    v38 := GenS5{F0: 255, F1: 2, F2: 65535}
    if ((v2 * v2) >= (float64(((v3 / v3) / (v3 * float32(v2)))) - 0.0)) {
        var v39 uint16 = 32767
        switch "foo bar" {
        case "hello":
            if v1 {
                if v37 {
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    var v40 uint16 = (v39 << ((0) & 63))
                    _ = v40
                }
            } else {
                v3 -= (float32(float64((v3 + v3))) - (v3 + 0.1))
                v36 = append(v36, -2147483648)
            }
        }
        _ = len(v36)
        _ = v39
    } else {
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
    }
    v41 := GenS4{F0: 18446744073709551615, F1: 1}
    var v42 uint = 738
    var v43 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v41
    _ = v42
    _ = v43
}
