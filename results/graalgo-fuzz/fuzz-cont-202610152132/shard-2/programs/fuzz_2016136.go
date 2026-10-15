package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 int; F2 int32 }
func genFn0(p0 uint32) int8 {
    var v5 uint16 = 933
    var v6 uint32 = (p0 + (^(p0)))
    for fi7 := 0; fi7 < 4; fi7++ {
        if false {
            v8 := GenS1{F0: 1}
            v6 = 489
            _ = v8
        } else {
            v9 := []int{fi7}
            _ = v9
        }
        v10 := GenS2{F0: 65535}
        v5 = v5
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int64 = 42
        v5 -= v5
        _ = fi11
        _ = v12
    }
    v6 = 42
    _ = p0
    _ = v5
    _ = v6
    return -104
}

func main() {
    var v1 float32 = ((1.5 - 1.5) / -2.25)
    v2 := []uint32{2147483648, 188, 4294967295, 42}
    v3 := []string{"foo bar", "\n\t", "\n\t"}
    var v4 uint8 = 127
    switch "" {
    case "\n\t":
        fmt.Println(v4)
    default:
        if ((3.14159 + (1e10 / float64((-2.25 * ((((1.5 + v1) + 1e10) * 1.5) * v1))))) < (float64(float32(3.14159)) + ((3.14159 * (3.14159 - 3.14159)) / (0.1 * 1.5)))) {
            var v13 float32 = v1
            _ = v13
        }
    }
    var v14 int64 = 110
    v15 := GenS1{F0: v4}
    if (42 > 127) {
        if ((!true) && true) {
            v1 += v1
            var v16 int = 42
            var v17 uint = 255
            _ = v16
            _ = v17
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v18 uint16 = 42
    if (!true) {
        v1 = 1.5
        if false {
            v19 := GenS2{F0: 255}
            switch 42 {
            case -635:
                if ((v4 & v4) < 117) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v4 = 128
                    fmt.Println(v14, v14)
                }
            case 127:
                var v20 float32 = v1
                _ = v20
            default:
                if false {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        switch (fi21 % (((fi21 & 32768)) | 1)) {
                        case -147:
                            if (255 <= v4) {
                                for fi22 := 0; fi22 < 15; fi22++ {
                                    var v23 float64 = 1e10
                                    v23 += 1.5
                                    _ = fi22
                                    _ = v23
                                }
                                for fi24 := 0; fi24 < 16; fi24++ {
                                    switch "foo bar" {
                                    case "a":
                                        if true {
                                            if true {
                                                var v25 int = fi24
                                                fmt.Println(v18)
                                                _ = v25
                                            }
                                        } else {
                                            v26 := GenS2{F0: 827}
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v27 := GenS3{F0: 65535, F1: 1}
                                            _ = v26
                                            _ = v27
                                        }
                                    default:
                                        v18 += (v18 ^ (v18 << ((18446744073709551615) & 63)))
                                    }
                                    fmt.Println(v18, fi21)
                                    if true {
                                        v18 = v18
                                    }
                                    _ = fi24
                                }
                            } else {
                                v4 = (^(v4))
                                switch fi21 {
                                case -535:
                                    switch "foo bar" {
                                    case "foo bar":
                                        v28 := GenS2{F0: 793}
                                        _ = v28
                                    }
                                }
                                var v29 string = ""
                                _ = v29
                            }
                        case 2147483647:
                            if false {
                                v1 += 0.1
                                _ = len(v3)
                            } else {
                                v1 = -2.25
                            }
                        case 32768:
                            var v30 string = "a"
                            _ = v30
                        default:
                            if (42 >= 2147483647) {
                                v1 -= (0.0 + (0.1 - ((v1 - ((3.14159 - (v1 * ((float32(0.1) / (((0.0 - (v1 * v1)) / (v1 * (v1 + v1))) + v1)) + (v1 + (float32(3.14159) + v1))))) - v1)) * (v1 / v1))))
                                var v31 int64 = (v14 ^ (-(v14)))
                                _ = v31
                            }
                        }
                        v14 = 42
                        _ = fi21
                    }
                    if (!true) {
                        v32 := GenS4{F0: true}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v32
                    }
                    var v33 string = "foo bar"
                    _ = v33
                }
            }
            _ = v19
        }
        fmt.Println(v18)
    }
    v4 = v4
    v1 = v1
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 uint8 = (v4 / (((v4 - 255)) | 1))
        if false {
            v36 := GenS5{F0: "foo bar", F1: 9223372036854775807, F2: -2147483648}
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 float32 = 3.14159
                v14 += (^(v14))
                _ = fi37
                _ = v38
            }
            if (0 > -1) {
                v4 = (^(v4))
            } else {
                var v39 uint64 = 42
                switch fi34 {
                case -727:
                    for fi40 := 0; fi40 < 8; fi40++ {
                        v1 += v1
                        var v41 uint32 = uint32(v39)
                        _ = fi40
                        _ = v41
                    }
                case -2147483648:
                    _ = len(v2)
                case 1:
                    v4 = (v35 << ((451) & 63))
                }
                for fi42 := 0; fi42 < 16; fi42++ {
                    v14 = v14
                    var v43 string = ((("a" + ("" + "a")) + "hello") + "hello")
                    fmt.Println(v43, v4, v43)
                    _ = fi42
                    _ = v43
                }
                _ = v39
            }
            _ = v36
        } else {
            switch (fi34 | (^(fi34))) {
            case -844:
                v44 := GenS2{F0: 425}
                _ = v44
            }
            fmt.Println(v4)
        }
        _ = fi34
        _ = v35
    }
    for fi45 := 0; fi45 < 4; fi45++ {
        v14 -= -863
        for fi46 := 0; fi46 < 5; fi46++ {
            var v47 int = fi46
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
