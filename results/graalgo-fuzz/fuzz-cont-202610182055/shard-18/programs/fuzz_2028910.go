package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 bool; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint; F1 bool; F2 uint32 }

func main() {
    var v1 uint16 = 871
    var v2 uint8 = 2
    var v3 int16 = 32767
    var v4 int16 = v3
    v5 := make(map[int]string)
    var v6 uint16 = 962
    v7 := GenS1{F0: 65535, F1: 100, F2: "foo bar"}
    v4 += (v4 | v4)
    v8 := GenS2{F0: 1}
    var v9 int16 = (^(v4))
    fmt.Println(v1, v2)
    var v10 int32 = -1
    switch ("hello" + "foo bar") {
    case "foo bar":
        var v11 int = -2147483648
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 5; fi12++ {
            if (v10 < (v10 / ((v10) | 1))) {
                var v13 string = ""
                if false {
                    switch fi12 {
                    case 9223372036854775807:
                        v14 := GenS2{F0: 922}
                        _ = v14
                    }
                    var v15 int64 = 1
                    switch fi12 {
                    case 4294967295:
                        for fi16 := 0; fi16 < 4; fi16++ {
                            v13 = ("a" + "")
                            _ = fi16
                        }
                    case 155:
                        if ((!(!true)) || (((fi12 / ((fi12) | 1)) <= (fi12 >> ((18446744073709551615) & 63))) && ((!(!(!(-2.25 <= 0.1)))) && false))) {
                            var v17 int64 = v15
                            var v18 uint8 = (^(v2))
                            _ = v17
                            _ = v18
                        } else {
                            if false {
                                if (v4 != v4) {
                                    for fi19 := 0; fi19 < 10; fi19++ {
                                        fmt.Println(v13)
                                        v20 := GenS3{F0: 735, F1: 0}
                                        v4 = (v4 * v4)
                                        if false {
                                            var v21 int32 = 2147483647
                                            if true {
                                                v22 := GenS3{F0: 432, F1: -478}
                                                switch v13 {
                                                case "":
                                                    var v23 float32 = 1.5
                                                    _ = v23
                                                case "\n\t":
                                                    v24 := GenS1{F0: 2147483647, F1: -54, F2: v13}
                                                    _ = v24
                                                default:
                                                    v25 := []int{127, (fi12 << ((213) & 63)), (^(fi19))}
                                                    _ = v25
                                                }
                                                _ = v22
                                            }
                                            fmt.Println(v3)
                                            _ = v21
                                        } else {
                                            v26 := GenS4{F0: false, F1: 2147483648, F2: (v4 | (-(v3)))}
                                            switch (v13 + v13) {
                                            case "hello":
                                                v10 = (v10 << ((32) & 63))
                                            case "":
                                                v5[(fi19 | -1)] = v13
                                            }
                                            switch fi19 {
                                            case 577:
                                                var v27 int32 = (v10 % (((v10 ^ v10)) | 1))
                                                _ = v27
                                            }
                                            _ = v26
                                        }
                                        if (v15 > v15) {
                                            var v28 int8 = 1
                                            if (true || (!false)) {
                                                v1 = (v1 ^ (v1 & v1))
                                            } else {
                                                v10 -= (v10 | 1)
                                                v10 = (v10 % (((v10 + v10)) | 1))
                                            }
                                            v29 := GenS2{F0: int64(fi19)}
                                            _ = v28
                                            _ = v29
                                        }
                                        var v30 bool = false
                                        _ = fi19
                                        _ = v20
                                        _ = v30
                                    }
                                }
                            }
                        }
                    case 658:
                        for fi31 := 0; fi31 < 17; fi31++ {
                            var v32 string = "foo bar"
                            _ = fi31
                            _ = v32
                        }
                    default:
                        v6 -= 65535
                    }
                    _ = v15
                } else {
                    v33 := GenS2{F0: -9223372036854775808}
                    v34 := []int{fi12}
                    v35 := v5[(fi12 & (-(fi12)))]
                    _ = v33
                    _ = v34
                    _ = v35
                }
                v4 -= 1
                _ = v13
            } else {
                for fi36 := 0; fi36 < 18; fi36++ {
                    if ((fi36 % ((fi12) | 1)) > 9223372036854775807) {
                        if (false && false) {
                            var v37 int8 = 100
                            _ = v37
                        }
                        v38 := []int{(fi12 << ((382) & 63)), (fi36 >> ((2) & 63))}
                        _ = len(v38)
                        var v39 uint64 = 4294967295
                        v6 -= v6
                        v40 := GenS5{F0: 4294967295, F1: ((fi12 >= 4294967295) || false), F2: 1}
                        v4 = (v9 & -484)
                        _ = v38
                        _ = v39
                        _ = v40
                    } else {
                        v5[32768] = "\n\t"
                        if true {
                            fmt.Println(v1, fi36, v4)
                            v41 := GenS4{F0: (fi36 == (fi12 + fi36)), F1: 1, F2: -32768}
                            _ = v41
                        }
                        var v42 bool = true
                        v43 := []int{}
                        if len(v43) >= 2 {
                            v43 = v43[1:len(v43)]
                        }
                        _ = v42
                        _ = v43
                    }
                    var v44 int32 = -2147483648
                    _ = fi36
                    _ = v44
                }
            }
            v5[(fi12 - (fi12 + fi12))] = "a"
            var v45 int16 = (v9 << ((326) & 63))
            _ = fi12
            _ = v45
        }
    case "\n\t":
        if ("hello" > ("foo bar" + "hello")) {
            switch 2147483647 {
            case 905:
                var v46 int8 = -44
                _ = v46
            }
            var v47 string = "a"
            _ = v47
        }
    default:
        var v48 int32 = (v10 & v10)
        _ = v48
    }
    fmt.Println(v4)
    v49 := []int{-1, 128}
    if len(v49) > 0 {
        v49[0] = 128
    }
    v50 := GenS1{F0: uint32(v1), F1: 1, F2: "hello"}
    v49 = append(v49, -1)
    if (739 == v6) {
        for fi51 := 0; fi51 < 6; fi51++ {
            var v52 uint = 13
            var v53 int = (^(fi51))
            _ = fi51
            _ = v52
            _ = v53
        }
    }
    var v54 uint16 = (v6 * v6)
    switch (("foo bar" + "") + (("" + "\n\t") + (("foo bar" + "\n\t") + "a"))) {
    case "":
        var v55 int = -9223372036854775808
        _ = v55
    case "hello":
        v4 -= v4
    default:
        var v56 uint = 132
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v54
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
