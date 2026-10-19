package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint16; F1 bool; F2 string }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 41
    v3 := []int{128, -9223372036854775808, 431, -608}
    var v4 string = ("" + "a")
    v5 := []string{(v4 + v4), ("hello" + "\n\t")}
    v6 := []uint32{}
    if (-1 == -2147483648) {
        switch 127 {
        case 367:
            for fi7 := 0; fi7 < 4; fi7++ {
                switch v4 {
                case "\n\t":
                    var v8 int32 = int32(v2)
                    _ = v8
                }
                _ = fi7
            }
        }
        fmt.Println(v2, v4)
        v3 = append(v3, -274)
    }
    if (false || true) {
        var v9 float64 = 1.5
        var v10 uint32 = 779
        for fi11 := 0; fi11 < 2; fi11++ {
            for fi12 := 0; fi12 < 16; fi12++ {
                fmt.Println(fi11)
                v13 := GenS1{F0: 1}
                v14 := GenS2{F0: 1}
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 uint = 463
            var v16 uint8 = 2
            _ = fi11
            _ = v15
            _ = v16
        }
        _ = v10
        _ = v9
    }
    v3 = append(v3, -9223372036854775808)
    fmt.Println(v4)
    v17 := GenS2{F0: 42}
    var v18 string = ("hello" + v4)
    _ = len(v3)
    var v19 int32 = 42
    for fi20 := 0; fi20 < 10; fi20++ {
        v1[(fi20 << (((v2 * (v2 | v2))) & 63))] = (510 >= 18446744073709551615)
        if (127 != 128) {
            if (1 > -1) {
                switch fi20 {
                case 32767:
                    v19 += (v19 - v19)
                case 9223372036854775807:
                    v19 = (^(v19))
                case 32768:
                    switch ("" + "a") {
                    case "\n\t":
                        v1[(fi20 & int(v2))] = (!(-9223372036854775808 > 1))
                    case "foo bar":
                        v2 = v2
                    }
                default:
                    v21 := GenS3{F0: 4294967295}
                    _ = v21
                }
                var v22 uint8 = 80
                v23 := GenS2{F0: -194}
                _ = v22
                _ = v23
            } else {
                var v24 int64 = 9223372036854775807
                for fi25 := 0; fi25 < 6; fi25++ {
                    _ = len(v5)
                    var v26 float32 = (3.14159 - (float32(0.0) * 1e10))
                    var v27 bool = false
                    v18 = ((("" + ((v4 + v18) + ((("foo bar" + v18) + "\n\t") + "a"))) + "foo bar") + v18)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = v24
            }
            v28 := GenS4{F0: -1, F1: false}
            switch "a" {
            case "":
                if len(v3) > 0 {
                    v3[0] = 127
                }
            case "\n\t":
                var v29 uint = (^(v2))
                _ = v29
            }
            _ = v28
        }
        _ = fi20
    }
    switch ("\n\t" + v18) {
    case "":
        var v30 uint16 = 754
        _ = v30
    }
    var v31 int = int(v2)
    for fi32 := 0; fi32 < 5; fi32++ {
        v33 := GenS1{F0: 4294967295}
        for fi34 := 0; fi34 < 15; fi34++ {
            v18 = v4
            switch (fi34 >> ((v2) & 63)) {
            case -451:
                var v35 int = (-(fi34))
                _ = v35
            }
            _ = fi34
        }
        switch (v31 << ((v2) & 63)) {
        case 2147483647:
            v4 = "\n\t"
        default:
            v18 = ("a" + v4)
        }
        _ = fi32
        _ = v33
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 bool = (!(!false))
        if v37 {
            fmt.Println(v4, v2, v19)
            fmt.Println(fi36, v4, v2)
            v2 = (v2 >> ((v2) & 63))
        }
        _ = fi36
        _ = v37
    }
    v38 := GenS5{F0: 1, F1: (4294967295 != 4294967295), F2: v18}
    switch (v31 / ((32767) | 1)) {
    case 32767:
        v5 = append(v5, "\n\t")
    default:
        if false {
            v19 -= (v19 - (^(v19)))
            for fi39 := 0; fi39 < 11; fi39++ {
                fmt.Println(fi39)
                v5 = append(v5, v4)
                _ = fi39
            }
        } else {
            var v40 float32 = 1.5
            _ = v40
        }
    }
    var v41 int32 = (^(v19))
    v41 -= (v19 ^ (v19 | (-(v19))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
