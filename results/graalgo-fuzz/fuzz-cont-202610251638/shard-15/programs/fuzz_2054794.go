package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 string; F1 uint32; F2 uint32 }
func genFn0() string {
    var v7 uint = 95
    v7 = v7
    switch ("a" + ("foo bar" + "a")) {
    case "":
        switch ("a" + "foo bar") {
        case "foo bar":
            v8 := GenS1{F0: -2147483648, F1: 564, F2: 42}
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 11; fi9++ {
                v10 := []int{(fi9 * fi9), (^(fi9)), 1, 127}
                if (!((65535 > 1) && false)) {
                    fmt.Println(v7, fi9)
                    var v11 uint8 = 166
                    _ = v11
                }
                fmt.Println(fi9, v7)
                _ = fi9
                _ = v10
            }
        default:
            if true {
                var v12 int8 = -128
                _ = v12
            }
        }
    case "a":
        v13 := GenS2{F0: 990, F1: (0.1 <= (((((((0.1 - (((1.5 + (float32(0.1) - 0.1)) * (1e10 - -2.25)) + 1e10)) * ((0.0 / 1e10) / (0.1 + 1e10))) / 3.14159) + 1e10) / 1.0) / ((((3.14159 * 3.14159) / ((0.0 / -2.25) + 0.1)) / float32(float64((1.5 + 1.5)))) + -2.25)) + 1e10)), F2: 1000}
        _ = v13
    case "foo bar":
        fmt.Println(v7)
    default:
        fmt.Println(v7)
    }
    _ = v7
    return ""
}
func genFn1(p0 string, p1 uint64) uint8 {
    for fi14 := 0; fi14 < 16; fi14++ {
        if (fi14 > 13) {
            fmt.Println(p1)
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            switch fi14 {
            case 0:
                p1 = 4294967295
            default:
                var v16 uint = 18446744073709551615
                _ = v16
            }
            _ = fi15
        }
        _ = fi14
    }
    p0 = ((p0 + p0) + (("a" + "foo bar") + (((p0 + (((p0 + "a") + "") + p0)) + p0) + "")))
    var v17 int8 = 42
    p0 = p0
    _ = p0
    _ = p1
    _ = v17
    return 128
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{-435}
    var v3 int64 = 860
    var v4 uint8 = 0
    var v5 uint8 = v4
    v6 := []uint32{2147483648, 4294967295, 42}
    if false {
        switch 127 {
        case -1:
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := GenS1{F0: fi18, F1: 32767, F2: v5}
                var v20 uint8 = (v5 + (^(v5)))
                v20 += v4
                _ = fi18
                _ = v19
                _ = v20
            }
        case 0:
            if false {
                v21 := GenS3{F0: (v3 & v3)}
                _ = v21
            } else {
                fmt.Println(v3, v4)
            }
        case 113:
            var v22 string = "foo bar"
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 int8 = 127
            _ = fi24
            _ = v25
        }
        v2 = append(v2, fi23)
        _ = fi23
    }
    switch ("hello" + "") {
    case "a":
        fmt.Println(v4, v3)
    }
    if (true && (77 <= 1)) {
        v26 := v1["hello"]
        for fi27 := 0; fi27 < 14; fi27++ {
            for fi28 := 0; fi28 < 1; fi28++ {
                v26 = v26
                switch "\n\t" {
                case "":
                    if (-2.25 >= 3.14159) {
                        v4 -= v4
                    } else {
                        var v29 uint8 = (v5 >> ((28) & 63))
                        if (1 != 808) {
                            fmt.Println(fi27, fi27)
                            if (v5 != genFn1((v26 + (((("\n\t" + ("\n\t" + (("hello" + ((v26 + v26) + v26)) + (v26 + v26)))) + v26) + (v26 + v26)) + v26)), 18446744073709551615)) {
                                v3 -= v3
                                switch "\n\t" {
                                case "":
                                    if false {
                                        v3 = v3
                                        var v30 string = "a"
                                        var v31 uint = 2
                                        _ = v30
                                        _ = v31
                                    }
                                }
                            }
                        } else {
                            fmt.Println(v3, fi28)
                            var v32 uint = 4294967295
                            if (true && (false && ((42 != (v32 & (v32 ^ uint(v29)))) || true))) {
                                var v33 int64 = v3
                                _ = v33
                            } else {
                                v32 = v32
                                v26 = "foo bar"
                                var v34 uint16 = 42
                                _ = v34
                            }
                            _ = v32
                        }
                        v1[v26] = ""
                        _ = v29
                    }
                }
                if false {
                    v2 = append(v2, fi27)
                } else {
                    v35 := GenS4{F0: 255, F1: genFn1(v26, 42)}
                    _ = v35
                }
                _ = fi28
            }
            _ = fi27
        }
        v26 = ""
        _ = v26
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 uint32 = 65535
        if true {
            switch 0 {
            case 0:
                v3 += (v3 / ((1) | 1))
            case -2147483648:
                fmt.Println(v3, v4)
            case -994:
                v6 = append(v6, v37)
            default:
                var v38 uint16 = 65535
                _ = v38
            }
        }
        v5 = (v4 / ((v4) | 1))
        _ = fi36
        _ = v37
    }
    v2 = append(v2, 1)
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    if len(v2) > 0 {
        v2[0] = -1
    }
    if (!(0 <= 42)) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v39 := GenS5{F0: "foo bar", F1: 627, F2: 1}
    var v40 uint64 = 0
    if (372 < 0) {
        if (true || true) {
            if len(v6) > 0 {
                v6[0] = 130
            }
            var v41 int64 = (v3 | -9223372036854775808)
            var v42 uint = 65535
            _ = v41
            _ = v42
        } else {
            for fi43 := 0; fi43 < 1; fi43++ {
                fmt.Println(v3, fi43, v4)
                fmt.Println(v3, v5)
                var v44 uint16 = 830
                _ = fi43
                _ = v44
            }
        }
        var v45 int = 32767
        v2 = append(v2, v45)
        _ = v45
    }
    _ = len(v2)
    v4 -= 128
    v1[(((("\n\t" + "hello") + "a") + "a") + "hello")] = ("foo bar" + ("hello" + "hello"))
    for fi46 := 0; fi46 < 14; fi46++ {
        for fi47 := 0; fi47 < 18; fi47++ {
            var v48 int8 = -92
            var v49 int8 = (-(v48))
            _ = fi47
            _ = v48
            _ = v49
        }
        var v50 string = "foo bar"
        _ = len(v1)
        _ = fi46
        _ = v50
    }
    for fi51 := 0; fi51 < 15; fi51++ {
        fmt.Println(fi51)
        v3 = (v3 & v3)
        _ = fi51
    }
    delete(v1, "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
