package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 string = "\n\t"
    var v2 uint64 = 4294967295
    var v3 int32 = 2147483647
    v1 = v1
    var v4 bool = true
    v2 = 0
    v1 = ((("" + v1) + v1) + "\n\t")
    if v4 {
        for fi5 := 0; fi5 < 7; fi5++ {
            if false {
                v6 := []int{1, 9223372036854775807}
                var v7 uint16 = 640
                if v4 {
                    var v8 uint = 18446744073709551615
                    v4 = (255 < 255)
                    _ = v8
                } else {
                    v9 := GenS1{F0: 27, F1: 65, F2: (v3 << ((65535) & 63))}
                    _ = v9
                }
                _ = v6
                _ = v7
            } else {
                v10 := GenS1{F0: -1, F1: 255, F2: (v3 * (v3 / ((0) | 1)))}
                fmt.Println(v4)
                _ = v10
            }
            _ = fi5
        }
        var v11 float32 = (0.0 - (((1.5 / 1.0) + ((3.14159 + 1e10) * (0.0 * 0.0))) + 3.14159))
        _ = v11
    } else {
        var v12 float64 = (3.14159 + 1e10)
        if v4 {
            for fi13 := 0; fi13 < 11; fi13++ {
                if v4 {
                    v14 := []int{2147483647, fi13}
                    _ = v14
                }
                v1 = v1
                var v15 float32 = (1.5 + float32(0.0))
                _ = fi13
                _ = v15
            }
        }
        _ = v12
    }
    v16 := []int{127, -1}
    v17 := GenS1{F0: 100, F1: 249, F2: v3}
    v18 := GenS2{F0: 42, F1: 0, F2: (!false)}
    switch (("foo bar" + v1) + "foo bar") {
    case "":
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        fmt.Println(v2)
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 int8 = 69
            v1 = "hello"
            _ = fi20
            _ = v21
        }
        var v22 uint64 = v2
        _ = fi19
        _ = v22
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        if v4 {
            if (v1 > "hello") {
                var v24 int16 = -1
                if ((((1e10 / 1.5) * float32(-2.25)) + (1e10 / (((0.1 - 0.1) - 1.5) / 1e10))) == 0.0) {
                    v24 = v24
                    if v4 {
                        for fi25 := 0; fi25 < 8; fi25++ {
                            v2 = v2
                            v1 = (v1 + v1)
                            v26 := GenS3{F0: v3}
                            _ = fi25
                            _ = v26
                        }
                        var v27 uint32 = 65535
                        v28 := GenS3{F0: v3}
                        _ = v27
                        _ = v28
                    } else {
                        var v29 uint32 = 876
                        for fi30 := 0; fi30 < 12; fi30++ {
                            var v31 float32 = ((3.14159 * (((-2.25 - 3.14159) + 1.5) + 1e10)) / 1.5)
                            v32 := GenS4{F0: 42, F1: 250}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        _ = v29
                    }
                    v24 = v24
                }
                v1 = (v1 + "\n\t")
                _ = v24
            } else {
                var v33 uint = 2
                _ = v33
            }
        }
        var v34 bool = (v4 && true)
        _ = fi23
        _ = v34
    }
    var v35 uint8 = 1
    v3 = (v3 ^ (v3 - v3))
    fmt.Println(v3)
    switch v1 {
    case "hello":
        if len(v16) > 0 {
            v16[0] = 9223372036854775807
        }
    default:
        for fi36 := 0; fi36 < 16; fi36++ {
            switch (fi36 >> ((42) & 63)) {
            case 9223372036854775807:
                v37 := GenS2{F0: -1, F1: 65535, F2: (v4 || v4)}
                _ = v37
            }
            fmt.Println(v35, v4)
            _ = fi36
        }
    }
    switch v1 {
    case "foo bar":
        var v38 uint16 = 110
        _ = v38
    case "":
        var v39 uint8 = v35
        _ = v39
    default:
        switch v1 {
        case "foo bar":
            var v40 int16 = -974
            _ = v40
        case "\n\t":
            v41 := GenS3{F0: (v3 | v3)}
            _ = v41
        case "hello":
            var v42 int8 = 0
            _ = v42
        }
    }
    v35 += v35
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
