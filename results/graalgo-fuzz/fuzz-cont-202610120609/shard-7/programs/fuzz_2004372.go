package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int64; F2 string }
type GenS5 struct { F0 bool; F1 int32 }

func main() {
    var v1 int = 127
    var v2 int64 = 589
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    var v6 float32 = ((3.14159 / 1.0) - ((((1.5 - 0.1) * (1.5 + 3.14159)) / 0.1) * 1.5))
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := v5[(v1 - (fi7 ^ (fi8 ^ fi7)))]
            v10 := GenS1{F0: v2, F1: v1}
            var v11 uint64 = 18446744073709551615
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        fmt.Println(fi7)
        var v12 uint8 = 42
        _ = fi7
        _ = v12
    }
    fmt.Println(v1, v1)
    switch v1 {
    case 128:
        fmt.Println(v1)
    case 2147483647:
        v1 -= -720
    case -42:
        var v13 int8 = 0
        _ = v13
    default:
        v14 := []int{}
        _ = v14
    }
    delete(v4, ("foo bar" + "a"))
    v1 = (v1 << ((584) & 63))
    v1 = v1
    v2 += 1
    if (((!true) || (!false)) && true) {
        var v15 int = 9223372036854775807
        var v16 uint32 = 0
        _ = v15
        _ = v16
    }
    switch "hello" {
    case "\n\t":
        for fi17 := 0; fi17 < 8; fi17++ {
            if false {
                v2 += -1
                v18 := GenS2{F0: -56}
                switch (("foo bar" + "hello") + ("hello" + "foo bar")) {
                case "hello":
                    v19 := GenS2{F0: -2147483648}
                    _ = v19
                case "foo bar":
                    v1 += (v1 ^ (fi17 >> ((18446744073709551615) & 63)))
                case "":
                    for fi20 := 0; fi20 < 10; fi20++ {
                        var v21 int8 = -128
                        v22 := GenS3{F0: ""}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                default:
                    if (-1 <= 284) {
                        var v23 float32 = (1.5 + (v6 - float32((float64(1.5) - (1e10 / -2.25)))))
                        v24 := []int{-745}
                        var v25 float64 = ((float64((v23 / ((float32(((0.1 / (1e10 - 0.1)) + 3.14159)) / 3.14159) - -2.25))) + ((1e10 * -2.25) / (-2.25 * 1e10))) * 1e10)
                        _ = v23
                        _ = v24
                        _ = v25
                    } else {
                        if (65535 <= 639) {
                            v26 := GenS4{F0: 2, F1: (-(v2)), F2: "\n\t"}
                            fmt.Println(fi17, v1)
                            _ = v26
                        } else {
                            if (32767 == 676) {
                                if ((-1 > -2147483648) && (0 >= 1)) {
                                    var v27 float32 = ((v6 + v6) + float32(3.14159))
                                    _ = v27
                                }
                            } else {
                                v1 -= fi17
                                _ = len(v4)
                                for fi28 := 0; fi28 < 6; fi28++ {
                                    v6 -= v6
                                    fmt.Println(v6)
                                    fmt.Println(v6, fi28)
                                    if true {
                                        v29 := []int{v1, 32768, (fi28 - 900), fi17}
                                        var v30 int8 = 100
                                        _ = len(v29)
                                        if (true || true) {
                                            if ((true && (1 != -442)) || true) {
                                                var v31 uint64 = 0
                                                _ = v31
                                            }
                                            v30 = v30
                                            fmt.Println(v30, v1, fi17)
                                        }
                                        var v32 uint8 = 1
                                        _ = v29
                                        _ = v30
                                        _ = v32
                                    } else {
                                        var v33 int8 = -1
                                        _ = v33
                                    }
                                    v4[""] = true
                                    _ = fi28
                                }
                            }
                            for fi34 := 0; fi34 < 14; fi34++ {
                                v2 += v2
                                v6 += (3.14159 - float32(3.14159))
                                switch v1 {
                                case -630:
                                    v35 := GenS1{F0: (v2 % (((v2 | v2)) | 1)), F1: fi34}
                                    _ = v35
                                case 32767:
                                    v6 = (((v6 * v6) - v6) + (v6 - v6))
                                default:
                                    v36 := GenS2{F0: -1}
                                    _ = v36
                                }
                                _ = fi34
                            }
                        }
                    }
                }
                _ = v18
            } else {
                var v37 uint = 255
                v1 += -765
                _ = v37
            }
            _ = fi17
        }
    case "foo bar":
        fmt.Println(v1)
    default:
        if false {
            _ = len(v5)
        } else {
            v2 += 9223372036854775807
            var v38 float32 = (((v6 + v6) - v6) - ((v6 - v6) * (v6 + v6)))
            _ = v38
        }
    }
    v39 := GenS3{F0: ""}
    var v40 string = "foo bar"
    v41 := GenS5{F0: false, F1: -92}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
