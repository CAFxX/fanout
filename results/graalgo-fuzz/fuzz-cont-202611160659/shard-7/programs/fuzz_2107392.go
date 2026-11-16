package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int; F2 bool }
func genFn0(p0 int64, p1 int, p2 int) uint8 {
    fmt.Println(p0, p1, p2)
    v7 := GenS1{F0: 128, F1: (true && true), F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 128
}
func genFn1(p0 int, p1 int32) int8 {
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS2{F0: 1, F1: 2}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        if (false || ((true || (!true)) || (!(2 < 42)))) {
            fmt.Println(p1)
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            if true {
                var v12 float64 = ((0.1 * 1e10) + (-2.25 + 3.14159))
                _ = v12
            }
            _ = fi11
        }
        for fi13 := 0; fi13 < 5; fi13++ {
            switch "\n\t" {
            case "foo bar":
                v14 := GenS3{F0: 2, F1: 1, F2: (1.5 != 0.1)}
                _ = v14
            default:
                p0 = fi13
                fmt.Println(fi10, fi10, p1)
                v15 := []int{42}
                _ = v15
            }
            switch (^(fi13)) {
            case 274:
                fmt.Println(fi10)
            default:
                v16 := []int{}
                _ = v16
            }
            _ = fi13
        }
        _ = fi10
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        p1 -= -246
        if true {
            for fi18 := 0; fi18 < 1; fi18++ {
                p0 = 9223372036854775807
                var v19 string = "hello"
                var v20 int32 = p1
                _ = fi18
                _ = v19
                _ = v20
            }
            fmt.Println(p1, p1, p1)
        }
        _ = fi17
    }
    if false {
        for fi21 := 0; fi21 < 8; fi21++ {
            p0 -= (^(fi21))
            _ = fi21
        }
        var v22 string = "\n\t"
        _ = v22
    }
    var v23 uint64 = 42
    _ = p0
    _ = p1
    _ = v23
    return -10
}

func main() {
    var v1 uint32 = 800
    var v2 bool = ((-1 == -1) && true)
    v3 := make(map[int]string)
    var v4 int = 1
    var v5 float32 = 0.0
    var v6 uint16 = 255
    if true {
        var v24 uint8 = uint8(v6)
        v24 = genFn0(115, -145, (-(v4)))
        v25 := []int{(v4 / (((v4 * v4)) | 1))}
        _ = v24
        _ = v25
    } else {
        if true {
            v2 = v2
            if (v2 && v2) {
                if ("\n\t" > "hello") {
                    v26 := []int{}
                    v1 -= (v1 >> ((42) & 63))
                    if len(v26) > 0 {
                        v26[0] = v4
                    }
                    _ = v26
                } else {
                    fmt.Println(v4, v6, v5)
                }
                if v2 {
                    var v27 int16 = 0
                    for fi28 := 0; fi28 < 16; fi28++ {
                        for fi29 := 0; fi29 < 8; fi29++ {
                            var v30 int32 = -458
                            var v31 int32 = v30
                            switch "a" {
                            case "\n\t":
                                var v32 bool = true
                                _ = v32
                            default:
                                if (false || false) {
                                    switch -722 {
                                    case -1:
                                        switch (fi28 % ((v4) | 1)) {
                                        case -33:
                                            v33 := GenS3{F0: 145, F1: 9223372036854775807, F2: (-2.25 != float64(float32(-2.25)))}
                                            _ = v33
                                        case -553:
                                            v4 = (v4 % ((fi29) | 1))
                                        }
                                    }
                                    v34 := []int{fi29, v4}
                                    v27 += v27
                                    var v35 int8 = 42
                                    _ = v34
                                    _ = v35
                                }
                            }
                            var v36 int = fi29
                            var v37 int16 = v27
                            _ = fi29
                            _ = v30
                            _ = v31
                            _ = v36
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 1; fi38++ {
                            v5 = (((v5 + float32(-2.25)) / 0.1) - v5)
                            _ = fi38
                        }
                        _ = fi28
                    }
                    _ = v27
                }
            }
            if v2 {
                if (v6 == v6) {
                    if (121 > -2147483648) {
                        v4 -= 694
                        var v39 uint16 = 1
                        _ = v39
                    } else {
                        v1 = (v1 << ((65535) & 63))
                        var v40 float64 = (0.1 * 0.1)
                        _ = v40
                    }
                }
                v41 := GenS2{F0: (v1 - v1), F1: 65535}
                var v42 uint = 0
                _ = v41
                _ = v42
            }
        }
        fmt.Println(v2, v4)
    }
    v6 = v6
    var v43 int8 = 1
    _ = len(v3)
    switch v4 {
    case 128:
        for fi44 := 0; fi44 < 9; fi44++ {
            for fi45 := 0; fi45 < 3; fi45++ {
                v46 := []int{fi45, -9223372036854775808, v4, -9223372036854775808}
                v4 += (fi45 ^ (fi44 * v4))
                v1 = (v1 / (((v1 / ((v1) | 1))) | 1))
                _ = fi45
                _ = v46
            }
            _ = fi44
        }
    case 9223372036854775807:
        _ = len(v3)
    case 32767:
        var v47 int32 = -528
        _ = v47
    }
    switch v4 {
    case -2147483648:
        _ = len(v3)
    case 4294967295:
        for fi48 := 0; fi48 < 15; fi48++ {
            for fi49 := 0; fi49 < 15; fi49++ {
                fmt.Println(v1, v4)
                _ = fi49
            }
            _ = fi48
        }
    case -1:
        var v50 string = ((((("foo bar" + ((("hello" + (((("hello" + "hello") + (("hello" + ("hello" + "foo bar")) + "")) + "") + ("" + ("a" + ("hello" + "hello"))))) + "a") + (("a" + (("" + ("foo bar" + "a")) + "")) + ""))) + (((("foo bar" + "") + "foo bar") + "foo bar") + "")) + (("a" + ("" + "\n\t")) + "a")) + "hello") + "a")
        _ = v50
    }
    v51 := []int{(v4 >> ((42) & 63)), -9223372036854775808, v4}
    v6 += v6
    for fi52 := 0; fi52 < 17; fi52++ {
        v53 := v3[(^(v4))]
        delete(v3, fi52)
        var v54 int16 = -1
        _ = fi52
        _ = v53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v51
    _ = v6
}
