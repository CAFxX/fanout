package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int8; F2 int64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
func genFn0(p0 int64) int32 {
    var v7 int16 = -212
    var v8 uint8 = 1
    v7 = 42
    var v9 float64 = (3.14159 + -2.25)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 1.5
    v3 := make(map[string]int)
    var v4 float64 = -2.25
    v5 := make(map[string]bool)
    var v6 uint = 0
    var v10 string = ("hello" + "a")
    v11 := GenS1{F0: 42}
    var v12 float64 = (v4 + 0.1)
    v13 := GenS2{F0: uint8(v6), F1: 1, F2: -784}
    v12 -= v12
    v12 = v4
    if (108 <= -1) {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int16 = 158
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS1{F0: -775}
            var v18 string = v10
            _ = fi16
            _ = v17
            _ = v18
        }
        if (2 <= 65535) {
            v19 := GenS3{F0: "a", F1: 1}
            var v20 int64 = -9223372036854775808
            _ = v19
            _ = v20
        }
    } else {
        v10 = v10
        v21 := GenS3{F0: ("\n\t" + v10), F1: 42}
        var v22 int32 = -1
        _ = v21
        _ = v22
    }
    if ("foo bar" != "") {
        for fi23 := 0; fi23 < 17; fi23++ {
            for fi24 := 0; fi24 < 16; fi24++ {
                switch 0 {
                case 127:
                    if ((-(fi23)) != (fi23 & (fi23 >> ((v6) & 63)))) {
                        switch fi23 {
                        case 2147483647:
                            var v25 uint = 65535
                            var v26 bool = true
                            _ = v25
                            _ = v26
                        case 1:
                            switch ("foo bar" + ((("a" + v10) + (((v10 + v10) + v10) + "hello")) + v10)) {
                            case "hello":
                                if (true && (972 >= 32767)) {
                                    var v27 int8 = 1
                                    _ = v27
                                }
                            case "a":
                                if (false || true) {
                                    v6 += 18446744073709551615
                                } else {
                                    delete(v5, v10)
                                    var v28 int32 = 0
                                    if (1 <= 1) {
                                        v29 := GenS4{F0: fi24}
                                        switch fi23 {
                                        case 127:
                                            v6 += 4294967295
                                        default:
                                            fmt.Println(v4, v12, v2)
                                        }
                                        _ = v29
                                    }
                                    fmt.Println(v4, fi23)
                                    _ = v28
                                }
                            case "":
                                v6 += (v6 & (v6 - 42))
                            default:
                                var v30 string = (((("hello" + v10) + ("hello" + ("hello" + "\n\t"))) + (v10 + ("hello" + (v10 + "hello")))) + "foo bar")
                                _ = v30
                            }
                        case 980:
                            v12 -= (v4 - v4)
                        default:
                            var v31 uint16 = 618
                            _ = v31
                        }
                        var v32 int = (fi23 << (((v6 | (v6 / (((^(v6))) | 1)))) & 63))
                        _ = v32
                    } else {
                        var v33 uint8 = 0
                        v34 := GenS5{F0: 202}
                        _ = v33
                        _ = v34
                    }
                default:
                    v10 = ("\n\t" + (v10 + "hello"))
                }
                v35 := GenS6{F0: (fi23 ^ fi23)}
                fmt.Println(v12, v10)
                _ = fi24
                _ = v35
            }
            var v36 int32 = -1
            v4 -= float64(0.0)
            _ = fi23
            _ = v36
        }
    }
    fmt.Println(v6)
    var v37 int32 = 42
    v12 -= v4
    fmt.Println(v6, v12, v6)
    var v38 int32 = -2147483648
    var v39 int = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
