package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint; F1 int16; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS7 struct { F0 bool }
func genFn0(p0 int) bool {
    v4 := []int{9223372036854775807, p0}
    var v5 int32 = 807
    var v6 bool = (289 < 942)
    v5 = v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 uint8, p1 uint16, p2 uint) int8 {
    fmt.Println(p1, p0, p1)
    switch 1 {
    case -1:
        if ((false && (true && true)) && (1 == 42)) {
            p2 -= (p2 << ((4294967295) & 63))
        }
    case -865:
        v7 := GenS1{F0: 9223372036854775807, F1: 1000}
        _ = v7
    case 947:
        var v8 uint32 = 42
        _ = v8
    default:
        v9 := []int{}
        _ = v9
    }
    if (-1 != -9223372036854775808) {
        for fi10 := 0; fi10 < 6; fi10++ {
            switch (fi10 - fi10) {
            case 32768:
                v11 := []int{(fi10 + fi10)}
                _ = v11
            case -81:
                for fi12 := 0; fi12 < 11; fi12++ {
                    v13 := GenS1{F0: -656, F1: -32768}
                    if (int16(fi10) <= 561) {
                        switch fi10 {
                        case 0:
                            fmt.Println(fi10, fi10)
                        default:
                            var v14 bool = true
                            _ = v14
                        }
                        v15 := GenS2{F0: p2, F1: 1, F2: 42}
                        var v16 int16 = 0
                        v16 += 42
                        _ = v15
                        _ = v16
                    }
                    var v17 uint16 = (p1 & (p1 << ((p2) & 63)))
                    _ = fi12
                    _ = v13
                    _ = v17
                }
            case -479:
                v18 := GenS2{F0: 861, F1: 1000, F2: 127}
                _ = v18
            default:
                p1 = p1
            }
            var v19 bool = false
            _ = fi10
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 100
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int8 = -69
    v3 := make(map[int]uint32)
    fmt.Println(v1, v2)
    delete(v3, -760)
    if (!(!true)) {
        v1 -= 2
        v20 := GenS3{F0: 188, F1: 127, F2: -327}
        switch -9223372036854775808 {
        case -188:
            if false {
                switch 32767 {
                case -1:
                    v1 += (v1 | v1)
                case 2147483647:
                    v1 = 2
                case 4294967295:
                    v2 = v2
                }
                v21 := GenS1{F0: 9223372036854775807, F1: 1}
                _ = v21
            } else {
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 int = 1
                    var v24 uint = 18446744073709551615
                    v1 += (v1 << (((^(v24))) & 63))
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v25 := GenS4{F0: 2, F1: (v2 - (v2 & int8(v1)))}
                _ = v25
            }
        case -2147483648:
            v26 := GenS4{F0: 4294967295, F1: (v2 % ((v2) | 1))}
            _ = v26
        }
        _ = v20
    } else {
        v27 := GenS5{F0: 2, F1: 890}
        v28 := GenS2{F0: 0, F1: 1, F2: v2}
        _ = v27
        _ = v28
    }
    if (!(255 <= 556)) {
        v29 := GenS6{F0: 127, F1: 846, F2: 2}
        var v30 int32 = 2147483647
        _ = v29
        _ = v30
    } else {
        var v31 uint32 = 65535
        switch "" {
        case "\n\t":
            _ = len(v3)
        }
        _ = v31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 int32 = -346
        fmt.Println(v2)
        var v34 uint64 = v1
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        for fi36 := 0; fi36 < 2; fi36++ {
            v37 := []int{int(v1), fi36, fi36, fi35}
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    var v38 uint = 1
    v38 += v38
    for fi39 := 0; fi39 < 5; fi39++ {
        for fi40 := 0; fi40 < 20; fi40++ {
            v41 := GenS7{F0: true}
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
}
