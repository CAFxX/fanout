package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 uint, p1 uint16) int32 {
    v5 := []int{}
    if false {
        v6 := GenS1{F0: 42}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        p1 += (p1 << (((p0 % ((65535) | 1))) & 63))
        _ = v6
    } else {
        var v7 int16 = 221
        p0 = (^(p0))
        p0 = uint(p1)
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        switch fi8 {
        case 32768:
            var v9 float64 = 1e10
            _ = v9
        default:
            if (4294967295 != 490) {
                if false {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v10 bool = true
                    _ = v10
                }
                if (!false) {
                    v11 := GenS1{F0: -1}
                    var v12 int32 = int32(p0)
                    if ((0.1 - (3.14159 * 1.5)) != 1.5) {
                        var v13 float64 = (3.14159 * 0.0)
                        for fi14 := 0; fi14 < 20; fi14++ {
                            p1 += (^(p1))
                            _ = fi14
                        }
                        switch fi8 {
                        case -2147483648:
                            var v15 float32 = -2.25
                            _ = v15
                        }
                        _ = v13
                    } else {
                        var v16 float64 = (((float64(0.1) - 0.0) + 3.14159) + 0.1)
                        fmt.Println(v16, v12)
                        if len(v5) > 0 {
                            v5[0] = fi8
                        }
                        _ = v16
                    }
                    _ = v11
                    _ = v12
                } else {
                    for fi17 := 0; fi17 < 7; fi17++ {
                        p0 = 1
                        v18 := GenS1{F0: -1}
                        fmt.Println(p0, p1, fi17)
                        _ = fi17
                        _ = v18
                    }
                }
                switch -786 {
                case 128:
                    p1 = 32767
                case -1:
                    var v19 uint8 = 56
                    _ = v19
                default:
                    var v20 float32 = 1e10
                    _ = v20
                }
            } else {
                p0 = (p0 + p0)
                p1 = 520
                var v21 int = (fi8 << (((^(p0))) & 63))
                _ = v21
            }
        }
        v22 := GenS1{F0: 9223372036854775807}
        _ = fi8
        _ = v22
    }
    p1 = 42
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 string, p1 int16) int {
    p1 -= p1
    fmt.Println(p1)
    if false {
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS1{F0: 9223372036854775807}
            for fi25 := 0; fi25 < 17; fi25++ {
                switch -68 {
                case 9223372036854775807:
                    v26 := []int{(fi25 << ((153) & 63)), (fi23 % ((127) | 1))}
                    _ = v26
                }
                if (15 >= -128) {
                    p0 = "foo bar"
                }
                var v27 bool = true
                _ = fi25
                _ = v27
            }
            _ = fi23
            _ = v24
        }
    } else {
        p1 = 42
        if (true && true) {
            v28 := GenS2{F0: false, F1: 1, F2: 4294967295}
            _ = v28
        }
    }
    v29 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v29
    return 736
}

func main() {
    var v1 string = ("a" + "a")
    var v2 int = 2147483647
    var v3 int = (v2 / (((v2 * v2)) | 1))
    var v4 float64 = float64(3.14159)
    fmt.Println(v1, v1, v4)
    v30 := GenS2{F0: (1 >= 42), F1: -9223372036854775808, F2: 245}
    var v31 bool = (0 == 0)
    v2 = (v3 ^ (v2 * 4294967295))
    for fi32 := 0; fi32 < 7; fi32++ {
        switch v3 {
        case -2147483648:
            var v33 int16 = 32767
            _ = v33
        case 2147483647:
            v34 := []int{(v3 >> ((4294967295) & 63)), (-(fi32)), (v3 ^ (fi32 + (-(fi32))))}
            _ = v34
        case -112:
            v35 := []int{fi32, v3, (v2 >> ((255) & 63))}
            _ = v35
        default:
            v3 += genFn1("a", 1000)
        }
        v4 = v4
        v4 -= v4
        _ = fi32
    }
    v36 := []int{genFn1(v1, 42), (v3 - (^(v2))), v3}
    switch ((v1 + "\n\t") + (((v1 + v1) + ("a" + "a")) + "foo bar")) {
    case "a":
        switch v3 {
        case -2147483648:
            var v37 int32 = -1
            _ = v37
        default:
            fmt.Println(v1)
        }
    }
    _ = len(v36)
    switch genFn1("\n\t", -480) {
    case 128:
        var v38 uint = 915
        _ = v38
    case 619:
        _ = len(v36)
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        v3 += genFn1(("foo bar" + (v1 + (v1 + v1))), 0)
        switch -1 {
        case -9223372036854775808:
            v40 := GenS3{F0: false, F1: v1}
            _ = v40
        case -1:
            var v41 string = v1
            _ = v41
        case 2147483647:
            v3 -= 2147483647
        }
        _ = fi39
    }
    var v42 uint8 = 241
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    if (!v31) {
        v1 = v1
    } else {
        var v43 int32 = 42
        _ = v43
    }
    switch (v2 & 32768) {
    case -995:
        var v44 int = v3
        _ = v44
    case 127:
        if true {
            var v45 uint8 = (v42 ^ (v42 & v42))
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
            _ = v45
        } else {
            var v46 int = v3
            if (float32(-2.25) != (float32(v4) + 0.0)) {
                if (!((2 != 0) && v31)) {
                    v2 += -9223372036854775808
                    v2 -= v46
                    v31 = v31
                } else {
                    if ((0.1 + v4) >= 0.0) {
                        fmt.Println(v4, v31, v2)
                    }
                    fmt.Println(v4)
                }
            } else {
                v1 = v1
                if len(v36) >= 2 {
                    v36 = v36[1:len(v36)]
                }
                v3 -= genFn1((((v1 + ("a" + "hello")) + v1) + (v1 + ("a" + v1))), -32768)
            }
            _ = v46
        }
    }
    if v31 {
        if v31 {
            for fi47 := 0; fi47 < 20; fi47++ {
                var v48 float64 = 1.5
                _ = fi47
                _ = v48
            }
        } else {
            fmt.Println(v1, v4)
            if false {
                var v49 int64 = -755
                switch (v1 + (("a" + (v1 + "foo bar")) + (((v1 + v1) + "\n\t") + ""))) {
                case "\n\t":
                    var v50 int64 = (v49 * int64(v2))
                    _ = v50
                default:
                    _ = len(v36)
                }
                _ = v49
            } else {
                var v51 uint8 = (^(v42))
                v31 = (!(211 != -1))
                for fi52 := 0; fi52 < 15; fi52++ {
                    var v53 int = (v2 | fi52)
                    _ = fi52
                    _ = v53
                }
                _ = v51
            }
        }
    }
    var v54 int32 = 860
    var v55 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v4
    _ = v42
    _ = v54
    _ = v55
}
