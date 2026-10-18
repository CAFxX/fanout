package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint64 }
func genFn0() uint {
    var v5 int16 = 42
    v5 += 32767
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{}
        var v8 uint64 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = v5
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 65535
    var v2 float32 = 1e10
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    var v9 int64 = 42
    var v10 float32 = (v2 - float32(-2.25))
    var v11 uint64 = 0
    v12 := []int{int(v9), -9223372036854775808, -2147483648}
    var v13 uint8 = 205
    v14 := GenS1{F0: 42, F1: 491, F2: 0}
    v13 = v13
    v2 += 3.14159
    var v15 int16 = 1
    switch "hello" {
    case "\n\t":
        v16 := GenS1{F0: uint(v1), F1: 128, F2: v1}
        _ = v16
    case "foo bar":
        if (v9 == (v9 % (((v9 % ((1) | 1))) | 1))) {
            var v17 uint = 65535
            v10 += (v10 * float32(float64(((v10 - float32(float64(v2))) / (float32(1.5) - v2)))))
            var v18 uint64 = v11
            _ = v17
            _ = v18
        }
    }
    switch -1 {
    case -291:
        if true {
            if ((("a" + "hello") < ((("\n\t" + "a") + "hello") + "a")) || true) {
                for fi19 := 0; fi19 < 11; fi19++ {
                    if false {
                        fmt.Println(v2, v2, v2)
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    }
                    var v20 int16 = v15
                    _ = fi19
                    _ = v20
                }
                var v21 float64 = 0.0
                if (!(594 == 559)) {
                    v3[""] = true
                    v22 := GenS2{F0: v1, F1: 0, F2: (v9 & (-(v9)))}
                    v21 = (((0.1 * v21) - v21) * v21)
                    _ = v22
                }
                _ = v21
            } else {
                var v23 int64 = 51
                _ = v23
            }
            switch -2147483648 {
            case 32767:
                for fi24 := 0; fi24 < 7; fi24++ {
                    v25 := GenS1{F0: 65535, F1: fi24, F2: (v1 | 0)}
                    var v26 int = fi24
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            case -9223372036854775808:
                v27 := GenS2{F0: 42, F1: (v11 & (v11 << ((1) & 63))), F2: (^(v9))}
                _ = v27
            default:
                fmt.Println(v2)
            }
            v28 := GenS3{F0: uint64(v13), F1: "hello"}
            _ = v28
        } else {
            v29 := GenS1{F0: 0, F1: int(v11), F2: v1}
            _ = v29
        }
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 string = "a"
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 float64 = (((1e10 - (((0.0 * 1e10) - (0.1 + 0.0)) - 3.14159)) - float64(v2)) - 0.0)
        switch -2147483648 {
        case 2147483647:
            v34 := GenS2{F0: v1, F1: (v11 + v11), F2: v9}
            _ = v34
        case 9223372036854775807:
            switch "\n\t" {
            case "a":
                v35 := GenS3{F0: (^(v11)), F1: "\n\t"}
                _ = v35
            case "foo bar":
                var v36 float32 = 0.1
                _ = v36
            }
        case 128:
            v33 += v33
        }
        _ = fi32
        _ = v33
    }
    var v37 string = "foo bar"
    var v38 float64 = float64(0.1)
    if false {
        var v39 string = v37
        switch 0 {
        case 128:
            v12 = append(v12, int(v13))
        case -514:
            var v40 uint64 = v11
            _ = v40
        }
        v37 = "\n\t"
        _ = v39
    } else {
        v37 = ((v37 + "foo bar") + v37)
        for fi41 := 0; fi41 < 9; fi41++ {
            v3[v37] = true
            var v42 uint16 = 0
            var v43 uint = 0
            _ = fi41
            _ = v42
            _ = v43
        }
        switch -407 {
        case 9223372036854775807:
            if (false || (127 <= v13)) {
                var v44 string = ("\n\t" + v37)
                v45 := GenS1{F0: 1, F1: 886, F2: (v1 >> ((genFn0()) & 63))}
                v13 += (v13 >> ((genFn0()) & 63))
                _ = v44
                _ = v45
            } else {
                var v46 string = v37
                _ = v46
            }
        case 0:
            v47 := GenS4{F0: uint64(v13)}
            _ = v47
        case -478:
            var v48 uint8 = 255
            _ = v48
        default:
            for fi49 := 0; fi49 < 20; fi49++ {
                var v50 int64 = v9
                _ = fi49
                _ = v50
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v9
}
