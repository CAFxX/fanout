package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int64 }
func genFn0(p0 bool, p1 uint8) uint16 {
    var v5 uint8 = 196
    var v6 int32 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint32, p1 bool) uint32 {
    switch 127 {
    case 32767:
        if ((false && p1) || (float32(1.5) >= (-2.25 / 3.14159))) {
            var v7 int64 = -9223372036854775808
            if (4294967295 <= 113) {
                v8 := []int{963, 127, 32768, 4294967295}
                v9 := GenS1{F0: 32768, F1: 32767, F2: v7}
                var v10 int64 = (v7 ^ v7)
                _ = v10
                _ = v8
                _ = v9
            } else {
                v11 := GenS1{F0: 581, F1: 42, F2: v7}
                _ = v11
            }
            _ = v7
        } else {
            v12 := []int{630, -9223372036854775808, 127, -9223372036854775808}
            var v13 string = (("a" + ("foo bar" + (("" + "foo bar") + ""))) + "")
            v14 := GenS1{F0: 32768, F1: 42, F2: 0}
            _ = v12
            _ = v13
            _ = v14
        }
    case -807:
        if ("a" != "") {
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 int16 = 1
                if p1 {
                    if (!("\n\t" == ("\n\t" + "foo bar"))) {
                        p1 = (!(!(!(!p1))))
                    } else {
                        var v17 float32 = 3.14159
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 1; fi18++ {
                        fmt.Println(p0)
                        _ = fi18
                    }
                }
                _ = fi15
                _ = v16
            }
            fmt.Println(p1, p1)
        } else {
            fmt.Println(p0)
        }
    case 4294967295:
        v19 := GenS1{F0: 4294967295, F1: 255, F2: -731}
        _ = v19
    }
    var v20 int8 = -1
    var v21 int16 = 32767
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return p0
}

func main() {
    var v1 int = -717
    v2 := make(map[string]string)
    var v3 int8 = -61
    var v4 float32 = 0.0
    switch "hello" {
    case "":
        var v22 uint16 = 65535
        _ = v22
    case "\n\t":
        var v23 bool = true
        _ = v23
    case "hello":
        v3 = (v3 + (v3 / (((v3 + -128)) | 1)))
    }
    v2["hello"] = "a"
    fmt.Println(v4)
    v1 -= v1
    fmt.Println(v1, v1, v3)
    delete(v2, ((("" + "a") + "foo bar") + ""))
    if (false || (v3 >= int8(v1))) {
        if false {
            v3 += 99
            if (!(4294967295 >= 128)) {
                fmt.Println(v4, v1, v4)
            } else {
                var v24 int8 = -128
                _ = v24
            }
            var v25 bool = false
            _ = v25
        } else {
            var v26 int8 = (^(v3))
            var v27 string = "\n\t"
            _ = v26
            _ = v27
        }
        v28 := []int{(v1 * 128), (v1 << ((491) & 63))}
        _ = v28
    } else {
        v29 := GenS1{F0: int(v3), F1: 65535, F2: 0}
        v30 := GenS1{F0: -619, F1: 916, F2: 42}
        _ = v29
        _ = v30
    }
    if (65535 < 4294967295) {
        var v31 int16 = 1
        _ = v31
    } else {
        v32 := GenS1{F0: (v1 >> ((2) & 63)), F1: 42, F2: 744}
        _ = v32
    }
    var v33 int = (v1 - v1)
    if false {
        v1 -= (^(v33))
    } else {
        if (!true) {
            if (int64(v3) >= -9223372036854775808) {
                v34 := GenS1{F0: (v33 % ((v33) | 1)), F1: genFn0(false, 127), F2: 443}
                _ = v34
            } else {
                var v35 int16 = 0
                switch "hello" {
                case "hello":
                    var v36 bool = false
                    _ = v36
                default:
                    fmt.Println(v3)
                }
                for fi37 := 0; fi37 < 4; fi37++ {
                    for fi38 := 0; fi38 < 12; fi38++ {
                        var v39 bool = ((0 != 32767) || true)
                        var v40 uint64 = 4294967295
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    var v41 string = "hello"
                    _ = fi37
                    _ = v41
                }
                _ = v35
            }
            switch 127 {
            case -2147483648:
                var v42 bool = false
                _ = v42
            case -9223372036854775808:
                var v43 uint = 35
                _ = v43
            }
        } else {
            fmt.Println(v33, v3)
            for fi44 := 0; fi44 < 13; fi44++ {
                v45 := GenS1{F0: (-(v33)), F1: 65535, F2: -9223372036854775808}
                var v46 int32 = 0
                _ = fi44
                _ = v45
                _ = v46
            }
        }
        v47 := GenS1{F0: v1, F1: 42, F2: 1}
        if true {
            for fi48 := 0; fi48 < 20; fi48++ {
                var v49 uint32 = 2147483647
                _ = fi48
                _ = v49
            }
        }
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
