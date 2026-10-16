package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 string; F2 bool }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint; F1 int8; F2 bool }
type GenS6 struct { F0 int; F1 uint32; F2 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int; F1 int8 }
func genFn0(p0 int8) string {
    v7 := GenS1{F0: 0}
    v8 := []int{}
    for fi9 := 0; fi9 < 10; fi9++ {
        switch ((("foo bar" + "\n\t") + (("" + "a") + "")) + ("a" + "\n\t")) {
        case "a":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return ("" + "foo bar")
}
func genFn1(p0 int32) uint32 {
    fmt.Println(p0)
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            if false {
                p0 -= (-(p0))
            } else {
                var v12 int16 = 1000
                if (((true && false) || true) || (fi10 != fi10)) {
                    var v13 uint16 = 42
                    switch -9223372036854775808 {
                    case 128:
                        var v14 float64 = (3.14159 / 3.14159)
                        _ = v14
                    case 0:
                        if (((3.14159 - 0.0) / 1e10) > (3.14159 / (0.1 - float32(-2.25)))) {
                            v15 := GenS1{F0: (v12 % ((v12) | 1))}
                            var v16 uint = 18446744073709551615
                            v13 += (v13 - 255)
                            _ = v15
                            _ = v16
                        } else {
                            switch "hello" {
                            case "\n\t":
                                v12 += (v12 | (v12 << ((211) & 63)))
                            case "":
                                var v17 uint32 = 1
                                _ = v17
                            }
                            var v18 float64 = (3.14159 / -2.25)
                            _ = v18
                        }
                        p0 = (-(p0))
                        if (2 >= 4294967295) {
                            var v19 int8 = 0
                            v13 = (v13 / ((v13) | 1))
                            if (1 > 42) {
                                v20 := GenS1{F0: v12}
                                if false {
                                    v13 = v13
                                    v13 = 255
                                    fmt.Println(v12)
                                    v21 := GenS2{F0: 2, F1: "\n\t", F2: (false && (!true))}
                                    var v22 int64 = 0
                                    v23 := GenS3{F0: (p0 | (p0 | 661)), F1: (2 > 1)}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                } else {
                                    switch fi10 {
                                    case 4294967295:
                                        v24 := []int{fi10, fi11, fi10}
                                        _ = v24
                                    }
                                    v25 := GenS2{F0: 0, F1: "hello", F2: (!(true || (!false)))}
                                    v26 := []int{(fi10 | (fi10 << ((18446744073709551615) & 63)))}
                                    _ = v25
                                    _ = v26
                                }
                                _ = v20
                            }
                            _ = v19
                        }
                    }
                    fmt.Println(fi10)
                    fmt.Println(p0, p0, fi10)
                    v13 -= 42
                    v27 := GenS4{F0: 32767, F1: int64(v13)}
                    _ = v13
                    _ = v27
                }
                if true {
                    v28 := []int{}
                    if false {
                        v12 = (v12 | v12)
                        var v29 int = (fi10 - fi10)
                        _ = v29
                    } else {
                        fmt.Println(p0, v12, fi10)
                        var v30 int32 = -908
                        p0 += v30
                        _ = v30
                    }
                    var v31 uint64 = 4294967295
                    _ = v28
                    _ = v31
                } else {
                    v32 := GenS5{F0: 42, F1: 127, F2: (703 >= 1)}
                    var v33 uint8 = 164
                    v34 := []int{}
                    _ = v32
                    _ = v33
                    _ = v34
                }
                _ = v12
            }
            v35 := GenS1{F0: -32768}
            _ = fi11
            _ = v35
        }
        _ = fi10
    }
    _ = p0
    return 636
}

func main() {
    var v1 bool = (42 <= 0)
    v2 := []uint32{4294967295}
    v3 := make(map[uint32]uint32)
    v4 := []string{"", "hello", "a", ("a" + ("hello" + "a"))}
    v5 := []int{-156, 319}
    v6 := []string{}
    for fi36 := 0; fi36 < 5; fi36++ {
        if ((v1 && v1) || false) {
            for fi37 := 0; fi37 < 3; fi37++ {
                fmt.Println(v1)
                var v38 uint = 4294967295
                _ = fi37
                _ = v38
            }
            if (42 <= 635) {
                delete(v3, 2147483648)
                switch ((("\n\t" + "\n\t") + "hello") + "\n\t") {
                case "":
                    if ((!v1) || (false && (v1 || (2 == 42)))) {
                        var v39 bool = (-128 == 42)
                        var v40 uint8 = 61
                        _ = v39
                        _ = v40
                    }
                case "\n\t":
                    v1 = v1
                case "hello":
                    for fi41 := 0; fi41 < 3; fi41++ {
                        var v42 float32 = 0.0
                        v43 := GenS3{F0: 2147483647, F1: true}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                default:
                    for fi44 := 0; fi44 < 14; fi44++ {
                        _ = len(v3)
                        v1 = v1
                        _ = fi44
                    }
                }
                var v45 bool = true
                _ = v45
            } else {
                fmt.Println(v1, fi36)
                _ = len(v3)
                for fi46 := 0; fi46 < 18; fi46++ {
                    v47 := v3[0]
                    fmt.Println(fi36, v1, v47)
                    _ = fi46
                    _ = v47
                }
            }
        } else {
            switch ("" + "foo bar") {
            case "":
                var v48 int64 = -9223372036854775808
                _ = v48
            }
            var v49 int16 = 1
            switch "a" {
            case "\n\t":
                var v50 uint16 = 42
                _ = v50
            case "hello":
                if (2 != genFn1(209)) {
                    var v51 uint8 = 2
                    fmt.Println(v1, v1)
                    _ = v51
                }
            default:
                var v52 int64 = -914
                _ = v52
            }
            _ = v49
        }
        v1 = v1
        v1 = ("" != ("" + ("foo bar" + "a")))
        _ = fi36
    }
    var v53 uint = 255
    for fi54 := 0; fi54 < 16; fi54++ {
        for fi55 := 0; fi55 < 14; fi55++ {
            if (1 < -127) {
                v53 -= 215
                v56 := GenS6{F0: fi55, F1: 825, F2: false}
                _ = v56
            }
            var v57 int = (fi55 / ((fi55) | 1))
            if v1 {
                if v1 {
                    switch "foo bar" {
                    case "foo bar":
                        if v1 {
                            var v58 float64 = (0.0 + ((0.0 * ((-2.25 - 0.0) * 1.5)) / -2.25))
                            _ = v58
                        }
                    case "\n\t":
                        v59 := GenS7{F0: 0}
                        _ = v59
                    }
                    var v60 uint8 = 255
                    v61 := GenS5{F0: uint(v60), F1: 42, F2: true}
                    _ = v60
                    _ = v61
                }
                var v62 float32 = (3.14159 * float32(0.0))
                _ = v62
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v63 := GenS8{F0: (v53 - v53)}
                _ = v63
            }
            var v64 uint64 = 1
            _ = fi55
            _ = v57
            _ = v64
        }
        v1 = (!v1)
        v3[2147483647] = 386
        _ = fi54
    }
    for fi65 := 0; fi65 < 9; fi65++ {
        var v66 string = "hello"
        v67 := GenS9{F0: fi65}
        fmt.Println(v66)
        _ = fi65
        _ = v66
        _ = v67
    }
    v1 = (v1 && true)
    v1 = v1
    if v1 {
        var v68 int8 = -109
        var v69 string = ("\n\t" + ("a" + "hello"))
        if (v53 > (v53 + v53)) {
            var v70 uint16 = 1
            switch (v69 + "a") {
            case "a":
                if (42 >= 4294967295) {
                    fmt.Println(v68)
                    fmt.Println(v70)
                    for fi71 := 0; fi71 < 20; fi71++ {
                        fmt.Println(v1, v1)
                        _ = fi71
                    }
                }
            case "foo bar":
                v3[1] = 831
            case "":
                for fi72 := 0; fi72 < 16; fi72++ {
                    var v73 int8 = (v68 << ((576) & 63))
                    _ = fi72
                    _ = v73
                }
            }
            _ = v70
        }
        _ = v68
        _ = v69
    }
    v53 = v53
    switch -98 {
    case 42:
        for fi74 := 0; fi74 < 9; fi74++ {
            fmt.Println(fi74, fi74, v53)
            _ = fi74
        }
    }
    var v75 bool = (1 < -1)
    v75 = (!v75)
    v76 := GenS3{F0: 42, F1: v1}
    var v77 uint64 = 2
    fmt.Println(v75, v75, v53)
    if (!true) {
        v53 = v53
        switch "" {
        case "\n\t":
            v53 = (^(v53))
        case "hello":
            v78 := GenS8{F0: v53}
            _ = v78
        case "a":
            if (!v1) {
                if (42 == 9223372036854775807) {
                    v77 -= (v77 | v77)
                    fmt.Println(v53, v77, v75)
                    for fi79 := 0; fi79 < 9; fi79++ {
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                        _ = fi79
                    }
                } else {
                    switch 4294967295 {
                    case 2147483647:
                        for fi80 := 0; fi80 < 5; fi80++ {
                            if true {
                                if false {
                                    v81 := GenS1{F0: -207}
                                    v53 += (v53 / ((944) | 1))
                                    v82 := v3[225]
                                    _ = v81
                                    _ = v82
                                }
                            } else {
                                var v83 bool = v1
                                if false {
                                    fmt.Println(v1)
                                    var v84 float32 = ((3.14159 - -2.25) + 1e10)
                                    v85 := GenS10{F0: "\n\t"}
                                    _ = v84
                                    _ = v85
                                }
                                _ = v83
                            }
                            _ = fi80
                        }
                    case 466:
                        v86 := GenS9{F0: -9223372036854775808}
                        _ = v86
                    case 32768:
                        for fi87 := 0; fi87 < 19; fi87++ {
                            if (0 > 491) {
                                v88 := GenS11{F0: fi87, F1: 0}
                                _ = v88
                            }
                            for fi89 := 0; fi89 < 8; fi89++ {
                                var v90 uint16 = 255
                                delete(v3, genFn1(2147483647))
                                var v91 bool = (588 == v53)
                                _ = fi89
                                _ = v90
                                _ = v91
                            }
                            _ = fi87
                        }
                    default:
                        for fi92 := 0; fi92 < 16; fi92++ {
                            if v1 {
                                var v93 string = (("a" + "a") + "foo bar")
                                _ = v93
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi92
                        }
                    }
                    var v94 int64 = 219
                    var v95 int64 = v94
                    _ = v94
                    _ = v95
                }
                if v1 {
                    var v96 uint = 106
                    v77 += (v77 % ((1) | 1))
                    _ = v96
                } else {
                    for fi97 := 0; fi97 < 11; fi97++ {
                        switch fi97 {
                        case -640:
                            v75 = (v75 || (!(!(v75 || v1))))
                        }
                        _ = fi97
                    }
                }
                v98 := GenS5{F0: v53, F1: 104, F2: v75}
                _ = v98
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v75
    _ = v76
    _ = v77
}
