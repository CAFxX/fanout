package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int64; F1 uint; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint, p1 uint) uint16 {
    var v6 uint = (p0 * (p1 - 948))
    var v7 bool = (!((32767 == 42) || (true && false)))
    v8 := GenS1{F0: 0, F1: 32767}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint = 4294967295
    var v2 int = 9223372036854775807
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    var v5 uint8 = 42
    var v9 int8 = 1
    var v10 uint64 = 4294967295
    switch ("" + "hello") {
    case "hello":
        switch (("a" + "a") + ("" + "\n\t")) {
        case "foo bar":
            v11 := GenS1{F0: -1, F1: int16(v1)}
            _ = v11
        case "a":
            v12 := GenS2{F0: (v2 | (-(v2))), F1: true}
            _ = v12
        case "\n\t":
            v9 += v9
        default:
            if (-2.25 >= ((((-2.25 + (3.14159 * (-2.25 + 1e10))) / 1.0) / -2.25) / 1.5)) {
                for fi13 := 0; fi13 < 20; fi13++ {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        if (true && true) {
                            v15 := GenS2{F0: (fi13 | fi14), F1: (958 < 1000)}
                            v16 := GenS1{F0: int32(fi13), F1: -750}
                            _ = v15
                            _ = v16
                        } else {
                            switch (fi14 - (v2 % (((fi13 << (((v1 / ((v1) | 1))) & 63))) | 1))) {
                            case 42:
                                v5 = (v5 << ((255) & 63))
                            default:
                                v17 := GenS3{F0: -997, F1: (v1 ^ (v1 * (v1 % ((v1) | 1)))), F2: (v9 << (((v1 % ((v1) | 1))) & 63))}
                                _ = v17
                            }
                            v10 = v10
                        }
                        v18 := GenS2{F0: 4294967295, F1: false}
                        v19 := GenS4{F0: int64(v10)}
                        _ = fi14
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS2{F0: fi13, F1: (v9 == (v9 / ((0) | 1)))}
                    _ = fi13
                    _ = v20
                }
            } else {
                var v21 float32 = 0.1
                v22 := []int{(v2 / ((v2) | 1)), (v2 - (v2 + v2))}
                _ = v21
                _ = v22
            }
        }
    case "":
        for fi23 := 0; fi23 < 19; fi23++ {
            for fi24 := 0; fi24 < 6; fi24++ {
                v9 -= v9
                v25 := []int{fi24, fi24}
                v5 -= uint8(v1)
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 10; fi26++ {
                switch fi23 {
                case 242:
                    var v27 uint16 = 620
                    v1 -= (v1 + v1)
                    v27 += genFn0((^(v1)), v1)
                    _ = v27
                default:
                    v5 -= (^(v5))
                }
                v28 := GenS1{F0: -2147483648, F1: -934}
                switch "hello" {
                case "a":
                    var v29 uint16 = genFn0(v1, (v1 - v1))
                    _ = v29
                case "hello":
                    v1 = v1
                case "":
                    v30 := GenS3{F0: 9223372036854775807, F1: v1, F2: v9}
                    _ = v30
                default:
                    var v31 uint8 = uint8(v10)
                    var v32 int64 = -417
                    _ = v31
                    _ = v32
                }
                _ = fi26
                _ = v28
            }
            switch "" {
            case "hello":
                if (false || true) {
                    v33 := []int{278}
                    _ = v33
                }
            default:
                switch "\n\t" {
                case "a":
                    v34 := GenS2{F0: (fi23 % ((int(v9)) | 1)), F1: false}
                    _ = v34
                case "":
                    var v35 int8 = v9
                    _ = v35
                default:
                    var v36 int16 = -489
                    _ = v36
                }
            }
            _ = fi23
        }
    default:
        if true {
            v37 := []int{v2, v2, v2, int(v5)}
            _ = v37
        }
    }
    v2 = (v2 | (v2 + v2))
    if false {
        var v38 uint16 = genFn0(v1, 589)
        _ = v38
    } else {
        var v39 uint8 = v5
        _ = v39
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        for fi41 := 0; fi41 < 11; fi41++ {
            v9 += v9
            _ = fi41
        }
        for fi42 := 0; fi42 < 19; fi42++ {
            v43 := GenS5{F0: "\n\t", F1: (("a" + (("a" + "\n\t") + ("\n\t" + "\n\t"))) + (("" + "foo bar") + "hello"))}
            fmt.Println(v1, v2, v5)
            if true {
                if true {
                    v44 := []int{fi40, v2}
                    _ = v44
                }
            }
            _ = fi42
            _ = v43
        }
        var v45 string = (((("hello" + (("a" + "a") + ("hello" + "foo bar"))) + "") + "") + "hello")
        _ = fi40
        _ = v45
    }
    v46 := GenS6{F0: (true || (2 < 4294967295)), F1: v5}
    if (-2.25 == 0.1) {
        var v47 bool = ((v1 >> (((^(v1))) & 63)) != v1)
        var v48 uint = (v1 / (((v1 * v1)) | 1))
        _ = v47
        _ = v48
    } else {
        if false {
            if false {
                v10 += 0
                var v49 int16 = 915
                _ = v49
            }
        } else {
            for fi50 := 0; fi50 < 1; fi50++ {
                if (42 >= (^(v10))) {
                    var v51 int = (v2 - 32768)
                    _ = v51
                }
                _ = fi50
            }
        }
        for fi52 := 0; fi52 < 18; fi52++ {
            switch (fi52 / (((v2 & v2)) | 1)) {
            case 955:
                v5 += v5
            case 42:
                v53 := GenS3{F0: 0, F1: v1, F2: v9}
                _ = v53
            case -644:
                var v54 float32 = -2.25
                _ = v54
            }
            var v55 uint8 = (v5 >> ((48) & 63))
            switch "a" {
            case "":
                v56 := []int{}
                _ = v56
            }
            _ = fi52
            _ = v55
        }
        var v57 uint8 = (^(v5))
        _ = v57
    }
    v58 := GenS5{F0: "foo bar", F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v58
    _ = v9
}
