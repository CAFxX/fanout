package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 bool }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    var v4 int64 = -653
    v5 := []string{"hello", ((("hello" + "a") + "a") + "\n\t")}
    v1 = v1
    v4 = (v4 << ((42) & 63))
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int32 = 0
        v6 = (1.5 / (3.14159 * float32(float64(v6))))
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(v8)
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    switch 32768 {
    case 9223372036854775807:
        var v10 uint = 1
        _ = v10
    case 0:
        v11 := v3[1]
        _ = v11
    case 32768:
        fmt.Println(v4)
    default:
        if (true || ("hello" < ("" + "a"))) {
            var v12 int32 = 148
            if (-74 != -85) {
                for fi13 := 0; fi13 < 10; fi13++ {
                    v6 += -2.25
                    v14 := GenS1{F0: "foo bar"}
                    _ = fi13
                    _ = v14
                }
                v15 := v3[2147483648]
                switch -2147483648 {
                case 4294967295:
                    if ((v4 < (v4 + (v4 / (((v4 / ((9223372036854775807) | 1))) | 1)))) && ((v12 / ((v12) | 1)) > int32(v4))) {
                        if false {
                            var v16 int32 = (v12 % (((v12 - int32(v4))) | 1))
                            v1 = v1
                            for fi17 := 0; fi17 < 16; fi17++ {
                                if false {
                                    v1 -= v1
                                    if true {
                                        var v18 string = "foo bar"
                                        switch v18 {
                                        case "foo bar":
                                            fmt.Println(v16)
                                        case "":
                                            var v19 int8 = 104
                                            _ = v19
                                        default:
                                            v20 := GenS1{F0: ("hello" + (v15 + "\n\t"))}
                                            _ = v20
                                        }
                                        _ = v18
                                    } else {
                                        fmt.Println(v16, fi17)
                                        switch (-(fi17)) {
                                        case 42:
                                            var v21 uint64 = 42
                                            _ = v21
                                        case 648:
                                            v6 = v6
                                        default:
                                            if false {
                                                switch "hello" {
                                                case "a":
                                                    v22 := GenS2{F0: 1000, F1: (false && false)}
                                                    _ = v22
                                                }
                                                if (true && false) {
                                                    var v23 float32 = v6
                                                    _ = v23
                                                }
                                                v24 := GenS1{F0: ((v15 + "a") + v15)}
                                                _ = v24
                                            } else {
                                                switch (fi17 & 1) {
                                                case -771:
                                                    v1 = float64(float32(v1))
                                                case 32767:
                                                    _ = len(v3)
                                                case 792:
                                                    var v25 bool = (v1 == (v1 / (v1 + (0.1 + 3.14159))))
                                                    _ = v25
                                                }
                                                for fi26 := 0; fi26 < 18; fi26++ {
                                                    var v27 uint8 = 2
                                                    _ = fi26
                                                    _ = v27
                                                }
                                            }
                                        }
                                    }
                                    var v28 float32 = ((v6 - v6) - float32(v1))
                                    _ = v28
                                }
                                v29 := GenS3{F0: false}
                                _ = fi17
                                _ = v29
                            }
                            _ = v16
                        } else {
                            fmt.Println(v15, v4)
                            var v30 int16 = 730
                            switch ((v15 + v15) + (v15 + "foo bar")) {
                            case "a":
                                fmt.Println(v1)
                            default:
                                fmt.Println(v12, v1)
                            }
                            _ = v30
                        }
                    } else {
                        var v31 float32 = float32(-2.25)
                        v32 := GenS1{F0: v15}
                        _ = v31
                        _ = v32
                    }
                case 32767:
                    var v33 string = v15
                    _ = v33
                }
                _ = v15
            }
            _ = v12
        }
    }
    v4 += v4
    v2[2147483647] = 427
    fmt.Println(v4, v1)
    v1 = float64(v6)
    if ((false && false) && (-1 <= (v4 >> ((255) & 63)))) {
        v6 = v6
    }
    if (!true) {
        var v34 uint32 = 65535
        v35 := GenS2{F0: int16(v34), F1: true}
        var v36 int16 = 1000
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
