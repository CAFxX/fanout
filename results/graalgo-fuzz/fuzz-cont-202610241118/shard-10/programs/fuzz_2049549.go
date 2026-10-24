package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 1
    v3 := []int{-1, -802}
    v2 += v2
    fmt.Println(v2)
    v2 = 2
    if false {
        for fi4 := 0; fi4 < 1; fi4++ {
            if (706 == 42) {
                switch "" {
                case "\n\t":
                    if false {
                        if false {
                            if false {
                                fmt.Println(fi4, v2)
                            } else {
                                v5 := GenS1{F0: 42, F1: -2147483648}
                                for fi6 := 0; fi6 < 6; fi6++ {
                                    v7 := GenS1{F0: 2147483647, F1: 42}
                                    v8 := GenS1{F0: 65535, F1: 2147483647}
                                    _ = len(v1)
                                    var v9 uint32 = 65535
                                    v10 := v1[""]
                                    _ = fi6
                                    _ = v10
                                    _ = v7
                                    _ = v8
                                    _ = v9
                                }
                                v11 := GenS1{F0: 509, F1: -494}
                                _ = v11
                                _ = v5
                            }
                            var v12 int16 = -32768
                            if true {
                                var v13 int = fi4
                                for fi14 := 0; fi14 < 6; fi14++ {
                                    if len(v3) > 0 {
                                        v3[0] = (fi14 & (v13 & fi4))
                                    }
                                    v15 := GenS1{F0: 0, F1: 2147483647}
                                    _ = fi14
                                    _ = v15
                                }
                                fmt.Println(v13, v2)
                                _ = v13
                            }
                            _ = v12
                        }
                    } else {
                        if (!(1 == 18446744073709551615)) {
                            _ = len(v1)
                            switch fi4 {
                            case -9223372036854775808:
                                v16 := GenS2{F0: 726, F1: 1}
                                _ = v16
                            case 446:
                                var v17 uint = 65535
                                _ = v17
                            default:
                                fmt.Println(v2, v2)
                            }
                            if false {
                                for fi18 := 0; fi18 < 9; fi18++ {
                                    var v19 int = 4294967295
                                    switch 0 {
                                    case -190:
                                        delete(v1, "a")
                                    case -1:
                                        fmt.Println(fi4, v2, v2)
                                    case 1:
                                        v19 = fi4
                                    }
                                    _ = fi18
                                    _ = v19
                                }
                                v2 = 32767
                            }
                        } else {
                            fmt.Println(fi4)
                            v2 -= (v2 + (v2 >> ((uint(v2)) & 63)))
                        }
                        _ = len(v3)
                    }
                case "foo bar":
                    switch ("a" + (((("hello" + "hello") + "") + "hello") + (("foo bar" + "") + "foo bar"))) {
                    case "foo bar":
                        v20 := GenS3{F0: 0, F1: 0, F2: 4294967295}
                        _ = v20
                    case "hello":
                        if (-633 < -1) {
                            v2 = (v2 % ((v2) | 1))
                            fmt.Println(fi4, fi4)
                            var v21 float32 = 0.0
                            _ = v21
                        }
                    }
                default:
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 uint64 = 433
                        if true {
                            if (!(32767 != 1)) {
                                v2 -= 255
                                var v24 string = ("" + "foo bar")
                                _ = v24
                            }
                            v25 := GenS2{F0: -32768, F1: -1}
                            v23 -= v23
                            _ = v25
                        }
                        _ = fi22
                        _ = v23
                    }
                }
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 float64 = 3.14159
                    _ = len(v3)
                    _ = fi26
                    _ = v27
                }
            } else {
                v28 := GenS4{F0: 1, F1: (fi4 - (fi4 % ((fi4) | 1)))}
                _ = v28
            }
            fmt.Println(fi4)
            _ = fi4
        }
    } else {
        var v29 uint = 868
        _ = v29
    }
    v2 -= (v2 + 755)
    fmt.Println(v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v30 := GenS5{F0: 65535}
    delete(v1, ("foo bar" + "hello"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
}
