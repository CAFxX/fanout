package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 bool = false
    var v2 float32 = (1.5 / 0.1)
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    var v5 uint32 = 1
    v6 := GenS1{F0: 11}
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint = 0
        _ = fi7
        _ = v8
    }
    v5 = (v5 / (((^(v5))) | 1))
    switch 1 {
    case 2147483647:
        var v9 int = 9223372036854775807
        _ = v9
    }
    v1 = ((!v1) && (3.14159 >= (1.5 / 1.0)))
    v10 := GenS1{F0: 65535}
    v2 -= 1.5
    for fi11 := 0; fi11 < 17; fi11++ {
        v1 = (((1e10 * v2) != (v2 + v2)) || (true || true))
        v3[32767] = "a"
        _ = fi11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v1 = (0.1 <= -2.25)
        switch (("hello" + "hello") + "foo bar") {
        case "\n\t":
            var v13 int16 = 1000
            _ = v13
        case "foo bar":
            for fi14 := 0; fi14 < 2; fi14++ {
                if (65535 <= 18446744073709551615) {
                    switch "hello" {
                    case "a":
                        var v15 uint64 = 999
                        _ = v15
                    case "foo bar":
                        if (fi14 > (fi12 ^ (^(fi14)))) {
                            switch -103 {
                            case 246:
                                switch fi12 {
                                case 4294967295:
                                    v16 := GenS2{F0: 2147483647, F1: (v5 * v5), F2: v5}
                                    _ = v16
                                case 0:
                                    var v17 int64 = 0
                                    _ = v17
                                case 42:
                                    v18 := []int{}
                                    _ = v18
                                default:
                                    if (101 == 1) {
                                        v19 := GenS2{F0: 0, F1: v5, F2: 42}
                                        _ = v19
                                    }
                                    var v20 int16 = 42
                                    _ = v20
                                }
                            default:
                                if v1 {
                                    v5 -= v5
                                }
                            }
                            var v21 string = ""
                            _ = v21
                        }
                    }
                    var v22 string = "\n\t"
                    v2 = (((v2 + ((3.14159 - v2) * v2)) + v2) + v2)
                    _ = v22
                } else {
                    v2 = 1e10
                    var v23 uint64 = 0
                    v24 := GenS2{F0: 502, F1: (v5 ^ v5), F2: v5}
                    _ = v23
                    _ = v24
                }
                _ = fi14
            }
        case "":
            v25 := GenS2{F0: 559, F1: (v5 ^ v5), F2: v5}
            _ = v25
        default:
            for fi26 := 0; fi26 < 8; fi26++ {
                if (!((!v1) && true)) {
                    if true {
                        fmt.Println(v5, v5)
                        if (77 >= 0) {
                            var v27 int64 = -500
                            switch (fi26 * (-(fi26))) {
                            case 42:
                                v28 := []int{(^(fi12)), fi26, fi12}
                                _ = v28
                            }
                            v3[32767] = "foo bar"
                            _ = v27
                        } else {
                            v29 := GenS2{F0: -1, F1: 291, F2: (v5 >> ((42) & 63))}
                            switch (("\n\t" + ("a" + "hello")) + ((((("\n\t" + "\n\t") + ("foo bar" + "")) + ("hello" + "hello")) + ("\n\t" + "\n\t")) + "\n\t")) {
                            case "\n\t":
                                var v30 float64 = (0.1 * 1e10)
                                _ = v30
                            default:
                                var v31 int64 = -9223372036854775808
                                _ = v31
                            }
                            if v1 {
                                v1 = (((684 >= -1) && (v5 >= v5)) && (!v1))
                                v3[-1] = "\n\t"
                            } else {
                                v32 := GenS3{F0: 128, F1: 698, F2: 32767}
                                _ = v32
                            }
                            _ = v29
                        }
                        var v33 uint8 = 255
                        v5 = 2147483648
                        _ = v33
                    }
                    if v1 {
                        v1 = v1
                        var v34 string = "a"
                        v1 = (v1 || (42 != 716))
                        _ = v34
                    }
                } else {
                    v2 = (v2 + 0.0)
                }
                v5 = 65535
                _ = fi26
            }
        }
        var v35 bool = ((1 > -983) && (!v1))
        _ = fi12
        _ = v35
    }
    if (869 >= 2) {
        if v1 {
            for fi36 := 0; fi36 < 8; fi36++ {
                var v37 int8 = -1
                for fi38 := 0; fi38 < 11; fi38++ {
                    v1 = v1
                    _ = fi38
                }
                var v39 int8 = (^(v37))
                _ = fi36
                _ = v37
                _ = v39
            }
            v40 := GenS1{F0: 32767}
            v41 := GenS4{F0: 42, F1: uint64(v5)}
            _ = v40
            _ = v41
        }
        var v42 uint8 = 0
        v43 := GenS5{F0: 128}
        _ = v42
        _ = v43
    }
    fmt.Println(v5, v1)
    var v44 int = 0
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
