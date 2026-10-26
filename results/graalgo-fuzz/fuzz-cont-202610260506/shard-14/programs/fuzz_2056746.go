package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int; F1 uint32; F2 string }
func genFn0(p0 int64, p1 int8, p2 int16) uint8 {
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            fmt.Println(p0)
            p1 += (p1 << ((949) & 63))
            fmt.Println(fi6, fi7)
            _ = fi7
        }
        _ = fi6
    }
    if (2 == 908) {
        v8 := GenS1{F0: p1}
        v9 := []int{32767, -373, -363, int(p2)}
        _ = v8
        _ = v9
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 bool = (!(!true))
    var v2 string = ("" + "foo bar")
    v3 := []uint32{}
    v4 := []uint32{739, 1, 42}
    v5 := make(map[uint32]string)
    v1 = v1
    for fi10 := 0; fi10 < 12; fi10++ {
        if (!false) {
            v1 = false
        }
        _ = fi10
    }
    v2 = v2
    fmt.Println(v2, v1)
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS2{F0: 1, F1: (-(fi11))}
        if (32767 != (fi11 & fi11)) {
            v13 := GenS3{F0: -997, F1: 1, F2: 127}
            var v14 uint64 = 18446744073709551615
            _ = v13
            _ = v14
        } else {
            var v15 float64 = 1.5
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            if false {
                var v17 int32 = 2147483647
                if v1 {
                    if (18446744073709551615 >= 372) {
                        if v1 {
                            var v18 uint8 = 42
                            var v19 bool = true
                            _ = v18
                            _ = v19
                        } else {
                            v5[1] = v2
                        }
                        var v20 float64 = (((1e10 + (-2.25 / 1e10)) * 0.1) - 1.5)
                        fmt.Println(v20, v17, v2)
                        _ = v20
                    } else {
                        if v1 {
                            v2 = (v2 + "foo bar")
                            var v21 int = (fi11 % (((fi16 >> ((0) & 63))) | 1))
                            if (32767 >= 0) {
                                v21 = (v21 | fi16)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v2 = v2
                            }
                            v22 := GenS3{F0: int16(v21), F1: 4294967295, F2: 42}
                            fmt.Println(v1, fi16, v1)
                            _ = v21
                            _ = v22
                        } else {
                            if v1 {
                                var v23 int16 = -940
                                var v24 int = fi11
                                v2 = "a"
                                _ = v23
                                _ = v24
                            } else {
                                v2 = v2
                                switch "hello" {
                                case "hello":
                                    v2 = v2
                                case "foo bar":
                                    var v25 uint32 = 346
                                    _ = v25
                                default:
                                    var v26 float64 = 1.5
                                    _ = v26
                                }
                            }
                        }
                        _ = len(v5)
                        fmt.Println(v17)
                    }
                } else {
                    if v1 {
                        fmt.Println(v1, fi16, v2)
                        _ = len(v5)
                        v27 := GenS4{F0: fi11}
                        v2 = ("a" + v2)
                        var v28 int64 = 0
                        _ = v27
                        _ = v28
                    } else {
                        v17 -= v17
                        v2 = v2
                    }
                    if (v1 && v1) {
                        var v29 int16 = 822
                        _ = v29
                    } else {
                        fmt.Println(fi11)
                        var v30 string = ("hello" + "a")
                        var v31 uint64 = 0
                        _ = v30
                        _ = v31
                    }
                    var v32 float32 = float32(((1.5 / 1.5) * (3.14159 - (0.0 - ((-2.25 - (float64(float32((-2.25 + float64(float32((3.14159 * (3.14159 * 1e10))))))) + -2.25)) + (((float64(float32((-2.25 / (((float64(((-2.25 * (0.0 + (1.5 - 1e10))) * ((float32((3.14159 / 1.0)) * float32(-2.25)) / 3.14159))) + 1.5) * ((0.0 + 1.5) * 0.1)) * ((3.14159 + (1.5 / 3.14159)) * (1.5 + float64(((1.5 / float32((-2.25 + 3.14159))) - (3.14159 + 1.5))))))))) * 0.1) / (3.14159 + float64(-2.25))) / 0.1))))))
                    _ = v32
                }
                var v33 float32 = 1e10
                _ = v17
                _ = v33
            } else {
                v34 := GenS1{F0: -128}
                var v35 uint8 = genFn0(int64(fi11), 100, -407)
                var v36 uint32 = 42
                if false {
                    fmt.Println(fi16, v1, v2)
                    fmt.Println(v2, v2)
                    fmt.Println(fi11, v35)
                    if v1 {
                        var v37 float32 = -2.25
                        _ = v37
                    }
                }
                _ = v34
                _ = v35
                _ = v36
            }
            v38 := GenS5{F0: -1, F1: 31}
            var v39 int8 = 100
            if (65535 > 777) {
                v40 := v5[2147483648]
                if false {
                    fmt.Println(v40)
                    delete(v5, 43)
                } else {
                    var v41 int32 = 0
                    v42 := GenS3{F0: 32767, F1: 2, F2: 127}
                    _ = v41
                    _ = v42
                }
                if v1 {
                    fmt.Println(fi16, v1, fi11)
                    var v43 int32 = 2147483647
                    v39 = (v39 >> ((255) & 63))
                    _ = v43
                }
                _ = v40
            }
            v39 -= (v39 % ((-94) | 1))
            _ = fi16
            _ = v38
            _ = v39
        }
        _ = fi11
        _ = v12
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        if (64 < 2147483648) {
            v45 := GenS4{F0: 0}
            _ = v45
        } else {
            switch (fi44 * (fi44 | 92)) {
            case 2147483647:
                var v46 int16 = -1
                _ = v46
            default:
                if (2147483648 != 4294967295) {
                    v47 := GenS5{F0: 0, F1: 18446744073709551615}
                    var v48 int16 = 879
                    v49 := GenS6{F0: (fi44 / ((fi44) | 1)), F1: 2147483648, F2: (v2 + v2)}
                    _ = v47
                    _ = v48
                    _ = v49
                }
            }
        }
        _ = fi44
    }
    fmt.Println(v1, v1)
    if false {
        var v50 float64 = 1.5
        v50 -= 0.0
        _ = v50
    } else {
        for fi51 := 0; fi51 < 7; fi51++ {
            _ = len(v5)
            fmt.Println(v1, v1, fi51)
            v1 = (v1 && false)
            _ = fi51
        }
        for fi52 := 0; fi52 < 13; fi52++ {
            var v53 int32 = -1
            _ = fi52
            _ = v53
        }
        var v54 uint32 = 2
        _ = v54
    }
    fmt.Println(v1, v1)
    var v55 bool = v1
    for fi56 := 0; fi56 < 5; fi56++ {
        v2 = ((("\n\t" + "hello") + "\n\t") + v2)
        v1 = (v2 != v2)
        _ = fi56
    }
    if v55 {
        v55 = false
    }
    v2 = "a"
    var v57 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v57
}
