package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 bool; F1 uint; F2 string }
func genFn0() int8 {
    var v6 uint64 = 18446744073709551615
    fmt.Println(v6)
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint32 = 79
        var v9 uint32 = (^(v8))
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 712
    if ((1 == 255) && (-2147483648 <= 42)) {
        v10 = (v10 & v10)
        fmt.Println(v10, v6)
    } else {
        var v11 uint16 = 146
        _ = v11
    }
    _ = v10
    _ = v6
    return -1
}
func genFn1() string {
    var v12 uint16 = 255
    v13 := GenS1{F0: 1}
    v12 -= (v12 / (((v12 << ((4294967295) & 63))) | 1))
    _ = v12
    _ = v13
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{65535}
    var v3 string = "foo bar"
    var v4 int16 = 1
    var v5 float64 = float64(3.14159)
    var v14 uint = 18446744073709551615
    switch (v3 + "foo bar") {
    case "":
        v5 = (((v5 - 1e10) * 1e10) / v5)
    }
    var v15 uint64 = 282
    var v16 string = "hello"
    if false {
        var v17 int16 = (-(v4))
        _ = v17
    } else {
        if (!true) {
            var v18 string = v3
            var v19 uint32 = 398
            if false {
                var v20 int16 = (v4 % ((v4) | 1))
                _ = v20
            } else {
                switch (v18 + ("" + v3)) {
                case "hello":
                    var v21 uint8 = 1
                    _ = v21
                case "":
                    if false {
                        if true {
                            if true {
                                if false {
                                    fmt.Println(v5, v3)
                                    fmt.Println(v16)
                                } else {
                                    v22 := GenS1{F0: 42}
                                    v23 := GenS1{F0: 127}
                                    var v24 int = 177
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                v19 = v19
                                v25 := v1[(v19 << ((v14) & 63))]
                                _ = v25
                            } else {
                                fmt.Println(v16, v18)
                                v16 = v3
                            }
                            v4 = (v4 * (v4 << ((v14) & 63)))
                            var v26 int16 = 512
                            _ = v26
                        } else {
                            v5 -= float64(0.0)
                            switch 127 {
                            case -284:
                                v18 = "hello"
                            default:
                                var v27 bool = false
                                _ = v27
                            }
                        }
                        v19 = 65535
                    } else {
                        var v28 int16 = v4
                        for fi29 := 0; fi29 < 17; fi29++ {
                            fmt.Println(v3)
                            _ = fi29
                        }
                        v19 = (v19 >> ((2) & 63))
                        _ = v28
                    }
                default:
                    v19 -= v19
                }
                v4 = (v4 * (v4 << ((v14) & 63)))
                v5 -= 0.1
            }
            _ = v18
            _ = v19
        } else {
            v14 -= v14
        }
    }
    v3 = (("foo bar" + "foo bar") + "hello")
    v5 += float64(1e10)
    if (773 > 9223372036854775807) {
        var v30 string = (v3 + v3)
        switch v16 {
        case "":
            switch 1 {
            case -626:
                if (v15 < v15) {
                    v3 = "\n\t"
                    _ = len(v1)
                    v5 += v5
                }
            }
        }
        _ = v30
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        if true {
            var v32 string = "foo bar"
            var v33 string = v32
            for fi34 := 0; fi34 < 16; fi34++ {
                v35 := GenS2{F0: 42, F1: 325}
                switch v33 {
                case "foo bar":
                    v5 -= v5
                case "hello":
                    if true {
                        v36 := GenS2{F0: 900, F1: v15}
                        var v37 string = v32
                        switch 32767 {
                        case 4294967295:
                            fmt.Println(v14, v15, v15)
                        case -538:
                            var v38 uint8 = 229
                            _ = v38
                        case 0:
                            var v39 int64 = -1
                            _ = v39
                        }
                        _ = v36
                        _ = v37
                    }
                default:
                    v40 := GenS3{F0: 0, F1: 9223372036854775807}
                    _ = v40
                }
                v41 := GenS3{F0: 127, F1: 1}
                _ = fi34
                _ = v35
                _ = v41
            }
            _ = v32
            _ = v33
        } else {
            switch fi31 {
            case 2147483647:
                var v42 int32 = -1
                _ = v42
            case 32767:
                _ = len(v1)
            }
        }
        if (v4 == (v4 + v4)) {
            if true {
                var v43 int64 = 9223372036854775807
                _ = v43
            }
        }
        if true {
            v44 := GenS4{F0: true, F1: v14, F2: (v16 + v3)}
            _ = v44
        }
        _ = fi31
    }
    var v45 bool = true
    var v46 uint8 = 255
    for fi47 := 0; fi47 < 14; fi47++ {
        v15 += 42
        _ = fi47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
