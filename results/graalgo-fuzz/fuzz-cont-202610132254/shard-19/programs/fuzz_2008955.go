package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int }
func genFn0() uint {
    var v7 uint32 = 65535
    v7 += v7
    v8 := GenS1{F0: (-58 >= -1), F1: 65535}
    v9 := []int{42}
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}
func genFn1() int8 {
    var v10 bool = true
    switch "a" {
    case "foo bar":
        var v11 uint16 = 32767
        _ = v11
    case "":
        for fi12 := 0; fi12 < 19; fi12++ {
            switch (fi12 * fi12) {
            case 384:
                if v10 {
                    switch 4294967295 {
                    case 127:
                        var v13 float32 = ((0.0 + 1.5) * (0.1 / 1.0))
                        _ = v13
                    default:
                        switch 1 {
                        case -1:
                            fmt.Println(v10)
                        case 4294967295:
                            v14 := GenS1{F0: (!(-9223372036854775808 <= -780)), F1: 954}
                            _ = v14
                        case 32767:
                            fmt.Println(v10, v10)
                        }
                    }
                    switch "\n\t" {
                    case "\n\t":
                        if (3.14159 > (0.1 * (3.14159 * -2.25))) {
                            v15 := []int{}
                            _ = v15
                        } else {
                            v16 := GenS1{F0: v10, F1: 65535}
                            var v17 uint16 = 1
                            _ = v16
                            _ = v17
                        }
                    }
                } else {
                    switch "" {
                    case "foo bar":
                        v10 = (false && v10)
                    }
                }
            case 32768:
                for fi18 := 0; fi18 < 8; fi18++ {
                    v19 := GenS2{F0: 127}
                    switch (-(fi18)) {
                    case 271:
                        var v20 float64 = 1.5
                        _ = v20
                    default:
                        switch fi12 {
                        case -481:
                            v21 := GenS2{F0: 0}
                            _ = v21
                        case 127:
                            v10 = (float64(0.1) < 1e10)
                            var v22 uint16 = 0
                            _ = v22
                        default:
                            var v23 uint64 = 4294967295
                            v24 := []int{}
                            v23 = (^(v23))
                            _ = v23
                            _ = v24
                        }
                    }
                    var v25 int64 = 997
                    _ = fi18
                    _ = v19
                    _ = v25
                }
            case -2147483648:
                v26 := []int{909, fi12, 128, (fi12 & 2147483647)}
                _ = v26
            }
            _ = fi12
        }
    case "a":
        fmt.Println(v10)
    }
    _ = v10
    return -1
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 int64 = -9223372036854775808
    var v5 int = 618
    v6 := []string{}
    v27 := GenS3{F0: false, F1: (v1 & v1)}
    _ = len(v2)
    delete(v3, 2)
    var v28 float64 = 0.0
    v29 := GenS2{F0: 88}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v30 int16 = -32768
    var v31 uint64 = 42
    if (true && true) {
        v32 := v3[0]
        var v33 uint = 612
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 6; fi34++ {
            for fi35 := 0; fi35 < 19; fi35++ {
                switch fi35 {
                case 9223372036854775807:
                    _ = len(v3)
                case 127:
                    switch fi34 {
                    case -127:
                        v4 = -1
                    }
                case 2147483647:
                    v36 := GenS3{F0: true, F1: genFn1()}
                    _ = v36
                }
                v6 = append(v6, ("hello" + "\n\t"))
                _ = fi35
            }
            var v37 int64 = (v4 ^ int64(v1))
            v38 := v3[464]
            _ = fi34
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 2; fi39++ {
            for fi40 := 0; fi40 < 17; fi40++ {
                v5 = (fi40 * (fi40 | (^(fi40))))
                v28 -= (v28 * v28)
                v5 -= fi40
                var v41 int32 = 2147483647
                var v42 int = (fi40 << ((65535) & 63))
                _ = fi40
                _ = v41
                _ = v42
            }
            switch "" {
            case "foo bar":
                v43 := GenS2{F0: 121}
                _ = v43
            }
            for fi44 := 0; fi44 < 2; fi44++ {
                var v45 int = v5
                _ = fi44
                _ = v45
            }
            _ = fi39
        }
    }
    v46 := v3[84]
    if (726 >= 877) {
        v47 := GenS1{F0: (127 == 0), F1: 255}
        for fi48 := 0; fi48 < 3; fi48++ {
            switch 42 {
            case -9223372036854775808:
                for fi49 := 0; fi49 < 6; fi49++ {
                    fmt.Println(v46, v1, v4)
                    _ = fi49
                }
            case 1:
                var v50 float64 = v28
                _ = v50
            default:
                v51 := GenS3{F0: v46, F1: 100}
                _ = v51
            }
            v52 := GenS4{F0: fi48}
            if true {
                v31 = v31
                fmt.Println(v46)
            }
            _ = fi48
            _ = v52
        }
        _ = v47
    } else {
        fmt.Println(v1, v4)
        v31 -= (v31 | (v31 << ((genFn0()) & 63)))
    }
    if true {
        v28 += v28
        v31 += v31
        var v53 bool = true
        _ = v53
    }
    fmt.Println(v4, v46)
    for fi54 := 0; fi54 < 16; fi54++ {
        v5 += (v5 >> ((1) & 63))
        _ = fi54
    }
    v46 = v46
    var v55 float64 = v28
    v56 := GenS3{F0: false, F1: genFn1()}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v46
    _ = v5
    _ = v55
    _ = v56
    _ = v6
}
