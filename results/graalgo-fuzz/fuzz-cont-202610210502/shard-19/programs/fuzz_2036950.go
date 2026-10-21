package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int16; F1 bool; F2 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool; F1 int; F2 uint32 }
type GenS9 struct { F0 uint8; F1 uint }
func genFn0(p0 int32, p1 uint64) int32 {
    switch -846 {
    case -1:
        fmt.Println(p1)
    case 881:
        p0 = -1
    default:
        fmt.Println(p1, p0)
    }
    if false {
        v5 := []int{}
        var v6 string = ("a" + ("\n\t" + "\n\t"))
        _ = len(v5)
        _ = v5
        _ = v6
    }
    v7 := []int{-602, 128, 2147483647, 32767}
    if len(v7) > 0 {
        v7[0] = -1
    }
    if true {
        var v8 float32 = 3.14159
        _ = len(v7)
        if true {
            fmt.Println(p1, v8, v8)
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 uint8 = 188
                p0 -= 884
                _ = fi9
                _ = v10
            }
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return p0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 0
    var v3 string = "foo bar"
    v4 := []int{378, 32768, -2147483648}
    var v11 string = v3
    v3 = v11
    switch v11 {
    case "foo bar":
        v3 = "a"
    }
    switch (((v11 + (v11 + "a")) + ((("hello" + ((v11 + v11) + "")) + v3) + v11)) + v3) {
    case "foo bar":
        if ((v2 & -9223372036854775808) > (v2 * v2)) {
            if true {
                v12 := GenS1{F0: 42}
                _ = v12
            }
            v1[0] = 65535
        }
    default:
        v2 -= -919
    }
    fmt.Println(v2)
    if len(v4) > 0 {
        v4[0] = 32768
    }
    var v13 int64 = 396
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(v3)
        v15 := GenS2{F0: 923, F1: 19}
        if true {
            if (!(!((fi14 | (fi14 * 127)) != (fi14 / ((4294967295) | 1))))) {
                var v16 int = int(v2)
                var v17 int32 = 2147483647
                var v18 int16 = -151
                _ = v16
                _ = v17
                _ = v18
            }
            _ = len(v1)
        }
        _ = fi14
        _ = v15
    }
    if ((!true) && ((300 != (-(v13))) && true)) {
        for fi19 := 0; fi19 < 13; fi19++ {
            v13 += (^(v13))
            _ = fi19
        }
        for fi20 := 0; fi20 < 9; fi20++ {
            _ = len(v4)
            _ = len(v4)
            v2 += v2
            _ = fi20
        }
    }
    switch (v3 + ("foo bar" + v11)) {
    case "hello":
        v3 = "a"
    case "a":
        v21 := GenS3{F0: 18446744073709551615, F1: 34}
        _ = v21
    default:
        fmt.Println(v11)
    }
    if (true || true) {
        switch v11 {
        case "a":
            switch "" {
            case "\n\t":
                if true {
                    if false {
                        v13 -= -137
                    }
                } else {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        fmt.Println(fi22)
                        _ = fi22
                    }
                    v2 = (v2 / ((v13) | 1))
                }
            }
        case "\n\t":
            var v23 uint16 = 65535
            _ = v23
        default:
            switch 42 {
            case 32768:
                var v24 uint16 = 2
                _ = v24
            case -9223372036854775808:
                switch 32767 {
                case 768:
                    if len(v4) > 0 {
                        v4[0] = 32767
                    }
                default:
                    fmt.Println(v2, v2)
                }
            default:
                if (-570 == 1) {
                    v25 := GenS1{F0: -1}
                    if false {
                        fmt.Println(v3, v11)
                    } else {
                        var v26 bool = true
                        var v27 float32 = ((1e10 / 3.14159) - -2.25)
                        _ = len(v4)
                        _ = v26
                        _ = v27
                    }
                    _ = v25
                } else {
                    var v28 int8 = 1
                    v29 := GenS4{F0: v2, F1: -1}
                    _ = v28
                    _ = v29
                }
            }
        }
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        switch fi30 {
        case 32767:
            v3 = v11
        case -182:
            switch (v11 + ("a" + (v11 + ((v3 + "") + "\n\t")))) {
            case "foo bar":
                v31 := GenS4{F0: (v2 << ((1) & 63)), F1: 1}
                _ = v31
            case "hello":
                v2 -= (v13 >> ((0) & 63))
            case "\n\t":
                v13 -= v2
            default:
                if (565 > 255) {
                    var v32 uint64 = 42
                    v33 := GenS5{F0: 42, F1: 127}
                    fmt.Println(v32)
                    _ = v32
                    _ = v33
                }
            }
        case 0:
            v34 := GenS6{F0: 386, F1: true, F2: 9223372036854775807}
            _ = v34
        default:
            var v35 uint16 = 1
            _ = v35
        }
        fmt.Println(v11)
        _ = fi30
    }
    v36 := GenS5{F0: 169, F1: 57}
    for fi37 := 0; fi37 < 17; fi37++ {
        v11 = ("\n\t" + "foo bar")
        if ((fi37 >> ((42) & 63)) <= fi37) {
            v38 := GenS3{F0: 4294967295, F1: -750}
            for fi39 := 0; fi39 < 10; fi39++ {
                if false {
                    v3 = v11
                    switch ((("" + v11) + v3) + v3) {
                    case "\n\t":
                        _ = len(v4)
                    default:
                        v40 := GenS7{F0: 336}
                        _ = v40
                    }
                    fmt.Println(v11)
                }
                v41 := GenS8{F0: false, F1: (fi39 - fi39), F2: 724}
                v4 = append(v4, (fi37 & fi37))
                _ = fi39
                _ = v41
            }
            var v42 int8 = -128
            _ = v38
            _ = v42
        } else {
            v13 = (v2 - v2)
            fmt.Println(v2, v3, v3)
            v43 := GenS4{F0: (v13 % (((v2 & v2)) | 1)), F1: 1}
            _ = v43
        }
        _ = fi37
    }
    var v44 int64 = v2
    for fi45 := 0; fi45 < 13; fi45++ {
        switch v11 {
        case "\n\t":
            switch ("" + (v11 + v3)) {
            case "foo bar":
                var v46 bool = true
                _ = v46
            case "a":
                var v47 float32 = 0.1
                _ = v47
            default:
                v44 = v13
            }
        }
        v48 := GenS9{F0: 0, F1: 255}
        if len(v4) > 0 {
            v4[0] = fi45
        }
        _ = fi45
        _ = v48
    }
    var v49 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v44
    _ = v49
}
