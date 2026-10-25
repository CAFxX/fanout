package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 string }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 string }
func genFn0(p0 uint) string {
    var v4 int16 = 32767
    var v5 int32 = 2147483647
    var v6 int32 = (v5 / (((v5 << ((255) & 63))) | 1))
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (("hello" + "hello") + "\n\t")
}
func genFn1(p0 uint, p1 int32, p2 int32) string {
    v7 := []int{9223372036854775807, 4294967295}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v8 := GenS1{F0: 4294967295, F1: 42}
    if (true && false) {
        p0 = 18446744073709551615
        if (true || (!false)) {
            for fi9 := 0; fi9 < 18; fi9++ {
                fmt.Println(p0, p1, p2)
                p1 += p1
                p0 = (p0 / ((p0) | 1))
                _ = fi9
            }
            if ((0.0 + 3.14159) != (3.14159 * (0.0 + float64(float32((0.1 + (((-2.25 * (0.1 * -2.25)) + 0.1) - (-2.25 + float64(0.1))))))))) {
                if false {
                    var v10 int = int(p0)
                    for fi11 := 0; fi11 < 9; fi11++ {
                        v12 := GenS1{F0: 2, F1: 127}
                        v13 := GenS1{F0: 2, F1: -1}
                        var v14 float32 = float32(1e10)
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 8; fi15++ {
                        p1 = 787
                        _ = fi15
                    }
                    _ = v10
                } else {
                    var v16 bool = false
                    _ = v16
                }
                var v17 uint32 = uint32(p0)
                fmt.Println(p1)
                _ = v17
            } else {
                var v18 uint16 = 638
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 uint16 = (v18 | v18)
                    v18 -= (v18 / ((32767) | 1))
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
        }
        fmt.Println(p0, p2)
    }
    p1 += (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 int64 = 0
    v2 := []uint32{941, 2147483647, 2147483647, 2147483647}
    var v3 uint16 = 0
    var v21 int16 = -32768
    for fi22 := 0; fi22 < 3; fi22++ {
        switch ("" + "\n\t") {
        case "":
            v1 += (v1 - v1)
        case "foo bar":
            switch (fi22 << ((255) & 63)) {
            case 1:
                if len(v2) > 0 {
                    v2[0] = 65535
                }
            }
        }
        var v23 int64 = v1
        v2 = append(v2, 887)
        _ = fi22
        _ = v23
    }
    v3 -= (^(v3))
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 float64 = (-2.25 / 1.0)
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 uint64 = 42
            fmt.Println(fi24, v27, v25)
            var v28 bool = ((false && true) || true)
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi24
        _ = v25
    }
    if (255 <= 2) {
        if false {
            var v29 float32 = 1e10
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 int32 = 2147483647
                _ = fi30
                _ = v31
            }
            _ = v29
        } else {
            switch 1 {
            case 128:
                var v32 int8 = -1
                _ = v32
            case -9223372036854775808:
                if true {
                    _ = len(v2)
                    fmt.Println(v21, v3, v21)
                }
            }
            var v33 int16 = (v21 - (v21 | (v21 + (v21 % (((^(v21))) | 1)))))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v33
        }
        for fi34 := 0; fi34 < 4; fi34++ {
            if ((float32(((float64(0.1) * 0.0) / 1.0)) + float32(-2.25)) > float32(0.0)) {
                v35 := GenS1{F0: 2, F1: 114}
                if false {
                    for fi36 := 0; fi36 < 18; fi36++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v21 += -32768
                        switch "" {
                        case "a":
                            v1 = (v1 << ((2) & 63))
                        case "":
                            switch fi34 {
                            case 32768:
                                fmt.Println(fi34, v21)
                            case -419:
                                v37 := GenS1{F0: 676, F1: 42}
                                _ = v37
                            }
                        default:
                            var v38 int = (fi36 * 1)
                            _ = v38
                        }
                        _ = fi36
                    }
                    v3 -= (v3 << ((255) & 63))
                    for fi39 := 0; fi39 < 8; fi39++ {
                        fmt.Println(v1, fi39)
                        _ = fi39
                    }
                }
                _ = v35
            }
            if (!false) {
                for fi40 := 0; fi40 < 6; fi40++ {
                    var v41 int64 = (v1 + (v1 & v1))
                    if true {
                        if ((v21 - (v21 ^ v21)) >= (v21 - (v21 + v21))) {
                            v21 = v21
                            var v42 int16 = (v21 + (v21 / (((v21 + 1)) | 1)))
                            _ = v42
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 0
                            }
                        }
                        if (false || false) {
                            v2 = append(v2, 2147483647)
                            switch "hello" {
                            case "foo bar":
                                v43 := GenS2{F0: true, F1: 255, F2: ("hello" + "")}
                                _ = v43
                            case "\n\t":
                                var v44 uint32 = 42
                                _ = v44
                            default:
                                var v45 int64 = (v41 + v41)
                                _ = v45
                            }
                            switch "" {
                            case "":
                                fmt.Println(v1, fi34, fi40)
                            }
                        }
                        var v46 float32 = (3.14159 * 3.14159)
                        _ = v46
                    }
                    _ = fi40
                    _ = v41
                }
            } else {
                v1 = v1
            }
            _ = fi34
        }
    }
    if true {
        for fi47 := 0; fi47 < 20; fi47++ {
            var v48 float64 = (0.1 + 3.14159)
            for fi49 := 0; fi49 < 16; fi49++ {
                v48 = (v48 * 1e10)
                var v50 uint8 = 255
                _ = fi49
                _ = v50
            }
            v51 := GenS1{F0: 539, F1: int8(v21)}
            _ = fi47
            _ = v48
            _ = v51
        }
        var v52 int = 1
        _ = v52
    }
    switch 128 {
    case -9223372036854775808:
        _ = len(v2)
    case 424:
        fmt.Println(v3, v1)
    default:
        v1 += v1
    }
    var v53 int64 = (v1 * 42)
    switch (("foo bar" + ("" + "a")) + "hello") {
    case "hello":
        for fi54 := 0; fi54 < 17; fi54++ {
            var v55 bool = false
            _ = fi54
            _ = v55
        }
    case "foo bar":
        if true {
            var v56 uint = 255
            var v57 int = 4294967295
            for fi58 := 0; fi58 < 14; fi58++ {
                var v59 int8 = 100
                _ = fi58
                _ = v59
            }
            _ = v56
            _ = v57
        }
    case "":
        for fi60 := 0; fi60 < 18; fi60++ {
            for fi61 := 0; fi61 < 1; fi61++ {
                v53 += v1
                if (2 < 4294967295) {
                    v53 = v1
                }
                _ = fi61
            }
            var v62 string = "hello"
            _ = fi60
            _ = v62
        }
    }
    v2 = append(v2, 0)
    v63 := GenS2{F0: false, F1: 0, F2: "\n\t"}
    switch 127 {
    case -9223372036854775808:
        v21 += (v21 / (((^(v21))) | 1))
    case 127:
        v64 := GenS3{F0: 2, F1: 0}
        _ = v64
    }
    switch "\n\t" {
    case "a":
        var v65 int32 = -635
        _ = v65
    default:
        if (true && (v1 >= v53)) {
            switch ("\n\t" + "foo bar") {
            case "hello":
                var v66 int = 523
                _ = v66
            case "":
                if ((v3 % (((^(v3))) | 1)) == 522) {
                    v21 = v21
                }
            }
        } else {
            _ = len(v2)
        }
    }
    v1 -= (v53 & (v53 & 42))
    v67 := GenS4{F0: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v53
    _ = v63
    _ = v67
}
