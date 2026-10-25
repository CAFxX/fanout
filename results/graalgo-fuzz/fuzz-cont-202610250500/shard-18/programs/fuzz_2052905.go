package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int16; F1 int32; F2 int32 }
func genFn0(p0 int8) int64 {
    switch 42 {
    case -1:
        fmt.Println(p0)
    case 538:
        if false {
            fmt.Println(p0)
        } else {
            var v5 uint8 = 223
            _ = v5
        }
    default:
        for fi6 := 0; fi6 < 15; fi6++ {
            p0 = 0
            p0 = -46
            _ = fi6
        }
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 32767}
        p0 = -1
        _ = fi7
        _ = v8
    }
    v9 := []int{32767, 42}
    var v10 uint16 = 255
    _ = p0
    _ = v10
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 uint32, p1 int16) uint {
    fmt.Println(p0, p0)
    var v11 float32 = 3.14159
    p0 += p0
    _ = p0
    _ = p1
    _ = v11
    return 916
}

func main() {
    var v1 float64 = 3.14159
    var v2 int = 2147483647
    v3 := make(map[uint32]uint32)
    var v4 string = ""
    delete(v3, 239)
    if true {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 bool = (false || (42 < 1))
            var v14 uint16 = 42
            switch ((((("foo bar" + (v4 + v4)) + (v4 + ((v4 + ((v4 + ((v4 + (v4 + (v4 + ""))) + v4)) + v4)) + "hello"))) + "") + "") + (v4 + (("hello" + v4) + ("\n\t" + (v4 + ((v4 + (v4 + "\n\t")) + ("hello" + v4))))))) {
            case "":
                v15 := GenS1{F0: 255}
                _ = v15
            case "\n\t":
                for fi16 := 0; fi16 < 8; fi16++ {
                    v1 += v1
                    if v13 {
                        var v17 string = v4
                        switch ("foo bar" + ("hello" + "\n\t")) {
                        case "":
                            var v18 float64 = v1
                            _ = v18
                        case "\n\t":
                            v19 := GenS1{F0: (v14 * 32767)}
                            _ = v19
                        default:
                            if false {
                                v20 := GenS2{F0: 390, F1: 2, F2: 979}
                                _ = v20
                            }
                        }
                        _ = v17
                    } else {
                        fmt.Println(v14, v14, v1)
                        var v21 uint8 = 128
                        _ = v21
                    }
                    _ = fi16
                }
            case "a":
                if (-32768 < 42) {
                    v22 := []int{(fi12 - (fi12 << ((1) & 63))), v2, 9223372036854775807, v2}
                    if (1 >= 10) {
                        var v23 uint = 1
                        switch "a" {
                        case "":
                            delete(v3, 2147483647)
                        case "hello":
                            delete(v3, 844)
                        }
                        _ = v23
                    }
                    v24 := GenS3{F0: 135, F1: 42}
                    _ = v22
                    _ = v24
                }
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        var v25 float32 = (3.14159 * (0.1 - float32(float64(-2.25))))
        _ = v25
    }
    v2 = v2
    var v26 int = v2
    v27 := GenS1{F0: 0}
    v3[443] = 2147483647
    var v28 uint8 = 128
    switch int(v28) {
    case 9223372036854775807:
        var v29 bool = true
        _ = v29
    case 128:
        v30 := v3[2147483647]
        _ = v30
    default:
        var v31 uint = 18446744073709551615
        _ = v31
    }
    if (genFn0(127) != genFn0(-1)) {
        v32 := GenS4{F0: -21, F1: -2147483648, F2: 0}
        _ = v32
    }
    if true {
        var v33 int = -350
        var v34 int64 = 806
        var v35 uint16 = 42
        _ = v33
        _ = v34
        _ = v35
    } else {
        if true {
            var v36 int16 = -32768
            _ = v36
        } else {
            v37 := GenS1{F0: 1}
            _ = v37
        }
    }
    if (0.0 <= ((0.1 - -2.25) * (float32(float64((0.1 / 1.0))) + (3.14159 * 1e10)))) {
        if (false && (2147483647 != 42)) {
            v38 := GenS3{F0: 185, F1: 2147483647}
            _ = v38
        }
        switch (v2 * v26) {
        case -2147483648:
            var v39 int32 = -179
            _ = v39
        case -9223372036854775808:
            var v40 uint64 = 914
            _ = v40
        }
    }
    var v41 int64 = 672
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
}
