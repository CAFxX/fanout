package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 bool; F1 uint32; F2 int64 }
type GenS8 struct { F0 int8; F1 uint; F2 uint8 }
func genFn0(p0 int, p1 uint) uint16 {
    var v7 uint64 = 0
    v7 += (^(v7))
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{}
    var v3 uint = 1
    v4 := make(map[int]string)
    v5 := []int{1, -1, 32768}
    var v6 int64 = 1
    delete(v4, 127)
    if (!(4294967295 >= 652)) {
        v8 := GenS1{F0: -1, F1: 102, F2: 1}
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := v1[24]
            v11 = (42 >= 42)
            var v12 float32 = -2.25
            _ = fi10
            _ = v11
            _ = v12
        }
        if true {
            if ((false && false) || (true && (("foo bar" + "\n\t") == (("foo bar" + "") + "hello")))) {
                delete(v4, (fi9 >> ((65535) & 63)))
            } else {
                v6 -= v6
                var v13 uint32 = 42
                _ = v13
            }
            switch "a" {
            case "hello":
                var v14 int16 = -890
                _ = v14
            case "a":
                var v15 uint16 = 1
                _ = v15
            default:
                fmt.Println(fi9, fi9)
            }
        } else {
            for fi16 := 0; fi16 < 6; fi16++ {
                fmt.Println(fi9, fi16, v3)
                _ = fi16
            }
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 uint8 = 0
                var v19 int = (fi17 & (fi9 >> ((1) & 63)))
                var v20 bool = true
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        }
        _ = fi9
    }
    _ = len(v2)
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 uint16 = 1
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 uint8 = 45
            if ("" > "a") {
                _ = len(v4)
            } else {
                var v25 uint = v3
                var v26 int64 = (v6 % ((1) | 1))
                _ = v25
                _ = v26
            }
            fmt.Println(v6, v22)
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    var v27 uint64 = 719
    v28 := GenS2{F0: false, F1: 100}
    v3 += (v3 / ((v3) | 1))
    v29 := GenS2{F0: true, F1: 1}
    for fi30 := 0; fi30 < 15; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            v3 = v3
            _ = fi31
        }
        _ = fi30
    }
    switch "a" {
    case "a":
        switch -1 {
        case -775:
            var v32 int = -9223372036854775808
            _ = v32
        case -918:
            v27 -= v27
        default:
            for fi33 := 0; fi33 < 2; fi33++ {
                for fi34 := 0; fi34 < 11; fi34++ {
                    var v35 uint = (v3 | (v3 % (((^(v3))) | 1)))
                    switch (("\n\t" + ("hello" + "foo bar")) + ("\n\t" + "\n\t")) {
                    case "":
                        v36 := GenS3{F0: 42}
                        _ = v36
                    case "foo bar":
                        v37 := GenS2{F0: (true || (127 > -1)), F1: -1}
                        _ = v37
                    case "a":
                        var v38 float32 = 0.1
                        _ = v38
                    default:
                        v2 = append(v2, -9223372036854775808)
                        var v39 float64 = 0.0
                        _ = v39
                    }
                    var v40 int8 = 1
                    _ = fi34
                    _ = v35
                    _ = v40
                }
                _ = fi33
            }
        }
    case "\n\t":
        var v41 uint32 = 2147483648
        _ = v41
    case "foo bar":
        if (!true) {
            v4[-195] = ("foo bar" + "")
            switch (("hello" + "a") + ("foo bar" + "foo bar")) {
            case "hello":
                fmt.Println(v6, v3, v27)
            case "\n\t":
                fmt.Println(v27)
            }
        }
    default:
        if true {
            v42 := GenS4{F0: -8}
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 4; fi43++ {
        for fi44 := 0; fi44 < 2; fi44++ {
            if false {
                v45 := GenS5{F0: -32768, F1: -40}
                v46 := GenS3{F0: 536}
                _ = v45
                _ = v46
            }
            switch fi43 {
            case 404:
                switch (("a" + ("\n\t" + "foo bar")) + "foo bar") {
                case "hello":
                    fmt.Println(fi44, v3)
                default:
                    var v47 int16 = 1
                    _ = v47
                }
            case -851:
                v27 = (v27 % (((v27 + (v27 / ((522) | 1)))) | 1))
            case 32768:
                if false {
                    _ = len(v5)
                }
            default:
                var v48 float64 = 0.0
                _ = v48
            }
            fmt.Println(v3)
            _ = fi44
        }
        v3 = 484
        if (false || true) {
            var v49 int64 = (v6 / ((42) | 1))
            v49 = -1
            v6 = (v6 + (^(v6)))
            _ = v49
        }
        _ = fi43
    }
    v3 = (v3 ^ v3)
    v50 := GenS2{F0: (1e10 != (-2.25 * 0.1)), F1: 0}
    v51 := GenS6{F0: (^(v27)), F1: -2147483648}
    for fi52 := 0; fi52 < 2; fi52++ {
        v53 := GenS2{F0: false, F1: 117}
        for fi54 := 0; fi54 < 17; fi54++ {
            v55 := GenS7{F0: (fi52 > 9223372036854775807), F1: 4294967295, F2: v6}
            v56 := GenS8{F0: 1, F1: 1, F2: 128}
            _ = fi54
            _ = v55
            _ = v56
        }
        _ = fi52
        _ = v53
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
