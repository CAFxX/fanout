package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint }
func genFn0(p0 string, p1 int32, p2 int64) uint8 {
    v7 := GenS1{F0: 83, F1: 826}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}

func main() {
    var v1 int32 = 828
    var v2 float32 = 3.14159
    v3 := make(map[string]uint32)
    v4 := []uint32{474}
    var v5 float32 = float32((1.5 - 0.1))
    var v6 uint64 = 42
    v8 := GenS1{F0: (-(v1)), F1: 126}
    v9 := GenS1{F0: v1, F1: 255}
    fmt.Println(v1, v5, v1)
    v1 = (v1 ^ (v1 << ((0) & 63)))
    var v10 string = ("foo bar" + "a")
    _ = len(v3)
    var v11 uint32 = 98
    switch ("hello" + "") {
    case "foo bar":
        if (!true) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case "a":
        v12 := GenS2{F0: 0}
        _ = v12
    default:
        fmt.Println(v11)
    }
    if true {
        v13 := v3[(v10 + v10)]
        v10 = v10
        _ = v13
    }
    switch "hello" {
    case "hello":
        var v14 uint16 = 167
        _ = v14
    case "\n\t":
        var v15 int16 = 1
        _ = v15
    default:
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS1{F0: v1, F1: 255}
            var v18 string = "a"
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    fmt.Println(v5, v1, v2)
    v19 := GenS3{F0: 987, F1: 65535}
    switch 4294967295 {
    case 42:
        fmt.Println(v1)
    }
    if (349 == 743) {
        var v20 int64 = 0
        fmt.Println(v11, v5, v10)
        if (1 <= v1) {
            switch ("\n\t" + v10) {
            case "\n\t":
                if true {
                    var v21 int64 = v20
                    v20 = (^(v20))
                    _ = v21
                } else {
                    var v22 float32 = (((3.14159 * (1.5 + float32(((0.1 / 1.0) / (((3.14159 - 3.14159) + float64(v5)) - ((((0.1 * (1e10 + (3.14159 * 0.0))) * (float64(v5) + 3.14159)) * 0.1) + 0.0)))))) - -2.25) * float32(3.14159))
                    _ = v22
                }
            }
            var v23 uint32 = (^(v11))
            _ = v23
        } else {
            var v24 uint = 65535
            if true {
                if false {
                    fmt.Println(v2, v10, v10)
                    if ((-594 == -1) && false) {
                        var v25 int8 = 96
                        var v26 uint16 = 0
                        if true {
                            v11 += v11
                            if (!(uint(v26) != (v24 << (((v24 % ((v24) | 1))) & 63)))) {
                                v5 -= float32(3.14159)
                            } else {
                                var v27 int = 42
                                if false {
                                    var v28 bool = (-32768 <= -289)
                                    v28 = true
                                    switch v27 {
                                    case -2147483648:
                                        for fi29 := 0; fi29 < 18; fi29++ {
                                            for fi30 := 0; fi30 < 16; fi30++ {
                                                var v31 int64 = (v20 | -9223372036854775808)
                                                _ = fi30
                                                _ = v31
                                            }
                                            if false {
                                                if v28 {
                                                    fmt.Println(v1, v10, v5)
                                                }
                                                if v28 {
                                                    v5 += v5
                                                    if len(v4) > 0 {
                                                        v4[0] = (v11 % ((v11) | 1))
                                                    }
                                                } else {
                                                    for fi32 := 0; fi32 < 10; fi32++ {
                                                        var v33 uint64 = v6
                                                        _ = fi32
                                                        _ = v33
                                                    }
                                                }
                                                v34 := GenS3{F0: v26, F1: 629}
                                                _ = v34
                                            }
                                            _ = fi29
                                        }
                                    default:
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(v10)
                                        default:
                                            var v35 uint64 = (v6 % (((v6 * v6)) | 1))
                                            _ = v35
                                        }
                                    }
                                    _ = v28
                                }
                                _ = v27
                            }
                            var v36 int64 = (v20 << ((65535) & 63))
                            _ = v36
                        }
                        _ = v25
                        _ = v26
                    }
                }
                var v37 uint64 = v6
                _ = v37
            } else {
                if (v6 != (v6 | v6)) {
                    var v38 int64 = v20
                    v20 = 0
                    _ = v38
                }
                var v39 uint = v24
                _ = v39
            }
            var v40 uint = (^(v24))
            _ = v24
            _ = v40
        }
        _ = v20
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        var v42 uint = 831
        _ = fi41
        _ = v42
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
