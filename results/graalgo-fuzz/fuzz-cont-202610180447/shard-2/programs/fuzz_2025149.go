package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint16; F2 int64 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8 }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]int)
    var v3 int8 = 1
    v4 := make(map[string]uint32)
    v5 := GenS1{F0: 255}
    if (2147483647 == -527) {
        fmt.Println(v3)
        var v6 int64 = 892
        _ = v6
    } else {
        switch 1 {
        case 128:
            if ((!true) || ((float64(1.5) / -2.25) > (1e10 * 1.5))) {
                for fi7 := 0; fi7 < 8; fi7++ {
                    switch "a" {
                    case "\n\t":
                        var v8 string = "a"
                        _ = v8
                    }
                    _ = fi7
                }
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 int64 = -9223372036854775808
                    fmt.Println(v10, v1)
                    v11 := GenS1{F0: 65535}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v12 := []int{0, -9223372036854775808, 0}
                _ = v12
            }
        case 42:
            var v13 bool = (261 < 1)
            _ = v13
        }
        v14 := GenS1{F0: 255}
        v15 := GenS1{F0: 0}
        _ = v14
        _ = v15
    }
    var v16 uint32 = 65535
    v17 := GenS2{F0: -9223372036854775808, F1: v1, F2: 42}
    if false {
        if false {
            var v18 int32 = -236
            if false {
                if false {
                    if true {
                        v19 := []int{-9223372036854775808, -9223372036854775808}
                        var v20 int32 = v18
                        _ = v19
                        _ = v20
                    }
                    if (((((("\n\t" + ("\n\t" + ("" + ""))) + "foo bar") + "foo bar") != "a") || (!true)) || ((v16 & v16) < v16)) {
                        switch "a" {
                        case "a":
                            v21 := []int{-899, 1, 15, 2147483647}
                            _ = v21
                        default:
                            var v22 string = ("a" + "a")
                            _ = v22
                        }
                    }
                }
                for fi23 := 0; fi23 < 6; fi23++ {
                    if ((false && false) || false) {
                        delete(v2, "hello")
                    }
                    _ = fi23
                }
                for fi24 := 0; fi24 < 19; fi24++ {
                    v1 = (^(v1))
                    if false {
                        var v25 int16 = 42
                        _ = v25
                    }
                    _ = fi24
                }
            } else {
                switch 9223372036854775807 {
                case -2147483648:
                    _ = len(v2)
                }
            }
            switch 847 {
            case 42:
                var v26 uint = 20
                _ = v26
            case -92:
                fmt.Println(v18)
            case 888:
                fmt.Println(v3)
            default:
                v3 -= (-(v3))
            }
            _ = v18
        } else {
            v1 = v1
            v27 := GenS1{F0: 4294967295}
            _ = v27
        }
        for fi28 := 0; fi28 < 17; fi28++ {
            fmt.Println(v1)
            v29 := GenS3{F0: "", F1: 1000}
            _ = fi28
            _ = v29
        }
        v30 := GenS2{F0: 1, F1: (v1 - v1), F2: 42}
        _ = v30
    } else {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v16)
        case "\n\t":
            var v31 int = 4294967295
            _ = v31
        case "a":
            if true {
                v32 := []int{-2147483648, -942}
                v1 -= v1
                for fi33 := 0; fi33 < 18; fi33++ {
                    if (v1 > 333) {
                        v16 += 1
                        if (3.14159 >= 1e10) {
                            v34 := GenS1{F0: 65535}
                            _ = v34
                        }
                    }
                    v3 = v3
                    _ = fi33
                }
                _ = v32
            } else {
                v35 := GenS4{F0: 0}
                var v36 int = 145
                _ = v35
                _ = v36
            }
        default:
            v1 += v1
        }
        switch "\n\t" {
        case "":
            v37 := GenS1{F0: 4294967295}
            _ = v37
        case "foo bar":
            var v38 int16 = -311
            _ = v38
        default:
            for fi39 := 0; fi39 < 17; fi39++ {
                v40 := GenS5{F0: -551}
                _ = fi39
                _ = v40
            }
        }
        switch (("foo bar" + "foo bar") + ("\n\t" + ("hello" + "\n\t"))) {
        case "foo bar":
            v2["\n\t"] = -69
        case "\n\t":
            for fi41 := 0; fi41 < 1; fi41++ {
                v3 += v3
                _ = fi41
            }
        }
    }
    v1 = 32767
    for fi42 := 0; fi42 < 5; fi42++ {
        v43 := GenS6{F0: 0}
        v44 := GenS2{F0: fi42, F1: v1, F2: 9223372036854775807}
        fmt.Println(v16)
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println(v3, v1)
    if (false || false) {
        if (1e10 < 0.1) {
            delete(v2, (("hello" + (("\n\t" + ("\n\t" + ("\n\t" + "\n\t"))) + "foo bar")) + "\n\t"))
        }
        var v45 uint = 250
        v16 = 2147483647
        _ = v45
    }
    if true {
        var v46 int = -1
        v1 -= v1
        _ = v46
    }
    v1 += (v1 >> ((4294967295) & 63))
    var v47 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
