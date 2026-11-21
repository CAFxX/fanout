package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int64, p1 uint64) int64 {
    v6 := []int{1}
    var v7 uint16 = 1
    if true {
        v8 := GenS1{F0: 543}
        switch "" {
        case "foo bar":
            var v9 uint32 = 2147483647
            _ = v9
        default:
            if false {
                v10 := GenS2{F0: p1, F1: 521, F2: "hello"}
                v11 := GenS3{F0: 4294967295}
                var v12 uint8 = 0
                _ = v10
                _ = v11
                _ = v12
            }
        }
        _ = v8
    }
    if (2 > 294) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if (0.1 > (-2.25 / 3.14159)) {
            var v13 bool = ((42 == -1) || (false || false))
            _ = v13
        } else {
            var v14 string = ("foo bar" + "hello")
            if false {
                if false {
                    if true {
                        var v15 float32 = (3.14159 - 0.1)
                        fmt.Println(v7, v7)
                        var v16 string = v14
                        _ = v15
                        _ = v16
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    v7 -= (^(v7))
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                p0 -= p0
            } else {
                if (!((p0 << ((0) & 63)) == p0)) {
                    v7 = (v7 >> ((1) & 63))
                    v14 = ""
                } else {
                    switch v14 {
                    case "hello":
                        var v17 string = ("hello" + "hello")
                        _ = v17
                    case "a":
                        var v18 uint8 = 10
                        _ = v18
                    case "":
                        v19 := GenS2{F0: p1, F1: 255, F2: v14}
                        _ = v19
                    default:
                        var v20 int16 = 1000
                        _ = v20
                    }
                }
                v7 = v7
                if true {
                    var v21 uint32 = 1
                    if true {
                        v21 -= 1
                        var v22 float32 = 3.14159
                        v22 += -2.25
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 19; fi23++ {
                            var v24 float64 = 0.0
                            _ = fi23
                            _ = v24
                        }
                        var v25 uint = 1
                        _ = v25
                    }
                    _ = v21
                }
            }
            _ = v14
        }
        var v26 uint = 135
        _ = v26
    } else {
        var v27 int = 32768
        _ = v27
    }
    var v28 uint64 = (p1 * 42)
    _ = p0
    _ = p1
    _ = v28
    _ = v6
    _ = v7
    return p0
}
func genFn1() uint64 {
    var v29 float64 = 3.14159
    fmt.Println(v29)
    v29 = v29
    for fi30 := 0; fi30 < 1; fi30++ {
        v31 := []int{(fi30 + fi30), 0}
        _ = fi30
        _ = v31
    }
    _ = v29
    return 1
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]string)
    var v3 uint8 = 2
    var v4 bool = true
    v5 := make(map[string]bool)
    switch 0 {
    case 9223372036854775807:
        if false {
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 int16 = v1
                _ = fi32
                _ = v33
            }
        } else {
            v34 := GenS3{F0: 42}
            for fi35 := 0; fi35 < 1; fi35++ {
                for fi36 := 0; fi36 < 5; fi36++ {
                    _ = len(v2)
                    v37 := []int{(fi36 ^ (fi36 / ((int(v1)) | 1))), fi35}
                    fmt.Println(fi36, v4)
                    _ = fi36
                    _ = v37
                }
                _ = fi35
            }
            _ = v34
        }
    case 6:
        var v38 int32 = -454
        _ = v38
    case 32767:
        fmt.Println(v3, v4, v1)
    default:
        v2[63] = "foo bar"
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        v4 = ((-2.25 / 1.0) == (1.5 - float32(-2.25)))
        if (633 < 2) {
            v40 := GenS4{F0: 0, F1: 933, F2: genFn1()}
            v41 := []int{0, (fi39 + fi39), fi39, fi39}
            switch (fi39 ^ fi39) {
            case -988:
                fmt.Println(v4, v3, fi39)
            case 775:
                v42 := GenS5{F0: 60}
                _ = v42
            case 128:
                v4 = (!(2 == 0))
            }
            _ = v40
            _ = v41
        }
        v43 := []int{(fi39 * (fi39 & fi39))}
        _ = fi39
        _ = v43
    }
    fmt.Println(v1, v4)
    for fi44 := 0; fi44 < 13; fi44++ {
        var v45 int = (fi44 ^ (fi44 / ((fi44) | 1)))
        var v46 int32 = int32(v3)
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 uint16 = 2
    v1 += v1
    v47 -= (v47 ^ uint16(v3))
    v1 += (v1 | v1)
    v48 := GenS6{F0: v3, F1: (v3 * 242)}
    switch 32768 {
    case -9223372036854775808:
        for fi49 := 0; fi49 < 10; fi49++ {
            var v50 int64 = 352
            _ = fi49
            _ = v50
        }
    }
    var v51 string = "foo bar"
    for fi52 := 0; fi52 < 2; fi52++ {
        v53 := []int{(fi52 << ((1) & 63))}
        _ = fi52
        _ = v53
    }
    v4 = (2 <= 118)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v51
}
