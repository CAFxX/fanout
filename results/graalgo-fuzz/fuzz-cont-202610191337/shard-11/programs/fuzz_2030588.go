package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint32; F2 int64 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint; F1 bool; F2 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64; F1 int32; F2 uint16 }
func genFn0(p0 uint) uint32 {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := []int{fi4, fi4, fi4, 221}
        v6 := GenS1{F0: 128, F1: (^(fi4)), F2: 207}
        _ = fi4
        _ = v5
        _ = v6
    }
    if (false && (42 > 0)) {
        var v7 int = 4294967295
        if true {
            var v8 string = "\n\t"
            p0 += (p0 << ((p0) & 63))
            _ = v8
        }
        fmt.Println(v7, p0)
        _ = v7
    } else {
        if (false && (2 == 418)) {
            v9 := GenS1{F0: 251, F1: 128, F2: 127}
            _ = v9
        }
        if false {
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := GenS2{F0: 1, F1: 1, F2: 597}
                var v12 int16 = 32767
                v13 := []int{fi10}
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            if false {
                p0 += (p0 ^ p0)
                var v14 int16 = -32768
                _ = v14
            }
        } else {
            var v15 int16 = -5
            switch ("\n\t" + (("a" + ("\n\t" + "a")) + ("a" + "foo bar"))) {
            case "hello":
                fmt.Println(v15, v15)
            case "foo bar":
                var v16 string = "hello"
                _ = v16
            }
            _ = v15
        }
    }
    if true {
        if ((!(uint8(p0) > 42)) || true) {
            v17 := []int{9223372036854775807}
            _ = v17
        } else {
            if ((255 > 221) && false) {
                var v18 uint8 = 1
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := []int{}
            if true {
                p0 = p0
                p0 = (p0 << (((p0 * p0)) & 63))
                var v21 int16 = 32767
                _ = v21
            } else {
                v22 := GenS3{F0: 9223372036854775807}
                switch "hello" {
                case "foo bar":
                    var v23 string = "a"
                    _ = v23
                case "hello":
                    v20 = append(v20, fi19)
                }
                _ = v22
            }
            for fi24 := 0; fi24 < 3; fi24++ {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                var v25 float32 = -2.25
                _ = fi24
                _ = v25
            }
            _ = fi19
            _ = v20
        }
        p0 = 928
    }
    _ = p0
    return 4294967295
}
func genFn1(p0 bool) int {
    var v26 float64 = 0.1
    fmt.Println(p0, p0)
    if (728 == 802) {
        v27 := []int{}
        _ = v27
    }
    p0 = p0
    _ = p0
    _ = v26
    return 42
}

func main() {
    var v1 uint64 = 880
    v2 := make(map[int]string)
    v3 := []uint32{}
    var v28 int32 = -491
    v3 = append(v3, 2147483648)
    v29 := v2[9223372036854775807]
    v30 := GenS4{F0: 2147483647, F1: 1, F2: 0}
    v29 = (v29 + "\n\t")
    v3 = append(v3, 65535)
    if (!false) {
        for fi31 := 0; fi31 < 19; fi31++ {
            v29 = "a"
            v29 = v29
            _ = fi31
        }
        v32 := GenS5{F0: -1, F1: genFn1((!true))}
        for fi33 := 0; fi33 < 5; fi33++ {
            fmt.Println(v29, v29, v28)
            v28 += v28
            v2[(fi33 | 157)] = "foo bar"
            _ = fi33
        }
        _ = v32
    } else {
        var v34 string = v29
        _ = v34
    }
    var v35 int32 = (v28 >> ((293) & 63))
    v28 = v35
    if false {
        for fi36 := 0; fi36 < 3; fi36++ {
            switch genFn1(false) {
            case 9223372036854775807:
                _ = len(v3)
            case 2147483647:
                v37 := GenS6{F0: 435, F1: true, F2: 1}
                _ = v37
            default:
                v38 := GenS7{F0: (fi36 % (((-(fi36))) | 1))}
                _ = v38
            }
            _ = fi36
        }
        for fi39 := 0; fi39 < 17; fi39++ {
            for fi40 := 0; fi40 < 9; fi40++ {
                v29 = (v29 + v29)
                _ = fi40
            }
            _ = fi39
        }
        if true {
            switch v29 {
            case "\n\t":
                v29 = "foo bar"
            case "":
                if (false || (!true)) {
                    var v41 uint8 = 127
                    if false {
                        var v42 uint8 = v41
                        _ = v42
                    }
                    _ = v41
                }
            case "a":
                switch 2147483647 {
                case -9223372036854775808:
                    v43 := GenS7{F0: 9223372036854775807}
                    _ = v43
                }
            }
            for fi44 := 0; fi44 < 16; fi44++ {
                v1 += v1
                v28 = 2147483647
                for fi45 := 0; fi45 < 3; fi45++ {
                    v29 = (v29 + (v29 + v29))
                    var v46 string = (((v29 + v29) + v29) + (v29 + v29))
                    _ = fi45
                    _ = v46
                }
                _ = fi44
            }
            var v47 uint8 = 2
            _ = v47
        } else {
            v48 := GenS2{F0: v1, F1: 237, F2: 340}
            v35 = v28
            _ = v48
        }
    } else {
        v3 = append(v3, 2)
        for fi49 := 0; fi49 < 10; fi49++ {
            v50 := GenS8{F0: (^(v1)), F1: v35, F2: 42}
            var v51 int = fi49
            _ = fi49
            _ = v50
            _ = v51
        }
    }
    for fi52 := 0; fi52 < 4; fi52++ {
        var v53 uint32 = 214
        v54 := GenS6{F0: 18446744073709551615, F1: (-313 >= genFn1((!true))), F2: -9223372036854775808}
        var v55 uint64 = v1
        _ = fi52
        _ = v53
        _ = v54
        _ = v55
    }
    v1 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
}
