package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint8; F2 bool }
type GenS3 struct { F0 uint; F1 int8; F2 int16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int) uint {
    switch "hello" {
    case "a":
        if false {
            p0 += p0
            if false {
                v6 := GenS1{F0: ("\n\t" + ("a" + "a"))}
                p0 = (p0 - p0)
                switch p0 {
                case 1:
                    v7 := GenS1{F0: "hello"}
                    _ = v7
                case -2147483648:
                    switch (p0 << ((1) & 63)) {
                    case 9223372036854775807:
                        for fi8 := 0; fi8 < 6; fi8++ {
                            v9 := []int{p0, fi8, (fi8 ^ fi8), -303}
                            for fi10 := 0; fi10 < 3; fi10++ {
                                if true {
                                    var v11 int = (p0 ^ -677)
                                    _ = v11
                                } else {
                                    p0 = p0
                                    _ = len(v9)
                                }
                                p0 = 128
                                _ = fi10
                            }
                            p0 = (p0 << ((657) & 63))
                            _ = fi8
                            _ = v9
                        }
                    case 32768:
                        switch p0 {
                        case 992:
                            p0 = (p0 / (((p0 + (p0 ^ (p0 + 32767)))) | 1))
                        default:
                            switch "\n\t" {
                            case "":
                                fmt.Println(p0)
                            case "hello":
                                var v12 int16 = -32768
                                _ = v12
                            case "foo bar":
                                v13 := GenS1{F0: "foo bar"}
                                _ = v13
                            }
                        }
                    case -2147483648:
                        p0 += (^(p0))
                    default:
                        var v14 uint = 4294967295
                        _ = v14
                    }
                default:
                    v15 := GenS2{F0: 449, F1: 13, F2: true}
                    _ = v15
                }
                _ = v6
            } else {
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 uint64 = 18446744073709551615
                    _ = fi16
                    _ = v17
                }
                if false {
                    v18 := []int{p0, (p0 >> ((42) & 63))}
                    v19 := GenS3{F0: 0, F1: 1, F2: -1}
                    _ = v18
                    _ = v19
                } else {
                    var v20 uint = 687
                    fmt.Println(p0, p0)
                    for fi21 := 0; fi21 < 5; fi21++ {
                        v22 := GenS2{F0: 42, F1: 128, F2: true}
                        v23 := []int{p0, p0, fi21, p0}
                        var v24 string = ("\n\t" + (("" + "") + ""))
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = v20
                }
                for fi25 := 0; fi25 < 2; fi25++ {
                    p0 = fi25
                    _ = fi25
                }
            }
        }
    }
    var v26 uint8 = 144
    _ = p0
    _ = v26
    return 755
}

func main() {
    var v1 float64 = (float64((1e10 / 3.14159)) + 0.1)
    var v2 uint64 = 42
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    var v5 int8 = 42
    v1 -= (0.0 / (((v1 - (((1e10 * (v1 / 1.5)) * 0.0) * v1)) + v1) - v1))
    var v27 uint = 65535
    var v28 uint8 = 116
    v29 := []int{867, int(v2), 42}
    if (true && false) {
        for fi30 := 0; fi30 < 1; fi30++ {
            fmt.Println(v2)
            _ = fi30
        }
        v2 += (v2 | 4294967295)
    } else {
        v2 += v2
        if true {
            if (((!true) && true) && (((v28 & v28) <= (v28 + (v28 & 193))) && true)) {
                v31 := GenS3{F0: v27, F1: (v5 / ((v5) | 1)), F2: -1}
                var v32 int16 = -563
                switch "\n\t" {
                case "hello":
                    fmt.Println(v1, v28)
                case "foo bar":
                    fmt.Println(v5)
                }
                _ = v31
                _ = v32
            } else {
                v29 = append(v29, -1)
            }
            for fi33 := 0; fi33 < 13; fi33++ {
                var v34 int8 = -13
                v2 -= v2
                var v35 int32 = -2147483648
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        v4["hello"] = "a"
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        for fi37 := 0; fi37 < 19; fi37++ {
            v38 := GenS2{F0: 42, F1: (^(v28)), F2: false}
            fmt.Println(v5)
            _ = fi37
            _ = v38
        }
        var v39 uint32 = 2147483647
        _ = fi36
        _ = v39
    }
    var v40 int8 = v5
    var v41 float32 = (float32(-2.25) * 0.0)
    if (!false) {
        var v42 float32 = v41
        v43 := GenS4{F0: -9223372036854775808}
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
