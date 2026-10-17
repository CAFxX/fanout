package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 string; F1 uint8; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int8 }
func genFn0() uint32 {
    v7 := GenS1{F0: 0, F1: 600}
    var v8 int = 657
    v8 = 1
    if (true && false) {
        v8 = 9223372036854775807
    }
    switch v8 {
    case 32768:
        var v9 uint64 = 4294967295
        _ = v9
    case 32767:
        v10 := GenS2{F0: "hello", F1: 1, F2: -2147483648}
        _ = v10
    }
    _ = v7
    _ = v8
    return 207
}

func main() {
    v1 := []string{"\n\t", (("a" + ("" + ("\n\t" + "hello"))) + (("\n\t" + "foo bar") + "a")), "hello", "foo bar"}
    v2 := []int{373, 1, 42, 953}
    v3 := make(map[uint32]string)
    v4 := []string{"foo bar", "foo bar"}
    var v5 float32 = (1.5 * (0.1 * (-2.25 / ((0.0 / 1.5) - float32(0.1)))))
    var v6 int = 4294967295
    v5 = 0.1
    if len(v4) > 0 {
        v4[0] = ""
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int64 = 749
        var v13 int64 = v12
        _ = fi11
        _ = v12
        _ = v13
    }
    v5 = v5
    _ = len(v3)
    v5 = 0.1
    var v14 uint16 = 42
    if ((!true) || true) {
        switch "a" {
        case "foo bar":
            var v15 uint64 = uint64(v14)
            _ = v15
        case "a":
            v16 := v3[genFn0()]
            _ = v16
        case "hello":
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS1{F0: 42, F1: -1}
                v6 += (fi17 % ((int(v14)) | 1))
                _ = fi17
                _ = v18
            }
        }
        v6 -= (v6 / (((v6 << ((65535) & 63))) | 1))
        for fi19 := 0; fi19 < 16; fi19++ {
            if (865 == 0) {
                if (true || false) {
                    v20 := GenS3{F0: (fi19 % ((fi19) | 1))}
                    v21 := GenS1{F0: 42, F1: -734}
                    var v22 float64 = (0.1 / (1.5 + (float64(v5) + 1e10)))
                    _ = v20
                    _ = v21
                    _ = v22
                }
                if ((v5 * v5) < 0.0) {
                    delete(v3, genFn0())
                    for fi23 := 0; fi23 < 16; fi23++ {
                        var v24 uint32 = 65535
                        _ = fi23
                        _ = v24
                    }
                    if (((771 >= 18446744073709551615) && (1e10 == (v5 * ((v5 + (v5 + v5)) * (v5 / 3.14159))))) && false) {
                        var v25 uint64 = 837
                        _ = v25
                    }
                } else {
                    v6 -= 42
                    for fi26 := 0; fi26 < 3; fi26++ {
                        if ((^(v14)) != (^(v14))) {
                            if len(v4) > 0 {
                                v4[0] = "hello"
                            }
                        } else {
                            v5 += (v5 - (0.1 * (float32(float64(v5)) * float32(0.0))))
                            var v27 int32 = -802
                            var v28 int64 = -400
                            _ = v27
                            _ = v28
                        }
                        v29 := GenS4{F0: 1, F1: -95}
                        _ = fi26
                        _ = v29
                    }
                    var v30 float32 = (v5 * v5)
                    _ = v30
                }
            }
            var v31 float32 = v5
            v14 -= 255
            _ = fi19
            _ = v31
        }
    } else {
        v5 += v5
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 uint16 = v14
            v14 = (v33 % ((v14) | 1))
            _ = fi32
            _ = v33
        }
        fmt.Println(v6)
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        if false {
            switch v6 {
            case -1:
                _ = len(v3)
            case 128:
                v6 += v6
            }
        } else {
            if (0.0 != 0.1) {
                v3[1] = ((("" + "") + ("a" + "hello")) + "\n\t")
                var v35 int32 = 2147483647
                _ = v35
            }
            if (!true) {
                var v36 int64 = 1
                for fi37 := 0; fi37 < 10; fi37++ {
                    switch (("" + "\n\t") + "foo bar") {
                    case "foo bar":
                        v6 -= fi37
                    case "\n\t":
                        v3[2147483647] = "hello"
                    default:
                        v2 = append(v2, fi34)
                    }
                    _ = fi37
                }
                _ = v36
            }
            v14 -= 1
        }
        var v38 uint64 = 1
        v39 := GenS4{F0: uint8(v14), F1: 79}
        _ = fi34
        _ = v38
        _ = v39
    }
    switch -544 {
    case 128:
        v40 := GenS5{F0: 665}
        _ = v40
    case -445:
        if true {
            var v41 int32 = 237
            _ = v41
        }
    default:
        for fi42 := 0; fi42 < 18; fi42++ {
            for fi43 := 0; fi43 < 3; fi43++ {
                fmt.Println(v5, v14)
                if true {
                    v5 = (1e10 + 1e10)
                    v14 = (^(v14))
                }
                switch "foo bar" {
                case "":
                    v6 -= (fi43 & fi42)
                }
                _ = fi43
            }
            switch 720 {
            case 2147483647:
                v44 := GenS6{F0: 18446744073709551615, F1: 42}
                _ = v44
            case -2147483648:
                if false {
                    var v45 uint32 = 0
                    var v46 int32 = 880
                    for fi47 := 0; fi47 < 16; fi47++ {
                        var v48 float32 = v5
                        if len(v4) > 0 {
                            v4[0] = ("" + "")
                        }
                        _ = fi47
                        _ = v48
                    }
                    _ = v45
                    _ = v46
                }
            case -374:
                switch (fi42 % (((^(v6))) | 1)) {
                case 32768:
                    v49 := GenS1{F0: 584, F1: -1}
                    _ = v49
                }
            }
            _ = fi42
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
