package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int8, p1 bool, p2 int16) int8 {
    var v4 int16 = -32768
    v5 := []int{2147483647, 4294967295, -103}
    var v6 int8 = p0
    for fi7 := 0; fi7 < 16; fi7++ {
        v5 = append(v5, 303)
        p0 -= v6
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (p0 | (p0 + v6))
}

func main() {
    var v1 bool = (65535 == 570)
    var v2 float32 = -2.25
    var v3 uint32 = 556
    fmt.Println(v2, v1, v2)
    var v8 int = 801
    v9 := []int{v8, v8, v8, (v8 + 32768)}
    var v10 float64 = (-2.25 + float64(v2))
    v11 := GenS1{F0: 0}
    v9 = append(v9, -2147483648)
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 int64 = 9223372036854775807
        v14 := GenS1{F0: v13}
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v8)
    var v15 uint64 = 1
    for fi16 := 0; fi16 < 2; fi16++ {
        switch (fi16 << ((360) & 63)) {
        case 127:
            v10 = v10
        case 32767:
            for fi17 := 0; fi17 < 16; fi17++ {
                if v1 {
                    var v18 uint32 = 2147483647
                    var v19 uint = 1
                    if len(v9) > 0 {
                        v9[0] = fi17
                    }
                    if v1 {
                        v15 = 2
                    }
                    v18 = v18
                    _ = v18
                    _ = v19
                } else {
                    if false {
                        fmt.Println(v10)
                    }
                    if false {
                        var v20 int32 = -1
                        if true {
                            v2 = 0.0
                            v20 -= 107
                            v15 = v15
                        }
                        var v21 uint64 = (v15 >> ((888) & 63))
                        _ = v20
                        _ = v21
                    }
                    if (v1 || (629 < uint16(v3))) {
                        if ((!(128 < 31)) && (-599 == -340)) {
                            v22 := GenS2{F0: 1, F1: 0, F2: 1}
                            _ = v22
                        } else {
                            v23 := GenS3{F0: -1, F1: 42}
                            var v24 bool = false
                            if ((^(v15)) != 42) {
                                if (3.14159 <= float64(-2.25)) {
                                    v10 = -2.25
                                    var v25 bool = (false || false)
                                    switch ("" + "a") {
                                    case "a":
                                        var v26 int16 = -524
                                        _ = v26
                                    case "\n\t":
                                        v1 = (-33 >= 0)
                                    }
                                    _ = v25
                                } else {
                                    _ = len(v9)
                                }
                                var v27 uint = 65535
                                fmt.Println(v1, fi17, v1)
                                _ = v27
                            }
                            _ = v23
                            _ = v24
                        }
                        switch ("" + "a") {
                        case "":
                            v3 = (v3 % (((v3 + 540)) | 1))
                        case "\n\t":
                            var v28 uint32 = uint32(v15)
                            _ = v28
                        case "hello":
                            if true {
                                v29 := GenS1{F0: -9223372036854775808}
                                var v30 int16 = 0
                                if false {
                                    v9 = append(v9, fi16)
                                    var v31 float64 = v10
                                    _ = v31
                                } else {
                                    v8 = (v8 % ((128) | 1))
                                }
                                if (!v1) {
                                    var v32 uint16 = 255
                                    v10 = (1.5 + (v10 * 3.14159))
                                    var v33 float32 = v2
                                    _ = v32
                                    _ = v33
                                }
                                v9 = append(v9, fi17)
                                _ = v29
                                _ = v30
                            } else {
                                v3 -= (^(v3))
                                if false {
                                    var v34 float64 = float64((3.14159 * float32(float64(0.0))))
                                    _ = v34
                                } else {
                                    v1 = (!v1)
                                    v15 = (v15 << ((366) & 63))
                                }
                                if v1 {
                                    v2 = v2
                                }
                            }
                        }
                    }
                }
                _ = fi17
            }
        default:
            v35 := GenS4{F0: (v3 % (((v3 - v3)) | 1)), F1: 76}
            _ = v35
        }
        _ = fi16
    }
    v1 = (true && v1)
    var v36 uint = 608
    v15 -= (v15 << ((v36) & 63))
    if ((-1 < 1000) && (0.0 <= ((v10 * float64(v2)) * 1.5))) {
        if true {
            switch 2147483647 {
            case 32768:
                var v37 string = ""
                _ = v37
            default:
                if len(v9) > 0 {
                    v9[0] = v8
                }
            }
        } else {
            for fi38 := 0; fi38 < 2; fi38++ {
                fmt.Println(v1, v2)
                v3 = 152
                var v39 float32 = v2
                _ = fi38
                _ = v39
            }
            if (-2.25 <= (v10 + v10)) {
                v40 := GenS2{F0: (v15 % ((v15) | 1)), F1: -1, F2: 54}
                _ = v40
            } else {
                for fi41 := 0; fi41 < 14; fi41++ {
                    for fi42 := 0; fi42 < 6; fi42++ {
                        fmt.Println(fi42)
                        v36 -= v36
                        _ = fi42
                    }
                    switch "\n\t" {
                    case "foo bar":
                        if (-1 < 42) {
                            var v43 int8 = genFn0(23, (-1 <= 1), -32768)
                            var v44 uint = 65535
                            _ = v43
                            _ = v44
                        }
                    default:
                        for fi45 := 0; fi45 < 1; fi45++ {
                            v46 := GenS5{F0: 255}
                            v15 = v15
                            var v47 uint64 = (^(v15))
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    }
                    _ = fi41
                }
            }
        }
        fmt.Println(v36)
        for fi48 := 0; fi48 < 14; fi48++ {
            var v49 float64 = 1e10
            fmt.Println(v10, v10, fi48)
            _ = fi48
            _ = v49
        }
    }
    v50 := GenS3{F0: 0, F1: 121}
    for fi51 := 0; fi51 < 4; fi51++ {
        var v52 uint = 255
        for fi53 := 0; fi53 < 7; fi53++ {
            if (32767 >= 2) {
                var v54 int = (fi51 * 2147483647)
                _ = len(v9)
                _ = v54
            }
            _ = fi53
        }
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v50
    _ = v8
    _ = v9
}
