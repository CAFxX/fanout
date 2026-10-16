package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
func genFn0() int8 {
    var v6 float32 = float32(1e10)
    fmt.Println(v6)
    for fi7 := 0; fi7 < 2; fi7++ {
        v6 += 0.0
        v6 += v6
        _ = fi7
    }
    _ = v6
    return -128
}

func main() {
    var v1 int = -2147483648
    var v2 float64 = (float64(1.5) + 0.0)
    var v3 uint64 = 361
    v4 := []int{(-(v1)), 2147483647, v1, v1}
    var v5 int8 = -128
    for fi8 := 0; fi8 < 12; fi8++ {
        _ = len(v4)
        if false {
            v2 += v2
        }
        _ = fi8
    }
    v1 = 42
    v9 := GenS1{F0: 899}
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS2{F0: 285}
        _ = fi10
        _ = v11
    }
    if (!(119 < 4294967295)) {
        v4 = append(v4, (v1 * v1))
    }
    v1 += v1
    fmt.Println(v3)
    switch (v1 >> ((0) & 63)) {
    case -867:
        var v12 string = "hello"
        _ = v12
    case 4294967295:
        fmt.Println(v2, v5)
    case -2147483648:
        var v13 float32 = -2.25
        _ = v13
    default:
        if false {
            if len(v4) > 0 {
                v4[0] = (-(v1))
            }
        } else {
            for fi14 := 0; fi14 < 18; fi14++ {
                if ((0.1 * 0.1) <= ((float32((0.0 + 1e10)) + float32((float64(float32((v2 * v2))) - 1.5))) + ((-2.25 * 1e10) + (0.1 / 0.1)))) {
                    if true {
                        fmt.Println(v2, v5, v3)
                        _ = len(v4)
                        switch ("hello" + "") {
                        case "a":
                            if true {
                                var v15 float64 = float64(3.14159)
                                var v16 float64 = 0.0
                                _ = v15
                                _ = v16
                            } else {
                                if true {
                                    if (!true) {
                                        v3 = 2
                                    }
                                }
                                v1 = (fi14 + fi14)
                            }
                        case "":
                            for fi17 := 0; fi17 < 19; fi17++ {
                                if (true || false) {
                                    var v18 uint8 = 127
                                    if false {
                                        v18 = v18
                                    } else {
                                        if ((("a" + "\n\t") + "foo bar") == "hello") {
                                            var v19 int16 = 369
                                            var v20 int64 = 1
                                            _ = v19
                                            _ = v20
                                        } else {
                                            v1 = (fi14 - fi14)
                                            _ = len(v4)
                                            v18 += (v18 / (((^(v18))) | 1))
                                        }
                                        if true {
                                            v3 = v3
                                            fmt.Println(v3, v5)
                                        }
                                        switch fi17 {
                                        case 128:
                                            fmt.Println(v2, v1)
                                        case -94:
                                            fmt.Println(fi17, v1)
                                        case 32767:
                                            fmt.Println(v18)
                                        default:
                                            v5 -= 127
                                        }
                                    }
                                    _ = v18
                                } else {
                                    var v21 string = "a"
                                    v5 = (v5 / (((v5 & v5)) | 1))
                                    _ = v21
                                }
                                switch 1 {
                                case 170:
                                    if true {
                                        _ = len(v4)
                                        if true {
                                            v22 := GenS2{F0: -1}
                                            if ((!(false && (!(true || (!(0 <= (v5 >> ((18446744073709551615) & 63)))))))) && (42 < 241)) {
                                                fmt.Println(v3)
                                                v4 = append(v4, (fi17 - (v1 * fi17)))
                                            } else {
                                                var v23 bool = (1 <= (^(v1)))
                                                var v24 int = v1
                                                _ = v23
                                                _ = v24
                                            }
                                            v1 += v1
                                            var v25 int16 = -32768
                                            fmt.Println(fi14, v25)
                                            _ = v22
                                            _ = v25
                                        }
                                        var v26 uint = 65535
                                        v27 := GenS1{F0: -135}
                                        switch fi14 {
                                        case 127:
                                            switch (v1 & (fi17 & (v1 % ((v1) | 1)))) {
                                            case 583:
                                                v28 := GenS1{F0: 9223372036854775807}
                                                _ = v28
                                            case -428:
                                                v29 := GenS2{F0: 557}
                                                _ = v29
                                            }
                                        default:
                                            var v30 string = ""
                                            _ = v30
                                        }
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                if (32767 > 0) {
                                    v31 := GenS2{F0: -122}
                                    fmt.Println(fi17, v5, fi14)
                                    _ = v31
                                } else {
                                    v3 = 0
                                    _ = len(v4)
                                    var v32 int = 1
                                    _ = v32
                                }
                                _ = fi17
                            }
                        default:
                            fmt.Println(v1)
                        }
                    }
                    for fi33 := 0; fi33 < 9; fi33++ {
                        if true {
                            v5 = v5
                            if (0 >= int64(v5)) {
                                v2 = 1e10
                                v34 := GenS3{F0: 2}
                                v4 = append(v4, (v1 & 127))
                                _ = v34
                            }
                        }
                        _ = fi33
                    }
                }
                _ = fi14
            }
            var v35 int = -79
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
