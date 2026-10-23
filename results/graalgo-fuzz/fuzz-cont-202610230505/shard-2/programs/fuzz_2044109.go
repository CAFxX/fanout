package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int16; F1 int32; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint64 = 129
    v2 := make(map[string]bool)
    var v3 uint32 = 641
    if true {
        if (((1 < -1) || ((v3 | v3) <= (v3 + v3))) || true) {
            fmt.Println(v3, v1)
            fmt.Println(v1, v1)
        } else {
            switch 128 {
            case -1:
                switch "foo bar" {
                case "a":
                    for fi4 := 0; fi4 < 2; fi4++ {
                        v5 := GenS1{F0: 2, F1: 4}
                        var v6 int32 = -2147483648
                        for fi7 := 0; fi7 < 2; fi7++ {
                            var v8 int8 = 3
                            switch ("hello" + "a") {
                            case "":
                                v9 := GenS1{F0: 1, F1: (v1 % ((v1) | 1))}
                                _ = v9
                            case "a":
                                v8 = (v8 | (v8 >> ((2) & 63)))
                                fmt.Println(v3, fi7, fi7)
                                if false {
                                    if (0 >= -848) {
                                        v10 := v2["\n\t"]
                                        v11 := GenS1{F0: 2, F1: (v1 % ((uint64(v3)) | 1))}
                                        _ = v10
                                        _ = v11
                                    }
                                    if false {
                                        v12 := []int{(fi4 ^ 42), (fi7 | fi4)}
                                        var v13 uint = 42
                                        _ = v12
                                        _ = v13
                                    }
                                    v1 += v1
                                    if (true || false) {
                                        var v14 uint16 = 65535
                                        v15 := []int{2147483647}
                                        v2["hello"] = (0 != 0)
                                        _ = v14
                                        _ = v15
                                    }
                                    if false {
                                        delete(v2, "hello")
                                    } else {
                                        delete(v2, ("" + ""))
                                        v1 += 401
                                    }
                                    v16 := GenS2{F0: 586, F1: v6, F2: v8}
                                    v17 := v2[("" + "foo bar")]
                                    _ = v16
                                    _ = v17
                                } else {
                                    v18 := GenS2{F0: 32767, F1: (v6 & v6), F2: (v8 & (v8 << ((2) & 63)))}
                                    var v19 uint32 = 2
                                    _ = v18
                                    _ = v19
                                }
                            }
                            fmt.Println(v8, v8, v1)
                            _ = fi7
                            _ = v8
                        }
                        _ = fi4
                        _ = v5
                        _ = v6
                    }
                case "\n\t":
                    for fi20 := 0; fi20 < 12; fi20++ {
                        if true {
                            var v21 float64 = 3.14159
                            _ = v21
                        } else {
                            var v22 int32 = -949
                            if (!false) {
                                v23 := []int{}
                                _ = v23
                            }
                            var v24 uint16 = 255
                            _ = v22
                            _ = v24
                        }
                        _ = fi20
                    }
                }
            case -2147483648:
                var v25 int32 = 703
                _ = v25
            case 32768:
                var v26 int32 = 192
                _ = v26
            }
            v27 := []int{}
            _ = v27
        }
        if (-32768 == -32768) {
            var v28 string = (("a" + "") + "foo bar")
            v3 = (v3 - 65535)
            _ = v28
        }
        fmt.Println(v1, v1)
    }
    v29 := []int{192, -9223372036854775808, 127, -1}
    fmt.Println(v3, v1)
    v1 += 0
    var v30 uint = 1
    v31 := GenS3{F0: "hello"}
    var v32 int = 2147483647
    var v33 bool = (false && true)
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 bool = false
        _ = fi34
        _ = v35
    }
    fmt.Println(v32, v30)
    fmt.Println(v1, v30)
    var v36 float64 = 0.0
    for fi37 := 0; fi37 < 7; fi37++ {
        if v33 {
            var v38 int = 786
            for fi39 := 0; fi39 < 18; fi39++ {
                if ((1.5 >= 3.14159) && (!(v33 && false))) {
                    switch v38 {
                    case 127:
                        var v40 uint = (v30 | (^(v30)))
                        _ = v40
                    }
                    v1 += v1
                    v41 := GenS4{F0: 235, F1: -2147483648}
                    _ = v41
                } else {
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                    var v42 float64 = -2.25
                    if false {
                        _ = len(v29)
                        v43 := GenS3{F0: "a"}
                        if v33 {
                            v38 += -2147483648
                        } else {
                            v36 += ((1e10 + v42) / float64(0.1))
                            var v44 int32 = 42
                            v32 -= v32
                            _ = v44
                        }
                        _ = v43
                    }
                    switch "" {
                    case "hello":
                        if (!false) {
                            switch 32767 {
                            case 4294967295:
                                var v45 float64 = v42
                                _ = v45
                            default:
                                var v46 uint32 = (v3 % (((v3 & (v3 & 65535))) | 1))
                                _ = v46
                            }
                        }
                        var v47 bool = ((v30 << ((v30) & 63)) != (v30 & 42))
                        _ = v47
                    default:
                        fmt.Println(v36, v30)
                    }
                    _ = len(v29)
                    _ = v42
                }
                _ = fi39
            }
            _ = v38
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v36
}
