package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int32; F2 int16 }
func genFn0(p0 int16, p1 int64, p2 uint16) uint64 {
    switch ("" + "foo bar") {
    case "\n\t":
        v6 := []int{-9223372036854775808, 127, 32768}
        _ = v6
    case "a":
        v7 := GenS1{F0: -128, F1: "foo bar"}
        _ = v7
    default:
        v8 := []int{0, 32767}
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: "\n\t"}
    if (((p2 ^ p2) != p2) && true) {
        v10 := []int{32767, 127, -9223372036854775808, 915}
        _ = v10
    }
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 4294967295
}
func genFn1(p0 bool) int16 {
    if ("hello" != "") {
        switch 9223372036854775807 {
        case 121:
            p0 = false
        default:
            p0 = p0
        }
        fmt.Println(p0)
    } else {
        if p0 {
            fmt.Println(p0)
            if (255 >= 42) {
                if true {
                    v11 := []int{-1, 32767, -1, 175}
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 14; fi12++ {
                        p0 = ((true || ((!true) && (-2147483648 > -1))) || true)
                        for fi13 := 0; fi13 < 15; fi13++ {
                            fmt.Println(fi12, fi13)
                            var v14 int64 = -9223372036854775808
                            _ = fi13
                            _ = v14
                        }
                        _ = fi12
                    }
                }
            }
        }
        p0 = (-50 <= 9223372036854775807)
    }
    var v15 int16 = 961
    v15 = (v15 / ((v15) | 1))
    v15 = (^(v15))
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := GenS2{F0: -893, F1: 255, F2: 42}
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = v15
    return 483
}

func main() {
    v1 := []uint32{2, 321}
    var v2 int16 = -1
    v3 := []uint32{}
    var v4 string = "hello"
    var v5 float64 = 1e10
    fmt.Println(v5, v2, v5)
    var v18 uint64 = 18446744073709551615
    for fi19 := 0; fi19 < 12; fi19++ {
        if (127 != -5) {
            v4 = ((((v4 + v4) + ("" + v4)) + "hello") + ("a" + v4))
        } else {
            switch fi19 {
            case 1:
                var v20 int32 = 0
                _ = v20
            case 42:
                _ = len(v1)
            case 455:
                v2 -= -707
            }
            if ((!(float32(-2.25) == -2.25)) && false) {
                var v21 int8 = 1
                var v22 float64 = v5
                _ = v21
                _ = v22
            } else {
                if (!(1 == 2147483647)) {
                    switch (fi19 / ((fi19) | 1)) {
                    case 9223372036854775807:
                        var v23 float64 = ((v5 / v5) + v5)
                        _ = v23
                    default:
                        v24 := GenS1{F0: -26, F1: "a"}
                        _ = v24
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v25 uint64 = v18
                    _ = v25
                }
                v18 += (v18 & 100)
                var v26 float32 = 3.14159
                _ = v26
            }
        }
        v27 := GenS1{F0: 127, F1: ""}
        if (42 <= 2147483647) {
            v28 := GenS3{F0: 156, F1: 2147483647}
            _ = v28
        } else {
            var v29 uint16 = 2
            var v30 int8 = -34
            _ = v29
            _ = v30
        }
        _ = fi19
        _ = v27
    }
    if len(v3) > 0 {
        v3[0] = 526
    }
    fmt.Println(v18)
    v1 = append(v1, 65535)
    if true {
        var v31 bool = true
        _ = v31
    } else {
        var v32 uint = 255
        if (true || (false && true)) {
            if true {
                v33 := GenS3{F0: 9223372036854775807, F1: 0}
                for fi34 := 0; fi34 < 14; fi34++ {
                    if (!((255 <= 857) && (false || true))) {
                        v32 -= 409
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v35 := GenS4{F0: uint32(v18), F1: -1, F2: genFn1((!true))}
                        _ = v35
                    }
                    _ = fi34
                }
                _ = v33
            } else {
                v2 += v2
                v32 = v32
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
            }
        } else {
            v32 += v32
            v36 := GenS2{F0: 42, F1: 2, F2: 2}
            if (42 <= 100) {
                var v37 string = ""
                var v38 int32 = 42
                _ = v37
                _ = v38
            } else {
                if false {
                    if (!(!true)) {
                        switch 391 {
                        case 0:
                            fmt.Println(v2, v32)
                        }
                    }
                    var v39 float64 = float64(3.14159)
                    _ = v39
                }
            }
            _ = v36
        }
        var v40 int8 = 0
        _ = v32
        _ = v40
    }
    var v41 int8 = -128
    var v42 uint8 = 181
    v2 = (v2 / ((v2) | 1))
    for fi43 := 0; fi43 < 8; fi43++ {
        v4 = v4
        var v44 float32 = float32(v5)
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v45 bool = (-2147483648 == -134)
            v46 := GenS4{F0: 2147483648, F1: 0, F2: (v2 & v2)}
            _ = v45
            _ = v46
        } else {
            switch (^(fi43)) {
            case 32767:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -2147483648:
                if (!true) {
                    v4 = v4
                    var v47 string = "a"
                    _ = v47
                }
            case 4294967295:
                switch fi43 {
                case 372:
                    if len(v1) > 0 {
                        v1[0] = 170
                    }
                case 1:
                    var v48 bool = false
                    _ = v48
                }
            }
            v1 = append(v1, 2147483648)
            v49 := GenS2{F0: -932, F1: 338, F2: 26}
            _ = v49
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
