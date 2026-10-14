package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 uint16 }
type GenS4 struct { F0 int; F1 int64 }

func main() {
    var v1 int = 128
    v2 := []string{"hello", "foo bar"}
    var v3 uint32 = 888
    var v4 bool = ((((v1 <= v1) || false) && true) || true)
    var v5 int64 = 9223372036854775807
    if ((!v4) || (v4 || true)) {
        if (32767 > 32767) {
            var v6 int32 = 42
            v7 := GenS1{F0: 749, F1: (v1 & v1), F2: v3}
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 int = (-(v1))
                _ = fi8
                _ = v9
            }
            _ = v6
            _ = v7
        }
    } else {
        switch "" {
        case "\n\t":
            for fi10 := 0; fi10 < 1; fi10++ {
                v3 = v3
                v11 := GenS2{F0: v4, F1: 1}
                _ = fi10
                _ = v11
            }
        case "a":
            var v12 uint16 = 42
            _ = v12
        case "foo bar":
            v1 = -788
        default:
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := GenS1{F0: 18446744073709551615, F1: fi13, F2: v3}
                v5 = (v5 % ((v5) | 1))
                v5 -= (v5 * (-(v5)))
                _ = fi13
                _ = v14
            }
        }
    }
    switch "foo bar" {
    case "":
        var v15 string = ("a" + "foo bar")
        _ = v15
    }
    v5 -= -1
    fmt.Println(v3, v1, v3)
    v4 = v4
    if v4 {
        if false {
            v3 += v3
            v16 := GenS2{F0: v4, F1: 2}
            var v17 uint32 = 2147483647
            _ = v16
            _ = v17
        } else {
            var v18 int = (v1 - (v1 | (v1 >> ((415) & 63))))
            _ = v18
        }
        if v4 {
            if (-2.25 > -2.25) {
                var v19 int8 = -1
                v2 = append(v2, "hello")
                _ = v19
            } else {
                var v20 int8 = int8(v3)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v21 int16 = 124
                _ = v20
                _ = v21
            }
            switch v1 {
            case 32767:
                var v22 uint = 684
                _ = v22
            default:
                for fi23 := 0; fi23 < 13; fi23++ {
                    if (0 == 75) {
                        for fi24 := 0; fi24 < 17; fi24++ {
                            v3 = v3
                            _ = fi24
                        }
                        var v25 uint64 = 956
                        v26 := GenS3{F0: (fi23 % ((42) | 1)), F1: v5, F2: 455}
                        _ = v25
                        _ = v26
                    } else {
                        v27 := GenS1{F0: 4294967295, F1: 42, F2: v3}
                        var v28 uint16 = 0
                        v3 = v3
                        _ = v27
                        _ = v28
                    }
                    _ = fi23
                }
            }
        }
        if (65535 > 32767) {
            var v29 uint16 = 964
            _ = v29
        }
    }
    var v30 uint64 = 2
    var v31 bool = false
    var v32 uint64 = v30
    fmt.Println(v5, v1, v5)
    var v33 float32 = ((0.0 * (0.1 * 0.1)) * 0.1)
    for fi34 := 0; fi34 < 2; fi34++ {
        switch ("\n\t" + ("a" + (("hello" + "foo bar") + "a"))) {
        case "":
            if (v31 || false) {
                v2 = append(v2, ("foo bar" + ""))
                var v35 uint64 = (v32 % ((v32) | 1))
                _ = v35
            }
        default:
            if (1 >= v3) {
                var v36 uint64 = 4294967295
                _ = v36
            } else {
                if (!(v30 <= (v32 - (^(v32))))) {
                    v37 := GenS4{F0: 127, F1: (v5 << ((465) & 63))}
                    if ((((0.1 - 3.14159) + 0.0) > -2.25) && (v1 > 128)) {
                        if v4 {
                            v31 = false
                        }
                        var v38 uint64 = (v32 | 2)
                        _ = v38
                    } else {
                        v31 = true
                        v39 := GenS2{F0: v4, F1: 42}
                        v3 = uint32(v32)
                        _ = v39
                    }
                    _ = v37
                } else {
                    fmt.Println(v32)
                    var v40 int = (fi34 / ((-703) | 1))
                    var v41 bool = (true && v31)
                    _ = v40
                    _ = v41
                }
                var v42 bool = (((v5 * (v5 / ((-9223372036854775808) | 1))) >= 42) && v31)
                _ = v42
            }
        }
        for fi43 := 0; fi43 < 8; fi43++ {
            var v44 int8 = 42
            var v45 uint16 = 255
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi43
            _ = v44
            _ = v45
        }
        fmt.Println(v32, v30)
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
