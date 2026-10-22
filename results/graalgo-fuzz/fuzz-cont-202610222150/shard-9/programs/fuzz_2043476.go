package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint; F2 int16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int64, p1 bool) string {
    var v7 int64 = (p0 >> ((2) & 63))
    if (2147483647 >= 0) {
        var v8 uint = 42
        var v9 float64 = float64((0.1 - 3.14159))
        v9 = (v9 + 1e10)
        _ = v8
        _ = v9
    } else {
        fmt.Println(p0, p1)
        if true {
            v10 := GenS1{F0: false}
            switch -1 {
            case -2147483648:
                v11 := GenS1{F0: (!(!false))}
                _ = v11
            }
            _ = v10
        }
        v12 := []int{-794, 4294967295, 9223372036854775807}
        _ = v12
    }
    p1 = p1
    if true {
        var v13 uint16 = 418
        _ = v13
    } else {
        var v14 bool = (877 > 4294967295)
        _ = v14
    }
    v7 -= (v7 * (p0 << ((255) & 63)))
    _ = p0
    _ = p1
    _ = v7
    return "a"
}
func genFn1(p0 bool, p1 int32, p2 int8) int64 {
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 uint32 = 42
        v16 = (^(v16))
        _ = fi15
        _ = v16
    }
    if (-482 != -431) {
        for fi17 := 0; fi17 < 5; fi17++ {
            p0 = false
            switch fi17 {
            case 896:
                p0 = (!(!(false || p0)))
            case 127:
                for fi18 := 0; fi18 < 8; fi18++ {
                    v19 := GenS2{F0: 0}
                    _ = fi18
                    _ = v19
                }
            case -1:
                p2 += p2
            default:
                p2 = -128
            }
            _ = fi17
        }
        var v20 bool = (675 == -300)
        _ = v20
    } else {
        if p0 {
            switch 42 {
            case 32767:
                fmt.Println(p1)
            case -524:
                fmt.Println(p2, p1)
            default:
                for fi21 := 0; fi21 < 20; fi21++ {
                    p2 = p2
                    if p0 {
                        p0 = (true || (!(32767 != -32768)))
                        var v22 int8 = (^(p2))
                        v23 := []int{-2147483648, fi21, fi21, int(p1)}
                        _ = v22
                        _ = v23
                    }
                    _ = fi21
                }
            }
            switch "foo bar" {
            case "foo bar":
                var v24 uint16 = 42
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 float64 = (float64(0.1) / 1.5)
                v27 := GenS1{F0: p0}
                var v28 uint16 = 2
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        }
        v29 := []int{}
        _ = v29
    }
    if p0 {
        for fi30 := 0; fi30 < 3; fi30++ {
            if p0 {
                if ((0.0 < -2.25) || p0) {
                    if p0 {
                        v31 := []int{(fi30 + 128), fi30, 128, 2147483647}
                        _ = v31
                    }
                }
                fmt.Println(p2, fi30)
                v32 := []int{1}
                _ = v32
            }
            if (!(!p0)) {
                if p0 {
                    for fi33 := 0; fi33 < 20; fi33++ {
                        v34 := GenS3{F0: 0, F1: 32767}
                        fmt.Println(fi30, p1)
                        p2 = 108
                        _ = fi33
                        _ = v34
                    }
                    fmt.Println(p2, p1)
                } else {
                    var v35 uint32 = 2
                    if p0 {
                        var v36 bool = p0
                        _ = v36
                    }
                    _ = v35
                }
                fmt.Println(p2, p2)
            } else {
                var v37 int16 = -1
                _ = v37
            }
            fmt.Println(fi30, p0)
            _ = fi30
        }
        switch 127 {
        case 701:
            var v38 uint32 = 428
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        p0 = (-32768 != -38)
        var v40 float64 = 0.0
        var v41 int64 = 541
        _ = fi39
        _ = v40
        _ = v41
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 int16 = 1
    v2 := []string{"", "foo bar", "a", ("\n\t" + "")}
    v3 := make(map[int]bool)
    var v4 string = "\n\t"
    v5 := []string{v4, "a"}
    var v6 int = -1
    v42 := GenS4{F0: 357, F1: 18446744073709551615, F2: v1}
    var v43 uint = 387
    var v44 float64 = 1.5
    fmt.Println(v1)
    fmt.Println(v1, v4)
    if true {
        for fi45 := 0; fi45 < 1; fi45++ {
            var v46 uint16 = 0
            fmt.Println(v6, v4, v44)
            _ = fi45
            _ = v46
        }
    } else {
        _ = len(v3)
    }
    for fi47 := 0; fi47 < 19; fi47++ {
        for fi48 := 0; fi48 < 12; fi48++ {
            if (!(1 <= -654)) {
                v43 = v43
                v44 = v44
            } else {
                switch (v4 + "a") {
                case "\n\t":
                    switch v4 {
                    case "a":
                        v4 = "\n\t"
                    default:
                        v3[(fi47 ^ (^(v6)))] = (((v43 << (((v43 >> ((65535) & 63))) & 63)) >= v43) || false)
                    }
                }
                switch (("foo bar" + v4) + "\n\t") {
                case "hello":
                    v43 -= v43
                case "\n\t":
                    fmt.Println(fi47)
                default:
                    var v49 bool = (2 <= 42)
                    _ = v49
                }
                var v50 uint8 = 127
                _ = v50
            }
            v4 = "\n\t"
            _ = fi48
        }
        v51 := GenS5{F0: 1}
        _ = fi47
        _ = v51
    }
    var v52 uint8 = 128
    var v53 int64 = 1
    fmt.Println(v52, v53, v6)
    if len(v2) > 0 {
        v2[0] = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
