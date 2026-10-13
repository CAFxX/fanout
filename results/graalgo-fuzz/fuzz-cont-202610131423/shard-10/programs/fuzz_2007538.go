package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 string }
type GenS3 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint8; F1 string; F2 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool; F1 uint16; F2 int16 }
type GenS8 struct { F0 uint64; F1 int16; F2 uint16 }
func genFn0() int8 {
    v7 := GenS1{F0: "hello", F1: 65535}
    var v8 float64 = -2.25
    switch "hello" {
    case "":
        v9 := GenS1{F0: ("hello" + "a"), F1: 32767}
        _ = v9
    case "a":
        fmt.Println(v8)
    default:
        v10 := GenS2{F0: 2, F1: 1, F2: ""}
        _ = v10
    }
    switch "" {
    case "a":
        v11 := []int{0, 0, -1}
        _ = v11
    }
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 int16, p1 bool, p2 uint) uint32 {
    v12 := GenS3{F0: 763, F1: 1, F2: 42}
    switch 0 {
    case 2147483647:
        for fi13 := 0; fi13 < 6; fi13++ {
            p2 -= p2
            var v14 int8 = 42
            var v15 int16 = p0
            _ = fi13
            _ = v14
            _ = v15
        }
    case 46:
        fmt.Println(p2)
    }
    p0 += 0
    p1 = ("\n\t" <= "")
    if p1 {
        if p1 {
            var v16 float32 = (((1.5 * (float32(1.5) - 3.14159)) / 1e10) + 1.5)
            _ = v16
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int8 = 100
            v19 := GenS3{F0: 0, F1: 1, F2: -2147483648}
            var v20 uint16 = uint16(p2)
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        v21 := GenS2{F0: 592, F1: 1, F2: "a"}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 2147483647
}

func main() {
    var v1 bool = (6 == 111)
    var v2 uint8 = 2
    var v3 int = -697
    var v4 int16 = 32767
    v5 := make(map[uint32]string)
    var v6 int8 = 100
    var v22 int32 = -243
    v23 := v5[0]
    v22 = (-(v22))
    fmt.Println(v1, v1, v6)
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 int = fi24
            var v27 int8 = -13
            _ = fi25
            _ = v26
            _ = v27
        }
        var v28 float64 = -2.25
        var v29 float32 = -2.25
        _ = fi24
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 string = ("a" + v23)
        v32 := GenS4{F0: 65535, F1: 18446744073709551615}
        fmt.Println(v1)
        _ = fi30
        _ = v31
        _ = v32
    }
    v1 = (v23 < v23)
    fmt.Println(v23)
    if (2147483647 == genFn1((v4 | v4), ((1e10 != (((1.5 - -2.25) / (1e10 + (3.14159 - 0.1))) - 0.0)) || v1), 862)) {
        switch (v3 >> ((4294967295) & 63)) {
        case -9223372036854775808:
            var v33 int = 346
            _ = v33
        }
        v34 := GenS4{F0: 65535, F1: 65535}
        _ = v34
    } else {
        v3 = (v3 >> ((2) & 63))
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        if true {
            v23 = v23
            v4 = v4
            if (924 != 0) {
                v4 = v4
                v36 := GenS4{F0: 2, F1: 4294967295}
                v22 = (v22 << ((951) & 63))
                _ = v36
            } else {
                for fi37 := 0; fi37 < 3; fi37++ {
                    v5[65535] = v23
                    v3 = fi37
                    var v38 uint32 = 419
                    if v1 {
                        if (!false) {
                            switch (v23 + "hello") {
                            case "\n\t":
                                v39 := GenS5{F0: 1, F1: (v23 + (v23 + (v23 + (v23 + ((v23 + ((v23 + (("hello" + v23) + ((("" + (("" + ("foo bar" + "hello")) + (v23 + v23))) + v23) + ((v23 + "\n\t") + ((v23 + "foo bar") + (v23 + v23)))))) + ("\n\t" + v23))) + "foo bar"))))), F2: int16(v38)}
                                _ = v39
                            case "hello":
                                _ = len(v5)
                            case "":
                                var v40 int16 = v4
                                _ = v40
                            default:
                                _ = len(v5)
                                if (v1 || (false && v1)) {
                                    v41 := []int{}
                                    var v42 int32 = (v22 | v22)
                                    _ = v41
                                    _ = v42
                                }
                                fmt.Println(v6)
                                switch v23 {
                                case "foo bar":
                                    if ((v38 | v38) < v38) {
                                        v43 := []int{}
                                        _ = v43
                                    }
                                    v38 -= 724
                                    var v44 uint64 = uint64(v38)
                                    _ = v44
                                case "":
                                    fmt.Println(v2)
                                default:
                                    v23 = v23
                                }
                                if (!(v1 || v1)) {
                                    switch (fi37 >> ((1) & 63)) {
                                    case -2147483648:
                                        v45 := GenS6{F0: 18446744073709551615}
                                        _ = v45
                                    case 498:
                                        v46 := []int{923, fi35}
                                        _ = v46
                                    case 42:
                                        fmt.Println(v3, v6)
                                    }
                                    fmt.Println(v3, fi37)
                                    fmt.Println(v6)
                                }
                            }
                            v47 := []int{fi37}
                            _ = v47
                        } else {
                            v48 := GenS7{F0: v1, F1: 822, F2: -1}
                            v49 := []int{fi35}
                            v50 := GenS8{F0: 2, F1: (v4 >> ((18446744073709551615) & 63)), F2: 2}
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                    }
                    fmt.Println(fi37, v1, v23)
                    _ = fi37
                    _ = v38
                }
                fmt.Println(v23)
                var v51 uint64 = 42
                _ = v51
            }
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
