package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 bool; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 uint8; F2 uint32 }
func genFn0() uint {
    if true {
        var v7 float32 = 0.1
        _ = v7
    } else {
        if (!((!(((1e10 - (3.14159 / 3.14159)) - 0.0) == 1e10)) || true)) {
            v8 := GenS1{F0: 42}
            _ = v8
        }
    }
    var v9 uint = 110
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 int64 = 42
            v13 := []int{(-(fi10))}
            v14 := GenS1{F0: 2}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        var v15 uint8 = 255
        _ = fi10
        _ = v15
    }
    _ = v9
    return 65535
}
func genFn1(p0 uint, p1 bool, p2 uint8) int16 {
    v16 := GenS1{F0: 460}
    p2 -= (p2 % ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -32768
}

func main() {
    var v1 int8 = -128
    v2 := []int{127, 32768, -734}
    v3 := []uint32{0, 65535, 2147483648}
    var v4 bool = ((0.0 - float64(-2.25)) == (0.0 * 0.0))
    var v5 int = -1
    v6 := make(map[uint32]uint32)
    v17 := GenS1{F0: 18446744073709551615}
    v18 := v6[216]
    switch ("foo bar" + "") {
    case "hello":
        if len(v3) > 0 {
            v3[0] = (v18 ^ v18)
        }
    case "\n\t":
        _ = len(v6)
    default:
        v4 = v4
    }
    fmt.Println(v1, v4)
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS1{F0: uint64(v18)}
        _ = len(v3)
        _ = len(v6)
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        switch ("hello" + "a") {
        case "hello":
            v1 += v1
        case "foo bar":
            for fi22 := 0; fi22 < 16; fi22++ {
                fmt.Println(fi21, fi21)
                var v23 uint8 = 255
                var v24 uint8 = 2
                _ = fi22
                _ = v23
                _ = v24
            }
        case "a":
            v18 = 2
        default:
            for fi25 := 0; fi25 < 6; fi25++ {
                switch (fi21 ^ fi25) {
                case 9223372036854775807:
                    switch (fi25 - v5) {
                    case 9223372036854775807:
                        var v26 int = (fi25 * (v5 << ((1) & 63)))
                        _ = v26
                    case 4294967295:
                        fmt.Println(v4)
                    case 822:
                        switch -505 {
                        case -300:
                            v18 = (v18 - v18)
                        case 4294967295:
                            _ = len(v6)
                        case 32768:
                            v27 := GenS2{F0: 2, F1: (!((1.5 * ((0.0 / 0.1) - 1e10)) != float64(1e10))), F2: 1}
                            _ = v27
                        }
                    }
                case 794:
                    v28 := GenS3{F0: fi21}
                    _ = v28
                case 0:
                    fmt.Println(v4)
                    var v29 float64 = 0.0
                    if true {
                        v30 := GenS2{F0: 138, F1: v4, F2: 0}
                        if (false && ((fi21 % ((128) | 1)) > (fi21 & fi21))) {
                            delete(v6, v18)
                            fmt.Println(v18, fi25)
                            var v31 float32 = -2.25
                            _ = v31
                        }
                        v32 := GenS4{F0: v1, F1: (fi21 >> ((2) & 63))}
                        fmt.Println(v1)
                        var v33 uint8 = 175
                        _ = v30
                        _ = v32
                        _ = v33
                    }
                    _ = v29
                default:
                    v34 := GenS3{F0: 32768}
                    _ = v34
                }
                _ = fi25
            }
        }
        _ = fi21
    }
    var v35 string = "hello"
    var v36 float32 = -2.25
    fmt.Println(v4)
    v37 := GenS1{F0: 0}
    v5 += (v5 % (((v5 & (v5 ^ 128))) | 1))
    v5 += -2147483648
    var v38 string = (("" + (v35 + (v35 + ((v35 + v35) + "a")))) + v35)
    if v4 {
        var v39 int8 = 42
        for fi40 := 0; fi40 < 19; fi40++ {
            v41 := GenS4{F0: (v39 ^ (v1 + v39)), F1: fi40}
            v38 = (("foo bar" + v38) + "a")
            switch v5 {
            case -248:
                delete(v6, 4294967295)
            case 141:
                v36 += float32(0.0)
            case 32767:
                if ((v1 >> ((42) & 63)) < (v39 / ((v1) | 1))) {
                    for fi42 := 0; fi42 < 17; fi42++ {
                        fmt.Println(v18, v5, v39)
                        v43 := GenS3{F0: v5}
                        var v44 uint32 = 2
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    fmt.Println(fi40, v4, fi40)
                    _ = len(v2)
                }
            }
            _ = fi40
            _ = v41
        }
        v45 := GenS5{F0: v18}
        _ = v39
        _ = v45
    } else {
        v46 := GenS4{F0: -7, F1: v5}
        delete(v6, v18)
        for fi47 := 0; fi47 < 18; fi47++ {
            if (!v4) {
                _ = len(v6)
                v48 := GenS6{F0: (!v4), F1: 89, F2: 65535}
                _ = v48
            } else {
                var v49 int16 = 1000
                _ = v49
            }
            _ = fi47
        }
        _ = v46
    }
    v1 -= (v1 / ((-3) | 1))
    fmt.Println(v35, v35, v38)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
