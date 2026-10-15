package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 string }
func genFn0(p0 int16, p1 uint) string {
    p1 = (^(p1))
    var v5 float64 = (1e10 - 1.5)
    var v6 uint64 = 688
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint8 = 128
        var v9 uint8 = 0
        var v10 int64 = 9223372036854775807
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    p0 += 42
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return ""
}
func genFn1(p0 int8) int32 {
    p0 += 100
    var v11 uint64 = 167
    v12 := []int{}
    _ = p0
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 int = -300
    var v2 int = v1
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    if (true && false) {
        if false {
            if (435 <= 42) {
                if (1 == 1) {
                    if (536 > int16(v2)) {
                        if true {
                            var v13 int32 = 42
                            var v14 uint16 = 65535
                            for fi15 := 0; fi15 < 18; fi15++ {
                                if (128 != 163) {
                                    fmt.Println(v2)
                                    v16 := []int{-2147483648}
                                    v17 := v4["a"]
                                    _ = v16
                                    _ = v17
                                } else {
                                    fmt.Println(v1, v14, v2)
                                    switch "" {
                                    case "":
                                        var v18 int64 = -9223372036854775808
                                        _ = v18
                                    case "foo bar":
                                        var v19 int16 = 787
                                        _ = v19
                                    case "\n\t":
                                        if (("\n\t" + ("\n\t" + ("\n\t" + "hello"))) == "foo bar") {
                                            var v20 int32 = v13
                                            v3[359] = true
                                            if (true && (!(false && (!(-837 != -713))))) {
                                                v2 = (^(fi15))
                                                v20 -= v13
                                            } else {
                                                fmt.Println(fi15, v14, v13)
                                            }
                                            _ = v20
                                        }
                                    }
                                }
                                v13 -= v13
                                v1 -= 0
                                _ = fi15
                            }
                            _ = v13
                            _ = v14
                        } else {
                            switch (v2 * v1) {
                            case 739:
                                for fi21 := 0; fi21 < 18; fi21++ {
                                    var v22 int32 = genFn1(0)
                                    v23 := GenS1{F0: 73, F1: ""}
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                            case -559:
                                var v24 uint16 = 380
                                _ = v24
                            case 426:
                                v2 += (v1 >> ((42) & 63))
                            default:
                                v25 := []int{v2, v1, v2}
                                _ = v25
                            }
                            fmt.Println(v1, v1)
                        }
                    } else {
                        v2 += 164
                    }
                    var v26 string = "a"
                    _ = v26
                }
                var v27 int16 = 42
                _ = v27
            } else {
                for fi28 := 0; fi28 < 5; fi28++ {
                    if (!(2 < 42)) {
                        for fi29 := 0; fi29 < 7; fi29++ {
                            v30 := GenS2{F0: 42}
                            fmt.Println(fi29, fi28)
                            v31 := GenS1{F0: 18446744073709551615, F1: "\n\t"}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 4; fi32++ {
                            v33 := GenS2{F0: 42}
                            var v34 int = fi28
                            switch ("a" + "\n\t") {
                            case "\n\t":
                                v2 -= 42
                            case "foo bar":
                                var v35 uint8 = 2
                                _ = v35
                            }
                            v1 = 32768
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    }
                    v1 -= (v1 << ((550) & 63))
                    _ = fi28
                }
                if true {
                    var v36 uint16 = 1
                    var v37 uint16 = 42
                    v2 -= v1
                    _ = v36
                    _ = v37
                } else {
                    v38 := GenS2{F0: 76}
                    _ = v38
                }
            }
            v39 := GenS2{F0: 0}
            _ = v39
        }
        _ = len(v3)
        if (128 > 152) {
            v40 := GenS3{F0: 42}
            _ = v40
        }
    } else {
        v1 += v2
    }
    var v41 int = (v2 * (v1 + v2))
    v2 += (^(v2))
    var v42 int8 = 100
    var v43 int32 = -1
    var v44 string = (("" + "hello") + "hello")
    _ = len(v4)
    var v45 bool = true
    if ((1.5 / ((-2.25 * (1.5 * 3.14159)) - (1.5 + 3.14159))) > (1e10 - -2.25)) {
        for fi46 := 0; fi46 < 19; fi46++ {
            if (!((42 >= 2) || false)) {
                v2 += v2
                v1 = fi46
                if (0 != 4294967295) {
                    var v47 uint32 = 590
                    v48 := GenS4{F0: (v42 ^ (v42 >> ((uint(v47)) & 63))), F1: v44}
                    v47 += 81
                    _ = v47
                    _ = v48
                }
            }
            var v49 int = (-(fi46))
            _ = fi46
            _ = v49
        }
    } else {
        var v50 string = ("" + "")
        _ = v50
    }
    _ = len(v3)
    v42 -= v42
    var v51 uint = 255
    v2 -= (v1 + v2)
    if ((!true) && false) {
        var v52 int64 = -9223372036854775808
        _ = v52
    } else {
        fmt.Println(v1, v42, v42)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v51
}
