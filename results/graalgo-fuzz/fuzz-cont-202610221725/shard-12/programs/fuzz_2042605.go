package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 int; F1 uint; F2 uint32 }
type GenS8 struct { F0 int32; F1 int32 }
func genFn0(p0 uint8) uint8 {
    p0 = p0
    if false {
        v6 := []int{-894, 161, 9223372036854775807, int(p0)}
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint16 = 2
            _ = fi7
            _ = v8
        }
        if true {
            fmt.Println(p0)
            fmt.Println(p0)
            var v9 bool = (("" + "foo bar") == "foo bar")
            _ = v9
        } else {
            var v10 int32 = 2147483647
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(v10, v10, v10)
                for fi12 := 0; fi12 < 9; fi12++ {
                    if true {
                        var v13 uint64 = 2
                        p0 += p0
                        var v14 int32 = (v10 + (v10 * (v10 + (v10 + (v10 << ((65535) & 63))))))
                        _ = v13
                        _ = v14
                    } else {
                        v15 := GenS1{F0: v10, F1: uint(p0)}
                        var v16 uint32 = 65535
                        switch "foo bar" {
                        case "\n\t":
                            var v17 int = int(p0)
                            _ = v17
                        }
                        _ = v15
                        _ = v16
                    }
                    _ = fi12
                }
                _ = fi11
            }
            if (3.14159 <= 3.14159) {
                var v18 uint32 = 42
                if (false && false) {
                    v6 = append(v6, 406)
                    _ = len(v6)
                    if true {
                        v19 := GenS2{F0: (true || (((!(!((v18 ^ 424) >= 75))) && true) && (false || (true && false)))), F1: 30}
                        _ = v19
                    } else {
                        var v20 uint8 = 2
                        _ = v20
                    }
                }
                _ = v18
            }
            _ = v10
        }
        _ = v6
    }
    var v21 uint8 = p0
    fmt.Println(v21, p0)
    if (true && (842 > 1000)) {
        for fi22 := 0; fi22 < 9; fi22++ {
            if false {
                v23 := GenS1{F0: -314, F1: 255}
                _ = v23
            }
            var v24 string = ("a" + ("foo bar" + "foo bar"))
            _ = fi22
            _ = v24
        }
        v25 := GenS1{F0: 1, F1: 4294967295}
        p0 = v21
        _ = v25
    } else {
        var v26 uint64 = 605
        _ = v26
    }
    _ = p0
    _ = v21
    return v21
}
func genFn1(p0 uint, p1 uint, p2 string) int32 {
    for fi27 := 0; fi27 < 17; fi27++ {
        p1 = (p1 / ((p1) | 1))
        p0 += 255
        p2 = p2
        _ = fi27
    }
    p1 += p0
    switch (("" + p2) + "\n\t") {
    case "hello":
        if (-2147483648 == 42) {
            p1 = p1
        } else {
            switch "foo bar" {
            case "a":
                p0 -= 256
            }
            v28 := GenS3{F0: false, F1: 1000}
            p0 = p1
            _ = v28
        }
    case "a":
        v29 := GenS1{F0: 809, F1: (p1 >> ((p1) & 63))}
        _ = v29
    case "":
        if (true && (false || (!(!false)))) {
            fmt.Println(p1)
            p0 += 42
            v30 := GenS1{F0: 1, F1: p0}
            _ = v30
        } else {
            p1 -= (p0 | (^(p1)))
            for fi31 := 0; fi31 < 11; fi31++ {
                switch (("hello" + p2) + p2) {
                case "foo bar":
                    if true {
                        var v32 int = fi31
                        var v33 uint32 = 2147483648
                        _ = v32
                        _ = v33
                    }
                case "\n\t":
                    v34 := GenS1{F0: -2147483648, F1: p0}
                    _ = v34
                default:
                    p0 = p0
                }
                if true {
                    v35 := []int{fi31, -693}
                    if len(v35) > 0 {
                        v35[0] = -867
                    }
                    v36 := GenS3{F0: false, F1: -398}
                    _ = v35
                    _ = v36
                } else {
                    var v37 float32 = 0.0
                    if (9223372036854775807 != 42) {
                        fmt.Println(p1, fi31)
                    } else {
                        v38 := GenS1{F0: -399, F1: 1}
                        if false {
                            fmt.Println(fi31)
                            p2 = "hello"
                            var v39 int32 = -2147483648
                            _ = v39
                        } else {
                            p0 -= 0
                        }
                        var v40 string = "a"
                        _ = v38
                        _ = v40
                    }
                    _ = v37
                }
                _ = fi31
            }
        }
    default:
        var v41 string = (p2 + "a")
        _ = v41
    }
    _ = p0
    _ = p1
    _ = p2
    return -851
}

func main() {
    v1 := []uint32{2147483648, 0}
    v2 := []string{("a" + ""), "hello", "", ""}
    v3 := []int{}
    v4 := []string{}
    var v5 float32 = 3.14159
    v5 += (float32(float64(float32(((1e10 / -2.25) / 1.5)))) * (v5 + v5))
    if len(v1) > 0 {
        v1[0] = 0
    }
    var v42 int = 523
    v43 := GenS2{F0: (false && true), F1: 42}
    v4 = append(v4, "")
    v4 = append(v4, (("hello" + "\n\t") + "foo bar"))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        v44 := GenS4{F0: 13, F1: 2147483647, F2: 100}
        v5 -= v5
        _ = v44
    } else {
        v42 -= (v42 % ((v42) | 1))
        var v45 int64 = -665
        switch v42 {
        case 32767:
            v5 -= 1.5
        case -788:
            if false {
                v46 := GenS3{F0: false, F1: 0}
                _ = v46
            } else {
                v45 = -1
                var v47 float64 = 0.1
                fmt.Println(v5, v45)
                _ = v47
            }
        case 453:
            v48 := GenS5{F0: (0 == 1000), F1: "\n\t"}
            _ = v48
        default:
            switch v42 {
            case 127:
                v45 -= v45
            case 4294967295:
                v49 := GenS6{F0: 42, F1: v42}
                _ = v49
            case 9223372036854775807:
                var v50 int16 = 42
                _ = v50
            }
        }
        _ = v45
    }
    var v51 uint8 = 30
    var v52 bool = (42 >= int16(v51))
    switch (("foo bar" + ((("foo bar" + "\n\t") + "") + "a")) + "hello") {
    case "a":
        if true {
            var v53 float64 = 0.0
            v5 = (0.1 - 1e10)
            _ = v53
        } else {
            if v52 {
                v5 += v5
            }
            var v54 bool = (127 != 74)
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 12; fi55++ {
        var v56 uint16 = 157
        v57 := GenS7{F0: (^(fi55)), F1: 2, F2: 1}
        _ = fi55
        _ = v56
        _ = v57
    }
    var v58 int8 = -37
    fmt.Println(v5, v58)
    v59 := GenS8{F0: 1, F1: 0}
    for fi60 := 0; fi60 < 12; fi60++ {
        v61 := GenS5{F0: (true || false), F1: ("foo bar" + ((("foo bar" + "a") + (("hello" + "") + "a")) + ""))}
        for fi62 := 0; fi62 < 11; fi62++ {
            var v63 string = (("hello" + "foo bar") + ("" + (("\n\t" + "\n\t") + "foo bar")))
            _ = fi62
            _ = v63
        }
        _ = fi60
        _ = v61
    }
    var v64 int32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v51
    _ = v52
    _ = v58
    _ = v59
    _ = v64
}
