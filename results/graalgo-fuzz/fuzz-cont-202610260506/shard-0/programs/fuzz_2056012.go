package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 int; F2 string }
func genFn0(p0 bool, p1 uint64, p2 uint16) uint32 {
    v4 := GenS1{F0: p0}
    v5 := GenS2{F0: 261, F1: -743, F2: 128}
    v6 := GenS1{F0: (0 > 395)}
    v7 := GenS2{F0: 2147483647, F1: 32767, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int16) string {
    v8 := []int{127, -340, -1}
    p0 = p0
    if true {
        for fi9 := 0; fi9 < 18; fi9++ {
            switch "\n\t" {
            case "hello":
                if (false && false) {
                    p0 += p0
                    fmt.Println(p0)
                } else {
                    p0 -= (p0 | 1)
                }
            case "foo bar":
                var v10 float64 = -2.25
                _ = v10
            }
            var v11 uint16 = 937
            v12 := GenS3{F0: (fi9 & fi9)}
            _ = fi9
            _ = v11
            _ = v12
        }
        p0 = p0
    }
    _ = p0
    _ = v8
    return ((("foo bar" + "a") + "foo bar") + "")
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{}
    var v3 float32 = 1e10
    var v13 uint32 = 1
    v3 += v3
    var v14 uint8 = 0
    switch ("foo bar" + (("a" + "hello") + ("" + ("\n\t" + "\n\t")))) {
    case "":
        var v15 string = (("a" + "a") + "hello")
        _ = v15
    default:
        v16 := v1["a"]
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        if false {
            v13 = 952
            v18 := GenS2{F0: (v13 / (((^(v13))) | 1)), F1: -32768, F2: (v14 ^ (v14 / (((^(v14))) | 1)))}
            if (!true) {
                for fi19 := 0; fi19 < 2; fi19++ {
                    v13 = (^(v13))
                    _ = fi19
                }
                for fi20 := 0; fi20 < 13; fi20++ {
                    v14 = (^(v14))
                    fmt.Println(fi20, v3)
                    var v21 string = "hello"
                    _ = fi20
                    _ = v21
                }
                var v22 uint64 = 42
                _ = v22
            }
            _ = v18
        }
        if ((fi17 >= (fi17 / ((fi17) | 1))) || (65535 >= 4294967295)) {
            v13 = v13
            var v23 int8 = -39
            _ = v23
        }
        var v24 int16 = -1
        _ = fi17
        _ = v24
    }
    v25 := GenS2{F0: 2147483648, F1: -106, F2: v14}
    switch int(v14) {
    case 128:
        var v26 int16 = 42
        _ = v26
    case 1:
        var v27 int64 = 9223372036854775807
        _ = v27
    }
    v3 = 3.14159
    var v28 uint16 = 65535
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 uint16 = (v28 % (((^(v28))) | 1))
        if false {
            if (true || (int64(v13) > -505)) {
                if (-406 == 843) {
                    fmt.Println(v3)
                }
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    v31 := GenS1{F0: false}
                    _ = v31
                case "a":
                    v13 += (^(v13))
                default:
                    for fi32 := 0; fi32 < 15; fi32++ {
                        var v33 uint64 = 2
                        switch 2147483647 {
                        case 32767:
                            if len(v2) > 0 {
                                v2[0] = 128
                            }
                        case 4294967295:
                            v34 := v1[""]
                            _ = v34
                        }
                        var v35 uint8 = (v14 >> ((42) & 63))
                        v30 = 32767
                        v36 := v1["hello"]
                        v37 := GenS3{F0: 32767}
                        _ = fi32
                        _ = v33
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                }
                v28 += (^(v28))
            } else {
                v1[""] = "hello"
            }
            if ((v3 + v3) <= float32(0.0)) {
                v30 = v28
                v38 := GenS2{F0: v13, F1: 1, F2: v14}
                switch fi29 {
                case 147:
                    v3 += (v3 / (float32(1e10) / v3))
                case -9223372036854775808:
                    switch "\n\t" {
                    case "\n\t":
                        v14 = v14
                    default:
                        v14 = (v14 << ((42) & 63))
                    }
                }
                _ = v38
            }
        }
        _ = fi29
        _ = v30
    }
    v3 += (1.5 / v3)
    var v39 uint8 = (v14 & v14)
    if true {
        _ = len(v1)
        v40 := GenS2{F0: (v13 ^ 2147483648), F1: -991, F2: v39}
        _ = v40
    }
    for fi41 := 0; fi41 < 2; fi41++ {
        fmt.Println(v13, fi41, v14)
        var v42 float32 = ((v3 / v3) - v3)
        v43 := GenS4{F0: 18446744073709551615, F1: fi41, F2: "a"}
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println(v13, v39, v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v39
}
