package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint32 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint32 }
func genFn0() int {
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 bool = (1 != 42)
            _ = fi8
            _ = v9
        }
        if false {
            if false {
                v10 := GenS1{F0: 126, F1: 0, F2: 2}
                var v11 int32 = -1
                var v12 uint64 = 1
                _ = v10
                _ = v11
                _ = v12
            }
        } else {
            switch "" {
            case "a":
                v13 := []int{(fi7 >> ((42) & 63)), 2147483647}
                _ = v13
            case "\n\t":
                var v14 string = (("hello" + "hello") + "\n\t")
                _ = v14
            default:
                fmt.Println(fi7)
            }
        }
        var v15 int16 = 1
        _ = fi7
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 int = (fi16 + 2147483647)
            var v19 string = ("" + "hello")
            if (1 == 25) {
                v18 -= v18
                v20 := []int{}
                _ = v20
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 int64 = 697
            _ = fi21
            _ = v22
        }
        var v23 uint64 = 0
        _ = fi16
        _ = v23
    }
    return 127
}

func main() {
    var v1 uint8 = 1
    var v2 float32 = 3.14159
    var v3 uint = 65535
    var v4 uint16 = 32767
    var v5 uint8 = (^(v1))
    var v6 int16 = -32768
    v6 -= v6
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS2{F0: "\n\t", F1: (v3 % ((uint(v5)) | 1))}
        _ = fi24
        _ = v25
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v3)
    case "":
        if false {
            v26 := GenS1{F0: 2, F1: 12, F2: uint32(v1)}
            v3 += (v3 << ((1) & 63))
            for fi27 := 0; fi27 < 1; fi27++ {
                v28 := GenS3{F0: uint32(v3), F1: -9223372036854775808}
                var v29 int64 = 1
                var v30 int16 = (v6 * v6)
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            _ = v26
        } else {
            var v31 uint32 = 0
            _ = v31
        }
    }
    var v32 int32 = 0
    v33 := []int{genFn0(), 127, 42, 42}
    v5 = v5
    v34 := GenS1{F0: (v3 & uint(v4)), F1: int8(v3), F2: 2}
    fmt.Println(v5, v5)
    var v35 uint32 = 1
    if (("a" + "hello") > "foo bar") {
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        for fi36 := 0; fi36 < 4; fi36++ {
            _ = len(v33)
            _ = fi36
        }
        var v37 int32 = (v32 >> ((42) & 63))
        _ = v37
    }
    var v38 int16 = v6
    if false {
        v39 := GenS4{F0: (v32 - v32), F1: -1}
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 int32 = (v32 % (((v32 / (((v32 % (((v32 & (^(v32)))) | 1))) | 1))) | 1))
            switch "hello" {
            case "a":
                v41 = v41
            case "hello":
                v33 = append(v33, (fi40 >> ((v3) & 63)))
            default:
                switch ("hello" + "") {
                case "hello":
                    for fi42 := 0; fi42 < 5; fi42++ {
                        var v43 int32 = v41
                        v35 += v35
                        if true {
                            v44 := GenS2{F0: "a", F1: (^(v3))}
                            _ = v44
                        } else {
                            v45 := GenS1{F0: v3, F1: 42, F2: (^(v35))}
                            _ = len(v33)
                            var v46 int32 = 678
                            var v47 bool = (1e10 < ((0.1 - -2.25) + 0.1))
                            if ((0 < 1) && v47) {
                                if true {
                                    v38 = v6
                                    var v48 float32 = (v2 * v2)
                                    _ = v48
                                } else {
                                    var v49 int = (fi40 & fi40)
                                    var v50 int8 = -8
                                    _ = v49
                                    _ = v50
                                }
                                var v51 int16 = 32767
                                var v52 string = "a"
                                _ = v51
                                _ = v52
                            } else {
                                v32 += 42
                                if v47 {
                                    var v53 string = "hello"
                                    _ = v53
                                } else {
                                    v38 += v6
                                    if (-176 != int64(v3)) {
                                        if len(v33) > 0 {
                                            v33[0] = (fi40 << (((v3 * uint(v4))) & 63))
                                        }
                                        fmt.Println(v2)
                                        _ = len(v33)
                                        v54 := GenS5{F0: v35}
                                        _ = v54
                                    }
                                }
                                var v55 int8 = 42
                                _ = v55
                            }
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        _ = fi42
                        _ = v43
                    }
                default:
                    var v56 uint8 = v1
                    _ = v56
                }
            }
            _ = fi40
            _ = v41
        }
        _ = v39
    }
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    fmt.Println(v2, v4, v5)
    var v57 uint8 = (v1 / ((v1) | 1))
    var v58 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
