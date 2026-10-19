package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint16; F1 string; F2 int32 }
type GenS5 struct { F0 uint16; F1 string; F2 uint32 }

func main() {
    v1 := []string{}
    var v2 int32 = 1
    var v3 uint16 = 65535
    var v4 int8 = int8(v3)
    v5 := []uint32{2, 4294967295, 2147483648, 4294967295}
    v2 += (v2 * v2)
    v5 = append(v5, 782)
    switch 714 {
    case -1:
        fmt.Println(v3, v4, v2)
    case 4294967295:
        var v6 int64 = 42
        _ = v6
    default:
        v7 := GenS1{F0: 1, F1: -396}
        _ = v7
    }
    if (true && true) {
        v2 = (v2 >> ((42) & 63))
        var v8 int8 = (v4 * v4)
        v9 := GenS1{F0: 228, F1: 113}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: -462}
    v3 += 32767
    v11 := GenS1{F0: 0, F1: 976}
    if (-2.25 >= (((((0.1 * (((0.0 - 0.1) - -2.25) * (0.0 - -2.25))) / (1.5 / float64((1e10 + (3.14159 + (3.14159 / 1.5)))))) - -2.25) + float64((1e10 + float32((-2.25 / 0.1))))) / 1.0)) {
        v3 += (v3 % (((v3 << ((0) & 63))) | 1))
    } else {
        if false {
            v12 := GenS2{F0: -1, F1: 65535, F2: (v2 | (v2 * (v2 << ((18446744073709551615) & 63))))}
            var v13 int16 = -32768
            _ = v12
            _ = v13
        } else {
            if ((32767 >= 1) && (!(((((v2 | v2) == 0) && false) && true) || false))) {
                v3 -= (v3 - 0)
                fmt.Println(v2, v2, v3)
            }
        }
    }
    var v14 uint8 = 127
    var v15 float64 = 1e10
    fmt.Println(v3, v15)
    var v16 int64 = 1
    for fi17 := 0; fi17 < 2; fi17++ {
        switch fi17 {
        case 32768:
            v15 -= float64((0.1 - 3.14159))
        case 755:
            switch fi17 {
            case 128:
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 uint = 65535
                    _ = fi18
                    _ = v19
                }
            default:
                fmt.Println(v15, v2, v4)
            }
        case 9223372036854775807:
            var v20 string = ("a" + ("foo bar" + "a"))
            _ = v20
        default:
            var v21 int64 = v16
            _ = v21
        }
        if (false || ((("hello" + ("\n\t" + "a")) + "\n\t") != "foo bar")) {
            v22 := GenS2{F0: v4, F1: 65535, F2: (v2 & -903)}
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS1{F0: 540, F1: 512}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v25 := GenS2{F0: v4, F1: 65535, F2: (v2 | (v2 ^ v2))}
                if (int(v14) < (fi23 - 265)) {
                    var v26 int32 = 1
                    if false {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = len(v5)
                        var v27 uint16 = 0
                        _ = v27
                    } else {
                        _ = len(v1)
                        v26 = (^(v26))
                        if true {
                            v28 := GenS2{F0: 0, F1: 556, F2: -249}
                            _ = v28
                        }
                    }
                    _ = v26
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            v16 += v16
            _ = v22
        } else {
            if true {
                _ = len(v5)
                var v29 uint64 = 18446744073709551615
                for fi30 := 0; fi30 < 18; fi30++ {
                    v3 = (v3 + v3)
                    _ = fi30
                }
                _ = v29
            } else {
                if true {
                    v4 = (v4 + v4)
                    v31 := GenS3{F0: v14, F1: true}
                    _ = v31
                }
            }
            var v32 uint64 = 786
            _ = v32
        }
        v33 := GenS4{F0: v3, F1: "hello", F2: v2}
        _ = fi17
        _ = v33
    }
    switch -65 {
    case 128:
        fmt.Println(v4, v16)
    default:
        var v34 int = -1
        _ = v34
    }
    _ = len(v1)
    for fi35 := 0; fi35 < 9; fi35++ {
        fmt.Println(v14, v16, v4)
        if true {
            var v36 uint32 = 1
            var v37 uint16 = (^(v3))
            for fi38 := 0; fi38 < 18; fi38++ {
                if true {
                    _ = len(v5)
                }
                switch fi38 {
                case 345:
                    v39 := GenS5{F0: v3, F1: "hello", F2: (v36 >> ((18446744073709551615) & 63))}
                    _ = v39
                }
                if false {
                    v2 = (v2 & 1)
                    var v40 bool = (!((1e10 >= float32(((v15 * v15) + v15))) || false))
                    if len(v5) > 0 {
                        v5[0] = (v36 | v36)
                    }
                    _ = v40
                }
                var v41 float32 = 0.1
                _ = fi38
                _ = v41
            }
            _ = v36
            _ = v37
        }
        v4 = int8(v14)
        _ = fi35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
