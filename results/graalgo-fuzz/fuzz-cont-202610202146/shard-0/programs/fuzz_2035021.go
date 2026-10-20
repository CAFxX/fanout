package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int8; F2 int }
func genFn0(p0 uint64) int32 {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float64 = -2.25
        _ = fi7
        _ = v8
    }
    p0 -= p0
    _ = p0
    return 42
}

func main() {
    var v1 int32 = -706
    v2 := make(map[string]bool)
    var v3 int16 = 64
    var v4 float32 = (float32(-2.25) + ((float32(float64(0.0)) * 0.0) - (1.5 - float32(0.0))))
    var v5 float32 = (v4 + float32(0.0))
    v6 := make(map[string]bool)
    var v9 float32 = v5
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            if true {
                var v12 bool = (!true)
                _ = v12
            }
            var v13 int = -1
            fmt.Println(v1)
            _ = fi11
            _ = v13
        }
        _ = fi10
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        switch "\n\t" {
        case "a":
            v15 := GenS1{F0: "a"}
            _ = v15
        case "\n\t":
            var v16 int = (fi14 << ((587) & 63))
            _ = v16
        case "":
            fmt.Println(v3, v5)
        }
        switch (fi14 ^ fi14) {
        case 128:
            for fi17 := 0; fi17 < 2; fi17++ {
                switch "a" {
                case "":
                    _ = len(v2)
                    v3 -= 42
                case "a":
                    v9 += v9
                case "foo bar":
                    var v18 int = (fi17 >> ((18446744073709551615) & 63))
                    if (((v1 + (v1 << ((2) & 63))) != 208) || (true && ((-38 > genFn0(4294967295)) || false))) {
                        fmt.Println(v3)
                    }
                    _ = v18
                }
                v19 := GenS1{F0: "foo bar"}
                var v20 int16 = (v3 / ((v3) | 1))
                _ = fi17
                _ = v19
                _ = v20
            }
        case -797:
            var v21 uint8 = 2
            _ = v21
        default:
            for fi22 := 0; fi22 < 13; fi22++ {
                fmt.Println(v5, v4)
                v23 := []int{int(v1), fi22, fi14, (^(fi14))}
                _ = fi22
                _ = v23
            }
        }
        _ = fi14
    }
    var v24 bool = (1 > 0)
    v9 = float32((0.1 / -2.25))
    var v25 int = 42
    for fi26 := 0; fi26 < 6; fi26++ {
        if (0 != 607) {
            if (2 < 19) {
                for fi27 := 0; fi27 < 12; fi27++ {
                    var v28 bool = v24
                    v29 := GenS1{F0: ""}
                    var v30 uint32 = 0
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                if (184 == 798) {
                    var v31 string = "\n\t"
                    v31 = v31
                    switch (v31 + v31) {
                    case "hello":
                        if (("hello" + (v31 + v31)) <= v31) {
                            if (-128 >= -14) {
                                var v32 float32 = -2.25
                                v5 -= 3.14159
                                _ = v32
                            } else {
                                var v33 int64 = 0
                                for fi34 := 0; fi34 < 2; fi34++ {
                                    v35 := []int{}
                                    _ = fi34
                                    _ = v35
                                }
                                _ = v33
                            }
                            v36 := []int{fi26, 32767, -9223372036854775808, (-(fi26))}
                            for fi37 := 0; fi37 < 14; fi37++ {
                                var v38 uint8 = 234
                                _ = fi37
                                _ = v38
                            }
                            _ = v36
                        } else {
                            switch v25 {
                            case 1:
                                var v39 float32 = v4
                                _ = v39
                            case -954:
                                v40 := []int{v25, (^(fi26))}
                                _ = v40
                            default:
                                fmt.Println(fi26, v9, v24)
                            }
                        }
                    case "":
                        v41 := v6[("foo bar" + (v31 + "a"))]
                        _ = v41
                    default:
                        switch fi26 {
                        case -2147483648:
                            fmt.Println(v3)
                        default:
                            v6[(v31 + (v31 + "a"))] = true
                        }
                    }
                    _ = v31
                } else {
                    v3 += (v3 >> ((228) & 63))
                }
                for fi42 := 0; fi42 < 20; fi42++ {
                    v24 = (true && v24)
                    _ = fi42
                }
            } else {
                v43 := []int{fi26, fi26, v25}
                v44 := GenS1{F0: (("\n\t" + "a") + "foo bar")}
                v45 := GenS2{F0: v24, F1: 100, F2: fi26}
                _ = v43
                _ = v44
                _ = v45
            }
            switch "hello" {
            case "\n\t":
                v46 := GenS1{F0: "a"}
                _ = v46
            case "":
                delete(v2, "hello")
            case "foo bar":
                var v47 uint16 = 42
                _ = v47
            }
            _ = len(v2)
        }
        _ = fi26
    }
    v48 := GenS2{F0: false, F1: -1, F2: (v25 >> ((1) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v9
}
