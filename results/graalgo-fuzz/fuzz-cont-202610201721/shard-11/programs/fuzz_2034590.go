package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 string; F2 int16 }
func genFn0(p0 uint64, p1 uint16, p2 int32) int8 {
    var v4 uint16 = p1
    v5 := GenS1{F0: false, F1: 127, F2: 42}
    var v6 int32 = (p2 + p2)
    var v7 uint64 = (p0 | 18446744073709551615)
    var v8 float64 = (-2.25 - 1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 bool, p1 int, p2 bool) bool {
    p1 = p1
    var v9 int64 = 473
    fmt.Println(p0)
    v9 = v9
    fmt.Println(v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p0 && (32767 <= (^(p1))))
}

func main() {
    var v1 bool = (!false)
    var v2 uint64 = 4294967295
    var v3 uint16 = 65535
    v10 := GenS1{F0: v1, F1: 195, F2: 47}
    if (18446744073709551615 > 2) {
        v11 := []int{470, -969}
        if false {
            var v12 int64 = 9223372036854775807
            for fi13 := 0; fi13 < 3; fi13++ {
                if true {
                    var v14 uint32 = 2147483647
                    _ = len(v11)
                    fmt.Println(v1)
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        var v16 bool = v1
                        _ = fi15
                        _ = v16
                    }
                    if (42 != 2) {
                        for fi17 := 0; fi17 < 13; fi17++ {
                            v3 += 963
                            if (v1 || (v1 || (!v1))) {
                                switch (fi13 << ((767) & 63)) {
                                case -479:
                                    v12 += 9223372036854775807
                                case 0:
                                    if (820 < 42) {
                                        var v18 int = fi13
                                        v12 = v12
                                        _ = v18
                                    } else {
                                        if v1 {
                                            fmt.Println(v1)
                                            if len(v11) >= 2 {
                                                v11 = v11[1:len(v11)]
                                            }
                                            if true {
                                                var v19 int = fi17
                                                v20 := GenS2{F0: (("\n\t" + ("" + (("\n\t" + "") + "foo bar"))) + "a")}
                                                _ = v19
                                                _ = v20
                                            } else {
                                                switch "a" {
                                                case "\n\t":
                                                    v1 = (!(2 <= v2))
                                                default:
                                                    var v21 int64 = v12
                                                    _ = v21
                                                }
                                                var v22 int32 = 471
                                                _ = len(v11)
                                                var v23 float32 = 3.14159
                                                var v24 int8 = 42
                                                _ = len(v11)
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            }
                                        }
                                        v3 -= (v3 + (v3 | (v3 << ((333) & 63))))
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                        if len(v11) > 0 {
                                            v11[0] = 4294967295
                                        }
                                        v2 = v2
                                    }
                                case 314:
                                    if (!false) {
                                        v11 = append(v11, fi17)
                                        var v25 float32 = 3.14159
                                        _ = v25
                                    }
                                }
                                v12 -= (v12 + -333)
                            } else {
                                v2 += v2
                            }
                            _ = fi17
                        }
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v27 := GenS3{F0: 42, F1: "a", F2: -32768}
                            _ = fi26
                            _ = v27
                        }
                    }
                }
                _ = fi13
            }
            _ = v12
        }
        for fi28 := 0; fi28 < 6; fi28++ {
            v3 -= v3
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi28
        }
        _ = v11
    } else {
        v2 += (^(v2))
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        for fi30 := 0; fi30 < 11; fi30++ {
            switch ("a" + "hello") {
            case "foo bar":
                v1 = ((("foo bar" + "\n\t") + "") <= "")
            default:
                var v31 uint32 = 2147483648
                _ = v31
            }
            v3 -= 2
            var v32 int8 = -128
            _ = fi30
            _ = v32
        }
        var v33 bool = v1
        _ = fi29
        _ = v33
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        if false {
            v35 := []int{fi34, (fi34 * 128)}
            _ = v35
        }
        _ = fi34
    }
    v3 += v3
    var v36 uint32 = 1
    v37 := GenS2{F0: ("" + ("hello" + "hello"))}
    fmt.Println(v1, v2, v1)
    if v1 {
        v38 := GenS1{F0: (-2.25 >= (0.1 - -2.25)), F1: 0, F2: 91}
        _ = v38
    }
    v2 = (v2 / (((v2 | v2)) | 1))
    v39 := []int{-304}
    v3 -= v3
    for fi40 := 0; fi40 < 20; fi40++ {
        fmt.Println(fi40)
        v36 = v36
        var v41 bool = (!v1)
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
}
