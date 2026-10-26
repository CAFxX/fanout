package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int64; F1 int64; F2 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint32; F1 string }
func genFn0(p0 int32, p1 uint, p2 uint8) int16 {
    p0 = 2147483647
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() int64 {
    v4 := []int{4294967295, 42, 127, 128}
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: -128, F1: 42, F2: 0}
        fmt.Println(fi5)
        v4 = append(v4, (fi5 % ((127) | 1)))
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: "hello"}
    if true {
        v8 := GenS1{F0: 127, F1: 0, F2: 127}
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 uint8 = 2
                _ = fi10
                _ = v11
            }
            v12 := GenS3{F0: (32767 > genFn0(2147483647, 65535, 2)), F1: ("a" + "foo bar")}
            var v13 uint = 18446744073709551615
            _ = fi9
            _ = v12
            _ = v13
        }
        _ = v8
    }
    _ = v4
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 1
    var v2 bool = (!false)
    var v3 float32 = 0.1
    switch "a" {
    case "hello":
        var v14 float64 = 0.1
        _ = v14
    default:
        var v15 uint32 = 223
        _ = v15
    }
    var v16 float32 = v3
    if (v2 || (((((float64(v3) + (1.5 * 1.5)) - (0.0 * float64((v16 + 1.5)))) * (1e10 * 0.1)) <= ((-2.25 / (-2.25 / 1.5)) + 0.1)) || (32767 < 0))) {
        v3 = (v3 + 3.14159)
        var v17 uint16 = 0
        _ = v17
    } else {
        fmt.Println(v2)
    }
    switch -2147483648 {
    case 128:
        v18 := GenS4{F0: genFn1(), F1: -532, F2: 2147483647}
        _ = v18
    case 42:
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int8 = int8(v1)
            v2 = v2
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v22 := []int{32767}
        v16 += (0.1 - (0.1 / ((v3 * float32(1e10)) / v3)))
        _ = fi21
        _ = v22
    }
    fmt.Println(v16, v16)
    for fi23 := 0; fi23 < 15; fi23++ {
        v24 := []int{fi23, fi23}
        var v25 int = (fi23 * 9223372036854775807)
        _ = fi23
        _ = v24
        _ = v25
    }
    if ((32767 == 0) && ((!((v1 >> ((1) & 63)) <= -761)) || v2)) {
        switch 927 {
        case -1:
            fmt.Println(v1)
        default:
            var v26 uint16 = 42
            _ = v26
        }
        var v27 bool = ((v3 - v16) != (float32(((float64((v16 - (1e10 * ((v16 + 3.14159) / v16)))) - 0.1) + float64((v3 * ((v16 + 0.1) + 1e10))))) - v3))
        switch "\n\t" {
        case "foo bar":
            if (((196 >= 255) || (!v2)) || false) {
                v16 = v16
                v3 += v16
            }
        case "":
            for fi28 := 0; fi28 < 18; fi28++ {
                if (true && (!(648 < 65535))) {
                    v29 := GenS5{F0: 204}
                    v16 -= (1.5 - ((0.1 * v3) / (v16 + 0.0)))
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 1; fi30++ {
                        v16 -= v16
                        var v31 int16 = v1
                        _ = fi30
                        _ = v31
                    }
                    fmt.Println(v3)
                }
                switch "hello" {
                case "hello":
                    v16 += v3
                case "\n\t":
                    var v32 int64 = 1
                    _ = v32
                }
                _ = fi28
            }
        }
        _ = v27
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 string = ""
        _ = fi33
        _ = v34
    }
    if (!(1 == -9223372036854775808)) {
        v1 = (v1 * genFn0(1, 65535, 42))
    } else {
        v16 -= v3
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        if v2 {
            switch "a" {
            case "a":
                var v36 int = (fi35 << ((4294967295) & 63))
                _ = v36
            case "":
                var v37 int16 = genFn0(0, 255, 128)
                _ = v37
            default:
                v38 := GenS5{F0: 42}
                _ = v38
            }
        } else {
            v39 := []int{}
            for fi40 := 0; fi40 < 3; fi40++ {
                switch fi35 {
                case 801:
                    var v41 uint16 = 42
                    _ = v41
                case -2147483648:
                    switch "" {
                    case "a":
                        fmt.Println(v2)
                    case "":
                        v42 := GenS2{F0: "a"}
                        _ = v42
                    case "\n\t":
                        v43 := GenS6{F0: 265, F1: "a"}
                        _ = v43
                    default:
                        var v44 bool = (0 > 2147483647)
                        _ = v44
                    }
                default:
                    v45 := GenS6{F0: 65535, F1: ("a" + "foo bar")}
                    _ = v45
                }
                _ = len(v39)
                _ = fi40
            }
            _ = v39
        }
        _ = fi35
    }
    if (255 < 427) {
        var v46 uint64 = 389
        var v47 uint = 255
        var v48 float32 = v16
        _ = v46
        _ = v47
        _ = v48
    } else {
        var v49 int8 = -128
        _ = v49
    }
    v2 = v2
    for fi50 := 0; fi50 < 15; fi50++ {
        fmt.Println(v1)
        v1 += -32768
        _ = fi50
    }
    v1 = v1
    if v2 {
        v1 = v1
        var v51 uint16 = 32767
        _ = v51
    } else {
        var v52 float32 = v3
        _ = v52
    }
    for fi53 := 0; fi53 < 13; fi53++ {
        v54 := GenS7{F0: 2, F1: "\n\t"}
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
