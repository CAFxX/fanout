package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = (42 == 330)
    var v3 uint = 18446744073709551615
    fmt.Println(v2, v2)
    v3 = v3
    fmt.Println(v2)
    var v4 bool = (42 != 42)
    if true {
        switch ("foo bar" + "") {
        case "hello":
            if v4 {
                if v2 {
                    if (!(772 <= 128)) {
                        v4 = v2
                        var v5 int16 = 1000
                        var v6 float64 = float64((float32(((1.5 + (0.0 * 1e10)) + 1.5)) - 0.1))
                        _ = v5
                        _ = v6
                    } else {
                        if v2 {
                            v7 := GenS1{F0: 128, F1: 488}
                            var v8 uint16 = 42
                            v9 := GenS2{F0: 127, F1: 2}
                            _ = v7
                            _ = v8
                            _ = v9
                        } else {
                            var v10 uint16 = 349
                            for fi11 := 0; fi11 < 13; fi11++ {
                                v12 := GenS3{F0: -583, F1: ("foo bar" + "\n\t")}
                                _ = fi11
                                _ = v12
                            }
                            _ = v10
                        }
                        v4 = true
                    }
                    v3 = (^(v3))
                } else {
                    var v13 int32 = 2147483647
                    _ = v13
                }
            }
        case "a":
            var v14 int64 = -9223372036854775808
            _ = v14
        }
        switch 433 {
        case -1:
            v15 := []int{740, 9223372036854775807, -1}
            _ = v15
        case 32767:
            v2 = ((128 == 18446744073709551615) && ("\n\t" == (("foo bar" + "\n\t") + "")))
        case -2147483648:
            v16 := GenS2{F0: -45, F1: 1}
            _ = v16
        }
        var v17 uint32 = 0
        _ = v17
    } else {
        fmt.Println(v3, v2)
        var v18 int8 = -36
        var v19 bool = v4
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v2 = (true && ((^(v3)) > (v3 << ((v3) & 63))))
        v21 := GenS4{F0: 127}
        _ = fi20
        _ = v21
    }
    var v22 float64 = (float64(0.1) - (-2.25 * (1e10 + -2.25)))
    _ = len(v1)
    v2 = v2
    switch (("foo bar" + ("\n\t" + "foo bar")) + ("a" + ("\n\t" + "\n\t"))) {
    case "foo bar":
        switch ("a" + "") {
        case "a":
            if (!false) {
                v22 -= (float64(-2.25) - -2.25)
                switch 0 {
                case 184:
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v24 := GenS2{F0: 100, F1: 19}
                        if true {
                            v3 += 434
                            var v25 int8 = 127
                            _ = v25
                        }
                        _ = fi23
                        _ = v24
                    }
                case 601:
                    for fi26 := 0; fi26 < 11; fi26++ {
                        if (-731 < 42) {
                            v27 := GenS2{F0: 0, F1: 65535}
                            v28 := GenS4{F0: 0}
                            v29 := []int{fi26, -291}
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            for fi30 := 0; fi30 < 3; fi30++ {
                                v2 = ((^(fi30)) < (fi26 * fi30))
                                var v31 uint64 = 1
                                _ = fi30
                                _ = v31
                            }
                            v3 += v3
                        }
                        v2 = v4
                        _ = fi26
                    }
                }
                v2 = (!v4)
            } else {
                var v32 uint32 = 1
                switch "\n\t" {
                case "foo bar":
                    v33 := []int{128, -210}
                    _ = v33
                default:
                    if ((v2 || (!false)) || (!true)) {
                        for fi34 := 0; fi34 < 7; fi34++ {
                            _ = len(v1)
                            _ = fi34
                        }
                    }
                }
                _ = v32
            }
        }
    }
    fmt.Println(v2, v2)
    switch (("foo bar" + "a") + ("\n\t" + "\n\t")) {
    case "foo bar":
        v3 = (^(v3))
    }
    var v35 int32 = 0
    switch 42 {
    case 127:
        for fi36 := 0; fi36 < 15; fi36++ {
            if (!v2) {
                if (!v2) {
                    fmt.Println(v22)
                    for fi37 := 0; fi37 < 10; fi37++ {
                        if (float64((float32(v22) - ((-2.25 + 3.14159) * 1e10))) < v22) {
                            var v38 int32 = v35
                            v38 = (v38 / ((v38) | 1))
                            _ = v38
                        } else {
                            var v39 int32 = (-(v35))
                            _ = v39
                        }
                        _ = fi37
                    }
                }
                var v40 string = "foo bar"
                v41 := GenS2{F0: int8(v35), F1: 284}
                _ = v40
                _ = v41
            }
            var v42 string = "foo bar"
            v43 := GenS4{F0: 131}
            _ = fi36
            _ = v42
            _ = v43
        }
    case -180:
        v44 := GenS2{F0: 1, F1: 42}
        _ = v44
    case -917:
        var v45 float32 = float32(1e10)
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v4
}
