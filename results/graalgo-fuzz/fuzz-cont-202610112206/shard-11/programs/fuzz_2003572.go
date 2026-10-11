package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint32; F2 int }
type GenS4 struct { F0 uint64; F1 uint16; F2 int64 }

func main() {
    var v1 float64 = 1e10
    v2 := []int{}
    var v3 int = 4294967295
    v4 := make(map[int]int)
    v5 := []int{32767, 0}
    var v6 string = ""
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v1, v6)
        _ = fi7
    }
    var v8 uint8 = 128
    if (!false) {
        if len(v5) > 0 {
            v5[0] = v3
        }
        v9 := GenS1{F0: (^(v8)), F1: int8(v8), F2: -1}
        var v10 float32 = float32(0.1)
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 bool = false
            if v12 {
                var v13 int32 = 2147483647
                _ = v13
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 string = v6
                _ = fi14
                _ = v15
            }
            _ = fi11
            _ = v12
        }
        var v16 float64 = (((v1 / v1) - v1) + v1)
        v17 := GenS1{F0: v8, F1: 42, F2: 100}
        _ = v16
        _ = v17
    }
    if true {
        fmt.Println(v6)
    } else {
        fmt.Println(v3, v6, v3)
        v18 := GenS1{F0: (v8 * (v8 >> ((1) & 63))), F1: -102, F2: -128}
        v3 += int(v8)
        _ = v18
    }
    switch v3 {
    case 800:
        for fi19 := 0; fi19 < 13; fi19++ {
            switch 128 {
            case 32768:
                switch (fi19 + (fi19 << ((1) & 63))) {
                case 127:
                    if (v8 < (v8 << ((617) & 63))) {
                        switch (fi19 << ((127) & 63)) {
                        case 4294967295:
                            if false {
                                var v20 string = v6
                                _ = v20
                            } else {
                                v3 = (-(fi19))
                                for fi21 := 0; fi21 < 11; fi21++ {
                                    if true {
                                        switch (v6 + v6) {
                                        case "foo bar":
                                            switch v3 {
                                            case 0:
                                                if true {
                                                    var v22 float32 = 0.1
                                                    if true {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    v6 = v6
                                                    v1 -= float64(((v22 + (1.5 * v22)) / v22))
                                                    if false {
                                                        delete(v4, fi19)
                                                        v3 = fi19
                                                        v23 := GenS1{F0: 1, F1: 103, F2: int8(fi19)}
                                                        _ = v23
                                                    } else {
                                                        var v24 int = v3
                                                        v25 := GenS1{F0: (v8 >> ((42) & 63)), F1: 127, F2: -1}
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                    _ = v22
                                                }
                                            case 128:
                                                switch (fi19 * v3) {
                                                case 2147483647:
                                                    v26 := v4[9223372036854775807]
                                                    _ = v26
                                                default:
                                                    v6 = (v6 + (((v6 + v6) + v6) + v6))
                                                }
                                            case 9223372036854775807:
                                                v2 = append(v2, v3)
                                            default:
                                                v1 += 3.14159
                                            }
                                        case "\n\t":
                                            var v27 float64 = v1
                                            _ = v27
                                        }
                                    }
                                    switch v6 {
                                    case "hello":
                                        if len(v5) > 0 {
                                            v5[0] = -803
                                        }
                                    case "":
                                        var v28 int16 = -32768
                                        _ = v28
                                    case "\n\t":
                                        var v29 float32 = 3.14159
                                        _ = v29
                                    }
                                    _ = fi21
                                }
                            }
                        }
                        v30 := GenS2{F0: 4294967295}
                        _ = v30
                    }
                case 2147483647:
                    var v31 uint = 990
                    _ = v31
                }
            case 42:
                var v32 bool = false
                _ = v32
            default:
                var v33 int8 = -1
                _ = v33
            }
            v6 = v6
            _ = fi19
        }
    }
    v3 -= v3
    v4[v3] = v3
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 uint = 18446744073709551615
        v6 = "a"
        var v36 string = "hello"
        _ = fi34
        _ = v35
        _ = v36
    }
    if ((^(v8)) <= v8) {
        for fi37 := 0; fi37 < 1; fi37++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v8, v1)
            _ = fi37
        }
        fmt.Println(v3, v6)
        v38 := GenS2{F0: 0}
        _ = v38
    } else {
        fmt.Println(v6, v6)
        fmt.Println(v3, v3, v6)
    }
    v3 += v3
    switch int(v8) {
    case 2147483647:
        var v39 uint = uint(v8)
        _ = v39
    default:
        var v40 int64 = 1
        _ = v40
    }
    v3 -= v3
    v41 := GenS3{F0: -251, F1: 2147483648, F2: (v3 * v3)}
    v1 = (1.5 - v1)
    var v42 float64 = (v1 - v1)
    switch v6 {
    case "a":
        if false {
            v43 := GenS4{F0: 4294967295, F1: 1, F2: 198}
            _ = v43
        } else {
            if (!(((0 > 510) || true) || false)) {
                var v44 uint8 = 202
                fmt.Println(v42, v1)
                var v45 int64 = 42
                _ = v44
                _ = v45
            } else {
                var v46 uint16 = 255
                _ = v46
            }
            for fi47 := 0; fi47 < 1; fi47++ {
                v3 = (v3 - fi47)
                _ = fi47
            }
            var v48 string = (((("a" + (v6 + "")) + v6) + (v6 + ((v6 + (v6 + v6)) + ("hello" + v6)))) + (("a" + v6) + v6))
            _ = v48
        }
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        var v49 int64 = 311
        _ = v49
    }
    var v50 uint8 = (^(v8))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v50
    _ = v6
    _ = v8
}
