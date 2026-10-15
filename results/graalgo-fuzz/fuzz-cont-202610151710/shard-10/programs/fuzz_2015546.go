package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int; F1 string }
func genFn0(p0 uint32, p1 bool, p2 bool) uint16 {
    v7 := []int{2147483647}
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(p2)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1() int64 {
    for fi9 := 0; fi9 < 13; fi9++ {
        if ((!(1.5 >= ((float32(float64(float32((((float64(0.1) + float64((-2.25 - (float32((1e10 - 0.1)) + (float32(3.14159) + (1e10 - float32((1e10 + 0.0)))))))) - float64(-2.25)) - 3.14159)))) * 1.5) / -2.25))) && (9223372036854775807 > -988)) {
            var v10 int32 = 2147483647
            fmt.Println(v10, v10)
            fmt.Println(fi9, fi9)
            _ = v10
        } else {
            v11 := []int{-963, fi9, -750}
            var v12 float32 = 1e10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    v13 := GenS1{F0: -32768}
    _ = v13
    return 1
}

func main() {
    v1 := []uint32{625}
    v2 := []string{}
    var v3 int32 = -750
    var v4 bool = false
    var v5 int16 = 1000
    var v6 int64 = 1
    var v14 float64 = (float64(((3.14159 * 0.0) / 1e10)) * 3.14159)
    v15 := GenS1{F0: (v5 + (v5 - -32768))}
    for fi16 := 0; fi16 < 16; fi16++ {
        if v4 {
            fmt.Println(v14, v14)
            var v17 uint16 = 65535
            for fi18 := 0; fi18 < 6; fi18++ {
                if v4 {
                    var v19 uint64 = 2
                    v17 += (^(v17))
                    var v20 float32 = (0.0 / ((((0.0 / 3.14159) - 0.0) / ((1.5 / 1.0) - 1e10)) - 1.5))
                    if v4 {
                        v21 := GenS2{F0: ((v4 || true) && (!v4))}
                        var v22 string = "\n\t"
                        _ = v21
                        _ = v22
                    }
                    if v4 {
                        var v23 float64 = ((float64((v20 - float32(((v14 - 1e10) * (v14 + v14))))) + v14) * v14)
                        var v24 bool = v4
                        v25 := GenS3{F0: 2147483648}
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = v19
                    _ = v20
                } else {
                    v17 = genFn0(uint32(v17), true, true)
                    if false {
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                        v26 := GenS4{F0: (fi18 * (fi18 | fi18))}
                        _ = v26
                    } else {
                        var v27 bool = (0.1 < ((v14 + (v14 - float64(3.14159))) * ((v14 * (v14 / float64((1e10 + 1e10)))) + (float64(((((0.0 / ((0.1 / 0.1) / 1e10)) / float32(((1e10 * 1.5) / -2.25))) - float32((((v14 + 0.0) + v14) - -2.25))) - (-2.25 * 1.5))) / v14))))
                        _ = v27
                    }
                    if ((fi18 | -9223372036854775808) <= fi16) {
                        _ = len(v1)
                        var v28 string = ""
                        if v4 {
                            if v4 {
                                var v29 int16 = int16(fi18)
                                v4 = v4
                                v17 = genFn0(489, v4, (v3 != v3))
                                v30 := GenS5{F0: v17, F1: 624}
                                _ = v29
                                _ = v30
                            }
                            if ((-1 <= v3) && ((v6 <= -264) && v4)) {
                                v31 := GenS1{F0: -1}
                                var v32 int16 = v5
                                _ = v31
                                _ = v32
                            } else {
                                switch ((v28 + "\n\t") + v28) {
                                case "":
                                    var v33 uint = 18446744073709551615
                                    _ = v33
                                case "a":
                                    if true {
                                        var v34 string = "\n\t"
                                        var v35 string = v28
                                        v35 = v35
                                        _ = v34
                                        _ = v35
                                    }
                                case "hello":
                                    var v36 float64 = 0.1
                                    _ = v36
                                }
                            }
                            v37 := GenS6{F0: fi18, F1: ("hello" + "\n\t")}
                            _ = v37
                        } else {
                            fmt.Println(v4, v6, v3)
                            fmt.Println(v6)
                        }
                        _ = v28
                    } else {
                        _ = len(v1)
                        switch "\n\t" {
                        case "":
                            v2 = append(v2, "foo bar")
                        case "a":
                            fmt.Println(fi16, v5, v5)
                        case "foo bar":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v6 = (v6 | v6)
                    }
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (2147483648 < 4294967295) {
                            if false {
                                switch (^(fi16)) {
                                case 127:
                                    var v38 int64 = (v6 << ((65535) & 63))
                                    _ = v38
                                case 463:
                                    fmt.Println(fi18)
                                }
                            }
                            var v39 int64 = (v6 | 9223372036854775807)
                            _ = v39
                        }
                    }
                }
                _ = fi18
            }
            _ = v17
        }
        _ = fi16
    }
    if (!((806 == 2) || v4)) {
        _ = len(v1)
        var v40 uint = 4294967295
        var v41 uint = (v40 / ((v40) | 1))
        _ = v40
        _ = v41
    } else {
        v3 = v3
        v6 -= v6
    }
    v14 = (0.0 + (float64(1.5) * v14))
    v1 = append(v1, 2147483648)
    var v42 int8 = 0
    var v43 float32 = (3.14159 + (-2.25 + 0.1))
    fmt.Println(v43, v42, v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
