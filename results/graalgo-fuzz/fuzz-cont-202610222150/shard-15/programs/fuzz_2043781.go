package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int; F1 int32 }
func genFn0() uint8 {
    var v5 string = ""
    fmt.Println(v5)
    for fi6 := 0; fi6 < 19; fi6++ {
        v5 = "a"
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    var v8 int = 127
    _ = v5
    _ = v8
    return 109
}
func genFn1(p0 string, p1 int64) int32 {
    var v9 float64 = 0.1
    var v10 uint8 = 128
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2147483647
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]string)
    v3 := []int{128, 468, 128, 31}
    var v4 int32 = 1
    for fi11 := 0; fi11 < 2; fi11++ {
        if true {
            switch (fi11 - fi11) {
            case -751:
                fmt.Println(v1, fi11)
            }
            if true {
                var v12 int64 = 1
                _ = v12
            } else {
                fmt.Println(fi11, fi11, fi11)
            }
            var v13 string = "foo bar"
            _ = v13
        } else {
            v14 := GenS1{F0: 0, F1: (v4 >> ((279) & 63)), F2: 0}
            _ = v14
        }
        _ = fi11
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint16 = 32767
        v16 += v16
        if true {
            var v17 uint64 = 0
            var v18 int8 = (^(v1))
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS1{F0: 536, F1: (v4 ^ 0), F2: 9223372036854775807}
                v16 -= (v16 ^ 608)
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v4 -= -2147483648
        _ = fi21
    }
    var v22 uint8 = 34
    fmt.Println(v4, v22)
    if (false || true) {
        if true {
            v23 := GenS2{F0: (v22 / (((v22 - v22)) | 1)), F1: 127}
            if false {
                v1 = v1
                for fi24 := 0; fi24 < 11; fi24++ {
                    for fi25 := 0; fi25 < 14; fi25++ {
                        fmt.Println(fi25)
                        var v26 uint8 = 2
                        var v27 int8 = v1
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    _ = fi24
                }
                if (!true) {
                    var v28 int8 = v1
                    v29 := GenS1{F0: 42, F1: (-(v4)), F2: 9223372036854775807}
                    _ = v28
                    _ = v29
                }
            } else {
                fmt.Println(v4, v1)
            }
            _ = v23
        } else {
            v22 = 1
        }
        var v30 int64 = 1
        var v31 uint = 598
        _ = v30
        _ = v31
    } else {
        v1 -= 127
        var v32 uint16 = 1
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 int16 = -238
            v35 := GenS2{F0: v22, F1: 127}
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = v32
    }
    switch "\n\t" {
    case "hello":
        switch -992 {
        case 42:
            v22 = v22
        case -331:
            var v36 uint = 1
            _ = v36
        case 2147483647:
            v37 := GenS3{F0: -1, F1: 123, F2: v1}
            _ = v37
        default:
            var v38 int32 = (v4 << ((18446744073709551615) & 63))
            _ = v38
        }
    case "foo bar":
        v39 := GenS4{F0: 18446744073709551615, F1: v22}
        _ = v39
    }
    v4 += v4
    v40 := v2[982]
    if (float64(0.1) > float64(3.14159)) {
        fmt.Println(v40)
    } else {
        if len(v3) > 0 {
            v3[0] = -203
        }
    }
    v3 = append(v3, 127)
    for fi41 := 0; fi41 < 20; fi41++ {
        switch v40 {
        case "a":
            for fi42 := 0; fi42 < 12; fi42++ {
                var v43 bool = true
                var v44 string = (("a" + "a") + ((((v40 + (v40 + (v40 + v40))) + "") + ("foo bar" + v40)) + "a"))
                _ = fi42
                _ = v43
                _ = v44
            }
        case "foo bar":
            fmt.Println(v40, v40)
        case "\n\t":
            var v45 int8 = v1
            _ = v45
        default:
            v46 := GenS5{F0: fi41, F1: genFn1("foo bar", 42)}
            _ = v46
        }
        for fi47 := 0; fi47 < 19; fi47++ {
            var v48 int8 = (v1 - 100)
            var v49 uint = 42
            _ = fi47
            _ = v48
            _ = v49
        }
        v1 = -128
        _ = fi41
    }
    v1 += (v1 % (((v1 % (((-(v1))) | 1))) | 1))
    var v50 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v50
}
