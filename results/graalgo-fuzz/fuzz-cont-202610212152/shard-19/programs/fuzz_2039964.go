package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int64; F2 int16 }
func genFn0() int {
    var v4 uint32 = 2
    var v5 uint16 = 32767
    var v6 int32 = 2147483647
    _ = v4
    _ = v5
    _ = v6
    return 262
}

func main() {
    var v1 uint16 = 753
    v2 := []int{-2147483648}
    v3 := []string{"hello"}
    v1 = 42
    var v7 int8 = 42
    if (-32768 == int16(v1)) {
        v1 += (v1 << ((4294967295) & 63))
        if len(v3) > 0 {
            v3[0] = ("" + "\n\t")
        }
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(v1)
            _ = fi9
        }
        _ = fi8
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        if (false || (fi10 > fi10)) {
            var v11 bool = ((^(v1)) > (v1 - 65535))
            switch "" {
            case "a":
                if (false || false) {
                    v12 := GenS1{F0: 0, F1: -43}
                    if v11 {
                        if v11 {
                            for fi13 := 0; fi13 < 4; fi13++ {
                                var v14 int = fi10
                                v14 = 0
                                switch 2147483647 {
                                case -1:
                                    v11 = (127 >= 128)
                                }
                                var v15 uint = 65535
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 17; fi16++ {
                                if len(v2) > 0 {
                                    v2[0] = (^(fi16))
                                }
                                fmt.Println(fi16, fi10, fi16)
                                _ = fi16
                            }
                            var v17 int16 = 254
                            _ = v17
                        } else {
                            v11 = true
                            var v18 uint8 = 2
                            _ = v18
                        }
                        v1 = (v1 >> ((0) & 63))
                    } else {
                        v19 := GenS2{F0: v1, F1: 2}
                        switch ("a" + "foo bar") {
                        case "\n\t":
                            var v20 uint = 1
                            _ = v20
                        }
                        _ = v19
                    }
                    _ = v12
                } else {
                    v1 -= (v1 | v1)
                }
            case "\n\t":
                v11 = v11
            case "hello":
                for fi21 := 0; fi21 < 3; fi21++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v22 bool = (v11 || (v1 <= v1))
                    var v23 float64 = (0.1 - float64(1.5))
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            default:
                var v24 uint = 4294967295
                _ = v24
            }
            v25 := GenS2{F0: v1, F1: 42}
            _ = v11
            _ = v25
        } else {
            if (42 >= 18446744073709551615) {
                v26 := GenS1{F0: -846, F1: 42}
                fmt.Println(v1)
                _ = v26
            } else {
                var v27 uint64 = 0
                v28 := GenS3{F0: 42}
                _ = v27
                _ = v28
            }
        }
        if (!false) {
            for fi29 := 0; fi29 < 9; fi29++ {
                v1 = (v1 & (v1 * v1))
                if (!true) {
                    fmt.Println(fi29, fi29, v7)
                }
                _ = fi29
            }
            var v30 float64 = (float64(float32((1e10 / (1.5 / 3.14159)))) / 0.1)
            if (0 <= 340) {
                v31 := GenS1{F0: -955, F1: 0}
                var v32 int32 = -2147483648
                _ = v31
                _ = v32
            } else {
                for fi33 := 0; fi33 < 17; fi33++ {
                    if true {
                        var v34 int8 = v7
                        if ("a" > "hello") {
                            if true {
                                if false {
                                    v34 -= v34
                                    v35 := GenS1{F0: 42, F1: 645}
                                    switch (fi33 >> ((65535) & 63)) {
                                    case -1:
                                        v3 = append(v3, "\n\t")
                                    case 200:
                                        fmt.Println(v7)
                                    default:
                                        fmt.Println(v34)
                                        var v36 uint32 = 1
                                        _ = v36
                                    }
                                    _ = v35
                                } else {
                                    var v37 bool = (float32(0.1) != 1.5)
                                    v38 := GenS4{F0: 18446744073709551615, F1: -1, F2: 110}
                                    switch "hello" {
                                    case "hello":
                                        v39 := GenS1{F0: 1, F1: -1}
                                        _ = v39
                                    }
                                    _ = v37
                                    _ = v38
                                }
                                v2 = append(v2, (fi10 / ((fi33) | 1)))
                                _ = len(v3)
                            } else {
                                var v40 int = fi10
                                _ = v40
                            }
                            switch (fi10 << ((459) & 63)) {
                            case -1000:
                                var v41 uint = 118
                                _ = v41
                            case -302:
                                fmt.Println(v1, v34)
                            }
                        }
                        _ = v34
                    }
                    if false {
                        v7 = 0
                    }
                    _ = fi33
                }
                v2 = append(v2, fi10)
            }
            _ = v30
        } else {
            var v42 int32 = 0
            if (!false) {
                var v43 int64 = 1
                _ = v43
            }
            _ = v42
        }
        var v44 uint16 = v1
        _ = fi10
        _ = v44
    }
    var v45 int32 = -2147483648
    v7 -= (v7 * (v7 + v7))
    v45 += v45
    for fi46 := 0; fi46 < 2; fi46++ {
        var v47 bool = (!(0 < 4294967295))
        _ = fi46
        _ = v47
    }
    var v48 string = "a"
    if (false || false) {
        v45 = -299
        v49 := GenS1{F0: 9223372036854775807, F1: -9223372036854775808}
        _ = v49
    } else {
        var v50 uint32 = 4294967295
        if (v50 == v50) {
            var v51 int = 254
            v45 = v45
            if true {
                _ = len(v3)
            } else {
                var v52 string = v48
                _ = v52
            }
            _ = v51
        }
        var v53 float64 = (0.0 - -2.25)
        _ = v50
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v48
    _ = v7
}
