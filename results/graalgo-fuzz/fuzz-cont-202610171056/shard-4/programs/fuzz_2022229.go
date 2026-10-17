package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 bool, p1 uint32, p2 int8) int {
    v6 := []int{386, 554, 4294967295, 47}
    fmt.Println(p1, p2, p0)
    v7 := GenS1{F0: 255, F1: (p2 * -58)}
    for fi8 := 0; fi8 < 12; fi8++ {
        if true {
            v9 := GenS1{F0: 899, F1: 0}
            _ = v9
        } else {
            p1 = p1
            v10 := GenS2{F0: 42, F1: p1}
            if len(v6) > 0 {
                v6[0] = (-(fi8))
            }
            _ = v10
        }
        fmt.Println(p0, p2, p2)
        _ = fi8
    }
    v11 := GenS2{F0: 0, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint64 = 2
    var v2 uint = 65535
    var v3 string = ("a" + "hello")
    v4 := []string{v3, v3}
    var v5 uint8 = uint8(v2)
    if (1 > 0) {
        var v12 float64 = 1e10
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v2 += (v2 % ((928) | 1))
            if len(v4) > 0 {
                v4[0] = (((v3 + "hello") + "foo bar") + (("\n\t" + v3) + ((v3 + "hello") + ((v3 + ("foo bar" + v3)) + (v3 + ("hello" + v3))))))
            }
        } else {
            fmt.Println(v12, v5)
            v3 = "a"
        }
        _ = v12
    } else {
        switch v3 {
        case "hello":
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 int = 1
                switch genFn0(false, 4294967295, 61) {
                case -1:
                    fmt.Println(fi13, v14)
                default:
                    var v15 uint32 = 98
                    _ = v15
                }
                v14 = 0
                _ = fi13
                _ = v14
            }
        }
        v5 += 2
        v16 := GenS1{F0: 837, F1: 100}
        _ = v16
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 uint32 = 2147483647
    if (((0 > v5) || (!((true && false) || false))) && (!(("" + (v3 + (v3 + ""))) >= (v3 + v3)))) {
        if false {
            var v18 int16 = 1000
            _ = v18
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 uint8 = v5
            v3 = v3
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if (false || true) {
            var v22 int16 = -32768
            _ = len(v4)
            _ = v22
        } else {
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 uint32 = v17
                v25 := GenS3{F0: 32767}
                _ = fi23
                _ = v24
                _ = v25
            }
            v2 = 1
        }
        if true {
            var v26 int16 = -32768
            var v27 float32 = ((3.14159 - (-2.25 / 1.5)) / 1e10)
            _ = v26
            _ = v27
        }
        _ = fi21
    }
    fmt.Println(v2, v3, v2)
    var v28 int64 = -469
    if false {
        if true {
            fmt.Println(v2)
            if (true || false) {
                for fi29 := 0; fi29 < 20; fi29++ {
                    var v30 uint = 2
                    for fi31 := 0; fi31 < 15; fi31++ {
                        var v32 uint8 = v5
                        _ = fi31
                        _ = v32
                    }
                    switch fi29 {
                    case 42:
                        switch fi29 {
                        case 0:
                            var v33 int = 477
                            _ = v33
                        default:
                            v30 += v2
                        }
                    case 127:
                        var v34 int64 = (v28 * (v28 & (v28 - (v28 % ((v28) | 1)))))
                        _ = v34
                    case -469:
                        v30 += 65535
                    default:
                        for fi35 := 0; fi35 < 18; fi35++ {
                            var v36 float64 = 0.0
                            var v37 float64 = v36
                            if true {
                                v5 += v5
                            }
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    }
                    _ = fi29
                    _ = v30
                }
                var v38 uint8 = (v5 / (((^(v5))) | 1))
                _ = v38
            } else {
                v28 = (v28 - v28)
                var v39 bool = (-1 > int16(v17))
                v17 += v17
                _ = v39
            }
            if (65535 < v2) {
                v1 += 920
            }
        } else {
            var v40 int = 541
            _ = v40
        }
    } else {
        var v41 int8 = 101
        v41 -= 1
        _ = v41
    }
    if ((false || ((v5 << ((v2) & 63)) != v5)) || true) {
        v4 = append(v4, (v3 + v3))
    }
    fmt.Println(v1, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
