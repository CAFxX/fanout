package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int16 }
type GenS2 struct { F0 int16; F1 int8; F2 int }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint16; F1 string }

func main() {
    var v1 uint = 255
    v2 := make(map[string]int)
    var v3 int32 = -2147483648
    for fi4 := 0; fi4 < 5; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            switch (("foo bar" + (("a" + ((("" + "hello") + ("a" + "\n\t")) + "\n\t")) + ("hello" + ("hello" + (("foo bar" + (("hello" + "hello") + "a")) + "a"))))) + "a") {
            case "foo bar":
                fmt.Println(v1, fi5, fi5)
            }
            if true {
                var v6 uint8 = 1
                var v7 int32 = v3
                _ = v6
                _ = v7
            }
            _ = fi5
        }
        fmt.Println(fi4, fi4)
        if (true || false) {
            if (202 <= 0) {
                var v8 int16 = 1000
                v1 -= 18446744073709551615
                delete(v2, "a")
                _ = v8
            } else {
                for fi9 := 0; fi9 < 19; fi9++ {
                    if ((1e10 <= (-2.25 * (((float64(-2.25) - (float64(1.5) / 3.14159)) + 1e10) * 0.0))) || false) {
                        switch fi9 {
                        case 128:
                            fmt.Println(fi4, fi4)
                        default:
                            _ = len(v2)
                            v1 -= 87
                            if true {
                                v2[("" + "")] = fi9
                                v1 += 868
                                v3 = v3
                                var v10 float32 = 0.0
                                _ = v10
                            }
                            if (0 < 70) {
                                v11 := []int{}
                                if ((0.1 - (((float32((((1.5 - (1.5 * (0.1 - 3.14159))) + (((-2.25 + 1e10) * 1.5) * 0.1)) - 1.5)) + 1e10) + 1e10) - 0.1)) == (1.5 - 0.1)) {
                                    var v12 uint = 0
                                    _ = v12
                                } else {
                                    v3 += (v3 ^ (-(v3)))
                                    var v13 string = "foo bar"
                                    _ = v13
                                }
                                v3 += (v3 / (((v3 - -646)) | 1))
                                v3 = int32(v1)
                                _ = v11
                            }
                        }
                        v14 := GenS1{F0: -1, F1: 32767, F2: -831}
                        fmt.Println(fi4, fi9, fi4)
                        v3 += v3
                        v1 = 65535
                        _ = v14
                    } else {
                        switch "a" {
                        case "\n\t":
                            v15 := []int{(fi9 + fi4), -841, 42}
                            var v16 uint64 = 1
                            _ = v15
                            _ = v16
                        }
                    }
                    switch 2147483647 {
                    case 38:
                        if false {
                            v1 = (^(v1))
                        } else {
                            v17 := GenS2{F0: 1, F1: 109, F2: fi9}
                            _ = v17
                        }
                    case -1:
                        if ((236 >= 42) || true) {
                            v18 := []int{0, -607, 817, fi4}
                            var v19 uint = (v1 | v1)
                            if true {
                                v20 := GenS2{F0: -1, F1: 100, F2: (fi4 - fi4)}
                                v19 += v19
                                _ = v20
                            }
                            fmt.Println(fi4, fi4, v1)
                            v21 := GenS2{F0: 0, F1: 42, F2: fi4}
                            _ = v18
                            _ = v19
                            _ = v21
                        }
                    }
                    _ = fi9
                }
                var v22 bool = (65535 == 65535)
                if true {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        switch (fi4 + fi4) {
                        case -642:
                            var v24 bool = false
                            _ = v24
                        default:
                            fmt.Println(v1)
                        }
                        _ = fi23
                    }
                    for fi25 := 0; fi25 < 19; fi25++ {
                        v26 := []int{(fi4 - (fi4 << ((v1) & 63)))}
                        var v27 float32 = 0.0
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS1{F0: -1, F1: -558, F2: 105}
                    _ = v28
                } else {
                    v29 := GenS1{F0: 41, F1: 1000, F2: -790}
                    var v30 uint16 = 1
                    switch "foo bar" {
                    case "\n\t":
                        v31 := GenS2{F0: 32767, F1: 127, F2: 0}
                        _ = v31
                    case "hello":
                        _ = len(v2)
                    default:
                        var v32 uint8 = uint8(v1)
                        _ = v32
                    }
                    _ = v29
                    _ = v30
                }
                _ = v22
            }
        }
        _ = fi4
    }
    v33 := GenS3{F0: 296, F1: 9}
    var v34 int32 = (-(v3))
    var v35 uint32 = 8
    v34 -= (v34 >> ((v1) & 63))
    for fi36 := 0; fi36 < 14; fi36++ {
        v37 := GenS4{F0: 18446744073709551615, F1: 1}
        v38 := GenS5{F0: 893, F1: "\n\t"}
        _ = fi36
        _ = v37
        _ = v38
    }
    v3 += (v3 % ((-2147483648) | 1))
    v3 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
