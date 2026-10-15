package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint32; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0(p0 uint) uint32 {
    p0 = (p0 - 18446744073709551615)
    var v6 int = 983
    var v7 int16 = 928
    p0 += p0
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []string{}
    v2 := []int{9223372036854775807, 2147483647, 231, -650}
    var v3 float32 = 3.14159
    v4 := make(map[string]bool)
    v5 := []int{9223372036854775807, -9223372036854775808, 4294967295, -1}
    fmt.Println(v3)
    switch 282 {
    case -2147483648:
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int32 = 0
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi8
            _ = v9
        }
    case -9223372036854775808:
        v3 = 1e10
    case 982:
        for fi10 := 0; fi10 < 1; fi10++ {
            if ((0.1 + ((1e10 * (1e10 + v3)) / 1e10)) < v3) {
                _ = len(v1)
                if false {
                    v3 -= 1e10
                    v3 -= v3
                    var v11 uint = 65535
                    _ = v11
                } else {
                    v3 = v3
                }
            }
            _ = fi10
        }
    default:
        v12 := GenS1{F0: 9223372036854775807}
        _ = v12
    }
    switch ("\n\t" + "") {
    case "\n\t":
        if false {
            var v13 float64 = (float64(1.5) - 1.5)
            var v14 int32 = 727
            if ((42 < 4294967295) && ((v14 + (v14 + -1)) != v14)) {
                v15 := GenS2{F0: 42, F1: 299, F2: "foo bar"}
                v13 += v13
                if true {
                    v16 := GenS1{F0: -856}
                    var v17 string = "a"
                    _ = v16
                    _ = v17
                }
                _ = v15
            }
            _ = v13
            _ = v14
        } else {
            v3 = (float32(1e10) * 3.14159)
        }
    case "a":
        v18 := GenS1{F0: 42}
        _ = v18
    case "foo bar":
        fmt.Println(v3)
    }
    var v19 float32 = (0.1 - v3)
    var v20 uint = 1
    var v21 uint64 = 42
    var v22 int32 = 860
    switch 32767 {
    case -687:
        for fi23 := 0; fi23 < 8; fi23++ {
            for fi24 := 0; fi24 < 9; fi24++ {
                v20 = (v20 / ((2) | 1))
                var v25 int64 = 206
                v26 := GenS3{F0: 9223372036854775807}
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi23
        }
    case 1:
        var v27 int16 = -1
        _ = v27
    default:
        v19 -= v19
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 5; fi29++ {
            switch "\n\t" {
            case "hello":
                var v30 string = "\n\t"
                _ = v30
            default:
                v31 := GenS4{F0: 32767}
                _ = v31
            }
            _ = fi29
        }
        delete(v4, "hello")
        var v32 uint64 = (v21 & 2)
        _ = fi28
        _ = v32
    }
    switch 0 {
    case -1:
        delete(v4, "hello")
    case -2147483648:
        for fi33 := 0; fi33 < 11; fi33++ {
            _ = len(v4)
            _ = fi33
        }
    case 0:
        for fi34 := 0; fi34 < 6; fi34++ {
            if (fi34 <= (-(fi34))) {
                for fi35 := 0; fi35 < 4; fi35++ {
                    v19 -= v3
                    v36 := GenS1{F0: 9223372036854775807}
                    var v37 float64 = (0.0 / (1.5 + ((1e10 + float64((v19 / v19))) * ((float64(v3) - 0.1) + ((((float64(v3) / -2.25) / 1.0) * float64(v3)) + -2.25)))))
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 -= (v3 * v19)
            }
            if true {
                v20 += v20
                v22 += (v22 >> (((v20 / (((^(v20))) | 1))) & 63))
                v20 = v20
            } else {
                var v38 float32 = (((v3 * 3.14159) * (v3 / v19)) * v19)
                v38 = -2.25
                var v39 int = (fi34 << ((255) & 63))
                _ = v38
                _ = v39
            }
            v20 = (v20 * 4294967295)
            _ = fi34
        }
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        var v41 float64 = 1.5
        var v42 bool = (!(!(18446744073709551615 < v21)))
        var v43 float32 = (v3 + v19)
        _ = fi40
        _ = v41
        _ = v42
        _ = v43
    }
    if false {
        v3 += (v3 - ((((float32(-2.25) / (1.5 * v3)) * 1e10) + v3) - 3.14159))
        fmt.Println(v19, v3, v19)
    }
    switch ("hello" + ("foo bar" + "")) {
    case "hello":
        var v44 uint32 = uint32(v21)
        _ = v44
    case "foo bar":
        v45 := GenS5{F0: 4294967295, F1: 0, F2: (^(v21))}
        _ = v45
    default:
        for fi46 := 0; fi46 < 20; fi46++ {
            if (!false) {
                v4[("hello" + ("\n\t" + ""))] = true
                for fi47 := 0; fi47 < 5; fi47++ {
                    _ = len(v4)
                    v20 = 255
                    if len(v2) > 0 {
                        v2[0] = fi46
                    }
                    _ = fi47
                }
                if (!((true || (255 <= 186)) || true)) {
                    var v48 uint32 = 2147483648
                    v20 = v20
                    v48 += (v48 + 1)
                    _ = v48
                }
            } else {
                var v49 uint16 = 65535
                _ = v49
            }
            fmt.Println(v21)
            _ = fi46
        }
    }
    var v50 string = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}
