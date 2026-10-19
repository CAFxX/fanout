package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint64; F1 int; F2 string }
type GenS4 struct { F0 int; F1 string; F2 int }
type GenS5 struct { F0 uint16; F1 uint; F2 int }
type GenS6 struct { F0 bool; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 1
    var v3 float32 = ((0.1 * (float32(1.5) * float32(float64(3.14159)))) * 1.5)
    v4 := []string{""}
    v5 := make(map[int]uint32)
    v6 := make(map[string]int)
    fmt.Println(v2, v3)
    switch 624 {
    case 32767:
        fmt.Println(v2)
    default:
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: 127, F1: -53}
            if true {
                for fi9 := 0; fi9 < 2; fi9++ {
                    v3 = (v3 - (float32(1e10) - 1.5))
                    var v10 int32 = -2147483648
                    v2 += 1
                    _ = fi9
                    _ = v10
                }
            }
            _ = fi7
            _ = v8
        }
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 string = "a"
        v13 := v1[fi11]
        if (42 >= -9223372036854775808) {
            var v14 uint64 = 4294967295
            v15 := GenS2{F0: -1, F1: (true || ((v2 >= 2) || false))}
            var v16 uint = 2
            _ = v14
            _ = v15
            _ = v16
        } else {
            switch "hello" {
            case "":
                var v17 uint64 = 113
                _ = v17
            case "foo bar":
                var v18 int32 = -549
                _ = v18
            case "\n\t":
                v3 -= (1e10 * v3)
            }
            v3 -= v3
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v2)
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        var v19 int32 = -1
        _ = v19
    default:
        var v20 uint64 = 1
        _ = v20
    }
    var v21 int8 = 42
    v22 := GenS1{F0: 1, F1: 42}
    switch "" {
    case "foo bar":
        for fi23 := 0; fi23 < 6; fi23++ {
            for fi24 := 0; fi24 < 18; fi24++ {
                _ = len(v5)
                _ = fi24
            }
            for fi25 := 0; fi25 < 1; fi25++ {
                fmt.Println(v2, v3, v21)
                _ = fi25
            }
            _ = fi23
        }
    default:
        switch ("foo bar" + ("" + "a")) {
        case "a":
            if len(v4) > 0 {
                v4[0] = ("hello" + "foo bar")
            }
        default:
            v26 := GenS3{F0: 42, F1: -895, F2: (("hello" + ((("foo bar" + ("" + ("hello" + "a"))) + "hello") + "")) + "hello")}
            _ = v26
        }
    }
    switch "foo bar" {
    case "":
        for fi27 := 0; fi27 < 6; fi27++ {
            switch fi27 {
            case 9223372036854775807:
                if false {
                    fmt.Println(v3, fi27, v21)
                } else {
                    var v28 int8 = (-(v21))
                    _ = v28
                }
            case 4294967295:
                if true {
                    for fi29 := 0; fi29 < 13; fi29++ {
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                        fmt.Println(v21, v21)
                        _ = fi29
                    }
                    var v30 uint = 0
                    _ = v30
                }
            case -2147483648:
                var v31 float32 = 0.0
                _ = v31
            }
            if false {
                if true {
                    var v32 bool = (false && ("hello" < "\n\t"))
                    _ = v32
                } else {
                    if ((float32(0.1) > v3) && false) {
                        v33 := GenS4{F0: (fi27 / ((fi27) | 1)), F1: "", F2: (^(fi27))}
                        _ = v33
                    } else {
                        var v34 int64 = 9223372036854775807
                        v35 := GenS5{F0: 243, F1: v2, F2: -695}
                        _ = v34
                        _ = v35
                    }
                    var v36 int32 = 2147483647
                    _ = v36
                }
                v37 := v5[fi27]
                var v38 string = "\n\t"
                _ = v37
                _ = v38
            }
            _ = fi27
        }
    }
    v39 := GenS6{F0: ("" != "foo bar"), F1: 4294967295, F2: -22}
    v40 := GenS4{F0: 234, F1: "a", F2: 32768}
    for fi41 := 0; fi41 < 2; fi41++ {
        for fi42 := 0; fi42 < 10; fi42++ {
            if ((false && true) && (42 == v21)) {
                var v43 bool = false
                fmt.Println(fi42, v2, v3)
                fmt.Println(v3, v43)
                _ = v43
            }
            var v44 uint16 = uint16(v2)
            var v45 int32 = 929
            _ = fi42
            _ = v44
            _ = v45
        }
        _ = fi41
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        for fi47 := 0; fi47 < 17; fi47++ {
            var v48 int64 = -9223372036854775808
            _ = fi47
            _ = v48
        }
        for fi49 := 0; fi49 < 1; fi49++ {
            delete(v1, fi46)
            v50 := GenS4{F0: (fi46 * (fi46 << (((^(v2))) & 63))), F1: ("a" + "hello"), F2: -783}
            v51 := GenS3{F0: 268, F1: (fi46 ^ fi46), F2: ("a" + ("\n\t" + ("\n\t" + "")))}
            _ = fi49
            _ = v50
            _ = v51
        }
        v3 = float32(1.5)
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
