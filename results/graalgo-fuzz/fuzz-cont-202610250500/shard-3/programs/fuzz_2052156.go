package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 string }
func genFn0() uint8 {
    var v4 int64 = 9223372036854775807
    v4 -= v4
    if (false || (2147483647 >= -2147483648)) {
        v5 := []int{460}
        v4 = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        switch 1 {
        case 1:
            v7 := GenS1{F0: -9223372036854775808}
            _ = v7
        case 2147483647:
            v4 = (^(v4))
        case 766:
            v8 := GenS1{F0: v4}
            _ = v8
        }
        if false {
            fmt.Println(fi6, v4)
            v9 := GenS1{F0: -1}
            _ = v9
        }
        v10 := []int{fi6}
        _ = fi6
        _ = v10
    }
    _ = v4
    return 241
}
func genFn1(p0 uint, p1 uint16) int {
    v11 := GenS1{F0: 0}
    if (!true) {
        p1 += 42
        fmt.Println(p0, p0)
        switch ("\n\t" + "") {
        case "":
            switch 128 {
            case 0:
                v12 := []int{974, -9223372036854775808, 1, 32767}
                _ = v12
            case 128:
                v13 := GenS2{F0: false, F1: -266}
                _ = v13
            case 127:
                v14 := []int{-9223372036854775808, 507}
                _ = v14
            }
        case "foo bar":
            var v15 int32 = -2147483648
            _ = v15
        case "\n\t":
            for fi16 := 0; fi16 < 18; fi16++ {
                for fi17 := 0; fi17 < 17; fi17++ {
                    switch ("\n\t" + "hello") {
                    case "a":
                        var v18 int16 = 32767
                        _ = v18
                    case "":
                        p1 = p1
                    case "hello":
                        v19 := GenS2{F0: false, F1: 0}
                        _ = v19
                    }
                    _ = fi17
                }
                if true {
                    switch 227 {
                    case 774:
                        if true {
                            p1 += 42
                            p0 += 18446744073709551615
                        } else {
                            v20 := GenS3{F0: 9223372036854775807}
                            _ = v20
                        }
                    case -191:
                        fmt.Println(fi16, p0)
                    case 32767:
                        v21 := []int{fi16, -736, (fi16 ^ (fi16 >> (((^(p0))) & 63)))}
                        _ = v21
                    }
                    v22 := []int{-9223372036854775808, -9223372036854775808, (^(fi16))}
                    _ = v22
                }
                p0 -= 1
                _ = fi16
            }
        default:
            if true {
                v23 := []int{-1}
                var v24 float32 = ((3.14159 + 1.5) + float32(1.5))
                if (2 != 65535) {
                    var v25 uint = p0
                    _ = v25
                } else {
                    fmt.Println(v24, v24)
                }
                _ = v23
                _ = v24
            }
        }
    }
    var v26 uint8 = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v26
    return -2147483648
}

func main() {
    var v1 int32 = 42
    var v2 uint64 = 42
    var v3 int64 = -280
    if false {
        v27 := GenS3{F0: (v3 ^ v3)}
        for fi28 := 0; fi28 < 2; fi28++ {
            fmt.Println(v3, fi28, v1)
            _ = fi28
        }
        _ = v27
    } else {
        v29 := GenS4{F0: 42, F1: (v3 ^ (v3 >> ((255) & 63)))}
        _ = v29
    }
    fmt.Println(v2)
    v30 := []int{0, -9223372036854775808}
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    v30 = append(v30, 2147483647)
    for fi31 := 0; fi31 < 15; fi31++ {
        for fi32 := 0; fi32 < 11; fi32++ {
            v33 := GenS1{F0: 9223372036854775807}
            _ = fi32
            _ = v33
        }
        for fi34 := 0; fi34 < 20; fi34++ {
            v1 = (v1 - 0)
            if (0.1 == ((-2.25 + 3.14159) * (1e10 / float64(1.5)))) {
                var v35 bool = (!true)
                var v36 uint = 0
                v37 := GenS1{F0: v3}
                if v35 {
                    v38 := GenS1{F0: -9223372036854775808}
                    var v39 bool = v35
                    _ = v38
                    _ = v39
                } else {
                    v1 += 588
                    v36 = v36
                    _ = len(v30)
                }
                _ = v35
                _ = v36
                _ = v37
            } else {
                if false {
                    if (!true) {
                        var v40 string = "a"
                        _ = v40
                    }
                } else {
                    _ = len(v30)
                    if len(v30) > 0 {
                        v30[0] = (-(fi34))
                    }
                }
            }
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            if len(v30) > 0 {
                v30[0] = (fi34 / ((genFn1(65535, 255)) | 1))
            }
            v2 += (v2 - v2)
            _ = fi34
        }
        switch 1 {
        case 471:
            for fi41 := 0; fi41 < 7; fi41++ {
                var v42 string = "foo bar"
                v43 := GenS3{F0: v3}
                _ = fi41
                _ = v42
                _ = v43
            }
        default:
            fmt.Println(fi31, fi31)
        }
        _ = fi31
    }
    var v44 float32 = 0.1
    if len(v30) > 0 {
        v30[0] = 32768
    }
    fmt.Println(v44)
    if (2 <= 133) {
        var v45 int16 = 32767
        var v46 uint64 = 182
        v30 = append(v30, 0)
        _ = v45
        _ = v46
    }
    v44 = (v44 + float32(-2.25))
    v30 = append(v30, genFn1(4294967295, 395))
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    if false {
        var v47 int16 = -32768
        v3 = int64(v47)
        switch 127 {
        case 149:
            var v48 float32 = (-2.25 + 0.1)
            _ = v48
        case 2147483647:
            fmt.Println(v47)
        case 575:
            v49 := GenS5{F0: "hello"}
            _ = v49
        default:
            v50 := GenS4{F0: 100, F1: v3}
            _ = v50
        }
        _ = v47
    } else {
        fmt.Println(v44, v44, v44)
    }
    v1 -= v1
    for fi51 := 0; fi51 < 18; fi51++ {
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        _ = fi51
    }
    v2 = (v2 / (((v2 ^ v2)) | 1))
    v44 = v44
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v44
}
