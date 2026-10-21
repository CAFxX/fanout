package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int8; F1 int; F2 bool }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 string; F1 uint; F2 uint32 }
func genFn0() uint8 {
    var v6 int32 = -2147483648
    var v7 int8 = 42
    var v8 uint = 2
    _ = v6
    _ = v7
    _ = v8
    return 202
}
func genFn1(p0 uint, p1 uint, p2 uint8) uint64 {
    switch ("\n\t" + "a") {
    case "hello":
        v9 := []int{370}
        _ = v9
    case "a":
        var v10 int = 889
        _ = v10
    }
    var v11 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    v3 := []uint32{894}
    var v4 int32 = v2
    v5 := make(map[uint32]string)
    v4 += v2
    v5[20] = "hello"
    fmt.Println(v2)
    v12 := GenS1{F0: "\n\t"}
    switch (("hello" + "a") + ("\n\t" + "hello")) {
    case "foo bar":
        v13 := GenS1{F0: "foo bar"}
        _ = v13
    case "\n\t":
        _ = len(v1)
    case "hello":
        v4 -= v4
    }
    switch "foo bar" {
    case "foo bar":
        if (2 <= 132) {
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := GenS1{F0: ("\n\t" + "hello")}
                switch (("foo bar" + "") + "a") {
                case "\n\t":
                    fmt.Println(fi14)
                case "hello":
                    fmt.Println(fi14, fi14)
                case "foo bar":
                    v4 = v4
                }
                _ = fi14
                _ = v15
            }
            if true {
                v2 -= (v4 - v4)
                var v16 uint8 = 127
                v17 := GenS2{F0: 2147483648, F1: 0}
                _ = v16
                _ = v17
            } else {
                for fi18 := 0; fi18 < 12; fi18++ {
                    if (true && (128 >= 83)) {
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                    } else {
                        var v19 uint = 255
                        _ = v19
                    }
                    v3 = append(v3, 857)
                    _ = fi18
                }
                var v20 int64 = -9223372036854775808
                _ = v20
            }
        }
    case "\n\t":
        v4 = v2
    default:
        if false {
            v2 = (^(v2))
            var v21 uint = 4294967295
            _ = v21
        } else {
            for fi22 := 0; fi22 < 13; fi22++ {
                if (("hello" == ("hello" + "")) || (("hello" + "") > "foo bar")) {
                    for fi23 := 0; fi23 < 6; fi23++ {
                        var v24 bool = (!true)
                        switch (fi22 >> ((4294967295) & 63)) {
                        case 42:
                            v24 = v24
                        case 9223372036854775807:
                            var v25 string = "foo bar"
                            _ = v25
                        case 2147483647:
                            if (float64((0.0 + (0.0 * float32(3.14159)))) > 1e10) {
                                v26 := GenS3{F0: (fi22 & (fi22 << ((18446744073709551615) & 63))), F1: 0}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if len(v3) > 0 {
                                    v3[0] = 65535
                                }
                                if (2 <= 42) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "a":
                                            v27 := GenS4{F0: 1000, F1: 18446744073709551615}
                                            _ = v27
                                        case "foo bar":
                                            v24 = (!(("\n\t" + "") <= "a"))
                                        case "\n\t":
                                            var v28 int64 = 42
                                            _ = v28
                                        }
                                    case "foo bar":
                                        if false {
                                            fmt.Println(fi22, fi23)
                                            v4 -= v4
                                            v2 += (v2 % ((v2) | 1))
                                            v29 := GenS5{F0: 127, F1: fi23, F2: (42 > 128)}
                                            _ = v29
                                        } else {
                                            var v30 string = "\n\t"
                                            v4 = v4
                                            var v31 float64 = ((3.14159 * (0.1 / 1e10)) + (0.1 / 3.14159))
                                            _ = v30
                                            _ = v31
                                        }
                                    case "hello":
                                        _ = len(v3)
                                    default:
                                        fmt.Println(fi22)
                                    }
                                    var v32 int8 = 100
                                    _ = v32
                                }
                                v33 := GenS6{F0: 42, F1: 2147483647}
                                _ = v26
                                _ = v33
                            } else {
                                switch fi22 {
                                case 0:
                                    if ((0.0 / 1.5) != -2.25) {
                                        v34 := GenS7{F0: -1, F1: 0}
                                        var v35 uint8 = 128
                                        var v36 uint64 = 0
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                case 627:
                                    var v37 int16 = 1000
                                    _ = v37
                                case 128:
                                    var v38 float64 = 1e10
                                    _ = v38
                                default:
                                    var v39 uint64 = 4294967295
                                    _ = v39
                                }
                                delete(v5, 2147483648)
                                switch (("foo bar" + ((("a" + "a") + "") + "foo bar")) + "foo bar") {
                                case "\n\t":
                                    v3 = append(v3, 936)
                                case "":
                                    v4 += (v4 % (((v4 + 1)) | 1))
                                case "hello":
                                    var v40 int32 = (^(v2))
                                    _ = v40
                                default:
                                    if ((-67 >= 1) && ((!(290 <= (v2 & v2))) || (837 < fi22))) {
                                        if v24 {
                                            fmt.Println(fi23, v4, v4)
                                            var v41 bool = ((4294967295 >= 1) && v24)
                                            _ = v41
                                        }
                                        v24 = (v24 || true)
                                        v24 = v24
                                    } else {
                                        if v24 {
                                            v2 -= (v4 - (v4 | (v4 / ((v4) | 1))))
                                            switch "" {
                                            case "":
                                                v42 := GenS1{F0: "a"}
                                                _ = v42
                                            }
                                        } else {
                                            switch fi22 {
                                            case 128:
                                                fmt.Println(v4)
                                            }
                                            var v43 float32 = float32(1.5)
                                            fmt.Println(v24, v4, fi23)
                                            _ = v43
                                        }
                                    }
                                }
                            }
                        }
                        if (0.1 != 1.5) {
                            fmt.Println(v24)
                        } else {
                            if false {
                                var v44 int16 = 0
                                fmt.Println(fi22, v2)
                                fmt.Println(v2, v2, fi23)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v45 := GenS8{F0: ("hello" + ""), F1: 91, F2: 2147483647}
                                _ = v44
                                _ = v45
                            } else {
                                fmt.Println(fi23)
                                v2 = (v4 << ((599) & 63))
                            }
                            v46 := GenS5{F0: 0, F1: (fi22 * (fi22 / ((51) | 1))), F2: v24}
                            _ = v46
                        }
                        _ = fi23
                        _ = v24
                    }
                    var v47 uint16 = 32767
                    fmt.Println(v2)
                    _ = v47
                } else {
                    var v48 int8 = 0
                    if true {
                        var v49 float64 = 1e10
                        _ = v49
                    }
                    _ = v48
                }
                var v50 uint16 = 199
                var v51 uint8 = 91
                _ = fi22
                _ = v50
                _ = v51
            }
            var v52 int = -9223372036854775808
            _ = v52
        }
    }
    var v53 int32 = v2
    for fi54 := 0; fi54 < 17; fi54++ {
        v4 = v2
        v4 = 0
        v1["hello"] = 535
        _ = fi54
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
