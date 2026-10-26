package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 int64; F1 int }
func genFn0() uint8 {
    if true {
        v4 := []int{32768, 2147483647, -490}
        switch 127 {
        case -991:
            var v5 int8 = 92
            _ = v5
        default:
            var v6 uint = 330
            _ = v6
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int16 = 1000
            for fi9 := 0; fi9 < 6; fi9++ {
                v8 = -34
                switch (("hello" + ("\n\t" + ("a" + ("hello" + "foo bar")))) + "") {
                case "":
                    if (9223372036854775807 == 53) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    var v10 int8 = -1
                    _ = v10
                }
                if len(v4) > 0 {
                    v4[0] = fi7
                }
                _ = fi9
            }
            if (!(!(!(!(0.1 > 3.14159))))) {
                var v11 uint16 = 2
                switch (fi7 * (fi7 % ((fi7) | 1))) {
                case 847:
                    fmt.Println(fi7, v11)
                }
                for fi12 := 0; fi12 < 16; fi12++ {
                    if false {
                        v8 = 264
                        fmt.Println(v11)
                        var v13 uint32 = 918
                        _ = v13
                    } else {
                        var v14 uint64 = 0
                        switch (fi12 << ((26) & 63)) {
                        case -887:
                            var v15 uint32 = 4294967295
                            _ = v15
                        default:
                            v16 := GenS1{F0: (v14 / (((v14 + v14)) | 1))}
                            _ = v16
                        }
                        _ = v14
                    }
                    _ = fi12
                }
                _ = v11
            } else {
                if len(v4) > 0 {
                    v4[0] = (-(fi7))
                }
                v8 -= v8
            }
            _ = fi7
            _ = v8
        }
        _ = v4
    } else {
        v17 := []int{42}
        _ = v17
    }
    var v18 int32 = 42
    fmt.Println(v18)
    _ = v18
    return 42
}
func genFn1(p0 uint8) int64 {
    v19 := GenS2{F0: -9223372036854775808, F1: -2147483648, F2: ((("\n\t" + "\n\t") + "") + "foo bar")}
    if true {
        for fi20 := 0; fi20 < 9; fi20++ {
            for fi21 := 0; fi21 < 7; fi21++ {
                v22 := []int{32767, (fi20 + fi20)}
                v23 := GenS1{F0: 42}
                _ = fi21
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 16; fi24++ {
                v25 := GenS2{F0: fi20, F1: -1, F2: "a"}
                _ = fi24
                _ = v25
            }
            _ = fi20
        }
    } else {
        p0 += (p0 - genFn0())
        p0 = (p0 + p0)
    }
    p0 = 0
    v26 := GenS1{F0: 2}
    v27 := GenS1{F0: 1}
    _ = p0
    _ = v19
    _ = v26
    _ = v27
    return 42
}

func main() {
    var v1 int64 = 0
    var v2 uint = 255
    v3 := []uint32{574}
    for fi28 := 0; fi28 < 11; fi28++ {
        v3 = append(v3, 0)
        if true {
            if (42 >= 2) {
                v29 := GenS2{F0: (fi28 + (fi28 * fi28)), F1: -89, F2: ""}
                for fi30 := 0; fi30 < 2; fi30++ {
                    v31 := GenS3{F0: v1, F1: fi30}
                    switch fi30 {
                    case 808:
                        var v32 int16 = 42
                        _ = v32
                    }
                    _ = fi30
                    _ = v31
                }
                _ = v29
            } else {
                v1 += (v1 % ((v1) | 1))
                fmt.Println(fi28, v2, v1)
            }
        } else {
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 string = "\n\t"
                _ = fi33
                _ = v34
            }
            v2 += (v2 * v2)
        }
        _ = fi28
    }
    var v35 uint16 = 2
    v3 = append(v3, 184)
    _ = len(v3)
    v36 := GenS1{F0: 43}
    var v37 int = 4294967295
    var v38 int16 = 0
    for fi39 := 0; fi39 < 1; fi39++ {
        v2 = 4294967295
        _ = fi39
    }
    if (true && (((v1 / (((v1 * genFn1(72))) | 1)) == (v1 % (((-(v1))) | 1))) || (!(42 > v2)))) {
        var v40 string = "hello"
        for fi41 := 0; fi41 < 2; fi41++ {
            var v42 float64 = (3.14159 - 0.0)
            if false {
                fmt.Println(v2)
                var v43 float64 = v42
                _ = v43
            }
            v40 = v40
            _ = fi41
            _ = v42
        }
        _ = v40
    } else {
        var v44 uint64 = 588
        switch v37 {
        case 1:
            for fi45 := 0; fi45 < 4; fi45++ {
                var v46 int64 = v1
                v47 := GenS3{F0: v1, F1: (v37 << ((42) & 63))}
                _ = fi45
                _ = v46
                _ = v47
            }
        }
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
