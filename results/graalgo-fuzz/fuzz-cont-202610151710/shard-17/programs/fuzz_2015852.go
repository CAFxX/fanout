package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int; F1 uint; F2 bool }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint; F1 int8; F2 int }
type GenS7 struct { F0 int64; F1 int64 }
func genFn0(p0 uint32) uint16 {
    v4 := GenS1{F0: 858, F1: "a"}
    var v5 uint32 = (^(p0))
    var v6 uint32 = p0
    var v7 int16 = -285
    v8 := GenS1{F0: 0, F1: ("\n\t" + "hello")}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int16, p1 int64) bool {
    var v9 int16 = p0
    v9 = (v9 - p0)
    var v10 bool = (-1 > (p1 | -9223372036854775808))
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (true && v10)
}

func main() {
    v1 := []string{"", "hello", "\n\t", ("a" + "hello")}
    var v2 int8 = -1
    var v3 uint64 = 979
    switch -739 {
    case 128:
        if (false && (2 >= v3)) {
            var v11 float64 = 1.5
            _ = v11
        } else {
            switch ("a" + "\n\t") {
            case "hello":
                switch -224 {
                case 2147483647:
                    v12 := GenS2{F0: (!false), F1: -127}
                    _ = v12
                case 4294967295:
                    _ = len(v1)
                case -9223372036854775808:
                    v3 -= v3
                default:
                    if ((-696 == -966) && (!false)) {
                        for fi13 := 0; fi13 < 9; fi13++ {
                            var v14 float32 = (0.0 * 1.5)
                            if len(v1) > 0 {
                                v1[0] = ("" + "foo bar")
                            }
                            if true {
                                for fi15 := 0; fi15 < 10; fi15++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v16 := GenS1{F0: 388, F1: "a"}
                                    v17 := GenS3{F0: (fi13 / ((32767) | 1)), F1: 230, F2: (!(42 >= 140))}
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                            }
                            _ = fi13
                            _ = v14
                        }
                    } else {
                        _ = len(v1)
                    }
                }
            }
            v18 := GenS4{F0: 65535, F1: 9223372036854775807}
            for fi19 := 0; fi19 < 13; fi19++ {
                switch -569 {
                case -249:
                    fmt.Println(v2)
                case 2147483647:
                    var v20 uint32 = 1
                    _ = v20
                }
                v1 = append(v1, ("foo bar" + "\n\t"))
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                _ = fi19
            }
            _ = v18
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v2 -= v2
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := GenS5{F0: (v3 / ((18446744073709551615) | 1)), F1: 2}
        _ = fi21
        _ = v22
    }
    v3 = v3
    v23 := GenS5{F0: (^(v3)), F1: genFn0(4294967295)}
    if (1 > -204) {
        var v24 int = -2147483648
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v24 -= 4294967295
        _ = v24
    } else {
        v2 += (^(v2))
        for fi25 := 0; fi25 < 18; fi25++ {
            if false {
                fmt.Println(v3, v3)
            }
            var v26 uint64 = (v3 - (v3 ^ (^(v3))))
            _ = len(v1)
            _ = fi25
            _ = v26
        }
        if ((v3 - (v3 >> ((42) & 63))) >= v3) {
            fmt.Println(v2, v2)
        } else {
            v2 = 100
            v3 = (v3 / ((42) | 1))
            _ = len(v1)
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v27 := GenS3{F0: -926, F1: 755, F2: true}
    for fi28 := 0; fi28 < 20; fi28++ {
        _ = len(v1)
        v3 -= 851
        if (920 != 0) {
            v3 += v3
            if (!(18446744073709551615 > 4294967295)) {
                switch (fi28 * (-(fi28))) {
                case -122:
                    v29 := GenS5{F0: v3, F1: 658}
                    _ = v29
                default:
                    if (8 < 172) {
                        if false {
                            v3 += (v3 / ((v3) | 1))
                        } else {
                            v2 = v2
                        }
                    } else {
                        for fi30 := 0; fi30 < 20; fi30++ {
                            fmt.Println(v2, fi30)
                            v3 = v3
                            _ = fi30
                        }
                        v3 = (v3 & v3)
                        switch fi28 {
                        case -2147483648:
                            v2 = v2
                        case 2147483647:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            if true {
                                v3 = 4294967295
                                for fi31 := 0; fi31 < 18; fi31++ {
                                    var v32 string = ("foo bar" + ("a" + ""))
                                    v2 -= v2
                                    _ = fi31
                                    _ = v32
                                }
                                for fi33 := 0; fi33 < 19; fi33++ {
                                    v34 := GenS5{F0: v3, F1: genFn0(2147483648)}
                                    v1 = append(v1, "\n\t")
                                    _ = fi33
                                    _ = v34
                                }
                            }
                        }
                    }
                }
                v2 -= 87
                var v35 float32 = float32(0.1)
                _ = v35
            }
        }
        _ = fi28
    }
    v3 = v3
    v2 = 42
    if (((!(2147483647 > 2147483647)) || (0.1 <= (((3.14159 - (1.5 / float32(3.14159))) + 1.5) + 3.14159))) || (!false)) {
        var v36 uint16 = 65535
        var v37 float32 = 3.14159
        _ = v36
        _ = v37
    } else {
        var v38 int8 = 42
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v38
    }
    v39 := GenS6{F0: 65535, F1: (v2 << ((4294967295) & 63)), F2: 4294967295}
    var v40 int8 = (v2 - v2)
    v2 += (^(v2))
    for fi41 := 0; fi41 < 19; fi41++ {
        fmt.Println(fi41)
        if false {
            fmt.Println(v2)
            var v42 uint = 255
            var v43 bool = (255 >= 0)
            _ = v42
            _ = v43
        }
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = fi41
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        switch (fi44 + (^(fi44))) {
        case 881:
            switch (fi44 - (-(fi44))) {
            case -1:
                var v45 int16 = -1
                _ = v45
            case -9223372036854775808:
                if (16 <= 127) {
                    v40 = v2
                    if true {
                        if (!((!true) || false)) {
                            v2 = 1
                            if false {
                                var v46 uint16 = genFn0(42)
                                var v47 uint16 = v46
                                if false {
                                    v2 = -128
                                    for fi48 := 0; fi48 < 15; fi48++ {
                                        switch fi48 {
                                        case 9223372036854775807:
                                            var v49 uint64 = (v3 | (v3 << ((18446744073709551615) & 63)))
                                            _ = v49
                                        case 42:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case -1:
                                            v3 = (^(v3))
                                        }
                                        var v50 int8 = (^(v2))
                                        _ = fi48
                                        _ = v50
                                    }
                                }
                                _ = v46
                                _ = v47
                            }
                        }
                    }
                }
            case 0:
                if true {
                    for fi51 := 0; fi51 < 3; fi51++ {
                        var v52 int = (fi44 & 4294967295)
                        v40 = (^(v40))
                        _ = fi51
                        _ = v52
                    }
                    fmt.Println(v2, fi44, fi44)
                }
            default:
                v53 := GenS7{F0: 0, F1: 1}
                _ = v53
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi44
    }
    var v54 int16 = -378
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v39
    _ = v40
    _ = v54
}
