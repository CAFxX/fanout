package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int32; F1 uint; F2 string }
type GenS4 struct { F0 string; F1 uint; F2 int64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int, p2 int32) uint64 {
    if false {
        var v7 uint = 255
        p0 = 2
        v8 := GenS1{F0: p1, F1: 255, F2: p0}
        _ = v7
        _ = v8
    }
    if (-2.25 < -2.25) {
        var v9 uint32 = 4294967295
        p0 += 42
        v10 := []int{0, (p1 << ((4294967295) & 63)), (-(p1)), (p1 / ((p1) | 1))}
        _ = v10
        _ = v9
    } else {
        var v11 uint = 1
        if true {
            fmt.Println(v11)
            for fi12 := 0; fi12 < 9; fi12++ {
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 uint8 = (p0 % (((p0 * p0)) | 1))
                    _ = fi13
                    _ = v14
                }
                v15 := GenS2{F0: 2147483647, F1: (2 >= 237)}
                v16 := GenS2{F0: 847, F1: false}
                _ = fi12
                _ = v15
                _ = v16
            }
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 969
}
func genFn1(p0 int16, p1 bool) uint16 {
    var v17 uint32 = 226
    var v18 bool = p1
    for fi19 := 0; fi19 < 7; fi19++ {
        switch "\n\t" {
        case "":
            var v20 int8 = -16
            _ = v20
        case "foo bar":
            var v21 int64 = 9223372036854775807
            _ = v21
        default:
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 float32 = (float32((1e10 / 0.1)) - 0.1)
                var v24 uint8 = uint8(v17)
                switch fi22 {
                case 1:
                    if false {
                        if (!false) {
                            v25 := GenS1{F0: fi22, F1: 0, F2: v24}
                            _ = v25
                        } else {
                            switch (fi19 << ((2) & 63)) {
                            case 4294967295:
                                var v26 int16 = 569
                                _ = v26
                            }
                        }
                        v24 = v24
                    } else {
                        v17 = (v17 / ((v17) | 1))
                        v27 := []int{(fi19 ^ 9223372036854775807), 128, (fi19 & 32768)}
                        _ = v27
                    }
                case 959:
                    var v28 uint8 = (^(v24))
                    _ = v28
                case 32767:
                    p1 = ((v18 && ("hello" == "a")) && v18)
                }
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        for fi29 := 0; fi29 < 16; fi29++ {
            v18 = (112 >= 856)
            var v30 float32 = -2.25
            switch fi29 {
            case -334:
                fmt.Println(p1)
            case 2147483647:
                var v31 uint64 = 494
                _ = v31
            default:
                var v32 int64 = 42
                _ = v32
            }
            _ = fi29
            _ = v30
        }
        _ = fi19
    }
    v17 = 2
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 442
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]int)
    var v3 int16 = -46
    v4 := make(map[uint32]bool)
    var v5 float64 = 1.5
    var v6 uint16 = 42
    v33 := GenS3{F0: 2147483647, F1: 1, F2: ("foo bar" + (("hello" + "foo bar") + ("hello" + "hello")))}
    v34 := GenS4{F0: ("a" + (((("hello" + (("\n\t" + (("hello" + "") + "foo bar")) + "foo bar")) + "a") + "a") + ("\n\t" + "\n\t"))), F1: 4294967295, F2: -700}
    if true {
        v1 += v1
    }
    if false {
        v3 = int16(v6)
        v35 := GenS3{F0: 1, F1: 4294967295, F2: "hello"}
        _ = v35
    }
    _ = len(v2)
    fmt.Println(v6, v3)
    v36 := GenS1{F0: -123, F1: 494, F2: 42}
    for fi37 := 0; fi37 < 18; fi37++ {
        v1 -= v1
        if (v1 < (v1 | v1)) {
            _ = len(v4)
            v38 := v2[(("a" + "foo bar") + ((("hello" + "\n\t") + "\n\t") + "foo bar"))]
            v6 = genFn1((v3 * 180), (!true))
            _ = v38
        }
        v1 = v1
        _ = fi37
    }
    if false {
        v2["foo bar"] = 136
    } else {
        v4[4294967295] = ((true || (!(("foo bar" + "") != ("a" + ("hello" + "hello"))))) && false)
        v39 := GenS4{F0: (("foo bar" + (("hello" + ("a" + "foo bar")) + (((("hello" + "a") + "foo bar") + "a") + "hello"))) + "foo bar"), F1: 1, F2: -9223372036854775808}
        for fi40 := 0; fi40 < 5; fi40++ {
            for fi41 := 0; fi41 < 20; fi41++ {
                v1 = (v1 % ((v1) | 1))
                _ = fi41
            }
            v6 = genFn1(v3, (0 >= 773))
            _ = fi40
        }
        _ = v39
    }
    var v42 int32 = -33
    for fi43 := 0; fi43 < 12; fi43++ {
        v6 = 65535
        v6 -= (^(v6))
        _ = fi43
    }
    v44 := GenS5{F0: 65535}
    var v45 int32 = v42
    v46 := v4[996]
    var v47 uint8 = 123
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
