package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint; F2 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 int16 = 34
    fmt.Println(v3)
    var v4 int32 = 2147483647
    switch -630 {
    case 127:
        fmt.Println(v3, v4)
    case 0:
        var v5 uint8 = 255
        _ = v5
    case -448:
        v3 = v3
    }
    if ((false || ((2147483647 != 1) || (!(((v4 & v4) > (v4 << ((169) & 63))) || true)))) && true) {
        var v6 int64 = 624
        var v7 uint16 = 65535
        _ = v6
        _ = v7
    } else {
        v3 = (v3 + (v3 * -530))
        var v8 int = 0
        _ = v8
    }
    v4 += (v4 >> ((1) & 63))
    if (!false) {
        var v9 uint = 1
        _ = v9
    } else {
        if false {
            switch -714 {
            case 9223372036854775807:
                switch -9223372036854775808 {
                case 1:
                    switch 32767 {
                    case 32768:
                        fmt.Println(v4)
                    default:
                        v3 += -32768
                    }
                case 32767:
                    var v10 uint16 = 1
                    _ = v10
                default:
                    for fi11 := 0; fi11 < 9; fi11++ {
                        v4 = v4
                        v4 += 1
                        var v12 uint = 2
                        _ = fi11
                        _ = v12
                    }
                }
            }
            v3 -= v3
        } else {
            var v13 int8 = 127
            _ = v13
        }
    }
    v14 := []int{32768}
    v4 -= v4
    switch (("hello" + "hello") + ("a" + "a")) {
    case "hello":
        if (32767 > 940) {
            var v15 int64 = 963
            _ = v15
        }
    case "foo bar":
        switch 2147483647 {
        case 1:
            fmt.Println(v4, v4)
        default:
            v4 += (v4 | 0)
        }
    case "":
        var v16 uint16 = 782
        _ = v16
    }
    var v17 float64 = ((3.14159 / 1.5) / 1.5)
    switch -2147483648 {
    case -1:
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := GenS1{F0: -12, F1: "a", F2: (v3 & (-(v3)))}
            fmt.Println(fi18, v4, v17)
            v20 := GenS2{F0: "foo bar", F1: false}
            _ = fi18
            _ = v19
            _ = v20
        }
    case 829:
        v17 = v17
    case 1:
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 uint64 = 42
            var v23 uint8 = 0
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v24 := GenS3{F0: (v3 * (v3 | v3)), F1: 128, F2: 986}
    fmt.Println(v17)
    v3 = 32767
    for fi25 := 0; fi25 < 11; fi25++ {
        if false {
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 uint8 = 90
                if len(v14) > 0 {
                    v14[0] = (fi25 | -9223372036854775808)
                }
                var v28 uint16 = 65535
                _ = fi26
                _ = v27
                _ = v28
            }
            v29 := GenS1{F0: 6, F1: "hello", F2: 1}
            _ = v29
        }
        if (true || true) {
            fmt.Println(v4, v3)
            fmt.Println(v3, v17)
            var v30 int32 = v4
            _ = v30
        }
        switch fi25 {
        case 127:
            if false {
                var v31 uint = 934
                _ = v31
            } else {
                if len(v14) > 0 {
                    v14[0] = (fi25 * fi25)
                }
                for fi32 := 0; fi32 < 7; fi32++ {
                    var v33 uint8 = 42
                    v33 += v33
                    v33 -= (v33 >> ((271) & 63))
                    _ = fi32
                    _ = v33
                }
                delete(v1, "a")
            }
        default:
            var v34 uint16 = 990
            _ = v34
        }
        _ = fi25
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        v36 := GenS4{F0: (-(v4)), F1: 896, F2: 551}
        if (!false) {
            v4 += (v4 & v4)
            var v37 int32 = (^(v4))
            _ = v37
        } else {
            var v38 int = 4294967295
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    switch ("a" + "hello") {
    case "\n\t":
        v4 += v4
    case "":
        for fi39 := 0; fi39 < 5; fi39++ {
            if false {
                switch 128 {
                case 73:
                    v1[""] = fi39
                default:
                    var v40 float32 = float32((float64(3.14159) + (v17 / v17)))
                    _ = v40
                }
                switch fi39 {
                case 4294967295:
                    v17 = (-2.25 * (float64(0.1) - v17))
                case 2147483647:
                    var v41 float64 = 1.5
                    _ = v41
                default:
                    if true {
                        v14 = append(v14, fi39)
                    } else {
                        for fi42 := 0; fi42 < 14; fi42++ {
                            var v43 int8 = 127
                            if (("hello" <= "hello") && true) {
                                fmt.Println(v17, v43, fi42)
                                var v44 uint64 = 42
                                v45 := GenS4{F0: v4, F1: 18446744073709551615, F2: v44}
                                _ = v44
                                _ = v45
                            } else {
                                var v46 uint16 = 717
                                _ = len(v2)
                                _ = v46
                            }
                            if false {
                                var v47 uint32 = 204
                                switch fi42 {
                                case 1:
                                    var v48 int16 = v3
                                    _ = v48
                                case -797:
                                    v49 := GenS3{F0: 32767, F1: 142, F2: 42}
                                    _ = v49
                                case -648:
                                    v4 += v4
                                    v43 = (v43 % ((-94) | 1))
                                }
                                _ = v47
                            }
                            _ = fi42
                            _ = v43
                        }
                    }
                }
            } else {
                v14 = append(v14, (fi39 ^ (fi39 - fi39)))
                if true {
                    var v50 string = ("hello" + "foo bar")
                    _ = v50
                }
            }
            fmt.Println(v3, v3)
            fmt.Println(fi39)
            _ = fi39
        }
    case "a":
        var v51 uint64 = 356
        _ = v51
    default:
        fmt.Println(v17, v4, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
