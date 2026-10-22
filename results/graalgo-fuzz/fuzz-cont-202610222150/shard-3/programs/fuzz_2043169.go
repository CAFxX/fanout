package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int64; F1 bool; F2 uint }
func genFn0() uint32 {
    var v5 bool = (!true)
    switch 42 {
    case 1:
        if (!(65535 != 0)) {
            v5 = (!((!v5) || v5))
            for fi6 := 0; fi6 < 6; fi6++ {
                if (1 == -1) {
                    v5 = ("a" < "hello")
                    v5 = (!v5)
                    for fi7 := 0; fi7 < 10; fi7++ {
                        var v8 int = (fi6 - fi7)
                        switch (fi7 ^ v8) {
                        case 9223372036854775807:
                            v9 := GenS1{F0: 65535, F1: 0}
                            var v10 int64 = 0
                            v10 = (v10 << ((673) & 63))
                            _ = v10
                            _ = v9
                        }
                        _ = fi7
                        _ = v8
                    }
                } else {
                    for fi11 := 0; fi11 < 13; fi11++ {
                        var v12 bool = (0 <= 32767)
                        _ = fi11
                        _ = v12
                    }
                }
                v13 := []int{fi6, fi6, -1}
                _ = fi6
                _ = v13
            }
        }
    case 193:
        if (((v5 && (!false)) && v5) && (685 <= 986)) {
            v14 := []int{-667, 525, 9223372036854775807, 4294967295}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                v16 := GenS2{F0: (fi15 & fi15), F1: 2}
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            fmt.Println(v5)
            if (false || v5) {
                var v17 uint32 = 128
                _ = v17
            } else {
                var v18 uint16 = 846
                _ = v18
            }
        }
    }
    _ = v5
    return 42
}

func main() {
    var v1 uint8 = 167
    var v2 uint64 = 42
    var v3 uint32 = 0
    v4 := []uint32{}
    if (v1 < v1) {
        if false {
            v4 = append(v4, (v3 % (((v3 ^ 779)) | 1)))
            v3 += (v3 % ((genFn0()) | 1))
            switch "a" {
            case "hello":
                fmt.Println(v2, v2)
            case "\n\t":
                if true {
                    if (487 >= 1) {
                        var v19 string = ("hello" + "")
                        v2 = (v2 & (v2 & v2))
                        for fi20 := 0; fi20 < 11; fi20++ {
                            v21 := GenS2{F0: (fi20 << ((42) & 63)), F1: 357}
                            v22 := GenS1{F0: (^(v3)), F1: v1}
                            v2 += v2
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        _ = v19
                    } else {
                        v23 := GenS3{F0: 932, F1: true, F2: 525}
                        _ = v23
                    }
                    v2 = uint64(v3)
                }
            default:
                for fi24 := 0; fi24 < 11; fi24++ {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 uint16 = 1
                        _ = fi25
                        _ = v26
                    }
                    _ = fi24
                }
            }
        } else {
            v1 -= 26
            if (3.14159 != 0.1) {
                v2 = 722
                var v27 string = "a"
                _ = v27
            }
            if (true || (!("a" != (("" + "\n\t") + ("hello" + "hello"))))) {
                if true {
                    if (false && false) {
                        _ = len(v4)
                    }
                    switch 187 {
                    case -2147483648:
                        var v28 uint64 = (v2 - 86)
                        _ = v28
                    case -46:
                        fmt.Println(v1, v2)
                    }
                }
            } else {
                v3 = v3
                v1 -= 0
            }
        }
    }
    var v29 string = ((("" + "foo bar") + "hello") + "foo bar")
    if false {
        if ("hello" == (v29 + v29)) {
            v2 -= 679
        }
    }
    _ = len(v4)
    switch v29 {
    case "\n\t":
        fmt.Println(v2, v29)
    }
    if true {
        var v30 uint16 = 2
        switch -9223372036854775808 {
        case -297:
            var v31 uint32 = v3
            _ = v31
        case 32767:
            _ = len(v4)
        }
        _ = v30
    } else {
        var v32 string = v29
        if false {
            var v33 uint16 = 32767
            v3 = genFn0()
            _ = v33
        }
        _ = v32
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 += v2
    fmt.Println(v3, v3)
    _ = len(v4)
    var v34 uint64 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v35 string = v29
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
