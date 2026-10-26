package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint32 }
func genFn0(p0 int16, p1 bool) int8 {
    p0 = (p0 >> ((255) & 63))
    for fi7 := 0; fi7 < 7; fi7++ {
        switch ("\n\t" + "") {
        case "foo bar":
            v8 := GenS1{F0: fi7}
            _ = v8
        }
        p0 += (p0 & -1)
        var v9 uint64 = 925
        _ = fi7
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := []int{int(p0)}
            var v13 float64 = 0.1
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    return -97
}
func genFn1(p0 uint16, p1 int64) int16 {
    var v14 int = 734
    p1 = p1
    _ = p0
    _ = p1
    _ = v14
    return -385
}

func main() {
    var v1 int16 = 42
    var v2 int = -2147483648
    var v3 int = (v2 | (v2 * v2))
    v4 := make(map[int]string)
    var v5 uint8 = 42
    var v6 string = (("a" + "hello") + "foo bar")
    var v15 int32 = 42
    if (!(2147483647 <= -56)) {
        for fi16 := 0; fi16 < 4; fi16++ {
            v5 -= (v5 >> ((964) & 63))
            for fi17 := 0; fi17 < 9; fi17++ {
                if true {
                    var v18 uint8 = v5
                    switch v2 {
                    case 128:
                        _ = len(v4)
                    }
                    _ = v18
                }
                var v19 uint32 = 514
                v20 := GenS1{F0: (v3 - (v3 << ((42) & 63)))}
                _ = fi17
                _ = v19
                _ = v20
            }
            switch v3 {
            case 128:
                for fi21 := 0; fi21 < 7; fi21++ {
                    v22 := GenS1{F0: (^(fi21))}
                    var v23 int16 = int16(v5)
                    var v24 uint16 = 354
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            case 841:
                var v25 uint32 = 2147483648
                _ = v25
            default:
                v15 = v15
            }
            _ = fi16
        }
    }
    v26 := []int{32767}
    for fi27 := 0; fi27 < 4; fi27++ {
        v28 := GenS1{F0: (-(v2))}
        if false {
            for fi29 := 0; fi29 < 12; fi29++ {
                v30 := GenS1{F0: (fi29 ^ (fi27 & (-(v2))))}
                _ = fi29
                _ = v30
            }
            v3 += 32768
            v5 += v5
        } else {
            var v31 uint16 = 497
            _ = v31
        }
        _ = fi27
        _ = v28
    }
    var v32 uint8 = 128
    for fi33 := 0; fi33 < 9; fi33++ {
        v6 = v6
        v1 = 1000
        var v34 string = "hello"
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        fmt.Println(v15, v2, v32)
        var v36 uint64 = 1
        _ = fi35
        _ = v36
    }
    if (!true) {
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    } else {
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    }
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    v2 += (v3 & (v2 & v2))
    v37 := GenS2{F0: 106, F1: 18446744073709551615}
    var v38 float32 = 3.14159
    if (352 < 201) {
        if true {
            if false {
                fmt.Println(v5, v32)
            } else {
                if false {
                    var v39 bool = (-128 > 42)
                    v40 := GenS3{F0: 4294967295, F1: int16(v2)}
                    var v41 int64 = -338
                    _ = v39
                    _ = v40
                    _ = v41
                }
                v42 := GenS2{F0: -1, F1: 0}
                v26 = append(v26, (v2 & (v3 / ((v2) | 1))))
                _ = v42
            }
            for fi43 := 0; fi43 < 8; fi43++ {
                for fi44 := 0; fi44 < 18; fi44++ {
                    if false {
                        v1 = 32767
                    }
                    v45 := GenS4{F0: (^(v5)), F1: 0, F2: 2147483647}
                    _ = fi44
                    _ = v45
                }
                _ = fi43
            }
            v32 = v32
        } else {
            for fi46 := 0; fi46 < 5; fi46++ {
                _ = len(v4)
                var v47 string = "foo bar"
                fmt.Println(v5)
                _ = fi46
                _ = v47
            }
            v15 = 1
        }
    } else {
        for fi48 := 0; fi48 < 13; fi48++ {
            switch (fi48 ^ 0) {
            case -9223372036854775808:
                var v49 uint16 = 42
                _ = v49
            case 32767:
                var v50 int8 = 0
                _ = v50
            }
            _ = fi48
        }
        if (2147483647 != 0) {
            v32 = (^(v5))
            for fi51 := 0; fi51 < 16; fi51++ {
                v52 := GenS4{F0: (v32 / ((127) | 1)), F1: int64(v15), F2: 2147483648}
                for fi53 := 0; fi53 < 7; fi53++ {
                    v15 = 42
                    _ = fi53
                }
                v1 = (v1 ^ (v1 % ((42) | 1)))
                _ = fi51
                _ = v52
            }
            fmt.Println(v5)
        }
    }
    v15 -= -1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
