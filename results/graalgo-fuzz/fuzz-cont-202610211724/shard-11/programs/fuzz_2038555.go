package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint64 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 bool; F1 bool; F2 bool }
type GenS4 struct { F0 uint; F1 uint8; F2 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 uint8; F1 int64 }
func genFn0(p0 bool) uint64 {
    if p0 {
        p0 = p0
    }
    v7 := []int{127, 0, 32767, 0}
    if p0 {
        if (255 <= 218) {
            if p0 {
                if p0 {
                    var v8 uint64 = 1
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 3; fi9++ {
                        var v10 uint8 = 2
                        fmt.Println(v10)
                        fmt.Println(v10)
                        _ = fi9
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 5; fi11++ {
                        if len(v7) > 0 {
                            v7[0] = fi11
                        }
                        p0 = p0
                        v12 := GenS1{F0: 641, F1: (1 != 65535), F2: 0}
                        _ = fi11
                        _ = v12
                    }
                }
                var v13 uint64 = 4294967295
                _ = v13
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            fmt.Println(p0)
        } else {
            fmt.Println(p0)
            var v14 uint16 = 327
            _ = v14
        }
        v15 := GenS2{F0: "", F1: ("foo bar" + ("hello" + "a"))}
        switch "" {
        case "a":
            var v16 int64 = -207
            _ = v16
        case "":
            if (107 >= 0) {
                fmt.Println(p0)
            }
        default:
            for fi17 := 0; fi17 < 16; fi17++ {
                v18 := GenS3{F0: p0, F1: p0, F2: (("foo bar" + "") != "\n\t")}
                p0 = (!p0)
                v7 = append(v7, 42)
                _ = fi17
                _ = v18
            }
        }
        _ = v15
    } else {
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS1{F0: 0, F1: true, F2: 824}
            _ = fi19
            _ = v20
        }
        var v21 int64 = -9223372036854775808
        _ = v21
    }
    p0 = true
    var v22 bool = p0
    _ = p0
    _ = v22
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 int16) int16 {
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := []int{32768, (-(fi23)), fi23}
        _ = fi23
        _ = v24
    }
    switch 32768 {
    case 2147483647:
        fmt.Println(p0, p1)
    default:
        p0 = 513
    }
    var v25 int32 = (p0 * (p0 >> ((116) & 63)))
    v26 := GenS4{F0: 864, F1: 127, F2: true}
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return (p1 - p1)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 1
    v3 := []string{"hello", "foo bar", "a"}
    var v4 uint64 = 4294967295
    v5 := make(map[uint32]string)
    var v6 uint = 18446744073709551615
    for fi27 := 0; fi27 < 19; fi27++ {
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 int8 = -125
            var v30 uint64 = uint64(v6)
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi27
    }
    var v31 int16 = -922
    v4 = 2
    v6 += 255
    var v32 uint16 = 249
    v31 -= v31
    if false {
        if true {
            var v33 int8 = 100
            if true {
                _ = len(v3)
                for fi34 := 0; fi34 < 14; fi34++ {
                    v35 := v5[2147483647]
                    switch "a" {
                    case "hello":
                        v32 -= (v32 - v32)
                    case "foo bar":
                        var v36 uint64 = 18446744073709551615
                        _ = v36
                    }
                    _ = fi34
                    _ = v35
                }
            } else {
                v37 := GenS1{F0: (v6 | v6), F1: false, F2: 470}
                fmt.Println(v32, v32)
                _ = v37
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v33
        }
        for fi38 := 0; fi38 < 3; fi38++ {
            v39 := GenS5{F0: 950}
            switch 128 {
            case 127:
                switch "hello" {
                case "hello":
                    if (true || (true || ((-(fi38)) != fi38))) {
                        var v40 int32 = 0
                        switch fi38 {
                        case -35:
                            v41 := GenS2{F0: "a", F1: "hello"}
                            _ = v41
                        default:
                            var v42 int8 = 1
                            _ = v42
                        }
                        _ = v40
                    } else {
                        if (!true) {
                            var v43 int16 = genFn1(2147483647, 1000)
                            v44 := GenS6{F0: v2, F1: 947}
                            v1[1] = (((!(!true)) || true) && (2147483647 <= -2147483648))
                            _ = v43
                            _ = v44
                        } else {
                            fmt.Println(v6, v4, fi38)
                            v32 = v32
                            var v45 uint32 = 456
                            _ = v45
                        }
                    }
                case "a":
                    switch 9223372036854775807 {
                    case 2147483647:
                        fmt.Println(v4)
                    case 574:
                        var v46 int = (-(fi38))
                        _ = v46
                    default:
                        v47 := v1[65535]
                        _ = v47
                    }
                }
            case 128:
                var v48 float64 = 0.1
                _ = v48
            case 4294967295:
                var v49 int16 = v31
                _ = v49
            default:
                if false {
                    if (false && (!(!false))) {
                        var v50 float32 = (-2.25 / 1.5)
                        _ = v50
                    } else {
                        fmt.Println(v2, v6, v6)
                        for fi51 := 0; fi51 < 10; fi51++ {
                            var v52 string = "foo bar"
                            fmt.Println(v4, v6, fi51)
                            var v53 string = v52
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        v6 = (v6 | (v6 + v6))
                    }
                    var v54 int = 1
                    for fi55 := 0; fi55 < 5; fi55++ {
                        var v56 int8 = int8(v2)
                        _ = fi55
                        _ = v56
                    }
                    _ = v54
                }
            }
            _ = fi38
            _ = v39
        }
    } else {
        for fi57 := 0; fi57 < 9; fi57++ {
            v32 += 278
            v31 = (v31 * 589)
            var v58 string = "hello"
            _ = fi57
            _ = v58
        }
    }
    if false {
        var v59 int16 = (^(v31))
        switch "\n\t" {
        case "foo bar":
            if (!true) {
                var v60 uint = v6
                _ = v60
            }
        default:
            var v61 float32 = ((0.1 + -2.25) * 0.1)
            _ = v61
        }
        v62 := GenS1{F0: v6, F1: ((true || false) || true), F2: (^(v4))}
        _ = v59
        _ = v62
    }
    v63 := GenS6{F0: v2, F1: 0}
    for fi64 := 0; fi64 < 9; fi64++ {
        v31 += (-(v31))
        _ = len(v5)
        var v65 int = fi64
        _ = fi64
        _ = v65
    }
    for fi66 := 0; fi66 < 20; fi66++ {
        for fi67 := 0; fi67 < 13; fi67++ {
            v6 = 65535
            var v68 float64 = (1.5 / 1.5)
            _ = fi67
            _ = v68
        }
        v69 := GenS1{F0: v6, F1: (float32((-2.25 / 0.1)) == 0.0), F2: genFn0((true || false))}
        _ = fi66
        _ = v69
    }
    for fi70 := 0; fi70 < 16; fi70++ {
        v4 = genFn0(false)
        _ = fi70
    }
    _ = len(v5)
    delete(v5, 1)
    switch (("hello" + ("\n\t" + "a")) + "a") {
    case "hello":
        var v71 string = (("hello" + ("a" + "hello")) + "")
        _ = v71
    case "foo bar":
        fmt.Println(v6, v31, v6)
    default:
        var v72 uint16 = (v32 + v32)
        _ = v72
    }
    v73 := GenS7{F0: 255, F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v73
}
