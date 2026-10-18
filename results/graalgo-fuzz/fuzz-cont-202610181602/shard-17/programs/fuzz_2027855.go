package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 bool }
type GenS2 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 string; F2 uint16 }

func main() {
    var v1 int16 = 564
    var v2 int16 = -245
    var v3 uint32 = 4294967295
    v2 += 219
    var v4 int8 = -114
    v5 := []int{}
    var v6 float64 = 0.0
    v7 := GenS1{F0: v1, F1: (^(v3)), F2: true}
    var v8 bool = true
    v9 := GenS1{F0: v1, F1: (v3 ^ 333), F2: (v6 == 1e10)}
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v6, v1)
        var v11 string = "a"
        _ = len(v5)
        _ = fi10
        _ = v11
    }
    var v12 int8 = (v4 ^ (-(v4)))
    if (0 < -1) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if v8 {
            var v13 uint8 = 4
            for fi14 := 0; fi14 < 18; fi14++ {
                v1 = (-(v1))
                if false {
                    var v15 bool = ((false || (v8 && false)) || v8)
                    var v16 bool = (!v15)
                    _ = v15
                    _ = v16
                }
                if len(v5) > 0 {
                    v5[0] = (fi14 << ((4294967295) & 63))
                }
                _ = fi14
            }
            for fi17 := 0; fi17 < 20; fi17++ {
                if true {
                    v18 := GenS1{F0: (v1 / ((v1) | 1)), F1: v3, F2: v8}
                    for fi19 := 0; fi19 < 8; fi19++ {
                        if len(v5) > 0 {
                            v5[0] = (fi19 * 32767)
                        }
                        if ("a" <= "a") {
                            v1 -= v2
                            var v20 int64 = -764
                            var v21 uint64 = 18446744073709551615
                            _ = v20
                            _ = v21
                        }
                        _ = fi19
                    }
                    v2 = -32768
                    _ = v18
                } else {
                    for fi22 := 0; fi22 < 9; fi22++ {
                        v1 = (-(v1))
                        var v23 uint = 4294967295
                        _ = fi22
                        _ = v23
                    }
                }
                _ = fi17
            }
            _ = v13
        } else {
            if ((!false) || false) {
                for fi24 := 0; fi24 < 1; fi24++ {
                    var v25 int = (fi24 % ((9223372036854775807) | 1))
                    switch "a" {
                    case "a":
                        v5 = append(v5, (v25 % ((v25) | 1)))
                    }
                    _ = fi24
                    _ = v25
                }
                var v26 int = 42
                _ = v26
            } else {
                v27 := GenS2{F0: 0, F1: 2, F2: 275}
                var v28 int64 = -571
                v12 += (v12 >> ((42) & 63))
                _ = v27
                _ = v28
            }
            fmt.Println(v12)
        }
    } else {
        fmt.Println(v1, v12)
        v29 := GenS1{F0: v1, F1: 370, F2: false}
        _ = v29
    }
    fmt.Println(v1)
    var v30 int8 = (v12 ^ 42)
    if v8 {
        if (v8 && false) {
            if true {
                var v31 bool = (((^(v2)) != (v1 << ((180) & 63))) || ("hello" <= "a"))
                v2 -= (^(v2))
                var v32 int = 495
                _ = v31
                _ = v32
            } else {
                if false {
                    v5 = append(v5, 1)
                }
                if len(v5) > 0 {
                    v5[0] = 9223372036854775807
                }
                if ("\n\t" > ("\n\t" + ("foo bar" + (("a" + "a") + (("hello" + ("foo bar" + ("\n\t" + "foo bar"))) + "foo bar"))))) {
                    if (v3 == (v3 % ((v3) | 1))) {
                        v33 := GenS1{F0: 32767, F1: (^(v3)), F2: (4294967295 <= 540)}
                        var v34 uint16 = 0
                        v30 = -128
                        _ = v33
                        _ = v34
                    }
                    v6 = (v6 / (float64(((float32((1e10 / v6)) * 1e10) - 0.0)) * -2.25))
                } else {
                    v30 -= 100
                }
            }
            if (v8 || (-2.25 <= 0.1)) {
                fmt.Println(v12)
                switch -754 {
                case -329:
                    for fi35 := 0; fi35 < 4; fi35++ {
                        switch (fi35 | (-(fi35))) {
                        case 9223372036854775807:
                            if v8 {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                for fi36 := 0; fi36 < 8; fi36++ {
                                    if v8 {
                                        v37 := GenS3{F0: int64(v3)}
                                        _ = v37
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = fi35
                                    }
                                    _ = fi36
                                }
                            } else {
                                var v38 int64 = int64(fi35)
                                var v39 int8 = 65
                                _ = v38
                                _ = v39
                            }
                        case 32767:
                            fmt.Println(v6)
                        default:
                            var v40 int8 = (v12 - (v12 >> ((95) & 63)))
                            _ = v40
                        }
                        v8 = (!true)
                        _ = fi35
                    }
                case -9223372036854775808:
                    if ((int64(v3) > 9223372036854775807) || v8) {
                        _ = len(v5)
                        for fi41 := 0; fi41 < 17; fi41++ {
                            v42 := GenS1{F0: (v2 ^ v1), F1: (v3 + v3), F2: false}
                            _ = fi41
                            _ = v42
                        }
                        for fi43 := 0; fi43 < 5; fi43++ {
                            if v8 {
                                v12 -= v4
                                v3 = v3
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            _ = fi43
                        }
                    }
                default:
                    for fi44 := 0; fi44 < 16; fi44++ {
                        v45 := GenS4{F0: "foo bar", F1: "hello", F2: 255}
                        var v46 bool = true
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                }
            } else {
                _ = len(v5)
            }
        } else {
            fmt.Println(v6, v12, v6)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
