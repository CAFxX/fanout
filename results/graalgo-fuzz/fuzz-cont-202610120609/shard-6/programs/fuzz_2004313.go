package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint8; F2 int }
type GenS3 struct { F0 bool; F1 bool; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 uint; F2 bool }
type GenS8 struct { F0 int8; F1 uint32 }
func genFn0() int16 {
    var v4 int8 = -1
    v5 := []int{42, 127, 4294967295, -725}
    var v6 int = -9223372036854775808
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1() int8 {
    var v7 float64 = 0.1
    v7 -= v7
    _ = v7
    return 47
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -370
    v3 := make(map[int]string)
    var v8 int = 9223372036854775807
    v2 = v2
    if (817 != (^(v2))) {
        v9 := GenS1{F0: 2147483647, F1: 0}
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 1; fi10++ {
            v2 -= v2
            _ = fi10
        }
    } else {
        delete(v3, v8)
    }
    v11 := []int{v8}
    if true {
        fmt.Println(v2, v8)
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        if true {
            v13 := GenS2{F0: (427 <= 0), F1: 255, F2: fi12}
            switch "" {
            case "":
                v2 = 42
            case "hello":
                switch -9223372036854775808 {
                case 2147483647:
                    v14 := GenS3{F0: false, F1: (((2147483648 <= 4294967295) && (v2 > 42)) && (1 < 127)), F2: -1}
                    _ = v14
                case 128:
                    v15 := GenS3{F0: true, F1: false, F2: -1}
                    _ = v15
                }
            case "foo bar":
                v16 := GenS3{F0: true, F1: (42 > 140), F2: 1}
                _ = v16
            }
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 int = fi17
                v19 := GenS4{F0: 1}
                v20 := v1[("hello" + "\n\t")]
                if v20 {
                    v18 += 2147483647
                    v21 := v1["a"]
                    if ("foo bar" != ("" + "\n\t")) {
                        fmt.Println(fi17)
                    } else {
                        fmt.Println(v18, v8)
                        if (((120 <= 1) && v21) || v20) {
                            v18 -= -9223372036854775808
                        } else {
                            v22 := GenS4{F0: 2147483647}
                            v23 := GenS5{F0: -1}
                            _ = v22
                            _ = v23
                        }
                        if v20 {
                            var v24 uint = 42
                            var v25 int32 = -549
                            if len(v11) > 0 {
                                v11[0] = 957
                            }
                            _ = v24
                            _ = v25
                        }
                    }
                    v8 += fi17
                    _ = v21
                } else {
                    var v26 uint16 = 2
                    fmt.Println(fi17, v2, v26)
                    v27 := GenS4{F0: -122}
                    _ = v26
                    _ = v27
                }
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = v13
        } else {
            var v28 int16 = 1000
            var v29 int32 = 1
            _ = v28
            _ = v29
        }
        v30 := GenS6{F0: 1}
        v8 -= (-(fi12))
        _ = fi12
        _ = v30
    }
    var v31 int8 = 0
    v32 := v1[(("foo bar" + "foo bar") + "\n\t")]
    v33 := GenS6{F0: v31}
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 float64 = float64((3.14159 / 3.14159))
        if true {
            v8 += (fi34 * (^(fi34)))
            var v36 int64 = (^(v2))
            _ = v36
        } else {
            switch "" {
            case "foo bar":
                fmt.Println(v32)
            }
            v2 = (^(v2))
            v37 := GenS2{F0: (4294967295 != 1), F1: 2, F2: (v8 % (((v8 * v8)) | 1))}
            _ = v37
        }
        v32 = v32
        _ = fi34
        _ = v35
    }
    var v38 int = v8
    fmt.Println(v2, v31)
    v38 += v38
    v39 := GenS3{F0: true, F1: true, F2: -1}
    v40 := GenS7{F0: 1, F1: 974, F2: false}
    for fi41 := 0; fi41 < 6; fi41++ {
        if ((v31 + v31) > 0) {
            if v32 {
                var v42 uint = 1
                switch "hello" {
                case "hello":
                    fmt.Println(fi41, v38, v32)
                case "":
                    v1["\n\t"] = false
                default:
                    delete(v3, (v8 - (v8 - fi41)))
                }
                if v32 {
                    v8 += (v8 ^ 4294967295)
                    v42 = (v42 % (((^(v42))) | 1))
                } else {
                    var v43 uint = (v42 - 18446744073709551615)
                    for fi44 := 0; fi44 < 12; fi44++ {
                        var v45 int16 = 39
                        v46 := GenS8{F0: genFn1(), F1: 2147483647}
                        var v47 uint = 1
                        v45 = int16(v31)
                        var v48 float32 = (float32(1e10) - (3.14159 / 3.14159))
                        v38 += v8
                        _ = fi44
                        _ = v45
                        _ = v46
                        _ = v47
                        _ = v48
                    }
                    _ = v43
                }
                _ = v42
            }
        } else {
            var v49 int32 = 2147483647
            var v50 float64 = (1e10 - 0.0)
            for fi51 := 0; fi51 < 7; fi51++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi51
            }
            _ = v49
            _ = v50
        }
        if len(v11) > 0 {
            v11[0] = fi41
        }
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v39
    _ = v40
    _ = v8
}
