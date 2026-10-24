package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 2
    var v3 uint64 = 1
    var v4 float32 = 1.5
    v5 := GenS1{F0: 238, F1: 9223372036854775807}
    switch 1 {
    case -2147483648:
        v6 := []int{-209}
        _ = v6
    case 32768:
        for fi7 := 0; fi7 < 17; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := []int{32768, -9223372036854775808}
                var v10 uint32 = 2147483647
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    default:
        v11 := GenS2{F0: (v3 - (v3 - 867)), F1: 4294967295}
        _ = v11
    }
    v2 += v2
    var v12 bool = true
    v13 := GenS2{F0: (^(v3)), F1: 744}
    var v14 uint32 = 42
    v15 := GenS2{F0: (v3 ^ 4294967295), F1: (v14 | (v14 & 4294967295))}
    var v16 bool = (3.14159 <= (float64(v4) - 0.0))
    for fi17 := 0; fi17 < 14; fi17++ {
        if false {
            fmt.Println(v3, v3, v2)
            v16 = ((fi17 ^ 128) >= (fi17 - fi17))
        }
        switch "hello" {
        case "foo bar":
            if (!v16) {
                for fi18 := 0; fi18 < 4; fi18++ {
                    v16 = (!v16)
                    _ = fi18
                }
                for fi19 := 0; fi19 < 6; fi19++ {
                    fmt.Println(v4)
                    fmt.Println(v16, v16)
                    if v12 {
                        if (0 != -2147483648) {
                            switch "foo bar" {
                            case "":
                                fmt.Println(fi17, v3, v16)
                                var v20 uint64 = v3
                                v16 = v16
                                _ = v20
                            }
                            if (0 < int32(v14)) {
                                fmt.Println(v12, v3)
                                var v21 float64 = 0.0
                                v3 = v3
                                _ = v21
                            }
                            var v22 float64 = -2.25
                            fmt.Println(v4, v12, v3)
                            v4 += (0.1 / 3.14159)
                            if (18446744073709551615 > 255) {
                                v14 = 1
                            }
                            _ = v22
                        }
                    }
                    _ = fi19
                }
            }
        case "\n\t":
            v16 = true
        case "hello":
            fmt.Println(fi17, fi17, v12)
        }
        _ = fi17
    }
    if v12 {
        v4 += (v4 - 0.1)
        var v23 int = 4294967295
        if v12 {
            fmt.Println(v12, v4, v16)
            v2 -= (v2 - v2)
        }
        _ = v23
    }
    fmt.Println(v3, v3, v2)
    for fi24 := 0; fi24 < 12; fi24++ {
        v12 = (v12 || v16)
        v4 = v4
        if ((v3 << ((1) & 63)) < 18446744073709551615) {
            var v25 string = "a"
            v26 := []int{}
            var v27 int = 201
            _ = v25
            _ = v26
            _ = v27
        } else {
            var v28 uint64 = v3
            for fi29 := 0; fi29 < 6; fi29++ {
                v16 = v16
                _ = fi29
            }
            for fi30 := 0; fi30 < 20; fi30++ {
                switch fi30 {
                case 557:
                    var v31 float64 = 3.14159
                    _ = v31
                case 32767:
                    v12 = false
                }
                v14 = (v14 | v14)
                _ = fi30
            }
            _ = v28
        }
        _ = fi24
    }
    switch 32768 {
    case 32767:
        switch (((("hello" + "\n\t") + ("foo bar" + "\n\t")) + "") + ((("foo bar" + (("foo bar" + "\n\t") + "a")) + "foo bar") + "\n\t")) {
        case "a":
            v32 := []int{}
            _ = v32
        case "hello":
            if (((1 < 1) || (((0.0 * 1e10) / float64(float32(0.1))) < 0.1)) || ((v14 * (v14 + (^(v14)))) != (v14 * 4294967295))) {
                v33 := v1[(v14 + 65535)]
                _ = v33
            } else {
                delete(v1, (^(v14)))
            }
        }
    case 730:
        var v34 string = ""
        _ = v34
    case -2147483648:
        if v12 {
            v35 := GenS1{F0: v14, F1: -379}
            var v36 int16 = 657
            _ = v35
            _ = v36
        } else {
            if false {
                var v37 float32 = ((((0.0 * v4) - (v4 * 0.0)) + v4) + v4)
                var v38 uint = uint(v2)
                _ = v37
                _ = v38
            } else {
                for fi39 := 0; fi39 < 16; fi39++ {
                    v40 := GenS2{F0: v3, F1: v14}
                    _ = fi39
                    _ = v40
                }
                var v41 uint32 = v14
                _ = v41
            }
            v42 := []int{1, 0}
            _ = v42
        }
    default:
        v12 = (255 <= 42)
    }
    var v43 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
