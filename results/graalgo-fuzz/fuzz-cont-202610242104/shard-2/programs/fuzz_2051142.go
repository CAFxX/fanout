package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int }

func main() {
    v1 := []uint32{2, 65535, 454, 2147483648}
    var v2 string = ("\n\t" + "foo bar")
    v3 := make(map[int]string)
    var v4 float32 = (1e10 * ((0.0 - 1e10) * 0.0))
    v4 += (v4 / 1.5)
    switch 1 {
    case 128:
        switch "a" {
        case "":
            for fi5 := 0; fi5 < 11; fi5++ {
                switch fi5 {
                case 127:
                    fmt.Println(v4, v2, v2)
                case 32768:
                    var v6 uint16 = 0
                    _ = v6
                case 1:
                    for fi7 := 0; fi7 < 20; fi7++ {
                        v4 = 1.5
                        v2 = v2
                        _ = fi7
                    }
                default:
                    if ((1.5 / 1.0) > (3.14159 / 3.14159)) {
                        if true {
                            for fi8 := 0; fi8 < 10; fi8++ {
                                switch fi5 {
                                case 351:
                                    switch (v2 + v2) {
                                    case "a":
                                        switch (v2 + "\n\t") {
                                        case "foo bar":
                                            var v9 uint64 = 4294967295
                                            _ = v9
                                        }
                                    case "hello":
                                        v10 := GenS1{F0: ((("" + v2) + "hello") + v2), F1: 32767}
                                        _ = v10
                                    }
                                case -9223372036854775808:
                                    v4 -= v4
                                case 128:
                                    switch (v2 + ("hello" + "hello")) {
                                    case "a":
                                        var v11 uint = 255
                                        _ = v11
                                    case "hello":
                                        var v12 uint = 0
                                        _ = v12
                                    }
                                }
                                v13 := GenS2{F0: -27}
                                _ = fi8
                                _ = v13
                            }
                        } else {
                            for fi14 := 0; fi14 < 5; fi14++ {
                                v15 := GenS3{F0: -937, F1: fi5}
                                _ = fi14
                                _ = v15
                            }
                        }
                        if (-842 >= 2147483647) {
                            v3[1] = (v2 + v2)
                        }
                        v16 := GenS1{F0: v2, F1: 65535}
                        _ = v16
                    }
                }
                v4 += 3.14159
                _ = fi5
            }
        case "a":
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 int32 = 0
                _ = fi17
                _ = v18
            }
        }
    case 127:
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 uint64 = 350
            _ = fi19
            _ = v20
        }
    case 1:
        var v21 uint64 = 4294967295
        _ = v21
    }
    v4 = v4
    switch 42 {
    case 967:
        switch v2 {
        case "\n\t":
            var v22 int16 = 42
            _ = v22
        case "a":
            _ = len(v3)
        case "hello":
            var v23 uint8 = 2
            _ = v23
        }
    case 2147483647:
        var v24 int = 2147483647
        _ = v24
    default:
        var v25 uint = 1
        _ = v25
    }
    v4 = v4
    switch "hello" {
    case "foo bar":
        v4 = (v4 / v4)
    default:
        var v26 uint16 = 295
        _ = v26
    }
    switch "" {
    case "":
        if true {
            var v27 int64 = -47
            var v28 uint32 = 65535
            var v29 int8 = 1
            _ = v27
            _ = v28
            _ = v29
        } else {
            for fi30 := 0; fi30 < 8; fi30++ {
                var v31 int32 = -1
                v32 := GenS2{F0: -1}
                v3[32768] = v2
                _ = fi30
                _ = v31
                _ = v32
            }
            v2 = v2
        }
    case "foo bar":
        var v33 float32 = (0.1 + (1.5 + v4))
        _ = v33
    case "hello":
        if true {
            switch -231 {
            case 1:
                switch -1 {
                case -9223372036854775808:
                    var v34 string = (((("a" + v2) + (((("a" + v2) + v2) + v2) + v2)) + v2) + v2)
                    _ = v34
                default:
                    switch (v2 + v2) {
                    case "":
                        var v35 float32 = v4
                        _ = v35
                    case "foo bar":
                        _ = len(v3)
                    case "\n\t":
                        for fi36 := 0; fi36 < 11; fi36++ {
                            var v37 uint32 = 65535
                            v38 := v3[fi36]
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    default:
                        var v39 int32 = 1
                        _ = v39
                    }
                }
            default:
                if ((127 == 42) && (!(true || true))) {
                    v4 -= (0.1 + (1.5 - v4))
                } else {
                    v40 := GenS3{F0: 42, F1: 0}
                    _ = v40
                }
            }
            v3[4294967295] = ((v2 + "\n\t") + "\n\t")
            var v41 float64 = float64(float32((1.5 / -2.25)))
            _ = v41
        } else {
            for fi42 := 0; fi42 < 17; fi42++ {
                v2 = v2
                _ = fi42
            }
            var v43 int64 = 1
            _ = v43
        }
    }
    v2 = (("hello" + ("a" + "a")) + "")
    var v44 uint32 = 42
    v45 := GenS1{F0: "a", F1: 2}
    var v46 uint16 = 389
    var v47 uint8 = 255
    v47 -= v47
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
