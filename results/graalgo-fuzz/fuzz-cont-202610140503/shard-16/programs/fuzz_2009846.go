package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int16, p2 bool) uint16 {
    v5 := []int{-2147483648, int(p0), 32768, -2147483648}
    fmt.Println(p0)
    if p2 {
        if p2 {
            fmt.Println(p0, p0, p2)
            switch (("hello" + "a") + "\n\t") {
            case "a":
                if (264 != -1) {
                    v6 := GenS1{F0: ("" + "a"), F1: 513}
                    _ = v6
                }
            default:
                for fi7 := 0; fi7 < 1; fi7++ {
                    if (fi7 >= -9223372036854775808) {
                        v8 := GenS2{F0: (p2 || false), F1: 247}
                        _ = v8
                    }
                    _ = fi7
                }
            }
        } else {
            var v9 uint = 978
            v9 = (^(v9))
            _ = v9
        }
        var v10 int32 = 418
        _ = v10
    } else {
        if (!p2) {
            p1 -= -1
            for fi11 := 0; fi11 < 13; fi11++ {
                v12 := GenS1{F0: "foo bar", F1: int(p0)}
                _ = fi11
                _ = v12
            }
            p1 = p1
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p0 & (p0 >> ((628) & 63)))
}

func main() {
    var v1 int8 = 127
    var v2 bool = ((("" + "a") + (("hello" + "") + "")) != "a")
    v3 := make(map[int]uint32)
    var v4 int = 42
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            fmt.Println(v4)
            switch 32767 {
            case -327:
                if (1 > 0) {
                    var v15 int8 = (v1 + (-(v1)))
                    switch (-(fi13)) {
                    case -2147483648:
                        v4 += fi13
                    case 2147483647:
                        var v16 uint = 42
                        _ = v16
                    case 1:
                        var v17 int8 = v1
                        _ = v17
                    default:
                        v18 := GenS3{F0: "a", F1: v2}
                        _ = v18
                    }
                    v15 += 0
                    _ = v15
                }
            case 1:
                var v19 uint8 = 2
                _ = v19
            case 128:
                v2 = false
            }
            _ = fi14
        }
        _ = fi13
    }
    switch 4294967295 {
    case 42:
        var v20 int8 = (v1 / ((0) | 1))
        _ = v20
    case -9223372036854775808:
        var v21 float32 = 1e10
        _ = v21
    }
    if false {
        if (!(-9223372036854775808 == 0)) {
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 uint = 2
                var v24 uint = v23
                if (v2 && v2) {
                    v25 := GenS4{F0: 42}
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 12; fi26++ {
                        v24 = v24
                        if (v2 || (1 == 2)) {
                            v3[127] = 675
                            var v27 uint64 = 42
                            _ = v27
                        }
                        _ = fi26
                    }
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = len(v3)
        } else {
            v1 += v1
            var v28 int8 = v1
            _ = v28
        }
        switch -198 {
        case 4294967295:
            v1 = v1
        case 42:
            switch (v4 ^ -2147483648) {
            case 2147483647:
                v2 = v2
            case -9223372036854775808:
                v1 = (v1 >> ((2) & 63))
            case 780:
                _ = len(v3)
            default:
                if (888 != 0) {
                    var v29 int8 = 100
                    switch ("hello" + "hello") {
                    case "foo bar":
                        v29 = v29
                    case "\n\t":
                        v3[(v4 | v4)] = 4294967295
                    default:
                        var v30 int16 = -792
                        _ = v30
                    }
                    v1 = (^(v29))
                    _ = v29
                }
            }
        default:
            fmt.Println(v4)
        }
    } else {
        switch v4 {
        case -2147483648:
            fmt.Println(v2, v2)
        default:
            if (true || (0 < 65535)) {
                for fi31 := 0; fi31 < 5; fi31++ {
                    v4 = 4294967295
                    v32 := []int{-9223372036854775808, (v4 ^ (fi31 - 127)), 1}
                    _ = fi31
                    _ = v32
                }
                if v2 {
                    v33 := GenS1{F0: "foo bar", F1: 690}
                    var v34 uint8 = 255
                    _ = v33
                    _ = v34
                }
            }
        }
    }
    var v35 uint16 = 971
    v36 := []int{128}
    var v37 uint = 42
    v37 = v37
    if (32767 >= int(v37)) {
        fmt.Println(v2, v2)
        v2 = false
        v35 += v35
    } else {
        var v38 uint64 = 356
        v4 = (v4 | (v4 | (v4 - v4)))
        _ = v38
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        v4 -= (fi39 & (fi39 + (fi39 | 4294967295)))
        fmt.Println(v1, v2)
        _ = fi39
    }
    v40 := GenS4{F0: 1}
    v2 = false
    v4 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
}
