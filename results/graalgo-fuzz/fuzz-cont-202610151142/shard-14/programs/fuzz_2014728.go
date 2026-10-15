package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int16; F1 uint16; F2 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8; F1 int16 }

func main() {
    var v1 string = ""
    var v2 uint32 = 2147483648
    var v3 int64 = 844
    v3 = (v3 / (((v3 | v3)) | 1))
    if (true || true) {
        v4 := GenS1{F0: 0, F1: -32768}
        _ = v4
    } else {
        v5 := GenS2{F0: 0, F1: (v2 >> ((1) & 63))}
        _ = v5
    }
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 6; fi7++ {
        v6 -= v6
        v1 = "foo bar"
        _ = fi7
    }
    fmt.Println(v3)
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS3{F0: 1}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(fi10, fi10)
        if (-972 <= fi10) {
            var v11 int32 = 412
            _ = v11
        } else {
            v12 := GenS3{F0: 153}
            var v13 uint32 = v2
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v2 = (v2 | v2)
        _ = fi14
    }
    v1 = (("" + v1) + ("a" + v1))
    if (!true) {
        if (42 == 0) {
            v15 := GenS4{F0: 1, F1: 272}
            _ = v15
        }
    } else {
        v3 -= v3
        v6 += 32767
    }
    v1 = v1
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint32 = 2147483647
        switch "hello" {
        case "hello":
            v2 -= v17
        case "foo bar":
            var v18 float32 = 0.0
            _ = v18
        case "\n\t":
            v19 := []int{fi16, (^(fi16)), fi16, 712}
            _ = v19
        default:
            v17 += (v2 * 475)
        }
        if (!true) {
            var v20 uint = 2
            v21 := []int{(fi16 | (fi16 - fi16))}
            var v22 int = fi16
            _ = v20
            _ = v21
            _ = v22
        } else {
            v17 = v2
        }
        _ = fi16
        _ = v17
    }
    v1 = ((v1 + "") + "hello")
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := []int{(fi23 >> ((42) & 63)), fi23, (fi23 - fi23), (-(fi23))}
        if true {
            if false {
                if true {
                    v6 -= v6
                    fmt.Println(v2, v1, v6)
                } else {
                    v25 := GenS5{F0: 0, F1: 255, F2: 115}
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                    var v26 uint16 = 65535
                    _ = v25
                    _ = v26
                }
            } else {
                if true {
                    var v27 uint8 = 0
                    v28 := GenS4{F0: 445, F1: 0}
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(v6, v2)
                    v29 := GenS3{F0: 42}
                    for fi30 := 0; fi30 < 5; fi30++ {
                        var v31 int = fi30
                        v32 := GenS6{F0: 583}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    _ = v29
                }
                if ((1.5 / 1e10) < -2.25) {
                    fmt.Println(fi23, v6)
                    fmt.Println(v1, v3)
                    if (!true) {
                        if (!false) {
                            switch fi23 {
                            case -2147483648:
                                if len(v24) > 0 {
                                    v24[0] = (fi23 - (fi23 / (((fi23 & fi23)) | 1)))
                                }
                            default:
                                v33 := GenS7{F0: -77, F1: 0}
                                _ = v33
                            }
                        }
                        if (!(1 < 42)) {
                            v3 += v3
                        }
                    }
                } else {
                    if (1e10 != (0.1 * (3.14159 + (float32(0.1) + -2.25)))) {
                        for fi34 := 0; fi34 < 7; fi34++ {
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                            if (v2 > (v2 % ((v2) | 1))) {
                                v6 = v6
                                v2 -= 1
                            } else {
                                _ = len(v24)
                            }
                            _ = fi34
                        }
                    } else {
                        v35 := GenS4{F0: 9223372036854775807, F1: 284}
                        _ = v35
                    }
                }
            }
        }
        _ = fi23
        _ = v24
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        v1 = "foo bar"
        _ = fi36
    }
    var v37 int8 = 48
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v6
}
