package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 uint16; F1 uint; F2 int32 }
func genFn0(p0 int16, p1 int8, p2 uint32) bool {
    p0 -= p0
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint8 = 42
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        switch (fi9 * fi9) {
        case 1:
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint8 = 42
                _ = fi10
                _ = v11
            }
        default:
            v12 := []int{fi9, fi9, (fi9 - fi9), (fi9 * fi9)}
            _ = v12
        }
        v13 := []int{}
        _ = fi9
        _ = v13
    }
    v14 := GenS1{F0: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return false
}
func genFn1(p0 int32, p1 uint16) int32 {
    v15 := []int{4294967295, 127, 128, 2147483647}
    switch -944 {
    case 32768:
        v16 := GenS2{F0: -9223372036854775808, F1: 42, F2: -1}
        _ = v16
    case 79:
        var v17 int8 = 100
        _ = v17
    case 2147483647:
        v15 = append(v15, -928)
    default:
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(p0)
            _ = fi18
        }
    }
    _ = p0
    _ = p1
    _ = v15
    return 0
}

func main() {
    var v1 bool = true
    var v2 uint32 = 2
    v3 := []uint32{v2, (v2 / (((v2 ^ v2)) | 1)), 5, 4294967295}
    var v4 float64 = -2.25
    v5 := make(map[int]uint32)
    var v6 bool = v1
    v19 := GenS3{F0: v2}
    for fi20 := 0; fi20 < 7; fi20++ {
        switch (^(fi20)) {
        case -400:
            for fi21 := 0; fi21 < 17; fi21++ {
                fmt.Println(fi20, v2, v1)
                fmt.Println(v6)
                _ = fi21
            }
        case 9223372036854775807:
            if v1 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v22 uint8 = 1
                _ = v22
            } else {
                var v23 uint16 = 0
                _ = v23
            }
        }
        switch "hello" {
        case "":
            fmt.Println(v4)
        }
        if v6 {
            v6 = (42 >= (^(fi20)))
            var v24 uint32 = 42
            _ = v24
        }
        _ = fi20
    }
    fmt.Println(v4, v1)
    for fi25 := 0; fi25 < 2; fi25++ {
        v4 += (v4 - v4)
        _ = len(v3)
        _ = fi25
    }
    switch 32767 {
    case -1:
        v26 := GenS4{F0: -571}
        _ = v26
    default:
        v6 = v6
    }
    if len(v3) > 0 {
        v3[0] = v2
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v28 := GenS1{F0: -561}
        v1 = (1000 != 0)
        _ = fi27
        _ = v28
    }
    var v29 string = ("foo bar" + ("\n\t" + "foo bar"))
    v4 += float64((0.1 - (3.14159 + 0.1)))
    v5[-1] = v2
    if true {
        v4 += v4
    }
    v30 := GenS4{F0: -539}
    if v6 {
        v31 := GenS5{F0: 1, F1: 42}
        var v32 int16 = int16(v2)
        _ = v31
        _ = v32
    } else {
        _ = len(v5)
        fmt.Println(v29, v4, v2)
        var v33 uint32 = (v2 * 0)
        _ = v33
    }
    v34 := GenS5{F0: 9223372036854775807, F1: 127}
    var v35 int64 = 42
    if (0.0 > -2.25) {
        for fi36 := 0; fi36 < 4; fi36++ {
            v4 = v4
            if (v4 < float64(1e10)) {
                _ = len(v3)
            }
            _ = fi36
        }
        for fi37 := 0; fi37 < 8; fi37++ {
            v38 := GenS6{F0: 2, F1: 984, F2: -1}
            if v1 {
                for fi39 := 0; fi39 < 5; fi39++ {
                    var v40 uint = 18446744073709551615
                    v41 := GenS6{F0: 32767, F1: 421, F2: 774}
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                v2 = 850
                var v42 int64 = 902
                _ = v42
            }
            var v43 uint32 = v2
            _ = fi37
            _ = v38
            _ = v43
        }
        fmt.Println(v1, v35, v2)
    } else {
        fmt.Println(v35, v35, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
