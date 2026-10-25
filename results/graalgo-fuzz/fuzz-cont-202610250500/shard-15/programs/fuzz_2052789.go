package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint32) uint32 {
    switch "foo bar" {
    case "\n\t":
        for fi7 := 0; fi7 < 1; fi7++ {
            p0 = p0
            _ = fi7
        }
    case "":
        p0 = 4294967295
    case "a":
        for fi8 := 0; fi8 < 14; fi8++ {
            p0 = p0
            _ = fi8
        }
    default:
        v9 := []int{-9223372036854775808, 188}
        _ = v9
    }
    var v10 bool = (1 > 42)
    v11 := GenS1{F0: "foo bar"}
    _ = p0
    _ = v10
    _ = v11
    return p0
}

func main() {
    v1 := []uint32{446}
    var v2 float32 = 1.5
    var v3 uint32 = 65535
    v4 := make(map[string]bool)
    v5 := make(map[string]uint32)
    v6 := make(map[int]string)
    v12 := GenS1{F0: ((("\n\t" + "a") + "") + "foo bar")}
    var v13 uint = 604
    var v14 int8 = 42
    if (!(true && (42 >= -247))) {
        v14 = (-(v14))
        for fi15 := 0; fi15 < 12; fi15++ {
            if (true && true) {
                v3 += (v3 >> (((v13 - 0)) & 63))
                var v16 int64 = int64(fi15)
                if (v14 < v14) {
                    for fi17 := 0; fi17 < 10; fi17++ {
                        v3 = v3
                        _ = fi17
                    }
                    v2 -= ((1e10 - v2) - -2.25)
                    fmt.Println(fi15, v2)
                } else {
                    switch (^(fi15)) {
                    case 444:
                        for fi18 := 0; fi18 < 9; fi18++ {
                            v3 = genFn0(genFn0(v3))
                            _ = fi18
                        }
                    case 42:
                        if (255 == 42) {
                            fmt.Println(fi15)
                            v3 = 233
                            v14 -= (v14 % (((v14 / (((-(v14))) | 1))) | 1))
                        }
                    case 128:
                        if false {
                            _ = len(v4)
                        }
                    default:
                        delete(v6, (fi15 - (fi15 << ((v13) & 63))))
                    }
                    var v19 float32 = v2
                    _ = v19
                }
                _ = v16
            } else {
                if false {
                    switch 0 {
                    case 486:
                        for fi20 := 0; fi20 < 1; fi20++ {
                            v13 -= v13
                            delete(v6, (fi15 * 685))
                            var v21 uint64 = 4294967295
                            _ = fi20
                            _ = v21
                        }
                    }
                    var v22 uint64 = 282
                    v3 += (v3 << (((^(v13))) & 63))
                    _ = v22
                }
                switch "\n\t" {
                case "foo bar":
                    var v23 uint = (v13 + (v13 - v13))
                    _ = v23
                case "hello":
                    var v24 int16 = 0
                    _ = v24
                }
            }
            _ = fi15
        }
    } else {
        _ = len(v1)
        v25 := v4["\n\t"]
        _ = v25
    }
    fmt.Println(v2, v13)
    if len(v1) > 0 {
        v1[0] = 615
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        v27 := v4["hello"]
        switch 4294967295 {
        case -1:
            v13 = (^(v13))
        case 4294967295:
            var v28 uint = v13
            _ = v28
        default:
            v3 = (v3 / (((^(v3))) | 1))
        }
        if v27 {
            if len(v1) > 0 {
                v1[0] = v3
            }
            for fi29 := 0; fi29 < 3; fi29++ {
                v5["hello"] = v3
                if false {
                    if v27 {
                        v30 := GenS1{F0: "\n\t"}
                        _ = v30
                    }
                    v4["foo bar"] = (32767 < 65535)
                    v6[fi29] = "\n\t"
                } else {
                    switch ("hello" + "") {
                    case "\n\t":
                        fmt.Println(v2)
                    default:
                        var v31 uint32 = 0
                        _ = v31
                    }
                }
                _ = len(v1)
                _ = fi29
            }
        } else {
            v1 = append(v1, (v3 / ((v3) | 1)))
            v32 := GenS1{F0: "a"}
            v33 := GenS1{F0: (("a" + "\n\t") + "\n\t")}
            _ = v32
            _ = v33
        }
        _ = fi26
        _ = v27
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        v35 := GenS1{F0: (("a" + "a") + "\n\t")}
        var v36 uint32 = v3
        v14 = (^(v14))
        _ = fi34
        _ = v35
        _ = v36
    }
    if (true || false) {
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 int64 = 0
            v39 := GenS2{F0: 0, F1: 0}
            var v40 string = "foo bar"
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
        var v41 int8 = v14
        v13 += v13
        _ = v41
    } else {
        v42 := GenS3{F0: 128}
        v3 -= genFn0((^(v3)))
        v13 += (v13 * v13)
        _ = v42
    }
    var v43 int8 = (v14 ^ (v14 << ((v13) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
