package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
type GenS4 struct { F0 string; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int32, p1 int32) int16 {
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 uint32 = 0
        fmt.Println(v6, p0, p1)
        _ = fi5
        _ = v6
    }
    p1 = p1
    p0 += 42
    switch "hello" {
    case "\n\t":
        if ((true && false) || true) {
            var v7 bool = (false || true)
            p1 -= p0
            _ = v7
        } else {
            var v8 uint64 = 710
            _ = v8
        }
    default:
        p0 = p0
    }
    p0 = p0
    _ = p0
    _ = p1
    return -32768
}
func genFn1(p0 uint32, p1 int32, p2 uint) uint {
    v9 := []int{127, 567, 0}
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(p1)
        var v11 uint32 = (p0 | (^(p0)))
        p2 -= (p2 / ((p2) | 1))
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p2 / (((p2 % (((^(p2))) | 1))) | 1))
}

func main() {
    v1 := []uint32{4294967295, 547, 65535}
    var v2 int16 = 1000
    v3 := make(map[int]string)
    var v4 int64 = -1
    v12 := GenS1{F0: 42, F1: 1, F2: 934}
    _ = len(v1)
    v13 := GenS1{F0: 1, F1: 246, F2: -374}
    var v14 int8 = 42
    if false {
        v2 = (v2 ^ v2)
    } else {
        switch -1 {
        case -9223372036854775808:
            v15 := GenS1{F0: 1, F1: -148, F2: int32(v4)}
            _ = v15
        }
    }
    if ("a" <= ("\n\t" + ("a" + "foo bar"))) {
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int = int(v2)
            var v18 int = 42
            v19 := GenS1{F0: 18446744073709551615, F1: -627, F2: 702}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS2{F0: 2, F1: -2147483648}
            _ = fi20
            _ = v21
        }
        v14 += v14
    } else {
        v2 = -853
        var v22 uint8 = 2
        v22 -= (v22 | (v22 ^ 42))
        _ = v22
    }
    var v23 uint = 18446744073709551615
    v24 := v3[105]
    switch "" {
    case "a":
        switch v24 {
        case "a":
            v3[32767] = v24
        }
    case "":
        for fi25 := 0; fi25 < 16; fi25++ {
            switch (fi25 % ((0) | 1)) {
            case 128:
                for fi26 := 0; fi26 < 2; fi26++ {
                    v4 = (^(v4))
                    if true {
                        switch v24 {
                        case "\n\t":
                            switch 610 {
                            case 394:
                                v23 = (v23 / (((v23 | v23)) | 1))
                            }
                        case "":
                            switch 32767 {
                            case -1:
                                fmt.Println(v4, v2, v4)
                            case 42:
                                switch v24 {
                                case "a":
                                    var v27 uint8 = 116
                                    _ = v27
                                case "hello":
                                    if (!((619 != (^(v4))) || (4294967295 == 720))) {
                                        v23 += v23
                                        var v28 uint8 = 104
                                        _ = v28
                                    }
                                case "\n\t":
                                    if (!(1.5 != (-2.25 * (0.1 / ((float32(3.14159) + (-2.25 - 1.5)) * 1e10))))) {
                                        v4 = v4
                                    } else {
                                        v24 = (v24 + (v24 + v24))
                                        v29 := GenS3{F0: 4294967295, F1: fi25, F2: 1}
                                        switch fi25 {
                                        case 42:
                                            v3[fi26] = v24
                                        default:
                                            var v30 float32 = 0.1
                                            _ = v30
                                        }
                                        _ = v29
                                    }
                                default:
                                    v31 := GenS1{F0: 498, F1: 1, F2: 0}
                                    _ = v31
                                }
                            case 1:
                                var v32 float32 = (0.1 * 3.14159)
                                _ = v32
                            }
                        case "a":
                            if ("hello" > (v24 + v24)) {
                                var v33 uint8 = 1
                                v14 = (^(v14))
                                v24 = "foo bar"
                                _ = v33
                            } else {
                                if false {
                                    var v34 uint32 = 2
                                    v35 := GenS2{F0: 127, F1: 1}
                                    _ = v34
                                    _ = v35
                                }
                                if (false || (((v14 / (((v14 * v14)) | 1)) != 42) && (v4 <= int64(v2)))) {
                                    var v36 uint32 = 42
                                    var v37 int16 = (v2 | (v2 * 0))
                                    _ = v36
                                    _ = v37
                                }
                                v38 := GenS2{F0: 1, F1: 0}
                                v39 := GenS4{F0: v24, F1: -724, F2: 2}
                                var v40 float32 = 0.0
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                        default:
                            var v41 bool = true
                            _ = v41
                        }
                        if ("a" <= v24) {
                            var v42 int8 = -107
                            var v43 string = (("hello" + (((v24 + v24) + "a") + "foo bar")) + v24)
                            _ = v42
                            _ = v43
                        } else {
                            v24 = v24
                            var v44 int32 = -1
                            _ = v44
                        }
                        v45 := GenS4{F0: v24, F1: -1, F2: 608}
                        _ = v45
                    }
                    v46 := v3[-1]
                    v24 = ("a" + v46)
                    v24 = v46
                    delete(v3, (fi25 + (fi25 * (fi25 << ((v23) & 63)))))
                    v24 = (v24 + ((v46 + v46) + v46))
                    v47 := GenS5{F0: 127}
                    _ = fi26
                    _ = v46
                    _ = v47
                }
            }
            var v48 int64 = v4
            v2 = v2
            _ = fi25
            _ = v48
        }
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v49 float32 = (float32(((((1.5 * 3.14159) + 0.0) + 3.14159) + 0.0)) - -2.25)
        _ = v49
    }
    v14 -= (v14 + (-(v14)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
