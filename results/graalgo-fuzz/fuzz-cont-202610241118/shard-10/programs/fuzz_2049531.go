package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool }

func main() {
    v1 := []int{128}
    v2 := make(map[int]string)
    var v3 int16 = 0
    var v4 uint64 = 4294967295
    var v5 uint32 = 42
    var v6 int8 = 42
    var v7 int64 = 9223372036854775807
    var v8 uint64 = v4
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            v6 += 127
            _ = fi10
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            switch fi11 {
            case 32767:
                v12 := GenS1{F0: v4, F1: v5, F2: 4294967295}
                _ = v12
            }
            var v13 int8 = v6
            var v14 bool = false
            _ = fi11
            _ = v13
            _ = v14
        }
        var v15 float32 = 0.1
        _ = fi9
        _ = v15
    }
    fmt.Println(v8, v8, v7)
    switch ("a" + "foo bar") {
    case "foo bar":
        if ((2 != 65535) || (!false)) {
            if (!(3.14159 != (float32(float64((0.0 - ((0.0 + 0.0) + (float32(3.14159) * (0.1 + (float32(float64((((-2.25 / 3.14159) / 1e10) * 1.5))) / (-2.25 - float32((float64(float32(3.14159)) / (0.1 * 1e10))))))))))) * 3.14159))) {
                if true {
                    var v16 uint = uint(v4)
                    fmt.Println(v6)
                    _ = v16
                } else {
                    var v17 int = 2147483647
                    _ = v17
                }
            } else {
                var v18 float64 = (1e10 * 1e10)
                _ = v18
            }
            for fi19 := 0; fi19 < 10; fi19++ {
                for fi20 := 0; fi20 < 15; fi20++ {
                    if (!(!false)) {
                        switch fi19 {
                        case 0:
                            v21 := GenS2{F0: 479, F1: ((v5 & 42) >= v5)}
                            _ = v21
                        case 127:
                            v22 := GenS2{F0: 65535, F1: true}
                            _ = v22
                        case 560:
                            var v23 uint8 = 1
                            _ = v23
                        default:
                            v4 = v8
                        }
                        if ((v5 + v5) >= 316) {
                            var v24 int64 = (^(v7))
                            v8 -= v8
                            _ = v24
                        } else {
                            fmt.Println(fi19)
                            var v25 uint32 = (v5 & v5)
                            switch (-(fi20)) {
                            case 42:
                                switch "a" {
                                case "":
                                    fmt.Println(v8)
                                }
                            case 0:
                                v3 = v3
                            case -9223372036854775808:
                                v1 = append(v1, -782)
                            default:
                                v6 += (v6 ^ 111)
                            }
                            _ = v25
                        }
                        fmt.Println(v4, fi19)
                        var v26 int8 = v6
                        _ = v26
                    } else {
                        fmt.Println(fi20, v8, v8)
                    }
                    v27 := v2[fi20]
                    if (false || true) {
                        var v28 float64 = (0.1 / (-2.25 * float64((0.0 - (0.0 + (1e10 - ((1.5 - (-2.25 - 1.5)) + (float32(1e10) - -2.25))))))))
                        var v29 string = (("foo bar" + ((v27 + "hello") + "")) + v27)
                        _ = v28
                        _ = v29
                    }
                    _ = fi20
                    _ = v27
                }
                v3 -= (v3 | (v3 / (((v3 / ((42) | 1))) | 1)))
                var v30 uint32 = v5
                _ = fi19
                _ = v30
            }
        } else {
            fmt.Println(v4)
            if true {
                v31 := GenS1{F0: 4294967295, F1: v5, F2: (^(v4))}
                _ = v31
            } else {
                v2[-501] = ""
                for fi32 := 0; fi32 < 10; fi32++ {
                    var v33 int16 = v3
                    var v34 uint = 372
                    for fi35 := 0; fi35 < 6; fi35++ {
                        _ = len(v2)
                        v36 := GenS3{F0: v3}
                        _ = fi35
                        _ = v36
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
        }
    case "hello":
        v37 := GenS4{F0: true}
        _ = v37
    case "\n\t":
        if (uint(v5) == uint(v4)) {
            var v38 uint8 = 7
            if (42 == 42) {
                var v39 bool = false
                delete(v2, int(v7))
                _ = v39
            }
            _ = v38
        }
    }
    var v40 float64 = -2.25
    var v41 float32 = 1.5
    var v42 uint32 = v5
    v6 -= int8(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
