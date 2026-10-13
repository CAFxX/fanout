package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int16; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint16, p1 int32) int32 {
    var v4 uint8 = 255
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: v4, F1: 4294967295, F2: -2147483648}
        var v7 bool = true
        _ = fi5
        _ = v6
        _ = v7
    }
    if (p0 <= 0) {
        for fi8 := 0; fi8 < 6; fi8++ {
            p0 = 0
            for fi9 := 0; fi9 < 1; fi9++ {
                v4 -= v4
                _ = fi9
            }
            _ = fi8
        }
        p0 -= p0
    } else {
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS2{F0: 2, F1: 4294967295}
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return (p1 >> ((255) & 63))
}

func main() {
    v1 := []uint32{42, 2147483647, 1, 2147483647}
    var v2 int = 852
    var v3 int = (^(v2))
    var v12 uint8 = 1
    switch (^(v3)) {
    case -132:
        v13 := GenS1{F0: v12, F1: 920, F2: 2147483647}
        _ = v13
    }
    v1 = append(v1, 482)
    for fi14 := 0; fi14 < 18; fi14++ {
        v12 += v12
        v15 := GenS3{F0: 65535, F1: (v3 - fi14), F2: v12}
        _ = fi14
        _ = v15
    }
    switch v2 {
    case 1:
        var v16 uint16 = 0
        _ = v16
    case -9223372036854775808:
        switch (v2 / ((v2) | 1)) {
        case 1:
            switch v3 {
            case 9223372036854775807:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                fmt.Println(v3, v3)
            }
        case 178:
            var v17 uint8 = 61
            _ = v17
        case 128:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            switch (("a" + ("" + "a")) + "hello") {
            case "\n\t":
                v3 -= (v3 * 2147483647)
            case "hello":
                v18 := GenS3{F0: 255, F1: -88, F2: v12}
                _ = v18
            case "foo bar":
                var v19 int8 = 1
                _ = v19
            default:
                for fi20 := 0; fi20 < 14; fi20++ {
                    v2 -= 161
                    _ = fi20
                }
            }
        }
    case 127:
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 18; fi22++ {
                if len(v1) > 0 {
                    v1[0] = 40
                }
                var v23 uint64 = 157
                var v24 int8 = 42
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
    }
    v3 -= (v3 ^ 1)
    for fi25 := 0; fi25 < 19; fi25++ {
        v2 = (v2 * v3)
        var v26 int64 = 1
        switch (fi25 ^ -125) {
        case -9223372036854775808:
            for fi27 := 0; fi27 < 7; fi27++ {
                v28 := GenS4{F0: 255, F1: 19, F2: -103}
                if true {
                    if len(v1) > 0 {
                        v1[0] = 260
                    }
                    v2 = fi25
                    var v29 bool = ((3.14159 / float64(3.14159)) <= float64((((1.5 / float32(0.1)) / 1.0) * ((1.5 + 1.5) / float32(float64(0.1))))))
                    _ = v29
                }
                _ = fi27
                _ = v28
            }
        case 9223372036854775807:
            switch -2147483648 {
            case 32767:
                var v30 float32 = 1.5
                _ = v30
            }
        default:
            if (false || true) {
                var v31 string = (("\n\t" + ("" + "a")) + "a")
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v31
            }
        }
        _ = fi25
        _ = v26
    }
    if (!(42 == 257)) {
        v32 := GenS4{F0: 42, F1: -539, F2: -99}
        v2 -= v3
        _ = v32
    } else {
        v33 := GenS5{F0: 255}
        _ = v33
    }
    v2 = -1
    fmt.Println(v12, v12, v12)
    for fi34 := 0; fi34 < 5; fi34++ {
        if (!((-128 == 127) && (1 >= 1))) {
            fmt.Println(v12, v3, v2)
            v35 := GenS2{F0: 618, F1: 65535}
            var v36 int8 = 97
            _ = v35
            _ = v36
        }
        _ = fi34
    }
    v37 := GenS3{F0: 622, F1: (v2 >> ((0) & 63)), F2: 42}
    switch 1 {
    case 9223372036854775807:
        for fi38 := 0; fi38 < 18; fi38++ {
            for fi39 := 0; fi39 < 8; fi39++ {
                fmt.Println(fi38, fi39, fi38)
                _ = fi39
            }
            v1 = append(v1, 42)
            _ = fi38
        }
    case -9223372036854775808:
        v3 += (v2 % ((v3) | 1))
    default:
        var v40 float64 = 1.5
        _ = v40
    }
    if len(v1) > 0 {
        v1[0] = 979
    }
    if true {
        var v41 int16 = -32768
        _ = v41
    } else {
        for fi42 := 0; fi42 < 20; fi42++ {
            switch (-(v3)) {
            case -908:
                switch v2 {
                case 128:
                    v43 := GenS6{F0: 0}
                    _ = v43
                case 72:
                    var v44 int8 = 0
                    _ = v44
                case -345:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    v12 += v12
                }
            default:
                v2 = v2
            }
            if (!(true || true)) {
                v2 = (v3 - fi42)
                v12 += v12
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
            }
            var v45 uint8 = (v12 & v12)
            _ = fi42
            _ = v45
        }
        v46 := GenS2{F0: 221, F1: 42}
        _ = v46
    }
    v2 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v37
}
