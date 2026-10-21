package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int8 }
func genFn0() uint8 {
    if (!(4294967295 != 0)) {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 int32 = 42
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 float64 = float64((1e10 - (1.5 * 0.0)))
                var v10 uint32 = 0
                var v11 int8 = 42
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS1{F0: -966}
                if ((!false) && true) {
                    if (!((true && true) || (false && (1 >= 42)))) {
                        switch 0 {
                        case -9223372036854775808:
                            v7 -= (^(v7))
                        case 2147483647:
                            v14 := GenS1{F0: 0}
                            _ = v14
                        case -1:
                            v15 := GenS1{F0: 1}
                            if true {
                                v7 += v7
                                fmt.Println(fi6, fi6, fi6)
                                var v16 uint = 4294967295
                                _ = v16
                            }
                            switch fi6 {
                            case 128:
                                if true {
                                    v7 = (v7 / (((v7 | v7)) | 1))
                                    v17 := []int{fi12}
                                    _ = v17
                                }
                            default:
                                v7 = (v7 ^ 42)
                                v18 := []int{}
                                _ = v18
                            }
                            var v19 string = "\n\t"
                            switch 2147483647 {
                            case 32768:
                                v19 = v19
                            case -864:
                                var v20 bool = false
                                _ = v20
                            }
                            _ = v15
                            _ = v19
                        }
                        v21 := []int{fi12, 491, (fi6 >> ((886) & 63))}
                        _ = v21
                    } else {
                        var v22 float64 = 1.5
                        if (((!(263 >= 2)) || (false && false)) && (127 < -13)) {
                            v22 = (v22 + v22)
                        }
                        v7 += v7
                        _ = v22
                    }
                } else {
                    var v23 float32 = 1e10
                    v7 = 0
                    switch "foo bar" {
                    case "hello":
                        if true {
                            v23 += (-2.25 / (v23 - v23))
                        } else {
                            v24 := GenS1{F0: -478}
                            _ = v24
                        }
                    }
                    _ = v23
                }
                v7 = (v7 + 484)
                v7 += -780
                var v25 int64 = 0
                var v26 uint8 = 255
                _ = fi12
                _ = v13
                _ = v25
                _ = v26
            }
            _ = fi6
            _ = v7
        }
        if (true && true) {
            v27 := []int{112, 1}
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 uint16 = 581
                _ = fi28
                _ = v29
            }
            _ = v27
        } else {
            v30 := GenS1{F0: 720}
            for fi31 := 0; fi31 < 3; fi31++ {
                fmt.Println(fi31)
                for fi32 := 0; fi32 < 3; fi32++ {
                    v33 := GenS1{F0: 0}
                    _ = fi32
                    _ = v33
                }
                var v34 int64 = 1
                _ = fi31
                _ = v34
            }
            var v35 uint16 = 32767
            _ = v30
            _ = v35
        }
    }
    v36 := []int{}
    var v37 string = ""
    _ = v36
    _ = v37
    return 127
}
func genFn1(p0 uint, p1 string, p2 int64) string {
    v38 := []int{-2147483648, -366}
    v39 := GenS2{F0: -32768, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    _ = v39
    return p1
}

func main() {
    var v1 int32 = -712
    var v2 uint64 = 0
    v3 := []uint32{315}
    v4 := []string{("\n\t" + "foo bar")}
    v5 := make(map[uint32]uint32)
    v1 = (v1 % (((v1 * (v1 * -2147483648))) | 1))
    v2 += 889
    _ = len(v3)
    v2 = (v2 + (^(v2)))
    var v40 int16 = 0
    v41 := GenS3{F0: 42, F1: 453, F2: -20}
    if (v1 < v1) {
        if len(v3) > 0 {
            v3[0] = uint32(v2)
        }
        if (2147483648 >= 819) {
            var v42 string = "hello"
            v5[42] = 932
            _ = v42
        } else {
            fmt.Println(v40, v2, v40)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    } else {
        for fi43 := 0; fi43 < 5; fi43++ {
            var v44 int32 = (v1 >> ((65535) & 63))
            _ = fi43
            _ = v44
        }
        fmt.Println(v2, v1, v1)
        switch 173 {
        case -726:
            var v45 int = 9223372036854775807
            _ = v45
        }
    }
    if (genFn0() < genFn0()) {
        for fi46 := 0; fi46 < 8; fi46++ {
            v1 += (-(v1))
            var v47 uint = 1
            _ = fi46
            _ = v47
        }
    } else {
        var v48 uint16 = 2
        _ = v48
    }
    v49 := v5[2147483647]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v49
    _ = v5
}
