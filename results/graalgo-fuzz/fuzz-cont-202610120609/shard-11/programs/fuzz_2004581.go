package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int; F1 int16; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 15
    var v3 uint32 = 2
    var v4 uint32 = 83
    var v5 float32 = -2.25
    v5 = 0.1
    switch "\n\t" {
    case "\n\t":
        v2 = (v2 % ((v2) | 1))
    case "foo bar":
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 uint = v2
            _ = fi6
            _ = v7
        }
    }
    var v8 int16 = -1
    if (!(false && true)) {
        v3 -= v3
        fmt.Println(v3, v5, v3)
        var v9 uint = v2
        _ = v9
    } else {
        if (!false) {
            v3 = v3
            if false {
                switch 0 {
                case -550:
                    if true {
                        v5 -= (v5 - (v5 + -2.25))
                    } else {
                        if true {
                            v10 := GenS1{F0: 42}
                            var v11 float64 = ((float64(v5) / 1e10) - 1.5)
                            _ = v10
                            _ = v11
                        } else {
                            v12 := GenS2{F0: 151, F1: ""}
                            _ = v12
                        }
                    }
                case 32768:
                    var v13 uint32 = v4
                    _ = v13
                case -760:
                    switch -1 {
                    case 128:
                        fmt.Println(v8)
                    case -515:
                        for fi14 := 0; fi14 < 3; fi14++ {
                            for fi15 := 0; fi15 < 15; fi15++ {
                                v5 += float32((1.5 / ((3.14159 + 1e10) * (1.5 / 0.1))))
                                v5 = 3.14159
                                _ = fi15
                            }
                            _ = fi14
                        }
                    default:
                        fmt.Println(v8)
                    }
                }
                v8 = (v8 << ((4294967295) & 63))
            }
        }
        var v16 int16 = int16(v2)
        _ = v16
    }
    var v17 int8 = -30
    v17 += (-(v17))
    switch ("a" + (("\n\t" + "") + "hello")) {
    case "hello":
        var v18 int8 = 0
        _ = v18
    case "\n\t":
        v19 := GenS2{F0: 0, F1: "hello"}
        _ = v19
    }
    if false {
        v20 := GenS3{F0: -2147483648, F1: (v8 * int16(v4)), F2: (v17 ^ (v17 & (v17 ^ (^(v17)))))}
        v21 := GenS2{F0: 42, F1: "\n\t"}
        if false {
            for fi22 := 0; fi22 < 9; fi22++ {
                fmt.Println(v3, v5, v17)
                fmt.Println(v5, v5, v2)
                v3 = (^(v4))
                _ = fi22
            }
        } else {
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 uint8 = 42
                _ = fi23
                _ = v24
            }
            if false {
                v25 := []int{4294967295}
                _ = v25
            } else {
                v2 += uint(v4)
                for fi26 := 0; fi26 < 7; fi26++ {
                    v27 := GenS2{F0: 0, F1: ""}
                    _ = fi26
                    _ = v27
                }
                v28 := GenS4{F0: 42}
                _ = v28
            }
            var v29 int32 = 17
            _ = v29
        }
        _ = v20
        _ = v21
    }
    var v30 int16 = 1000
    fmt.Println(v5)
    if true {
        if true {
            fmt.Println(v17, v5)
            for fi31 := 0; fi31 < 17; fi31++ {
                for fi32 := 0; fi32 < 15; fi32++ {
                    var v33 float32 = v5
                    v34 := []int{int(v2), fi32, (fi31 * -1), (fi31 ^ (-(fi31)))}
                    v35 := GenS5{F0: (fi31 ^ 2147483647)}
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                }
                _ = fi31
            }
            v36 := []int{9223372036854775807, 0}
            _ = v36
        } else {
            delete(v1, (v3 & 700))
            for fi37 := 0; fi37 < 17; fi37++ {
                var v38 int64 = 1
                v5 = float32(0.1)
                v4 -= v3
                _ = fi37
                _ = v38
            }
        }
    } else {
        fmt.Println(v30, v3, v30)
    }
    if true {
        var v39 int64 = 42
        var v40 uint16 = 32767
        _ = len(v1)
        _ = v39
        _ = v40
    }
    v2 = (^(v2))
    delete(v1, 2147483648)
    v41 := GenS3{F0: 872, F1: v30, F2: -128}
    v42 := v1[(v4 + 326)]
    if true {
        var v43 int64 = 1
        v4 = 2147483648
        switch 9223372036854775807 {
        case 127:
            v17 = (^(v17))
        case 538:
            v2 -= 1
        default:
            fmt.Println(v17)
        }
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v8
}
