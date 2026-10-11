package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 string; F2 bool }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint }
func genFn0(p0 int64, p1 int64, p2 uint32) int {
    if ((p2 - (p2 % ((966) | 1))) > (p2 ^ p2)) {
        for fi6 := 0; fi6 < 10; fi6++ {
            p2 = p2
            fmt.Println(p2, p0, fi6)
            _ = fi6
        }
    } else {
        v7 := []int{-2147483648, -457, 2147483647}
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: (p1 / (((p1 & p1)) | 1))}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi8
            _ = v9
        }
        if len(v7) > 0 {
            v7[0] = 42
        }
        _ = v7
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if true {
            var v11 uint16 = 255
            var v12 uint64 = 18446744073709551615
            _ = v11
            _ = v12
        }
        if (127 >= 18) {
            if true {
                v13 := []int{}
                p2 = p2
                if ((^(p1)) < p0) {
                    var v14 float64 = (((0.0 + 0.1) * -2.25) / 3.14159)
                    _ = v14
                } else {
                    p0 = (p1 / ((p0) | 1))
                    var v15 int = fi10
                    _ = v15
                }
                _ = v13
            } else {
                switch -9223372036854775808 {
                case 117:
                    var v16 int32 = 42
                    _ = v16
                case -391:
                    p2 += p2
                case -2147483648:
                    p1 = (p1 + (p0 >> ((170) & 63)))
                }
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                p0 = 1
                fmt.Println(fi17, p2)
                _ = fi17
            }
            v18 := []int{fi10, (fi10 + fi10), (fi10 + (fi10 << ((173) & 63))), int(p1)}
            _ = v18
        }
        var v19 uint = 2
        _ = fi10
        _ = v19
    }
    if (0 < 65535) {
        if (false && ((!(!(100 == 1))) && true)) {
            p2 = (p2 + p2)
            for fi20 := 0; fi20 < 20; fi20++ {
                p0 += p0
                _ = fi20
            }
            switch 0 {
            case 695:
                if ((65535 > (p2 / ((p2) | 1))) && ((1 != 42) || (235 < 4))) {
                    v21 := GenS1{F0: p1}
                    var v22 uint16 = 65535
                    if true {
                        var v23 int8 = 100
                        p2 -= p2
                        v24 := []int{-9223372036854775808, 529, 2147483647}
                        _ = v23
                        _ = v24
                    }
                    _ = v21
                    _ = v22
                }
            case 128:
                var v25 int = 9223372036854775807
                _ = v25
            }
        } else {
            var v26 uint16 = 42
            v27 := GenS1{F0: p0}
            if (true && (p1 >= (p1 | int64(v26)))) {
                var v28 uint8 = 87
                var v29 uint32 = p2
                _ = v28
                _ = v29
            }
            _ = v26
            _ = v27
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}
func genFn1() int8 {
    var v30 float64 = (float64(-2.25) * -2.25)
    fmt.Println(v30)
    _ = v30
    return 0
}

func main() {
    var v1 bool = true
    var v2 int64 = -1
    var v3 string = ("foo bar" + (("\n\t" + ("foo bar" + "hello")) + "a"))
    var v4 int16 = 1
    var v5 int64 = v2
    var v31 int64 = -9223372036854775808
    v32 := GenS1{F0: (v31 / (((v5 - 530)) | 1))}
    if v1 {
        var v33 uint = 2
        var v34 uint64 = 2
        _ = v33
        _ = v34
    } else {
        if (127 <= 128) {
            var v35 float32 = 3.14159
            _ = v35
        }
        var v36 float32 = ((1.5 + (-2.25 / 1.5)) / (1.5 * 3.14159))
        _ = v36
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        switch v3 {
        case "foo bar":
            fmt.Println(v4)
        default:
            v3 = "hello"
        }
        if (!v1) {
            v38 := GenS1{F0: (v2 << ((1) & 63))}
            _ = v38
        } else {
            for fi39 := 0; fi39 < 3; fi39++ {
                switch (fi39 | (^(fi39))) {
                case 0:
                    v40 := []int{(^(fi39))}
                    _ = v40
                case 273:
                    v3 = "\n\t"
                case 4294967295:
                    var v41 uint8 = 128
                    _ = v41
                }
                _ = fi39
            }
            v31 = v2
        }
        var v42 int32 = -1
        _ = fi37
        _ = v42
    }
    v1 = (275 != 2147483648)
    v43 := []int{genFn0((v2 << ((0) & 63)), (v2 ^ v5), 65535), -9223372036854775808, 0}
    _ = len(v43)
    fmt.Println(v3, v5)
    for fi44 := 0; fi44 < 11; fi44++ {
        v45 := GenS2{F0: 2147483648, F1: (((((v3 + v3) + "") + "foo bar") + "a") + ((v3 + ("hello" + "foo bar")) + v3)), F2: (32767 >= (v4 | 32767))}
        fmt.Println(v1, v4)
        var v46 float64 = 3.14159
        _ = fi44
        _ = v45
        _ = v46
    }
    v4 += (v4 % (((v4 << ((506) & 63))) | 1))
    if len(v43) > 0 {
        v43[0] = 128
    }
    var v47 int64 = 42
    v47 += 733
    if v1 {
        v43 = append(v43, -2147483648)
        v2 -= (v31 * (v31 - v2))
        if len(v43) > 0 {
            v43[0] = -830
        }
    }
    switch 66 {
    case 384:
        var v48 uint32 = 42
        _ = v48
    case -583:
        var v49 int16 = 0
        _ = v49
    default:
        switch v3 {
        case "hello":
            for fi50 := 0; fi50 < 11; fi50++ {
                _ = len(v43)
                _ = fi50
            }
        case "":
            v51 := GenS3{F0: 0, F1: 32, F2: 74}
            _ = v51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v43
    _ = v47
    _ = v5
}
