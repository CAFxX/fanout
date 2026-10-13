package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS4 struct { F0 int16; F1 int8; F2 int64 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int, p1 uint16) int8 {
    v7 := GenS1{F0: -233, F1: (55 != 255), F2: 42}
    if (42 > 0) {
        var v8 int = (p0 * 0)
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint = 4294967295
            var v11 int8 = 127
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint16 = 227
        _ = v12
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return -1
}

func main() {
    var v1 uint8 = 145
    var v2 uint = 65535
    v3 := make(map[string]uint32)
    var v4 uint32 = 4294967295
    v5 := []string{"hello"}
    var v6 string = ""
    v13 := GenS2{F0: 4294967295}
    v1 = v1
    if (!false) {
        switch ((((v6 + "") + "hello") + "a") + v6) {
        case "hello":
            if true {
                v14 := GenS3{F0: (v4 >> ((v2) & 63)), F1: 350, F2: 42}
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS4{F0: 393, F1: 1, F2: int64(v1)}
                    var v17 int = (fi15 * 350)
                    var v18 int8 = -128
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v14
            }
        case "\n\t":
            v4 -= 2147483647
        }
    }
    fmt.Println(v2)
    var v19 int = 32767
    switch ((v6 + v6) + v6) {
    case "\n\t":
        fmt.Println(v1, v4)
    case "foo bar":
        switch v6 {
        case "a":
            var v20 uint = (v2 >> ((v2) & 63))
            _ = v20
        case "\n\t":
            for fi21 := 0; fi21 < 14; fi21++ {
                if (!(2 != 809)) {
                    if ((v2 * 144) >= v2) {
                        for fi22 := 0; fi22 < 11; fi22++ {
                            v6 = ("\n\t" + "hello")
                            _ = fi22
                        }
                        if (uint64(v2) == 4294967295) {
                            var v23 bool = (-1 != 42)
                            v1 -= v1
                            _ = v23
                        }
                    } else {
                        var v24 float32 = -2.25
                        _ = v24
                    }
                    fmt.Println(v6)
                    v25 := GenS4{F0: 215, F1: 42, F2: 1}
                    _ = v25
                } else {
                    v6 = v6
                }
                v2 += (v2 ^ v2)
                _ = fi21
            }
        case "foo bar":
            if false {
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v6, v4)
                    default:
                        v26 := GenS1{F0: 0, F1: false, F2: v2}
                        _ = v26
                    }
                    fmt.Println(v19, v1)
                    var v27 uint64 = 213
                    _ = v27
                }
            }
        }
    default:
        var v28 bool = true
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 int = fi29
        if true {
            if (264 < 631) {
                var v31 uint8 = (v1 ^ v1)
                for fi32 := 0; fi32 < 4; fi32++ {
                    v31 += uint8(v4)
                    var v33 uint16 = 0
                    var v34 int8 = 64
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                switch (((v6 + ((v6 + "foo bar") + (("foo bar" + v6) + (v6 + "")))) + v6) + "hello") {
                case "":
                    v4 = (v4 - v4)
                default:
                    v2 = (^(v2))
                }
                _ = v31
            }
        } else {
            var v35 int16 = 1000
            _ = v35
        }
        for fi36 := 0; fi36 < 13; fi36++ {
            var v37 uint16 = 464
            v38 := GenS4{F0: 37, F1: -1, F2: 9223372036854775807}
            if true {
                _ = len(v5)
                if (true || true) {
                    fmt.Println(v37)
                } else {
                    v37 += v37
                    var v39 uint8 = (v1 - v1)
                    _ = v39
                }
            } else {
                delete(v3, ("foo bar" + (((v6 + "a") + (("" + "\n\t") + (v6 + v6))) + "\n\t")))
                fmt.Println(fi36)
            }
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi29
        _ = v30
    }
    v2 = v2
    for fi40 := 0; fi40 < 4; fi40++ {
        for fi41 := 0; fi41 < 1; fi41++ {
            v42 := v3[(v6 + "\n\t")]
            v43 := GenS4{F0: -415, F1: 127, F2: -409}
            fmt.Println(v6, fi41)
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi40
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        v45 := GenS5{F0: 255}
        var v46 uint = (v2 / ((v2) | 1))
        var v47 int64 = -9223372036854775808
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    v48 := GenS2{F0: 271}
    if true {
        v2 = (v2 & v2)
        v2 = 910
    } else {
        fmt.Println(v2, v6)
        var v49 string = ("a" + "hello")
        fmt.Println(v6)
        _ = v49
    }
    var v50 int64 = -765
    v6 = v6
    if (!(2147483647 <= v19)) {
        if (-1 == -1) {
            v51 := GenS3{F0: v4, F1: 65535, F2: 2147483647}
            fmt.Println(v6)
            var v52 uint8 = (v1 % ((v1) | 1))
            _ = v51
            _ = v52
        }
        v53 := GenS5{F0: 801}
        _ = v53
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v50
    _ = v6
}
