package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint32, p1 string, p2 uint16) int64 {
    fmt.Println(p1, p0, p2)
    v7 := []int{-2147483648, 32767, 4294967295, 764}
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: "hello", F1: -1, F2: 1}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: (p1 + (((p1 + "\n\t") + p1) + p1)), F1: -1, F2: 4}
    var v11 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    return 186
}
func genFn1(p0 uint16, p1 string) int32 {
    v12 := GenS1{F0: (p1 + p1), F1: 0, F2: 1000}
    p0 -= (p0 * 255)
    p0 = (p0 - (p0 + p0))
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    v3 := []int{}
    var v4 int16 = 1
    var v5 bool = false
    v6 := []string{"hello", "", ""}
    var v13 int8 = -1
    var v14 string = "a"
    v15 := GenS2{F0: 1, F1: 42}
    fmt.Println(v14, v5)
    switch 32768 {
    case -152:
        if (float64((1e10 / (1.5 * float32(-2.25)))) == 0.1) {
            v2[-9223372036854775808] = 42
            v4 += (v4 & (v4 - v4))
        } else {
            fmt.Println(v5, v5)
            if (!true) {
                var v16 uint32 = 4294967295
                _ = v16
            } else {
                for fi17 := 0; fi17 < 17; fi17++ {
                    v1[(fi17 | (fi17 | -2147483648))] = fi17
                    if ((v5 && ((!v5) && (v5 || true))) && (63 != (-(v13)))) {
                        v5 = true
                        var v18 bool = (!v5)
                        var v19 int32 = 624
                        _ = v18
                        _ = v19
                    } else {
                        fmt.Println(v13, v5, v4)
                    }
                    v5 = true
                    _ = fi17
                }
            }
        }
    }
    fmt.Println(v4)
    fmt.Println(v13)
    v5 = v5
    fmt.Println(v4, v5)
    v20 := GenS3{F0: -321}
    v21 := GenS3{F0: -549}
    if (!v5) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v13 += 127
        switch 32767 {
        case 32768:
            _ = len(v6)
        case -9223372036854775808:
            if v5 {
                if len(v3) > 0 {
                    v3[0] = 815
                }
                v4 = (-(v4))
                switch (v14 + v14) {
                case "a":
                    for fi22 := 0; fi22 < 20; fi22++ {
                        for fi23 := 0; fi23 < 19; fi23++ {
                            v5 = false
                            _ = fi23
                        }
                        _ = fi22
                    }
                case "":
                    for fi24 := 0; fi24 < 14; fi24++ {
                        var v25 float64 = (float64(0.0) / 1.0)
                        v13 = (^(v13))
                        _ = fi24
                        _ = v25
                    }
                case "foo bar":
                    var v26 int = -244
                    _ = v26
                default:
                    var v27 int16 = int16(v13)
                    _ = v27
                }
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        case 451:
            if ((v13 << ((18446744073709551615) & 63)) != v13) {
                v28 := GenS3{F0: 9223372036854775807}
                _ = v28
            }
        default:
            var v29 int16 = v4
            _ = v29
        }
    } else {
        v13 -= (^(v13))
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v14 = v14
        if len(v6) > 0 {
            v6[0] = v14
        }
        _ = fi30
    }
    var v31 int32 = 0
    for fi32 := 0; fi32 < 3; fi32++ {
        switch fi32 {
        case 32768:
            switch (^(fi32)) {
            case -930:
                for fi33 := 0; fi33 < 5; fi33++ {
                    delete(v1, 9223372036854775807)
                    _ = fi33
                }
            default:
                delete(v2, fi32)
            }
        case -758:
            var v34 uint64 = 4294967295
            _ = v34
        case 4294967295:
            if (fi32 < 32767) {
                if false {
                    delete(v1, (fi32 >> ((65535) & 63)))
                    v35 := v1[(fi32 - fi32)]
                    fmt.Println(v31, v4, fi32)
                    _ = v35
                } else {
                    v36 := GenS4{F0: 41, F1: 2147483647}
                    for fi37 := 0; fi37 < 12; fi37++ {
                        _ = len(v1)
                        v31 = 2147483647
                        fmt.Println(fi37, v14, v31)
                        v2[32767] = 1
                        _ = fi37
                    }
                    v38 := GenS2{F0: 65535, F1: 0}
                    _ = v36
                    _ = v38
                }
            } else {
                var v39 int = fi32
                var v40 int16 = 65
                _ = v39
                _ = v40
            }
        default:
            for fi41 := 0; fi41 < 6; fi41++ {
                switch ("" + v14) {
                case "foo bar":
                    if v5 {
                        v31 -= -2147483648
                        switch ((v14 + ("a" + (v14 + (v14 + v14)))) + "hello") {
                        case "a":
                            var v42 int = (fi41 ^ 720)
                            _ = v42
                        }
                    }
                case "":
                    var v43 int64 = -9223372036854775808
                    _ = v43
                case "a":
                    var v44 uint32 = 967
                    _ = v44
                default:
                    v31 += genFn1(1, (v14 + ""))
                }
                v45 := GenS1{F0: "\n\t", F1: (-(v13)), F2: (v4 & v4)}
                switch "\n\t" {
                case "foo bar":
                    var v46 uint16 = 65535
                    _ = v46
                case "":
                    _ = len(v3)
                }
                _ = fi41
                _ = v45
            }
        }
        var v47 int16 = v4
        _ = fi32
        _ = v47
    }
    var v48 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
