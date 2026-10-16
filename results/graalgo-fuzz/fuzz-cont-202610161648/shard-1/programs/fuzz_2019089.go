package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int8 }
func genFn0() string {
    var v6 uint32 = 85
    switch -2147483648 {
    case 1:
        v6 = v6
    case 980:
        if false {
            var v7 int = 32768
            if false {
                v6 -= v6
                v6 -= v6
                v8 := []int{v7, v7, (^(v7))}
                _ = v8
            }
            _ = v7
        } else {
            v9 := []int{42, 4294967295, 2147483647, 0}
            _ = v9
        }
    case 871:
        v10 := []int{0, -152, 0, 2147483647}
        _ = v10
    }
    var v11 uint8 = 83
    var v12 uint64 = 140
    _ = v11
    _ = v12
    _ = v6
    return ("a" + ("hello" + "foo bar"))
}
func genFn1(p0 int8) uint8 {
    fmt.Println(p0)
    p0 = p0
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := []int{(fi13 & (fi13 + fi13)), fi13, fi13}
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS1{F0: 0, F1: 1, F2: 1}
            fmt.Println(fi15, p0)
            switch fi13 {
            case -9223372036854775808:
                fmt.Println(p0, p0)
            }
            p0 -= 127
            if false {
                v17 := GenS1{F0: 293, F1: 2, F2: 127}
                _ = v17
            }
            fmt.Println(fi15)
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    return 0
}

func main() {
    var v1 bool = false
    v2 := []int{4294967295, -2147483648, 32768}
    v3 := []int{-9223372036854775808, -972, 4294967295}
    v4 := []string{"a", ("hello" + "foo bar")}
    var v5 uint32 = 4294967295
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            _ = len(v2)
            var v20 float64 = 1.5
            v21 := GenS2{F0: 2, F1: 1, F2: 4294967295}
            var v22 uint = 1
            if true {
                v20 = (-2.25 + ((float64(0.0) - float64(float32((v20 + ((0.0 / 3.14159) / (-2.25 * v20)))))) / v20))
            }
            if v1 {
                v23 := GenS3{F0: 847}
                var v24 uint32 = (v5 * 1)
                v25 := GenS4{F0: 42}
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = len(v2)
        _ = fi18
    }
    _ = len(v4)
    v26 := GenS4{F0: 42}
    var v27 bool = ((true || (-887 <= 1)) && v1)
    var v28 float64 = 3.14159
    if (!(70 == -128)) {
        if v27 {
            v29 := GenS3{F0: 1}
            _ = v29
        }
    }
    var v30 float32 = 0.1
    for fi31 := 0; fi31 < 8; fi31++ {
        v30 = (-2.25 - v30)
        _ = fi31
    }
    switch 15 {
    case 4294967295:
        v30 = v30
    case -2147483648:
        for fi32 := 0; fi32 < 1; fi32++ {
            var v33 int64 = -1
            if (v27 && true) {
                v33 = v33
                v34 := GenS4{F0: 4294967295}
                _ = v34
            }
            for fi35 := 0; fi35 < 8; fi35++ {
                v27 = (((v28 > v28) && (!(v27 && true))) || v27)
                fmt.Println(v33, fi35, v28)
                fmt.Println(v28, v28, v5)
                _ = fi35
            }
            _ = fi32
            _ = v33
        }
    case 32768:
        switch 32768 {
        case -414:
            v28 += 1.5
        case 9223372036854775807:
            for fi36 := 0; fi36 < 6; fi36++ {
                v37 := GenS5{F0: ("\n\t" + "a"), F1: 79}
                var v38 int32 = -2147483648
                if false {
                    for fi39 := 0; fi39 < 15; fi39++ {
                        var v40 uint16 = 2
                        var v41 int16 = 490
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    fmt.Println(v28, v28)
                    v28 += 1e10
                }
                _ = fi36
                _ = v37
                _ = v38
            }
        case -550:
            if len(v3) > 0 {
                v3[0] = 42
            }
        }
    }
    v30 += v30
    switch "a" {
    case "foo bar":
        var v42 int = -1
        _ = v42
    default:
        switch -1 {
        case -2147483648:
            for fi43 := 0; fi43 < 5; fi43++ {
                v30 -= v30
                for fi44 := 0; fi44 < 19; fi44++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi44
                }
                _ = fi43
            }
        default:
            if len(v2) > 0 {
                v2[0] = 293
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
