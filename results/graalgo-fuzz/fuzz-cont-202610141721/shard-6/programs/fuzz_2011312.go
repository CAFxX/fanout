package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 string; F2 bool }
type GenS5 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint) uint64 {
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(fi6, p0)
        p0 = p0
        v7 := []int{(-(fi6)), (fi6 << ((p0) & 63))}
        _ = fi6
        _ = v7
    }
    if (int32(p0) > 968) {
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: 186, F1: 127, F2: 13}
            if true {
                v10 := GenS2{F0: (p0 >> (((p0 / (((p0 % ((p0) | 1))) | 1))) & 63))}
                fmt.Println(p0, fi8)
                _ = v10
            }
            var v11 bool = (1 > 42)
            _ = fi8
            _ = v11
            _ = v9
        }
    }
    p0 = (p0 ^ 18446744073709551615)
    _ = p0
    return 365
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 string = ("foo bar" + "")
    var v3 int32 = -282
    v4 := make(map[string]int)
    var v5 string = v2
    if (100 <= 0) {
        v12 := []int{54, 0}
        if (65535 >= 255) {
            v3 = (^(v3))
            if (int16(v3) != 1000) {
                v5 = ""
            }
        } else {
            var v13 bool = true
            _ = v13
        }
        _ = v12
    }
    if true {
        v14 := GenS3{F0: 60}
        v15 := GenS4{F0: true, F1: v2, F2: false}
        v2 = v5
        _ = v14
        _ = v15
    } else {
        v3 = v3
    }
    v1 = v1
    v5 = (v2 + v5)
    fmt.Println(v3, v2, v2)
    if true {
        if ((false && true) && true) {
            if (!((false && false) && ((9223372036854775807 == -519) || false))) {
                fmt.Println(v1, v2)
                if (2 <= 0) {
                    var v16 uint32 = 2
                    var v17 int16 = -32768
                    _ = v16
                    _ = v17
                }
                v18 := v4[v5]
                _ = v18
            }
            var v19 uint = 255
            _ = v19
        }
        if (((((-2.25 - 1.5) - ((-2.25 - float32((3.14159 / float64((1.5 - ((1e10 - (1.5 * 1.5)) + ((-2.25 * float32(float64(float32(3.14159)))) / 1.5))))))) / 0.1)) * (1.5 * 1e10)) > 1.5) || true) {
            v2 = v5
            switch 127 {
            case -885:
                v2 = ((v2 + (v5 + "")) + v2)
            case 1:
                var v20 int = -2147483648
                _ = v20
            }
        } else {
            fmt.Println(v1, v1, v3)
        }
        fmt.Println(v1)
    } else {
        switch -2147483648 {
        case -212:
            for fi21 := 0; fi21 < 11; fi21++ {
                _ = len(v4)
                _ = fi21
            }
        default:
            v22 := GenS5{F0: v1, F1: 2}
            _ = v22
        }
        if true {
            var v23 string = v5
            var v24 uint64 = (v1 / (((v1 / ((v1) | 1))) | 1))
            v25 := GenS2{F0: 65535}
            _ = v23
            _ = v24
            _ = v25
        }
    }
    switch (((v2 + v5) + "foo bar") + ("" + "hello")) {
    case "foo bar":
        var v26 bool = true
        _ = v26
    default:
        switch ((v5 + (("" + (v2 + ((v5 + v5) + v5))) + v5)) + "hello") {
        case "":
            v1 -= (v1 & v1)
        case "\n\t":
            v2 = (v5 + "")
        default:
            var v27 float32 = (1.5 / -2.25)
            _ = v27
        }
    }
    var v28 uint64 = (v1 * 42)
    v28 = (^(v28))
    if ((-573 != int(v1)) && (!(80 < 42))) {
        v5 = "a"
        v29 := v4[v5]
        switch v29 {
        case 0:
            var v30 int32 = (v3 | 617)
            _ = v30
        case -92:
            if (1 > 537) {
                v5 = v2
                switch int(v28) {
                case 906:
                    for fi31 := 0; fi31 < 14; fi31++ {
                        v28 -= 1
                        v3 += v3
                        v2 = v2
                        _ = fi31
                    }
                case 32767:
                    if (-1 >= 1) {
                        fmt.Println(v1)
                    }
                case 42:
                    fmt.Println(v1, v3, v2)
                default:
                    v1 -= 939
                }
            }
        default:
            _ = len(v4)
        }
        _ = v29
    } else {
        var v32 int = -2147483648
        var v33 int8 = 70
        var v34 uint8 = 127
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 uint8 = 255
    if (1000 != -653) {
        v4[(((((v5 + v2) + ("hello" + v2)) + ("hello" + (v5 + v2))) + v2) + ((v5 + v2) + v5))] = -9223372036854775808
        var v36 int32 = 861
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
