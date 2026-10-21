package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint64, p1 bool, p2 string) uint32 {
    fmt.Println(p0, p1, p0)
    for fi4 := 0; fi4 < 3; fi4++ {
        fmt.Println(p1, p0, p0)
        var v5 float32 = float32(0.1)
        var v6 int32 = 0
        _ = fi4
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    v3 := []string{"\n\t"}
    switch -64 {
    case -2147483648:
        var v7 int = 136
        _ = v7
    }
    if (271 >= 615) {
        var v8 uint64 = 18446744073709551615
        v8 = (v8 ^ v8)
        var v9 int8 = 42
        _ = v8
        _ = v9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch ("a" + ("\n\t" + "")) {
    case "foo bar":
        var v10 int8 = -83
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 float32 = (0.1 / 1.0)
            if false {
                v12 += (1e10 / -2.25)
                for fi13 := 0; fi13 < 11; fi13++ {
                    if true {
                        var v14 int8 = -128
                        if (true || (float64(float32(1.5)) <= float64(v12))) {
                            fmt.Println(fi11, v14, v12)
                            switch -9223372036854775808 {
                            case 496:
                                switch (fi11 % (((-(fi11))) | 1)) {
                                case -1:
                                    if false {
                                        var v15 uint = 4294967295
                                        var v16 int8 = v14
                                        v17 := v2[(("foo bar" + "hello") + "hello")]
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                    } else {
                                        v18 := GenS1{F0: 1000, F1: "hello", F2: true}
                                        if (!(!true)) {
                                            switch fi13 {
                                            case 9223372036854775807:
                                                v19 := GenS2{F0: 32767, F1: false}
                                                _ = v19
                                            default:
                                                v20 := GenS3{F0: 527, F1: 0}
                                                _ = v20
                                            }
                                        }
                                        var v21 uint8 = 213
                                        _ = v18
                                        _ = v21
                                    }
                                case -625:
                                    switch "foo bar" {
                                    case "":
                                        v22 := GenS1{F0: 32767, F1: "", F2: false}
                                        _ = v22
                                    }
                                }
                            case -682:
                                v12 = v12
                            case 9223372036854775807:
                                switch fi13 {
                                case 2147483647:
                                    v23 := GenS2{F0: -32768, F1: true}
                                    _ = v23
                                case 991:
                                    delete(v1, 4294967295)
                                }
                            default:
                                switch (fi13 * fi11) {
                                case -9223372036854775808:
                                    v12 -= (1e10 * (v12 * float32(((float64(v12) * float64(v12)) / (1.5 + float64(3.14159))))))
                                case 1:
                                    v24 := GenS3{F0: 32767, F1: 255}
                                    var v25 uint32 = 1
                                    var v26 float64 = 1.5
                                    switch (^(fi13)) {
                                    case -9223372036854775808:
                                        _ = len(v2)
                                    case 914:
                                        var v27 uint16 = 42
                                        _ = v27
                                    default:
                                        var v28 uint16 = 65535
                                        _ = v28
                                    }
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                case -1:
                                    var v29 float32 = v12
                                    _ = v29
                                default:
                                    var v30 int = (fi11 & fi13)
                                    _ = v30
                                }
                            }
                        } else {
                            var v31 bool = (42 == 18446744073709551615)
                            var v32 string = "\n\t"
                            _ = v31
                            _ = v32
                        }
                        _ = v14
                    }
                    switch (fi13 - (fi11 - 807)) {
                    case -1:
                        fmt.Println(fi11, v12, v12)
                    case 1:
                        var v33 uint = 2
                        _ = v33
                    case 32768:
                        v34 := GenS3{F0: 2, F1: 128}
                        _ = len(v2)
                        _ = v34
                    }
                    _ = fi13
                }
                v12 += (((v12 * -2.25) * -2.25) - 3.14159)
            } else {
                fmt.Println(fi11, fi11)
            }
            _ = fi11
            _ = v12
        }
    }
    v35 := GenS1{F0: 42, F1: "hello", F2: ((1.5 * (float32((((float64(3.14159) / 0.1) / (float64(float32(1e10)) + 0.1)) / ((3.14159 / ((0.0 / float64(3.14159)) + 1.5)) * float64((0.1 + float32((0.0 / 0.1))))))) + ((0.1 - 0.0) + 0.0))) <= 1.5)}
    for fi36 := 0; fi36 < 1; fi36++ {
        var v37 float32 = ((0.0 * 3.14159) * float32(0.1))
        _ = fi36
        _ = v37
    }
    var v38 uint32 = genFn0(2, false, "foo bar")
    switch -1 {
    case -9223372036854775808:
        v38 = (v38 + v38)
    }
    var v39 string = "foo bar"
    var v40 uint16 = 42
    v40 += (v40 / ((798) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v40
}
