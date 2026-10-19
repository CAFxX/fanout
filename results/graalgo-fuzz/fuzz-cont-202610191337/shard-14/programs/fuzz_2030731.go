package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int64; F1 string; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint) uint16 {
    p1 += p1
    for fi4 := 0; fi4 < 20; fi4++ {
        if true {
            p0 = 917
            for fi5 := 0; fi5 < 7; fi5++ {
                var v6 uint8 = 0
                fmt.Println(v6)
                if ((v6 | v6) != (v6 << ((uint(p0)) & 63))) {
                    if false {
                        var v7 uint = 0
                        _ = v7
                    } else {
                        v6 = v6
                        var v8 int64 = int64(p0)
                        v9 := []int{}
                        v10 := GenS1{F0: 42, F1: p1}
                        var v11 float32 = 3.14159
                        switch "foo bar" {
                        case "foo bar":
                            var v12 uint = (p1 - 648)
                            _ = v12
                        case "":
                            v13 := GenS1{F0: -2147483648, F1: (p1 >> ((p1) & 63))}
                            _ = v13
                        default:
                            v11 -= v11
                        }
                        v9 = append(v9, -9223372036854775808)
                        var v14 bool = true
                        _ = v10
                        _ = v11
                        _ = v14
                        _ = v8
                        _ = v9
                    }
                    p1 += (^(p1))
                } else {
                    v15 := []int{fi5, 9223372036854775807, -935}
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = len(v15)
                    _ = v15
                }
                fmt.Println(fi4, v6, p0)
                var v16 bool = false
                _ = fi5
                _ = v16
                _ = v6
            }
            p1 -= (p1 ^ (p1 + 4294967295))
        } else {
            v17 := []int{fi4, fi4, 127}
            _ = v17
        }
        v18 := GenS1{F0: -339, F1: 1}
        switch "foo bar" {
        case "hello":
            var v19 int8 = -70
            _ = v19
        case "":
            v20 := GenS2{F0: -1, F1: "a", F2: 2147483648}
            _ = v20
        case "a":
            switch "" {
            case "\n\t":
                v21 := GenS3{F0: -1, F1: (p1 / (((p1 << ((p1) & 63))) | 1)), F2: 1}
                _ = v21
            default:
                switch (fi4 << ((p1) & 63)) {
                case -188:
                    fmt.Println(p1, p1, p1)
                case 128:
                    var v22 float32 = float32(((1.5 - 0.0) + 3.14159))
                    _ = v22
                case 9223372036854775807:
                    fmt.Println(p0)
                }
            }
        default:
            var v23 int16 = -32768
            _ = v23
        }
        _ = fi4
        _ = v18
    }
    p0 = p0
    _ = p0
    _ = p1
    return 899
}
func genFn1(p0 int) uint {
    fmt.Println(p0)
    v24 := GenS4{F0: 1}
    _ = p0
    _ = v24
    return 42
}

func main() {
    var v1 int8 = 24
    var v2 int16 = -32768
    v3 := []uint32{2147483647, 2147483648, 0}
    v25 := GenS2{F0: 416, F1: "", F2: 1}
    var v26 float64 = 1.5
    var v27 float64 = 0.0
    v27 = (((3.14159 + float64(1e10)) / 3.14159) * (float64(3.14159) + (float64((((1e10 * ((((-2.25 - ((-2.25 + -2.25) - 3.14159)) * (0.0 * (0.0 - (float32(((v26 - 1e10) * 1e10)) - -2.25)))) / float32(float64(3.14159))) - 0.1)) - 3.14159) - float32(v26))) * (v26 - ((0.1 / float64(3.14159)) / (v26 - ((v27 + v26) + v27)))))))
    var v28 int16 = -1
    var v29 string = ""
    var v30 int = 4294967295
    var v31 uint32 = 665
    v29 = "a"
    var v32 string = (v29 + v29)
    for fi33 := 0; fi33 < 15; fi33++ {
        if (false || (!true)) {
            _ = len(v3)
            fmt.Println(v27, v30)
            v30 = v30
        } else {
            v3 = append(v3, v31)
            for fi34 := 0; fi34 < 2; fi34++ {
                fmt.Println(v27)
                if len(v3) > 0 {
                    v3[0] = (v31 * v31)
                }
                var v35 uint64 = 724
                v32 = v29
                _ = fi34
                _ = v35
            }
            v31 = v31
        }
        switch v29 {
        case "":
            var v36 uint8 = 255
            _ = v36
        case "foo bar":
            if false {
                v37 := GenS1{F0: 101, F1: 65535}
                if ((v29 + v32) == v29) {
                    v38 := GenS5{F0: -32768}
                    v39 := GenS1{F0: -1, F1: genFn1((v30 >> ((378) & 63)))}
                    v28 += 255
                    _ = v38
                    _ = v39
                } else {
                    var v40 int8 = v1
                    for fi41 := 0; fi41 < 10; fi41++ {
                        var v42 uint = 123
                        v32 = v32
                        v43 := GenS4{F0: 18446744073709551615}
                        v42 += v42
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    v40 += v1
                    _ = v40
                }
                v44 := GenS5{F0: v28}
                _ = v37
                _ = v44
            }
        case "hello":
            var v45 float64 = (v27 + (float64(1e10) + 1e10))
            _ = v45
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v31 = (v31 * 214)
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
