package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint8; F2 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS8 struct { F0 int8; F1 uint64; F2 uint32 }

func main() {
    var v1 int = 677
    var v2 int = v1
    var v3 float32 = float32((1.5 / 3.14159))
    var v4 int = (^(v2))
    var v5 bool = false
    fmt.Println(v2)
    switch (v2 ^ (v1 ^ v1)) {
    case 2147483647:
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 bool = v5
            v3 = (v3 / (v3 * v3))
            _ = fi6
            _ = v7
        }
    case 4294967295:
        if v5 {
            fmt.Println(v3)
            v8 := []int{(-(v1))}
            var v9 int64 = -1
            _ = v8
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 3; fi10++ {
            fmt.Println(fi10, v3)
            for fi11 := 0; fi11 < 18; fi11++ {
                v3 = float32(0.0)
                _ = fi11
            }
            var v12 float32 = float32(0.1)
            _ = fi10
            _ = v12
        }
    }
    v5 = (false && (!((v5 || (-1 > 1000)) && ("hello" == "a"))))
    v2 = (-(v1))
    v5 = (2 < 2147483648)
    var v13 uint = 2
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := []int{(fi14 / ((v1) | 1)), -9223372036854775808, v1, v4}
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        var v16 int64 = 9223372036854775807
        _ = fi14
        _ = v15
        _ = v16
    }
    if ((278 != 772) && true) {
        if v5 {
            var v17 float32 = v3
            for fi18 := 0; fi18 < 2; fi18++ {
                v17 -= float32(((1e10 / -2.25) / float64(v17)))
                v19 := GenS1{F0: 966, F1: -32768}
                var v20 float64 = -2.25
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        } else {
            v21 := GenS2{F0: 42, F1: -426}
            if (1 <= 42) {
                var v22 int32 = int32(v13)
                v23 := []int{v1}
                _ = v22
                _ = v23
            }
            _ = v21
        }
        for fi24 := 0; fi24 < 4; fi24++ {
            v25 := GenS2{F0: 1, F1: -1}
            _ = fi24
            _ = v25
        }
    } else {
        if (!(v5 && ((1e10 * 0.1) > 0.0))) {
            switch ((("\n\t" + ("foo bar" + ("hello" + ""))) + "\n\t") + "\n\t") {
            case "hello":
                var v26 uint64 = 402
                _ = v26
            case "foo bar":
                v27 := []int{v2}
                _ = v27
            case "\n\t":
                switch "\n\t" {
                case "a":
                    var v28 int8 = -1
                    _ = v28
                case "foo bar":
                    v29 := []int{(^(v1)), v2, v4}
                    _ = v29
                default:
                    if ((("a" + "foo bar") == (("foo bar" + "foo bar") + ("hello" + (("" + "") + "\n\t")))) || (v5 || v5)) {
                        for fi30 := 0; fi30 < 10; fi30++ {
                            v31 := GenS3{F0: 127, F1: v1}
                            v1 = fi30
                            _ = fi30
                            _ = v31
                        }
                        var v32 uint64 = 2
                        v33 := []int{(-(v4))}
                        _ = v32
                        _ = v33
                    } else {
                        for fi34 := 0; fi34 < 4; fi34++ {
                            if false {
                                v35 := GenS4{F0: 131}
                                fmt.Println(v13)
                                v36 := GenS2{F0: 362, F1: 1}
                                _ = v35
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 1; fi37++ {
                                if (!true) {
                                    v38 := []int{}
                                    v39 := GenS5{F0: 306, F1: 42, F2: fi34}
                                    _ = v38
                                    _ = v39
                                } else {
                                    switch -2147483648 {
                                    case 187:
                                        v40 := []int{fi34, v4, (v1 * fi34), (^(v1))}
                                        _ = v40
                                    case 9223372036854775807:
                                        var v41 bool = (float64((v3 * v3)) != 1.5)
                                        _ = v41
                                    case 166:
                                        var v42 uint8 = 135
                                        _ = v42
                                    }
                                    v43 := []int{219, (v4 | 989), (fi37 & int(v13)), v1}
                                    v4 = (fi37 % ((-253) | 1))
                                    _ = v43
                                }
                                var v44 float64 = 3.14159
                                v45 := []int{127, 84, (-(fi37))}
                                _ = fi37
                                _ = v44
                                _ = v45
                            }
                            _ = fi34
                        }
                        if (v13 < v13) {
                            switch ((("" + "") + "a") + "hello") {
                            case "hello":
                                v5 = v5
                            case "foo bar":
                                v3 -= 1e10
                            default:
                                v2 += (v1 >> ((v13) & 63))
                            }
                            fmt.Println(v13, v5, v3)
                        }
                        v46 := GenS6{F0: 66}
                        _ = v46
                    }
                }
            }
            v47 := []int{v2}
            _ = v47
        } else {
            var v48 int8 = 42
            switch "hello" {
            case "hello":
                v49 := GenS7{F0: 532, F1: 1, F2: 0}
                _ = v49
            case "\n\t":
                var v50 int32 = 233
                _ = v50
            case "a":
                v48 = (v48 << ((v13) & 63))
            }
            fmt.Println(v5)
            _ = v48
        }
        fmt.Println(v1, v4)
    }
    for fi51 := 0; fi51 < 17; fi51++ {
        fmt.Println(v3)
        switch -1 {
        case 4294967295:
            var v52 uint8 = 193
            _ = v52
        case 32767:
            v53 := GenS3{F0: 42, F1: 1}
            _ = v53
        default:
            v13 = (v13 % (((^(v13))) | 1))
        }
        _ = fi51
    }
    if (-2.25 < (0.0 / 3.14159)) {
        var v54 int16 = 986
        _ = v54
    }
    v55 := GenS3{F0: 45, F1: (-(v1))}
    for fi56 := 0; fi56 < 5; fi56++ {
        v57 := []int{(v2 << ((v13) & 63)), fi56, (-(fi56)), 901}
        _ = fi56
        _ = v57
    }
    var v58 uint32 = 689
    v59 := GenS8{F0: -104, F1: 0, F2: 0}
    for fi60 := 0; fi60 < 2; fi60++ {
        v5 = v5
        var v61 uint8 = 127
        v62 := []int{}
        _ = fi60
        _ = v61
        _ = v62
    }
    v1 = (v1 - 42)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v58
    _ = v59
}
