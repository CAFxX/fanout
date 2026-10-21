package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int64; F1 uint64; F2 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 bool; F2 uint32 }
func genFn0(p0 uint, p1 int) bool {
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := []int{127}
        var v7 int = p1
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v8
    return false
}

func main() {
    var v1 bool = true
    var v2 uint32 = 535
    var v3 int32 = 0
    v4 := make(map[string]string)
    var v9 int8 = -119
    v10 := GenS1{F0: 338}
    v3 += (v3 * (v3 + v3))
    var v11 string = "foo bar"
    v12 := GenS2{F0: 42, F1: 127, F2: 128}
    for fi13 := 0; fi13 < 16; fi13++ {
        v2 -= 4294967295
        if v1 {
            v14 := v4[(((((v11 + v11) + (("a" + v11) + "")) + v11) + ((v11 + ("a" + ("\n\t" + "\n\t"))) + "")) + v11)]
            _ = v14
        }
        switch v11 {
        case "a":
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 float32 = -2.25
                _ = fi15
                _ = v16
            }
        case "hello":
            for fi17 := 0; fi17 < 12; fi17++ {
                fmt.Println(fi13, v11, fi17)
                _ = fi17
            }
        default:
            for fi18 := 0; fi18 < 7; fi18++ {
                v9 = v9
                _ = len(v4)
                _ = fi18
            }
        }
        _ = fi13
    }
    v19 := GenS3{F0: 255, F1: 1}
    fmt.Println(v3)
    if v1 {
        var v20 bool = (v9 <= v9)
        if ((1 <= 18446744073709551615) && (65535 < uint(v2))) {
            switch int(v3) {
            case -2147483648:
                if ((("foo bar" + ("\n\t" + "hello")) + ((v11 + (v11 + v11)) + v11)) != v11) {
                    if (false && v1) {
                        v4[((v11 + (v11 + (v11 + (v11 + (v11 + "hello"))))) + "\n\t")] = ((v11 + v11) + "foo bar")
                    }
                } else {
                    fmt.Println(v3, v20, v3)
                    v21 := GenS3{F0: 0, F1: 255}
                    v9 = v9
                    _ = v21
                }
            default:
                fmt.Println(v11, v3, v2)
            }
            var v22 int16 = 32767
            _ = v22
        } else {
            for fi23 := 0; fi23 < 20; fi23++ {
                _ = len(v4)
                if ("foo bar" != "a") {
                    for fi24 := 0; fi24 < 14; fi24++ {
                        v9 = (v9 ^ v9)
                        v4[("foo bar" + v11)] = (("\n\t" + (v11 + v11)) + v11)
                        var v25 int32 = (v3 | v3)
                        if false {
                            _ = len(v4)
                            var v26 int8 = v9
                            _ = v26
                        }
                        _ = fi24
                        _ = v25
                    }
                } else {
                    fmt.Println(v2, v9)
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v28 := GenS1{F0: (v2 - v2)}
                        _ = fi27
                        _ = v28
                    }
                }
                delete(v4, "hello")
                _ = fi23
            }
        }
        var v29 string = ""
        _ = v20
        _ = v29
    } else {
        if true {
            if v1 {
                v3 = (v3 * v3)
            } else {
                for fi30 := 0; fi30 < 3; fi30++ {
                    v31 := []int{(^(fi30)), (fi30 << ((4294967295) & 63)), 32767}
                    v2 = (v2 << ((900) & 63))
                    _ = fi30
                    _ = v31
                }
                switch (v11 + v11) {
                case "\n\t":
                    v3 += v3
                }
                v11 = ("hello" + "")
            }
            v32 := GenS4{F0: 923, F1: 0, F2: 1}
            _ = v32
        }
    }
    switch v11 {
    case "":
        for fi33 := 0; fi33 < 6; fi33++ {
            if v1 {
                if v1 {
                    v2 = 4294967295
                    for fi34 := 0; fi34 < 16; fi34++ {
                        if (790 > 614) {
                            v35 := GenS5{F0: (v11 + v11)}
                            v3 = int32(v2)
                            v11 = (v11 + ("" + v11))
                            _ = v35
                        } else {
                            var v36 uint8 = 0
                            var v37 int32 = v3
                            switch "" {
                            case "\n\t":
                                fmt.Println(v1, v1)
                            default:
                                fmt.Println(v3, v37, v1)
                            }
                            _ = v36
                            _ = v37
                        }
                        v38 := v4[(((v11 + ((v11 + v11) + v11)) + ((((v11 + "\n\t") + v11) + (v11 + (v11 + ((v11 + v11) + v11)))) + v11)) + (v11 + ((("hello" + (v11 + "")) + v11) + v11)))]
                        fmt.Println(v38, v1)
                        fmt.Println(v38, v2)
                        _ = fi34
                        _ = v38
                    }
                    v39 := GenS2{F0: 220, F1: v9, F2: 42}
                    _ = v39
                }
                v40 := []int{(-(fi33)), -487}
                fmt.Println(v9)
                _ = v40
            } else {
                v41 := []int{(fi33 / ((fi33) | 1))}
                var v42 int = fi33
                _ = v41
                _ = v42
            }
            v2 = (v2 ^ 4294967295)
            _ = fi33
        }
    case "a":
        v3 = v3
    case "hello":
        v43 := GenS5{F0: v11}
        _ = v43
    default:
        if v1 {
            v2 -= (v2 & 0)
        }
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        var v45 int64 = 9223372036854775807
        for fi46 := 0; fi46 < 18; fi46++ {
            v47 := v4[v11]
            _ = fi46
            _ = v47
        }
        v9 -= (v9 * 0)
        _ = fi44
        _ = v45
    }
    v48 := GenS3{F0: 42, F1: 574}
    var v49 int16 = -581
    v50 := GenS4{F0: -190, F1: 2, F2: 842}
    v11 = (v11 + v11)
    for fi51 := 0; fi51 < 2; fi51++ {
        fmt.Println(v2)
        var v52 uint16 = 854
        _ = fi51
        _ = v52
    }
    if v1 {
        for fi53 := 0; fi53 < 1; fi53++ {
            v1 = v1
            v54 := GenS6{F0: -186, F1: (v1 || v1), F2: 963}
            v55 := []int{fi53}
            _ = fi53
            _ = v54
            _ = v55
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v50
    _ = v9
}
