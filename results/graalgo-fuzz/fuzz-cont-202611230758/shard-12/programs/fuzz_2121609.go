package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 bool; F2 int32 }
type GenS6 struct { F0 string; F1 uint16; F2 string }
type GenS7 struct { F0 int64; F1 int8 }
func genFn0(p0 bool, p1 int, p2 uint16) int8 {
    var v6 int32 = -2147483648
    v7 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 string, p2 int) bool {
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 bool = false
        var v10 int = 127
        switch "\n\t" {
        case "foo bar":
            v11 := GenS2{F0: (^(p0)), F1: 4294967295, F2: 1}
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    switch (p2 % ((p2) | 1)) {
    case 42:
        p1 = ""
    case -675:
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int64 = 1
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 int8 = 0
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    default:
        v16 := []int{p2, (^(p2)), (p2 >> ((p0) & 63)), (p2 & (p2 - 32767))}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 int64 = 42
    var v2 uint = 2
    v3 := make(map[uint32]uint32)
    var v4 float64 = (3.14159 - (0.1 * 0.0))
    var v5 int8 = 1
    if ((^(v5)) < (v5 >> (((v2 >> ((v2) & 63))) & 63))) {
        _ = len(v3)
        v17 := GenS3{F0: 1}
        var v18 uint8 = 33
        _ = v17
        _ = v18
    } else {
        v2 = v2
        v2 -= (v2 ^ (^(v2)))
    }
    v4 += float64(0.0)
    v5 -= (v5 | -62)
    var v19 int64 = (v1 + v1)
    switch ("foo bar" + "hello") {
    case "foo bar":
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS4{F0: (("foo bar" + ("hello" + "\n\t")) + "a")}
            _ = fi20
            _ = v21
        }
    }
    var v22 uint = 255
    switch -221 {
    case 0:
        v22 = (v22 << ((v2) & 63))
    case 127:
        fmt.Println(v19, v22)
    case 32767:
        var v23 int8 = 127
        _ = v23
    }
    fmt.Println(v22, v1, v5)
    for fi24 := 0; fi24 < 17; fi24++ {
        if false {
            if true {
                if (((-(fi24)) != int(v1)) && (!true)) {
                    for fi25 := 0; fi25 < 19; fi25++ {
                        if (-335 > v19) {
                            var v26 uint = v22
                            if (!false) {
                                switch "foo bar" {
                                case "foo bar":
                                    if false {
                                        var v27 uint = 279
                                        if true {
                                            v28 := GenS5{F0: 233, F1: false, F2: 1}
                                            var v29 uint64 = 0
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 int8 = -30
                                            _ = v30
                                        }
                                        _ = v27
                                    }
                                case "a":
                                    fmt.Println(v4, v4, fi24)
                                }
                            } else {
                                var v31 int32 = -2147483648
                                _ = v31
                            }
                            _ = v26
                        } else {
                            fmt.Println(fi24, v1)
                            var v32 int = 127
                            switch (fi25 & (v32 | fi24)) {
                            case 1:
                                v2 += v2
                            case -1:
                                v33 := GenS6{F0: "", F1: 42, F2: (("a" + (("hello" + ("hello" + "foo bar")) + "\n\t")) + ("\n\t" + "hello"))}
                                _ = v33
                            case 32767:
                                v2 = v22
                            default:
                                v34 := []int{0, (fi24 - v32), (fi24 - -825)}
                                _ = v34
                            }
                            _ = v32
                        }
                        if (true && (4294967295 <= 0)) {
                            var v35 float32 = 1e10
                            v36 := []int{fi24, (fi25 * fi25)}
                            _ = v35
                            _ = v36
                        } else {
                            v37 := GenS7{F0: 0, F1: (v5 | genFn0((((!false) && false) || false), fi25, 1))}
                            fmt.Println(v2)
                            v5 += (v5 * (^(v5)))
                            _ = v37
                        }
                        _ = fi25
                    }
                    v38 := GenS2{F0: 255, F1: 854, F2: 945}
                    _ = v38
                }
            }
        } else {
            var v39 int = fi24
            var v40 int8 = 42
            _ = v39
            _ = v40
        }
        var v41 bool = (2 == 127)
        _ = fi24
        _ = v41
    }
    fmt.Println(v19, v19)
    switch 32767 {
    case 32768:
        var v42 uint = v2
        _ = v42
    case 44:
        fmt.Println(v1)
    }
    v43 := []int{0}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
