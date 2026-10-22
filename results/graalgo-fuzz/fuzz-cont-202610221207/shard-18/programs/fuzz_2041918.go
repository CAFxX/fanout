package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int16; F1 int64; F2 int }
type GenS4 struct { F0 uint32; F1 int; F2 uint16 }
func genFn0(p0 uint16, p1 uint8, p2 uint16) int32 {
    var v4 int32 = 353
    v5 := []int{32768, 438}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}
func genFn1() bool {
    for fi6 := 0; fi6 < 8; fi6++ {
        switch (fi6 % ((fi6) | 1)) {
        case 127:
            switch "\n\t" {
            case "foo bar":
                var v7 bool = false
                _ = v7
            case "\n\t":
                fmt.Println(fi6)
            case "":
                fmt.Println(fi6)
            default:
                v8 := GenS1{F0: 582}
                _ = v8
            }
        case -9223372036854775808:
            v9 := GenS1{F0: 42}
            _ = v9
        case 32768:
            var v10 int8 = 8
            _ = v10
        default:
            var v11 int = fi6
            _ = v11
        }
        _ = fi6
    }
    var v12 uint8 = 255
    _ = v12
    return (0.1 <= 0.0)
}

func main() {
    var v1 int32 = 42
    var v2 int = 32768
    v3 := make(map[uint32]uint32)
    for fi13 := 0; fi13 < 9; fi13++ {
        fmt.Println(v2, fi13, v1)
        _ = fi13
    }
    v2 = (v2 ^ v2)
    var v14 int = (^(v2))
    _ = len(v3)
    fmt.Println(v14)
    fmt.Println(v14)
    var v15 bool = (((2 <= 186) || true) && (690 <= 89))
    if v15 {
        v16 := GenS2{F0: 0, F1: "\n\t"}
        v15 = v15
        _ = v16
    } else {
        switch "\n\t" {
        case "\n\t":
            v15 = true
        case "":
            if (!v15) {
                v1 = (v1 % (((v1 + v1)) | 1))
                v17 := []int{426, 772}
                var v18 int16 = -55
                _ = v17
                _ = v18
            }
        }
        if v15 {
            for fi19 := 0; fi19 < 9; fi19++ {
                if (!(false && (true || false))) {
                    v20 := []int{}
                    if v15 {
                        var v21 uint8 = 0
                        if (1.5 > float64((3.14159 / 3.14159))) {
                            for fi22 := 0; fi22 < 12; fi22++ {
                                fmt.Println(v21, fi19, v15)
                                _ = fi22
                            }
                            var v23 uint8 = (^(v21))
                            v2 = 128
                            _ = v23
                        }
                        var v24 bool = (-41 <= -128)
                        _ = v21
                        _ = v24
                    } else {
                        _ = len(v3)
                        v25 := GenS3{F0: 59, F1: 9223372036854775807, F2: 0}
                        var v26 float32 = 1e10
                        _ = v25
                        _ = v26
                    }
                    _ = v20
                } else {
                    if false {
                        switch (v14 & (-(fi19))) {
                        case -659:
                            var v27 bool = (!(v15 && (42 <= -67)))
                            _ = v27
                        case 128:
                            switch ("\n\t" + (("a" + "a") + ("a" + "\n\t"))) {
                            case "foo bar":
                                delete(v3, 2)
                            case "a":
                                fmt.Println(fi19)
                            default:
                                v2 += (v2 * v14)
                            }
                        case 464:
                            switch "\n\t" {
                            case "foo bar":
                                var v28 int64 = 9223372036854775807
                                _ = v28
                            case "\n\t":
                                for fi29 := 0; fi29 < 11; fi29++ {
                                    v30 := v3[529]
                                    switch ("" + ((("" + ("a" + "hello")) + "a") + "\n\t")) {
                                    case "\n\t":
                                        var v31 int64 = 1
                                        _ = v31
                                    case "hello":
                                        if (!(-1 > -555)) {
                                            var v32 uint32 = (v30 + (v30 << ((65535) & 63)))
                                            _ = v32
                                        }
                                    case "":
                                        v33 := GenS4{F0: 2147483648, F1: 32768, F2: 356}
                                        _ = v33
                                    }
                                    if (!(128 != fi29)) {
                                        var v34 int = fi19
                                        _ = v34
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                            case "hello":
                                var v35 int32 = (v1 << ((305) & 63))
                                _ = v35
                            default:
                                v36 := GenS4{F0: 535, F1: -2147483648, F2: 1}
                                _ = v36
                            }
                        default:
                            if v15 {
                                v37 := GenS3{F0: 0, F1: 0, F2: (v14 * (fi19 * (fi19 % ((v14) | 1))))}
                                _ = v37
                            }
                        }
                        var v38 float32 = 3.14159
                        _ = v38
                    }
                    if v15 {
                        _ = len(v3)
                        fmt.Println(v15, v15)
                    }
                }
                _ = fi19
            }
            for fi39 := 0; fi39 < 19; fi39++ {
                var v40 int64 = 1
                v2 = (v2 << ((250) & 63))
                fmt.Println(v40, v15)
                _ = fi39
                _ = v40
            }
        } else {
            v41 := GenS2{F0: 550, F1: (("a" + ("foo bar" + "\n\t")) + (("a" + "") + ("\n\t" + "a")))}
            v42 := GenS1{F0: 42}
            _ = v41
            _ = v42
        }
    }
    v43 := []int{32768, (v14 >> ((207) & 63)), (v14 >> ((1) & 63))}
    var v44 int = (^(v14))
    fmt.Println(v14, v14, v14)
    v45 := v3[638]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
}
