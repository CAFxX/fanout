package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int; F1 uint; F2 int16 }
func genFn0(p0 bool, p1 bool) int {
    var v4 int32 = 0
    if ((1 <= (v4 / ((v4) | 1))) && p0) {
        for fi5 := 0; fi5 < 4; fi5++ {
            fmt.Println(fi5, v4, p0)
            fmt.Println(p1, v4)
            switch ("hello" + "foo bar") {
            case "":
                v6 := GenS1{F0: "", F1: (-(v4)), F2: 128}
                _ = v6
            }
            _ = fi5
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 32768
}
func genFn1(p0 int64) bool {
    var v7 int64 = (p0 << ((2) & 63))
    switch (("foo bar" + "foo bar") + ("\n\t" + "hello")) {
    case "\n\t":
        v7 = v7
    case "":
        var v8 bool = true
        _ = v8
    }
    _ = p0
    _ = v7
    return (0.1 > 1e10)
}

func main() {
    var v1 int = 42
    v2 := make(map[string]bool)
    var v3 float64 = 3.14159
    v9 := GenS2{F0: 0, F1: 641, F2: 518}
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int32 = 527
        v12 := GenS3{F0: -939, F1: 637, F2: -762}
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 = v3
    if (!(45 != 42)) {
        var v13 float32 = float32(((0.1 + v3) * v3))
        _ = v13
    } else {
        var v14 int64 = -9223372036854775808
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v1 -= fi15
        for fi16 := 0; fi16 < 9; fi16++ {
            switch genFn0((float64((((((float32(v3) / -2.25) + 3.14159) * 1e10) + 1.5) * -2.25)) > float64(1e10)), false) {
            case -390:
                var v17 int16 = -1
                _ = v17
            case 1:
                v1 -= (fi15 ^ v1)
            }
            _ = fi16
        }
        _ = fi15
    }
    if (4294967295 == 0) {
        v3 = 0.1
        var v18 uint64 = 991
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(fi19)
            var v20 int = 2147483647
            if false {
                fmt.Println(fi19, fi19, fi19)
                var v21 uint = 160
                _ = v21
            }
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        for fi22 := 0; fi22 < 18; fi22++ {
            v1 += fi22
            var v23 uint = 244
            _ = fi22
            _ = v23
        }
    }
    v24 := GenS2{F0: -242, F1: 42, F2: 2147483648}
    if true {
        v1 = v1
        v25 := GenS2{F0: 0, F1: 909, F2: 4294967295}
        _ = v25
    } else {
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 int = -143
            var v28 float32 = 1.5
            _ = fi26
            _ = v27
            _ = v28
        }
        if (0 <= -2147483648) {
            var v29 int32 = 2147483647
            v1 = (v1 ^ v1)
            var v30 int32 = (v29 << ((65535) & 63))
            _ = v29
            _ = v30
        } else {
            v1 = v1
        }
        switch 42 {
        case -475:
            switch v1 {
            case -527:
                v1 -= 2147483647
            case 127:
                v31 := v2["foo bar"]
                _ = v31
            case 32767:
                for fi32 := 0; fi32 < 11; fi32++ {
                    v33 := GenS1{F0: ("hello" + "hello"), F1: 109, F2: 128}
                    if (true && (1.5 == ((0.1 * 0.0) / 3.14159))) {
                        v1 = v1
                        for fi34 := 0; fi34 < 13; fi34++ {
                            v1 = v1
                            v35 := GenS1{F0: "a", F1: -1, F2: 2}
                            v36 := GenS3{F0: (fi32 % ((32767) | 1)), F1: 0, F2: -629}
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    } else {
                        for fi37 := 0; fi37 < 9; fi37++ {
                            v38 := GenS2{F0: 42, F1: 688, F2: 2147483648}
                            fmt.Println(fi37)
                            _ = fi37
                            _ = v38
                        }
                        v3 += 1.5
                    }
                    _ = fi32
                    _ = v33
                }
            }
        case -66:
            switch "\n\t" {
            case "hello":
                var v39 int8 = 50
                _ = v39
            default:
                var v40 int = v1
                _ = v40
            }
        case 42:
            switch genFn0(true, false) {
            case 127:
                var v41 int8 = 42
                _ = v41
            case 345:
                var v42 uint8 = 1
                _ = v42
            case 858:
                v3 -= 1e10
            }
        default:
            delete(v2, "")
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v9
}
