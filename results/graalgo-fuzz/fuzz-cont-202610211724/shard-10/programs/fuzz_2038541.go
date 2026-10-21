package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint32; F1 int64; F2 int64 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int64; F2 uint16 }
func genFn0(p0 int32, p1 string) bool {
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: 18446744073709551615, F1: 32767}
        if (((-123 > 0) && false) && false) {
            p0 += p0
            for fi6 := 0; fi6 < 18; fi6++ {
                p0 = p0
                var v7 uint64 = 42
                _ = fi6
                _ = v7
            }
        }
        var v8 int16 = 80
        _ = fi4
        _ = v5
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        switch p1 {
        case "foo bar":
            p0 -= (p0 ^ (p0 + p0))
        case "":
            fmt.Println(fi9, p0, fi9)
        case "a":
            var v10 string = ((("hello" + (("foo bar" + "hello") + "hello")) + ("\n\t" + p1)) + "hello")
            _ = v10
        default:
            fmt.Println(fi9)
        }
        switch (-(fi9)) {
        case 4294967295:
            p0 += 1
        }
        _ = fi9
    }
    p0 = -1
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 bool = (-1 == 496)
    var v2 int = -605
    var v3 uint8 = 255
    v2 -= v2
    v11 := GenS2{F0: 852, F1: 305, F2: -979}
    v1 = ((!v1) || v1)
    var v12 int16 = -831
    switch (v2 << ((42) & 63)) {
    case -2147483648:
        v2 -= v2
    case -121:
        fmt.Println(v2, v2)
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v2 -= 319
        fmt.Println(v3, v2)
        _ = fi13
    }
    v14 := []int{(v2 >> ((255) & 63)), 122}
    switch (("foo bar" + "a") + ("\n\t" + "foo bar")) {
    case "hello":
        v3 -= v3
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            _ = len(v14)
            v17 := GenS2{F0: 399, F1: 9223372036854775807, F2: -219}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            var v18 uint = 818
            v19 := GenS3{F0: (v18 + 65535), F1: 65535, F2: v18}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        if (!(true || (2147483647 == 2147483647))) {
            if len(v14) > 0 {
                v14[0] = (v2 / (((fi15 % ((128) | 1))) | 1))
            }
            if true {
                var v20 uint64 = 30
                v3 += v3
                switch ((("a" + (("foo bar" + "hello") + "\n\t")) + "") + "\n\t") {
                case "\n\t":
                    for fi21 := 0; fi21 < 19; fi21++ {
                        if true {
                            var v22 uint64 = (v20 ^ 0)
                            if (true && (v3 <= 147)) {
                                v23 := GenS2{F0: 570, F1: 148, F2: -177}
                                switch "" {
                                case "a":
                                    if ((fi15 % ((fi21) | 1)) >= v2) {
                                        var v24 uint16 = 255
                                        if v1 {
                                            if (uint32(v24) >= 543) {
                                                v12 -= v12
                                                _ = len(v14)
                                                var v25 uint = 0
                                                var v26 uint32 = 942
                                                _ = v25
                                                _ = v26
                                            }
                                        }
                                        _ = v24
                                    } else {
                                        switch v2 {
                                        case 42:
                                            if len(v14) >= 2 {
                                                v14 = v14[1:len(v14)]
                                            }
                                        case 4294967295:
                                            if (!v1) {
                                                switch fi15 {
                                                case 32767:
                                                    v1 = (134 >= (v20 ^ v20))
                                                }
                                                var v27 uint = 18446744073709551615
                                                _ = v27
                                            }
                                        default:
                                            v28 := GenS4{F0: (v1 && v1), F1: 4294967295}
                                            _ = v28
                                        }
                                        if (4294967295 > 4294967295) {
                                            v2 = int(v22)
                                            var v29 uint64 = (v22 + 931)
                                            v30 := GenS5{F0: v12}
                                            _ = v29
                                            _ = v30
                                        }
                                        var v31 bool = v1
                                        _ = v31
                                    }
                                case "hello":
                                    v1 = (128 == v3)
                                case "\n\t":
                                    if false {
                                        v20 = v20
                                        v14 = append(v14, (fi15 * fi21))
                                        fmt.Println(v3, v12)
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                        var v32 bool = ((fi15 & 42) != fi15)
                                        _ = v32
                                    }
                                }
                                _ = v23
                            } else {
                                if (v12 >= v12) {
                                    v12 = v12
                                }
                            }
                            _ = len(v14)
                            if (!(0.0 > (0.0 * float64((1.5 / -2.25))))) {
                                v14 = append(v14, 32768)
                                switch "\n\t" {
                                case "hello":
                                    v33 := GenS6{F0: uint(v3)}
                                    _ = v33
                                }
                                var v34 float64 = (1e10 * 1.5)
                                _ = v34
                            } else {
                                v12 = (v12 | v12)
                                v35 := GenS5{F0: int16(fi21)}
                                _ = v35
                            }
                            _ = v22
                        }
                        v20 -= v20
                        _ = fi21
                    }
                case "a":
                    v3 = 127
                }
                _ = v20
            }
        } else {
            fmt.Println(fi15, v2)
            for fi36 := 0; fi36 < 8; fi36++ {
                v2 -= (fi15 >> ((65535) & 63))
                _ = fi36
            }
        }
        _ = fi15
    }
    fmt.Println(v1)
    v1 = (v1 && (v1 || true))
    v37 := GenS7{F0: 127, F1: -9223372036854775808, F2: 41}
    fmt.Println(v1)
    for fi38 := 0; fi38 < 10; fi38++ {
        fmt.Println(v2, fi38)
        var v39 float32 = float32(1.5)
        v12 = v12
        _ = fi38
        _ = v39
    }
    fmt.Println(v12, v12, v1)
    var v40 bool = (949 < 2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v40
}
