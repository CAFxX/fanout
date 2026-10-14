package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint }
type GenS3 struct { F0 bool; F1 bool; F2 int16 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 uint32 = 65535
    var v2 float32 = 1e10
    var v3 int32 = 29
    switch "" {
    case "\n\t":
        v4 := []int{42, 4294967295, 524}
        _ = v4
    }
    switch "" {
    case "hello":
        v3 += (v3 & v3)
    case "a":
        var v5 int8 = 127
        _ = v5
    case "":
        v6 := GenS1{F0: (("\n\t" + "hello") + "foo bar"), F1: 0}
        _ = v6
    }
    if (!(0.0 < float64(v2))) {
        v7 := []int{}
        var v8 uint8 = 2
        _ = v7
        _ = v8
    } else {
        switch 634 {
        case 502:
            var v9 uint = 255
            _ = v9
        case 2147483647:
            fmt.Println(v3, v2, v2)
        case 32767:
            v2 -= 0.1
        }
        var v10 int = 2147483647
        _ = v10
    }
    switch 651 {
    case 4294967295:
        var v11 string = (("\n\t" + "\n\t") + ("\n\t" + ("\n\t" + "foo bar")))
        _ = v11
    case -2147483648:
        v3 = (v3 - v3)
    case -9223372036854775808:
        var v12 uint16 = 1
        _ = v12
    default:
        for fi13 := 0; fi13 < 8; fi13++ {
            v2 = 1e10
            v14 := []int{}
            _ = fi13
            _ = v14
        }
    }
    if (32767 < 2) {
        fmt.Println(v3, v2, v1)
    }
    v15 := GenS1{F0: "a", F1: 2}
    var v16 float32 = (3.14159 - v2)
    var v17 uint32 = v1
    v1 += v1
    if (!(!true)) {
        var v18 uint16 = 255
        _ = v18
    } else {
        v19 := []int{9223372036854775807, -773, -783}
        _ = v19
    }
    if ("a" == ("foo bar" + ("hello" + "foo bar"))) {
        switch (("\n\t" + (((("hello" + ("\n\t" + ("a" + "hello"))) + ("a" + "a")) + "") + "hello")) + ((("\n\t" + "") + "\n\t") + "a")) {
        case "foo bar":
            for fi20 := 0; fi20 < 14; fi20++ {
                if (4294967295 <= 1) {
                    var v21 uint32 = v17
                    if false {
                        var v22 int16 = -411
                        switch (fi20 | 128) {
                        case -655:
                            v23 := GenS1{F0: "foo bar", F1: 636}
                            _ = v23
                        case 251:
                            v2 = ((v16 / (v16 + 0.0)) + ((-2.25 - v2) + v2))
                        default:
                            v17 += 2
                        }
                        switch ((("\n\t" + ("foo bar" + "hello")) + ("foo bar" + ("hello" + "\n\t"))) + "foo bar") {
                        case "\n\t":
                            v2 = (v16 * (((-2.25 / 0.1) + (v16 / (v2 / v16))) + ((v16 * (1e10 + v16)) / ((3.14159 + 3.14159) * -2.25))))
                        case "foo bar":
                            switch "foo bar" {
                            case "\n\t":
                                v24 := GenS2{F0: v1, F1: 2, F2: 42}
                                _ = v24
                            }
                        case "a":
                            var v25 uint16 = 42
                            _ = v25
                        default:
                            var v26 int8 = -118
                            _ = v26
                        }
                        _ = v22
                    }
                    fmt.Println(v17, fi20)
                    _ = v21
                } else {
                    v2 += 1.5
                }
                fmt.Println(v3)
                _ = fi20
            }
        default:
            if (-106 > -123) {
                v27 := GenS3{F0: (v2 < 0.1), F1: (!false), F2: -32768}
                var v28 int8 = -1
                _ = v27
                _ = v28
            }
        }
    } else {
        v29 := []int{171, -408, 4294967295, 0}
        _ = v29
    }
    v30 := []int{128, 32767, 879, 127}
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    v31 := GenS3{F0: (1 < 255), F1: false, F2: 32767}
    if false {
        v32 := GenS4{F0: 1}
        if false {
            v2 = float32(0.0)
            if (65535 == v17) {
                v33 := GenS4{F0: 127}
                for fi34 := 0; fi34 < 12; fi34++ {
                    v35 := GenS2{F0: 212, F1: 32767, F2: 1}
                    _ = fi34
                    _ = v35
                }
                _ = v33
            } else {
                v3 += (v3 << ((65535) & 63))
            }
        } else {
            v36 := GenS4{F0: 1}
            var v37 int32 = v3
            v37 = v3
            _ = v36
            _ = v37
        }
        var v38 int64 = 9223372036854775807
        _ = v32
        _ = v38
    } else {
        var v39 uint8 = 42
        v39 += v39
        _ = v39
    }
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    v17 += (v17 + (v17 >> ((65535) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
