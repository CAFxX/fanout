package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 string; F1 uint; F2 bool }
type GenS6 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int, p1 uint16) uint {
    var v7 int = p0
    if true {
        v7 -= -2147483648
    } else {
        v8 := GenS1{F0: (!(-1 <= 32767)), F1: v7}
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: (v7 != p0), F1: (p0 & v7)}
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: (((p0 << ((890) & 63)) <= (p0 / ((-9223372036854775808) | 1))) && (1.5 == ((1e10 / 3.14159) - ((1e10 * 1e10) - float64((1.5 / -2.25)))))), F1: (p0 | p0)}
        _ = v11
        _ = v8
    }
    v12 := GenS1{F0: true, F1: (p0 / (((v7 / ((p0) | 1))) | 1))}
    v13 := []int{}
    v14 := GenS1{F0: false, F1: 128}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v7
    return 286
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 uint = 2
    var v4 float32 = float32((float64(float32(0.0)) / (3.14159 + 1.5)))
    var v5 uint = v3
    var v6 uint16 = 1
    fmt.Println(v5, v4, v6)
    v3 -= 96
    v15 := GenS2{F0: 4294967295}
    var v16 int32 = 0
    var v17 uint = 255
    var v18 int16 = 824
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS2{F0: 42}
        if true {
            for fi21 := 0; fi21 < 6; fi21++ {
                v22 := GenS3{F0: v6}
                _ = fi21
                _ = v22
            }
        }
        _ = fi19
        _ = v20
    }
    if false {
        v4 -= float32(float64(((v4 - (v4 - (v4 - 0.0))) + 1e10)))
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 int8 = 60
            _ = fi23
            _ = v24
        }
        v25 := GenS1{F0: false, F1: -547}
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := []int{0, (fi26 + (fi26 & (fi26 | (fi26 / ((int(v16)) | 1))))), -1, 9223372036854775807}
        fmt.Println(fi26, v6, fi26)
        _ = fi26
        _ = v27
    }
    v28 := GenS3{F0: (v6 & v6)}
    var v29 bool = false
    switch ("foo bar" + "hello") {
    case "a":
        v30 := GenS4{F0: 0, F1: (32768 != 662)}
        _ = v30
    default:
        var v31 bool = (3.14159 != v4)
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        if v29 {
            switch 32767 {
            case -306:
                v6 += (v6 - (v6 / ((v6) | 1)))
            case 766:
                v33 := []int{(fi32 * (fi32 << ((v17) & 63))), (fi32 + fi32), (fi32 ^ (-(fi32)))}
                _ = v33
            default:
                if v29 {
                    if (v4 == (v4 - (v4 - (1e10 - v4)))) {
                        fmt.Println(fi32, v4, v3)
                        var v34 uint = (v5 * (v3 + (^(v5))))
                        v35 := GenS5{F0: "foo bar", F1: v34, F2: (127 != 42)}
                        _ = v34
                        _ = v35
                    } else {
                        var v36 float32 = 1.5
                        v37 := GenS3{F0: 65535}
                        _ = v36
                        _ = v37
                    }
                    v18 += v18
                }
            }
            fmt.Println(v6)
            fmt.Println(v3, v6)
        } else {
            v17 += v17
            if true {
                fmt.Println(v4, fi32)
                var v38 int8 = 13
                _ = v38
            } else {
                v17 += (v5 % (((v17 ^ v5)) | 1))
                for fi39 := 0; fi39 < 7; fi39++ {
                    delete(v1, 168)
                    fmt.Println(v4, v16, v6)
                    _ = fi39
                }
            }
            if (((v16 + (v16 | (v16 / ((1) | 1)))) > v16) || (!v29)) {
                fmt.Println(v6, v4)
            } else {
                var v40 uint32 = 0
                v2[850] = true
                _ = v40
            }
        }
        fmt.Println(v18, v6, v6)
        v41 := GenS6{F0: 42, F1: 112}
        _ = fi32
        _ = v41
    }
    v17 = genFn0(1, (v6 & v6))
    switch 252 {
    case 1:
        if v29 {
            switch (("" + ("a" + "hello")) + "") {
            case "hello":
                v6 = v6
            default:
                v42 := []int{}
                _ = v42
            }
            var v43 float64 = (3.14159 * 3.14159)
            _ = v43
        } else {
            var v44 float32 = ((1.5 + v4) - 3.14159)
            for fi45 := 0; fi45 < 9; fi45++ {
                v46 := []int{}
                fmt.Println(v29)
                var v47 uint64 = 8
                _ = fi45
                _ = v46
                _ = v47
            }
            _ = v44
        }
    default:
        var v48 int16 = (v18 + (v18 >> ((v3) & 63)))
        _ = v48
    }
    v29 = (v29 && ((!true) || (true || (387 >= 2))))
    if (1 < 9223372036854775807) {
        var v49 int8 = 100
        _ = v49
    } else {
        v50 := []int{9223372036854775807, -2147483648}
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
