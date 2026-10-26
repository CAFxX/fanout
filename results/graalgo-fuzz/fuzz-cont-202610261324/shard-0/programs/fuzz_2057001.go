package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint16; F2 string }
func genFn0(p0 bool) uint16 {
    for fi7 := 0; fi7 < 1; fi7++ {
        p0 = true
        var v8 uint8 = 78
        for fi9 := 0; fi9 < 7; fi9++ {
            v8 = v8
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    v10 := GenS1{F0: 444}
    _ = p0
    _ = v10
    return 42
}

func main() {
    var v1 bool = (127 == 437)
    var v2 float32 = 0.1
    var v3 int = 127
    var v4 string = ""
    var v5 uint64 = 674
    var v6 int = v3
    v11 := []int{v6, -258}
    v12 := GenS1{F0: 42}
    var v13 uint8 = 2
    switch (v6 % (((v6 + (^(v3)))) | 1)) {
    case 79:
        switch (v6 / ((-2147483648) | 1)) {
        case 4294967295:
            var v14 float64 = (-2.25 - 1e10)
            _ = v14
        case -2147483648:
            v15 := GenS1{F0: -32768}
            _ = v15
        case -9223372036854775808:
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(fi16, v1)
                var v17 bool = (v6 == fi16)
                var v18 float32 = 1.5
                _ = fi16
                _ = v17
                _ = v18
            }
        default:
            _ = len(v11)
        }
    case 42:
        v13 = v13
    case 394:
        for fi19 := 0; fi19 < 1; fi19++ {
            _ = len(v11)
            for fi20 := 0; fi20 < 17; fi20++ {
                fmt.Println(v13)
                _ = len(v11)
                _ = fi20
            }
            _ = fi19
        }
    default:
        v13 -= v13
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS2{F0: (v1 && v1), F1: (v13 / ((128) | 1)), F2: -1}
            fmt.Println(v1)
            var v24 int = (^(fi22))
            _ = fi22
            _ = v23
            _ = v24
        }
        v25 := GenS1{F0: 299}
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi21
        _ = v25
    }
    if true {
        v26 := GenS3{F0: 1, F1: v6}
        var v27 int = (v6 % ((v6) | 1))
        v28 := GenS4{F0: v1, F1: -1}
        _ = v26
        _ = v27
        _ = v28
    }
    var v29 uint32 = 629
    v13 -= 127
    v30 := GenS5{F0: 644, F1: 982, F2: ("foo bar" + "\n\t")}
    var v31 uint16 = 483
    v32 := GenS5{F0: v31, F1: genFn0(((v1 || (1.5 < (0.0 / (1e10 / (0.1 * float64(1e10)))))) || (((v4 <= (("a" + "\n\t") + v4)) && (18446744073709551615 != 65535)) && v1))), F2: v4}
    if false {
        v33 := GenS3{F0: uint32(v31), F1: v6}
        _ = v33
    }
    var v34 float32 = ((v2 - 3.14159) - float32(0.1))
    var v35 uint16 = v31
    if (!(!v1)) {
        for fi36 := 0; fi36 < 16; fi36++ {
            v3 -= (-(v3))
            v2 -= v34
            v29 = v29
            _ = fi36
        }
    }
    var v37 bool = (32768 > (v3 >> ((365) & 63)))
    var v38 uint = 255
    var v39 uint32 = v29
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
