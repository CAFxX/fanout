package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16; F1 int8; F2 bool }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint; F2 int64 }
type GenS7 struct { F0 uint; F1 int; F2 uint32 }
type GenS8 struct { F0 int; F1 int16 }

func main() {
    var v1 uint8 = 11
    var v2 uint32 = 2147483648
    var v3 uint16 = 32767
    switch "\n\t" {
    case "a":
        switch ((("\n\t" + ("hello" + "a")) + ("" + "a")) + "") {
        case "a":
            for fi4 := 0; fi4 < 20; fi4++ {
                fmt.Println(fi4)
                _ = fi4
            }
        case "foo bar":
            if (v2 == 872) {
                var v5 float32 = 3.14159
                if (((1.5 + (1e10 + 3.14159)) + -2.25) != (-2.25 + (float64(v5) + ((((float64(0.1) + -2.25) + 0.1) + 0.0) / 3.14159)))) {
                    v6 := []int{128, 9223372036854775807}
                    var v7 uint8 = 42
                    var v8 int32 = -1
                    _ = v6
                    _ = v7
                    _ = v8
                } else {
                    if (!(-1 != -1)) {
                        var v9 float32 = v5
                        v5 = v5
                        fmt.Println(v9, v2, v1)
                        _ = v9
                    } else {
                        if true {
                            var v10 int32 = 785
                            v11 := GenS1{F0: -865, F1: 551, F2: v10}
                            var v12 int16 = 42
                            _ = v10
                            _ = v11
                            _ = v12
                        } else {
                            if (false && (839 < 18446744073709551615)) {
                                switch "foo bar" {
                                case "a":
                                    for fi13 := 0; fi13 < 6; fi13++ {
                                        for fi14 := 0; fi14 < 5; fi14++ {
                                            v15 := []int{2147483647, fi13, 2147483647, (fi14 ^ 127)}
                                            var v16 uint = 939
                                            _ = len(v15)
                                            v17 := GenS2{F0: -32768, F1: 127, F2: false}
                                            _ = fi14
                                            _ = v15
                                            _ = v16
                                            _ = v17
                                        }
                                        v1 = (v1 >> ((674) & 63))
                                        _ = fi13
                                    }
                                case "\n\t":
                                    v5 = 3.14159
                                case "":
                                    v18 := GenS3{F0: 42, F1: (v2 % (((v2 % (((^(v2))) | 1))) | 1))}
                                    _ = v18
                                default:
                                    if ((!(1 < 114)) && (42 >= 638)) {
                                        fmt.Println(v1, v3, v1)
                                        for fi19 := 0; fi19 < 5; fi19++ {
                                            v20 := GenS4{F0: 42}
                                            if ((v3 - (v3 * 42)) >= (^(v3))) {
                                                fmt.Println(v1, v3)
                                                var v21 int32 = -469
                                                if true {
                                                    for fi22 := 0; fi22 < 12; fi22++ {
                                                        if (!true) {
                                                            v23 := []int{(fi19 ^ int(v21)), fi19, fi22}
                                                            var v24 int64 = 1
                                                            _ = len(v23)
                                                            _ = v23
                                                            _ = v24
                                                        }
                                                        _ = fi22
                                                    }
                                                    var v25 uint32 = v2
                                                    v3 = uint16(v1)
                                                    _ = v25
                                                }
                                                _ = v21
                                            } else {
                                                v3 -= 65535
                                            }
                                            v26 := GenS5{F0: ("a" + (("\n\t" + ("a" + "foo bar")) + "")), F1: 42, F2: 522}
                                            _ = fi19
                                            _ = v20
                                            _ = v26
                                        }
                                        var v27 uint = 18446744073709551615
                                        _ = v27
                                    }
                                }
                            }
                        }
                    }
                    v28 := []int{716, 2147483647, -9223372036854775808}
                    _ = v28
                }
                _ = v5
            } else {
                if false {
                    v29 := GenS1{F0: 4294967295, F1: 954, F2: -34}
                    _ = v29
                }
            }
        default:
            v30 := []int{-2147483648, 32768, 217, 42}
            _ = v30
        }
    default:
        var v31 uint8 = 33
        _ = v31
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        fmt.Println(v3, v1, v2)
        _ = fi32
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 uint32 = (v2 / ((v2) | 1))
        var v35 int64 = 42
        _ = fi33
        _ = v34
        _ = v35
    }
    switch 1 {
    case -1:
        var v36 uint32 = v2
        _ = v36
    case 42:
        switch 32767 {
        case 42:
            v37 := []int{-9223372036854775808, 2147483647, 9223372036854775807, 4294967295}
            _ = v37
        case 127:
            if (-902 <= -140) {
                for fi38 := 0; fi38 < 19; fi38++ {
                    v39 := GenS4{F0: (v2 ^ (v2 & 25))}
                    _ = fi38
                    _ = v39
                }
                v40 := GenS6{F0: (^(v3)), F1: 18446744073709551615, F2: -493}
                _ = v40
            } else {
                v41 := GenS1{F0: 549, F1: 205, F2: 0}
                if (!false) {
                    var v42 uint8 = (v1 * v1)
                    v42 += (v1 & v1)
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 4; fi43++ {
                        v44 := GenS7{F0: 65535, F1: -544, F2: v2}
                        _ = fi43
                        _ = v44
                    }
                    v45 := GenS4{F0: (v2 << ((871) & 63))}
                    _ = v45
                }
                _ = v41
            }
        case 2147483647:
            switch "foo bar" {
            case "":
                for fi46 := 0; fi46 < 16; fi46++ {
                    for fi47 := 0; fi47 < 5; fi47++ {
                        switch (fi47 >> ((255) & 63)) {
                        case 1:
                            v2 = (v2 * v2)
                        default:
                            v48 := []int{-530, (fi46 ^ (fi46 - fi47)), (fi46 - (fi46 >> ((0) & 63))), fi47}
                            _ = v48
                        }
                        v1 = (^(v1))
                        var v49 int64 = -69
                        _ = fi47
                        _ = v49
                    }
                    v2 = v2
                    _ = fi46
                }
            case "hello":
                v50 := GenS8{F0: 4294967295, F1: 0}
                _ = v50
            case "a":
                var v51 bool = (54 >= 18446744073709551615)
                _ = v51
            }
        default:
            v1 = 1
        }
    }
    v52 := []int{2147483647}
    if (v1 == v1) {
        v2 += (v2 & uint32(v1))
        v1 = v1
    } else {
        v2 += v2
        v3 = (v3 % ((v3) | 1))
    }
    fmt.Println(v3, v3, v2)
    v2 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v52:", len(v52), v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
}
