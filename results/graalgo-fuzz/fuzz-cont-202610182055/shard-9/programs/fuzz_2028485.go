package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 int8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint64) int16 {
    v6 := []int{221}
    if ((42 > 0) || (!(0 > 42))) {
        v7 := GenS1{F0: true, F1: 32767, F2: 580}
        var v8 uint64 = (p1 & (p1 * (p1 | (p1 >> ((2) & 63)))))
        _ = v7
        _ = v8
    }
    var v9 int8 = -128
    switch ("a" + ("" + "a")) {
    case "hello":
        var v10 int = 128
        _ = v10
    case "\n\t":
        switch "" {
        case "a":
            v11 := GenS1{F0: (true && (!true)), F1: -1, F2: -1}
            _ = v11
        case "\n\t":
            p1 = p1
        case "foo bar":
            var v12 uint32 = 2147483647
            _ = v12
        }
    case "":
        for fi13 := 0; fi13 < 8; fi13++ {
            if (255 != (p0 % (((p0 + p0)) | 1))) {
                var v14 int16 = -32768
                for fi15 := 0; fi15 < 11; fi15++ {
                    v9 = v9
                    var v16 int = 1
                    _ = fi15
                    _ = v16
                }
                var v17 uint = 2
                _ = v14
                _ = v17
            } else {
                if (true && (736 <= 883)) {
                    if true {
                        var v18 uint16 = (^(p0))
                        _ = v18
                    }
                } else {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        var v20 uint = 549
                        v9 = v9
                        v21 := GenS1{F0: false, F1: 1, F2: 0}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v9 = (v9 - v9)
                    if false {
                        p1 = (^(p1))
                        fmt.Println(p1, v9)
                    } else {
                        var v22 int16 = 0
                        _ = v22
                    }
                }
                if true {
                    v23 := GenS2{F0: p0, F1: -762, F2: v9}
                    var v24 uint = 351
                    fmt.Println(v9, fi13)
                    _ = v23
                    _ = v24
                } else {
                    var v25 int16 = -322
                    if true {
                        v26 := GenS1{F0: true, F1: 1000, F2: 222}
                        _ = v26
                    } else {
                        if true {
                            var v27 int = (-(fi13))
                            p0 += p0
                            _ = v27
                        }
                        var v28 float32 = 1.5
                        _ = v28
                    }
                    _ = v25
                }
                if (!true) {
                    fmt.Println(fi13, p1)
                }
            }
            p0 += 42
            _ = fi13
        }
    }
    v29 := GenS1{F0: (false || (true && (!true))), F1: -32768, F2: 42}
    _ = p0
    _ = p1
    _ = v29
    _ = v6
    _ = v9
    return -321
}
func genFn1(p0 int32, p1 int) uint16 {
    for fi30 := 0; fi30 < 6; fi30++ {
        switch (((("foo bar" + "a") + "hello") + ("foo bar" + "a")) + "hello") {
        case "a":
            var v31 uint16 = 42
            _ = v31
        case "hello":
            var v32 uint64 = 0
            _ = v32
        }
        _ = fi30
    }
    var v33 int64 = 684
    _ = p0
    _ = p1
    _ = v33
    return 440
}

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 42
    var v4 uint = 4294967295
    var v5 int32 = v1
    var v34 int8 = 100
    v2[2147483648] = 664
    delete(v2, 0)
    var v35 int = 0
    v36 := GenS1{F0: false, F1: genFn0(2, 18446744073709551615), F2: (v1 ^ 578)}
    delete(v2, 213)
    fmt.Println(v4, v3)
    if false {
        v4 = (v4 % (((v4 % ((2) | 1))) | 1))
        fmt.Println(v35)
    } else {
        switch "" {
        case "hello":
            v35 = (v35 & v35)
        default:
            var v37 uint64 = 977
            _ = v37
        }
        if (818 < (-(v1))) {
            delete(v2, 4294967295)
        }
    }
    v38 := GenS2{F0: 255, F1: genFn0(2, 1), F2: v34}
    var v39 int64 = 42
    var v40 int64 = -9223372036854775808
    v1 = (-(v1))
    var v41 int8 = 1
    var v42 float64 = 3.14159
    v43 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
