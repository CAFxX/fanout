package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int; F1 uint64; F2 uint8 }
type GenS5 struct { F0 string; F1 int; F2 int8 }
type GenS6 struct { F0 int; F1 int8; F2 uint16 }

func main() {
    var v1 bool = (2147483648 != 847)
    var v2 int8 = 127
    var v3 uint8 = 127
    v4 := make(map[string]uint32)
    var v5 int16 = -253
    var v6 bool = v1
    v5 -= (^(v5))
    v7 := GenS1{F0: 2147483648, F1: (v2 << ((42) & 63))}
    switch ("\n\t" + "a") {
    case "":
        v2 += v2
    case "foo bar":
        fmt.Println(v2, v5, v5)
    default:
        var v8 int8 = v2
        _ = v8
    }
    v2 -= (v2 << ((42) & 63))
    var v9 uint8 = v3
    if v6 {
        switch (("" + ("foo bar" + ("a" + "\n\t"))) + "hello") {
        case "":
            v10 := GenS2{F0: 814}
            _ = v10
        case "\n\t":
            v5 -= v5
        default:
            v11 := GenS3{F0: false, F1: 1}
            _ = v11
        }
        var v12 int32 = 1
        _ = v12
    }
    v9 -= v9
    v13 := GenS3{F0: v6, F1: 1}
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(v9, v5, fi14)
        fmt.Println(v6, v6)
        v15 := GenS4{F0: 698, F1: 273, F2: v3}
        _ = fi14
        _ = v15
    }
    switch "" {
    case "hello":
        v16 := GenS3{F0: (-762 < -1), F1: -948}
        _ = v16
    case "a":
        if ((float64(1.5) + 0.1) != ((0.0 + 3.14159) + 1e10)) {
            switch ("hello" + "\n\t") {
            case "hello":
                v17 := []int{}
                _ = v17
            case "foo bar":
                var v18 int16 = v5
                _ = v18
            case "\n\t":
                v19 := GenS4{F0: -75, F1: 1, F2: v3}
                _ = v19
            }
        }
    default:
        v9 += (v3 & 159)
    }
    var v20 int64 = 0
    v9 -= v9
    var v21 string = ""
    switch v21 {
    case "":
        var v22 string = ("a" + ((v21 + v21) + v21))
        _ = v22
    default:
        if (int64(v9) < (-(v20))) {
            var v23 int32 = 42
            if v1 {
                if v6 {
                    var v24 uint64 = 411
                    v25 := []int{32768, int(v23), -1}
                    var v26 float64 = (3.14159 * 0.0)
                    _ = v24
                    _ = v25
                    _ = v26
                }
                fmt.Println(v23, v21)
            }
            switch v21 {
            case "a":
                var v27 uint16 = 32767
                _ = v27
            case "\n\t":
                _ = len(v4)
            }
            _ = v23
        } else {
            v28 := []int{42, 125}
            v29 := v4[v21]
            _ = v28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        v3 = 2
        if v6 {
            var v31 string = (((v21 + (v21 + "foo bar")) + (v21 + v21)) + v21)
            fmt.Println(v5, v3, v6)
            _ = v31
        } else {
            var v32 int8 = (v2 >> ((2) & 63))
            v33 := GenS3{F0: v1, F1: 0}
            fmt.Println(v9)
            _ = v32
            _ = v33
        }
        v34 := v4[v21]
        _ = fi30
        _ = v34
    }
    v35 := GenS5{F0: (v21 + v21), F1: -2147483648, F2: v2}
    fmt.Println(v3)
    if v6 {
        v36 := GenS3{F0: (1 < v20), F1: 42}
        switch ((v21 + v21) + (("foo bar" + v21) + v21)) {
        case "hello":
            for fi37 := 0; fi37 < 4; fi37++ {
                v9 = v3
                v38 := GenS6{F0: (fi37 - fi37), F1: (v2 % ((v2) | 1)), F2: uint16(v3)}
                if (!(true || ((("foo bar" + (("" + ((v21 + v21) + v21)) + "a")) + "foo bar") < v21))) {
                    var v39 uint64 = 1
                    if true {
                        for fi40 := 0; fi40 < 2; fi40++ {
                            var v41 bool = v1
                            fmt.Println(v41, v5)
                            _ = len(v4)
                            _ = fi40
                            _ = v41
                        }
                        for fi42 := 0; fi42 < 15; fi42++ {
                            v9 -= (v9 % (((v3 >> ((49) & 63))) | 1))
                            var v43 uint = 18446744073709551615
                            _ = fi42
                            _ = v43
                        }
                        var v44 uint = 65535
                        _ = v44
                    } else {
                        delete(v4, v21)
                        switch (v21 + "") {
                        case "a":
                            var v45 float64 = 1e10
                            _ = v45
                        case "":
                            var v46 uint32 = 747
                            _ = v46
                        case "\n\t":
                            var v47 int = (fi37 - (fi37 >> ((18446744073709551615) & 63)))
                            _ = v47
                        default:
                            v9 = (v9 - v3)
                        }
                    }
                    _ = v39
                }
                _ = fi37
                _ = v38
            }
        case "\n\t":
            switch v21 {
            case "a":
                v48 := GenS4{F0: -539, F1: 0, F2: v3}
                _ = v48
            default:
                switch v21 {
                case "\n\t":
                    v49 := GenS4{F0: -637, F1: 4294967295, F2: (v3 ^ v3)}
                    _ = v49
                case "hello":
                    var v50 int64 = (^(v20))
                    _ = v50
                }
            }
        default:
            v51 := []int{2147483647, -1}
            _ = v51
        }
        for fi52 := 0; fi52 < 4; fi52++ {
            for fi53 := 0; fi53 < 1; fi53++ {
                v9 = (v3 / (((v9 >> ((2) & 63))) | 1))
                _ = fi53
            }
            _ = len(v4)
            _ = fi52
        }
        _ = v36
    } else {
        v54 := []int{32767}
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
