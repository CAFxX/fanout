package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint16) uint64 {
    var v6 bool = (!((!(42 != p0)) && true))
    fmt.Println(p0, p0)
    if true {
        var v7 float64 = 1.5
        v8 := GenS1{F0: 9223372036854775807, F1: 127}
        _ = v7
        _ = v8
    } else {
        if (!true) {
            v6 = v6
            p0 += p0
            fmt.Println(p0)
        } else {
            var v9 int8 = 127
            v10 := GenS1{F0: -1, F1: v9}
            _ = v10
            _ = v9
        }
    }
    v6 = (!(false || v6))
    switch 826 {
    case 32768:
        v11 := GenS1{F0: 32768, F1: 1}
        _ = v11
    case 42:
        var v12 int16 = -192
        _ = v12
    case 293:
        switch (((("\n\t" + (("a" + "") + "hello")) + ("\n\t" + "\n\t")) + "\n\t") + (("" + ("a" + "hello")) + "\n\t")) {
        case "foo bar":
            var v13 uint64 = 18446744073709551615
            _ = v13
        case "a":
            p0 = p0
        case "":
            switch 127 {
            case -9223372036854775808:
                fmt.Println(v6)
            case 4294967295:
                var v14 uint8 = 42
                _ = v14
            }
        }
    default:
        p0 += p0
    }
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 int8 = 0
    var v2 string = "foo bar"
    v3 := []string{(v2 + v2)}
    v4 := []uint32{4294967295}
    var v5 float32 = 1.5
    fmt.Println(v5, v5, v1)
    fmt.Println(v2, v5, v1)
    if false {
        if true {
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 int = 1
                    var v18 uint = 991
                    var v19 float64 = (1e10 + 3.14159)
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v20 := GenS1{F0: (fi15 >> ((255) & 63)), F1: 1}
                _ = fi15
                _ = v20
            }
            var v21 string = v2
            if len(v4) > 0 {
                v4[0] = 428
            }
            _ = v21
        } else {
            switch 32768 {
            case 127:
                v22 := GenS2{F0: v1, F1: 209}
                _ = v22
            case 32767:
                var v23 int = 32767
                _ = v23
            case -792:
                if true {
                    var v24 float64 = 1.5
                    if false {
                        var v25 int32 = -1
                        _ = v25
                    }
                    _ = v24
                }
            default:
                v1 -= -128
            }
            switch "hello" {
            case "foo bar":
                var v26 string = "\n\t"
                _ = v26
            default:
                var v27 int32 = -1
                _ = v27
            }
            if true {
                var v28 bool = false
                v1 += (v1 - v1)
                for fi29 := 0; fi29 < 13; fi29++ {
                    v30 := GenS2{F0: (-(v1)), F1: 0}
                    v31 := GenS3{F0: 0}
                    if true {
                        v4 = append(v4, 1)
                        v3 = append(v3, (v2 + v2))
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = v28
            } else {
                var v32 uint64 = 4294967295
                _ = v32
            }
        }
        var v33 int32 = 1
        _ = v33
    }
    if (198 <= 507) {
        if true {
            v34 := GenS4{F0: 823}
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        fmt.Println(v2, fi35, fi35)
        _ = fi35
    }
    v36 := GenS5{F0: 2}
    v37 := GenS6{F0: 42}
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 int16 = 42
        if false {
            var v40 float32 = 0.1
            _ = v40
        } else {
            for fi41 := 0; fi41 < 15; fi41++ {
                if (42 != 4294967295) {
                    var v42 int = 128
                    var v43 int32 = -2147483648
                    v39 = (v39 & (v39 % ((v39) | 1)))
                    _ = v42
                    _ = v43
                } else {
                    var v44 int8 = v1
                    var v45 bool = false
                    v3 = append(v3, v2)
                    var v46 int8 = (v1 >> ((0) & 63))
                    var v47 int8 = 42
                    _ = v44
                    _ = v45
                    _ = v46
                    _ = v47
                }
                v39 = (v39 >> ((65535) & 63))
                var v48 uint = 0
                _ = fi41
                _ = v48
            }
            v39 = (^(v39))
            for fi49 := 0; fi49 < 19; fi49++ {
                if ((-2.25 < 0.0) || false) {
                    fmt.Println(v5, v39, v1)
                    if false {
                        v2 = (("\n\t" + v2) + (v2 + ((v2 + "") + "foo bar")))
                        fmt.Println(fi49, fi38)
                    }
                    switch "" {
                    case "a":
                        v50 := GenS2{F0: (v1 ^ 42), F1: (v39 & v39)}
                        _ = v50
                    case "":
                        v39 -= 32767
                        v1 -= v1
                        var v51 bool = (1 >= (v1 >> ((898) & 63)))
                        _ = v51
                    case "hello":
                        if ((-236 > 9223372036854775807) && (!(-705 != -9223372036854775808))) {
                            fmt.Println(fi49, v5, v5)
                        }
                    default:
                        if (((2 < 4294967295) || true) && false) {
                            v5 -= v5
                        }
                    }
                }
                _ = fi49
            }
        }
        _ = fi38
        _ = v39
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
