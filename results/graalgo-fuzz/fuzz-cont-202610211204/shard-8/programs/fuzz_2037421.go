package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint; F2 int16 }
type GenS4 struct { F0 uint; F1 bool; F2 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64 }
func genFn0() string {
    if false {
        var v4 int32 = 42
        for fi5 := 0; fi5 < 2; fi5++ {
            v4 = v4
            v6 := GenS1{F0: 1, F1: int16(fi5)}
            var v7 int64 = -9223372036854775808
            _ = fi5
            _ = v6
            _ = v7
        }
        v4 -= v4
        _ = v4
    }
    if (452 > 697) {
        v8 := GenS2{F0: 505, F1: 255}
        _ = v8
    } else {
        var v9 uint = 1
        _ = v9
    }
    return "hello"
}

func main() {
    v1 := []string{"\n\t", "hello", "a", "foo bar"}
    v2 := make(map[string]int)
    var v3 uint = 4294967295
    v10 := GenS1{F0: -1, F1: -583}
    var v11 float32 = 1e10
    v12 := GenS3{F0: 100, F1: v3, F2: 32767}
    var v13 uint64 = 680
    v1 = append(v1, ("\n\t" + ""))
    var v14 int = 128
    switch ((("hello" + "") + "hello") + "") {
    case "foo bar":
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 uint = 18446744073709551615
            v17 := GenS4{F0: (v16 & 2), F1: ((v13 + (v13 & v13)) != (v13 + (v13 ^ (v13 | v13)))), F2: "foo bar"}
            _ = fi15
            _ = v16
            _ = v17
        }
    case "hello":
        switch (v14 ^ 667) {
        case -729:
            for fi18 := 0; fi18 < 7; fi18++ {
                if ((float32(float64(v11)) == (-2.25 * ((v11 / v11) - v11))) || ((false && (286 > fi18)) && (9223372036854775807 < -413))) {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        var v20 uint8 = 127
                        _ = fi19
                        _ = v20
                    }
                    if true {
                        var v21 int16 = 0
                        v22 := GenS5{F0: 42}
                        _ = v21
                        _ = v22
                    } else {
                        var v23 uint8 = 1
                        _ = v23
                    }
                }
                if (v3 < (^(v3))) {
                    delete(v2, "\n\t")
                    v13 = (v13 ^ v13)
                } else {
                    for fi24 := 0; fi24 < 3; fi24++ {
                        if (253 >= 160) {
                            var v25 uint = (^(v3))
                            var v26 int64 = 42
                            _ = v25
                            _ = v26
                        } else {
                            v27 := GenS4{F0: v3, F1: true, F2: ("foo bar" + ("hello" + ("foo bar" + (("hello" + "hello") + ((("" + "foo bar") + "\n\t") + "foo bar")))))}
                            _ = v27
                        }
                        var v28 uint64 = (v13 - 18446744073709551615)
                        v2["a"] = fi24
                        _ = fi24
                        _ = v28
                    }
                }
                var v29 int16 = 32767
                _ = fi18
                _ = v29
            }
        case 32767:
            for fi30 := 0; fi30 < 16; fi30++ {
                if false {
                    v13 += (v13 / ((v13) | 1))
                }
                v14 = fi30
                _ = fi30
            }
        default:
            fmt.Println(v3)
        }
    case "a":
        if ((1 < int64(v14)) && ((1.5 + (v11 - v11)) > (v11 - v11))) {
            if true {
                var v31 uint = 2
                var v32 int64 = 9223372036854775807
                v3 -= 0
                _ = v31
                _ = v32
            }
            v13 += (v13 % ((v13) | 1))
            v33 := GenS6{F0: 255}
            _ = v33
        } else {
            v3 -= (v3 >> ((v3) & 63))
            v34 := GenS3{F0: 100, F1: v3, F2: 42}
            if false {
                fmt.Println(v11)
                var v35 float32 = v11
                var v36 uint32 = 65535
                _ = v35
                _ = v36
            } else {
                if false {
                    if true {
                        if ((2 > 42) || (0 >= -32768)) {
                            fmt.Println(v13, v13)
                            for fi37 := 0; fi37 < 13; fi37++ {
                                v11 = float32(3.14159)
                                if (!(v3 == 2)) {
                                    var v38 string = (((("a" + "foo bar") + "") + "") + "")
                                    v3 = 579
                                    _ = v38
                                } else {
                                    v14 = int(v3)
                                }
                                v11 = v11
                                _ = fi37
                            }
                        } else {
                            v14 += (v14 % ((v14) | 1))
                            var v39 uint64 = (^(v13))
                            _ = v39
                        }
                    }
                    var v40 uint64 = v13
                    _ = v40
                }
                switch 545 {
                case 2147483647:
                    v1 = append(v1, ("hello" + ""))
                default:
                    var v41 int16 = 1000
                    _ = v41
                }
                var v42 float64 = 0.1
                _ = v42
            }
            _ = v34
        }
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        var v44 uint8 = 128
        v45 := GenS7{F0: v13}
        v11 = float32(float64(v11))
        _ = fi43
        _ = v44
        _ = v45
    }
    v46 := v2["hello"]
    if false {
        if false {
            v46 = (^(v14))
            v14 = 127
        }
        if (false || (!((^(v3)) <= (v3 % ((v3) | 1))))) {
            for fi47 := 0; fi47 < 3; fi47++ {
                var v48 float32 = float32(-2.25)
                _ = len(v2)
                v13 = v13
                _ = fi47
                _ = v48
            }
        }
    }
    if true {
        if (-80 > 1) {
            switch v14 {
            case 0:
                fmt.Println(v3)
            case 1:
                for fi49 := 0; fi49 < 15; fi49++ {
                    for fi50 := 0; fi50 < 8; fi50++ {
                        v13 = v13
                        _ = fi50
                    }
                    _ = fi49
                }
            default:
                var v51 uint64 = (v13 >> ((255) & 63))
                _ = v51
            }
        }
        var v52 uint64 = v13
        _ = v52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v46
}
