package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint8; F1 int; F2 int32 }

func main() {
    var v1 int16 = 0
    var v2 int64 = 130
    v3 := []string{"a", ("hello" + "a"), "", "foo bar"}
    var v4 int = 42
    var v5 uint8 = 0
    v6 := GenS1{F0: 0}
    _ = len(v3)
    v7 := GenS2{F0: v4}
    v1 += (^(v1))
    var v8 int8 = 0
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(v2)
        _ = fi9
    }
    v8 -= (v8 * (v8 ^ (v8 - (v8 ^ (-(v8))))))
    v2 = v2
    fmt.Println(v8, v4)
    var v10 float64 = float64(3.14159)
    if ((v4 >> ((42) & 63)) != v4) {
        if true {
            var v11 int16 = (v1 + (v1 % ((v1) | 1)))
            if false {
                if ((v1 & int16(v4)) != (-(v1))) {
                    v12 := GenS1{F0: 1}
                    _ = v12
                } else {
                    v4 += v4
                }
                v13 := GenS2{F0: 127}
                var v14 float64 = v10
                _ = v13
                _ = v14
            } else {
                switch "\n\t" {
                case "a":
                    v4 = (v4 & -203)
                case "\n\t":
                    var v15 uint = 1
                    _ = v15
                case "":
                    v2 = (v2 ^ (^(v2)))
                default:
                    if true {
                        if (!(v8 >= (v8 ^ v8))) {
                            for fi16 := 0; fi16 < 19; fi16++ {
                                fmt.Println(fi16, v4, v5)
                                v17 := GenS1{F0: 577}
                                _ = fi16
                                _ = v17
                            }
                        }
                    } else {
                        v11 -= -862
                    }
                }
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
            }
            v5 -= (^(v5))
            _ = v11
        } else {
            fmt.Println(v8)
        }
        if (144 < 4294967295) {
            if (uint16(v5) >= 32767) {
                for fi18 := 0; fi18 < 9; fi18++ {
                    if ((v10 == float64(-2.25)) || true) {
                        for fi19 := 0; fi19 < 19; fi19++ {
                            var v20 int32 = 2147483647
                            _ = fi19
                            _ = v20
                        }
                        if true {
                            if true {
                                var v21 string = "foo bar"
                                if ((((true && ((false || false) || true)) || false) || (true && true)) && ((100 == v8) && true)) {
                                    var v22 uint64 = 2
                                    v23 := GenS1{F0: 735}
                                    _ = v22
                                    _ = v23
                                } else {
                                    for fi24 := 0; fi24 < 15; fi24++ {
                                        var v25 float64 = float64(-2.25)
                                        switch ((v21 + "hello") + (v21 + v21)) {
                                        case "foo bar":
                                            var v26 int8 = v8
                                            _ = v26
                                        case "":
                                            switch (-(fi18)) {
                                            case -145:
                                                fmt.Println(v2)
                                            case -130:
                                                var v27 uint16 = 65535
                                                _ = v27
                                            }
                                        default:
                                            v28 := GenS2{F0: fi24}
                                            _ = v28
                                        }
                                        _ = fi24
                                        _ = v25
                                    }
                                    fmt.Println(v8, v10)
                                    var v29 uint = 0
                                    _ = v29
                                }
                                var v30 uint8 = v5
                                _ = v21
                                _ = v30
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = "foo bar"
                                }
                                v31 := GenS1{F0: 196}
                                for fi32 := 0; fi32 < 1; fi32++ {
                                    v33 := GenS1{F0: 42}
                                    fmt.Println(v2, fi18)
                                    _ = fi32
                                    _ = v33
                                }
                                _ = v31
                            }
                            var v34 uint16 = 202
                            _ = v34
                        }
                    }
                    _ = fi18
                }
                var v35 uint32 = 1
                v10 = (0.1 - (v10 * 1e10))
                _ = v35
            }
            var v36 float32 = -2.25
            var v37 uint16 = 0
            _ = v36
            _ = v37
        } else {
            for fi38 := 0; fi38 < 12; fi38++ {
                fmt.Println(v2)
                v8 -= v8
                _ = fi38
            }
            if (((v10 - 3.14159) != (float64(1e10) * ((((0.0 - v10) + 3.14159) / 1e10) + (v10 * v10)))) || false) {
                fmt.Println(v8, v4, v5)
                v1 = v1
                fmt.Println(v1, v5, v5)
            } else {
                if len(v3) > 0 {
                    v3[0] = (((("foo bar" + ("foo bar" + ("foo bar" + "hello"))) + "foo bar") + ("hello" + "\n\t")) + "")
                }
                v8 -= 42
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    }
    v39 := GenS3{F0: 42}
    var v40 bool = (0 < 116)
    var v41 uint16 = 2
    for fi42 := 0; fi42 < 13; fi42++ {
        if false {
            for fi43 := 0; fi43 < 2; fi43++ {
                if (!(v4 >= (fi42 + fi42))) {
                    var v44 uint32 = 724
                    _ = v44
                } else {
                    v3 = append(v3, "\n\t")
                    v45 := GenS4{F0: "foo bar", F1: 820, F2: 65535}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v45
                }
                _ = fi43
            }
            if (117 <= -444) {
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                v40 = ("foo bar" <= "hello")
                var v46 int32 = -1
                _ = v46
            } else {
                for fi47 := 0; fi47 < 6; fi47++ {
                    var v48 int16 = 32767
                    v1 = v48
                    _ = fi47
                    _ = v48
                }
                v2 = v2
                var v49 float32 = 0.0
                _ = v49
            }
        } else {
            for fi50 := 0; fi50 < 14; fi50++ {
                if ("hello" < "") {
                    var v51 int32 = int32(v5)
                    switch fi42 {
                    case 2147483647:
                        if v40 {
                            v52 := GenS3{F0: 0}
                            fmt.Println(v41, v8)
                            _ = v52
                        }
                    case 962:
                        v8 -= (v8 * v8)
                    default:
                        var v53 uint8 = (^(v5))
                        _ = v53
                    }
                    _ = v51
                }
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if len(v3) > 0 {
                        v3[0] = ("a" + "a")
                    }
                    v54 := GenS5{F0: (v5 + v5), F1: (fi42 & -546), F2: 2147483647}
                    _ = v54
                }
                _ = fi50
            }
            fmt.Println(v1, v10)
            v3 = append(v3, "a")
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
