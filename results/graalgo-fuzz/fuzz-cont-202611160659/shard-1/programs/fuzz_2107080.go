package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int8 }

func main() {
    var v1 bool = (65535 > 42)
    var v2 string = "foo bar"
    var v3 int32 = 308
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    var v6 int16 = -1
    v7 := []int{}
    if v1 {
        v6 = 32767
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int16 = v6
            v7 = append(v7, (fi8 | (^(fi8))))
            for fi10 := 0; fi10 < 2; fi10++ {
                v9 = v9
                v2 = "hello"
                fmt.Println(v2, v2)
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
    } else {
        var v11 int16 = 1
        _ = v11
    }
    if (-13 <= 100) {
        switch v2 {
        case "foo bar":
            switch "\n\t" {
            case "":
                v2 = (v2 + v2)
            case "a":
                for fi12 := 0; fi12 < 7; fi12++ {
                    if (!((!(int64(v6) <= -1)) || true)) {
                        v5[(fi12 / ((fi12) | 1))] = (1.5 < (0.1 - 0.0))
                        v13 := GenS1{F0: v2}
                        for fi14 := 0; fi14 < 2; fi14++ {
                            fmt.Println(v3)
                            v15 := GenS1{F0: v2}
                            _ = fi14
                            _ = v15
                        }
                        _ = v13
                    } else {
                        var v16 uint16 = 0
                        v16 -= (v16 * (v16 / ((v16) | 1)))
                        v16 = v16
                        _ = v16
                    }
                    _ = fi12
                }
            }
        case "\n\t":
            for fi17 := 0; fi17 < 9; fi17++ {
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 bool = (v2 != v2)
                    v6 -= (v6 - (v6 + 42))
                    _ = fi18
                    _ = v19
                }
                var v20 int32 = (-(v3))
                _ = fi17
                _ = v20
            }
        case "hello":
            var v21 uint32 = 2147483648
            _ = v21
        }
    }
    if false {
        var v22 int32 = 2147483647
        _ = v22
    } else {
        var v23 int16 = v6
        if (!(100 != 100)) {
            v1 = (!v1)
            switch 32767 {
            case 42:
                var v24 uint = 42
                _ = v24
            case 104:
                for fi25 := 0; fi25 < 20; fi25++ {
                    switch (^(fi25)) {
                    case 32767:
                        v26 := GenS2{F0: 2, F1: 1}
                        _ = v26
                    case -2147483648:
                        if v1 {
                            var v27 float64 = (0.1 / 1.5)
                            v2 = v2
                            _ = v27
                        } else {
                            v3 = (v3 % (((v3 ^ (v3 * v3))) | 1))
                        }
                    }
                    v23 = v6
                    var v28 string = "a"
                    _ = fi25
                    _ = v28
                }
            }
        } else {
            if v1 {
                switch "hello" {
                case "\n\t":
                    v29 := GenS1{F0: v2}
                    _ = v29
                case "a":
                    for fi30 := 0; fi30 < 6; fi30++ {
                        var v31 uint16 = 599
                        _ = fi30
                        _ = v31
                    }
                }
                if (v3 <= 2147483647) {
                    v32 := GenS1{F0: ""}
                    var v33 uint16 = 159
                    for fi34 := 0; fi34 < 17; fi34++ {
                        var v35 bool = (!(1e10 <= 1.5))
                        _ = fi34
                        _ = v35
                    }
                    _ = v32
                    _ = v33
                } else {
                    for fi36 := 0; fi36 < 7; fi36++ {
                        if len(v7) > 0 {
                            v7[0] = (-(fi36))
                        }
                        v3 -= (v3 + -1)
                        if (0 <= 522) {
                            var v37 uint32 = 2147483647
                            var v38 int32 = (v3 % ((v3) | 1))
                            var v39 int = 9223372036854775807
                            _ = v37
                            _ = v38
                            _ = v39
                        } else {
                            var v40 float64 = 0.0
                            _ = v40
                        }
                        _ = fi36
                    }
                    fmt.Println(v23, v2)
                }
                v2 = v2
            } else {
                if (-447 != 1) {
                    for fi41 := 0; fi41 < 3; fi41++ {
                        if v1 {
                            v1 = v1
                            v5[fi41] = ((v3 | (v3 - (v3 << ((77) & 63)))) == v3)
                        }
                        var v42 int = (fi41 ^ fi41)
                        for fi43 := 0; fi43 < 9; fi43++ {
                            var v44 int16 = (v23 >> ((474) & 63))
                            _ = fi43
                            _ = v44
                        }
                        _ = fi41
                        _ = v42
                    }
                    v45 := GenS3{F0: 65535}
                    v1 = v1
                    _ = v45
                } else {
                    for fi46 := 0; fi46 < 6; fi46++ {
                        v3 += v3
                        delete(v5, (fi46 + (fi46 << ((1) & 63))))
                        for fi47 := 0; fi47 < 11; fi47++ {
                            if v1 {
                                v6 += 1000
                            } else {
                                _ = len(v4)
                            }
                            _ = fi47
                        }
                        _ = fi46
                    }
                    v6 = (v6 * 1)
                }
                v3 = int32(v23)
            }
        }
        for fi48 := 0; fi48 < 9; fi48++ {
            v2 = v2
            v1 = v1
            _ = fi48
        }
        _ = v23
    }
    _ = len(v4)
    v49 := GenS4{F0: v2, F1: -1}
    var v50 float32 = 3.14159
    v51 := GenS3{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
}
