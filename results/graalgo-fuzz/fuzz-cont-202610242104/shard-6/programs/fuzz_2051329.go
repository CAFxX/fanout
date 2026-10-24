package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS2 struct { F0 int; F1 int; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    switch -9223372036854775808 {
    case 128:
        var v8 int16 = -37
        _ = v8
    default:
        if false {
            var v9 uint64 = 2
            for fi10 := 0; fi10 < 15; fi10++ {
                if (1.5 < float32(-2.25)) {
                    var v11 uint8 = 1
                    _ = v11
                } else {
                    v12 := GenS1{F0: 1, F1: 239, F2: v9}
                    var v13 int64 = -402
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
            _ = v9
        } else {
            var v14 string = ""
            fmt.Println(v14)
            _ = v14
        }
    }
    var v15 float32 = float32(1.5)
    v16 := GenS1{F0: 18446744073709551615, F1: 1000, F2: 1}
    var v17 float32 = ((0.1 * 0.1) / (v15 + v15))
    _ = v15
    _ = v16
    _ = v17
    return 2
}

func main() {
    var v1 uint = 65535
    v2 := []string{"foo bar"}
    var v3 uint32 = 65535
    var v4 float64 = 1.5
    var v5 int8 = -114
    var v6 int32 = 42
    for fi18 := 0; fi18 < 4; fi18++ {
        v5 = v5
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 bool = ("a" > ("hello" + "hello"))
            v21 := GenS1{F0: 4294967295, F1: -1, F2: 0}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    if (!((((("a" != "foo bar") || false) || (42 != -56)) || true) || false)) {
        v22 := GenS1{F0: 18446744073709551615, F1: -4, F2: 4294967295}
        _ = v22
    } else {
        fmt.Println(v6, v5, v3)
        fmt.Println(v4)
    }
    var v23 uint16 = 1
    var v24 bool = true
    v5 = -43
    v4 = (v4 + (((1e10 / v4) / v4) - ((((v4 + 0.1) * (-2.25 + -2.25)) / (v4 + float64(float32(1e10)))) + v4)))
    if v24 {
        for fi25 := 0; fi25 < 18; fi25++ {
            fmt.Println(v23)
            _ = fi25
        }
        switch -9223372036854775808 {
        case 1:
            if (-2.25 > (1.5 * 0.1)) {
                fmt.Println(v6)
                var v26 bool = v24
                var v27 int32 = v6
                _ = v26
                _ = v27
            }
        case 4294967295:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v28 float64 = v4
            _ = v28
        }
        if v24 {
            for fi29 := 0; fi29 < 9; fi29++ {
                switch ("foo bar" + ("a" + "\n\t")) {
                case "\n\t":
                    v30 := GenS1{F0: 18446744073709551615, F1: -32768, F2: 0}
                    _ = v30
                case "":
                    v5 += v5
                }
                if true {
                    v31 := GenS2{F0: (fi29 & fi29), F1: fi29, F2: v3}
                    for fi32 := 0; fi32 < 4; fi32++ {
                        switch fi29 {
                        case 128:
                            var v33 uint32 = (v3 / ((uint32(v23)) | 1))
                            _ = v33
                        case -133:
                            var v34 float64 = (1e10 / v4)
                            _ = v34
                        case 285:
                            v35 := GenS1{F0: 42, F1: -1, F2: 4294967295}
                            _ = v35
                        }
                        v24 = v24
                        var v36 bool = (!(true || v24))
                        _ = fi32
                        _ = v36
                    }
                    _ = v31
                }
                _ = fi29
            }
        }
    }
    if ((v6 & (v6 / (((-(v6))) | 1))) > (v6 % (((-(v6))) | 1))) {
        v23 += v23
        var v37 uint = uint(v3)
        var v38 uint8 = 48
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        v4 = v4
        _ = fi39
    }
    if (-372 > 0) {
        var v40 int32 = 1
        v41 := GenS3{F0: -128, F1: 214}
        var v42 int = -1
        _ = v40
        _ = v41
        _ = v42
    } else {
        v6 += (v6 >> ((uint(v23)) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
