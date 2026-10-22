package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32) bool {
    var v7 uint64 = 2
    v8 := []int{2147483647, 220}
    var v9 bool = (false || (!(-9223372036854775808 >= 32767)))
    v7 = v7
    v7 = (v7 & 1)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (255 < 2)
}

func main() {
    var v1 float32 = (((1.5 / 1.5) / float32((1e10 * ((1e10 / 1e10) * -2.25)))) - (((float32((0.0 / float64((3.14159 - -2.25)))) / 3.14159) - -2.25) - 3.14159))
    var v2 uint = 880
    var v3 int8 = -1
    var v4 uint = (v2 | (v2 / (((v2 << ((496) & 63))) | 1)))
    var v5 int32 = -994
    v6 := make(map[uint32]int)
    if false {
        if (!false) {
            var v10 uint16 = 635
            switch 356 {
            case 1:
                switch -755 {
                case -360:
                    var v11 uint64 = 16
                    _ = v11
                case -9223372036854775808:
                    _ = len(v6)
                }
            default:
                var v12 uint8 = 59
                _ = v12
            }
            _ = v10
        } else {
            v13 := []int{1, -2147483648}
            v5 = 42
            _ = v13
        }
    }
    v5 = -850
    if (1 != 32767) {
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 string = "hello"
            v3 = (v3 << (((v4 ^ (^(v2)))) & 63))
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(v3, v4)
            v17 := []int{(-(fi16)), fi16, fi16, fi16}
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            if false {
                var v19 int16 = 115
                if false {
                    switch ((("" + ("foo bar" + "a")) + (("a" + ("hello" + "\n\t")) + "hello")) + (("hello" + "") + ("foo bar" + ("a" + ("a" + "hello"))))) {
                    case "":
                        if true {
                            for fi20 := 0; fi20 < 7; fi20++ {
                                var v21 int32 = (v5 / (((v5 >> ((255) & 63))) | 1))
                                _ = fi20
                                _ = v21
                            }
                        }
                    case "\n\t":
                        var v22 int16 = (v19 | 1)
                        _ = v22
                    case "a":
                        v23 := v6[2]
                        _ = v23
                    }
                } else {
                    v24 := GenS1{F0: 324}
                    for fi25 := 0; fi25 < 8; fi25++ {
                        var v26 string = "foo bar"
                        v2 = v2
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                }
                var v27 int8 = v3
                _ = v19
                _ = v27
            } else {
                switch ("hello" + "a") {
                case "hello":
                    fmt.Println(fi18, v3)
                case "":
                    v28 := []int{(fi18 - -675)}
                    _ = v28
                case "foo bar":
                    if true {
                        var v29 int64 = 9223372036854775807
                        switch fi18 {
                        case 9223372036854775807:
                            if (-32768 != 0) {
                                if false {
                                    v30 := GenS1{F0: (v4 | v2)}
                                    v31 := GenS2{F0: 458, F1: (v3 >> (((v4 >> (((v4 % ((v4) | 1))) & 63))) & 63))}
                                    _ = v30
                                    _ = v31
                                }
                            } else {
                                if (v29 <= -191) {
                                    for fi32 := 0; fi32 < 4; fi32++ {
                                        v2 += v4
                                        fmt.Println(fi32)
                                        _ = fi32
                                    }
                                    v1 = 1e10
                                    fmt.Println(v1)
                                }
                            }
                        case -2147483648:
                            v33 := []int{fi18}
                            _ = v33
                        }
                        var v34 uint16 = 0
                        _ = v29
                        _ = v34
                    }
                }
            }
            v35 := GenS1{F0: (^(v2))}
            _ = fi18
            _ = v35
        }
    } else {
        v3 -= -128
        var v36 int32 = v5
        v4 += v4
        _ = v36
    }
    _ = len(v6)
    if false {
        var v37 uint32 = 0
        var v38 float32 = (3.14159 - v1)
        fmt.Println(v3, v2, v3)
        _ = v37
        _ = v38
    } else {
        var v39 int16 = 32767
        _ = v39
    }
    v5 = -1
    v5 -= 0
    if false {
        v40 := []int{93, 4294967295, -1, 464}
        _ = len(v40)
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
        _ = v40
    }
    var v41 int = -177
    v42 := GenS3{F0: 0}
    var v43 float64 = (3.14159 / 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
