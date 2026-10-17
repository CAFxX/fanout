package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint32; F1 int32; F2 int8 }
type GenS5 struct { F0 uint32; F1 int; F2 int8 }
type GenS6 struct { F0 int64; F1 int16; F2 uint }
func genFn0(p0 string, p1 int16, p2 bool) int64 {
    var v4 int32 = 0
    p0 = ""
    fmt.Println(p0, p1)
    var v5 uint64 = 18446744073709551615
    p2 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 uint8, p1 uint16) int {
    if (65535 != 4294967295) {
        var v6 int64 = genFn0(("" + "a"), -32768, false)
        switch "foo bar" {
        case "":
            var v7 int16 = 499
            _ = v7
        case "hello":
            v8 := GenS1{F0: 4294967295, F1: "", F2: 100}
            _ = v8
        case "foo bar":
            v9 := GenS1{F0: 42, F1: "hello", F2: 100}
            _ = v9
        }
        _ = v6
    } else {
        v10 := GenS2{F0: 566, F1: ((float64((((1.5 + (((1.5 * 1e10) / ((0.1 / 3.14159) + (1e10 * (-2.25 * 1e10)))) + (0.0 + float32(0.1)))) / 1.5) / -2.25)) + float64((0.0 / float32(float64(1e10))))) <= 0.0)}
        if ("a" <= ("" + "")) {
            if (false || (true || true)) {
                p1 = p1
                for fi11 := 0; fi11 < 16; fi11++ {
                    v12 := GenS1{F0: 2, F1: "hello", F2: 0}
                    _ = fi11
                    _ = v12
                }
                fmt.Println(p1)
            } else {
                if (!true) {
                    var v13 int16 = -165
                    switch 0 {
                    case -141:
                        var v14 uint32 = 780
                        _ = v14
                    case -1:
                        switch "foo bar" {
                        case "hello":
                            if true {
                                var v15 int8 = -117
                                _ = v15
                            }
                        case "\n\t":
                            switch ((("" + "foo bar") + "foo bar") + ((("hello" + "hello") + ("foo bar" + "hello")) + "foo bar")) {
                            case "foo bar":
                                if (genFn0("", (^(v13)), (p1 <= p1)) != 0) {
                                    fmt.Println(v13)
                                } else {
                                    p1 = 255
                                    v13 = 42
                                    var v16 int64 = -427
                                    _ = v16
                                }
                            case "hello":
                                v17 := []int{0, 32767, 0}
                                _ = v17
                            default:
                                var v18 int32 = -771
                                _ = v18
                            }
                        case "foo bar":
                            p0 += p0
                        default:
                            v19 := []int{9223372036854775807, 0, 730, -2147483648}
                            _ = v19
                        }
                    }
                    _ = v13
                } else {
                    var v20 int = -1
                    var v21 uint8 = p0
                    _ = v20
                    _ = v21
                }
                p0 += (p0 | 128)
            }
            switch ("a" + "hello") {
            case "a":
                p1 += p1
            case "hello":
                var v22 string = ("a" + "foo bar")
                _ = v22
            default:
                var v23 uint16 = p1
                _ = v23
            }
            v24 := GenS3{F0: -32768, F1: 0}
            _ = v24
        } else {
            if true {
                if ((-1 > 1) && (((-1 >= 32767) || false) && (!(!(!((true && false) && false)))))) {
                    var v25 uint32 = 2147483647
                    v25 = v25
                    _ = v25
                }
                v26 := GenS4{F0: 4294967295, F1: -1, F2: 1}
                p1 = (^(p1))
                _ = v26
            }
            fmt.Println(p0, p0)
            if (!(822 < 595)) {
                switch (("\n\t" + (("hello" + ("\n\t" + "")) + ("\n\t" + "hello"))) + "foo bar") {
                case "a":
                    p0 -= p0
                }
            }
        }
        var v27 int8 = 0
        _ = v10
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        if (!(false && (((42 >= 0) && (!true)) || true))) {
            var v29 uint8 = (p0 % ((0) | 1))
            _ = v29
        }
        _ = fi28
    }
    if (false && (!true)) {
        var v30 uint32 = 960
        _ = v30
    } else {
        v31 := []int{}
        v32 := GenS5{F0: uint32(p0), F1: 4294967295, F2: 100}
        var v33 string = ("foo bar" + "foo bar")
        _ = v31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 string = "foo bar"
        v36 := []int{}
        for fi37 := 0; fi37 < 11; fi37++ {
            if false {
                if ("hello" >= v35) {
                    v36 = append(v36, fi34)
                    var v38 int32 = -1
                    var v39 string = ("" + ("foo bar" + ("\n\t" + ((v35 + "hello") + ("" + v35)))))
                    v40 := GenS4{F0: 2147483647, F1: (v38 % ((v38) | 1)), F2: 42}
                    _ = v38
                    _ = v39
                    _ = v40
                }
                var v41 uint64 = 18446744073709551615
                var v42 int16 = -32768
                _ = v41
                _ = v42
            }
            switch v35 {
            case "":
                v35 = "foo bar"
            case "hello":
                var v43 string = v35
                _ = v43
            default:
                p1 = (p1 + 665)
            }
            fmt.Println(p1, v35)
            _ = fi37
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    _ = p0
    _ = p1
    return 32768
}

func main() {
    var v1 bool = false
    var v2 float64 = 0.0
    var v3 uint64 = 2
    var v44 uint32 = 362
    if false {
        v45 := GenS4{F0: v44, F1: 734, F2: -119}
        v46 := []int{}
        _ = v45
        _ = v46
    }
    switch "\n\t" {
    case "foo bar":
        for fi47 := 0; fi47 < 1; fi47++ {
            for fi48 := 0; fi48 < 14; fi48++ {
                v49 := GenS6{F0: genFn0("", 32767, (543 != -1)), F1: -283, F2: 2}
                var v50 uint64 = (v3 >> ((1) & 63))
                v51 := []int{1, (fi48 >> ((42) & 63))}
                _ = fi48
                _ = v49
                _ = v50
                _ = v51
            }
            if true {
                fmt.Println(fi47, v44, v44)
            }
            _ = fi47
        }
    }
    var v52 uint32 = v44
    var v53 uint16 = 625
    v52 += v44
    var v54 int64 = genFn0(("foo bar" + "\n\t"), 84, (1 == 878))
    v1 = v1
    switch "a" {
    case "a":
        v53 = v53
    default:
        fmt.Println(v53)
    }
    if false {
        if v1 {
            v55 := GenS6{F0: (-(v54)), F1: -410, F2: 0}
            _ = v55
        } else {
            for fi56 := 0; fi56 < 15; fi56++ {
                var v57 int8 = 0
                v54 = genFn0("hello", 1000, ((2 < 65535) && v1))
                _ = fi56
                _ = v57
            }
            var v58 string = "hello"
            v59 := []int{-1, -514}
            _ = v58
            _ = v59
        }
        var v60 uint16 = (v53 >> ((42) & 63))
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v52
    _ = v53
    _ = v54
}
