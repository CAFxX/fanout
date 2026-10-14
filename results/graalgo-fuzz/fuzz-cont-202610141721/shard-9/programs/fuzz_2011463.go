package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int32 }
type GenS2 struct { F0 bool; F1 uint; F2 int32 }
type GenS3 struct { F0 uint16 }
func genFn0() int8 {
    var v4 string = ""
    var v5 uint = 255
    v6 := []int{4294967295, -219}
    var v7 int32 = 2147483647
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -128
}
func genFn1() string {
    if false {
        var v8 int = -863
        var v9 int64 = -1
        fmt.Println(v8, v8)
        _ = v8
        _ = v9
    } else {
        if (-1 != 9223372036854775807) {
            if false {
                var v10 bool = (1 < 42)
                _ = v10
            }
            v11 := GenS1{F0: 42, F1: 255, F2: -2147483648}
            var v12 uint64 = 782
            _ = v11
            _ = v12
        } else {
            v13 := GenS1{F0: -526, F1: 229, F2: -2147483648}
            switch "\n\t" {
            case "hello":
                v14 := GenS1{F0: -473, F1: 128, F2: 729}
                _ = v14
            case "a":
                var v15 int = 810
                _ = v15
            }
            _ = v13
        }
        v16 := []int{-9223372036854775808, 9223372036854775807, -9223372036854775808, -2147483648}
        for fi17 := 0; fi17 < 8; fi17++ {
            if (127 < 253) {
                if (true || false) {
                    if (!true) {
                        if false {
                            fmt.Println(fi17)
                            var v18 float64 = float64(((0.0 * 0.1) - -2.25))
                            v19 := GenS1{F0: 1, F1: 243, F2: -55}
                            _ = v18
                            _ = v19
                        }
                        var v20 int64 = 1
                        v20 += 93
                        _ = v20
                    }
                    if true {
                        if (1 <= 963) {
                            _ = len(v16)
                            v21 := GenS1{F0: 32767, F1: 42, F2: -2147483648}
                            for fi22 := 0; fi22 < 7; fi22++ {
                                var v23 int64 = -554
                                var v24 int = fi22
                                v25 := GenS2{F0: false, F1: 42, F2: 42}
                                _ = fi22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            _ = v21
                        }
                        var v26 float32 = 0.0
                        _ = v26
                    }
                }
            }
            _ = len(v16)
            _ = fi17
        }
        _ = v16
    }
    if true {
        var v27 int64 = -1
        v27 += v27
        if ((((("\n\t" + ("a" + "\n\t")) + "hello") + "") + "hello") != "\n\t") {
            v28 := GenS1{F0: 1, F1: 240, F2: 42}
            var v29 float64 = (3.14159 * 0.0)
            v29 -= 1.5
            _ = v28
            _ = v29
        }
        _ = v27
    }
    return "foo bar"
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]string)
    v3 := []uint32{}
    switch 32768 {
    case 337:
        if false {
            v30 := GenS2{F0: (!false), F1: 255, F2: 991}
            _ = v30
        }
    case -303:
        if true {
            var v31 float32 = (3.14159 / v1)
            v31 = v31
            _ = v31
        } else {
            switch 571 {
            case -481:
                v32 := GenS2{F0: true, F1: 18446744073709551615, F2: 42}
                _ = v32
            case -1:
                _ = len(v3)
            case -9223372036854775808:
                delete(v2, ("hello" + ("foo bar" + "hello")))
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                v1 = v1
                var v34 int16 = 1000
                fmt.Println(v34)
                _ = fi33
                _ = v34
            }
            v35 := v2[(((("\n\t" + "") + (((("hello" + "foo bar") + "\n\t") + "foo bar") + "a")) + "hello") + "\n\t")]
            _ = v35
        }
    }
    v36 := GenS1{F0: 0, F1: 42, F2: -197}
    v3 = append(v3, 335)
    v1 += v1
    var v37 uint32 = 271
    var v38 uint16 = 565
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 uint16 = v38
        switch fi39 {
        case 32768:
            if (2147483647 <= 42) {
                v41 := GenS3{F0: (v38 % (((v38 ^ v38)) | 1))}
                _ = v41
            }
        }
        fmt.Println(v38, v37)
        _ = fi39
        _ = v40
    }
    v38 -= v38
    fmt.Println(v37, v38, v38)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
