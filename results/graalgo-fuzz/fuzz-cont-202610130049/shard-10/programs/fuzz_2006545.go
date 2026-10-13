package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int; F2 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 int }
func genFn0() uint {
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint = 65535
            v8 = v8
            switch "foo bar" {
            case "a":
                v9 := GenS1{F0: 994, F1: 0, F2: 0}
                _ = v9
            case "hello":
                var v10 uint32 = 2
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v11 uint8 = 85
    v11 = (v11 + 108)
    v12 := GenS2{F0: 65535, F1: "\n\t"}
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int8 = 42
        v15 := GenS3{F0: 45}
        if (!(0 <= 42)) {
            v16 := GenS3{F0: 2}
            _ = v16
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v11
    _ = v12
    return 541
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    v3 := []string{"\n\t", "hello", "hello"}
    var v4 int8 = 101
    v5 := make(map[string]uint32)
    v4 -= (-(v4))
    v17 := GenS4{F0: -32768}
    v18 := GenS4{F0: 32767}
    v4 -= v4
    var v19 int64 = 1
    var v20 uint32 = 0
    var v21 int8 = (v4 | (v4 ^ v4))
    switch 305 {
    case 32768:
        switch ("hello" + "foo bar") {
        case "foo bar":
            if false {
                fmt.Println(v20)
            }
        default:
            if (0 >= 525) {
                for fi22 := 0; fi22 < 15; fi22++ {
                    v23 := GenS2{F0: 255, F1: "a"}
                    for fi24 := 0; fi24 < 9; fi24++ {
                        v21 += -1
                        switch "" {
                        case "\n\t":
                            var v25 float32 = 0.0
                            _ = v25
                        case "":
                            var v26 int32 = -40
                            _ = v26
                        default:
                            var v27 uint = 1
                            _ = v27
                        }
                        _ = fi24
                    }
                    v28 := GenS5{F0: 0, F1: fi22, F2: v19}
                    _ = fi22
                    _ = v23
                    _ = v28
                }
            } else {
                switch int(v4) {
                case 255:
                    v19 += v19
                case -52:
                    var v29 float64 = 1e10
                    _ = v29
                case 32767:
                    v1[v20] = false
                }
                var v30 int64 = v19
                _ = v30
            }
        }
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        switch "" {
        case "hello":
            for fi32 := 0; fi32 < 14; fi32++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (32767 == 1000) {
                    if true {
                        v33 := GenS5{F0: 317, F1: (fi32 ^ (fi32 >> ((4294967295) & 63))), F2: v19}
                        v21 += (v4 + v4)
                        if (true && (!false)) {
                            _ = len(v1)
                            if false {
                                _ = len(v2)
                            }
                        }
                        _ = v33
                    } else {
                        _ = len(v3)
                    }
                }
                _ = fi32
            }
        default:
            v21 = v4
        }
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 float64 = 1.5
            var v36 int64 = 1
            _ = fi34
            _ = v35
            _ = v36
        }
        for fi37 := 0; fi37 < 20; fi37++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v38 := GenS4{F0: -67}
            _ = fi37
            _ = v38
        }
        _ = fi31
    }
    fmt.Println(v21)
    v39 := v1[v20]
    for fi40 := 0; fi40 < 1; fi40++ {
        for fi41 := 0; fi41 < 4; fi41++ {
            _ = len(v1)
            v4 += (v4 << ((4294967295) & 63))
            v20 -= 42
            _ = fi41
        }
        v42 := GenS6{F0: fi40}
        fmt.Println(v20, v19, v39)
        _ = fi40
        _ = v42
    }
    v43 := GenS3{F0: 2}
    v44 := GenS3{F0: 2}
    switch "foo bar" {
    case "foo bar":
        v45 := GenS7{F0: 18446744073709551615}
        _ = v45
    default:
        for fi46 := 0; fi46 < 2; fi46++ {
            switch "foo bar" {
            case "hello":
                _ = len(v2)
            default:
                v19 = (v19 ^ v19)
            }
            v47 := GenS8{F0: 842, F1: fi46}
            _ = fi46
            _ = v47
        }
    }
    v2["a"] = -2147483648
    v48 := v2[("\n\t" + ("a" + (("foo bar" + "\n\t") + (("a" + ("\n\t" + ("\n\t" + "foo bar"))) + "foo bar"))))]
    v21 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v48
    _ = v5
}
