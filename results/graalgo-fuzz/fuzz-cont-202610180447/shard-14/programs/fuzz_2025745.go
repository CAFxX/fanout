package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint) uint16 {
    v7 := GenS1{F0: 1, F1: 127}
    var v8 uint8 = 1
    var v9 float64 = 1e10
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 16
}
func genFn1(p0 int8, p1 int16) string {
    var v10 int8 = 1
    for fi11 := 0; fi11 < 15; fi11++ {
        if true {
            fmt.Println(fi11, fi11, v10)
        }
        switch "foo bar" {
        case "foo bar":
            p0 = int8(fi11)
        default:
            if false {
                for fi12 := 0; fi12 < 1; fi12++ {
                    if false {
                        switch -293 {
                        case 32768:
                            p1 = p1
                        case 32767:
                            fmt.Println(p0)
                        case -1:
                            fmt.Println(fi12, p1, p1)
                        default:
                            var v13 float64 = (((1.5 / 3.14159) - 1.5) + 1.5)
                            _ = v13
                        }
                    }
                    _ = fi12
                }
                v14 := []int{fi11}
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 int64 = -1
                    fmt.Println(v10, fi15)
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return "\n\t"
}

func main() {
    var v1 int64 = -452
    var v2 uint32 = 65535
    var v3 int8 = 1
    v4 := []string{"foo bar", "foo bar"}
    var v5 uint8 = 255
    var v6 uint32 = 308
    switch "" {
    case "a":
        switch 529 {
        case 116:
            v17 := GenS2{F0: 386, F1: "foo bar"}
            _ = v17
        case -2147483648:
            var v18 int8 = v3
            _ = v18
        case 42:
            v19 := GenS1{F0: v3, F1: (v3 << ((143) & 63))}
            _ = v19
        default:
            for fi20 := 0; fi20 < 11; fi20++ {
                v21 := GenS1{F0: v3, F1: v3}
                _ = fi20
                _ = v21
            }
        }
    case "":
        v3 = v3
    }
    v6 = v2
    v6 += (v2 - v2)
    switch "\n\t" {
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
    }
    switch "foo bar" {
    case "a":
        var v22 uint8 = (v5 & v5)
        _ = v22
    case "foo bar":
        v6 += v2
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    default:
        if true {
            var v23 int32 = 2147483647
            for fi24 := 0; fi24 < 16; fi24++ {
                fmt.Println(v6)
                for fi25 := 0; fi25 < 5; fi25++ {
                    if true {
                        fmt.Println(fi25)
                        var v26 int16 = 773
                        fmt.Println(v6, fi24, v1)
                        _ = v26
                    }
                    if true {
                        v3 = (^(v3))
                        v27 := GenS3{F0: (("" + "hello") + "foo bar"), F1: 65535}
                        _ = v27
                    } else {
                        _ = len(v4)
                        v1 += -9223372036854775808
                        v28 := GenS4{F0: (v5 + 69)}
                        if true {
                            var v29 int8 = (v3 / (((v3 % ((0) | 1))) | 1))
                            _ = v29
                        } else {
                            switch (fi24 ^ -619) {
                            case 0:
                                var v30 int64 = (v1 + v1)
                                _ = v30
                            }
                            var v31 bool = (false && (884 <= v2))
                            v2 = v6
                            _ = v31
                        }
                        _ = v28
                    }
                    v3 += v3
                    _ = fi25
                }
                var v32 bool = true
                _ = fi24
                _ = v32
            }
            _ = v23
        }
    }
    fmt.Println(v2, v2)
    v6 += (v2 / ((v6) | 1))
    for fi33 := 0; fi33 < 14; fi33++ {
        v34 := GenS2{F0: 2, F1: ((("\n\t" + "") + ("foo bar" + "a")) + ("a" + ""))}
        _ = fi33
        _ = v34
    }
    var v35 int = 32767
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "a" {
    case "":
        v36 := GenS3{F0: "foo bar", F1: 933}
        _ = v36
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        if false {
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 int8 = v3
                v5 = v5
                _ = fi37
                _ = v38
            }
        }
    }
    var v39 string = "a"
    var v40 bool = (800 == 190)
    if false {
        v3 -= v3
        var v41 int32 = 0
        _ = v41
    }
    v1 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
