package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int; F2 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint16) bool {
    var v7 float64 = 1.5
    v7 = (v7 / v7)
    switch ("" + "hello") {
    case "hello":
        p0 = p0
    default:
        var v8 uint = 4294967295
        _ = v8
    }
    if false {
        v9 := GenS1{F0: 861}
        v10 := GenS1{F0: 154}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    return true
}
func genFn1(p0 uint16) uint32 {
    v11 := []int{871, 32767}
    var v12 string = ""
    p0 = 32767
    _ = p0
    _ = v11
    _ = v12
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 1000
    v3 := []int{32768, 42, -518}
    v4 := []uint32{}
    var v5 float64 = (-2.25 - 1.5)
    var v6 uint8 = 2
    var v13 uint16 = 255
    if ((1 <= 1) || true) {
        v14 := GenS2{F0: 1, F1: 42, F2: 823}
        fmt.Println(v13, v6, v13)
        _ = v14
    } else {
        var v15 int = 42
        v13 = (v13 | 2)
        _ = v15
    }
    v4 = append(v4, uint32(v2))
    var v16 uint64 = (v2 + (v2 << ((42) & 63)))
    var v17 uint16 = v13
    fmt.Println(v5, v13, v16)
    v2 -= 570
    var v18 int32 = 0
    var v19 uint64 = uint64(v17)
    v20 := GenS3{F0: 2, F1: -250, F2: 171}
    if (!((591 >= 42) && false)) {
        if (v6 == (v6 * v6)) {
            if (!false) {
                var v21 uint64 = v19
                for fi22 := 0; fi22 < 2; fi22++ {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if true {
                            fmt.Println(v21, v18, v2)
                        }
                        var v24 int32 = v18
                        v25 := GenS4{F0: 1}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    var v26 uint = 1
                    _ = fi22
                    _ = v26
                }
                if (65535 == 42) {
                    v2 = 18446744073709551615
                    if false {
                        if false {
                            var v27 bool = (uint(v13) >= 4294967295)
                            _ = v27
                        } else {
                            var v28 int64 = int64(v18)
                            if true {
                                var v29 uint = 42
                                for fi30 := 0; fi30 < 12; fi30++ {
                                    if false {
                                        if len(v4) > 0 {
                                            v4[0] = 4294967295
                                        }
                                    }
                                    if true {
                                        if true {
                                            if (true && (v13 != uint16(v2))) {
                                                delete(v1, 837)
                                            } else {
                                                if true {
                                                    v31 := GenS4{F0: 255}
                                                    _ = v31
                                                } else {
                                                    v13 -= v17
                                                    for fi32 := 0; fi32 < 20; fi32++ {
                                                        var v33 uint64 = v19
                                                        v6 += (v6 | 231)
                                                        if (true || ((0.0 > 3.14159) || false)) {
                                                            v19 = v16
                                                            fmt.Println(v19)
                                                        }
                                                        _ = fi32
                                                        _ = v33
                                                    }
                                                }
                                                v6 += v6
                                                if len(v3) > 0 {
                                                    v3[0] = (fi30 << (((v29 * (v29 ^ v29))) & 63))
                                                }
                                            }
                                        }
                                        var v34 float64 = (v5 + v5)
                                        var v35 float32 = 0.1
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = fi30
                                }
                                _ = v29
                            } else {
                                v17 += v13
                            }
                            var v36 int16 = 1
                            _ = v28
                            _ = v36
                        }
                    }
                } else {
                    var v37 uint = 976
                    delete(v1, 0)
                    _ = v37
                }
                _ = v21
            } else {
                var v38 uint16 = 287
                var v39 uint64 = (v2 & 1)
                _ = v38
                _ = v39
            }
            if ((v6 << ((667) & 63)) != v6) {
                fmt.Println(v19, v13)
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v17, v18)
            var v40 float64 = v5
            _ = v40
        }
        switch 128 {
        case -337:
            var v41 uint32 = 2147483647
            _ = v41
        default:
            fmt.Println(v17, v13)
        }
        var v42 int = 127
        _ = v42
    } else {
        var v43 int16 = -1
        var v44 int = -791
        _ = v43
        _ = v44
    }
    var v45 uint = 284
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
