package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -110
    v3 := []uint32{2, 42}
    var v4 uint64 = 4294967295
    switch "a" {
    case "\n\t":
        fmt.Println(v4)
    case "hello":
        v5 := GenS1{F0: 935, F1: 42}
        _ = v5
    default:
        fmt.Println(v2, v2)
    }
    v4 = v4
    v6 := GenS2{F0: v4}
    v2 -= v2
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint8 = 2
        for fi9 := 0; fi9 < 20; fi9++ {
            v2 -= (v2 + v2)
            var v10 uint = 699
            _ = fi9
            _ = v10
        }
        var v11 string = ("hello" + "foo bar")
        _ = fi7
        _ = v11
        _ = v8
    }
    switch ((("" + "hello") + ("foo bar" + "hello")) + "\n\t") {
    case "\n\t":
        for fi12 := 0; fi12 < 11; fi12++ {
            fmt.Println(v2)
            _ = fi12
        }
    case "":
        if (!(9 <= 1)) {
            if true {
                var v13 float32 = (float32(-2.25) * (0.0 * float32(1e10)))
                v13 -= 1e10
                _ = v13
            } else {
                v2 -= (v2 * v2)
            }
            var v14 uint64 = (v4 / ((v4) | 1))
            fmt.Println(v2, v4, v2)
            _ = v14
        }
    }
    fmt.Println(v2, v4)
    if (0 == 255) {
        fmt.Println(v2, v4)
        v15 := GenS1{F0: 1, F1: 0}
        v2 = (v2 & (v2 & (v2 ^ 99)))
        _ = v15
    }
    v16 := GenS1{F0: -1, F1: -1}
    v2 = v2
    v17 := GenS3{F0: uint16(v4), F1: 0, F2: 42}
    _ = len(v1)
    var v18 uint32 = 2147483648
    fmt.Println(v18)
    for fi19 := 0; fi19 < 16; fi19++ {
        if ((v4 << ((702) & 63)) != v4) {
            _ = len(v3)
        } else {
            var v20 uint16 = 255
            for fi21 := 0; fi21 < 7; fi21++ {
                fmt.Println(fi21, v4)
                switch "foo bar" {
                case "\n\t":
                    if false {
                        v22 := v1[fi19]
                        v2 -= (-(v2))
                        if (0 <= -67) {
                            var v23 uint16 = v20
                            var v24 uint = 532
                            _ = v23
                            _ = v24
                        }
                        var v25 int8 = 127
                        fmt.Println(v25)
                        var v26 uint8 = uint8(v22)
                        if (805 != uint(v26)) {
                            if (!(1 > 223)) {
                                var v27 int = fi19
                                if false {
                                    fmt.Println(fi21, fi19)
                                    var v28 uint32 = (v18 * 324)
                                    var v29 int8 = v25
                                    _ = v28
                                    _ = v29
                                }
                                _ = v27
                            } else {
                                v25 = -27
                                var v30 string = ""
                                switch (fi21 | 32768) {
                                case 0:
                                    if false {
                                        delete(v1, fi21)
                                        if (("hello" + "\n\t") >= ((v30 + v30) + v30)) {
                                            v26 += (v26 << ((42) & 63))
                                            var v31 int64 = 0
                                            var v32 int32 = 2147483647
                                            if true {
                                                v4 -= 0
                                                fmt.Println(fi19, fi21, v30)
                                                if ((((v30 + (v30 + "foo bar")) + "") + "a") < (v30 + v30)) {
                                                    fmt.Println(fi19)
                                                }
                                                v33 := GenS1{F0: -32768, F1: v31}
                                                _ = v33
                                            } else {
                                                var v34 uint8 = (^(v26))
                                                if len(v3) > 0 {
                                                    v3[0] = v18
                                                }
                                                _ = v34
                                            }
                                            _ = v31
                                            _ = v32
                                        } else {
                                            switch ("\n\t" + v30) {
                                            case "\n\t":
                                                v18 = (v22 * 987)
                                            case "":
                                                var v35 string = ("\n\t" + v30)
                                                _ = v35
                                            }
                                        }
                                    }
                                default:
                                    switch fi19 {
                                    case -982:
                                        var v36 int = 42
                                        _ = v36
                                    }
                                }
                                _ = v30
                            }
                            v22 += (v22 & v18)
                        }
                        v37 := GenS4{F0: v26}
                        switch (fi19 ^ (fi19 ^ 9223372036854775807)) {
                        case -48:
                            var v38 string = ""
                            _ = v38
                        case 2147483647:
                            switch ("a" + "foo bar") {
                            case "\n\t":
                                _ = len(v1)
                            case "":
                                if false {
                                    var v39 uint32 = 4294967295
                                    fmt.Println(v22)
                                    _ = v39
                                } else {
                                    fmt.Println(fi19)
                                }
                            }
                        }
                        fmt.Println(fi21, v25, fi21)
                        _ = v22
                        _ = v25
                        _ = v26
                        _ = v37
                    } else {
                        v20 = v20
                        var v40 float64 = -2.25
                        _ = v40
                    }
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "hello":
                    var v41 uint32 = v18
                    v3 = append(v3, (v18 % (((v41 & (^(v18)))) | 1)))
                    _ = v41
                }
                _ = fi21
            }
            var v42 uint16 = (v20 | v20)
            _ = v20
            _ = v42
        }
        v1[-112] = 65535
        _ = fi19
    }
    var v43 uint8 = 2
    for fi44 := 0; fi44 < 6; fi44++ {
        for fi45 := 0; fi45 < 2; fi45++ {
            v18 -= 1
            _ = fi45
        }
        if true {
            fmt.Println(v18, v2)
            for fi46 := 0; fi46 < 15; fi46++ {
                v18 += v18
                var v47 uint8 = (v43 * (v43 >> ((1) & 63)))
                fmt.Println(fi44)
                var v48 uint8 = 255
                v48 = (v47 % (((^(v47))) | 1))
                if ("\n\t" != "") {
                    var v49 uint32 = 65535
                    _ = v49
                } else {
                    v43 += v47
                    switch fi46 {
                    case 127:
                        v50 := GenS5{F0: v4}
                        _ = v50
                    default:
                        var v51 int64 = 0
                        _ = v51
                    }
                    switch (^(fi44)) {
                    case 32768:
                        fmt.Println(v18)
                    case -669:
                        v52 := GenS6{F0: uint(v47), F1: (v2 & v2)}
                        _ = v52
                    }
                    var v53 float64 = (0.1 - float64(0.0))
                    v47 = (v48 ^ (v43 * (^(v43))))
                    _ = v53
                }
                fmt.Println(v18)
                _ = fi46
                _ = v47
                _ = v48
            }
        }
        var v54 bool = true
        _ = fi44
        _ = v54
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v6
}
