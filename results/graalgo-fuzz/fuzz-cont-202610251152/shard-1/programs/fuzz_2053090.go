package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS2 struct { F0 int16; F1 int32; F2 uint8 }
type GenS3 struct { F0 string; F1 int32 }
func genFn0() bool {
    v6 := []int{-9223372036854775808, -9223372036854775808}
    var v7 bool = ((!false) || false)
    v8 := GenS1{F0: 2147483647, F1: 0, F2: -128}
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 string = (((("a" + "foo bar") + ("" + ("hello" + ("" + ("a" + "hello"))))) + "foo bar") + "foo bar")
        var v11 float64 = 1e10
        _ = fi9
        _ = v10
        _ = v11
    }
    if (!false) {
        for fi12 := 0; fi12 < 11; fi12++ {
            switch (fi12 | 42) {
            case 2147483647:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case 127:
                v13 := GenS2{F0: -1, F1: -1, F2: 128}
                _ = v13
            }
            _ = fi12
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v14 := GenS1{F0: 259, F1: 524, F2: 100}
        _ = v14
    }
    _ = v6
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 int16) int8 {
    var v15 int8 = 1
    fmt.Println(p0, p0)
    _ = p0
    _ = v15
    return 42
}

func main() {
    var v1 uint16 = 65535
    var v2 int8 = -1
    v3 := make(map[uint32]uint32)
    v4 := []string{}
    v5 := []uint32{uint32(v1), 262, 4294967295, 2147483647}
    for fi16 := 0; fi16 < 14; fi16++ {
        v2 = (v2 & v2)
        _ = len(v3)
        if ((1 >= genFn1(32767)) || true) {
            v2 += (v2 / ((v2) | 1))
            v17 := GenS2{F0: 1000, F1: -789, F2: 128}
            _ = v17
        } else {
            var v18 int = 9223372036854775807
            _ = len(v3)
            _ = v18
        }
        _ = fi16
    }
    fmt.Println(v2, v1)
    if (0 < 0) {
        v19 := GenS1{F0: 2, F1: (v1 * (v1 - v1)), F2: 0}
        _ = v19
    }
    if false {
        if (((0.0 > 0.0) && (!false)) && false) {
            v20 := GenS2{F0: 32767, F1: 0, F2: 214}
            var v21 uint64 = 18446744073709551615
            v21 = 4294967295
            _ = v20
            _ = v21
        } else {
            if (42 >= 2) {
                v1 = v1
                if false {
                    v3[0] = 522
                    v2 += (v2 ^ -125)
                    var v22 uint16 = 255
                    _ = v22
                }
                v1 -= (v1 << ((255) & 63))
            } else {
                var v23 uint = 0
                if false {
                    v1 = 0
                }
                _ = v23
            }
            if false {
                var v24 uint32 = 2147483648
                var v25 int64 = 42
                v2 -= v2
                _ = v24
                _ = v25
            } else {
                v26 := GenS2{F0: -274, F1: 1, F2: 63}
                v27 := GenS1{F0: 100, F1: v1, F2: 1}
                _ = v26
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 1; fi28++ {
            fmt.Println(v2, fi28)
            var v29 int32 = 1
            _ = fi28
            _ = v29
        }
        v2 = v2
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        if true {
            var v31 uint16 = (v1 << ((611) & 63))
            _ = v31
        }
        _ = fi30
    }
    var v32 float32 = 0.1
    fmt.Println(v1, v1, v2)
    var v33 uint = 921
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v3)
    fmt.Println(v32, v32)
    if (!(true && (42 == 1))) {
        switch 167 {
        case 2147483647:
            fmt.Println(v1, v2)
        case -828:
            if ((!(v33 != (^(v33)))) || true) {
                var v34 int8 = v2
                for fi35 := 0; fi35 < 1; fi35++ {
                    if len(v4) > 0 {
                        v4[0] = ("a" + "a")
                    }
                    v32 = v32
                    var v36 bool = (128 <= 168)
                    _ = fi35
                    _ = v36
                }
                var v37 uint8 = 128
                _ = v34
                _ = v37
            }
        default:
            if false {
                v38 := GenS2{F0: 42, F1: -2147483648, F2: 2}
                v33 = v33
                var v39 int8 = 127
                _ = v38
                _ = v39
            } else {
                v40 := GenS1{F0: 4294967295, F1: v1, F2: genFn1(32767)}
                v41 := GenS1{F0: 590, F1: 255, F2: (v2 ^ -45)}
                var v42 uint = v33
                _ = v40
                _ = v41
                _ = v42
            }
        }
        v43 := GenS3{F0: "\n\t", F1: -1}
        _ = v43
    } else {
        switch "a" {
        case "foo bar":
            var v44 uint16 = 32767
            _ = v44
        }
        _ = len(v4)
    }
    delete(v3, 2)
    if (uint(v1) > v33) {
        v1 = (v1 & (v1 % (((v1 & v1)) | 1)))
        for fi45 := 0; fi45 < 10; fi45++ {
            v2 = v2
            fmt.Println(fi45, v32)
            _ = fi45
        }
        if false {
            if (v2 >= (^(v2))) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if true {
                    v33 -= (v33 | (v33 * 0))
                    v46 := v3[65535]
                    _ = v46
                } else {
                    fmt.Println(v1)
                    for fi47 := 0; fi47 < 13; fi47++ {
                        var v48 int16 = 0
                        if true {
                            v49 := GenS3{F0: "a", F1: -789}
                            delete(v3, 65535)
                            _ = v49
                        } else {
                            fmt.Println(v48)
                            fmt.Println(v1, v1, v2)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        v1 = v1
                        _ = fi47
                        _ = v48
                    }
                    var v50 uint8 = 255
                    _ = v50
                }
                switch 4294967295 {
                case 638:
                    var v51 int8 = 42
                    _ = v51
                case 127:
                    v2 -= v2
                case -2147483648:
                    v33 = v33
                default:
                    var v52 int16 = -32768
                    _ = v52
                }
            }
        }
    }
    v32 = ((v32 * 1.5) / 1.0)
    v32 -= 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
