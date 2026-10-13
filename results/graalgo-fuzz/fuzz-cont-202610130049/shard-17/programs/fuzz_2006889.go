package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint8 }
func genFn0(p0 uint32) int {
    var v4 uint64 = 2
    var v5 float32 = (1e10 + 3.14159)
    var v6 int16 = 0
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint64 = v4
        v9 := GenS1{F0: p0, F1: "a", F2: 127}
        if (65535 != 255) {
            switch (fi7 - fi7) {
            case 0:
                var v10 uint = 4294967295
                _ = v10
            case 127:
                var v11 float32 = v5
                _ = v11
            default:
                var v12 uint32 = (p0 * p0)
                _ = v12
            }
            v13 := []int{}
            if true {
                fmt.Println(v4, v4)
            } else {
                var v14 uint16 = 42
                _ = v14
            }
            _ = v13
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if (((float32(((1.5 - 0.1) + (-2.25 - 3.14159))) + float32((0.0 / -2.25))) / v5) > (v5 / ((float32(0.1) - ((-2.25 + ((float32(-2.25) * v5) - (0.1 / v5))) - v5)) / float32(0.1)))) {
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := []int{32768, fi15, fi15}
            _ = fi15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 19; fi17++ {
            if false {
                var v18 int8 = -128
                p0 = 0
                for fi19 := 0; fi19 < 5; fi19++ {
                    v20 := GenS1{F0: p0, F1: (("hello" + "a") + ("hello" + "")), F2: 42}
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
            _ = fi17
        }
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -212
}
func genFn1(p0 int32, p1 bool) uint32 {
    var v21 float32 = ((((1.5 / 3.14159) + (float32(((1e10 - 1e10) / (0.1 / 1e10))) * 0.1)) - (1.5 / (0.1 / 3.14159))) + -2.25)
    v22 := GenS1{F0: 2, F1: "\n\t", F2: 42}
    var v23 uint = 18446744073709551615
    switch ("foo bar" + ("foo bar" + "")) {
    case "hello":
        if true {
            switch 0 {
            case 878:
                v23 = v23
            case -9223372036854775808:
                fmt.Println(v23, p0, v23)
            }
            v21 = float32(1.5)
        } else {
            p1 = false
            var v24 int64 = -662
            _ = v24
        }
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v23
    return 1
}

func main() {
    var v1 float32 = (1e10 + 1e10)
    var v2 uint16 = 2
    var v3 bool = false
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint = 0
        var v27 int16 = 42
        fmt.Println(v3, fi25, v2)
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        if (float32(0.0) != 1e10) {
            v29 := GenS1{F0: 1, F1: "", F2: 138}
            if (!(!((v3 && (0.0 > (0.1 - 1.5))) || (((-2147483648 != 42) && v3) || true)))) {
                v30 := GenS1{F0: 696, F1: "\n\t", F2: 0}
                var v31 int64 = 375
                _ = v30
                _ = v31
            }
            v32 := GenS1{F0: 1, F1: ("foo bar" + ("a" + "hello")), F2: 192}
            _ = v29
            _ = v32
        } else {
            v33 := []int{}
            _ = v33
        }
        var v34 float64 = -2.25
        if v3 {
            if (157 == 127) {
                switch ("\n\t" + "\n\t") {
                case "hello":
                    switch fi28 {
                    case 1:
                        v35 := []int{fi28}
                        _ = v35
                    case 0:
                        v36 := GenS2{F0: 1, F1: (fi28 * fi28), F2: 42}
                        _ = v36
                    case 778:
                        v37 := []int{}
                        _ = v37
                    }
                default:
                    var v38 uint64 = 18446744073709551615
                    _ = v38
                }
                if true {
                    switch ("a" + "a") {
                    case "\n\t":
                        var v39 bool = ((!(234 > -9223372036854775808)) || (v3 && (-19 != 127)))
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 7; fi40++ {
                        v1 = v1
                        var v41 int8 = -109
                        _ = fi40
                        _ = v41
                    }
                    var v42 int8 = 42
                    _ = v42
                }
            }
            switch "" {
            case "a":
                fmt.Println(v3)
            }
            if (!v3) {
                var v43 int64 = 896
                v44 := []int{genFn0(2147483648), 127}
                if v3 {
                    v45 := GenS2{F0: 32767, F1: fi28, F2: 331}
                    _ = v45
                } else {
                    v46 := GenS1{F0: 4294967295, F1: "\n\t", F2: 42}
                    var v47 int32 = 1
                    fmt.Println(v3, v34, v3)
                    _ = v46
                    _ = v47
                }
                _ = v43
                _ = v44
            } else {
                if v3 {
                    fmt.Println(v2, v1, fi28)
                }
                v48 := GenS3{F0: 42, F1: (v2 - (v2 ^ v2)), F2: 128}
                _ = v48
            }
        }
        _ = fi28
        _ = v34
    }
    var v49 int8 = -128
    v50 := []int{9223372036854775807, -146}
    if (1 > 196) {
        for fi51 := 0; fi51 < 11; fi51++ {
            for fi52 := 0; fi52 < 13; fi52++ {
                _ = len(v50)
                _ = fi52
            }
            var v53 int = genFn0(2147483648)
            v53 -= fi51
            _ = fi51
            _ = v53
        }
        var v54 float32 = ((v1 / ((((0.0 * v1) + -2.25) - 0.1) / 1.0)) - 1.5)
        _ = v54
    } else {
        var v55 float32 = v1
        if v3 {
            var v56 float64 = 3.14159
            _ = v56
        }
        _ = v55
    }
    for fi57 := 0; fi57 < 18; fi57++ {
        v2 = (v2 & (^(v2)))
        var v58 int = (^(fi57))
        fmt.Println(v49, v2, v3)
        _ = fi57
        _ = v58
    }
    fmt.Println(v3, v49)
    if (359 != genFn1(-2147483648, true)) {
        v50 = append(v50, 503)
        v1 = 3.14159
    } else {
        v1 += -2.25
    }
    v59 := GenS3{F0: 0, F1: 0, F2: 0}
    switch 127 {
    case -807:
        var v60 float64 = 0.1
        _ = v60
    case 4294967295:
        var v61 float64 = 1.5
        _ = v61
    }
    var v62 int8 = (v49 >> ((255) & 63))
    if len(v50) > 0 {
        v50[0] = 127
    }
    v62 = (-(v62))
    var v63 int64 = 9223372036854775807
    v49 -= (v49 >> ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v59
    _ = v62
    _ = v63
}
