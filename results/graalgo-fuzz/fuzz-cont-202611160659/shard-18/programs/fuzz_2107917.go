package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int64; F1 string; F2 uint32 }

func main() {
    var v1 float64 = (0.1 / 1.0)
    v2 := make(map[uint32]string)
    var v3 float32 = 0.0
    v4 := []string{("hello" + ((("" + "foo bar") + "\n\t") + "hello"))}
    var v5 int16 = 842
    var v6 uint64 = 2
    v6 += v6
    v5 -= (-(v5))
    v3 = v3
    v7 := GenS1{F0: ("hello" + ("hello" + "a"))}
    fmt.Println(v3)
    v5 = v5
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: "\n\t"}
        v5 += (-(v5))
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi10
    }
    if (-280 > 0) {
        var v11 int32 = -2147483648
        if false {
            for fi12 := 0; fi12 < 7; fi12++ {
                v13 := GenS2{F0: 321, F1: 65535, F2: (v6 / ((v6) | 1))}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 6; fi14++ {
                for fi15 := 0; fi15 < 5; fi15++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch (fi15 << ((42) & 63)) {
                    case -768:
                        var v16 float32 = float32(float64(v3))
                        _ = v16
                    case 0:
                        _ = len(v4)
                    case 127:
                        var v17 int32 = 290
                        _ = v17
                    }
                    _ = fi15
                }
                if (2 < 4294967295) {
                    fmt.Println(v11, v11)
                    if true {
                        for fi18 := 0; fi18 < 9; fi18++ {
                            switch (fi18 << ((65535) & 63)) {
                            case 42:
                                switch "" {
                                case "foo bar":
                                    v19 := GenS3{F0: 912}
                                    _ = v19
                                }
                            case 1:
                                v6 = (v6 | 2)
                            default:
                                v5 = int16(v11)
                            }
                            var v20 int8 = 42
                            v4 = append(v4, ("foo bar" + ""))
                            _ = fi18
                            _ = v20
                        }
                        v4 = append(v4, "")
                        v3 = v3
                    } else {
                        v21 := GenS1{F0: "foo bar"}
                        var v22 uint8 = 2
                        _ = v21
                        _ = v22
                    }
                }
                var v23 int32 = v11
                _ = fi14
                _ = v23
            }
        } else {
            v24 := GenS4{F0: (v11 * (v11 - -1)), F1: (v6 ^ (v6 * (v6 * (^(v6)))))}
            _ = v24
        }
        var v25 float32 = v3
        _ = v11
        _ = v25
    } else {
        v1 = v1
    }
    if (("\n\t" <= "hello") && false) {
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 uint64 = (v6 << ((65535) & 63))
            var v28 uint64 = (v27 - (^(v6)))
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 uint8 = 42
        _ = v29
    } else {
        v5 += (-(v5))
        var v30 uint64 = 4294967295
        _ = v30
    }
    if (32767 <= -1) {
        v31 := GenS1{F0: "foo bar"}
        _ = len(v4)
        v32 := GenS5{F0: 487, F1: "a", F2: 2147483647}
        _ = v31
        _ = v32
    }
    v6 -= 940
    v4 = append(v4, ("hello" + ((("a" + "") + ((("foo bar" + ("hello" + ("foo bar" + "foo bar"))) + "foo bar") + "hello")) + (("foo bar" + "foo bar") + (("hello" + "") + ("a" + "hello"))))))
    for fi33 := 0; fi33 < 9; fi33++ {
        switch 32767 {
        case 4294967295:
            v1 += v1
        case 863:
            v34 := GenS5{F0: -1, F1: "foo bar", F2: 2147483648}
            _ = v34
        case -142:
            v35 := GenS1{F0: "a"}
            _ = v35
        default:
            fmt.Println(v3, v3, v3)
        }
        v5 -= v5
        var v36 uint32 = 5
        _ = fi33
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
