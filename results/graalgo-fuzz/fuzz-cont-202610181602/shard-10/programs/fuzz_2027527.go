package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int; F1 int16; F2 int }
type GenS3 struct { F0 bool }
func genFn0() string {
    v6 := []int{4294967295, 1, -816, 128}
    var v7 uint16 = 2
    switch 4294967295 {
    case 4294967295:
        var v8 int = -2147483648
        _ = v8
    default:
        v7 -= (v7 - 0)
    }
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    var v1 int8 = -1
    var v2 uint16 = 2
    var v3 float32 = -2.25
    var v4 int = -9223372036854775808
    v5 := make(map[int]int)
    var v9 int16 = 32767
    v10 := []int{v4}
    fmt.Println(v2, v2)
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if len(v10) > 0 {
        v10[0] = (v4 % (((^(v4))) | 1))
    }
    var v11 uint = 2
    switch ("" + "foo bar") {
    case "":
        v12 := GenS1{F0: (^(v11)), F1: "hello"}
        _ = v12
    case "hello":
        v1 = (v1 % (((v1 | v1)) | 1))
    case "foo bar":
        v13 := GenS2{F0: v4, F1: v9, F2: (v4 - (v4 % ((int(v2)) | 1)))}
        _ = v13
    }
    v14 := GenS2{F0: (v4 / ((-621) | 1)), F1: 1, F2: (v4 >> ((v11) & 63))}
    switch (^(v4)) {
    case 128:
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS2{F0: v4, F1: v9, F2: (fi15 * (fi15 + -2147483648))}
            _ = fi15
            _ = v16
        }
    case 127:
        if (!false) {
            v3 = 3.14159
        } else {
            for fi17 := 0; fi17 < 10; fi17++ {
                if true {
                    if (!(-816 < -819)) {
                        if false {
                            if (!(((v2 - (v2 | (v2 & v2))) < (v2 % (((v2 | (v2 * 394))) | 1))) || ((!((178 > uint8(v2)) && (true && (-449 >= -453)))) && true))) {
                                var v18 uint64 = 444
                                _ = v18
                            } else {
                                for fi19 := 0; fi19 < 10; fi19++ {
                                    var v20 int32 = 0
                                    switch 1 {
                                    case 4294967295:
                                        if ((v20 >= v20) || true) {
                                            fmt.Println(v2, fi19, v4)
                                            if true {
                                                var v21 int16 = 0
                                                var v22 float32 = float32((3.14159 + (float64((((float32(((3.14159 + 1e10) / (3.14159 + 0.0))) + v3) / float32(-2.25)) * 3.14159)) * 1e10)))
                                                fmt.Println(v3)
                                                _ = v21
                                                _ = v22
                                            }
                                            var v23 int32 = -2147483648
                                            v4 = (v4 >> (((v11 << (((v11 ^ v11)) & 63))) & 63))
                                            _ = v23
                                        } else {
                                            v24 := GenS3{F0: true}
                                            var v25 bool = (v9 != 42)
                                            _ = v24
                                            _ = v25
                                        }
                                    default:
                                        switch "" {
                                        case "":
                                            v26 := GenS3{F0: false}
                                            _ = v26
                                        default:
                                            switch -1 {
                                            case 42:
                                                if (2 != 545) {
                                                    var v27 uint8 = 1
                                                    v4 = fi19
                                                    _ = v27
                                                }
                                            case 32767:
                                                var v28 uint8 = 127
                                                _ = v28
                                            }
                                        }
                                    }
                                    _ = fi19
                                    _ = v20
                                }
                            }
                            for fi29 := 0; fi29 < 3; fi29++ {
                                v10 = append(v10, (fi17 << (((v11 - v11)) & 63)))
                                v11 = v11
                                _ = fi29
                            }
                        } else {
                            for fi30 := 0; fi30 < 11; fi30++ {
                                var v31 uint8 = 128
                                fmt.Println(v4)
                                _ = fi30
                                _ = v31
                            }
                            v11 += v11
                            fmt.Println(v11, v9)
                        }
                        fmt.Println(v2, v3, v4)
                        var v32 bool = true
                        _ = v32
                    }
                    _ = len(v10)
                    v3 = 3.14159
                }
                _ = fi17
            }
            var v33 int16 = v9
            _ = v33
        }
    case 9223372036854775807:
        var v34 string = "foo bar"
        _ = v34
    }
    if false {
        v35 := GenS1{F0: (v11 << ((v11) & 63)), F1: "a"}
        var v36 uint8 = 102
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
