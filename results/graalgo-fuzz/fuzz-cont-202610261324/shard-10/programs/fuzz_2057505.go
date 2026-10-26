package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 bool; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 string; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint16; F1 int8; F2 uint16 }
func genFn0(p0 uint64) int32 {
    if true {
        p0 = (p0 ^ 18446744073709551615)
        p0 = (p0 + 852)
        p0 = p0
    }
    if true {
        v7 := GenS1{F0: -1}
        _ = v7
    } else {
        var v8 string = "foo bar"
        v9 := []int{}
        p0 = 765
        _ = v8
        _ = v9
    }
    v10 := []int{}
    _ = len(v10)
    _ = p0
    _ = v10
    return 1
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v4 string = ("hello" + "foo bar")
    var v5 uint32 = 65535
    var v6 float32 = (0.1 * 1e10)
    v1 += v1
    var v11 string = v4
    v1 -= float64((v6 * float32((0.1 / 1.5))))
    v3[v11] = ((((v4 + v4) + v11) + "\n\t") + v4)
    v6 += (((((((v6 / v6) * (0.0 * ((v6 * 0.1) - v6))) / 1e10) * float32(v1)) + -2.25) * v6) + v6)
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int64 = -460
        var v14 uint64 = 0
        _ = len(v2)
        _ = fi12
        _ = v13
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 7; fi15++ {
            v11 = "foo bar"
            _ = fi15
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            for fi17 := 0; fi17 < 14; fi17++ {
                v18 := GenS2{F0: (fi17 & (fi17 >> ((842) & 63))), F1: (686 == 0), F2: 42}
                switch (fi16 + (fi16 & fi17)) {
                case 434:
                    v4 = v11
                case 2147483647:
                    fmt.Println(v5, fi17, v4)
                case 32768:
                    v19 := GenS2{F0: (-(fi16)), F1: true, F2: 18446744073709551615}
                    _ = v19
                default:
                    var v20 int8 = 0
                    _ = v20
                }
                v4 = ("foo bar" + "hello")
                _ = fi17
                _ = v18
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS2{F0: fi21, F1: true, F2: 0}
                var v23 bool = (false && false)
                switch (fi21 & 185) {
                case 349:
                    if (v23 || (0 >= 801)) {
                        v6 += 1e10
                        if v23 {
                            var v24 string = "a"
                            _ = v24
                        }
                        v5 = 0
                        v25 := GenS1{F0: 127}
                        var v26 uint8 = 1
                        _ = v25
                        _ = v26
                    }
                case 560:
                    v5 += v5
                case 2147483647:
                    v11 = (v4 + "\n\t")
                }
                fmt.Println(fi16, v4)
                v6 = (1.5 - v6)
                var v27 string = v4
                _ = fi21
                _ = v22
                _ = v23
                _ = v27
            }
            _ = fi16
        }
    } else {
        if true {
            for fi28 := 0; fi28 < 13; fi28++ {
                v29 := []int{fi28, (fi28 + 1), fi28, -9223372036854775808}
                for fi30 := 0; fi30 < 20; fi30++ {
                    if len(v29) > 0 {
                        v29[0] = fi30
                    }
                    var v31 float64 = (v1 - ((v1 * v1) / (v1 / v1)))
                    var v32 uint64 = 30
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = fi28
                _ = v29
            }
            v33 := GenS3{F0: (!false)}
            _ = v33
        }
        v5 += 892
    }
    v34 := GenS2{F0: 0, F1: (18446744073709551615 >= 18446744073709551615), F2: 862}
    v35 := []int{2147483647, 4294967295, -497}
    var v36 uint = 42
    var v37 int16 = 890
    var v38 int32 = -2147483648
    v11 = ("foo bar" + v11)
    v39 := GenS4{F0: true, F1: v36}
    fmt.Println(v37, v6)
    var v40 float32 = v6
    var v41 int8 = -16
    if true {
        for fi42 := 0; fi42 < 8; fi42++ {
            for fi43 := 0; fi43 < 10; fi43++ {
                var v44 int32 = -107
                v36 = v36
                switch "a" {
                case "":
                    var v45 uint = 65535
                    v46 := GenS5{F0: (v37 & (-(v37))), F1: v44}
                    var v47 float64 = (v1 / v1)
                    _ = v45
                    _ = v46
                    _ = v47
                case "foo bar":
                    v38 -= v38
                }
                _ = fi43
                _ = v44
            }
            if (v4 > "foo bar") {
                _ = len(v35)
                if (true || true) {
                    fmt.Println(v5, v6, v37)
                    v48 := GenS6{F0: (("hello" + "hello") + "\n\t"), F1: 2, F2: 9223372036854775807}
                    _ = v48
                } else {
                    v49 := GenS2{F0: fi42, F1: true, F2: (v36 / ((v36) | 1))}
                    _ = v49
                }
            } else {
                if true {
                    v50 := GenS1{F0: v41}
                    fmt.Println(v37, v38)
                    _ = v50
                }
                switch int(v36) {
                case 1:
                    v51 := GenS7{F0: 591, F1: v41, F2: 2}
                    _ = v51
                case 419:
                    v52 := GenS6{F0: v11, F1: 1, F2: 0}
                    _ = v52
                case 2147483647:
                    v53 := GenS7{F0: uint16(v36), F1: 42, F2: 2}
                    _ = v53
                default:
                    if (false || (v4 <= "foo bar")) {
                        v54 := GenS5{F0: v37, F1: 2147483647}
                        for fi55 := 0; fi55 < 5; fi55++ {
                            _ = len(v35)
                            fmt.Println(v41, v11, v5)
                            _ = fi55
                        }
                        fmt.Println(v1, fi42, v6)
                        _ = v54
                    }
                }
            }
            _ = fi42
        }
        var v56 uint8 = 42
        for fi57 := 0; fi57 < 17; fi57++ {
            var v58 uint32 = (v5 % ((v5) | 1))
            _ = fi57
            _ = v58
        }
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
