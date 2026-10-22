package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int16, p1 uint16, p2 int32) uint32 {
    fmt.Println(p1)
    fmt.Println(p2)
    v6 := GenS1{F0: 2147483648, F1: false, F2: 0}
    p0 += 0
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 381, F1: false, F2: -9223372036854775808}
        fmt.Println(fi7, p0)
        p0 -= -32768
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}
func genFn1(p0 int32) bool {
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int = 32768
        if (607 < 0) {
            v11 := []int{(fi9 ^ v10), 128}
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    if (!false) {
        p0 += 0
        if ("hello" != "\n\t") {
            v12 := GenS1{F0: 42, F1: (0 > 2), F2: -9223372036854775808}
            v13 := GenS2{F0: 42}
            p0 = p0
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := []int{}
            v16 := GenS3{F0: 77}
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        if false {
            for fi17 := 0; fi17 < 16; fi17++ {
                v18 := []int{2147483647, (fi17 * (fi17 - (-(fi17)))), -248}
                if true {
                    fmt.Println(p0, p0)
                } else {
                    if (255 < 42) {
                        fmt.Println(fi17, p0)
                        fmt.Println(fi17)
                    }
                    var v19 int64 = -830
                    if false {
                        var v20 uint8 = 69
                        v20 -= (^(v20))
                        var v21 float32 = 3.14159
                        _ = v20
                        _ = v21
                    } else {
                        fmt.Println(v19, p0, fi17)
                        v22 := GenS2{F0: 127}
                        _ = v22
                    }
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
        }
        v23 := []int{75}
        _ = len(v23)
        _ = v23
    }
    _ = p0
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    var v5 int64 = 0
    v24 := GenS3{F0: 0}
    v25 := v1[""]
    if true {
        if true {
            v25 = ((v25 + v25) + (((v25 + "") + v25) + "hello"))
            switch 813 {
            case 127:
                for fi26 := 0; fi26 < 11; fi26++ {
                    switch 574 {
                    case -575:
                        v27 := v4[("foo bar" + (v25 + ("\n\t" + "foo bar")))]
                        _ = v27
                    case 127:
                        var v28 int16 = 1
                        _ = v28
                    }
                    var v29 float64 = float64((0.1 * ((float32(1.5) + 0.1) - 0.0)))
                    for fi30 := 0; fi30 < 14; fi30++ {
                        _ = len(v3)
                        v2 = v2
                        var v31 int32 = 0
                        _ = fi30
                        _ = v31
                    }
                    _ = fi26
                    _ = v29
                }
            }
            if (128 > 128) {
                var v32 bool = false
                _ = v32
            }
        }
    }
    v33 := []int{128, 9223372036854775807, 9223372036854775807, -786}
    v5 -= (v5 ^ v5)
    var v34 float64 = 1.5
    var v35 uint64 = (v2 * v2)
    for fi36 := 0; fi36 < 7; fi36++ {
        v35 += v35
        _ = fi36
    }
    if true {
        v34 = 1.5
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        var v38 int = -721
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        _ = fi37
        _ = v38
    }
    switch 127 {
    case -4:
        var v39 float32 = 0.1
        _ = v39
    case 9223372036854775807:
        v40 := GenS3{F0: 697}
        _ = v40
    default:
        v34 = (v34 - v34)
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        var v42 int32 = 1
        var v43 int = (^(fi41))
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
