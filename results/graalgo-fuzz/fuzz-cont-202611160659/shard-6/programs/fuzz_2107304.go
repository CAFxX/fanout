package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }

func main() {
    var v1 int32 = 1
    v2 := []int{-40, 127, 346, -2147483648}
    var v3 float32 = -2.25
    var v4 uint64 = 0
    var v5 int32 = v1
    if (!false) {
        var v6 float32 = v3
        _ = v6
    } else {
        var v7 int16 = -688
        _ = v7
    }
    switch "hello" {
    case "a":
        var v8 uint8 = 81
        _ = v8
    }
    v5 += v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        var v9 uint16 = 32767
        var v10 uint16 = v9
        if (1.5 < 1e10) {
            if (-890 == -9223372036854775808) {
                var v11 uint16 = (v9 % ((32767) | 1))
                _ = v11
            } else {
                v9 += v9
            }
        }
        _ = v10
        _ = v9
    } else {
        if ((795 != 4294967295) || (!((1000 <= 730) && true))) {
            fmt.Println(v1, v1, v1)
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 float32 = v3
                v14 := GenS1{F0: 42, F1: 42, F2: "\n\t"}
                _ = fi12
                _ = v13
                _ = v14
            }
        } else {
            var v15 uint64 = v4
            if true {
                v2 = append(v2, 9223372036854775807)
                var v16 uint64 = (^(v4))
                if false {
                    if false {
                        if len(v2) > 0 {
                            v2[0] = -617
                        }
                        v4 -= (v4 * (v15 / (((^(v4))) | 1)))
                    } else {
                        fmt.Println(v4)
                        if (127 > 5) {
                            for fi17 := 0; fi17 < 13; fi17++ {
                                _ = len(v2)
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    var v19 int64 = -9223372036854775808
                                    if false {
                                        var v20 string = "foo bar"
                                        fmt.Println(fi17, fi18)
                                        var v21 uint32 = 0
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if (true && true) {
                                            fmt.Println(v16)
                                        }
                                    }
                                    _ = fi18
                                    _ = v19
                                }
                                _ = fi17
                            }
                            for fi22 := 0; fi22 < 3; fi22++ {
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    var v24 uint32 = 0
                                    _ = fi23
                                    _ = v24
                                }
                                for fi25 := 0; fi25 < 9; fi25++ {
                                    var v26 uint = 339
                                    _ = fi25
                                    _ = v26
                                }
                                _ = fi22
                            }
                            if false {
                                fmt.Println(v5, v5)
                                v4 = v16
                                var v27 bool = false
                                _ = v27
                            } else {
                                var v28 float64 = float64(float32(0.0))
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v28
                            }
                        }
                        var v29 bool = (!(71 <= 65535))
                        _ = v29
                    }
                    fmt.Println(v15, v4)
                }
                _ = v16
            }
            for fi30 := 0; fi30 < 16; fi30++ {
                for fi31 := 0; fi31 < 2; fi31++ {
                    v3 = (v3 / v3)
                    if len(v2) > 0 {
                        v2[0] = fi30
                    }
                    _ = fi31
                }
                v32 := GenS1{F0: 100, F1: 811, F2: "hello"}
                _ = fi30
                _ = v32
            }
            _ = v15
        }
        var v33 int32 = v5
        _ = v33
    }
    if len(v2) > 0 {
        v2[0] = 303
    }
    fmt.Println(v5, v4)
    v34 := GenS1{F0: 42, F1: 1, F2: "hello"}
    v5 = (v5 + 1)
    if false {
        var v35 int16 = 1
        var v36 uint = 4294967295
        _ = v35
        _ = v36
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch "foo bar" {
    case "a":
        for fi37 := 0; fi37 < 5; fi37++ {
            var v38 uint16 = 346
            v3 += 1.5
            v39 := GenS1{F0: 42, F1: 0, F2: "foo bar"}
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    switch 128 {
    case 128:
        if (true || false) {
            var v40 float32 = (v3 - float32(((3.14159 * (((-2.25 - 1.5) / float64(float32(3.14159))) - float64((v3 + float32(0.0))))) / 1.5)))
            var v41 int64 = 9223372036854775807
            var v42 int64 = -83
            _ = v40
            _ = v41
            _ = v42
        } else {
            v4 = v4
        }
    case 9223372036854775807:
        if (v4 < 480) {
            if (1 < 71) {
                if false {
                    var v43 bool = ((v5 == 0) || false)
                    _ = v43
                } else {
                    var v44 int16 = 42
                    var v45 uint32 = 466
                    _ = v44
                    _ = v45
                }
                v1 = (v1 - 0)
            }
            var v46 int8 = -18
            v47 := GenS1{F0: v46, F1: 18446744073709551615, F2: ("foo bar" + "a")}
            _ = v46
            _ = v47
        }
    case -2147483648:
        v5 += v5
    default:
        v4 = v4
    }
    switch 42 {
    case 9223372036854775807:
        var v48 bool = true
        _ = v48
    }
    switch "a" {
    case "":
        if true {
            v3 += 0.0
        }
    case "a":
        var v49 float64 = 1.5
        _ = v49
    default:
        var v50 string = (("hello" + ("" + "foo bar")) + "hello")
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
