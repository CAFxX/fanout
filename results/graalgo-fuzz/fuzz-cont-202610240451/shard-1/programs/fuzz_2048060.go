package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint64; F1 int16; F2 int8 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 bool }
func genFn0() int32 {
    var v4 int8 = -128
    v5 := []int{42}
    fmt.Println(v4)
    for fi6 := 0; fi6 < 13; fi6++ {
        if true {
            var v7 uint16 = 42
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 uint8 = 42
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        fmt.Println(fi6, v4)
        _ = fi6
    }
    v4 += v4
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int16) uint8 {
    v10 := GenS1{F0: 4294967295, F1: 729, F2: 42}
    if (!(false && true)) {
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := []int{(fi11 * fi11), fi11, -1}
            p0 += p0
            if (!false) {
                p0 = (p0 / ((p0) | 1))
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := GenS2{F0: 42, F1: -104}
                    p0 -= p0
                    var v15 float32 = (-2.25 * ((1e10 / ((-2.25 - (1.5 - ((float32(1.5) - 1.5) + 0.0))) - 0.1)) + 0.0))
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := GenS3{F0: 435, F1: 1, F2: 65}
                _ = v16
            }
            _ = fi11
            _ = v12
        }
    } else {
        fmt.Println(p0)
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS4{F0: ((fi17 | fi17) < (fi17 * 32768)), F1: 4294967295}
        v19 := []int{(^(fi17)), (fi17 % ((32767) | 1)), 4294967295, 128}
        if false {
            v20 := GenS3{F0: 1, F1: p0, F2: int8(p0)}
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 int16 = (p0 ^ 32767)
                _ = fi21
                _ = v22
            }
            p0 = (p0 * 0)
            _ = v20
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = v10
    return 36
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    if false {
        var v23 int64 = -1
        v3[(v1 % (((v1 + (v1 >> ((315) & 63)))) | 1))] = (true && (true || false))
        _ = v23
    }
    var v24 float32 = -2.25
    if true {
        if true {
            v25 := GenS5{F0: (v1 + (v1 | 628)), F1: 18}
            fmt.Println(v1, v1)
            var v26 uint = 270
            _ = v25
            _ = v26
        } else {
            v27 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: 0}
            _ = v27
        }
        v28 := v3[4294967295]
        _ = v28
    } else {
        var v29 int16 = 0
        var v30 uint8 = genFn1(v29)
        _ = v29
        _ = v30
    }
    v24 = v24
    for fi31 := 0; fi31 < 18; fi31++ {
        v32 := GenS6{F0: (("\n\t" + "a") + "\n\t")}
        for fi33 := 0; fi33 < 17; fi33++ {
            switch (fi33 / ((42) | 1)) {
            case 1:
                v34 := GenS2{F0: 42, F1: -97}
                _ = v34
            case -726:
                var v35 uint64 = uint64(v1)
                _ = v35
            default:
                var v36 int = 128
                _ = v36
            }
            var v37 float64 = 1e10
            fmt.Println(v24)
            _ = fi33
            _ = v37
        }
        _ = fi31
        _ = v32
    }
    if true {
        v38 := GenS5{F0: 2, F1: -47}
        fmt.Println(v24)
        for fi39 := 0; fi39 < 6; fi39++ {
            switch fi39 {
            case -2147483648:
                for fi40 := 0; fi40 < 3; fi40++ {
                    if ((fi40 << ((387) & 63)) < (fi39 & 856)) {
                        v1 = (v1 / ((v1) | 1))
                        v41 := v3[(^(v1))]
                        v42 := []int{fi39}
                        _ = v41
                        _ = v42
                    }
                    if true {
                        v1 = 42
                    }
                    _ = len(v2)
                    if ("foo bar" != ("hello" + "\n\t")) {
                        var v43 int8 = 110
                        var v44 int8 = (v43 / (((v43 - (v43 + (v43 * v43)))) | 1))
                        fmt.Println(v24)
                        _ = v43
                        _ = v44
                    } else {
                        fmt.Println(fi40)
                        v45 := GenS7{F0: 4294967295, F1: "hello"}
                        _ = v45
                    }
                    var v46 int16 = 32767
                    _ = fi40
                    _ = v46
                }
            }
            _ = fi39
        }
        _ = v38
    }
    var v47 uint = 2
    v1 = v1
    if true {
        v47 -= (v47 >> ((v47) & 63))
        for fi48 := 0; fi48 < 4; fi48++ {
            var v49 int = fi48
            _ = fi48
            _ = v49
        }
    } else {
        if false {
            fmt.Println(v24, v24)
        }
        v50 := GenS8{F0: true}
        _ = v50
    }
    var v51 uint64 = 18446744073709551615
    v52 := GenS7{F0: (v47 + 2), F1: ("hello" + "hello")}
    switch 1 {
    case 32768:
        if false {
            v24 = ((((v24 + 0.1) - 0.0) * 0.0) / 1.5)
            v3[2147483648] = true
        } else {
            var v53 int8 = 101
            v54 := v3[(v1 * (v1 + (^(v1))))]
            _ = v53
            _ = v54
        }
    default:
        fmt.Println(v1, v1, v1)
    }
    for fi55 := 0; fi55 < 18; fi55++ {
        var v56 uint8 = 128
        _ = fi55
        _ = v56
    }
    fmt.Println(v24, v47)
    v47 -= (v47 ^ v47)
    v57 := GenS8{F0: (!true)}
    var v58 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    fmt.Println("v51:", v51)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v47
    _ = v51
    _ = v52
    _ = v57
    _ = v58
}
