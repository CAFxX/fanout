package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int32 }
type GenS2 struct { F0 int32; F1 int32; F2 int32 }
type GenS3 struct { F0 bool; F1 int64; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int32; F2 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 uint8; F2 int64 }
func genFn0(p0 int32) int8 {
    var v5 string = "hello"
    var v6 int8 = 24
    fmt.Println(p0)
    var v7 int64 = 0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 int) uint64 {
    fmt.Println(p0)
    v8 := []int{(p0 + 4294967295)}
    _ = p0
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 int16 = -32768
    var v4 bool = ((944 != 0) && true)
    if v4 {
        var v9 float32 = (3.14159 * 0.1)
        var v10 uint16 = 1
        _ = v10
        _ = v9
    }
    v4 = (-2147483648 == 42)
    var v11 int16 = v3
    fmt.Println(v4, v4)
    var v12 uint8 = 128
    v13 := GenS1{F0: (("" + (("" + "") + ("" + ((("hello" + ("hello" + (("a" + "foo bar") + "a"))) + "hello") + "a")))) + (("\n\t" + "a") + ("\n\t" + "foo bar"))), F1: -1, F2: -758}
    v14 := GenS1{F0: "a", F1: 1, F2: 1}
    v4 = (3.14159 < 1e10)
    if (1 == 2) {
        v4 = (true && v4)
        for fi15 := 0; fi15 < 15; fi15++ {
            if (v4 && v4) {
                var v16 uint64 = genFn1(2147483647)
                v17 := GenS2{F0: 42, F1: -1, F2: -558}
                switch "a" {
                case "a":
                    fmt.Println(v4, v12)
                }
                _ = v16
                _ = v17
            } else {
                switch "" {
                case "":
                    var v18 uint32 = 824
                    _ = v18
                default:
                    var v19 float64 = ((1e10 - 0.0) + float64((3.14159 / ((((-2.25 - 3.14159) + 1.5) / (1.5 * 0.1)) - float32(1e10)))))
                    _ = v19
                }
                if (0 >= 652) {
                    delete(v2, "")
                    for fi20 := 0; fi20 < 16; fi20++ {
                        v11 = 32767
                        switch ("a" + "") {
                        case "foo bar":
                            if (466 != 42) {
                                if v4 {
                                    switch (fi15 - -2147483648) {
                                    case 0:
                                        v21 := []int{fi15}
                                        _ = v21
                                    }
                                } else {
                                    switch fi15 {
                                    case -9223372036854775808:
                                        var v22 uint = 42
                                        _ = v22
                                    case 887:
                                        v23 := GenS1{F0: (((((("\n\t" + ("a" + "")) + "foo bar") + "\n\t") + "a") + "\n\t") + "hello"), F1: int8(v12), F2: 2147483647}
                                        _ = v23
                                    case -764:
                                        v24 := []int{82}
                                        _ = v24
                                    default:
                                        fmt.Println(v3)
                                    }
                                }
                                v25 := GenS2{F0: 0, F1: 1, F2: -1}
                                _ = v25
                            }
                        case "":
                            v26 := GenS3{F0: v4, F1: -692, F2: (v12 & v12)}
                            _ = v26
                        default:
                            var v27 uint32 = 941
                            v28 := GenS4{F0: int8(v12), F1: 829, F2: 255}
                            v11 += 1000
                            v3 = (v3 + -153)
                            _ = len(v2)
                            _ = v27
                            _ = v28
                        }
                        _ = fi20
                    }
                }
            }
            _ = fi15
        }
    } else {
        _ = len(v2)
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := []int{128, (fi29 - 127), (fi29 * fi29)}
        delete(v1, 71)
        v30 = append(v30, (fi29 - (fi29 >> ((1) & 63))))
        _ = fi29
        _ = v30
    }
    delete(v2, "a")
    switch "a" {
    case "hello":
        v31 := GenS1{F0: "hello", F1: 0, F2: 2147483647}
        _ = v31
    default:
        if (!(744 <= 42)) {
            for fi32 := 0; fi32 < 20; fi32++ {
                fmt.Println(fi32, fi32, v3)
                v11 = v3
                v33 := []int{fi32, fi32}
                _ = fi32
                _ = v33
            }
        } else {
            v11 += (v11 >> ((2) & 63))
            v3 = v11
            v12 += 187
        }
    }
    v34 := GenS5{F0: 42, F1: 64, F2: 611}
    if v4 {
        if false {
            v35 := GenS3{F0: (0 == genFn0(191)), F1: -928, F2: (v12 + (v12 | v12))}
            v12 += 128
            var v36 string = ("foo bar" + "")
            _ = v35
            _ = v36
        } else {
            if v4 {
                for fi37 := 0; fi37 < 10; fi37++ {
                    v4 = (12 > 0)
                    v2[("\n\t" + (("" + "foo bar") + "foo bar"))] = v4
                    if (507 != 4294967295) {
                        for fi38 := 0; fi38 < 13; fi38++ {
                            var v39 uint = 65535
                            v4 = (!v4)
                            if v4 {
                                var v40 int8 = -1
                                v3 -= v11
                                var v41 bool = v4
                                _ = v40
                                _ = v41
                            } else {
                                fmt.Println(fi37, fi38)
                                var v42 float64 = (0.0 + ((3.14159 * 3.14159) / ((-2.25 / 0.1) - ((3.14159 - 1e10) - float64(3.14159)))))
                                v43 := []int{(fi38 / (((^(fi37))) | 1)), (fi37 >> ((v39) & 63))}
                                _ = v42
                                _ = v43
                            }
                            _ = fi38
                            _ = v39
                        }
                        for fi44 := 0; fi44 < 11; fi44++ {
                            v45 := GenS4{F0: 42, F1: 42, F2: v12}
                            _ = fi44
                            _ = v45
                        }
                    } else {
                        v4 = true
                        v2["hello"] = (488 >= 65535)
                        v46 := GenS6{F0: 65535}
                        _ = v46
                    }
                    _ = fi37
                }
                for fi47 := 0; fi47 < 3; fi47++ {
                    _ = len(v1)
                    var v48 int = (fi47 >> ((1) & 63))
                    _ = fi47
                    _ = v48
                }
            } else {
                fmt.Println(v11, v4)
            }
            v49 := GenS7{F0: 4294967295, F1: v12, F2: -305}
            _ = v49
        }
        v11 -= 32767
        if ((v3 + -180) < (v3 >> ((616) & 63))) {
            if (1e10 < float32(float64(0.1))) {
                var v50 int16 = 42
                _ = v50
            } else {
                var v51 string = "foo bar"
                if (!(0 != uint(v12))) {
                    v11 = v11
                    switch 127 {
                    case -769:
                        fmt.Println(v3)
                    case 9223372036854775807:
                        if (!(("hello" + (v51 + v51)) <= "a")) {
                            var v52 int16 = 0
                            for fi53 := 0; fi53 < 20; fi53++ {
                                var v54 uint32 = 2147483647
                                v55 := GenS1{F0: v51, F1: genFn0(-2147483648), F2: -2147483648}
                                var v56 uint64 = 1
                                _ = fi53
                                _ = v54
                                _ = v55
                                _ = v56
                            }
                            _ = v52
                        } else {
                            var v57 uint64 = 671
                            _ = v57
                        }
                    }
                    var v58 int8 = -128
                    _ = v58
                } else {
                    v11 += v3
                    if true {
                        v51 = v51
                    }
                }
                if (-128 < -95) {
                    var v59 int32 = -347
                    _ = v59
                }
                _ = v51
            }
        }
    }
    v60 := GenS3{F0: true, F1: 42, F2: (v12 * 42)}
    v12 = v12
    var v61 float64 = 0.0
    var v62 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v60
    _ = v61
    _ = v62
}
