package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint32 }
func genFn0(p0 bool) int {
    var v5 int32 = 2147483647
    var v6 uint = 316
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 uint8 = 176
    var v2 int64 = 0
    v3 := []string{("\n\t" + "foo bar"), "\n\t", "hello"}
    v4 := make(map[int]uint32)
    for fi7 := 0; fi7 < 3; fi7++ {
        if false {
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                v8 := GenS1{F0: 950, F1: 86, F2: 128}
                switch ((("" + ("foo bar" + "hello")) + "a") + "\n\t") {
                case "\n\t":
                    var v9 string = "hello"
                    _ = v9
                case "hello":
                    switch "a" {
                    case "a":
                        if false {
                            fmt.Println(v1, v2)
                        } else {
                            switch "" {
                            case "hello":
                                v10 := GenS1{F0: 2, F1: 42, F2: 0}
                                _ = v10
                            }
                            v11 := GenS1{F0: 255, F1: -128, F2: (v1 ^ (v1 << ((255) & 63)))}
                            _ = v11
                        }
                    }
                case "foo bar":
                    fmt.Println(v1, fi7, v2)
                default:
                    v2 += (v2 + 42)
                }
                v2 += 1
                _ = v8
            }
            var v12 int16 = -801
            v12 = v12
            _ = v12
        } else {
            v13 := GenS1{F0: 2, F1: 41, F2: v1}
            _ = v13
        }
        v14 := GenS2{F0: 42}
        var v15 int32 = 42
        _ = fi7
        _ = v14
        _ = v15
    }
    v16 := GenS3{F0: 0}
    var v17 uint8 = 2
    var v18 uint16 = 42
    v2 -= (v2 ^ (v2 | -168))
    for fi19 := 0; fi19 < 6; fi19++ {
        if true {
            for fi20 := 0; fi20 < 10; fi20++ {
                v1 += 100
                switch 9223372036854775807 {
                case -656:
                    if (0.1 >= float32(0.0)) {
                        var v21 int64 = v2
                        if (!((("hello" + "foo bar") > "\n\t") || (!(1.5 == ((0.0 + (1.5 + ((0.0 + float64(((0.0 * 1e10) / (3.14159 - (1.5 / float32(float64(1e10))))))) * 0.1))) / (0.0 + (-2.25 + 0.0))))))) {
                            if false {
                                _ = len(v4)
                            }
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            var v22 int32 = 711
                            _ = v22
                        } else {
                            v23 := GenS4{F0: 255, F1: 2}
                            var v24 int16 = 0
                            v25 := GenS4{F0: 65535, F1: 2147483647}
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        if true {
                            v1 = (v17 << ((1) & 63))
                            var v26 uint64 = 152
                            v17 = (v17 - (v1 | v1))
                            switch fi19 {
                            case 4294967295:
                                v27 := GenS3{F0: 42}
                                _ = len(v3)
                                if (((fi20 == fi19) || ((v1 < (v1 << ((uint(v18)) & 63))) || ((0 > (^(v18))) || ((!(!false)) && true)))) && (v18 <= v18)) {
                                    fmt.Println(v1, v21, v2)
                                    v1 = v1
                                    var v28 float32 = -2.25
                                    _ = v28
                                }
                                v2 = (v2 | 1)
                                var v29 uint32 = 4294967295
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                                _ = len(v4)
                                _ = v27
                                _ = v29
                            case 32767:
                                v2 += -102
                            }
                            _ = v26
                        }
                        _ = v21
                    }
                case 9223372036854775807:
                    v2 = 1
                }
                var v30 int64 = int64(v1)
                _ = fi20
                _ = v30
            }
            for fi31 := 0; fi31 < 18; fi31++ {
                v18 += (v18 / ((v18) | 1))
                fmt.Println(v2)
                _ = fi31
            }
        } else {
            fmt.Println(v1, v1, fi19)
            var v32 bool = (0 >= 17)
            v1 -= (v17 ^ (v17 / (((v1 >> ((126) & 63))) | 1)))
            _ = v32
        }
        switch (fi19 ^ (-(fi19))) {
        case -2147483648:
            v17 += 255
        }
        _ = fi19
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 int32 = -2147483648
        var v35 int32 = -2147483648
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := v4[0]
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
