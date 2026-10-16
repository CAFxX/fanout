package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 string }
type GenS2 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 uint32) bool {
    v6 := []int{-850, -9223372036854775808}
    v7 := GenS1{F0: 0, F1: -758, F2: ("hello" + "")}
    p0 += (p0 - p0)
    _ = p0
    _ = v6
    _ = v7
    return ("" > ("hello" + (("\n\t" + "foo bar") + "hello")))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 int32 = 2147483647
    v4 := make(map[uint32]int)
    v5 := []uint32{2}
    v2[-192] = "a"
    v8 := GenS1{F0: 42, F1: 995, F2: ("" + ("hello" + "\n\t"))}
    var v9 int = 196
    v3 += v3
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(v3, fi10)
            v12 := v4[4294967295]
            _ = fi11
            _ = v12
        }
        v3 = v3
        v9 += (v9 / (((v9 - 127)) | 1))
        _ = fi10
    }
    fmt.Println(v3)
    _ = len(v4)
    v3 -= -1
    var v13 int8 = 1
    v13 = 0
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            fmt.Println(v3, v13, v3)
            v16 := GenS2{F0: 255, F1: 1, F2: 1}
            switch int(v13) {
            case 681:
                v3 = 0
            default:
                fmt.Println(v9, fi14, v3)
            }
            switch (fi15 >> ((18446744073709551615) & 63)) {
            case 1:
                v9 = 927
            case -2147483648:
                var v17 int64 = -9223372036854775808
                fmt.Println(fi15, v13)
                if (!(true || true)) {
                    switch "foo bar" {
                    case "":
                        var v18 bool = ((!false) || false)
                        _ = v18
                    default:
                        v9 = (fi15 + -2147483648)
                    }
                    v19 := GenS3{F0: 2, F1: 652}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v20 int8 = 1
                    _ = v19
                    _ = v20
                }
                v3 = (^(v3))
                switch ("a" + "\n\t") {
                case "a":
                    fmt.Println(v3)
                }
                _ = v17
            }
            v5 = append(v5, 2)
            switch "hello" {
            case "":
                v5 = append(v5, 842)
            default:
                if (!(false || true)) {
                    v1[42] = 0
                    _ = len(v5)
                    if false {
                        fmt.Println(fi15, fi14, fi15)
                        switch "\n\t" {
                        case "\n\t":
                            if len(v5) > 0 {
                                v5[0] = 4294967295
                            }
                        default:
                            switch 1 {
                            case 42:
                                var v21 bool = true
                                fmt.Println(v3, v13, fi15)
                                _ = v21
                            case 127:
                                v22 := GenS1{F0: 550, F1: -475, F2: "hello"}
                                _ = v22
                            case 9223372036854775807:
                                fmt.Println(fi15)
                            default:
                                v2[127] = "a"
                            }
                        }
                        var v23 int16 = int16(fi15)
                        var v24 float64 = 3.14159
                        var v25 uint = 919
                        _ = v23
                        _ = v24
                        _ = v25
                    } else {
                        var v26 int64 = 1
                        if (false && true) {
                            v9 = fi14
                            v3 -= 874
                        }
                        _ = v26
                    }
                } else {
                    if ((0.0 * -2.25) <= ((-2.25 + 1e10) - 1e10)) {
                        _ = len(v5)
                        var v27 bool = false
                        switch (fi15 % ((360) | 1)) {
                        case 9223372036854775807:
                            v3 = (v3 % ((v3) | 1))
                        }
                        v9 = (fi15 % ((fi15) | 1))
                        v9 = (v9 / ((-1) | 1))
                        switch (v9 >> ((65535) & 63)) {
                        case -2147483648:
                            _ = len(v1)
                        case 32768:
                            v27 = v27
                        }
                        var v28 float32 = 3.14159
                        _ = v27
                        _ = v28
                    } else {
                        v13 -= v13
                        var v29 bool = false
                        v13 -= (v13 & (v13 + (^(v13))))
                        _ = v29
                    }
                    var v30 uint64 = 994
                    v9 -= v9
                    _ = v30
                }
            }
            var v31 string = "a"
            _ = fi15
            _ = v16
            _ = v31
        }
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
