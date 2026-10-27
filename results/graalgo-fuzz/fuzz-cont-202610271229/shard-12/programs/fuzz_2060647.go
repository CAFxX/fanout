package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int; F1 int16; F2 uint64 }
type GenS3 struct { F0 string; F1 int32; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string }
func genFn0() uint32 {
    var v4 string = ((("a" + "foo bar") + "foo bar") + ((("hello" + ((("" + "hello") + ("" + "a")) + (("foo bar" + ("hello" + "a")) + "\n\t"))) + (("hello" + "hello") + "foo bar")) + "foo bar"))
    v4 = v4
    v5 := GenS1{F0: -128, F1: 0, F2: 89}
    _ = v4
    _ = v5
    return 560
}
func genFn1(p0 uint32) int8 {
    p0 = p0
    if (127 < 0) {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS2{F0: 32767, F1: 0, F2: 42}
            if true {
                v8 := GenS2{F0: 32767, F1: 1000, F2: 4294967295}
                var v9 bool = false
                v10 := GenS3{F0: ("foo bar" + ("foo bar" + "foo bar")), F1: -1, F2: (fi6 >> ((255) & 63))}
                _ = v10
                _ = v8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    v11 := GenS3{F0: "a", F1: 2147483647, F2: 2147483647}
    v12 := GenS2{F0: -2147483648, F1: -1, F2: 0}
    _ = p0
    _ = v11
    _ = v12
    return -128
}

func main() {
    var v1 uint = 65535
    var v2 string = "a"
    var v3 float32 = 1.5
    fmt.Println(v1)
    var v13 int64 = 0
    var v14 bool = true
    var v15 uint16 = 2
    v1 += (v1 * (^(v1)))
    v14 = (v14 || true)
    if v14 {
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS4{F0: 0}
            if v14 {
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := GenS1{F0: 90, F1: 18446744073709551615, F2: 146}
                    _ = fi18
                    _ = v19
                }
            }
            v15 = (v15 | 42)
            _ = fi16
            _ = v17
        }
        switch v2 {
        case "foo bar":
            fmt.Println(v13, v1, v1)
        case "a":
            var v20 string = v2
            _ = v20
        case "\n\t":
            v1 -= v1
        }
        if v14 {
            var v21 uint64 = 18446744073709551615
            _ = v21
        } else {
            v22 := []int{9223372036854775807, 623}
            v15 -= v15
            _ = v22
        }
    }
    v23 := GenS2{F0: 2147483647, F1: 1, F2: 2}
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := []int{425, (-(fi24))}
        if len(v25) > 0 {
            v25[0] = (fi24 | 9223372036854775807)
        }
        var v26 int8 = 1
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := []int{}
    v28 := GenS1{F0: int8(v15), F1: 2, F2: 42}
    switch 128 {
    case 32767:
        for fi29 := 0; fi29 < 11; fi29++ {
            for fi30 := 0; fi30 < 17; fi30++ {
                if v14 {
                    if len(v27) > 0 {
                        v27[0] = fi29
                    }
                    v31 := GenS1{F0: 42, F1: 2, F2: 128}
                    v32 := GenS2{F0: (fi30 ^ 128), F1: -32768, F2: 2}
                    _ = v31
                    _ = v32
                } else {
                    if (fi29 != -310) {
                        v2 = "\n\t"
                    }
                    _ = len(v27)
                    v2 = ("\n\t" + v2)
                }
                _ = fi30
            }
            _ = len(v27)
            _ = fi29
        }
    }
    fmt.Println(v1)
    for fi33 := 0; fi33 < 16; fi33++ {
        v34 := GenS4{F0: 0}
        switch (v2 + v2) {
        case "hello":
            if true {
                if (v3 == 3.14159) {
                    switch (fi33 >> ((468) & 63)) {
                    case 2147483647:
                        var v35 bool = (float32(-2.25) < 0.1)
                        _ = v35
                    case 32767:
                        v27 = append(v27, (fi33 ^ (fi33 ^ fi33)))
                    }
                    if (("foo bar" <= v2) || true) {
                        var v36 int = 42
                        _ = v36
                    } else {
                        _ = len(v27)
                    }
                } else {
                    var v37 int8 = 123
                    var v38 int = 32767
                    _ = v37
                    _ = v38
                }
            } else {
                if (0 < 82) {
                    var v39 uint16 = 2
                    var v40 float32 = v3
                    _ = v39
                    _ = v40
                } else {
                    if len(v27) > 0 {
                        v27[0] = fi33
                    }
                }
            }
        case "a":
            for fi41 := 0; fi41 < 6; fi41++ {
                var v42 float32 = 0.0
                fmt.Println(v14, v3)
                _ = fi41
                _ = v42
            }
        }
        v14 = true
        _ = fi33
        _ = v34
    }
    var v43 int8 = -81
    if (v14 && (v14 || v14)) {
        for fi44 := 0; fi44 < 10; fi44++ {
            if len(v27) > 0 {
                v27[0] = (fi44 & fi44)
            }
            _ = fi44
        }
    } else {
        _ = len(v27)
        for fi45 := 0; fi45 < 7; fi45++ {
            switch 2147483647 {
            case 32768:
                v46 := GenS1{F0: 100, F1: 25, F2: 0}
                _ = v46
            default:
                for fi47 := 0; fi47 < 1; fi47++ {
                    var v48 uint32 = 42
                    var v49 int16 = 0
                    if true {
                        v48 = (v48 % (((^(v48))) | 1))
                        var v50 int32 = -2147483648
                        _ = v50
                    }
                    if (!true) {
                        v3 = ((v3 - v3) + v3)
                        if v14 {
                            if false {
                                fmt.Println(v13, v3, v3)
                            } else {
                                var v51 int8 = v43
                                fmt.Println(v13, v15, v15)
                                _ = v51
                            }
                            v52 := GenS3{F0: v2, F1: 2147483647, F2: (^(fi47))}
                            _ = v52
                        }
                    } else {
                        v53 := GenS5{F0: v2}
                        _ = v53
                    }
                    var v54 string = (v2 + "a")
                    v3 += 0.0
                    _ = fi47
                    _ = v48
                    _ = v49
                    _ = v54
                }
            }
            _ = fi45
        }
    }
    var v55 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v43
    _ = v55
}
