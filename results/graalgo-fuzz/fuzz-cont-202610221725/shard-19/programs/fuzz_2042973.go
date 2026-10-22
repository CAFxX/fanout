package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int, p1 uint64) string {
    p1 = (p1 & p1)
    switch 119 {
    case 0:
        var v6 int16 = 32767
        _ = v6
    default:
        v7 := []int{(p0 * (-(p0))), 390}
        _ = v7
    }
    p1 -= (^(p1))
    var v8 bool = false
    var v9 string = "a"
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1() int32 {
    if true {
        switch "\n\t" {
        case "foo bar":
            var v10 int16 = 1
            _ = v10
        }
        var v11 int8 = 28
        _ = v11
    } else {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 uint8 = 42
            if false {
                var v14 string = ("\n\t" + "")
                _ = v14
            }
            var v15 int8 = -117
            _ = fi12
            _ = v13
            _ = v15
        }
    }
    var v16 uint16 = 377
    v17 := []int{}
    _ = v16
    _ = v17
    return 401
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]bool)
    v3 := []uint32{4294967295, 2147483648}
    var v4 int32 = 42
    var v5 uint = 1
    for fi18 := 0; fi18 < 12; fi18++ {
        v1 = 42
        if false {
            _ = len(v2)
        } else {
            var v19 uint8 = 2
            var v20 string = ("\n\t" + "a")
            _ = v19
            _ = v20
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        switch (fi21 >> (((v5 % (((^(v5))) | 1))) & 63)) {
        case 42:
            switch "a" {
            case "foo bar":
                if true {
                    _ = len(v2)
                    v1 = (v1 & v1)
                    switch 128 {
                    case 32767:
                        if false {
                            var v22 uint16 = 65535
                            var v23 int32 = (v4 / ((2147483647) | 1))
                            _ = v22
                            _ = v23
                        }
                    case 2147483647:
                        if false {
                            v4 = (v4 / ((v4) | 1))
                            if false {
                                for fi24 := 0; fi24 < 11; fi24++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v25 := GenS1{F0: ((v4 * (v4 + (v4 | v4))) < (v4 ^ v4)), F1: "a", F2: 2147483647}
                                    v4 -= (v4 & v4)
                                    _ = fi24
                                    _ = v25
                                }
                                if (!false) {
                                    v26 := GenS1{F0: (856 == 18446744073709551615), F1: "foo bar", F2: 65535}
                                    _ = v26
                                }
                            }
                            fmt.Println(v1, v5, v1)
                        }
                    case -357:
                        _ = len(v2)
                    default:
                        switch fi21 {
                        case 128:
                            v27 := GenS1{F0: (true || (!(3.14159 == -2.25))), F1: "", F2: 30}
                            _ = v27
                        case -861:
                            v4 -= (-(v4))
                        case 960:
                            v28 := GenS1{F0: true, F1: "hello", F2: 667}
                            _ = v28
                        default:
                            v4 += v4
                        }
                    }
                } else {
                    for fi29 := 0; fi29 < 11; fi29++ {
                        var v30 uint32 = 2
                        var v31 string = ("foo bar" + (("" + "a") + ""))
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    var v32 int = 42
                    v1 += v1
                    _ = v32
                }
            case "a":
                for fi33 := 0; fi33 < 7; fi33++ {
                    var v34 uint64 = 0
                    _ = fi33
                    _ = v34
                }
            case "\n\t":
                v35 := GenS2{F0: 9223372036854775807}
                _ = v35
            default:
                var v36 int8 = 0
                _ = v36
            }
        default:
            var v37 uint16 = 1
            _ = v37
        }
        var v38 bool = false
        var v39 int = fi21
        _ = fi21
        _ = v38
        _ = v39
    }
    switch 2147483647 {
    case -1:
        if true {
            var v40 int16 = 1
            v5 = (v5 << ((v5) & 63))
            _ = v40
        }
    }
    v5 += v5
    if true {
        _ = len(v3)
        fmt.Println(v1, v5, v5)
        fmt.Println(v4, v1, v4)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 += (v4 >> ((v5) & 63))
    fmt.Println(v1, v4)
    var v41 uint16 = 32767
    switch ("foo bar" + (("hello" + "hello") + "a")) {
    case "":
        fmt.Println(v4, v4, v5)
    default:
        v5 = v5
    }
    var v42 uint16 = v41
    for fi43 := 0; fi43 < 5; fi43++ {
        v1 = 1
        for fi44 := 0; fi44 < 14; fi44++ {
            fmt.Println(v42, v42)
            fmt.Println(fi44, fi44)
            fmt.Println(v42, v41, v5)
            _ = fi44
        }
        var v45 uint64 = 917
        _ = fi43
        _ = v45
    }
    v46 := GenS2{F0: (v1 % (((v1 / ((v1) | 1))) | 1))}
    v1 = v1
    v47 := GenS1{F0: false, F1: "a", F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v46
    _ = v47
    _ = v5
}
