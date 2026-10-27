package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 1
    v3 := make(map[int]uint32)
    v4 := []int{v2}
    var v5 string = (("hello" + ("foo bar" + ("a" + "hello"))) + "foo bar")
    fmt.Println(v2, v5)
    var v6 uint = 68
    v7 := GenS1{F0: int16(v6), F1: (v6 + (v6 * (^(v6))))}
    v6 = 18446744073709551615
    v5 = ((("foo bar" + "") + "a") + (v5 + "hello"))
    for fi8 := 0; fi8 < 20; fi8++ {
        if false {
            v9 := GenS1{F0: -945, F1: 1}
            fmt.Println(v5)
            var v10 int32 = 1
            _ = v10
            _ = v9
        } else {
            v11 := GenS2{F0: (^(v2))}
            _ = v11
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int32 = 8
        var v14 uint8 = 0
        var v15 uint64 = 4294967295
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v6 = (v6 / (((v6 | 1)) | 1))
    switch (((v5 + ("foo bar" + "")) + v5) + (v5 + v5)) {
    case "\n\t":
        var v16 bool = ((((1e10 != 0.1) || (true && false)) && true) && (0.1 < 0.1))
        _ = v16
    case "hello":
        if (!(false || false)) {
            var v17 float32 = 1.5
            v5 = v5
            switch v5 {
            case "":
                v18 := v1[(-(v2))]
                _ = v18
            case "foo bar":
                if false {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        if (!false) {
                            var v20 string = "a"
                            v2 -= (v2 << (((^(v6))) & 63))
                            _ = v20
                        } else {
                            if len(v4) > 0 {
                                v4[0] = v2
                            }
                            if true {
                                fmt.Println(v2, fi19)
                                v2 -= 1
                            }
                        }
                        for fi21 := 0; fi21 < 3; fi21++ {
                            var v22 int8 = -97
                            v6 += v6
                            _ = fi21
                            _ = v22
                        }
                        _ = fi19
                    }
                } else {
                    v4 = append(v4, 1)
                    var v23 bool = false
                    _ = v23
                }
            default:
                fmt.Println(v6, v2, v17)
            }
            _ = v17
        }
    default:
        v2 = -427
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        switch ("hello" + v5) {
        case "":
            v25 := GenS3{F0: 2147483648, F1: 0}
            _ = v25
        case "\n\t":
            v26 := v3[fi24]
            _ = v26
        case "hello":
            var v27 int = (v2 * 930)
            _ = v27
        }
        v28 := GenS2{F0: (-(v2))}
        v29 := GenS1{F0: 1, F1: (^(v6))}
        _ = fi24
        _ = v28
        _ = v29
    }
    if ((v6 >> ((803) & 63)) > 1) {
        v30 := GenS2{F0: (v2 & (-(v2)))}
        var v31 int16 = 68
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
