package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint8; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint32, p1 int8) string {
    v5 := []int{42}
    for fi6 := 0; fi6 < 18; fi6++ {
        switch fi6 {
        case -1:
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 string = ("foo bar" + "\n\t")
                var v9 uint16 = 0
                v10 := GenS1{F0: -1}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        case 32767:
            var v11 int16 = 1
            _ = v11
        case -218:
            v5 = append(v5, (^(fi6)))
        }
        if true {
            _ = len(v5)
            if (((32767 > 2) && false) || (938 > 65535)) {
                p0 = 0
                v12 := GenS1{F0: -220}
                v13 := GenS1{F0: -1}
                _ = v12
                _ = v13
            }
            var v14 uint64 = 802
            _ = v14
        } else {
            fmt.Println(p0, p0)
            var v15 int8 = p1
            _ = v15
        }
        p1 -= (p1 - p1)
        _ = fi6
    }
    var v16 bool = false
    _ = p0
    _ = p1
    _ = v16
    _ = v5
    return (("a" + "hello") + "\n\t")
}

func main() {
    var v1 bool = (!true)
    var v2 uint8 = 42
    var v3 uint64 = 2
    v4 := []uint32{65535}
    switch "a" {
    case "":
        var v17 uint64 = (v3 * 18446744073709551615)
        _ = v17
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            if (784 > 149) {
                for fi18 := 0; fi18 < 2; fi18++ {
                    switch "\n\t" {
                    case "hello":
                        v19 := GenS1{F0: 218}
                        _ = v19
                    case "\n\t":
                        v4 = append(v4, 364)
                    default:
                        if len(v4) > 0 {
                            v4[0] = 2
                        }
                    }
                    _ = fi18
                }
                v3 = 18446744073709551615
            } else {
                v2 -= 128
                if (-1 != -266) {
                    var v20 string = "hello"
                    v21 := GenS1{F0: 2147483647}
                    var v22 uint16 = 32767
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS2{F0: -9223372036854775808, F1: (v2 - 1), F2: true}
                    v3 -= v3
                    v2 = (v2 % (((v2 & v2)) | 1))
                    _ = v23
                }
            }
        case "":
            v3 = v3
        default:
            var v24 bool = ((1 == 2) || true)
            _ = v24
        }
    case "a":
        var v25 int32 = 913
        _ = v25
    }
    v2 = 42
    var v26 int16 = 0
    switch "" {
    case "":
        v1 = (("foo bar" + "foo bar") == "a")
    case "hello":
        if ((v2 >> ((255) & 63)) > (^(v2))) {
            v27 := GenS2{F0: 42, F1: (v2 % ((v2) | 1)), F2: (!v1)}
            _ = v27
        }
    case "\n\t":
        var v28 float32 = (1.5 / (3.14159 * (3.14159 - -2.25)))
        _ = v28
    default:
        var v29 uint32 = 1
        _ = v29
    }
    var v30 int64 = -1
    v4 = append(v4, 65535)
    if (2147483647 == 34) {
        var v31 float32 = 1e10
        if (!true) {
            for fi32 := 0; fi32 < 4; fi32++ {
                v26 -= (-(v26))
                v31 -= ((1e10 + (v31 + 3.14159)) / (float32(0.1) * v31))
                _ = fi32
            }
        }
        _ = v31
    } else {
        for fi33 := 0; fi33 < 11; fi33++ {
            if len(v4) > 0 {
                v4[0] = 696
            }
            for fi34 := 0; fi34 < 19; fi34++ {
                v35 := GenS3{F0: v30}
                v36 := GenS4{F0: 163, F1: v26, F2: 588}
                if false {
                    v30 += v30
                    switch (fi33 / ((-696) | 1)) {
                    case 1:
                        _ = len(v4)
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v26 = 783
                }
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = fi33
        }
        switch -943 {
        case -551:
            fmt.Println(v3)
        case 42:
            v37 := GenS5{F0: 127}
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        v39 := GenS1{F0: 2147483647}
        for fi40 := 0; fi40 < 17; fi40++ {
            fmt.Println(v30)
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
}
