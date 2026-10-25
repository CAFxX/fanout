package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }
type GenS2 struct { F0 uint32; F1 int16; F2 int16 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 int; F2 int32 }
type GenS9 struct { F0 int }

func main() {
    var v1 bool = true
    v2 := []string{}
    var v3 int16 = -32768
    var v4 bool = (32767 == (v3 * (^(v3))))
    if (32767 <= 272) {
        fmt.Println(v3, v3)
        v2 = append(v2, "a")
    } else {
        if v4 {
            for fi5 := 0; fi5 < 13; fi5++ {
                v6 := GenS1{F0: -9223372036854775808, F1: ("" + "a"), F2: -1}
                v4 = true
                _ = fi5
                _ = v6
            }
            for fi7 := 0; fi7 < 17; fi7++ {
                switch fi7 {
                case 32767:
                    _ = len(v2)
                case 2147483647:
                    var v8 int = fi7
                    _ = v8
                case 9223372036854775807:
                    var v9 bool = v1
                    _ = v9
                }
                v10 := GenS1{F0: -686, F1: ("\n\t" + ("" + ("\n\t" + ""))), F2: (v3 ^ -1)}
                switch (fi7 / (((fi7 - (fi7 % (((fi7 & (-(fi7)))) | 1)))) | 1)) {
                case 9223372036854775807:
                    var v11 float64 = 3.14159
                    _ = v11
                case 2147483647:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi12
                    }
                default:
                    var v13 int = fi7
                    _ = v13
                }
                _ = fi7
                _ = v10
            }
        }
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    }
    var v14 float64 = (0.1 - (0.1 - 0.1))
    v4 = ((1.5 + v14) > v14)
    v14 = (v14 / 1.0)
    fmt.Println(v4, v4, v4)
    if false {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 uint16 = 1
            switch "hello" {
            case "hello":
                if len(v2) > 0 {
                    v2[0] = "a"
                }
            default:
                var v17 float32 = ((3.14159 - (0.1 / 0.1)) - -2.25)
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        v2 = append(v2, "hello")
        if v4 {
            if (-71 <= 100) {
                if false {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v4 = v1
                        fmt.Println(v3, v1, v1)
                        _ = fi18
                    }
                    v3 += -32768
                    v19 := GenS1{F0: -828, F1: (("foo bar" + (("\n\t" + "hello") + "foo bar")) + ""), F2: v3}
                    _ = v19
                } else {
                    v20 := GenS1{F0: 128, F1: "\n\t", F2: (-(v3))}
                    _ = v20
                }
                if len(v2) > 0 {
                    v2[0] = ("" + "hello")
                }
            }
            var v21 uint = 255
            _ = v21
        } else {
            switch "" {
            case "foo bar":
                var v22 string = "hello"
                _ = v22
            }
            v23 := GenS2{F0: 4294967295, F1: v3, F2: 42}
            _ = v23
        }
    }
    if (42 != 42) {
        var v24 int32 = -464
        v24 += (v24 << ((678) & 63))
        v1 = v4
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        for fi26 := 0; fi26 < 18; fi26++ {
            var v27 uint64 = 456
            v3 = v3
            v3 -= (v3 | v3)
            _ = fi26
            _ = v27
        }
        switch "foo bar" {
        case "\n\t":
            v28 := GenS3{F0: 1, F1: "a"}
            _ = v28
        case "foo bar":
            for fi29 := 0; fi29 < 19; fi29++ {
                switch -9223372036854775808 {
                case 42:
                    v3 -= (v3 / ((226) | 1))
                default:
                    var v30 uint32 = 42
                    _ = v30
                }
                v31 := GenS4{F0: 968}
                _ = fi29
                _ = v31
            }
        case "a":
            if v1 {
                v1 = (115 != 255)
            }
        default:
            var v32 uint64 = 108
            _ = v32
        }
        _ = fi25
    }
    v3 = -542
    v4 = v1
    v3 += 42
    fmt.Println(v4)
    for fi33 := 0; fi33 < 5; fi33++ {
        v34 := GenS4{F0: 875}
        _ = len(v2)
        if (v4 || v4) {
            switch "foo bar" {
            case "hello":
                var v35 bool = (v4 || (((((float64(-2.25) + -2.25) == (v14 / float64(1e10))) && (684 >= 2)) || (true || (2 == 21))) || v1))
                _ = v35
            case "a":
                v3 += (v3 / ((v3) | 1))
            case "":
                var v36 int16 = (v3 << ((10) & 63))
                _ = v36
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v37 := GenS1{F0: fi33, F1: "a", F2: 32767}
            _ = v37
        } else {
            v38 := GenS1{F0: (fi33 % ((fi33) | 1)), F1: "hello", F2: (v3 / ((302) | 1))}
            v2 = append(v2, (((("\n\t" + ("foo bar" + "")) + "foo bar") + (("foo bar" + "\n\t") + ("a" + "hello"))) + "hello"))
            _ = v38
        }
        _ = fi33
        _ = v34
    }
    v39 := GenS5{F0: (!(!true))}
    if (!false) {
        v1 = (true && (-2147483648 < 2147483647))
        fmt.Println(v14)
    }
    v4 = (-2.25 != (v14 / v14))
    if true {
        if v4 {
            if v1 {
                if v4 {
                    v40 := GenS3{F0: 269, F1: "hello"}
                    for fi41 := 0; fi41 < 4; fi41++ {
                        _ = len(v2)
                        _ = fi41
                    }
                    v42 := GenS6{F0: true, F1: v1}
                    _ = v40
                    _ = v42
                }
            }
            fmt.Println(v3, v1, v4)
        } else {
            v14 -= 3.14159
            for fi43 := 0; fi43 < 13; fi43++ {
                switch "\n\t" {
                case "":
                    v44 := GenS7{F0: v4}
                    _ = v44
                default:
                    var v45 uint32 = 2
                    _ = v45
                }
                _ = fi43
            }
        }
        for fi46 := 0; fi46 < 20; fi46++ {
            v47 := GenS8{F0: (42 >= 242), F1: fi46, F2: 592}
            _ = fi46
            _ = v47
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch -2147483648 {
        case 1:
            v48 := GenS9{F0: -1}
            _ = v48
        case 4294967295:
            v14 = -2.25
        }
        for fi49 := 0; fi49 < 2; fi49++ {
            var v50 int32 = -2147483648
            _ = fi49
            _ = v50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
}
