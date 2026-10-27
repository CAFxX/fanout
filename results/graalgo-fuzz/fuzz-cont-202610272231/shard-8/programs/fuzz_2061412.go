package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int8; F1 string; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 string = ("" + "hello")
    v3 := make(map[uint32]string)
    v4 := make(map[int]bool)
    var v5 int16 = 32767
    var v6 string = (("foo bar" + (v2 + (v2 + v2))) + (v2 + v2))
    v5 = 0
    var v7 uint8 = 255
    if true {
        fmt.Println(v5, v5, v5)
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                if (1e10 >= (1.5 * (3.14159 + ((1e10 - (0.0 - 1e10)) - ((3.14159 - (-2.25 * (-2.25 * 1.5))) - (1.5 * (-2.25 + 1e10))))))) {
                    _ = len(v1)
                    switch v2 {
                    case "foo bar":
                        v2 = v6
                    default:
                        fmt.Println(v6, v7)
                    }
                    if (fi8 > 32767) {
                        fmt.Println(v2, v2, v7)
                        switch -249 {
                        case -547:
                            v10 := v4[32768]
                            _ = v10
                        case -60:
                            v5 -= (v5 | v5)
                        default:
                            v6 = v2
                        }
                        if true {
                            v11 := GenS1{F0: (fi9 | 32767)}
                            if (42 != 1) {
                                v5 = v5
                                var v12 int32 = 0
                                _ = v12
                            }
                            _ = v11
                        }
                        fmt.Println(v6, v2)
                    }
                    var v13 uint16 = uint16(v7)
                    v14 := GenS2{F0: 1, F1: fi8}
                    switch v2 {
                    case "foo bar":
                        v6 = v2
                    case "":
                        v15 := []int{}
                        _ = v15
                    default:
                        var v16 uint32 = 724
                        _ = v16
                    }
                    _ = v13
                    _ = v14
                }
                var v17 uint8 = 27
                _ = fi9
                _ = v17
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                if (42 != 782) {
                    var v19 uint64 = 4294967295
                    v7 = 244
                    switch v2 {
                    case "hello":
                        var v20 int = 462
                        _ = v20
                    }
                    var v21 int8 = 42
                    _ = v19
                    _ = v21
                }
                _ = fi18
            }
            _ = fi8
        }
        v22 := []int{4294967295, 127, 127}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 13; fi23++ {
            fmt.Println(v5)
            var v24 int16 = (v5 + (v5 - (^(v5))))
            switch 9223372036854775807 {
            case 42:
                for fi25 := 0; fi25 < 5; fi25++ {
                    v26 := GenS3{F0: 128, F1: ((false || true) || false)}
                    _ = fi25
                    _ = v26
                }
            case 127:
                var v27 uint8 = (^(v7))
                _ = v27
            }
            _ = fi23
            _ = v24
        }
    }
    v1[-512] = ""
    if true {
        if true {
            var v28 uint8 = 42
            v29 := GenS3{F0: int(v7), F1: (v5 != (v5 & v5))}
            _ = v28
            _ = v29
        } else {
            if true {
                v6 = ((((v2 + v6) + ("" + v6)) + (v6 + "\n\t")) + v2)
                fmt.Println(v5, v5, v5)
                var v30 uint8 = 127
                _ = v30
            } else {
                v31 := v3[1]
                _ = v31
            }
        }
    } else {
        var v32 int32 = -928
        _ = v32
    }
    if false {
        if false {
            v5 -= (v5 >> ((42) & 63))
            var v33 bool = (!true)
            if (-1 > 0) {
                switch -885 {
                case -2147483648:
                    for fi34 := 0; fi34 < 6; fi34++ {
                        v7 += v7
                        fmt.Println(v7, v6)
                        v35 := GenS1{F0: fi34}
                        _ = fi34
                        _ = v35
                    }
                case 42:
                    v36 := []int{1}
                    _ = v36
                case 304:
                    v5 += (^(v5))
                }
            } else {
                var v37 string = v2
                _ = v37
            }
            _ = v33
        } else {
            if false {
                var v38 uint32 = 294
                for fi39 := 0; fi39 < 6; fi39++ {
                    v5 = (v5 - (v5 % (((v5 + 42)) | 1)))
                    switch "" {
                    case "\n\t":
                        fmt.Println(v6, v7)
                    }
                    v40 := GenS1{F0: fi39}
                    _ = fi39
                    _ = v40
                }
                fmt.Println(v7, v38, v2)
                _ = v38
            } else {
                var v41 uint16 = 42
                _ = v41
            }
            v42 := GenS4{F0: 42, F1: "a", F2: v5}
            _ = v42
        }
        if (false && (false && true)) {
            v2 = ("a" + ((v2 + "foo bar") + v2))
            v5 += v5
            var v43 uint16 = 255
            _ = v43
        }
    } else {
        for fi44 := 0; fi44 < 15; fi44++ {
            switch (("a" + "\n\t") + "foo bar") {
            case "":
                v45 := v4[-568]
                _ = v45
            case "hello":
                delete(v3, 745)
            case "a":
                for fi46 := 0; fi46 < 20; fi46++ {
                    if (!false) {
                        fmt.Println(v7, fi46, v7)
                        if false {
                            switch (v2 + (("hello" + (v6 + "a")) + v6)) {
                            case "":
                                var v47 string = (v6 + (v2 + (("\n\t" + "hello") + "\n\t")))
                                _ = v47
                            case "\n\t":
                                v2 = ""
                                var v48 float32 = -2.25
                                v5 += (v5 / ((427) | 1))
                                _ = v48
                            case "a":
                                v49 := GenS1{F0: (fi46 + fi44)}
                                _ = v49
                            default:
                                v6 = v6
                            }
                            if true {
                                var v50 bool = true
                                v5 += (^(v5))
                                var v51 float32 = 1.5
                                var v52 int64 = 875
                                var v53 float64 = 3.14159
                                _ = v50
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                            v6 = "foo bar"
                        }
                        v2 = v2
                    } else {
                        fmt.Println(fi46)
                        switch v6 {
                        case "a":
                            if (fi44 > fi46) {
                                var v54 uint32 = 2147483647
                                _ = v54
                            }
                        }
                        if true {
                            v7 = 115
                            fmt.Println(v7)
                            var v55 int8 = 42
                            var v56 float32 = ((3.14159 * 1e10) - 0.1)
                            _ = v55
                            _ = v56
                        }
                    }
                    fmt.Println(fi44, v5)
                    fmt.Println(fi44, fi44, fi44)
                    _ = fi46
                }
            default:
                var v57 bool = true
                _ = v57
            }
            _ = fi44
        }
        v58 := []int{-1, 32767}
        _ = len(v1)
        _ = v58
    }
    fmt.Println(v2, v5)
    v5 += 744
    var v59 uint64 = 0
    var v60 int32 = 0
    switch 0 {
    case 578:
        v2 = v6
    case 0:
        var v61 int = 128
        _ = v61
    case 391:
        fmt.Println(v2)
    default:
        v62 := []int{2147483647, -9223372036854775808}
        _ = v62
    }
    v4[32767] = false
    if false {
        fmt.Println(v59, v60)
        fmt.Println(v60, v5)
        var v63 float32 = (-2.25 + (-2.25 / ((float32(3.14159) - (((1e10 - 0.1) - 1e10) * 0.1)) - 1.5)))
        _ = v63
    }
    v64 := GenS4{F0: -128, F1: "a", F2: v5}
    var v65 bool = false
    v59 += uint64(v7)
    var v66 uint8 = (v7 + v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v64
    _ = v65
    _ = v66
    _ = v7
}
