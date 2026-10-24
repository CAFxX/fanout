package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS2 struct { F0 int; F1 uint8; F2 int16 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int32; F1 uint; F2 int16 }
type GenS5 struct { F0 int16; F1 int64; F2 uint16 }
func genFn0(p0 bool, p1 bool) int {
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS1{F0: 139, F1: -1, F2: 336}
            var v10 int = fi8
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    var v11 int64 = 225
    var v12 uint32 = 4294967295
    if (p0 && (-9223372036854775808 == -499)) {
        if (-167 == v11) {
            var v13 uint64 = 0
            _ = v13
        }
        p0 = p1
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 4294967295
}

func main() {
    var v1 bool = true
    var v2 bool = (!true)
    var v3 uint32 = 933
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    var v6 int = int(v3)
    switch ("\n\t" + "") {
    case "hello":
        var v14 int8 = 0
        _ = v14
    }
    switch "foo bar" {
    case "foo bar":
        for fi15 := 0; fi15 < 17; fi15++ {
            for fi16 := 0; fi16 < 14; fi16++ {
                if (uint(v3) != 867) {
                    v1 = v2
                }
                v6 -= (v6 + fi16)
                v17 := GenS1{F0: 191, F1: 42, F2: 781}
                v18 := GenS1{F0: 0, F1: 919, F2: 0}
                _ = fi16
                _ = v17
                _ = v18
            }
            v19 := []int{(^(v6))}
            _ = fi15
            _ = v19
        }
    case "":
        v20 := GenS2{F0: (-(v6)), F1: 14, F2: 32767}
        _ = v20
    case "\n\t":
        for fi21 := 0; fi21 < 5; fi21++ {
            if false {
                v22 := GenS3{F0: ("hello" < "hello"), F1: 825}
                _ = v22
            }
            v1 = (!(0.1 >= 1.5))
            var v23 uint64 = 0
            _ = fi21
            _ = v23
        }
    default:
        var v24 int8 = 100
        _ = v24
    }
    if v1 {
        if (uint8(v3) < 1) {
            var v25 int64 = -120
            _ = v25
        } else {
            if v1 {
                v1 = (255 < 2)
            }
            if (255 <= 1) {
                fmt.Println(v2)
                var v26 uint = 255
                _ = v26
            }
            switch ("foo bar" + "foo bar") {
            case "hello":
                v1 = false
            default:
                delete(v5, v3)
            }
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := GenS4{F0: -1, F1: 18446744073709551615, F2: 32767}
            _ = fi27
            _ = v28
        }
    } else {
        v2 = v1
        var v29 int8 = 82
        _ = v29
    }
    if (1e10 < -2.25) {
        var v30 int = v6
        for fi31 := 0; fi31 < 19; fi31++ {
            if false {
                fmt.Println(fi31)
                v30 += (v6 & (v6 % (((v30 & -1)) | 1)))
                v32 := GenS5{F0: 1, F1: 42, F2: 0}
                _ = v32
            } else {
                v3 = 1
            }
            _ = len(v5)
            var v33 int16 = 42
            _ = fi31
            _ = v33
        }
        _ = v30
    } else {
        for fi34 := 0; fi34 < 6; fi34++ {
            fmt.Println(v1)
            _ = fi34
        }
    }
    if v1 {
        for fi35 := 0; fi35 < 6; fi35++ {
            if v2 {
                v36 := []int{}
                _ = v36
            }
            for fi37 := 0; fi37 < 16; fi37++ {
                if true {
                    switch ("a" + "hello") {
                    case "":
                        if false {
                            var v38 bool = true
                            v2 = (4294967295 != v3)
                            fmt.Println(fi35)
                            _ = v38
                        } else {
                            v6 = 128
                            var v39 int32 = 0
                            _ = v39
                        }
                    case "\n\t":
                        var v40 uint32 = (v3 ^ (v3 ^ (v3 / (((^(v3))) | 1))))
                        _ = v40
                    }
                }
                _ = fi37
            }
            _ = len(v4)
            _ = fi35
        }
        v6 += v6
        fmt.Println(v1, v2, v2)
    } else {
        for fi41 := 0; fi41 < 10; fi41++ {
            if ((511 != 0) && (v2 || v1)) {
                fmt.Println(v6, v2, v3)
                v3 = v3
            } else {
                switch "a" {
                case "":
                    var v42 int8 = 74
                    _ = v42
                }
                fmt.Println(fi41, fi41, v3)
                v43 := GenS5{F0: -1, F1: 42, F2: 255}
                _ = v43
            }
            _ = fi41
        }
        var v44 uint64 = 42
        v1 = (-32768 > 1)
        _ = v44
    }
    v45 := GenS4{F0: 42, F1: 2, F2: 0}
    v3 -= v3
    var v46 bool = ((0.1 - 3.14159) == (((((1.5 * (((float64(-2.25) / 3.14159) + 1e10) + float64(-2.25))) - -2.25) * ((0.0 - (float64(1.5) / 3.14159)) / (0.1 * -2.25))) - 1e10) / 1.5))
    v47 := []int{v6}
    switch v6 {
    case 879:
        v47 = append(v47, v6)
    }
    v4[""] = v3
    if len(v47) > 0 {
        v47[0] = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
