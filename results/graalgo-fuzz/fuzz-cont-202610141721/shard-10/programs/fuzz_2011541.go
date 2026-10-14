package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int8; F1 string; F2 int8 }
type GenS3 struct { F0 int32; F1 uint; F2 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 bool; F2 string }
func genFn0(p0 uint, p1 int, p2 string) int8 {
    switch 32768 {
    case 2147483647:
        switch (-(p1)) {
        case 958:
            var v7 float64 = 1e10
            _ = v7
        case 32768:
            if (true && false) {
                v8 := GenS1{F0: p1, F1: 32767}
                p0 = p0
                _ = v8
            } else {
                if false {
                    v9 := []int{}
                    _ = v9
                }
                p2 = ("foo bar" + p2)
            }
        case 72:
            v10 := []int{}
            _ = v10
        default:
            var v11 uint16 = 42
            _ = v11
        }
    case 1:
        v12 := []int{p1}
        _ = v12
    default:
        var v13 int8 = 100
        _ = v13
    }
    if false {
        v14 := GenS2{F0: 127, F1: "a", F2: 1}
        if (p1 != 32768) {
            var v15 string = ("hello" + "")
            fmt.Println(p2)
            var v16 int64 = -669
            _ = v15
            _ = v16
        } else {
            if (-326 >= -1) {
                var v17 bool = false
                var v18 bool = (float32(((1.5 - float64(1.5)) * 0.0)) > 0.0)
                _ = v17
                _ = v18
            } else {
                switch -1 {
                case 325:
                    if false {
                        switch (p1 / ((p1) | 1)) {
                        case 1:
                            fmt.Println(p2, p0)
                        }
                    } else {
                        v19 := []int{p1, 1}
                        _ = len(v19)
                        _ = v19
                    }
                case 1:
                    for fi20 := 0; fi20 < 16; fi20++ {
                        switch (p2 + "foo bar") {
                        case "\n\t":
                            v21 := GenS3{F0: 42, F1: 18446744073709551615, F2: 930}
                            _ = v21
                        case "hello":
                            v22 := []int{}
                            _ = v22
                        case "foo bar":
                            fmt.Println(fi20, p1)
                        }
                        p2 = p2
                        _ = fi20
                    }
                }
                if (!((p1 * 128) >= (p1 | (^(p1))))) {
                    var v23 int16 = 367
                    _ = v23
                }
            }
            p1 = (p1 | (p1 << (((p0 + (p0 * 42))) & 63)))
        }
        _ = v14
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(p0, p2, p1)
        _ = fi24
    }
    _ = p0
    _ = p1
    _ = p2
    return 27
}

func main() {
    v1 := []string{}
    var v2 uint32 = 2147483648
    v3 := make(map[string]int)
    v4 := make(map[int]int)
    var v5 int = -1
    var v6 uint = 651
    fmt.Println(v6)
    v6 = v6
    v5 = v5
    v6 -= (v6 - v6)
    v4[v5] = v5
    if false {
        v2 = (v2 + 902)
        var v25 int16 = 792
        var v26 uint64 = 2
        _ = v25
        _ = v26
    }
    fmt.Println(v5, v2)
    switch "foo bar" {
    case "foo bar":
        v27 := GenS3{F0: 0, F1: v6, F2: 18446744073709551615}
        _ = v27
    case "a":
        fmt.Println(v2, v6, v2)
    case "hello":
        switch (^(v5)) {
        case 0:
            switch "foo bar" {
            case "hello":
                var v28 int64 = 9223372036854775807
                _ = v28
            }
        case 907:
            v29 := GenS2{F0: 1, F1: (("hello" + "\n\t") + ""), F2: -69}
            _ = v29
        }
    }
    _ = len(v1)
    if (!false) {
        v6 = v6
        var v30 bool = false
        delete(v4, v5)
        _ = v30
    } else {
        fmt.Println(v2, v5)
        for fi31 := 0; fi31 < 11; fi31++ {
            fmt.Println(v2, fi31)
            fmt.Println(v5, v2, v5)
            v6 = 255
            _ = fi31
        }
        v5 = v5
    }
    if true {
        if (127 != -128) {
            v32 := GenS2{F0: 127, F1: "a", F2: 42}
            _ = v32
        }
        if false {
            if (42 > -9223372036854775808) {
                v33 := GenS4{F0: (100 == genFn0(487, (v5 ^ 32767), "a"))}
                for fi34 := 0; fi34 < 11; fi34++ {
                    v3["foo bar"] = v5
                    fmt.Println(fi34, fi34)
                    v2 += 894
                    _ = fi34
                }
                v2 = v2
                _ = v33
            } else {
                v6 = (^(v6))
            }
            var v35 uint = (v6 % (((^(v6))) | 1))
            _ = v35
        }
        v36 := GenS2{F0: -1, F1: ("\n\t" + ""), F2: -1}
        _ = v36
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    switch ((("hello" + "foo bar") + "foo bar") + ("a" + "foo bar")) {
    case "foo bar":
        fmt.Println(v5, v5)
    case "\n\t":
        var v37 uint64 = 42
        _ = v37
    case "":
        if false {
            fmt.Println(v5, v5, v2)
            if ((^(v5)) < (^(v5))) {
                fmt.Println(v2)
            } else {
                v38 := GenS1{F0: (v5 + (v5 % ((v5) | 1))), F1: 32767}
                _ = v38
            }
            for fi39 := 0; fi39 < 18; fi39++ {
                switch "foo bar" {
                case "a":
                    for fi40 := 0; fi40 < 19; fi40++ {
                        v2 = (v2 % ((v2) | 1))
                        if true {
                            v2 -= v2
                            var v41 uint64 = 324
                            _ = v41
                        }
                        _ = fi40
                    }
                default:
                    v2 -= v2
                }
                var v42 uint64 = uint64(v6)
                v5 += fi39
                _ = fi39
                _ = v42
            }
        } else {
            for fi43 := 0; fi43 < 17; fi43++ {
                v44 := GenS5{F0: v2, F1: true, F2: "hello"}
                v45 := GenS3{F0: 1, F1: v6, F2: 42}
                fmt.Println(v6, v5)
                _ = fi43
                _ = v44
                _ = v45
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
