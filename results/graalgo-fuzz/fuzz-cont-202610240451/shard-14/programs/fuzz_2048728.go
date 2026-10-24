package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint; F2 int }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 int32; F1 int16; F2 int8 }
type GenS7 struct { F0 int; F1 uint16; F2 uint8 }
type GenS8 struct { F0 int8; F1 uint64; F2 int }
func genFn0(p0 int8, p1 bool, p2 uint8) bool {
    p2 = p2
    if p1 {
        v5 := GenS1{F0: 0, F1: p2}
        var v6 uint16 = 32767
        if true {
            switch 9223372036854775807 {
            case 696:
                fmt.Println(p2)
            default:
                v7 := GenS1{F0: 255, F1: p2}
                _ = v7
            }
        }
        _ = v5
        _ = v6
    } else {
        fmt.Println(p2, p1)
    }
    fmt.Println(p1)
    v8 := []int{505, 0, 886, -917}
    if (p1 && true) {
        var v9 bool = p1
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := GenS1{F0: p2, F1: 115}
                var v13 int = fi11
                if (p1 || true) {
                    v13 = 165
                    v14 := GenS2{F0: p2}
                    var v15 int32 = 42
                    _ = v14
                    _ = v15
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p1
}
func genFn1() int64 {
    v16 := GenS3{F0: 42, F1: 32767}
    switch 624 {
    case -1:
        switch ("foo bar" + "foo bar") {
        case "a":
            var v17 uint16 = 2
            _ = v17
        case "\n\t":
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 uint = 42
                var v20 uint8 = 127
                var v21 uint16 = 2
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        default:
            var v22 float64 = (3.14159 / 1.0)
            _ = v22
        }
    }
    _ = v16
    return -9223372036854775808
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]int)
    var v3 int8 = -128
    var v4 uint64 = 184
    var v23 int32 = -1
    var v24 int16 = 42
    for fi25 := 0; fi25 < 3; fi25++ {
        if ((false || ((((float64(0.0) / 1.0) + 3.14159) * (0.1 / 3.14159)) > (0.0 - 1e10))) && true) {
            if true {
                for fi26 := 0; fi26 < 17; fi26++ {
                    v24 = v24
                    fmt.Println(v1, v23)
                    var v27 uint = 0
                    v4 += 17
                    _ = fi26
                    _ = v27
                }
            }
            v1 -= v1
            v28 := []int{(fi25 ^ (v1 + 32767)), (-(v1)), (v1 & fi25)}
            _ = v28
        } else {
            for fi29 := 0; fi29 < 14; fi29++ {
                v30 := []int{(v1 ^ (-(v1))), v1, -112, fi25}
                switch "hello" {
                case "hello":
                    var v31 int16 = v24
                    _ = v31
                case "":
                    _ = len(v30)
                default:
                    _ = len(v2)
                }
                _ = len(v2)
                _ = fi29
                _ = v30
            }
        }
        _ = fi25
    }
    v23 -= (v23 - -1)
    v32 := GenS1{F0: 128, F1: 2}
    for fi33 := 0; fi33 < 6; fi33++ {
        var v34 float32 = ((3.14159 - 3.14159) / 1e10)
        v23 += v23
        _ = fi33
        _ = v34
    }
    v35 := GenS4{F0: 2, F1: 255, F2: v1}
    if (!true) {
        if true {
            if true {
                var v36 float64 = 1.5
                v37 := GenS5{F0: 4294967295, F1: (v1 % (((v1 << ((4294967295) & 63))) | 1))}
                _ = v36
                _ = v37
            }
            if true {
                var v38 string = ((("\n\t" + "foo bar") + ("hello" + ("foo bar" + "hello"))) + ((("hello" + (((("\n\t" + ("foo bar" + "foo bar")) + "foo bar") + (("a" + "foo bar") + ("foo bar" + "\n\t"))) + "foo bar")) + "") + "a"))
                _ = v38
            }
            v3 -= (v3 / ((v3) | 1))
        }
        for fi39 := 0; fi39 < 4; fi39++ {
            v40 := GenS6{F0: (-(v23)), F1: -213, F2: (v3 * 51)}
            fmt.Println(v1, v1)
            fmt.Println(v24, v23)
            _ = fi39
            _ = v40
        }
        fmt.Println(v24, v4, v1)
    }
    var v41 uint32 = 4294967295
    v42 := GenS4{F0: 1, F1: 193, F2: -9223372036854775808}
    v43 := v2[v1]
    if true {
        var v44 int8 = 127
        _ = len(v2)
        _ = v44
    }
    v23 = (^(v23))
    for fi45 := 0; fi45 < 4; fi45++ {
        var v46 int16 = 32767
        fmt.Println(v46, v46)
        if true {
            v47 := []int{(v1 | 4294967295), (v43 - int(v4))}
            if true {
                var v48 uint = 944
                _ = v48
            }
            v1 -= fi45
            _ = v47
        }
        _ = fi45
        _ = v46
    }
    for fi49 := 0; fi49 < 14; fi49++ {
        var v50 int32 = 42
        fmt.Println(fi49, v4, v3)
        _ = fi49
        _ = v50
    }
    if (!true) {
        fmt.Println(v3, v1, v24)
        if false {
            v41 += v41
            if false {
                var v51 int16 = (v24 + v24)
                v52 := []int{(v1 ^ (v43 + v43))}
                var v53 int8 = v3
                _ = v51
                _ = v52
                _ = v53
            } else {
                v23 += (v23 % (((v23 * v23)) | 1))
                v54 := GenS7{F0: v43, F1: 255, F2: uint8(v4)}
                _ = v54
            }
        }
    } else {
        v23 = v23
        if false {
            v4 -= (v4 + (^(v4)))
        } else {
            var v55 uint64 = 4294967295
            fmt.Println(v43)
            v56 := GenS8{F0: (v3 * v3), F1: 2, F2: (v1 - v43)}
            _ = v55
            _ = v56
        }
        var v57 int8 = v3
        _ = v57
    }
    switch v43 {
    case 2147483647:
        v58 := []int{}
        _ = v58
    default:
        var v59 int64 = 42
        _ = v59
    }
    v60 := GenS3{F0: v23, F1: (v24 * -1)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v60
}
