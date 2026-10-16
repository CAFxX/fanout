package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int16; F2 int }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint; F1 string; F2 uint }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string; F1 int16; F2 int16 }
func genFn0(p0 bool) int {
    p0 = ((!true) || false)
    fmt.Println(p0)
    p0 = true
    var v5 float32 = float32(1.5)
    _ = p0
    _ = v5
    return 9223372036854775807
}
func genFn1(p0 int) uint8 {
    if false {
        var v6 int32 = -1
        switch "hello" {
        case "foo bar":
            v7 := GenS1{F0: -128, F1: 1000, F2: ((v6 / ((163) | 1)) < 275)}
            _ = v7
        case "":
            for fi8 := 0; fi8 < 10; fi8++ {
                p0 = 4294967295
                switch (fi8 - genFn0((!((2 > 4294967295) && false)))) {
                case -253:
                    switch ((((("\n\t" + "") + "") + "a") + "hello") + (("a" + "\n\t") + "\n\t")) {
                    case "foo bar":
                        if true {
                            v6 -= (v6 | 2147483647)
                        } else {
                            if false {
                                var v9 int8 = -39
                                v9 = (v9 ^ v9)
                                v10 := GenS2{F0: (v9 - 100)}
                                _ = v10
                                _ = v9
                            } else {
                                var v11 float32 = float32(0.1)
                                for fi12 := 0; fi12 < 12; fi12++ {
                                    var v13 float64 = 1.5
                                    var v14 int8 = -128
                                    _ = fi12
                                    _ = v13
                                    _ = v14
                                }
                                var v15 int32 = v6
                                _ = v11
                                _ = v15
                            }
                            var v16 int32 = (v6 | (-(v6)))
                            _ = v16
                        }
                    case "hello":
                        if true {
                            var v17 bool = ("a" == "a")
                            _ = v17
                        } else {
                            var v18 float32 = (1.5 - 0.1)
                            v18 = (v18 - ((1.5 * 1e10) + ((v18 * float32((1.5 + 3.14159))) * v18)))
                            _ = v18
                        }
                    }
                case -920:
                    if ((1.5 > 0.0) || false) {
                        if false {
                            v6 = v6
                            fmt.Println(fi8)
                        } else {
                            v19 := GenS2{F0: 0}
                            var v20 bool = (1 != 557)
                            _ = v19
                            _ = v20
                        }
                    } else {
                        v21 := GenS2{F0: -128}
                        fmt.Println(p0, fi8, p0)
                        _ = v21
                    }
                case -9223372036854775808:
                    var v22 uint64 = 0
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 uint8 = 255
                        if (1.5 != 0.0) {
                            var v25 uint64 = 341
                            v26 := []int{(fi8 >> ((18446744073709551615) & 63))}
                            _ = v25
                            _ = v26
                        }
                        _ = fi23
                        _ = v24
                    }
                }
                _ = fi8
            }
        default:
            var v27 uint64 = 985
            _ = v27
        }
        v28 := []int{(p0 / (((p0 * -140)) | 1))}
        _ = v28
        _ = v6
    } else {
        var v29 uint8 = 127
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        fmt.Println(p0, p0)
        _ = fi30
    }
    switch 9223372036854775807 {
    case -1:
        var v31 int64 = -895
        _ = v31
    case 1:
        var v32 int32 = 0
        _ = v32
    case 32767:
        for fi33 := 0; fi33 < 17; fi33++ {
            v34 := []int{fi33, p0, genFn0(false)}
            _ = fi33
            _ = v34
        }
    }
    var v35 string = "hello"
    for fi36 := 0; fi36 < 11; fi36++ {
        if false {
            var v37 int64 = 1
            v38 := GenS3{F0: v35}
            v37 = (v37 >> ((2) & 63))
            _ = v37
            _ = v38
        } else {
            v39 := GenS4{F0: fi36, F1: 147, F2: (p0 % (((fi36 ^ p0)) | 1))}
            fmt.Println(fi36, p0)
            _ = v39
        }
        _ = fi36
    }
    _ = p0
    _ = v35
    return 127
}

func main() {
    var v1 int32 = 1
    v2 := []uint32{}
    v3 := make(map[int]int)
    var v4 int64 = 42
    delete(v3, 42)
    v40 := GenS5{F0: 1, F1: false}
    delete(v3, 4294967295)
    for fi41 := 0; fi41 < 8; fi41++ {
        v42 := GenS3{F0: "a"}
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 20; fi43++ {
        v44 := v3[fi43]
        _ = len(v3)
        if false {
            v45 := GenS6{F0: 190, F1: "\n\t", F2: 255}
            _ = v45
        }
        _ = fi43
        _ = v44
    }
    v2 = append(v2, 2147483647)
    v46 := GenS7{F0: true, F1: -126}
    for fi47 := 0; fi47 < 15; fi47++ {
        if (594 != 42) {
            var v48 uint16 = 32767
            if ((((1.5 + 1e10) + (3.14159 - float32((1.5 * ((0.1 + 0.0) - ((((-2.25 + ((float64((1.5 + float32(0.0))) * float64(0.1)) / 1.0)) / (0.1 / float64(3.14159))) / 1e10) / 1e10)))))) / 1.0) < 0.1) {
                v48 = v48
                for fi49 := 0; fi49 < 10; fi49++ {
                    v1 = v1
                    var v50 uint64 = 4294967295
                    v51 := GenS6{F0: 212, F1: ("hello" + "a"), F2: 255}
                    _ = fi49
                    _ = v50
                    _ = v51
                }
            } else {
                switch (("\n\t" + "") + "hello") {
                case "":
                    v48 += v48
                case "foo bar":
                    fmt.Println(v1)
                default:
                    switch "foo bar" {
                    case "hello":
                        if true {
                            var v52 uint8 = 1
                            v53 := GenS8{F0: 1}
                            _ = v52
                            _ = v53
                        } else {
                            if false {
                                var v54 int = (fi47 << ((11) & 63))
                                v4 += (v4 | v4)
                                if ((0.1 * 1.5) < 1e10) {
                                    fmt.Println(fi47, v4, v1)
                                    if ((^(v48)) >= (^(v48))) {
                                        v4 = v4
                                        switch (("a" + ((("a" + ("hello" + "a")) + "a") + "a")) + "hello") {
                                        case "hello":
                                            v55 := GenS2{F0: 0}
                                            _ = v55
                                        }
                                        if (false || true) {
                                            var v56 int8 = 0
                                            v57 := GenS7{F0: true, F1: -9223372036854775808}
                                            _ = v56
                                            _ = v57
                                        } else {
                                            switch ("hello" + ("\n\t" + "foo bar")) {
                                            case "a":
                                                v58 := GenS8{F0: 32}
                                                _ = v58
                                            }
                                        }
                                    }
                                    v1 = v1
                                }
                                _ = v54
                            } else {
                                var v59 uint16 = v48
                                v60 := GenS4{F0: (-(fi47)), F1: -32768, F2: genFn0(false)}
                                _ = v59
                                _ = v60
                            }
                            if (!((0 < -128) && true)) {
                                for fi61 := 0; fi61 < 1; fi61++ {
                                    var v62 uint = 2
                                    _ = fi61
                                    _ = v62
                                }
                                v63 := GenS9{F0: ("\n\t" + ("\n\t" + ("hello" + "foo bar"))), F1: -1, F2: 0}
                                var v64 int64 = (v4 ^ int64(v1))
                                _ = v63
                                _ = v64
                            }
                            fmt.Println(v48, v4)
                        }
                    case "":
                        switch "" {
                        case "":
                            for fi65 := 0; fi65 < 15; fi65++ {
                                delete(v3, fi47)
                                v66 := GenS9{F0: "hello", F1: 42, F2: 32767}
                                _ = len(v2)
                                _ = fi65
                                _ = v66
                            }
                        case "a":
                            if false {
                                v67 := GenS9{F0: ("hello" + "\n\t"), F1: -20, F2: 32767}
                                var v68 int32 = (v1 + -1)
                                _ = v67
                                _ = v68
                            }
                        }
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = v48
        } else {
            fmt.Println(fi47, v4)
        }
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v46
}
