package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint) int32 {
    var v6 uint64 = 4294967295
    switch "a" {
    case "a":
        if (42 <= 42) {
            p0 -= -1
            v7 := []int{127, -9223372036854775808}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 18; fi8++ {
                v6 -= (v6 + 4294967295)
                _ = fi8
            }
            var v9 uint32 = 42
            _ = v9
        }
    case "foo bar":
        p0 += (p0 + (p0 & 2147483647))
    default:
        fmt.Println(p0)
    }
    if (false && true) {
        v6 -= v6
    }
    v10 := GenS1{F0: (p0 | p0)}
    v11 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    return p0
}
func genFn1(p0 int16) uint8 {
    v12 := []int{-966, -922}
    var v13 uint = 1
    _ = p0
    _ = v12
    _ = v13
    return 197
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := []int{-1, -9223372036854775808, 2147483647, 32768}
    var v4 int16 = 0
    var v5 int = 128
    v14 := GenS1{F0: 1}
    if true {
        v15 := GenS1{F0: -2147483648}
        fmt.Println(v5, v5)
        switch 0 {
        case 128:
            var v16 float32 = 1.5
            _ = v16
        case -9223372036854775808:
            switch "foo bar" {
            case "a":
                v17 := GenS1{F0: 2147483647}
                _ = v17
            case "foo bar":
                fmt.Println(v4, v5)
            default:
                var v18 uint16 = 32767
                _ = v18
            }
        case 4294967295:
            if true {
                for fi19 := 0; fi19 < 16; fi19++ {
                    v20 := GenS1{F0: -1}
                    switch "foo bar" {
                    case "\n\t":
                        for fi21 := 0; fi21 < 18; fi21++ {
                            var v22 uint64 = 2
                            _ = fi21
                            _ = v22
                        }
                    }
                    _ = fi19
                    _ = v20
                }
                v5 -= (v5 | (v5 % (((v5 << ((1) & 63))) | 1)))
                fmt.Println(v5, v4)
            }
        default:
            for fi23 := 0; fi23 < 12; fi23++ {
                fmt.Println(v5, fi23)
                _ = fi23
            }
        }
        _ = v15
    }
    v4 += v4
    switch ("\n\t" + ("a" + "a")) {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = (v5 << ((42) & 63))
        }
    case "a":
        v24 := GenS2{F0: 30, F1: "\n\t", F2: genFn1((v4 + 737))}
        _ = v24
    }
    switch (v5 + int(v4)) {
    case -2147483648:
        if false {
            _ = len(v3)
        }
    case 150:
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 int8 = 42
            switch "hello" {
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "foo bar":
                for fi27 := 0; fi27 < 19; fi27++ {
                    var v28 int32 = 1
                    _ = fi27
                    _ = v28
                }
            case "a":
                switch (v5 * (v5 % (((-(v5))) | 1))) {
                case 1:
                    switch "\n\t" {
                    case "":
                        _ = len(v3)
                    case "a":
                        fmt.Println(v5, v4)
                    case "\n\t":
                        v5 -= (v5 | v5)
                    default:
                        var v29 uint32 = 0
                        _ = v29
                    }
                case -9223372036854775808:
                    fmt.Println(fi25)
                }
            default:
                var v30 uint16 = 1
                _ = v30
            }
            if ((42 <= 42) || true) {
                if len(v3) > 0 {
                    v3[0] = (v5 << ((0) & 63))
                }
            }
            _ = fi25
            _ = v26
        }
    case 2147483647:
        v2[(v5 | (v5 + (v5 << ((35) & 63))))] = ("\n\t" + (("a" + "") + "hello"))
    }
    var v31 uint32 = 899
    v31 = 1
    var v32 uint16 = 693
    var v33 uint64 = 18446744073709551615
    if (3.14159 < 0.1) {
        var v34 int32 = -2147483648
        for fi35 := 0; fi35 < 20; fi35++ {
            var v36 string = (("a" + "\n\t") + "hello")
            var v37 uint = 18446744073709551615
            v38 := GenS3{F0: v32, F1: (v37 % ((v37) | 1)), F2: 9223372036854775807}
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        v5 += 127
        _ = v34
    }
    v5 = v5
    var v39 uint8 = 2
    fmt.Println(v5)
    if true {
        v40 := v2[int(v4)]
        v41 := GenS4{F0: (v5 % ((-1) | 1)), F1: 123}
        _ = v40
        _ = v41
    } else {
        switch "a" {
        case "foo bar":
            switch "hello" {
            case "\n\t":
                var v42 uint32 = v31
                _ = v42
            case "hello":
                v43 := GenS2{F0: (v33 - 0), F1: "foo bar", F2: (v39 - v39)}
                _ = v43
            }
        case "hello":
            fmt.Println(v39, v33)
        case "":
            for fi44 := 0; fi44 < 3; fi44++ {
                v45 := GenS4{F0: v5, F1: v33}
                _ = fi44
                _ = v45
            }
        default:
            switch (-(v5)) {
            case 127:
                if false {
                    if (!true) {
                        fmt.Println(v5)
                        v46 := GenS1{F0: -1}
                        var v47 uint32 = v31
                        _ = v46
                        _ = v47
                    }
                    var v48 uint16 = 255
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v48
                } else {
                    if false {
                        v31 = v31
                        v5 -= v5
                        v3 = append(v3, 2147483647)
                    }
                }
            case 4294967295:
                if ((1e10 / float32(1.5)) >= 0.1) {
                    v49 := GenS5{F0: 1}
                    var v50 uint64 = (v33 % ((668) | 1))
                    _ = v49
                    _ = v50
                }
            default:
                switch (v5 % ((v5) | 1)) {
                case 2147483647:
                    delete(v1, v31)
                case -524:
                    var v51 string = "foo bar"
                    _ = v51
                default:
                    var v52 string = ((("\n\t" + ("hello" + "foo bar")) + "") + ("a" + "foo bar"))
                    _ = v52
                }
            }
        }
        var v53 int32 = 0
        _ = v53
    }
    if (false && (!((v32 >> ((uint(v32)) & 63)) > 0))) {
        v54 := GenS3{F0: (v32 - 32767), F1: 1, F2: -184}
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v5
}
