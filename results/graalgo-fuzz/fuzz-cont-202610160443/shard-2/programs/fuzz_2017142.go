package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 uint }
func genFn0(p0 uint8, p1 uint32) uint64 {
    var v4 uint16 = 1
    v5 := GenS1{F0: "a", F1: 32767, F2: 32767}
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS2{F0: 32767, F1: 809}
        if true {
            fmt.Println(fi6, v4)
            if (0 >= 480) {
                var v8 bool = (p0 > uint8(v4))
                _ = v8
            }
        } else {
            var v9 uint64 = 4294967295
            v10 := []int{(-(fi6)), -1, (-(fi6)), fi6}
            v4 -= 2
            _ = v10
            _ = v9
        }
        p0 += 90
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 int32) int16 {
    p0 += (^(p0))
    var v11 int32 = -2147483648
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := []int{fi12, fi12, fi12}
        var v14 string = "foo bar"
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: "\n\t", F1: 31, F2: 42}
    p0 = p0
    _ = p0
    _ = v11
    _ = v15
    return 944
}

func main() {
    var v1 bool = ((4294967295 >= 2) && (true || ((!(100 < 100)) && true)))
    var v2 int32 = -852
    v3 := []int{547, 128}
    v16 := GenS3{F0: 1000, F1: 128, F2: 42}
    var v17 int = 898
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 int32 = (v2 / ((v2) | 1))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch 128 {
        case 754:
            for fi20 := 0; fi20 < 14; fi20++ {
                switch (("foo bar" + "foo bar") + "\n\t") {
                case "hello":
                    if (178 > 201) {
                        if ("" > "foo bar") {
                            fmt.Println(v17)
                            v21 := GenS2{F0: -32768, F1: -484}
                            v1 = ((v19 & v19) == 2147483647)
                            if (!(true || v1)) {
                                var v22 int8 = 100
                                switch (("hello" + (("foo bar" + "foo bar") + ("a" + "foo bar"))) + ("\n\t" + ("a" + "\n\t"))) {
                                case "":
                                    var v23 uint32 = 0
                                    _ = v23
                                case "foo bar":
                                    var v24 int32 = -150
                                    _ = v24
                                }
                                _ = v22
                            }
                            _ = v21
                        } else {
                            var v25 uint8 = 255
                            var v26 float32 = (-2.25 + 0.0)
                            _ = v25
                            _ = v26
                        }
                        fmt.Println(v2)
                        v2 = (v19 % (((v19 * (v2 ^ 42))) | 1))
                    }
                case "\n\t":
                    var v27 int8 = -1
                    _ = v27
                }
                _ = fi20
            }
        case 2147483647:
            v2 -= v19
        case -2147483648:
            fmt.Println(v17, v19, fi18)
        }
        _ = fi18
        _ = v19
    }
    var v28 uint32 = 578
    if len(v3) > 0 {
        v3[0] = v17
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        if (!v1) {
            if len(v3) > 0 {
                v3[0] = fi29
            }
            switch "" {
            case "":
                v2 = -818
            case "foo bar":
                v30 := GenS4{F0: "foo bar", F1: 0}
                _ = v30
            case "\n\t":
                var v31 uint16 = 1
                _ = v31
            }
        } else {
            for fi32 := 0; fi32 < 19; fi32++ {
                if (2 > 1) {
                    v28 = v28
                    _ = len(v3)
                    if v1 {
                        if v1 {
                            var v33 uint = 0
                            var v34 int64 = 9223372036854775807
                            _ = v33
                            _ = v34
                        } else {
                            var v35 bool = (false || (v1 || v1))
                            v17 += fi32
                            switch 613 {
                            case 0:
                                var v36 float64 = 1.5
                                _ = v36
                            default:
                                v35 = v1
                            }
                            _ = v35
                        }
                        var v37 int32 = (v2 | -1)
                        _ = v37
                    } else {
                        var v38 uint32 = (v28 | (v28 + (v28 + (v28 & (v28 & 1)))))
                        if (970 == v38) {
                            v39 := GenS1{F0: "", F1: 1, F2: -1}
                            v40 := GenS5{F0: 4294967295}
                            fmt.Println(fi29)
                            _ = v39
                            _ = v40
                        } else {
                            v41 := GenS6{F0: 4294967295, F1: 42}
                            var v42 uint16 = 694
                            _ = v41
                            _ = v42
                        }
                        _ = v38
                    }
                }
                var v43 int64 = 42
                v28 -= (v28 - (v28 + (v28 & (^(v28)))))
                _ = fi32
                _ = v43
            }
        }
        _ = fi29
    }
    var v44 uint64 = uint64(v28)
    var v45 float64 = 0.1
    var v46 int = (v17 ^ 15)
    var v47 uint16 = 255
    if false {
        fmt.Println(v45)
        v48 := GenS4{F0: "\n\t", F1: 0}
        v28 += (v28 << ((18446744073709551615) & 63))
        _ = v48
    } else {
        var v49 float32 = 3.14159
        v46 += v46
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
