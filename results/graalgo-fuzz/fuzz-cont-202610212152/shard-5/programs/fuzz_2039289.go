package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 string = ("a" + (("" + ("\n\t" + (("foo bar" + "") + "foo bar"))) + (("a" + "a") + "hello")))
    var v2 int = 42
    var v3 bool = false
    var v4 uint32 = 196
    v5 := []int{-1, (v2 | (v2 << ((222) & 63))), (^(v2))}
    if (false && false) {
        var v6 float64 = (float64(((0.1 - ((3.14159 * -2.25) / 3.14159)) + 3.14159)) * 1e10)
        _ = v6
    } else {
        v7 := GenS1{F0: v1, F1: ("a" + "hello")}
        _ = v7
    }
    switch v2 {
    case 9223372036854775807:
        var v8 int64 = 1
        _ = v8
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v5 = append(v5, fi9)
        if (!v3) {
            v10 := GenS1{F0: v1, F1: "foo bar"}
            if v3 {
                v11 := GenS1{F0: v1, F1: ("hello" + (v1 + "foo bar"))}
                var v12 uint = 65535
                v13 := GenS1{F0: "", F1: "foo bar"}
                _ = v11
                _ = v12
                _ = v13
            } else {
                v14 := GenS2{F0: 42, F1: 0, F2: 18446744073709551615}
                _ = v14
            }
            switch (fi9 << ((193) & 63)) {
            case -772:
                if (float64(-2.25) != 0.1) {
                    var v15 uint64 = 2
                    var v16 uint8 = 255
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint64 = 0
                    fmt.Println(v1, v3, fi9)
                    var v18 uint32 = v4
                    _ = v17
                    _ = v18
                }
            default:
                var v19 bool = true
                _ = v19
            }
            _ = v10
        }
        if (-797 <= 1) {
            fmt.Println(fi9, v3)
        } else {
            var v20 bool = v3
            if len(v5) > 0 {
                v5[0] = v2
            }
            var v21 float64 = (-2.25 / (float64((-2.25 - 0.0)) + 0.1))
            _ = v20
            _ = v21
        }
        _ = fi9
    }
    v1 = (v1 + "\n\t")
    var v22 bool = v3
    var v23 uint = 18446744073709551615
    for fi24 := 0; fi24 < 4; fi24++ {
        if v3 {
            var v25 int16 = 943
            for fi26 := 0; fi26 < 6; fi26++ {
                if true {
                    var v27 int16 = v25
                    _ = v27
                } else {
                    switch (((v1 + v1) + "hello") + "\n\t") {
                    case "a":
                        var v28 uint8 = 221
                        _ = v28
                    default:
                        var v29 uint = v23
                        _ = v29
                    }
                    v22 = v22
                }
                fmt.Println(v2, v1, fi24)
                _ = fi26
            }
            var v30 int = -483
            _ = v25
            _ = v30
        }
        if false {
            var v31 float64 = (3.14159 - -2.25)
            if (0 != 194) {
                v23 = (v23 - v23)
                fmt.Println(v31)
                for fi32 := 0; fi32 < 11; fi32++ {
                    var v33 uint16 = uint16(v4)
                    v5 = append(v5, (fi24 ^ fi24))
                    _ = fi32
                    _ = v33
                }
            }
            v34 := GenS1{F0: "", F1: "a"}
            _ = v31
            _ = v34
        } else {
            v35 := GenS2{F0: 255, F1: 25, F2: 1}
            if v3 {
                v36 := GenS2{F0: 17, F1: 28, F2: 2}
                _ = v36
            } else {
                for fi37 := 0; fi37 < 2; fi37++ {
                    var v38 uint8 = 101
                    _ = fi37
                    _ = v38
                }
            }
            v5 = append(v5, (v2 & 0))
            _ = v35
        }
        _ = fi24
    }
    switch (v2 - v2) {
    case 1:
        v4 = v4
    }
    v4 -= 2
    if (!(-1 < 2147483647)) {
        var v39 uint8 = 93
        if (1 < 1) {
            for fi40 := 0; fi40 < 7; fi40++ {
                v2 += v2
                for fi41 := 0; fi41 < 12; fi41++ {
                    switch v1 {
                    case "":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "hello":
                        v42 := GenS3{F0: 778, F1: 2}
                        _ = v42
                    case "\n\t":
                        v3 = v3
                    }
                    v3 = ((fi41 - (fi41 - (fi40 % (((fi41 & 0)) | 1)))) != v2)
                    _ = fi41
                }
                _ = fi40
            }
            v5 = append(v5, 150)
            switch (v2 | (-(v2))) {
            case 971:
                fmt.Println(v1)
            }
        } else {
            for fi43 := 0; fi43 < 20; fi43++ {
                v1 = (((v1 + v1) + ("foo bar" + "a")) + (("\n\t" + ("a" + "a")) + "a"))
                for fi44 := 0; fi44 < 3; fi44++ {
                    var v45 uint8 = (^(v39))
                    v4 += v4
                    v46 := GenS3{F0: 2, F1: 320}
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                if len(v5) > 0 {
                    v5[0] = (v2 % ((70) | 1))
                }
                _ = fi43
            }
            var v47 float32 = 0.0
            var v48 float32 = v47
            _ = v47
            _ = v48
        }
        _ = v39
    }
    for fi49 := 0; fi49 < 14; fi49++ {
        var v50 uint16 = uint16(v4)
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
