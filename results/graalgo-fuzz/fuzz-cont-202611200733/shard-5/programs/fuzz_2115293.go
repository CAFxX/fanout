package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint; F1 int; F2 uint16 }
func genFn0(p0 uint, p1 int16) uint8 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int = (-(fi4))
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint16 = 382
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: (p1 - (p1 >> (((p0 << ((p0) & 63))) & 63))), F1: 1, F2: (("foo bar" + "foo bar") + ("" + "foo bar"))}
        _ = fi6
        _ = v9
    }
    v10 := GenS1{F0: (p1 ^ (p1 >> ((237) & 63))), F1: 1, F2: "a"}
    p1 += -1
    _ = p0
    _ = p1
    _ = v10
    return 0
}
func genFn1(p0 uint16, p1 uint) uint16 {
    v11 := GenS2{F0: "a"}
    v12 := GenS2{F0: "\n\t"}
    v13 := GenS3{F0: -2147483648, F1: 1}
    p1 = p1
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return p0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 0
    v3 := make(map[uint32]string)
    v14 := []int{0}
    if ((!false) || ("\n\t" <= ("\n\t" + "a"))) {
        v2 = v2
    }
    v2 = (v2 & (v2 & 398))
    v2 -= 1
    var v15 uint16 = 1
    for fi16 := 0; fi16 < 15; fi16++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        if true {
            _ = len(v1)
            var v17 string = ("" + "hello")
            var v18 uint8 = 0
            _ = v17
            _ = v18
        } else {
            var v19 uint8 = 86
            switch -9223372036854775808 {
            case 0:
                var v20 int64 = 9223372036854775807
                _ = v20
            default:
                if true {
                    _ = len(v14)
                }
            }
            var v21 uint64 = (^(v2))
            _ = v19
            _ = v21
        }
        _ = fi16
    }
    fmt.Println(v2, v2)
    v3[65535] = ""
    var v22 uint16 = 953
    if true {
        switch "foo bar" {
        case "foo bar":
            v22 += (v22 / ((v15) | 1))
        }
    } else {
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS4{F0: 21, F1: 0}
            _ = fi23
            _ = v24
        }
        if ((32767 != 127) && (!(true || true))) {
            v15 = (v15 | genFn1(v15, 2))
            v25 := GenS1{F0: 42, F1: 255, F2: ("a" + "hello")}
            _ = len(v3)
            _ = v25
        } else {
            if (617 <= 891) {
                if (!(!false)) {
                    v26 := v1[1]
                    var v27 string = "a"
                    _ = v26
                    _ = v27
                }
                if (3.14159 < float32((1.5 - 1e10))) {
                    var v28 int16 = 1
                    fmt.Println(v2, v28)
                    _ = v28
                } else {
                    var v29 float32 = ((((0.1 * ((-2.25 + (1e10 - 1.5)) * float32((float64((1.5 / 1e10)) * ((float64(3.14159) + 0.1) + 3.14159))))) * float32(-2.25)) * -2.25) * 0.1)
                    fmt.Println(v15, v29, v2)
                    if false {
                        fmt.Println(v2)
                        var v30 int16 = 32767
                        _ = v30
                    }
                    _ = v29
                }
            }
            switch -1 {
            case 0:
                v31 := GenS5{F0: uint(v15), F1: 0, F2: v22}
                _ = v31
            case 128:
                v2 = v2
            case 4294967295:
                v2 = 2
            }
            v32 := v1[42]
            _ = v32
        }
    }
    var v33 float32 = 0.1
    if len(v14) > 0 {
        v14[0] = 127
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        v35 := v1[612]
        var v36 string = ("hello" + "")
        _ = fi34
        _ = v35
        _ = v36
    }
    switch 32767 {
    case 128:
        var v37 int = 127
        _ = v37
    }
    v22 = (v15 + (v15 - 255))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v33
}
