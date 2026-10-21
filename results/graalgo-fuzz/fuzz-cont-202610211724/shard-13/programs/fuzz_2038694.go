package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int16) int32 {
    for fi4 := 0; fi4 < 20; fi4++ {
        for fi5 := 0; fi5 < 3; fi5++ {
            switch (fi4 << ((1) & 63)) {
            case 1:
                var v6 uint32 = 40
                _ = v6
            case 0:
                p0 = p0
            case -482:
                switch "" {
                case "hello":
                    v7 := GenS1{F0: 1}
                    v8 := GenS1{F0: 620}
                    _ = v7
                    _ = v8
                case "a":
                    switch (fi5 ^ (fi5 | fi5)) {
                    case 707:
                        switch fi4 {
                        case 127:
                            p0 = -1
                        default:
                            v9 := GenS2{F0: 127}
                            _ = v9
                        }
                    default:
                        fmt.Println(fi5, fi4, p0)
                    }
                }
            default:
                switch 836 {
                case -197:
                    p0 += p0
                case -9223372036854775808:
                    fmt.Println(fi4, fi4, fi5)
                case -2147483648:
                    if (2 <= 1) {
                        fmt.Println(p0)
                    } else {
                        fmt.Println(fi5, fi4, p0)
                    }
                default:
                    var v10 int8 = -61
                    _ = v10
                }
            }
            switch "foo bar" {
            case "hello":
                v11 := GenS3{F0: 4294967295, F1: (p0 % ((p0) | 1)), F2: 1}
                _ = v11
            case "\n\t":
                fmt.Println(p0)
            case "a":
                p0 = -540
            default:
                var v12 float64 = 0.1
                _ = v12
            }
            _ = fi5
        }
        var v13 float32 = 1e10
        _ = fi4
        _ = v13
    }
    p0 += (p0 * p0)
    _ = p0
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    v3 := []string{(("foo bar" + ("a" + ("foo bar" + "foo bar"))) + "a")}
    v2[458] = "\n\t"
    var v14 uint16 = 610
    var v15 bool = false
    var v16 uint32 = 2147483648
    v14 = (^(v14))
    v16 = v16
    if v15 {
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS4{F0: v14, F1: -128}
            _ = fi17
            _ = v18
        }
        var v19 int = 2147483647
        var v20 int = v19
        _ = v19
        _ = v20
    } else {
        var v21 uint32 = uint32(v14)
        v22 := GenS1{F0: 18446744073709551615}
        var v23 int = -220
        _ = v21
        _ = v22
        _ = v23
    }
    if v15 {
        _ = len(v3)
        _ = len(v1)
        fmt.Println(v15, v14, v14)
    } else {
        if (-9223372036854775808 < 1) {
            v24 := v2[0]
            for fi25 := 0; fi25 < 1; fi25++ {
                v26 := GenS5{F0: v14}
                for fi27 := 0; fi27 < 8; fi27++ {
                    _ = len(v2)
                    switch v24 {
                    case "\n\t":
                        v14 = (^(v14))
                    default:
                        v1[(fi25 + int(v14))] = (v24 + v24)
                    }
                    var v28 string = ("foo bar" + v24)
                    v29 := GenS3{F0: (v16 % (((v16 << ((2) & 63))) | 1)), F1: 618, F2: 42}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = fi25
                _ = v26
            }
            _ = v24
        } else {
            var v30 string = ("foo bar" + ((("foo bar" + "") + ("foo bar" + "")) + "\n\t"))
            v30 = (v30 + (v30 + "\n\t"))
            if len(v3) > 0 {
                v3[0] = v30
            }
            _ = v30
        }
        var v31 int16 = -384
        _ = v31
    }
    var v32 uint8 = 0
    v33 := GenS5{F0: v14}
    var v34 float64 = (((1.5 * -2.25) / 1.5) + ((((0.0 - 3.14159) / float64(float32(float64(1.5)))) + ((((3.14159 - 3.14159) - -2.25) * -2.25) / ((((1.5 - 0.0) * 0.0) + float64(3.14159)) - -2.25))) * -2.25))
    v32 += (v32 >> ((884) & 63))
    v34 += (((v34 * v34) / v34) / v34)
    if (-2147483648 != 42) {
        v15 = v15
        if v15 {
            var v35 int8 = 100
            var v36 uint64 = 2
            var v37 int32 = -265
            _ = v35
            _ = v36
            _ = v37
        }
    }
    v15 = (!true)
    var v38 string = "hello"
    switch 32768 {
    case 128:
        switch v38 {
        case "a":
            for fi39 := 0; fi39 < 7; fi39++ {
                var v40 bool = ((^(v32)) == (v32 ^ v32))
                v34 += (v34 * v34)
                v32 = 105
                _ = fi39
                _ = v40
            }
        case "\n\t":
            delete(v1, 976)
        case "foo bar":
            var v41 uint16 = (v14 - v14)
            _ = v41
        }
    default:
        v34 = v34
    }
    if (4294967295 > 70) {
        switch 0 {
        case 0:
            v3 = append(v3, v38)
        default:
            _ = len(v1)
        }
    } else {
        v34 = 0.0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v38
}
