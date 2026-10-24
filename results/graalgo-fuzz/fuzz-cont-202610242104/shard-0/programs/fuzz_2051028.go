package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int8; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint; F1 int32; F2 uint16 }
func genFn0() int8 {
    var v6 int8 = 100
    for fi7 := 0; fi7 < 12; fi7++ {
        v6 -= 127
        _ = fi7
    }
    _ = v6
    return 100
}
func genFn1(p0 uint16, p1 uint) bool {
    p0 += uint16(p1)
    switch ("" + "a") {
    case "":
        p0 += (p0 % ((p0) | 1))
    case "hello":
        switch 127 {
        case -392:
            var v8 int8 = 127
            _ = v8
        }
    case "\n\t":
        p1 += 1
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := []int{-246, 0, 426, -9223372036854775808}
    var v2 uint32 = 4294967295
    var v3 int16 = 3
    var v4 int = 32767
    var v5 float32 = -2.25
    for fi9 := 0; fi9 < 5; fi9++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v1) > 0 {
            v1[0] = v4
        }
        v4 = (fi9 % ((fi9) | 1))
        _ = fi9
    }
    v10 := GenS1{F0: 5, F1: v2}
    v11 := GenS2{F0: 4294967295}
    if false {
        if len(v1) > 0 {
            v1[0] = v4
        }
        if true {
            var v12 bool = false
            switch "a" {
            case "":
                if (-326 != -9223372036854775808) {
                    if (2 >= (^(v2))) {
                        if v12 {
                            var v13 int8 = 77
                            fmt.Println(v3, v5, v2)
                            fmt.Println(v5, v4)
                            _ = v13
                        }
                        v14 := GenS3{F0: 255, F1: -28, F2: genFn0()}
                        _ = len(v1)
                        _ = v14
                    }
                } else {
                    for fi15 := 0; fi15 < 2; fi15++ {
                        var v16 int = (-(v4))
                        _ = len(v1)
                        _ = fi15
                        _ = v16
                    }
                    v2 = 4294967295
                }
            case "hello":
                switch v4 {
                case 2147483647:
                    v2 = (v2 | (v2 + (^(v2))))
                case -579:
                    v3 -= v3
                case -942:
                    if v12 {
                        var v17 int8 = 127
                        _ = v17
                    }
                default:
                    v18 := GenS3{F0: 949, F1: -94, F2: int8(v2)}
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 16; fi19++ {
                    fmt.Println(v2)
                    _ = fi19
                }
            }
            _ = v12
        } else {
            v4 = v4
            switch (v4 ^ (v4 / ((v4) | 1))) {
            case 32767:
                var v20 uint32 = v2
                _ = v20
            case -1:
                var v21 uint32 = v2
                _ = v21
            default:
                for fi22 := 0; fi22 < 14; fi22++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi23 := 0; fi23 < 17; fi23++ {
                        var v24 int8 = 127
                        _ = fi23
                        _ = v24
                    }
                    v3 = v3
                    _ = fi22
                }
            }
            v25 := GenS4{F0: 65535, F1: 0, F2: 42}
            _ = v25
        }
    }
    if false {
        fmt.Println(v4)
        switch -935 {
        case 364:
            v1 = append(v1, v4)
        case 0:
            if true {
                fmt.Println(v5)
                var v26 int32 = -872
                _ = v26
            } else {
                if true {
                    if ((v2 > v2) || ((-(v3)) > 752)) {
                        if (!("hello" > (((("foo bar" + ("hello" + ("hello" + "\n\t"))) + "hello") + "foo bar") + ("" + "foo bar")))) {
                            for fi27 := 0; fi27 < 2; fi27++ {
                                v28 := GenS2{F0: (^(v2))}
                                var v29 uint16 = 985
                                for fi30 := 0; fi30 < 13; fi30++ {
                                    v4 += int(v2)
                                    fmt.Println(v4, v29)
                                    if len(v1) > 0 {
                                        v1[0] = 9223372036854775807
                                    }
                                    v31 := GenS2{F0: 664}
                                    _ = fi30
                                    _ = v31
                                }
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            fmt.Println(v5, v4)
                            v4 += v4
                        } else {
                            var v32 uint64 = 312
                            _ = v32
                        }
                        v4 = (-(v4))
                    } else {
                        v1 = append(v1, v4)
                        v33 := GenS3{F0: 0, F1: 0, F2: 42}
                        v34 := GenS5{F0: 205, F1: -1, F2: 0}
                        _ = v33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 15; fi35++ {
                        var v36 int32 = 976
                        v37 := GenS5{F0: 4294967295, F1: 804, F2: uint16(v2)}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                }
                switch (v4 % ((128) | 1)) {
                case 32767:
                    var v38 bool = false
                    _ = v38
                case 484:
                    var v39 int8 = 1
                    _ = v39
                }
                var v40 string = "a"
                _ = v40
            }
        case 458:
            v2 -= v2
        }
    }
    if len(v1) > 0 {
        v1[0] = v4
    }
    v5 = (v5 / 1.0)
    for fi41 := 0; fi41 < 17; fi41++ {
        switch "foo bar" {
        case "\n\t":
            if false {
                if len(v1) > 0 {
                    v1[0] = (fi41 / ((fi41) | 1))
                }
            } else {
                var v42 uint32 = 430
                _ = v42
            }
        case "":
            switch 956 {
            case 1:
                v4 = (v4 << ((4294967295) & 63))
            default:
                v43 := GenS4{F0: 32767, F1: 294, F2: 127}
                _ = v43
            }
        default:
            fmt.Println(v4, v4)
        }
        v1 = append(v1, fi41)
        _ = fi41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
