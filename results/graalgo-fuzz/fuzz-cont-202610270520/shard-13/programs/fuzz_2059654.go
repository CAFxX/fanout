package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint16 }

func main() {
    var v1 int16 = 801
    var v2 int16 = 42
    var v3 uint64 = 4294967295
    v4 := []int{}
    var v5 uint64 = v3
    var v6 float64 = 1e10
    for fi7 := 0; fi7 < 14; fi7++ {
        if true {
            if len(v4) > 0 {
                v4[0] = (fi7 * (fi7 / (((fi7 + 32767)) | 1)))
            }
            switch fi7 {
            case 32768:
                for fi8 := 0; fi8 < 5; fi8++ {
                    var v9 float32 = 0.1
                    v3 -= v3
                    _ = fi8
                    _ = v9
                }
            case -9223372036854775808:
                v10 := GenS1{F0: 1, F1: "hello", F2: 255}
                _ = v10
            }
        } else {
            var v11 uint = 0
            v11 = 2
            _ = v11
        }
        if (32767 == v1) {
            var v12 uint64 = v5
            _ = v12
        }
        _ = fi7
    }
    var v13 uint = 2
    if len(v4) > 0 {
        v4[0] = int(v13)
    }
    v14 := GenS1{F0: -2147483648, F1: "hello", F2: 255}
    if true {
        switch ("\n\t" + ("a" + "foo bar")) {
        case "hello":
            v3 = v5
        default:
            if (!true) {
                var v15 bool = true
                var v16 uint32 = 42
                _ = v15
                _ = v16
            }
        }
        if false {
            v3 = (^(v5))
            if false {
                fmt.Println(v6, v13, v5)
                for fi17 := 0; fi17 < 2; fi17++ {
                    v18 := GenS1{F0: 0, F1: "\n\t", F2: 0}
                    for fi19 := 0; fi19 < 2; fi19++ {
                        v2 -= v2
                        _ = fi19
                    }
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v1 += int16(v3)
                        _ = fi20
                    }
                    _ = fi17
                    _ = v18
                }
            }
        } else {
            for fi21 := 0; fi21 < 11; fi21++ {
                fmt.Println(v1)
                var v22 int32 = 0
                v23 := GenS1{F0: v22, F1: "hello", F2: 65535}
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 uint32 = 1
            switch (("hello" + "foo bar") + "foo bar") {
            case "\n\t":
                fmt.Println(v2, v6)
            }
            _ = v24
        }
    } else {
        var v25 uint16 = 954
        v26 := GenS1{F0: -787, F1: ((("\n\t" + "\n\t") + "") + "hello"), F2: 0}
        _ = v25
        _ = v26
    }
    v2 += -575
    v5 = (^(v3))
    var v27 int16 = -1
    var v28 uint32 = 4294967295
    v6 = v6
    if false {
        if true {
            v29 := GenS1{F0: -491, F1: "\n\t", F2: 42}
            switch 128 {
            case -9223372036854775808:
                v28 += 360
            }
            _ = v29
        }
        v28 -= 586
    } else {
        fmt.Println(v3, v6, v6)
        var v30 uint64 = v5
        var v31 uint64 = v30
        _ = v30
        _ = v31
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v32 int32 = -1
    for fi33 := 0; fi33 < 13; fi33++ {
        v5 = (^(v3))
        v4 = append(v4, (fi33 ^ int(v5)))
        _ = fi33
    }
    if (1.5 <= ((0.1 + 1e10) - ((1.5 / float32(3.14159)) - 0.0))) {
        var v34 bool = false
        for fi35 := 0; fi35 < 4; fi35++ {
            if ((v13 & 42) != 65535) {
                var v36 string = "\n\t"
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if v34 {
                    var v37 int = (fi35 | (fi35 & fi35))
                    _ = v37
                } else {
                    v38 := GenS2{F0: v36}
                    for fi39 := 0; fi39 < 8; fi39++ {
                        v36 = v36
                        v3 += 1
                        fmt.Println(v27, v6)
                        _ = fi39
                    }
                    var v40 float64 = v6
                    _ = v38
                    _ = v40
                }
                _ = v36
            } else {
                for fi41 := 0; fi41 < 14; fi41++ {
                    switch 32768 {
                    case -1:
                        var v42 uint16 = 204
                        _ = v42
                    case 127:
                        v34 = v34
                    default:
                        v43 := GenS3{F0: 0, F1: 827}
                        _ = v43
                    }
                    v4 = append(v4, fi35)
                    _ = fi41
                }
                for fi44 := 0; fi44 < 2; fi44++ {
                    var v45 uint16 = 32767
                    var v46 float32 = 0.1
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                if true {
                    fmt.Println(v13, v28, v5)
                }
            }
            v6 += 0.1
            _ = fi35
        }
        _ = v34
    }
    if (false && ((^(v27)) < (v27 / ((v27) | 1)))) {
        var v47 uint64 = 933
        _ = v47
    }
    fmt.Println(v3)
    var v48 uint8 = 125
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
