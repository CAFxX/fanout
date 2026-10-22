package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint16, p1 int, p2 bool) bool {
    p0 = 652
    v6 := []int{p1, (-(p1)), (p1 >> ((255) & 63)), (p1 << ((221) & 63))}
    var v7 int = (^(p1))
    _ = len(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (!(2147483647 < 65535))
}
func genFn1() uint64 {
    var v8 int32 = 674
    fmt.Println(v8)
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint = 65535
    v2 := []uint32{514}
    var v3 string = ""
    var v4 uint8 = 127
    var v5 int64 = 0
    var v9 uint8 = (v4 + v4)
    switch ((("\n\t" + v3) + ((v3 + (((v3 + v3) + "") + "")) + v3)) + v3) {
    case "\n\t":
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                if false {
                    v12 := GenS1{F0: 0, F1: 4294967295}
                    _ = v12
                }
                if len(v2) > 0 {
                    v2[0] = 0
                }
                _ = fi11
            }
            v13 := GenS2{F0: v1, F1: 100, F2: -1}
            _ = fi10
            _ = v13
        }
    }
    v4 += v9
    v14 := GenS1{F0: 70, F1: 2147483648}
    _ = len(v2)
    var v15 float64 = 0.1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        v16 := GenS2{F0: v1, F1: -128, F2: -128}
        v9 = 2
        _ = v16
    } else {
        v1 += (v1 + 1)
    }
    if ((!false) || true) {
        if false {
            var v17 uint32 = 65535
            _ = v17
        } else {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 float32 = (1.5 / 1.5)
                _ = fi18
                _ = v19
            }
            var v20 string = v3
            _ = v20
        }
        _ = len(v2)
        v4 += (v9 % ((0) | 1))
    }
    if true {
        v15 -= v15
    } else {
        var v21 float32 = 0.1
        if false {
            switch 395 {
            case -320:
                var v22 float32 = v21
                _ = v22
            case 4294967295:
                var v23 int16 = 42
                _ = v23
            case 127:
                var v24 float32 = v21
                _ = v24
            default:
                if ((0.1 - (v21 + (0.1 + v21))) <= v21) {
                    var v25 string = v3
                    if false {
                        for fi26 := 0; fi26 < 8; fi26++ {
                            var v27 uint32 = 2
                            fmt.Println(v27, v15, v4)
                            v27 = (v27 / ((v27) | 1))
                            _ = fi26
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 2; fi28++ {
                            if true {
                                v29 := GenS1{F0: 1, F1: 323}
                                if (-2147483648 == 2147483647) {
                                    v21 = 0.1
                                    v15 = 0.0
                                } else {
                                    var v30 int16 = int16(fi28)
                                    fmt.Println(v25, fi28)
                                    if len(v2) > 0 {
                                        v2[0] = 4294967295
                                    }
                                    _ = v30
                                }
                                v4 = (v4 % ((v4) | 1))
                                _ = v29
                            }
                            _ = fi28
                        }
                    }
                    v31 := GenS1{F0: 127, F1: 2147483648}
                    _ = v25
                    _ = v31
                }
            }
        } else {
            switch -2147483648 {
            case -9223372036854775808:
                var v32 uint = v1
                _ = v32
            default:
                v33 := GenS3{F0: v4, F1: 1000, F2: 1}
                _ = v33
            }
            if (true || false) {
                var v34 string = "hello"
                if (false || true) {
                    v2 = append(v2, 309)
                    v35 := GenS3{F0: (v9 << (((v1 * (v1 + v1))) & 63)), F1: 0, F2: 923}
                    _ = v35
                } else {
                    if len(v2) > 0 {
                        v2[0] = uint32(v9)
                    }
                    v36 := GenS2{F0: 255, F1: 100, F2: -1}
                    v37 := GenS3{F0: v9, F1: -1, F2: genFn1()}
                    _ = v36
                    _ = v37
                }
                _ = v34
            } else {
                v15 = -2.25
            }
            v38 := GenS2{F0: (v1 % (((v1 ^ v1)) | 1)), F1: 0, F2: 127}
            _ = v38
        }
        v9 += v9
        _ = v21
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        switch v3 {
        case "\n\t":
            var v40 uint8 = (^(v9))
            _ = v40
        case "":
            if (false || (3.14159 != 0.1)) {
                var v41 int = (fi39 % (((fi39 % (((fi39 * (fi39 - 127))) | 1))) | 1))
                var v42 uint8 = (^(v4))
                _ = v41
                _ = v42
            } else {
                v43 := GenS2{F0: v1, F1: 127, F2: 94}
                _ = v43
            }
        default:
            v2 = append(v2, 65535)
        }
        if (-32768 != 778) {
            var v44 int8 = 100
            v15 -= (v15 + v15)
            _ = v44
        } else {
            for fi45 := 0; fi45 < 17; fi45++ {
                var v46 uint32 = 2147483648
                var v47 float32 = 1e10
                fmt.Println(v1)
                _ = fi45
                _ = v46
                _ = v47
            }
            if true {
                var v48 string = v3
                v49 := GenS1{F0: -43, F1: 15}
                v50 := GenS4{F0: v4, F1: -1}
                _ = v48
                _ = v49
                _ = v50
            } else {
                v3 = ("\n\t" + "\n\t")
            }
        }
        _ = fi39
    }
    v51 := GenS2{F0: v1, F1: -104, F2: 1}
    if (!true) {
        var v52 uint16 = 255
        _ = v52
    }
    fmt.Println(v4)
    fmt.Println(v1, v15)
    var v53 uint = (v1 & v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v53
    _ = v9
}
