package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint16) string {
    var v4 int = 4294967295
    v4 += (^(v4))
    v4 += (v4 ^ (v4 + v4))
    _ = p0
    _ = v4
    return ""
}

func main() {
    var v1 int8 = -46
    var v2 int32 = -527
    v3 := []int{-9223372036854775808, -875, -2147483648, 127}
    var v5 int64 = -209
    v6 := GenS1{F0: 1, F1: 0, F2: v1}
    var v7 uint64 = 18446744073709551615
    if false {
        if len(v3) > 0 {
            v3[0] = -1
        }
    } else {
        if (int64(v7) < v5) {
            var v8 uint = 1
            v8 = (v8 & (v8 - 177))
            _ = v8
        }
        _ = len(v3)
    }
    v9 := GenS1{F0: v1, F1: -43, F2: v1}
    var v10 uint = 255
    var v11 string = (("hello" + "\n\t") + "hello")
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            switch (fi13 * fi13) {
            case 42:
                v10 += (v10 >> ((65535) & 63))
            case -442:
                v2 = v2
            default:
                var v14 string = ""
                _ = v14
            }
            _ = fi13
        }
        var v15 bool = (fi12 == fi12)
        _ = fi12
        _ = v15
    }
    v2 += 42
    switch (v11 + ("" + v11)) {
    case "foo bar":
        for fi16 := 0; fi16 < 11; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                switch (v11 + (((v11 + "hello") + (v11 + v11)) + "a")) {
                case "foo bar":
                    v18 := GenS2{F0: "hello", F1: 32767, F2: 135}
                    _ = v18
                case "a":
                    var v19 bool = (float64(0.1) == (3.14159 / (3.14159 - float64(1e10))))
                    _ = v19
                default:
                    if (!(0 < 42)) {
                        switch ("a" + v11) {
                        case "a":
                            if ((v5 >> (((v10 ^ v10)) & 63)) >= 9223372036854775807) {
                                v2 = (v2 % ((v2) | 1))
                                fmt.Println(v2, v7)
                                var v20 int32 = (v2 | 1)
                                _ = v20
                            }
                            fmt.Println(v10)
                            if len(v3) > 0 {
                                v3[0] = fi17
                            }
                        case "\n\t":
                            if (42 < 0) {
                                if ((1.5 < 1e10) && (255 < 0)) {
                                    var v21 float32 = 1.5
                                    switch (fi16 >> (((v10 ^ 2)) & 63)) {
                                    case 1:
                                        if (v21 != v21) {
                                            var v22 int32 = (v2 & 42)
                                            var v23 int16 = 1000
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                    v21 -= ((v21 + v21) + ((float32(0.1) - ((v21 / (v21 + (v21 + v21))) * v21)) / (3.14159 - v21)))
                                    var v24 uint32 = 2147483648
                                    v1 -= v1
                                    _ = v21
                                    _ = v24
                                } else {
                                    if (false && false) {
                                        if len(v3) > 0 {
                                            v3[0] = fi17
                                        }
                                    }
                                    v3 = append(v3, (fi16 & 0))
                                    switch (fi17 >> ((v10) & 63)) {
                                    case -9223372036854775808:
                                        if ("foo bar" >= v11) {
                                            fmt.Println(v5)
                                            v1 += (v1 << (((v10 | 4294967295)) & 63))
                                            if len(v3) > 0 {
                                                v3[0] = int(v10)
                                            }
                                        }
                                    case -2147483648:
                                        if false {
                                            switch (^(fi17)) {
                                            case 127:
                                                var v25 float64 = -2.25
                                                _ = v25
                                            }
                                        } else {
                                            v11 = v11
                                            switch (v11 + v11) {
                                            case "a":
                                                v10 += v10
                                            case "\n\t":
                                                v26 := GenS3{F0: v10}
                                                _ = v26
                                            case "hello":
                                                v7 += v7
                                            default:
                                                if true {
                                                    v27 := GenS1{F0: (v1 >> (((v10 % ((0) | 1))) & 63)), F1: (fi16 / ((-990) | 1)), F2: (-(v1))}
                                                    _ = v27
                                                }
                                            }
                                            fmt.Println(v11)
                                        }
                                        fmt.Println(fi17, v2)
                                    }
                                }
                            } else {
                                fmt.Println(v11)
                            }
                            if true {
                                var v28 bool = (!(!false))
                                var v29 int32 = v2
                                _ = v28
                                _ = v29
                            }
                            if false {
                                var v30 uint32 = 794
                                v31 := GenS4{F0: (^(v2))}
                                _ = v30
                                _ = v31
                            }
                            var v32 uint16 = 692
                            v33 := GenS3{F0: 65535}
                            v5 += (v5 % ((int64(fi17)) | 1))
                            var v34 int = 32768
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        v5 += v5
                    } else {
                        v11 = v11
                        fmt.Println(v5, v11, v10)
                        if false {
                            var v35 int64 = v5
                            fmt.Println(v1, fi17, v2)
                            _ = v35
                        }
                    }
                }
                _ = fi17
            }
            v36 := GenS4{F0: (-(v2))}
            _ = fi16
            _ = v36
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
