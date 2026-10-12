package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int32; F2 uint16 }
func genFn0() uint16 {
    v4 := []int{}
    var v5 float32 = (0.0 - (1e10 - 1.5))
    switch "hello" {
    case "":
        v6 := GenS1{F0: 32767, F1: 595, F2: -1}
        _ = v6
    }
    _ = v4
    _ = v5
    return 514
}
func genFn1(p0 uint8, p1 int32, p2 uint32) uint32 {
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS2{F0: 1}
        p1 -= p1
        if ((0.0 + 1.5) == 3.14159) {
            p2 -= p2
            v9 := GenS1{F0: -85, F1: -697, F2: -9223372036854775808}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    if ((p2 + (p2 / ((65535) | 1))) != (p2 * p2)) {
        switch -1 {
        case 0:
            var v10 float64 = 1.5
            _ = v10
        case 4294967295:
            var v11 int = -9223372036854775808
            _ = v11
        }
        v12 := []int{}
        switch 9223372036854775807 {
        case 32768:
            fmt.Println(p1, p1)
        case 0:
            switch 1000 {
            case 4294967295:
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 uint8 = p0
                    _ = fi13
                    _ = v14
                }
            }
        case 42:
            if true {
                for fi15 := 0; fi15 < 2; fi15++ {
                    p1 += (p1 - p1)
                    fmt.Println(p1)
                    switch fi15 {
                    case 0:
                        p0 = (^(p0))
                    case -9223372036854775808:
                        v16 := GenS3{F0: 0}
                        _ = v16
                    case -572:
                        fmt.Println(fi15, p0, p2)
                    }
                    _ = fi15
                }
                p2 = p2
            } else {
                for fi17 := 0; fi17 < 16; fi17++ {
                    p2 += 420
                    _ = fi17
                }
                var v18 string = ("\n\t" + "hello")
                _ = v18
            }
        default:
            switch ("\n\t" + ("foo bar" + "")) {
            case "":
                v19 := GenS4{F0: false, F1: p1, F2: genFn0()}
                _ = v19
            case "a":
                var v20 uint64 = 338
                _ = v20
            case "hello":
                p0 -= p0
            default:
                for fi21 := 0; fi21 < 8; fi21++ {
                    v12 = append(v12, (fi21 >> ((2) & 63)))
                    _ = fi21
                }
            }
        }
        _ = v12
    }
    var v22 float32 = (0.0 * (float32(1.5) + 1.5))
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return p2
}

func main() {
    v1 := []int{-697, 128, -777, -9223372036854775808}
    var v2 string = ""
    var v3 int = 2147483647
    v23 := GenS3{F0: 108}
    var v24 uint16 = 209
    v25 := GenS4{F0: true, F1: 1, F2: (v24 / (((v24 / ((1) | 1))) | 1))}
    fmt.Println(v24, v2)
    v1 = append(v1, -461)
    if (!(float32((-2.25 / (1e10 / (0.1 - 0.0)))) >= -2.25)) {
        if (1 >= 65535) {
            v24 = (v24 & v24)
        } else {
            v24 = v24
            var v26 bool = (false || true)
            fmt.Println(v24, v26, v3)
            _ = v26
        }
        if (2147483647 < 0) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v24 -= 32767
        }
        switch v3 {
        case 128:
            var v27 bool = false
            _ = v27
        case -2147483648:
            v24 = v24
        }
    } else {
        var v28 int16 = -872
        if false {
            for fi29 := 0; fi29 < 11; fi29++ {
                var v30 float32 = 3.14159
                if (!(!true)) {
                    v2 = v2
                    var v31 int = (-(fi29))
                    _ = v31
                } else {
                    v2 = ("a" + v2)
                }
                for fi32 := 0; fi32 < 10; fi32++ {
                    switch v2 {
                    case "hello":
                        var v33 uint16 = genFn0()
                        _ = v33
                    }
                    _ = fi32
                }
                _ = fi29
                _ = v30
            }
        }
        _ = v28
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ("hello" + (v2 + ("hello" + (v2 + (v2 + "a"))))) {
    case "foo bar":
        v2 = v2
    }
    if true {
        var v34 int32 = -880
        v24 = 42
        _ = v34
    }
    fmt.Println(v2)
    switch v2 {
    case "a":
        v1 = append(v1, -21)
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            v1 = append(v1, (v3 - (v3 >> ((4294967295) & 63))))
        case "a":
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        default:
            var v35 string = (v2 + v2)
            _ = v35
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
