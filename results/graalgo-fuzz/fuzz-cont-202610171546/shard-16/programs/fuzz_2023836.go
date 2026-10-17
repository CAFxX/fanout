package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int64; F1 bool; F2 int }
type GenS4 struct { F0 uint; F1 bool; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int; F1 uint16; F2 uint64 }
func genFn0(p0 uint64, p1 uint32) int16 {
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            p0 += p0
            _ = fi5
        }
        _ = fi4
    }
    if (42 > 0) {
        fmt.Println(p1, p1)
        var v6 int64 = 969
        for fi7 := 0; fi7 < 2; fi7++ {
            if ((p1 >> ((1) & 63)) <= p1) {
                var v8 uint16 = 65535
                _ = v8
            }
            _ = fi7
        }
        _ = v6
    }
    var v9 float32 = 0.1
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint = 65535
        v12 := []int{fi10, fi10}
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v9
    return -240
}
func genFn1(p0 int64) uint64 {
    if ((false || (int32(p0) == 42)) || false) {
        var v13 float32 = ((3.14159 - (0.0 * 0.0)) - 3.14159)
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS1{F0: 12, F1: (p0 % ((p0) | 1))}
        v16 := GenS1{F0: -75, F1: p0}
        var v17 bool = true
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: -1, F1: (p0 + 456)}
    var v19 uint32 = 0
    var v20 float64 = 1e10
    _ = p0
    _ = v18
    _ = v19
    _ = v20
    return 629
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -291
    var v3 int8 = 100
    v21 := GenS2{F0: 255, F1: 0}
    switch -926 {
    case 1:
        if ((false || true) && (1e10 >= float32(3.14159))) {
            var v22 int16 = -32768
            var v23 int64 = -342
            v1[65535] = (("hello" + "hello") + "a")
            _ = v22
            _ = v23
        } else {
            v24 := GenS2{F0: 99, F1: 2147483647}
            switch "foo bar" {
            case "hello":
                v25 := GenS3{F0: 256, F1: false, F2: 42}
                _ = v25
            case "\n\t":
                v26 := GenS4{F0: 4294967295, F1: (!true), F2: -1}
                _ = v26
            }
            v27 := GenS5{F0: 584, F1: 42}
            _ = v24
            _ = v27
        }
    default:
        var v28 int = -390
        _ = v28
    }
    v29 := GenS6{F0: -966, F1: 594, F2: 1}
    var v30 uint8 = 255
    if false {
        fmt.Println(v2)
        _ = len(v1)
    }
    v31 := []int{9223372036854775807, 9223372036854775807, -62, -43}
    var v32 int8 = v3
    var v33 uint16 = 34
    switch 127 {
    case -611:
        v34 := GenS1{F0: (v32 & (v32 ^ v3)), F1: -1}
        _ = v34
    default:
        var v35 float32 = 1.5
        _ = v35
    }
    switch "" {
    case "\n\t":
        for fi36 := 0; fi36 < 6; fi36++ {
            var v37 float32 = 1e10
            var v38 uint32 = 1
            v37 += -2.25
            _ = fi36
            _ = v37
            _ = v38
        }
    case "foo bar":
        for fi39 := 0; fi39 < 4; fi39++ {
            _ = len(v1)
            _ = fi39
        }
    case "hello":
        v40 := GenS2{F0: v33, F1: 4294967295}
        _ = v40
    default:
        var v41 uint32 = 2
        _ = v41
    }
    var v42 bool = false
    v2 = (-(v2))
    v30 -= v30
    v43 := v1[2147483647]
    for fi44 := 0; fi44 < 1; fi44++ {
        if v42 {
            var v45 int8 = v3
            for fi46 := 0; fi46 < 8; fi46++ {
                var v47 uint8 = 255
                v42 = (v42 || false)
                if (76 >= 65535) {
                    v32 = -7
                } else {
                    v30 = 255
                    var v48 uint = 2
                    fmt.Println(v33, v42, v42)
                    _ = v48
                }
                _ = fi46
                _ = v47
            }
            if ((v33 >> ((30) & 63)) >= (^(v33))) {
                switch v43 {
                case "\n\t":
                    v42 = (v30 <= (v30 | v30))
                default:
                    var v49 uint = 255
                    _ = v49
                }
                v32 -= v45
                if len(v31) >= 2 {
                    v31 = v31[1:len(v31)]
                }
            }
            _ = v45
        }
        v2 = (v2 & v2)
        _ = fi44
    }
    var v50 bool = v42
    var v51 uint64 = 226
    v30 -= (v30 - (v30 + v30))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v42
    _ = v43
    _ = v50
    _ = v51
}
