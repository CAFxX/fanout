package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0() uint16 {
    var v5 bool = false
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{(fi6 % ((fi6) | 1)), -202, (fi6 & -9223372036854775808), fi6}
        var v8 int64 = int64(fi6)
        fmt.Println(v5)
        _ = fi6
        _ = v7
        _ = v8
    }
    v5 = v5
    _ = v5
    return 32767
}
func genFn1(p0 uint) int {
    v9 := []int{127, 42, 0, -9223372036854775808}
    if (-32768 != 32767) {
        var v10 bool = false
        v10 = v10
        v10 = (v10 && v10)
        _ = v10
    } else {
        p0 += (^(p0))
        p0 -= p0
        switch (("foo bar" + "foo bar") + "") {
        case "hello":
            var v11 uint = (p0 >> ((p0) & 63))
            _ = v11
        case "\n\t":
            p0 = (^(p0))
        default:
            switch 32768 {
            case 2147483647:
                if (441 >= 0) {
                    v12 := GenS1{F0: 234}
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v14 := GenS1{F0: 255}
                        _ = fi13
                        _ = v14
                    }
                    var v15 uint32 = 673
                    _ = v12
                    _ = v15
                } else {
                    p0 = (p0 ^ p0)
                    if false {
                        for fi16 := 0; fi16 < 5; fi16++ {
                            switch "a" {
                            case "a":
                                if (691 > 1) {
                                    var v17 int64 = -980
                                    var v18 uint = 178
                                    v19 := GenS1{F0: 199}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                            case "":
                                var v20 int16 = 1000
                                _ = v20
                            }
                            p0 += 18446744073709551615
                            for fi21 := 0; fi21 < 2; fi21++ {
                                if len(v9) > 0 {
                                    v9[0] = (fi16 << (((^(p0))) & 63))
                                }
                                v22 := GenS1{F0: 1}
                                p0 = (p0 ^ 42)
                                p0 += (p0 * p0)
                                _ = fi21
                                _ = v22
                            }
                            _ = fi16
                        }
                        _ = len(v9)
                        var v23 int8 = 76
                        _ = v23
                    } else {
                        fmt.Println(p0)
                    }
                }
            case 42:
                for fi24 := 0; fi24 < 14; fi24++ {
                    if false {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    } else {
                        for fi25 := 0; fi25 < 19; fi25++ {
                            var v26 float32 = 0.0
                            p0 += (^(p0))
                            v27 := GenS1{F0: 245}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 13; fi28++ {
                            v29 := GenS1{F0: 18}
                            _ = fi28
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 6; fi30++ {
                            v31 := GenS1{F0: 0}
                            _ = fi30
                            _ = v31
                        }
                    }
                    fmt.Println(p0, p0)
                    _ = fi24
                }
            }
        }
    }
    p0 = 255
    _ = p0
    _ = v9
    return 32768
}

func main() {
    var v1 int8 = 100
    v2 := []uint32{2147483648, 1000, 2147483648}
    var v3 int16 = 1
    v4 := []uint32{713}
    for fi32 := 0; fi32 < 15; fi32++ {
        v33 := GenS2{F0: (v1 - v1)}
        _ = fi32
        _ = v33
    }
    v3 = (v3 / (((v3 + v3)) | 1))
    v3 += (v3 + 1)
    v3 = v3
    var v34 int8 = 0
    v3 = v3
    if false {
        v3 += v3
        v1 = v34
        var v35 bool = (float64(0.0) == float64(((float32((((1e10 + 3.14159) * 0.1) * (-2.25 + (-2.25 + ((3.14159 * 1e10) - (0.0 - 0.0)))))) / (1e10 * 3.14159)) / 0.1)))
        _ = v35
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        var v37 int = fi36
        for fi38 := 0; fi38 < 1; fi38++ {
            v39 := GenS3{F0: -2147483648, F1: (false || true)}
            var v40 uint16 = 702
            v3 += v3
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi36
        _ = v37
    }
    v1 = (v34 ^ (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
