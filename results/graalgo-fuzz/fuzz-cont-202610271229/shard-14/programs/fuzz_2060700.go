package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint16) int8 {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint16 = p0
        _ = fi6
        _ = v7
    }
    if (int32(p0) == -2147483648) {
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{-9223372036854775808, 32768, 1}
            _ = fi8
            _ = v9
        }
        p0 = (p0 / (((p0 + (p0 << ((0) & 63)))) | 1))
        switch "foo bar" {
        case "foo bar":
            p0 = 20
        default:
            p0 = p0
        }
    }
    v10 := []int{}
    if (!((true && true) || false)) {
        for fi11 := 0; fi11 < 6; fi11++ {
            p0 = (p0 / ((2) | 1))
            var v12 int64 = 9223372036854775807
            _ = fi11
            _ = v12
        }
        var v13 bool = false
        v14 := GenS1{F0: ((((65535 < 4294967295) || v13) && (1 < 171)) || (-32768 < -32768)), F1: 32767}
        _ = v13
        _ = v14
    } else {
        p0 -= 42
    }
    var v15 uint = 647
    _ = p0
    _ = v10
    _ = v15
    return 0
}

func main() {
    v1 := []uint32{819, 2147483647, 4294967295}
    var v2 int64 = -741
    var v3 int = 127
    v4 := make(map[uint32]string)
    var v5 int64 = v2
    for fi16 := 0; fi16 < 1; fi16++ {
        v5 = (v5 >> ((65535) & 63))
        _ = fi16
    }
    v3 = v3
    var v17 uint16 = 304
    if false {
        v3 = (-(v3))
        v18 := GenS1{F0: ((v5 & v2) >= (v5 >> ((954) & 63))), F1: 42}
        for fi19 := 0; fi19 < 1; fi19++ {
            fmt.Println(v2, v5, v17)
            v1 = append(v1, 2147483647)
            _ = fi19
        }
        _ = v18
    }
    if ("\n\t" > (("foo bar" + ("\n\t" + (((((("a" + "\n\t") + "foo bar") + "foo bar") + ((("" + ("a" + ("foo bar" + ("\n\t" + ("hello" + "\n\t"))))) + "a") + "")) + ("a" + ("hello" + "a"))) + "\n\t"))) + (("a" + ("hello" + "a")) + ("" + "hello")))) {
        var v20 uint = 255
        _ = v20
    } else {
        fmt.Println(v2, v3)
    }
    if false {
        if len(v1) > 0 {
            v1[0] = 0
        }
    } else {
        if true {
            var v21 uint16 = (^(v17))
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 int8 = 1
                for fi24 := 0; fi24 < 20; fi24++ {
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        if ((true && (true && true)) && false) {
                            if false {
                                var v25 float32 = ((1.5 * 3.14159) / (float32(3.14159) + 1.5))
                                if false {
                                    fmt.Println(v3, fi24, v23)
                                    var v26 uint16 = v21
                                    _ = v26
                                } else {
                                    var v27 uint16 = v21
                                    _ = v27
                                }
                                _ = v25
                            } else {
                                _ = len(v4)
                                fmt.Println(v23)
                                v2 = v5
                            }
                        }
                    }
                    if (1.5 != (3.14159 / 0.1)) {
                        var v28 string = "a"
                        _ = v28
                    } else {
                        v23 = (^(v23))
                    }
                    _ = fi24
                }
                v29 := GenS2{F0: -32768}
                _ = fi22
                _ = v23
                _ = v29
            }
            _ = v21
        }
    }
    fmt.Println(v17, v17)
    for fi30 := 0; fi30 < 8; fi30++ {
        v2 -= (^(v5))
        var v31 int16 = 42
        _ = fi30
        _ = v31
    }
    v32 := GenS3{F0: int16(v3), F1: (v17 / ((855) | 1)), F2: 42}
    switch "\n\t" {
    case "a":
        v33 := v4[667]
        _ = v33
    case "":
        v17 += v17
    default:
        for fi34 := 0; fi34 < 9; fi34++ {
            v3 += 127
            _ = fi34
        }
    }
    v35 := GenS4{F0: (!(1 != 42)), F1: -972}
    var v36 uint16 = v17
    delete(v4, 2)
    v37 := v4[884]
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if (4294967295 < 827) {
        for fi38 := 0; fi38 < 15; fi38++ {
            var v39 uint8 = 2
            _ = fi38
            _ = v39
        }
        var v40 uint16 = (v17 + 65535)
        if true {
            for fi41 := 0; fi41 < 11; fi41++ {
                fmt.Println(v40, v40, v17)
                _ = fi41
            }
            if ((v5 - -1) != v5) {
                fmt.Println(v3)
                v42 := GenS5{F0: 18446744073709551615}
                v40 = (v40 + v40)
                _ = v42
            }
            switch v3 {
            case 4294967295:
                if (!(true && true)) {
                    switch (-(v3)) {
                    case 127:
                        var v43 int32 = -2147483648
                        _ = v43
                    case 9223372036854775807:
                        v44 := GenS2{F0: 1}
                        _ = v44
                    default:
                        var v45 bool = (!false)
                        _ = v45
                    }
                } else {
                    var v46 bool = (128 != 1)
                    _ = v46
                }
            default:
                if true {
                    _ = len(v4)
                }
            }
        }
        _ = v40
    } else {
        delete(v4, 0)
    }
    delete(v4, 2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
