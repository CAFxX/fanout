package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint32; F1 int; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string }
func genFn0(p0 int16) uint32 {
    if false {
        var v6 int = 42
        v6 += (v6 >> ((0) & 63))
        _ = v6
    }
    var v7 string = "a"
    var v8 uint64 = 2
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 42
    var v3 int32 = -2147483648
    var v4 float32 = 1.5
    v5 := make(map[string]uint32)
    if true {
        var v9 float32 = v4
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        switch (fi10 - (fi10 + (^(fi10)))) {
        case -9223372036854775808:
            fmt.Println(v3, v3, v2)
        case 311:
            for fi11 := 0; fi11 < 4; fi11++ {
                v2 += 38
                _ = fi11
            }
        }
        _ = fi10
    }
    v12 := GenS1{F0: 4294967295, F1: -1, F2: false}
    var v13 uint = 802
    var v14 int64 = 1
    var v15 uint = v13
    fmt.Println(v3, v14, v3)
    v16 := GenS2{F0: ""}
    fmt.Println(v4, v15)
    if true {
        v4 -= (v4 + 0.0)
        if (v3 > (v3 - v3)) {
            if false {
                v2 -= v2
                for fi17 := 0; fi17 < 7; fi17++ {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        delete(v5, "")
                        _ = len(v1)
                        v19 := GenS2{F0: "foo bar"}
                        var v20 int32 = (v3 << (((v15 | v13)) & 63))
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    var v21 int32 = 395
                    _ = fi17
                    _ = v21
                }
            }
        } else {
            switch "hello" {
            case "a":
                var v22 float64 = (1e10 / 1e10)
                _ = v22
            case "":
                v23 := GenS3{F0: 42, F1: 128}
                _ = v23
            default:
                if (-1 > 53) {
                    v24 := GenS4{F0: v15, F1: int32(v13)}
                    _ = len(v1)
                    v14 = v14
                    _ = v24
                } else {
                    v25 := GenS5{F0: 663, F1: -9223372036854775808, F2: 546}
                    fmt.Println(v15, v13, v14)
                    for fi26 := 0; fi26 < 16; fi26++ {
                        if true {
                            var v27 float32 = -2.25
                            if true {
                                fmt.Println(v3, v13, v15)
                                if false {
                                    v13 = 661
                                    for fi28 := 0; fi28 < 20; fi28++ {
                                        v29 := []int{fi26, fi28, (fi26 & fi28), fi28}
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                            }
                            _ = v27
                        }
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v4, v15, v4)
                        case "hello":
                            var v30 int8 = 42
                            _ = v30
                        default:
                            v13 = (v15 + 1)
                        }
                        _ = fi26
                    }
                    _ = v25
                }
            }
            var v31 uint8 = v2
            var v32 uint32 = 4294967295
            _ = v31
            _ = v32
        }
        switch "" {
        case "hello":
            if ((!(!false)) || true) {
                var v33 int = 1
                _ = v33
            }
        default:
            var v34 bool = ((v14 * -1) < (-(v14)))
            _ = v34
        }
    } else {
        switch 128 {
        case 32768:
            var v35 uint = v13
            _ = v35
        }
        v14 += (v14 & (v14 * (v14 >> ((18446744073709551615) & 63))))
    }
    if false {
        v36 := GenS6{F0: v3}
        _ = v36
    } else {
        v37 := []int{}
        if (true || ((!true) && false)) {
            if len(v37) > 0 {
                v37[0] = 32767
            }
        }
        _ = v37
    }
    switch 42 {
    case -2147483648:
        v38 := []int{}
        _ = v38
    case 891:
        v39 := GenS3{F0: v14, F1: (v2 / (((^(v2))) | 1))}
        _ = v39
    default:
        for fi40 := 0; fi40 < 4; fi40++ {
            switch ("foo bar" + ("a" + (("hello" + "a") + ""))) {
            case "hello":
                v15 -= (v15 & uint(v2))
            case "a":
                var v41 int8 = -65
                _ = v41
            case "":
                for fi42 := 0; fi42 < 14; fi42++ {
                    var v43 int = fi40
                    v44 := GenS7{F0: "a"}
                    switch -9223372036854775808 {
                    case 2147483647:
                        v45 := []int{32767, fi42, (-(fi42))}
                        _ = v45
                    }
                    v46 := []int{(fi40 + v43), fi40, (-(fi42))}
                    v14 -= v14
                    v43 -= v43
                    _ = fi42
                    _ = v43
                    _ = v44
                    _ = v46
                }
            default:
                v15 -= v13
            }
            _ = fi40
        }
    }
    v15 += v15
    if true {
        var v47 int64 = -9223372036854775808
        fmt.Println(v15)
        v48 := GenS2{F0: "foo bar"}
        _ = v47
        _ = v48
    }
    var v49 int = 9223372036854775807
    for fi50 := 0; fi50 < 2; fi50++ {
        var v51 uint16 = 0
        _ = fi50
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
