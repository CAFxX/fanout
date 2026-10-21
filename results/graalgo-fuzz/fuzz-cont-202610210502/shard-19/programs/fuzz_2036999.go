package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 int64; F2 uint64 }
type GenS8 struct { F0 int16; F1 int16; F2 uint }
func genFn0(p0 bool, p1 uint64, p2 int64) bool {
    for fi4 := 0; fi4 < 13; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            if p0 {
                p0 = p0
            }
            _ = fi5
        }
        _ = fi4
    }
    v6 := GenS1{F0: 394}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (0 == -19)
}
func genFn1(p0 uint) uint32 {
    var v7 bool = false
    var v8 bool = v7
    var v9 int64 = int64(p0)
    v10 := GenS1{F0: 1}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 831
}

func main() {
    var v1 int64 = -906
    var v2 string = "hello"
    var v3 bool = (42 < 255)
    var v11 uint = 365
    switch -486 {
    case 0:
        v12 := GenS2{F0: 42, F1: -62}
        _ = v12
    case 302:
        v13 := GenS2{F0: 1, F1: 2147483647}
        _ = v13
    }
    v11 = v11
    v14 := []int{4294967295}
    if (940 > 42) {
        switch ((v2 + (("hello" + v2) + (v2 + "hello"))) + "a") {
        case "":
            if (!v3) {
                v15 := GenS1{F0: 42}
                v2 = "hello"
                var v16 uint16 = 65535
                _ = v15
                _ = v16
            }
        case "hello":
            var v17 int8 = 1
            _ = v17
        }
        switch -460 {
        case 50:
            _ = len(v14)
        case -2147483648:
            v3 = v3
        case 893:
            var v18 uint64 = 42
            _ = v18
        default:
            v3 = (550 > 42)
        }
    }
    var v19 uint32 = 2
    switch -1 {
    case -1:
        switch v2 {
        case "hello":
            fmt.Println(v3)
        }
    case 4294967295:
        var v20 float64 = 0.1
        _ = v20
    default:
        v21 := GenS3{F0: v19, F1: 65535, F2: genFn1(v11)}
        _ = v21
    }
    v22 := GenS4{F0: (v1 >> ((v11) & 63))}
    var v23 uint8 = 27
    switch 32768 {
    case 1:
        v23 -= v23
    }
    var v24 uint64 = 0
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := GenS5{F0: v23, F1: (v23 * 192), F2: genFn1((v11 | (v11 | (v11 * 387))))}
        if false {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            v27 := GenS6{F0: (^(v1))}
            switch 976 {
            case -1:
                if len(v14) > 0 {
                    v14[0] = (-(fi25))
                }
            case -224:
                _ = len(v14)
            default:
                var v28 string = v2
                _ = v28
            }
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    if (!true) {
        var v29 int64 = (v1 ^ v1)
        _ = v29
    } else {
        if (0.0 >= float32(float64(float32(1.5)))) {
            for fi30 := 0; fi30 < 17; fi30++ {
                var v31 int64 = 1
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 9; fi32++ {
                v14 = append(v14, (fi32 | (-(fi32))))
                var v33 int32 = int32(v1)
                _ = fi32
                _ = v33
            }
            if (0.0 == ((((float64((-2.25 * 0.1)) / 3.14159) + ((-2.25 + float64(0.0)) + float64(3.14159))) + -2.25) * (0.1 - 1e10))) {
                v34 := GenS7{F0: 152, F1: v1, F2: (v24 + (v24 & v24))}
                _ = v34
            } else {
                v3 = (!v3)
                var v35 int = -1
                if v3 {
                    v2 = v2
                    v36 := GenS3{F0: (v19 ^ 65535), F1: 450, F2: 193}
                    if v3 {
                        v37 := GenS6{F0: 42}
                        v38 := GenS5{F0: (v23 | 75), F1: v23, F2: 61}
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v37
                        _ = v38
                    }
                    _ = v36
                }
                _ = v35
            }
        }
        v3 = (v3 && false)
        switch "" {
        case "foo bar":
            v39 := GenS8{F0: 1000, F1: 32767, F2: v11}
            _ = v39
        default:
            if v3 {
                fmt.Println(v23, v19, v3)
            }
        }
    }
    if len(v14) > 0 {
        v14[0] = -384
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        var v41 int32 = 2147483647
        v1 = int64(fi40)
        _ = fi40
        _ = v41
    }
    var v42 string = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v11
    _ = v14
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v42
}
