package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int }
type GenS5 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 uint; F1 uint; F2 int64 }
type GenS10 struct { F0 int8 }

func main() {
    v1 := []string{((("foo bar" + "hello") + "a") + "foo bar")}
    var v2 int = 1
    var v3 string = ("" + "a")
    var v4 uint8 = 42
    v5 := GenS1{F0: 389, F1: -211, F2: 42}
    switch ((v3 + v3) + v3) {
    case "a":
        switch (v2 | -722) {
        case -1:
            var v6 uint16 = 0
            _ = v6
        case 2147483647:
            v2 = v2
        case 42:
            switch (v2 - (v2 + (-(v2)))) {
            case 128:
                v4 = 128
            case 9223372036854775807:
                for fi7 := 0; fi7 < 7; fi7++ {
                    if (-128 >= 127) {
                        v8 := GenS1{F0: 220, F1: 0, F2: -1}
                        switch (fi7 * -2147483648) {
                        case -1:
                            v1 = append(v1, ((((v3 + v3) + ("\n\t" + (v3 + v3))) + (v3 + v3)) + (v3 + ("\n\t" + ("hello" + v3)))))
                        case 682:
                            var v9 float32 = (((-2.25 - 1e10) * 3.14159) + (float32(1.5) * 0.1))
                            _ = v9
                        case 1:
                            v10 := GenS2{F0: v4, F1: 1, F2: -825}
                            _ = v10
                        }
                        fmt.Println(v2, fi7)
                        _ = v8
                    } else {
                        v11 := GenS3{F0: 2147483647, F1: 1}
                        fmt.Println(v4, v2)
                        if false {
                            v12 := GenS2{F0: 2, F1: 1, F2: 38}
                            _ = v12
                        }
                        _ = v11
                    }
                    _ = fi7
                }
            case -972:
                var v13 string = v3
                _ = v13
            }
        default:
            var v14 uint32 = 42
            _ = v14
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch v3 {
    case "\n\t":
        var v15 int = 849
        _ = v15
    case "foo bar":
        var v16 uint32 = 4294967295
        _ = v16
    case "":
        v17 := GenS3{F0: -2147483648, F1: 42}
        _ = v17
    default:
        fmt.Println(v2, v3, v4)
    }
    var v18 uint32 = 42
    fmt.Println(v18)
    v19 := GenS4{F0: 720, F1: (v4 * (v4 & v4)), F2: v2}
    v18 = v18
    if false {
        if (true || false) {
            var v20 float32 = (0.0 - 0.1)
            _ = v20
        } else {
            v21 := GenS4{F0: 501, F1: 180, F2: (v2 + v2)}
            fmt.Println(v4)
            v4 -= 127
            _ = v21
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi23 := 0; fi23 < 11; fi23++ {
                v24 := GenS5{F0: 634, F1: 2, F2: 100}
                fmt.Println(fi22, fi23)
                v2 = fi22
                _ = fi23
                _ = v24
            }
            var v25 int64 = 9223372036854775807
            _ = fi22
            _ = v25
        }
        v26 := GenS6{F0: 0}
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 uint = 18446744073709551615
        var v29 bool = false
        for fi30 := 0; fi30 < 3; fi30++ {
            v31 := GenS7{F0: 1}
            _ = fi30
            _ = v31
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    if false {
        for fi32 := 0; fi32 < 5; fi32++ {
            if (!false) {
                v18 = (v18 + (^(v18)))
                var v33 int8 = 1
                v18 = (v18 << ((875) & 63))
                _ = v33
            }
            fmt.Println(v2, v18)
            switch v3 {
            case "hello":
                for fi34 := 0; fi34 < 6; fi34++ {
                    fmt.Println(v3, v4, v3)
                    _ = fi34
                }
            case "foo bar":
                switch (v3 + v3) {
                case "\n\t":
                    var v35 int = 201
                    _ = v35
                case "":
                    if true {
                        var v36 int32 = 42
                        fmt.Println(fi32)
                        var v37 int64 = 16
                        _ = v36
                        _ = v37
                    }
                default:
                    if ((1 > 78) && (-32768 <= 42)) {
                        v38 := GenS4{F0: 42, F1: v4, F2: v2}
                        v39 := GenS8{F0: 264, F1: -817}
                        _ = v38
                        _ = v39
                    }
                }
            }
            _ = fi32
        }
        v40 := GenS1{F0: 1, F1: 120, F2: 0}
        if (1 > -2147483648) {
            if (true && (((v4 / (((^(v4))) | 1)) >= v4) || ((v2 - -905) != (v2 << ((1) & 63))))) {
                for fi41 := 0; fi41 < 9; fi41++ {
                    if (18446744073709551615 < 2) {
                        v42 := GenS9{F0: 580, F1: 0, F2: -9223372036854775808}
                        v43 := GenS10{F0: 112}
                        v4 = 42
                        _ = v42
                        _ = v43
                    }
                    _ = fi41
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi44 := 0; fi44 < 12; fi44++ {
                    v18 += 415
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi44
                }
            }
            v45 := GenS10{F0: 0}
            fmt.Println(v3)
            _ = v45
        } else {
            fmt.Println(v4, v18)
            switch (v2 - v2) {
            case -2147483648:
                var v46 float64 = 1e10
                _ = v46
            case 4294967295:
                for fi47 := 0; fi47 < 8; fi47++ {
                    var v48 uint = 2
                    _ = fi47
                    _ = v48
                }
            case 32768:
                v2 = (v2 - (v2 % (((^(v2))) | 1)))
            default:
                v3 = "\n\t"
            }
        }
        _ = v40
    }
    var v49 uint64 = 1
    v2 = -110
    for fi50 := 0; fi50 < 8; fi50++ {
        fmt.Println(fi50, v3, v2)
        var v51 uint64 = v49
        _ = fi50
        _ = v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
