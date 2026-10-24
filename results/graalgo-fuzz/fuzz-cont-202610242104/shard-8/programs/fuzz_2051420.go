package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 string }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int32, p1 uint, p2 int32) uint32 {
    p0 += p2
    v4 := []int{990, 2147483647, -9223372036854775808, 127}
    _ = len(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2147483647
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    var v3 uint32 = 298
    v5 := GenS1{F0: 2, F1: v3, F2: ""}
    v6 := []int{}
    fmt.Println(v1, v1)
    if false {
        v1 = true
        if (v1 || v1) {
            var v7 int16 = 718
            if v1 {
                for fi8 := 0; fi8 < 18; fi8++ {
                    if ((100 != 127) && (-2.25 > (3.14159 - 0.1))) {
                        v1 = v1
                        _ = len(v6)
                    } else {
                        switch fi8 {
                        case -2147483648:
                            var v9 bool = false
                            _ = v9
                        case -1:
                            _ = len(v2)
                        case -9223372036854775808:
                            var v10 uint8 = 0
                            _ = v10
                        }
                        if true {
                            fmt.Println(fi8, fi8)
                            fmt.Println(v7, fi8, v7)
                            var v11 uint16 = 151
                            _ = v11
                        } else {
                            switch "foo bar" {
                            case "":
                                fmt.Println(v7)
                            case "a":
                                if v1 {
                                    var v12 string = "a"
                                    v13 := GenS2{F0: 42, F1: v7}
                                    _ = v12
                                    _ = v13
                                }
                            }
                            var v14 uint8 = 128
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 10; fi15++ {
                            var v16 int = fi8
                            v17 := GenS2{F0: 32767, F1: v7}
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                    _ = fi8
                }
                if (!((65535 != 4294967295) && (128 == 0))) {
                    var v18 int16 = v7
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi19 := 0; fi19 < 11; fi19++ {
                        if (1e10 >= float32(-2.25)) {
                            v20 := GenS2{F0: 690, F1: (^(v7))}
                            _ = v20
                        } else {
                            if v1 {
                                switch "" {
                                case "\n\t":
                                    v7 += (v18 * (^(v18)))
                                default:
                                    switch "hello" {
                                    case "foo bar":
                                        v21 := GenS1{F0: 4294967295, F1: v3, F2: ""}
                                        _ = v21
                                    default:
                                        fmt.Println(v7, v3)
                                    }
                                }
                            }
                        }
                        _ = fi19
                    }
                    _ = v18
                }
                var v22 int32 = 2147483647
                _ = v22
            }
            _ = v7
        }
    }
    switch (("" + "foo bar") + "foo bar") {
    case "foo bar":
        v3 += (v3 + (v3 ^ genFn0(2147483647, 255, 0)))
    case "\n\t":
        for fi23 := 0; fi23 < 7; fi23++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v24 uint16 = 65535
            v25 := GenS2{F0: v24, F1: 97}
            _ = fi23
            _ = v24
            _ = v25
        }
    default:
        v26 := GenS3{F0: 0}
        _ = v26
    }
    v1 = ((!v1) || v1)
    var v27 uint8 = 188
    var v28 uint = 270
    var v29 uint16 = 42
    v28 -= v28
    var v30 float32 = 1e10
    var v31 int = -1
    fmt.Println(v3, v29, v31)
    var v32 string = (("" + ("foo bar" + "foo bar")) + "")
    if (v27 == v27) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v31 -= (v31 | 32768)
    } else {
        fmt.Println(v28)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v5
    _ = v6
}
