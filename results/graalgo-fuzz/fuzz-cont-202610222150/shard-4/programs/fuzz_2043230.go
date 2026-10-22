package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint32; F2 string }
func genFn0(p0 uint32, p1 uint16, p2 uint32) int64 {
    if (false || (!(-752 != 371))) {
        fmt.Println(p1, p0, p2)
        v7 := []int{}
        _ = v7
    } else {
        v8 := GenS1{F0: 180, F1: p1, F2: 255}
        _ = v8
    }
    var v9 int32 = 1
    v10 := GenS1{F0: 42, F1: (p1 * (p1 << ((890) & 63))), F2: 77}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 42
}
func genFn1() int8 {
    var v11 int32 = -1
    v12 := GenS2{F0: 1}
    var v13 int16 = 1
    _ = v11
    _ = v12
    _ = v13
    return int8(v11)
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]string)
    var v3 int64 = 42
    var v4 bool = true
    var v5 float64 = 3.14159
    v6 := make(map[string]uint32)
    var v14 uint16 = 32767
    if v4 {
        v15 := []int{-705, -280, 9223372036854775807, 127}
        if (v4 || (!v4)) {
            var v16 float64 = v5
            fmt.Println(v5, v14, v14)
            if true {
                var v17 uint16 = (^(v14))
                if (!v4) {
                    if ((2 < 2) || (!v4)) {
                        v5 += (((1.5 / 3.14159) / float64(1.5)) + v5)
                        switch 128 {
                        case -772:
                            if v4 {
                                v18 := GenS3{F0: 1, F1: 65535, F2: ((("hello" + "foo bar") + "a") + "\n\t")}
                                _ = v18
                            }
                        }
                    } else {
                        if v4 {
                            var v19 int64 = genFn0(915, 65535, 2147483648)
                            var v20 uint16 = v17
                            _ = v19
                            _ = v20
                        } else {
                            var v21 float64 = (v16 + (v5 - v5))
                            delete(v2, 397)
                            _ = v21
                        }
                        switch ("a" + "foo bar") {
                        case "":
                            var v22 float32 = (-2.25 + (1.5 * (0.1 * -2.25)))
                            _ = v22
                        case "\n\t":
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        }
                        _ = len(v15)
                    }
                    if v4 {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        for fi23 := 0; fi23 < 19; fi23++ {
                            var v24 int32 = -2147483648
                            _ = len(v6)
                            for fi25 := 0; fi25 < 7; fi25++ {
                                var v26 uint = 317
                                _ = fi25
                                _ = v26
                            }
                            _ = fi23
                            _ = v24
                        }
                        if false {
                            var v27 int = 9223372036854775807
                            _ = v27
                        }
                    }
                }
                switch 1 {
                case 32767:
                    fmt.Println(v17, v4, v3)
                }
                _ = v17
            }
            _ = v16
        } else {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        }
        _ = v15
    } else {
        fmt.Println(v1, v1)
        if v4 {
            if false {
                var v28 int64 = -1
                _ = v28
            }
        } else {
            fmt.Println(v3, v4)
            fmt.Println(v3)
            v29 := GenS3{F0: 846, F1: 2, F2: "a"}
            _ = v29
        }
        fmt.Println(v3, v4)
    }
    var v30 uint = 18446744073709551615
    for fi31 := 0; fi31 < 8; fi31++ {
        v32 := GenS3{F0: 2147483648, F1: 65535, F2: "foo bar"}
        v33 := []int{fi31}
        switch 4294967295 {
        case 2147483647:
            var v34 float64 = (3.14159 + 1.5)
            _ = v34
        default:
            for fi35 := 0; fi35 < 5; fi35++ {
                switch (("" + ("" + "hello")) + "foo bar") {
                case "\n\t":
                    var v36 int16 = 1000
                    _ = v36
                case "a":
                    var v37 float32 = 0.0
                    _ = v37
                case "foo bar":
                    var v38 int64 = (-(v3))
                    _ = v38
                default:
                    fmt.Println(fi35)
                }
                _ = fi35
            }
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        var v40 uint8 = 2
        v3 -= -731
        _ = fi39
        _ = v40
    }
    var v41 float64 = v5
    for fi42 := 0; fi42 < 13; fi42++ {
        switch fi42 {
        case 1:
            if false {
                var v43 int8 = v1
                _ = v43
            }
        }
        v44 := GenS3{F0: 4294967295, F1: 414, F2: "a"}
        _ = fi42
        _ = v44
    }
    v45 := []int{4294967295, 0, 9223372036854775807, -894}
    for fi46 := 0; fi46 < 5; fi46++ {
        var v47 uint = 861
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v45
    _ = v5
    _ = v6
}
