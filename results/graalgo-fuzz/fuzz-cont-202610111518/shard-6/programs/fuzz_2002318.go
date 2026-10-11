package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 string }

func main() {
    var v1 int16 = 739
    v2 := make(map[string]bool)
    var v3 float32 = (((3.14159 + 1.5) + 1.5) * 1e10)
    v3 += v3
    v3 += v3
    var v4 uint8 = 195
    v5 := GenS1{F0: 908, F1: 255, F2: 0}
    var v6 int16 = (v1 >> ((18446744073709551615) & 63))
    if (1 != 42) {
        for fi7 := 0; fi7 < 4; fi7++ {
            if true {
                if true {
                    var v8 int64 = 125
                    v9 := []int{fi7, 335}
                    v9 = append(v9, (fi7 ^ int(v4)))
                    _ = v8
                    _ = v9
                } else {
                    v10 := GenS1{F0: 32767, F1: 2, F2: 1}
                    _ = v10
                }
            } else {
                v6 = -1
                if (((v4 % ((v4) | 1)) > v4) && false) {
                    var v11 int32 = 42
                    _ = v11
                } else {
                    switch 2147483647 {
                    case -94:
                        var v12 float32 = float32((1e10 + 3.14159))
                        _ = v12
                    case 32767:
                        v6 += v6
                    case 140:
                        v13 := GenS1{F0: 504, F1: 269, F2: 42}
                        _ = v13
                    default:
                        if true {
                            v1 = -1
                            v14 := []int{}
                            _ = v14
                        }
                    }
                }
                fmt.Println(v3, v3)
            }
            v15 := []int{}
            _ = fi7
            _ = v15
        }
        v16 := []int{-2147483648}
        _ = v16
    } else {
        switch (("\n\t" + "a") + ("foo bar" + "")) {
        case "hello":
            switch "hello" {
            case "\n\t":
                fmt.Println(v1, v4, v6)
            default:
                v4 = v4
            }
        case "\n\t":
            if (false && true) {
                if (false || false) {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        if false {
                            for fi18 := 0; fi18 < 17; fi18++ {
                                fmt.Println(v1, v4, fi17)
                                _ = fi18
                            }
                            v19 := []int{(fi17 | 673), fi17, fi17, (fi17 & fi17)}
                            if (v6 <= v1) {
                                fmt.Println(v1, v4)
                                v4 += v4
                                v19 = append(v19, fi17)
                            } else {
                                var v20 uint32 = uint32(v4)
                                _ = v20
                            }
                            _ = v19
                        } else {
                            var v21 float64 = (1e10 - 0.1)
                            _ = v21
                        }
                        v4 -= v4
                        _ = fi17
                    }
                } else {
                    if ("a" < "") {
                        for fi22 := 0; fi22 < 10; fi22++ {
                            var v23 int = 9223372036854775807
                            v1 -= 42
                            v24 := []int{fi22, (fi22 + (v23 / ((-181) | 1))), v23, 127}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        fmt.Println(v6, v3)
                        for fi25 := 0; fi25 < 1; fi25++ {
                            v26 := GenS2{F0: 1, F1: fi25}
                            _ = fi25
                            _ = v26
                        }
                    } else {
                        v1 += (v1 / (((v6 - (v6 & v1))) | 1))
                    }
                }
                var v27 int64 = 0
                v28 := GenS1{F0: 255, F1: 1, F2: 104}
                _ = v27
                _ = v28
            } else {
                v29 := []int{}
                var v30 string = ""
                _ = v29
                _ = v30
            }
        case "a":
            _ = len(v2)
        default:
            var v31 uint64 = 842
            _ = v31
        }
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 uint64 = 0
            _ = fi32
            _ = v33
        }
        switch 9223372036854775807 {
        case 130:
            v34 := v2["a"]
            _ = v34
        default:
            v35 := []int{127, -1, 0, 55}
            _ = v35
        }
    }
    if false {
        fmt.Println(v6)
        _ = len(v2)
        var v36 int64 = -1
        _ = v36
    } else {
        if false {
            v1 = (v1 / ((1) | 1))
            v37 := GenS2{F0: 81, F1: 590}
            v38 := GenS3{F0: "foo bar"}
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 1; fi39++ {
            v40 := GenS2{F0: 531, F1: (^(fi39))}
            switch int(v1) {
            case 26:
                var v41 string = ("foo bar" + "hello")
                _ = v41
            case 127:
                v4 -= (v4 + v4)
            case 32768:
                var v42 uint32 = 708
                _ = v42
            default:
                var v43 bool = (2 < 65535)
                _ = v43
            }
            _ = fi39
            _ = v40
        }
    }
    var v44 uint = 1
    v45 := []int{32768, -320, 9223372036854775807, 32768}
    v46 := GenS3{F0: ("foo bar" + ("\n\t" + "\n\t"))}
    if ((435 <= 428) && false) {
        v47 := v2["foo bar"]
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
