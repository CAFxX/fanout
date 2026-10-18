package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0(p0 uint64, p1 int) bool {
    p1 += p1
    if true {
        fmt.Println(p1)
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 uint64 = 784
    var v2 int8 = 36
    var v3 int32 = 999
    v4 := []string{"a", "hello", ("" + ("" + (("" + ("\n\t" + ("foo bar" + (((("foo bar" + "a") + ("foo bar" + ("hello" + "a"))) + "foo bar") + "\n\t")))) + ("" + "foo bar"))))}
    var v5 bool = (!((!(false && true)) || true))
    v6 := []int{42, 128}
    var v7 uint32 = 2147483648
    v8 := GenS1{F0: (v3 ^ -2147483648), F1: ((1e10 * 1e10) <= float64(1.5)), F2: v7}
    v9 := GenS1{F0: -420, F1: v5, F2: uint32(v1)}
    var v10 uint32 = v7
    if true {
        var v11 uint16 = 969
        v12 := GenS1{F0: v3, F1: v5, F2: (^(v10))}
        _ = v11
        _ = v12
    }
    var v13 int32 = (v3 >> ((0) & 63))
    var v14 int32 = int32(v2)
    if (866 > 255) {
        switch 0 {
        case -2147483648:
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 uint8 = 210
                if (true || v5) {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        v18 := GenS1{F0: -323, F1: ((v2 / (((v2 & 121)) | 1)) > -52), F2: v7}
                        var v19 uint8 = (^(v16))
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                fmt.Println(v16)
                _ = fi15
                _ = v16
            }
        case 297:
            v1 -= (v1 & (^(v1)))
        default:
            var v20 int64 = -9223372036854775808
            _ = v20
        }
        fmt.Println(v13, v13)
    } else {
        switch 861 {
        case 127:
            v21 := GenS1{F0: (v3 << ((1) & 63)), F1: (!((((((0.0 * ((1.5 + ((1e10 - 1.5) / (float32(1e10) - (float32(0.0) / (1e10 - (3.14159 * 3.14159)))))) / 1e10)) / float32(1.5)) / (1e10 * ((0.0 + (0.0 - 0.1)) - -2.25))) - (1e10 - (((((0.0 - -2.25) + 1.5) + (0.0 / (0.1 * (0.0 + ((0.1 - (float32((0.1 / 1.0)) * (((1.5 * 0.1) + 0.1) * 0.0))) / (1e10 * 3.14159)))))) - ((((0.1 / -2.25) * 1.5) + (3.14159 - -2.25)) + 0.1)) * (0.1 * ((1e10 - 3.14159) + 3.14159))))) + (0.1 * 1.5)) != 0.0)), F2: 59}
            _ = v21
        case 2147483647:
            fmt.Println(v10, v3, v2)
        case -2147483648:
            var v22 uint64 = 0
            _ = v22
        default:
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        }
        var v23 int64 = -1
        var v24 uint32 = v7
        _ = v23
        _ = v24
    }
    v6 = append(v6, -163)
    if v5 {
        v4 = append(v4, "foo bar")
        for fi25 := 0; fi25 < 1; fi25++ {
            switch ("foo bar" + "\n\t") {
            case "":
                var v26 uint16 = 65535
                _ = v26
            case "a":
                v2 += v2
            }
            if false {
                for fi27 := 0; fi27 < 9; fi27++ {
                    var v28 float64 = -2.25
                    v13 += (-(v3))
                    v3 = (v3 % ((0) | 1))
                    var v29 int32 = (v13 % ((42) | 1))
                    v14 = (v3 % ((v3) | 1))
                    v5 = v5
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            _ = len(v6)
            _ = fi25
        }
    }
    if true {
        for fi30 := 0; fi30 < 16; fi30++ {
            switch -1 {
            case 32767:
                var v31 float64 = 0.1
                _ = v31
            case -2147483648:
                switch fi30 {
                case -313:
                    for fi32 := 0; fi32 < 14; fi32++ {
                        v33 := GenS1{F0: v13, F1: true, F2: v7}
                        v34 := GenS1{F0: (v3 >> ((236) & 63)), F1: true, F2: 2147483648}
                        switch (fi30 / ((fi32) | 1)) {
                        case -621:
                            v35 := GenS2{F0: 181}
                            _ = v35
                        case 32767:
                            v1 += v1
                        default:
                            var v36 int32 = 1
                            _ = v36
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                default:
                    v14 -= int32(fi30)
                }
            }
            _ = fi30
        }
        if v5 {
            v13 = 42
        } else {
            v5 = (3.14159 > 1.5)
            v37 := GenS2{F0: 510}
            _ = v37
        }
    }
    var v38 uint64 = 18446744073709551615
    var v39 uint16 = 32767
    switch ("a" + ("\n\t" + "")) {
    case "foo bar":
        var v40 uint64 = (v1 / (((v38 << ((501) & 63))) | 1))
        _ = v40
    default:
        for fi41 := 0; fi41 < 9; fi41++ {
            var v42 int64 = 1
            switch int(v1) {
            case 127:
                v6 = append(v6, (-(fi41)))
            case 32767:
                for fi43 := 0; fi43 < 11; fi43++ {
                    if false {
                        fmt.Println(v5, fi43)
                    }
                    v44 := GenS3{F0: 32767}
                    _ = fi43
                    _ = v44
                }
            case 404:
                for fi45 := 0; fi45 < 6; fi45++ {
                    v46 := GenS4{F0: 210, F1: (-(v3))}
                    _ = len(v6)
                    _ = fi45
                    _ = v46
                }
            }
            for fi47 := 0; fi47 < 10; fi47++ {
                var v48 float32 = 1e10
                if (203 < 32767) {
                    fmt.Println(v39)
                    v4 = append(v4, (("hello" + "foo bar") + ("hello" + "\n\t")))
                    v49 := GenS4{F0: uint(v10), F1: 0}
                    _ = v49
                } else {
                    var v50 int16 = -1
                    _ = v50
                }
                _ = fi47
                _ = v48
            }
            _ = fi41
            _ = v42
        }
    }
    v5 = v5
    v10 = v7
    v3 = (v3 ^ v13)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
