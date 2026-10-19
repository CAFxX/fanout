package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }

func main() {
    var v1 uint16 = 735
    var v2 bool = true
    var v3 uint64 = 2
    var v4 bool = v2
    var v5 uint32 = 65535
    v6 := GenS1{F0: v3, F1: v3, F2: 18446744073709551615}
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{(fi7 + (fi7 / ((fi7) | 1))), (^(fi7))}
        _ = fi7
        _ = v8
    }
    var v9 float64 = (float64(1.5) - 0.0)
    v10 := []int{}
    v10 = append(v10, -9223372036854775808)
    v11 := GenS1{F0: v3, F1: v3, F2: 65535}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if ((v4 && v2) || (!(0 >= 42))) {
        v4 = ((-1 > 42) && (!v2))
    } else {
        switch 9223372036854775807 {
        case 32767:
            var v12 string = "foo bar"
            _ = v12
        case -2147483648:
            fmt.Println(v3, v2, v4)
        default:
            fmt.Println(v4)
        }
        switch "a" {
        case "a":
            v13 := GenS2{F0: ("a" + "foo bar")}
            _ = v13
        case "foo bar":
            var v14 bool = (-1 <= 2147483647)
            _ = v14
        default:
            if v2 {
                fmt.Println(v9, v1)
                fmt.Println(v4)
                for fi15 := 0; fi15 < 14; fi15++ {
                    if (v3 > (v3 ^ uint64(v5))) {
                        if v2 {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            fmt.Println(v2)
                            v5 = v5
                        }
                        v4 = (128 <= (fi15 >> ((2) & 63)))
                    }
                    _ = fi15
                }
            }
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 string = "\n\t"
            if (792 == 255) {
                var v18 bool = true
                _ = v18
            } else {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            if len(v10) > 0 {
                v10[0] = 128
            }
            _ = fi16
            _ = v17
        }
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            v5 = (v5 / (((v5 % (((v5 >> ((255) & 63))) | 1))) | 1))
            v10 = append(v10, fi20)
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi20
        }
        if v4 {
            var v21 float64 = v9
            v9 = (v9 - v21)
            v21 += float64(0.1)
            _ = v21
        } else {
            var v22 bool = v2
            var v23 uint8 = 74
            _ = v22
            _ = v23
        }
        _ = fi19
    }
    v24 := GenS3{F0: ("" + (("" + "a") + ("hello" + "a")))}
    v9 = 0.1
    if ((1 > 0) && v2) {
        for fi25 := 0; fi25 < 9; fi25++ {
            if v2 {
                if v2 {
                    fmt.Println(v5, v9, v3)
                    v26 := GenS3{F0: ("hello" + "a")}
                    var v27 int32 = 2147483647
                    _ = v26
                    _ = v27
                } else {
                    if false {
                        v1 = v1
                        var v28 bool = v4
                        _ = v28
                    } else {
                        if (-1 > -1) {
                            if ((-2.25 >= (3.14159 - 0.0)) || v2) {
                                v5 = v5
                            }
                            var v29 int64 = 42
                            _ = v29
                        } else {
                            switch "foo bar" {
                            case "hello":
                                for fi30 := 0; fi30 < 3; fi30++ {
                                    switch (fi25 | (fi30 & 127)) {
                                    case 294:
                                        if (!(-1 != 19)) {
                                            var v31 uint64 = (v3 % (((v3 << ((255) & 63))) | 1))
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                v2 = (0 <= -1)
                                            }
                                            _ = v31
                                        } else {
                                            var v32 float64 = (v9 + ((v9 + v9) + v9))
                                            switch -1 {
                                            case 0:
                                                v32 += float64((3.14159 / 0.1))
                                            case 561:
                                                var v33 float32 = -2.25
                                                _ = v33
                                            }
                                            v10 = append(v10, (fi30 + (fi30 + fi30)))
                                            if (-1 < 42) {
                                                if false {
                                                    v5 = (v5 - 2147483647)
                                                    var v34 int8 = 1
                                                    _ = v34
                                                }
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                fmt.Println(fi30, v5)
                                            } else {
                                                v5 = v5
                                                v35 := GenS3{F0: ""}
                                                fmt.Println(fi30, fi25)
                                                _ = v35
                                            }
                                            v4 = v2
                                            _ = v32
                                        }
                                    case -9223372036854775808:
                                        var v36 string = ""
                                        _ = v36
                                    case -1:
                                        var v37 uint64 = 708
                                        _ = v37
                                    }
                                    if (1 == 496) {
                                        fmt.Println(v9, v2, v5)
                                        var v38 bool = (("" + (("hello" + ("" + "")) + "")) != "foo bar")
                                        v3 += v3
                                        _ = v38
                                    } else {
                                        var v39 float64 = (float64(1.5) * v9)
                                        var v40 int32 = -1
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        v40 -= int32(v3)
                                        _ = v39
                                        _ = v40
                                    }
                                    _ = fi30
                                }
                            case "":
                                if true {
                                    v10 = append(v10, (fi25 / ((fi25) | 1)))
                                }
                            case "foo bar":
                                var v41 bool = (((("\n\t" + ("" + "")) + "") + "") <= "a")
                                _ = v41
                            }
                            v10 = append(v10, (fi25 + (fi25 + int(v1))))
                            switch fi25 {
                            case 0:
                                var v42 uint16 = (v1 << ((2) & 63))
                                _ = v42
                            case 93:
                                var v43 uint = 0
                                _ = v43
                            case -139:
                                _ = len(v10)
                            }
                        }
                    }
                    v9 = v9
                }
            } else {
                v9 = (float64(((3.14159 + 0.0) - 3.14159)) * float64(float32(v9)))
                v44 := GenS2{F0: ""}
                _ = v44
            }
            var v45 uint = 493
            fmt.Println(fi25, v2, v2)
            _ = fi25
            _ = v45
        }
        v46 := GenS2{F0: "a"}
        _ = v46
    } else {
        var v47 int32 = -495
        var v48 int = 0
        _ = v47
        _ = v48
    }
    v5 = (v5 * v5)
    v1 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
