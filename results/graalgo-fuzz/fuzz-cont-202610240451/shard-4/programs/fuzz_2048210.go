package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int32; F1 int32; F2 int64 }
func genFn0() int32 {
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 32767, F1: 0, F2: -1}
        fmt.Println(v6, fi7)
        _ = fi7
        _ = v8
    }
    v9 := []int{-302, 943}
    var v10 uint8 = 133
    _ = v10
    _ = v6
    _ = v9
    return -1
}

func main() {
    var v1 bool = true
    v2 := []uint32{4294967295, 1}
    v3 := make(map[int]bool)
    var v4 string = ""
    v5 := []int{}
    var v11 uint16 = 835
    var v12 int32 = 0
    v13 := GenS2{F0: (!v1), F1: 162, F2: true}
    if v1 {
        v12 = 0
    } else {
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 int = -882
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 bool = v1
                _ = fi16
                _ = v17
            }
            fmt.Println(v4, fi14, v11)
            _ = fi14
            _ = v15
        }
        v18 := v3[-974]
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v18
    }
    var v19 int16 = 0
    v4 = (("foo bar" + v4) + v4)
    var v20 float64 = float64(1e10)
    if ((false && (!v1)) && false) {
        var v21 bool = v1
        for fi22 := 0; fi22 < 11; fi22++ {
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 int64 = 42
                var v25 uint = 18446744073709551615
                v26 := GenS3{F0: 65535}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi22
        }
        _ = v21
    } else {
        var v27 uint64 = 851
        switch (v4 + v4) {
        case "\n\t":
            v20 = (v20 * v20)
        case "":
            switch 0 {
            case 974:
                if v1 {
                    v3[103] = (true || v1)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    for fi28 := 0; fi28 < 16; fi28++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi28
                    }
                } else {
                    if true {
                        delete(v3, -177)
                    } else {
                        v29 := GenS4{F0: uint8(v11), F1: v27}
                        _ = v29
                    }
                    var v30 int64 = 42
                    _ = v30
                }
            case 42:
                for fi31 := 0; fi31 < 12; fi31++ {
                    v11 += v11
                    if false {
                        _ = len(v3)
                    } else {
                        v12 = genFn0()
                        fmt.Println(v4)
                    }
                    for fi32 := 0; fi32 < 19; fi32++ {
                        var v33 uint16 = v11
                        delete(v3, (fi32 * 424))
                        _ = fi32
                        _ = v33
                    }
                    _ = fi31
                }
            case -139:
                if (!((v12 >> ((2) & 63)) < v12)) {
                    var v34 int = -2147483648
                    _ = v34
                } else {
                    if (!(v1 && (!(!(3.14159 == (3.14159 * -2.25)))))) {
                        switch -1 {
                        case 4294967295:
                            fmt.Println(v11, v1)
                        case -9223372036854775808:
                            var v35 uint = 18446744073709551615
                            _ = v35
                        }
                        for fi36 := 0; fi36 < 16; fi36++ {
                            fmt.Println(v20, v1, fi36)
                            v3[(fi36 * fi36)] = false
                            fmt.Println(v20)
                            _ = fi36
                        }
                        if false {
                            if (1 == v27) {
                                v37 := GenS5{F0: int32(v27), F1: -1, F2: 0}
                                _ = v37
                            } else {
                                if (v1 && (-2147483648 <= -9223372036854775808)) {
                                    for fi38 := 0; fi38 < 6; fi38++ {
                                        v27 = v27
                                        _ = fi38
                                    }
                                    for fi39 := 0; fi39 < 4; fi39++ {
                                        v20 -= 1e10
                                        v2 = append(v2, 42)
                                        _ = fi39
                                    }
                                    if v1 {
                                        if ((-2.25 * float32(v20)) <= 0.1) {
                                            v40 := GenS5{F0: v12, F1: (v12 >> ((379) & 63)), F2: -639}
                                            _ = v40
                                        } else {
                                            if v1 {
                                                v5 = append(v5, 42)
                                            } else {
                                                _ = len(v3)
                                                v12 = genFn0()
                                                if v1 {
                                                    v12 = genFn0()
                                                    v5 = append(v5, -2147483648)
                                                }
                                            }
                                            var v41 bool = v1
                                            _ = v41
                                        }
                                    } else {
                                        var v42 int8 = -128
                                        fmt.Println(v4, v12, v27)
                                        fmt.Println(v20)
                                        _ = v42
                                    }
                                }
                                _ = len(v5)
                            }
                            v27 = 1
                            _ = len(v5)
                        }
                    }
                    for fi43 := 0; fi43 < 5; fi43++ {
                        var v44 int32 = (v12 * genFn0())
                        var v45 int64 = 9223372036854775807
                        if ((0.0 - 3.14159) == 1e10) {
                            v27 = (v27 - (v27 - (v27 * (v27 + (v27 << ((450) & 63))))))
                        } else {
                            v4 = v4
                            v45 -= v45
                        }
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                }
            }
        default:
            var v46 int16 = -204
            _ = v46
        }
        _ = v27
    }
    v5 = append(v5, 4294967295)
    v11 = (v11 - v11)
    v47 := v3[42]
    if v47 {
        fmt.Println(v4, v12)
        v48 := GenS3{F0: v11}
        _ = v48
    }
    var v49 uint = 255
    v49 = (^(v49))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v5
}
