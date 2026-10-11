package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 bool }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint; F1 int; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 int32; F1 uint64; F2 int }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 int8 }
func genFn0(p0 int64) uint8 {
    fmt.Println(p0)
    if true {
        p0 += (^(p0))
    } else {
        v7 := GenS1{F0: 508, F1: 329, F2: true}
        fmt.Println(p0)
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: -1, F1: 2, F2: false}
            v10 := []int{fi8, (fi8 | -959)}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    var v11 uint = 1
    _ = p0
    _ = v11
    return 2
}
func genFn1(p0 string, p1 int16, p2 uint16) int {
    if ((p2 << ((2) & 63)) > (p2 + p2)) {
        switch -425 {
        case 128:
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(p0, fi12, p1)
                var v13 float32 = (1.5 / float32((1.5 + 0.0)))
                _ = fi12
                _ = v13
            }
        case 127:
            var v14 uint64 = 2
            _ = v14
        }
        var v15 uint32 = 234
        fmt.Println(p0)
        _ = v15
    } else {
        for fi16 := 0; fi16 < 9; fi16++ {
            p2 -= (p2 & 32767)
            var v17 uint64 = 42
            for fi18 := 0; fi18 < 19; fi18++ {
                v19 := GenS2{F0: -1, F1: (!true)}
                v20 := GenS2{F0: 315, F1: true}
                switch (p0 + p0) {
                case "":
                    if false {
                        if true {
                            v21 := GenS3{F0: 834, F1: (fi16 & (fi16 - fi18)), F2: fi16}
                            v17 -= 544
                            var v22 int = fi16
                            _ = v21
                            _ = v22
                        }
                    } else {
                        v23 := GenS2{F0: 1, F1: false}
                        var v24 float32 = float32(0.0)
                        _ = v23
                        _ = v24
                    }
                case "foo bar":
                    v17 = 0
                case "hello":
                    if ((fi18 ^ 32768) > 0) {
                        v25 := GenS4{F0: p0}
                        v17 = 18446744073709551615
                        _ = v25
                    }
                default:
                    p0 = "foo bar"
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 string = ((p0 + p0) + p0)
            v29 := []int{fi27, (-(fi27))}
            if false {
                var v30 int32 = 42
                v31 := GenS2{F0: v30, F1: ((!(true && (true || true))) && true)}
                p0 = (p0 + (p0 + v28))
                p1 = (p1 & (p1 | 42))
                var v32 uint8 = genFn0(42)
                var v33 int64 = -951
                _ = v30
                _ = v31
                _ = v32
                _ = v33
            } else {
                var v34 uint = 2
                v35 := GenS5{F0: 1, F1: 418}
                if len(v29) > 0 {
                    v29[0] = -502
                }
                _ = v34
                _ = v35
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 uint64 = 278
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    var v4 int8 = -67
    v5 := make(map[string]string)
    var v6 uint16 = 667
    switch 9223372036854775807 {
    case -9223372036854775808:
        fmt.Println(v4)
    }
    fmt.Println(v4, v6, v1)
    v1 -= uint64(v6)
    v36 := GenS1{F0: -484, F1: (v1 & v1), F2: true}
    v1 = 785
    for fi37 := 0; fi37 < 18; fi37++ {
        for fi38 := 0; fi38 < 17; fi38++ {
            var v39 uint64 = 18446744073709551615
            if ((fi38 >> ((395) & 63)) != fi38) {
                switch genFn1("a", 814, (v6 % ((255) | 1))) {
                case -143:
                    if true {
                        v40 := GenS3{F0: 1, F1: genFn1(("\n\t" + "a"), 1000, (v6 ^ 255)), F2: genFn1("\n\t", -743, (v6 + (v6 & (v6 * (v6 - 91)))))}
                        v41 := GenS2{F0: 42, F1: true}
                        _ = v40
                        _ = v41
                    } else {
                        var v42 uint64 = (^(v1))
                        _ = v42
                    }
                case -2147483648:
                    var v43 int8 = (^(v4))
                    _ = v43
                }
                _ = len(v5)
                v6 = (v6 << ((18446744073709551615) & 63))
                var v44 int = fi38
                v4 += (v4 / ((v4) | 1))
                _ = v44
            }
            _ = fi38
            _ = v39
        }
        switch (("hello" + "foo bar") + ("" + "")) {
        case "hello":
            v4 += v4
        default:
            v6 += v6
        }
        _ = fi37
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        v46 := GenS6{F0: 42, F1: (v1 | v1), F2: fi45}
        _ = fi45
        _ = v46
    }
    v47 := GenS1{F0: 42, F1: (v1 << ((18446744073709551615) & 63)), F2: (2147483647 > 42)}
    delete(v5, "")
    var v48 uint16 = v6
    var v49 string = (("\n\t" + "hello") + ("" + "\n\t"))
    if true {
        var v50 int16 = 0
        v4 = (v4 | 127)
        _ = v50
    } else {
        switch 0 {
        case 42:
            fmt.Println(v4, v1, v49)
        case 32767:
            switch -137 {
            case 42:
                v4 = (v4 + v4)
            case 32767:
                v6 = (v48 * (v48 << ((856) & 63)))
            case -9223372036854775808:
                v49 = "hello"
            default:
                v1 -= 305
            }
        default:
            v49 = "a"
        }
    }
    v51 := GenS1{F0: 2147483647, F1: v1, F2: ((1 > 32767) && ((genFn1(v49, 1000, v48) > -851) || false))}
    fmt.Println(v48, v4)
    for fi52 := 0; fi52 < 2; fi52++ {
        for fi53 := 0; fi53 < 6; fi53++ {
            var v54 bool = false
            fmt.Println(v4, fi53)
            switch v49 {
            case "\n\t":
                var v55 uint32 = 0
                _ = v55
            case "":
                if (((v1 & v1) != v1) || false) {
                    switch fi52 {
                    case 7:
                        if v54 {
                            if true {
                                var v56 int64 = 0
                                var v57 int32 = 0
                                _ = len(v5)
                                var v58 int = 4294967295
                                v59 := v3[(fi53 + (fi52 * 128))]
                                _ = v56
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                        }
                    }
                } else {
                    v60 := GenS7{F0: v49, F1: v1}
                    _ = v60
                }
            }
            _ = fi53
            _ = v54
        }
        v61 := GenS8{F0: -125}
        _ = len(v2)
        _ = fi52
        _ = v61
    }
    if (185 != 2) {
        for fi62 := 0; fi62 < 7; fi62++ {
            switch 1 {
            case 32768:
                var v63 int16 = 0
                _ = v63
            case 716:
                if true {
                    v64 := []int{}
                    if (389 < 643) {
                        var v65 bool = (false || (false || false))
                        var v66 int = (-(fi62))
                        fmt.Println(v6)
                        _ = v65
                        _ = v66
                    }
                    _ = v64
                }
            case 9223372036854775807:
                switch (fi62 / ((fi62) | 1)) {
                case -147:
                    var v67 int = (^(fi62))
                    _ = v67
                case 42:
                    v68 := []int{}
                    _ = v68
                }
            }
            var v69 uint16 = 65535
            v1 -= (^(v1))
            _ = fi62
            _ = v69
        }
        var v70 uint8 = genFn0(1)
        v70 = 128
        _ = v70
    } else {
        var v71 int32 = int32(v1)
        v72 := GenS3{F0: 4294967295, F1: 1, F2: 4294967295}
        _ = v71
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v51
    _ = v6
}
