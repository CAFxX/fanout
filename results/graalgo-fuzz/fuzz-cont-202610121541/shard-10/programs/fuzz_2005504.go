package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 string; F2 int64 }
func genFn0(p0 int64, p1 uint16) string {
    v5 := GenS1{F0: 0}
    v6 := []int{255, 128, 0, 0}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    v1 := []uint32{42, 491}
    var v2 int64 = -349
    var v3 string = ""
    var v4 bool = false
    if v4 {
        for fi7 := 0; fi7 < 20; fi7++ {
            if v4 {
                var v8 int64 = v2
                _ = v8
            }
            v4 = (1 == 18446744073709551615)
            var v9 int = (fi7 / ((-9223372036854775808) | 1))
            _ = fi7
            _ = v9
        }
        switch -9223372036854775808 {
        case -575:
            for fi10 := 0; fi10 < 10; fi10++ {
                for fi11 := 0; fi11 < 17; fi11++ {
                    v4 = v4
                    v3 = v3
                    v12 := GenS1{F0: 128}
                    _ = fi11
                    _ = v12
                }
                var v13 uint32 = 2147483647
                v13 = (^(v13))
                _ = fi10
                _ = v13
            }
        case 629:
            switch 2147483647 {
            case 32767:
                var v14 int32 = 136
                _ = v14
            case 2147483647:
                if v4 {
                    v2 -= (v2 - -1)
                }
            case 128:
                var v15 uint32 = 65535
                _ = v15
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    }
    v16 := GenS2{F0: 1, F1: (v3 + (v3 + ("" + (v3 + (("foo bar" + ("a" + v3)) + ("" + v3)))))), F2: (^(v2))}
    v1 = append(v1, 881)
    if true {
        var v17 int = -1
        v18 := GenS1{F0: v17}
        _ = v17
        _ = v18
    } else {
        v4 = v4
        v4 = (2 <= 65535)
    }
    v19 := GenS2{F0: 0, F1: "foo bar", F2: (v2 | (v2 & v2))}
    switch 2147483647 {
    case 127:
        var v20 int = 32768
        _ = v20
    case 1:
        if false {
            var v21 int = 127
            if (!(v4 && (42 != 987))) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch v21 {
                case 635:
                    fmt.Println(v4, v21)
                }
                var v22 uint64 = 681
                _ = v22
            } else {
                var v23 float64 = 0.0
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v23 -= (1e10 - float64(float32(1.5)))
                _ = v23
            }
            v3 = (v3 + v3)
            _ = v21
        } else {
            switch 2147483647 {
            case -674:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 1:
                v1 = append(v1, 42)
            }
            v4 = (127 < 0)
            if (!(("foo bar" + "hello") >= v3)) {
                for fi24 := 0; fi24 < 6; fi24++ {
                    for fi25 := 0; fi25 < 4; fi25++ {
                        var v26 string = "\n\t"
                        fmt.Println(v4, fi24)
                        _ = fi25
                        _ = v26
                    }
                    switch (^(fi24)) {
                    case 127:
                        _ = len(v1)
                    case 551:
                        for fi27 := 0; fi27 < 9; fi27++ {
                            fmt.Println(v4)
                            _ = fi27
                        }
                    default:
                        v28 := GenS2{F0: 42, F1: "a", F2: 0}
                        _ = v28
                    }
                    _ = fi24
                }
                var v29 int64 = v2
                v30 := GenS2{F0: 1, F1: v3, F2: (v29 & (v2 % (((v2 >> ((1) & 63))) | 1)))}
                _ = v29
                _ = v30
            } else {
                for fi31 := 0; fi31 < 6; fi31++ {
                    if true {
                        v32 := GenS2{F0: 0, F1: ("hello" + "hello"), F2: v2}
                        for fi33 := 0; fi33 < 19; fi33++ {
                            if (-644 == int32(fi31)) {
                                v2 = 809
                                var v34 uint32 = 0
                                v35 := GenS1{F0: (fi33 | fi31)}
                                var v36 int16 = 1
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            _ = fi33
                        }
                        _ = v32
                    }
                    _ = fi31
                }
            }
        }
    }
    var v37 uint64 = 553
    var v38 bool = (!(v2 <= (v2 << ((534) & 63))))
    for fi39 := 0; fi39 < 13; fi39++ {
        for fi40 := 0; fi40 < 16; fi40++ {
            v2 = 9223372036854775807
            if (127 >= 128) {
                v4 = true
                var v41 int16 = -297
                _ = v41
            }
            _ = fi40
        }
        var v42 uint64 = 2
        _ = fi39
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
