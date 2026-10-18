package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int16; F2 int8 }

func main() {
    var v1 float32 = 1.5
    v2 := []string{("hello" + ("a" + "a")), "\n\t", "a"}
    var v3 float64 = float64(v1)
    var v4 uint32 = 2147483647
    v1 -= ((3.14159 * ((-2.25 / v1) - v1)) + float32(float64((1.5 / 1.5))))
    v3 -= (3.14159 + (-2.25 / 1e10))
    fmt.Println(v4, v4, v3)
    fmt.Println(v3)
    var v5 float64 = (v3 * (-2.25 + ((v3 / v3) * (-2.25 * 1.5))))
    for fi6 := 0; fi6 < 3; fi6++ {
        v1 += -2.25
        v7 := GenS1{F0: 796}
        fmt.Println(v3)
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: 65535}
    fmt.Println(v1, v5)
    v5 = v3
    switch 4294967295 {
    case 861:
        if (128 > 2147483647) {
            if (1 <= 69) {
                for fi9 := 0; fi9 < 20; fi9++ {
                    v5 = 0.0
                    _ = fi9
                }
                v3 += v5
            } else {
                var v10 uint32 = v4
                var v11 uint16 = 255
                _ = v10
                _ = v11
            }
            v3 = v3
            if true {
                fmt.Println(v1, v3)
            }
        } else {
            v12 := GenS1{F0: 1}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := GenS2{F0: 0}
                var v15 string = ("a" + "hello")
                for fi16 := 0; fi16 < 11; fi16++ {
                    if (127 >= -128) {
                        v15 = "hello"
                    }
                    v4 += v4
                    if ((true && false) && false) {
                        var v17 int8 = -1
                        if (false || false) {
                            if (18446744073709551615 <= 49) {
                                v3 = v5
                                var v18 uint = 4294967295
                                _ = v18
                            } else {
                                v5 = v5
                            }
                            var v19 bool = ((!false) && false)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v19
                        }
                        v20 := GenS1{F0: 42}
                        if (!(float64(v1) <= (v5 * (-2.25 - float64(v1))))) {
                            var v21 int8 = (v17 * -124)
                            var v22 int8 = (v21 << ((42) & 63))
                            _ = v21
                            _ = v22
                        } else {
                            if false {
                                var v23 uint64 = 4294967295
                                v17 = (v17 & v17)
                                v1 -= 0.0
                                _ = v23
                            }
                        }
                        v5 = float64(-2.25)
                        _ = v17
                        _ = v20
                    } else {
                        v24 := GenS3{F0: 1, F1: -865, F2: 127}
                        if true {
                            var v25 float32 = 0.1
                            fmt.Println(v25)
                            _ = v25
                        }
                        fmt.Println(v1)
                        _ = v24
                    }
                    var v26 uint16 = 1
                    _ = fi16
                    _ = v26
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
    case 128:
        var v27 int32 = -622
        _ = v27
    case 127:
        var v28 uint8 = 127
        _ = v28
    default:
        v2 = append(v2, ("" + "\n\t"))
    }
    fmt.Println(v4)
    var v29 uint = 18446744073709551615
    if ((-2147483648 == 2147483647) && (718 <= -1)) {
        for fi30 := 0; fi30 < 7; fi30++ {
            fmt.Println(v4, fi30, v1)
            var v31 uint32 = v4
            for fi32 := 0; fi32 < 16; fi32++ {
                v5 = ((0.0 - (v5 + v3)) + v5)
                var v33 int64 = -704
                if ((false || true) && false) {
                    var v34 int64 = 616
                    _ = v34
                } else {
                    v3 -= (v3 / (v3 / float64(-2.25)))
                    var v35 float32 = v1
                    _ = v35
                }
                var v36 uint16 = 488
                v5 = v5
                _ = fi32
                _ = v33
                _ = v36
            }
            _ = fi30
            _ = v31
        }
        var v37 uint32 = v4
        if (!((false || ((v37 % ((v4) | 1)) > 65535)) && true)) {
            v5 += (((float64(v1) + v3) + (((v3 * 1e10) - 1e10) / (float64(((v1 + (3.14159 * (v1 * ((0.0 - 1e10) / v1)))) / v1)) / (1e10 - float64(v1))))) * 0.1)
        }
        _ = v37
    } else {
        for fi38 := 0; fi38 < 12; fi38++ {
            var v39 float32 = (1.5 / v1)
            if (-2147483648 >= -66) {
                if true {
                    v29 += 1
                }
            }
            v4 -= v4
            _ = fi38
            _ = v39
        }
        for fi40 := 0; fi40 < 18; fi40++ {
            var v41 uint8 = 0
            v29 = 0
            _ = fi40
            _ = v41
        }
        v2 = append(v2, ("a" + ("a" + "foo bar")))
    }
    var v42 int16 = int16(v29)
    fmt.Println(v5, v4, v1)
    var v43 string = ("" + ("" + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v8
}
