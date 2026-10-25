package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0() uint16 {
    var v6 float32 = (-2.25 * (-2.25 / 1e10))
    v6 = v6
    var v7 int32 = -2147483648
    _ = v6
    _ = v7
    return 135
}
func genFn1(p0 int, p1 bool, p2 uint16) int64 {
    if false {
        var v8 uint32 = 2147483648
        _ = v8
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 int32 = 0
                if (uint32(p2) != 4294967295) {
                    p2 = 32767
                } else {
                    v12 := GenS1{F0: 0}
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            p1 = (((-1 < 70) && p1) || (542 == 1))
            _ = fi9
        }
        var v13 int8 = -1
        _ = v13
    }
    fmt.Println(p0, p1, p1)
    v14 := []int{(p0 / ((-963) | 1)), (p0 * p0), 2147483647, 750}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return -1
}

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{752, 1, 2, 42}
    v3 := make(map[uint32]string)
    var v4 uint32 = 2147483647
    var v5 float32 = ((1e10 - 1.5) * ((float32((v1 - v1)) + (1e10 - ((3.14159 / 1e10) - 1e10))) / 3.14159))
    v15 := GenS1{F0: 0}
    for fi16 := 0; fi16 < 2; fi16++ {
        switch (("a" + "\n\t") + "") {
        case "":
            var v17 int = 4294967295
            _ = v17
        case "hello":
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 uint = 1
                var v20 uint32 = 256
                v21 := GenS1{F0: 1}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        case "\n\t":
            switch fi16 {
            case -525:
                fmt.Println(v4, fi16, fi16)
            case -200:
                var v22 int16 = 32767
                _ = v22
            case -62:
                var v23 int = 127
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 int8 = 1
            switch "\n\t" {
            case "foo bar":
                v5 = (1e10 * v5)
            case "a":
                switch "" {
                case "foo bar":
                    if (1 > 0) {
                        if ((-(fi24)) < fi16) {
                            fmt.Println(v4, v5, v5)
                            v25 += 127
                        } else {
                            fmt.Println(fi16)
                        }
                    } else {
                        v26 := GenS2{F0: 477, F1: 255}
                        _ = v26
                    }
                case "\n\t":
                    if (!(!(191 <= 742))) {
                        v1 = (0.0 - (v1 * -2.25))
                        v27 := GenS3{F0: (v25 << ((42) & 63))}
                        _ = v27
                    } else {
                        v2 = append(v2, (^(v4)))
                        v28 := GenS2{F0: 501, F1: 0}
                        if (65535 > 421) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v29 := v3[(v4 * v4)]
                            _ = v29
                        }
                        _ = v28
                    }
                    v4 += 2147483648
                    if true {
                        if false {
                            var v30 bool = false
                            var v31 uint = 611
                            var v32 int = 2147483647
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            if false {
                                var v33 float64 = (v1 / (0.1 - float64((-2.25 - (v5 + v5)))))
                                _ = v33
                            }
                            v34 := GenS3{F0: v25}
                            _ = v34
                        }
                    } else {
                        var v35 float32 = float32(0.0)
                        _ = v35
                    }
                default:
                    v25 = 100
                    v36 := GenS4{F0: (("\n\t" + ("foo bar" + ((("\n\t" + ("\n\t" + "foo bar")) + "") + "a"))) + ("hello" + "foo bar")), F1: (v4 - v4)}
                    v37 := GenS3{F0: -121}
                    _ = v36
                    _ = v37
                }
            }
            fmt.Println(fi24, v25, v5)
            _ = fi24
            _ = v25
        }
        _ = fi16
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        v1 -= 1.5
        v39 := GenS1{F0: 100}
        _ = fi38
        _ = v39
    }
    v40 := GenS4{F0: "foo bar", F1: (^(v4))}
    if (false || true) {
        fmt.Println(v5)
        var v41 int32 = 245
        v42 := GenS2{F0: 255, F1: 0}
        _ = v41
        _ = v42
    }
    v1 = (((v1 / v1) * (v1 * (((v1 / v1) / (((v1 * (v1 - float64(1e10))) + v1) - float64(3.14159))) + 0.0))) / -2.25)
    switch ("foo bar" + ("\n\t" + "foo bar")) {
    case "foo bar":
        v4 = 0
    case "":
        fmt.Println(v1, v4)
    }
    var v43 uint64 = 2
    for fi44 := 0; fi44 < 15; fi44++ {
        v43 += 554
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v5
}
