package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int8 }
func genFn0(p0 uint, p1 bool, p2 int8) uint32 {
    v4 := GenS1{F0: -1}
    switch 2147483647 {
    case -579:
        var v5 uint32 = 748
        _ = v5
    }
    v6 := []int{955, -392}
    fmt.Println(p2, p0, p0)
    switch -993 {
    case -9223372036854775808:
        p2 = p2
    case -1:
        _ = len(v6)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 1
}

func main() {
    v1 := []int{2147483647, 128, -12, 175}
    v2 := []int{127, 53}
    v3 := []uint32{536, 2147483648, 2, 4294967295}
    if true {
        var v7 float32 = float32(0.0)
        var v8 int32 = 2147483647
        if len(v3) > 0 {
            v3[0] = genFn0(57, true, -20)
        }
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -103}
    var v10 uint8 = 154
    var v11 uint8 = 1
    if (false && (32767 == -687)) {
        v10 -= v11
        var v12 int32 = -1
        v13 := GenS1{F0: (v12 << ((255) & 63))}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        _ = len(v3)
        var v15 uint8 = v10
        _ = fi14
        _ = v15
    }
    _ = len(v1)
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint8 = (^(v11))
        v18 := GenS2{F0: "foo bar"}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: -2147483648}
    if (!true) {
        var v20 int16 = 0
        if true {
            if true {
                v20 = v20
                if (128 >= (v11 | v10)) {
                    if (42 != 1) {
                        v10 = 0
                        v20 = v20
                        switch int(v20) {
                        case 42:
                            for fi21 := 0; fi21 < 11; fi21++ {
                                var v22 int16 = (v20 / ((32767) | 1))
                                switch ("" + "a") {
                                case "hello":
                                    var v23 int = 2147483647
                                    _ = v23
                                case "\n\t":
                                    v10 += (v11 ^ (v10 / (((v11 * v11)) | 1)))
                                default:
                                    v22 -= 576
                                }
                                if false {
                                    for fi24 := 0; fi24 < 10; fi24++ {
                                        v25 := GenS2{F0: "foo bar"}
                                        _ = fi24
                                        _ = v25
                                    }
                                    if (0 == 98) {
                                        var v26 uint32 = 65535
                                        _ = v26
                                    } else {
                                        var v27 uint = 42
                                        for fi28 := 0; fi28 < 18; fi28++ {
                                            if true {
                                                var v29 uint16 = 65535
                                                _ = v29
                                            } else {
                                                var v30 int8 = 121
                                                _ = v30
                                            }
                                            v22 = 0
                                            _ = fi28
                                        }
                                        if true {
                                            v11 -= (v10 & v10)
                                            for fi31 := 0; fi31 < 7; fi31++ {
                                                _ = len(v1)
                                                _ = fi31
                                            }
                                            v20 += v20
                                        }
                                        _ = v27
                                    }
                                    var v32 uint32 = 463
                                    _ = v32
                                }
                                _ = fi21
                                _ = v22
                            }
                        case 743:
                            switch ("a" + ("a" + "")) {
                            case "a":
                                var v33 int16 = v20
                                _ = v33
                            }
                        default:
                            v34 := GenS3{F0: "\n\t", F1: -128}
                            _ = v34
                        }
                    }
                    var v35 uint16 = 676
                    _ = v35
                } else {
                    var v36 uint64 = 0
                    v36 -= 2
                    _ = v36
                }
            } else {
                var v37 int16 = (v20 << ((42) & 63))
                v3 = append(v3, 4294967295)
                v10 = (v11 | (v11 ^ 127))
                _ = v37
            }
        }
        _ = v20
    } else {
        var v38 uint = 1
        for fi39 := 0; fi39 < 9; fi39++ {
            switch ("" + "") {
            case "":
                v38 += (v38 + (v38 - 18446744073709551615))
            default:
                fmt.Println(fi39, fi39)
            }
            var v40 int = -1
            _ = fi39
            _ = v40
        }
        var v41 int16 = 1000
        _ = v38
        _ = v41
    }
    if true {
        var v42 int = -1
        _ = v42
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        fmt.Println(v11, fi43, v10)
        v11 -= (v11 * 2)
        var v44 bool = true
        _ = fi43
        _ = v44
    }
    var v45 uint16 = 1
    var v46 uint16 = v45
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v9
}
