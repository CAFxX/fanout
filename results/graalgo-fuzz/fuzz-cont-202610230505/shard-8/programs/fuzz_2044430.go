package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 string; F1 uint32; F2 int }
func genFn0(p0 string, p1 uint) uint32 {
    fmt.Println(p0, p1)
    var v4 int64 = -9223372036854775808
    switch 128 {
    case -9223372036854775808:
        fmt.Println(p1)
    case -2147483648:
        fmt.Println(p1)
    case 0:
        p0 = "hello"
    default:
        var v5 uint64 = 42
        _ = v5
    }
    v6 := GenS1{F0: 638}
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 2
}
func genFn1() int64 {
    v7 := GenS1{F0: 93}
    if (255 != 42) {
        if (!true) {
            if (-284 > 32767) {
                for fi8 := 0; fi8 < 11; fi8++ {
                    fmt.Println(fi8)
                    var v9 float64 = (-2.25 * ((1e10 - 3.14159) - 0.0))
                    _ = fi8
                    _ = v9
                }
            }
            var v10 uint32 = 42
            _ = v10
        }
        v11 := []int{}
        var v12 int64 = -496
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 472}
    var v14 int8 = 127
    if false {
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 19; fi16++ {
                v14 -= v14
                _ = fi16
            }
            v17 := []int{}
            for fi18 := 0; fi18 < 9; fi18++ {
                if true {
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    var v19 uint64 = 42
                    _ = v19
                } else {
                    var v20 bool = (true || true)
                    _ = v20
                }
                _ = fi18
            }
            _ = fi15
            _ = v17
        }
        v14 += (-(v14))
    }
    _ = v13
    _ = v14
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := []int{32767, 32768, 4294967295, 127}
    v2 := []uint32{42, 720}
    var v3 float32 = -2.25
    if (733 <= 1) {
        v3 -= 0.1
    } else {
        if len(v1) > 0 {
            v1[0] = 42
        }
    }
    v3 = v3
    v3 = v3
    v3 = v3
    if (42 != -191) {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 uint64 = 42
            for fi23 := 0; fi23 < 14; fi23++ {
                fmt.Println(v3, v3, v22)
                v22 = (^(v22))
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v1) > 0 {
            v1[0] = 629
        }
    }
    var v24 string = ""
    switch -760 {
    case 306:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -2147483648:
        for fi25 := 0; fi25 < 9; fi25++ {
            if (!true) {
                var v26 string = v24
                if (!true) {
                    v27 := GenS3{F0: 2, F1: 127, F2: 9}
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 16; fi28++ {
                        if false {
                            if true {
                                v24 = v26
                                var v29 uint32 = 2147483647
                                fmt.Println(v29, v3, v3)
                                _ = v29
                            }
                        } else {
                            v30 := GenS2{F0: 950}
                            if (4294967295 >= 18446744073709551615) {
                                v3 = -2.25
                                switch v24 {
                                case "hello":
                                    v31 := GenS3{F0: 255, F1: 79, F2: 127}
                                    _ = v31
                                case "foo bar":
                                    if true {
                                        v32 := GenS3{F0: 242, F1: 1, F2: 1}
                                        _ = v32
                                    } else {
                                        if true {
                                            var v33 int32 = 0
                                            _ = v33
                                        } else {
                                            v34 := GenS1{F0: 2}
                                            _ = v34
                                        }
                                        v24 = (((v24 + ("\n\t" + ("a" + ("\n\t" + (v26 + (v26 + (v24 + ("" + ("a" + "\n\t"))))))))) + v26) + "a")
                                        var v35 int32 = 103
                                        if false {
                                            var v36 uint = 42
                                            switch "\n\t" {
                                            case "":
                                                v2 = append(v2, genFn0(v24, 42))
                                            case "hello":
                                                var v37 float32 = 1.5
                                                _ = v37
                                            }
                                            var v38 bool = (true && (!(127 > 216)))
                                            _ = v36
                                            _ = v38
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = fi25
                                            }
                                            var v39 uint8 = 127
                                            _ = v39
                                        }
                                        var v40 float32 = float32(-2.25)
                                        _ = v35
                                        _ = v40
                                    }
                                case "a":
                                    v41 := GenS4{F0: false, F1: 100}
                                    _ = v41
                                }
                                v42 := GenS5{F0: "", F1: 1, F2: (fi28 >> ((255) & 63))}
                                _ = v42
                            }
                            _ = v30
                        }
                        v43 := GenS2{F0: 2}
                        if ((18446744073709551615 > 18446744073709551615) || false) {
                            v44 := GenS5{F0: "\n\t", F1: 42, F2: fi28}
                            _ = v44
                        }
                        _ = fi28
                        _ = v43
                    }
                    fmt.Println(v3, v26)
                    v45 := GenS2{F0: 2}
                    _ = v45
                }
                _ = v26
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v3 = 0.0
            _ = fi25
        }
    case 1:
        fmt.Println(v24, v24)
    }
    for fi46 := 0; fi46 < 2; fi46++ {
        v47 := GenS2{F0: 4294967295}
        _ = fi46
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
