package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
func genFn0() int16 {
    var v5 float64 = 3.14159
    var v6 bool = (("hello" + ("hello" + ("foo bar" + "hello"))) >= "foo bar")
    if ((v5 / 0.1) <= 0.1) {
        v7 := GenS1{F0: 32767}
        v5 += (v5 / v5)
        _ = v7
    } else {
        var v8 bool = ((2147483648 < 2147483647) || v6)
        var v9 float64 = v5
        var v10 uint64 = 393
        _ = v10
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 42
    v3 := []string{}
    var v4 uint32 = 0
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            v4 -= v4
            switch -338 {
            case 0:
                v2 -= v2
                delete(v1, v4)
            default:
                switch 540 {
                case 246:
                    if false {
                        switch (fi11 ^ (fi11 * (fi11 + fi11))) {
                        case 2147483647:
                            delete(v1, 2)
                        case 4294967295:
                            switch "foo bar" {
                            case "foo bar":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        default:
                            var v13 uint64 = uint64(v4)
                            switch "\n\t" {
                            case "":
                                fmt.Println(v2)
                            case "\n\t":
                                var v14 uint64 = v13
                                _ = v14
                            default:
                                var v15 uint8 = 127
                                _ = v15
                            }
                            v16 := GenS1{F0: 237}
                            var v17 int64 = 0
                            if (true || false) {
                                v2 -= (v2 | (v2 >> ((581) & 63)))
                            }
                            _ = v13
                            _ = v16
                            _ = v17
                        }
                        if ((-1 == -1) || true) {
                            v2 -= (v2 ^ (v2 + v2))
                            if false {
                                if false {
                                    var v18 int8 = (v2 | v2)
                                    switch (fi12 - fi11) {
                                    case 0:
                                        if ("foo bar" == "a") {
                                            v3 = append(v3, "a")
                                            if true {
                                                var v19 uint = 65535
                                                _ = v19
                                            } else {
                                                v20 := GenS1{F0: 32767}
                                                var v21 int = (fi12 ^ (fi12 * (fi12 >> ((255) & 63))))
                                                _ = v20
                                                _ = v21
                                            }
                                        } else {
                                            v22 := GenS2{F0: "hello", F1: 2147483647}
                                            v23 := GenS3{F0: -1, F1: -2147483648, F2: (v18 * v2)}
                                            v24 := v1[v4]
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        }
                                    default:
                                        var v25 uint = 528
                                        _ = v25
                                    }
                                    _ = v18
                                } else {
                                    if (!((v4 / (((v4 >> ((521) & 63))) | 1)) >= v4)) {
                                        var v26 string = "foo bar"
                                        _ = v26
                                    }
                                    v27 := GenS1{F0: -1}
                                    _ = v27
                                }
                            } else {
                                fmt.Println(v4, v2, fi12)
                                var v28 int64 = int64(v2)
                                v28 = 42
                                _ = v28
                            }
                        } else {
                            var v29 int64 = -9223372036854775808
                            switch "foo bar" {
                            case "foo bar":
                                v30 := GenS3{F0: 2147483647, F1: -788, F2: (v2 ^ int8(v29))}
                                _ = v30
                            case "hello":
                                _ = len(v1)
                            default:
                                v29 = (v29 | v29)
                            }
                            var v31 uint32 = (v4 - (v4 & 42))
                            _ = v29
                            _ = v31
                        }
                        var v32 uint = 167
                        _ = v32
                    }
                case 32767:
                    fmt.Println(v4, fi11, fi11)
                }
            }
            if ((0 <= 128) || true) {
                v33 := GenS2{F0: "a", F1: 2147483647}
                _ = v33
            }
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println(v2, v4)
    v2 = (v2 ^ (-(v2)))
    v2 = -128
    v2 = v2
    v1[(v4 << ((4294967295) & 63))] = -1
    v34 := GenS1{F0: 802}
    delete(v1, v4)
    fmt.Println(v4, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
