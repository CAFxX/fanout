package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 string; F1 int32 }
func genFn0(p0 uint16, p1 uint, p2 uint32) int16 {
    if ((-360 >= 42) || (4294967295 == p1)) {
        v6 := GenS1{F0: p2}
        _ = v6
    } else {
        switch -9223372036854775808 {
        case 32767:
            v7 := GenS1{F0: p2}
            _ = v7
        case 42:
            v8 := []int{42, 32768, -1}
            _ = v8
        case -2147483648:
            p1 -= p1
        default:
            if (!(true || (true && true))) {
                var v9 uint8 = 127
                _ = v9
            } else {
                switch 42 {
                case 314:
                    p1 -= (p1 << ((p1) & 63))
                default:
                    v10 := []int{-829}
                    _ = v10
                }
                v11 := GenS2{F0: 142, F1: (p2 % ((p2) | 1))}
                p0 -= p0
                _ = v11
            }
        }
        v12 := GenS2{F0: p1, F1: p2}
        p2 -= uint32(p1)
        _ = v12
    }
    v13 := GenS3{F0: 42}
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 42
}
func genFn1(p0 uint8, p1 uint) int32 {
    v14 := []int{2147483647, -999}
    _ = len(v14)
    fmt.Println(p1)
    v15 := GenS4{F0: 128, F1: 643}
    var v16 uint16 = 2
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return -1
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t", ("foo bar" + "hello"), ("a" + "")}
    var v4 int16 = 538
    var v5 int32 = -215
    if (!((!(1 <= v4)) || (!(-9223372036854775808 > -1)))) {
        if false {
            var v17 uint = 0
            _ = v17
        } else {
            for fi18 := 0; fi18 < 11; fi18++ {
                for fi19 := 0; fi19 < 4; fi19++ {
                    if (!((-(fi19)) <= (-(fi19)))) {
                        v20 := GenS4{F0: 2147483648, F1: 985}
                        var v21 float32 = 1e10
                        _ = v20
                        _ = v21
                    } else {
                        var v22 float32 = 0.0
                        v22 += float32(float64(v22))
                        v3 = append(v3, "foo bar")
                        _ = v22
                    }
                    var v23 uint64 = 14
                    var v24 uint64 = 18446744073709551615
                    v1 = 851
                    var v25 float32 = 1e10
                    _ = fi19
                    _ = v23
                    _ = v24
                    _ = v25
                }
                var v26 uint64 = 222
                for fi27 := 0; fi27 < 13; fi27++ {
                    switch (^(fi18)) {
                    case 1:
                        var v28 int8 = 42
                        _ = v28
                    default:
                        v29 := GenS5{F0: (fi27 ^ fi18), F1: "hello"}
                        _ = v29
                    }
                    _ = fi27
                }
                _ = fi18
                _ = v26
            }
        }
        if ((2147483648 < 2) || (0.1 >= (0.0 - (float32(1e10) / -2.25)))) {
            switch "" {
            case "":
                _ = len(v3)
            case "foo bar":
                v30 := GenS6{F0: (!("foo bar" != "")), F1: 909}
                _ = v30
            case "\n\t":
                if ((!((-2.25 == 1.5) || (true && (255 != 922)))) || (0 <= 427)) {
                    var v31 uint = 42
                    _ = v31
                } else {
                    var v32 uint16 = 0
                    _ = v32
                }
            }
            v33 := v2[649]
            for fi34 := 0; fi34 < 2; fi34++ {
                fmt.Println(v33)
                v35 := GenS4{F0: 4294967295, F1: 1}
                _ = fi34
                _ = v35
            }
            _ = v33
        } else {
            var v36 uint8 = 2
            for fi37 := 0; fi37 < 11; fi37++ {
                v4 = (v4 - -1)
                _ = fi37
            }
            _ = v36
        }
        var v38 bool = ((true || true) && (!(42 >= 32)))
        _ = v38
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        v4 += v4
    } else {
        v39 := GenS4{F0: 494, F1: 1}
        if (18446744073709551615 <= 1) {
            fmt.Println(v1)
        } else {
            v3 = append(v3, ((("a" + "") + ("\n\t" + "a")) + "\n\t"))
        }
        for fi40 := 0; fi40 < 5; fi40++ {
            v3 = append(v3, "a")
            var v41 uint = 18446744073709551615
            _ = fi40
            _ = v41
        }
        _ = v39
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 += (v4 + genFn0(255, 2, 302))
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "a")
    }
    v1 += -32768
    v4 = v1
    if false {
        var v42 uint8 = 127
        _ = v42
    }
    var v43 float64 = (1.5 - (3.14159 - -2.25))
    var v44 string = ""
    if (344 < 2147483648) {
        v43 = v43
    } else {
        if (true && (int64(v1) != 9223372036854775807)) {
            if (!false) {
                if true {
                    v44 = (v44 + v44)
                    fmt.Println(v4, v4)
                    if (!true) {
                        v45 := GenS2{F0: 65535, F1: 65535}
                        _ = v45
                    } else {
                        v1 = v1
                    }
                } else {
                    v46 := GenS7{F0: 127, F1: int16(v5)}
                    v47 := GenS2{F0: 65535, F1: 117}
                    _ = v46
                    _ = v47
                }
                v48 := GenS8{F0: "foo bar", F1: genFn1(95, 18446744073709551615)}
                _ = v48
            } else {
                v43 -= ((v43 / ((float64(-2.25) / 0.1) - 3.14159)) + 3.14159)
                var v49 bool = (!false)
                v50 := GenS4{F0: 4294967295, F1: 4294967295}
                _ = v49
                _ = v50
            }
            var v51 uint64 = 2
            _ = v51
        }
    }
    switch -2147483648 {
    case 9223372036854775807:
        for fi52 := 0; fi52 < 20; fi52++ {
            for fi53 := 0; fi53 < 17; fi53++ {
                var v54 uint64 = 4294967295
                _ = fi53
                _ = v54
            }
            _ = fi52
        }
    case 331:
        v43 -= (3.14159 + v43)
    case 0:
        for fi55 := 0; fi55 < 18; fi55++ {
            var v56 float32 = -2.25
            _ = fi55
            _ = v56
        }
    default:
        var v57 string = ("foo bar" + v44)
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
