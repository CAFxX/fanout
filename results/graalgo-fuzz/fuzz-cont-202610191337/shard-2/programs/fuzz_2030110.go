package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
func genFn0(p0 uint16, p1 int64, p2 bool) int64 {
    fmt.Println(p1)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (^(p1))
}
func genFn1() string {
    var v8 uint64 = 289
    var v9 int = 4294967295
    v10 := []int{(^(v9)), v9, (v9 + (^(v9))), (v9 % (((-(v9))) | 1))}
    if true {
        for fi11 := 0; fi11 < 5; fi11++ {
            if false {
                fmt.Println(fi11, v9)
            } else {
                var v12 uint64 = (v8 & v8)
                for fi13 := 0; fi13 < 4; fi13++ {
                    if len(v10) > 0 {
                        v10[0] = (fi13 << ((2) & 63))
                    }
                    fmt.Println(v12, v8)
                    _ = fi13
                }
                fmt.Println(fi11, fi11)
                _ = v12
            }
            v8 += v8
            fmt.Println(v9, fi11, v8)
            _ = fi11
        }
    } else {
        var v14 string = ""
        var v15 float32 = ((0.0 * (1e10 * 0.1)) - (0.0 - 1e10))
        _ = len(v10)
        _ = v14
        _ = v15
    }
    _ = v10
    _ = v8
    _ = v9
    return ("\n\t" + "")
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (!((651 > -1) || true))
    v3 := []int{}
    v4 := []uint32{65535}
    var v5 int = 967
    var v6 string = ""
    fmt.Println(v6, v2)
    v16 := GenS1{F0: 128}
    v5 += v5
    if v2 {
        fmt.Println(v6)
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 uint32 = 42
                v20 := GenS2{F0: v6, F1: 0, F2: 0}
                v21 := v1[(^(v19))]
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 uint8 = 0
            _ = fi17
            _ = v22
        }
        delete(v1, 2147483647)
        var v23 uint8 = 1
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        if (false && (v2 || (-128 >= 0))) {
            fmt.Println(v5, v2, v2)
        }
        if v2 {
            for fi25 := 0; fi25 < 8; fi25++ {
                var v26 int8 = 1
                _ = fi25
                _ = v26
            }
        } else {
            if (!v2) {
                v27 := GenS3{F0: 1, F1: 32767}
                if (!v2) {
                    var v28 uint64 = 0
                    _ = v28
                }
                _ = v27
            }
            if (9223372036854775807 < 1) {
                v5 = 32768
            } else {
                _ = len(v4)
                for fi29 := 0; fi29 < 1; fi29++ {
                    var v30 uint16 = 1
                    _ = fi29
                    _ = v30
                }
            }
            if (((v6 + v6) + "hello") > "foo bar") {
                var v31 string = ("hello" + v6)
                v6 = v31
                var v32 bool = v2
                _ = v31
                _ = v32
            } else {
                var v33 bool = false
                fmt.Println(fi24, fi24, v6)
                _ = v33
            }
        }
        _ = fi24
    }
    v2 = false
    switch v6 {
    case "a":
        fmt.Println(v5)
    }
    switch "" {
    case "":
        var v34 int8 = -1
        _ = v34
    default:
        v6 = "a"
    }
    var v35 int16 = 1
    for fi36 := 0; fi36 < 2; fi36++ {
        delete(v1, 2147483647)
        fmt.Println(v35)
        _ = fi36
    }
    var v37 uint64 = 549
    if v2 {
        v38 := GenS4{F0: 509}
        if (176 != 2) {
            switch v5 {
            case 651:
                switch v6 {
                case "a":
                    if v2 {
                        if ("\n\t" < ((v6 + "a") + v6)) {
                            if v2 {
                                for fi39 := 0; fi39 < 6; fi39++ {
                                    var v40 uint8 = 2
                                    v40 = 97
                                    _ = fi39
                                    _ = v40
                                }
                                for fi41 := 0; fi41 < 20; fi41++ {
                                    for fi42 := 0; fi42 < 18; fi42++ {
                                        var v43 int32 = 2147483647
                                        v3 = append(v3, 128)
                                        var v44 uint64 = 42
                                        _ = fi42
                                        _ = v43
                                        _ = v44
                                    }
                                    if v2 {
                                        var v45 int8 = 127
                                        _ = v45
                                    }
                                    var v46 bool = v2
                                    _ = fi41
                                    _ = v46
                                }
                            } else {
                                if false {
                                    v6 = ""
                                }
                            }
                            v47 := GenS1{F0: 1}
                            _ = v47
                        }
                        v35 -= v35
                    } else {
                        v48 := GenS4{F0: v35}
                        _ = v48
                    }
                default:
                    v49 := GenS3{F0: 283, F1: 1}
                    _ = v49
                }
            default:
                v37 = v37
            }
            fmt.Println(v35, v5)
        }
        _ = v38
    } else {
        for fi50 := 0; fi50 < 20; fi50++ {
            var v51 int = (v5 * (fi50 + fi50))
            switch (("foo bar" + ("hello" + "hello")) + ((v6 + (("a" + v6) + v6)) + "a")) {
            case "a":
                var v52 int = -2147483648
                _ = v52
            case "foo bar":
                v2 = false
            case "hello":
                var v53 uint8 = 1
                _ = v53
            default:
                v54 := GenS1{F0: 2}
                _ = v54
            }
            for fi55 := 0; fi55 < 6; fi55++ {
                switch "" {
                case "foo bar":
                    var v56 int8 = -128
                    _ = v56
                case "a":
                    var v57 uint8 = 9
                    _ = v57
                default:
                    v58 := GenS5{F0: v6}
                    _ = v58
                }
                var v59 float64 = 0.0
                if (false && v2) {
                    var v60 int16 = v35
                    if v2 {
                        v37 = v37
                        if v2 {
                            if (-32768 < -1) {
                                fmt.Println(v60, v6, v35)
                                v37 -= 636
                            } else {
                                var v61 uint = 699
                                _ = v61
                            }
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                        } else {
                            v62 := GenS2{F0: "a", F1: (-(v35)), F2: (v37 << ((65535) & 63))}
                            _ = v62
                        }
                    } else {
                        v1[125] = v6
                        v63 := GenS3{F0: (^(v37)), F1: 1}
                        _ = v63
                    }
                    _ = v60
                }
                _ = fi55
                _ = v59
            }
            _ = fi50
            _ = v51
        }
        var v64 uint64 = 42
        v65 := GenS1{F0: 127}
        _ = v64
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
