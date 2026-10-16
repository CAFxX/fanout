package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 string }

func main() {
    var v1 uint8 = 1
    v2 := []uint32{}
    var v3 uint64 = 569
    var v4 uint = 42
    var v5 string = ("a" + "foo bar")
    if ((!false) && false) {
        v6 := GenS1{F0: -123, F1: 42, F2: false}
        if false {
            var v7 string = ""
            _ = v7
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            switch fi8 {
            case 127:
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 uint16 = 193
                    fmt.Println(v1, v5, fi8)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi9
                    _ = v10
                }
            case 61:
                v3 = v3
            }
            var v11 string = v5
            _ = fi8
            _ = v11
        }
        _ = v6
    } else {
        var v12 int = 936
        var v13 string = (v5 + v5)
        var v14 float32 = (0.0 - 0.0)
        _ = v12
        _ = v13
        _ = v14
    }
    v3 = v3
    var v15 uint8 = uint8(v3)
    if (false || true) {
        v15 += 42
        v5 = (v5 + v5)
        switch 127 {
        case -393:
            var v16 uint8 = 0
            _ = v16
        case 9223372036854775807:
            fmt.Println(v3, v4, v4)
        case 0:
            switch (v5 + v5) {
            case "\n\t":
                var v17 uint = 2
                _ = v17
            }
        }
    }
    fmt.Println(v3)
    if (0.1 == 3.14159) {
        v2 = append(v2, 1)
    } else {
        var v18 int64 = 42
        switch -1 {
        case 878:
            v19 := GenS1{F0: 1, F1: v18, F2: true}
            _ = v19
        case 128:
            v2 = append(v2, 42)
        case 0:
            var v20 float32 = ((1.5 - (float32((0.1 * (0.1 - (-2.25 * 1.5)))) - (-2.25 + 0.0))) - (-2.25 + -2.25))
            _ = v20
        }
        _ = v18
    }
    v15 += v1
    v3 += v3
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 int = fi21
        if true {
            switch (v22 % (((fi21 % (((v22 << (((v4 & (^(v4)))) & 63))) | 1))) | 1)) {
            case 350:
                for fi23 := 0; fi23 < 16; fi23++ {
                    if (0.1 == (3.14159 - 1e10)) {
                        var v24 uint64 = 18446744073709551615
                        v25 := GenS1{F0: -57, F1: -1, F2: false}
                        _ = v24
                        _ = v25
                    }
                    v26 := GenS1{F0: 100, F1: 42, F2: (false && true)}
                    var v27 int = 4294967295
                    v3 += v3
                    _ = len(v2)
                    _ = fi23
                    _ = v26
                    _ = v27
                }
            case 4294967295:
                if (!true) {
                    v22 += fi21
                } else {
                    var v28 bool = ((v3 - 1) <= v3)
                    v29 := GenS1{F0: 101, F1: 680, F2: false}
                    var v30 bool = true
                    _ = v28
                    _ = v29
                    _ = v30
                }
            default:
                var v31 bool = true
                _ = v31
            }
            switch (("foo bar" + v5) + ("hello" + "foo bar")) {
            case "\n\t":
                v2 = append(v2, 0)
            default:
                fmt.Println(fi21, v22)
            }
            v1 += v1
        } else {
            var v32 float32 = 1e10
            if true {
                if ((false && false) && ((false || false) || (!true))) {
                    var v33 uint = 2
                    _ = v33
                } else {
                    if (v15 < v15) {
                        for fi34 := 0; fi34 < 2; fi34++ {
                            fmt.Println(v4)
                            var v35 uint = v4
                            _ = fi34
                            _ = v35
                        }
                        v3 = v3
                        v36 := GenS2{F0: (v3 << ((18446744073709551615) & 63)), F1: false, F2: (v15 >> ((v4) & 63))}
                        _ = v36
                    }
                    var v37 uint32 = 2147483648
                    _ = v37
                }
                var v38 uint32 = 42
                if false {
                    fmt.Println(v32, v15, fi21)
                } else {
                    for fi39 := 0; fi39 < 1; fi39++ {
                        v32 = 0.0
                        v1 -= v1
                        var v40 uint32 = (v38 + uint32(v4))
                        _ = fi39
                        _ = v40
                    }
                    var v41 float32 = v32
                    _ = v41
                }
                _ = v38
            } else {
                v42 := GenS3{F0: v3}
                var v43 uint64 = (^(v3))
                _ = v42
                _ = v43
            }
            v44 := GenS4{F0: (v1 / (((^(v15))) | 1))}
            _ = v32
            _ = v44
        }
        _ = fi21
        _ = v22
    }
    if false {
        var v45 int32 = int32(v15)
        var v46 uint16 = 2
        _ = v45
        _ = v46
    } else {
        v1 = v15
    }
    fmt.Println(v3, v1)
    v47 := GenS5{F0: 357}
    _ = len(v2)
    v15 += v1
    if (!true) {
        v48 := GenS6{F0: 42, F1: (v5 + v5)}
        if false {
            v3 -= v3
            v49 := GenS2{F0: v3, F1: (42 <= 2147483647), F2: v1}
            _ = v49
        } else {
            for fi50 := 0; fi50 < 11; fi50++ {
                v5 = ((v5 + "hello") + v5)
                var v51 uint64 = v3
                var v52 uint16 = 32767
                _ = fi50
                _ = v51
                _ = v52
            }
        }
        var v53 uint16 = 788
        _ = v48
        _ = v53
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        for fi54 := 0; fi54 < 12; fi54++ {
            var v55 uint16 = 2
            for fi56 := 0; fi56 < 9; fi56++ {
                fmt.Println(v5)
                _ = fi56
            }
            var v57 string = ((v5 + "a") + v5)
            _ = fi54
            _ = v55
            _ = v57
        }
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v3, v4, v4)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
