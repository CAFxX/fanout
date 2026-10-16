package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int }
type GenS2 struct { F0 uint64; F1 int8; F2 int16 }
type GenS3 struct { F0 string; F1 uint32; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 int; F1 int; F2 int64 }
func genFn0(p0 uint8, p1 int8, p2 int8) int32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 uint16 = 2
        var v7 int64 = 9223372036854775807
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 928 {
    case -364:
        v8 := []int{446, 1, -210, 546}
        _ = v8
    default:
        v9 := []int{-1, 4294967295}
        _ = v9
    }
    fmt.Println(p1)
    var v10 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -2147483648
}
func genFn1(p0 uint32, p1 uint64) string {
    fmt.Println(p0, p1)
    fmt.Println(p0, p0)
    if (!((p1 * p1) > (^(p1)))) {
        switch "" {
        case "\n\t":
            fmt.Println(p1, p0)
        }
    } else {
        if (65535 > 65535) {
            if true {
                var v11 uint8 = uint8(p0)
                v12 := GenS1{F0: 0, F1: -9223372036854775808, F2: -2147483648}
                _ = v11
                _ = v12
            }
            switch ((("foo bar" + "foo bar") + "") + "a") {
            case "hello":
                p0 = p0
            case "":
                var v13 int32 = 814
                _ = v13
            case "foo bar":
                v14 := []int{327, -813}
                _ = v14
            default:
                if (1.5 != 3.14159) {
                    if (-2.25 < float64(float32((1.5 - 0.1)))) {
                        if (243 >= 32767) {
                            if (float64(3.14159) == -2.25) {
                                for fi15 := 0; fi15 < 17; fi15++ {
                                    v16 := GenS2{F0: p1, F1: 127, F2: 1}
                                    _ = fi15
                                    _ = v16
                                }
                            }
                            v17 := GenS3{F0: "a", F1: p0, F2: 1000}
                            p0 -= p0
                            _ = v17
                        }
                        v18 := []int{32767, -2147483648, -910}
                        var v19 uint8 = 255
                        _ = v18
                        _ = v19
                    }
                    p1 = 0
                } else {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 uint = 42
                        v22 := []int{fi20, (-(fi20)), (fi20 ^ fi20), (fi20 + fi20)}
                        fmt.Println(v21, p0, p0)
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v23 := []int{-599, 9223372036854775807, 692}
                    p0 += 65535
                    _ = v23
                }
            }
            v24 := GenS3{F0: "hello", F1: (^(p0)), F2: -267}
            _ = v24
        } else {
            switch -1 {
            case 127:
                if (true && false) {
                    for fi25 := 0; fi25 < 2; fi25++ {
                        var v26 int16 = 32767
                        v27 := GenS4{F0: (26 >= 1)}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    p1 -= (p1 << ((65535) & 63))
                } else {
                    p0 -= p0
                    p0 = 0
                }
            case 9223372036854775807:
                v28 := GenS2{F0: (p1 + (p1 >> ((0) & 63))), F1: -110, F2: 96}
                _ = v28
            }
        }
        var v29 int16 = -1
        var v30 uint64 = p1
        _ = v29
        _ = v30
    }
    v31 := GenS4{F0: false}
    v32 := []int{127}
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    return "\n\t"
}

func main() {
    v1 := []int{0, -937}
    v2 := []string{"", (("hello" + (("foo bar" + ("hello" + "foo bar")) + ("foo bar" + "hello"))) + "\n\t"), ("a" + "a")}
    var v3 float32 = 0.1
    var v4 uint32 = 42
    v33 := GenS1{F0: 255, F1: -689, F2: -9223372036854775808}
    for fi34 := 0; fi34 < 6; fi34++ {
        for fi35 := 0; fi35 < 3; fi35++ {
            if (3.14159 != (1e10 - 1.5)) {
                var v36 uint = 2
                if len(v1) > 0 {
                    v1[0] = fi34
                }
                _ = v36
            } else {
                var v37 bool = (((1e10 * 0.1) + (0.0 / 0.1)) != 0.1)
                _ = v37
            }
            switch (fi34 ^ fi34) {
            case 0:
                if true {
                    switch (fi35 % ((127) | 1)) {
                    case 127:
                        v38 := GenS5{F0: 100, F1: v4}
                        _ = v38
                    case 128:
                        if true {
                            fmt.Println(v3, v3, fi34)
                        }
                    }
                    var v39 uint32 = v4
                    fmt.Println(fi35)
                    _ = v39
                } else {
                    var v40 uint = 0
                    var v41 uint64 = 1
                    var v42 uint32 = v4
                    _ = v40
                    _ = v41
                    _ = v42
                }
            case -282:
                switch "a" {
                case "a":
                    v3 = v3
                }
            case -1:
                v43 := GenS6{F0: 2, F1: 4294967295}
                _ = v43
            }
            _ = fi35
        }
        _ = fi34
    }
    v4 -= (^(v4))
    var v44 uint8 = 0
    v44 -= v44
    v45 := GenS2{F0: 18446744073709551615, F1: 0, F2: -1}
    v1 = append(v1, 127)
    for fi46 := 0; fi46 < 8; fi46++ {
        switch "foo bar" {
        case "foo bar":
            v3 = v3
        }
        _ = fi46
    }
    if (!(true && false)) {
        var v47 float32 = 0.0
        var v48 uint8 = (v44 - 127)
        _ = v47
        _ = v48
    } else {
        if (!(true || true)) {
            v44 += v44
        }
        v44 = v44
        v3 = ((v3 / v3) / v3)
    }
    fmt.Println(v3, v44, v4)
    for fi49 := 0; fi49 < 5; fi49++ {
        switch fi49 {
        case -9223372036854775808:
            v50 := GenS7{F0: (fi49 - (fi49 << ((uint(v44)) & 63))), F1: fi49, F2: 9223372036854775807}
            _ = v50
        default:
            if (false || (3.14159 == 3.14159)) {
                v44 += 39
                var v51 float64 = float64(v3)
                fmt.Println(fi49, v4)
                _ = v51
            } else {
                for fi52 := 0; fi52 < 12; fi52++ {
                    var v53 int32 = -1
                    _ = fi52
                    _ = v53
                }
                v3 += (v3 * (v3 - (v3 / (1e10 * float32(3.14159)))))
                for fi54 := 0; fi54 < 10; fi54++ {
                    v55 := GenS6{F0: 0, F1: 233}
                    if true {
                        var v56 int32 = 2147483647
                        _ = v56
                    } else {
                        var v57 uint16 = 65535
                        var v58 uint16 = (v57 - v57)
                        _ = v57
                        _ = v58
                    }
                    v44 += 104
                    _ = fi54
                    _ = v55
                }
            }
        }
        _ = fi49
    }
    var v59 uint32 = v4
    v60 := GenS4{F0: (!true)}
    var v61 string = (("\n\t" + ("a" + "foo bar")) + "a")
    for fi62 := 0; fi62 < 15; fi62++ {
        v4 = 2
        v59 -= 350
        _ = fi62
    }
    v4 -= (^(v59))
    for fi63 := 0; fi63 < 2; fi63++ {
        var v64 float32 = 0.0
        var v65 uint = 65535
        _ = fi63
        _ = v64
        _ = v65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v44
    _ = v45
    _ = v59
    _ = v60
    _ = v61
}
