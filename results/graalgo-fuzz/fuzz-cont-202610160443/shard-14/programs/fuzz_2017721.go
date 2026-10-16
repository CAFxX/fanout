package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 string) uint64 {
    p0 = (p0 + p0)
    if (32767 > -372) {
        v5 := []int{4294967295, 784, 1, 0}
        _ = v5
    } else {
        v6 := GenS1{F0: 18446744073709551615}
        p0 = p0
        p0 = ((((((p0 + "") + ("a" + (("foo bar" + (p0 + "foo bar")) + p0))) + (p0 + p0)) + "hello") + "\n\t") + p0)
        _ = v6
    }
    p0 = p0
    p0 = "a"
    _ = p0
    return 698
}
func genFn1() int16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float64 = (float64(0.1) / float64((1e10 + 3.14159)))
        _ = fi7
        _ = v8
    }
    v9 := []int{64, 9223372036854775807}
    var v10 uint64 = 1
    var v11 int64 = -253
    v12 := GenS1{F0: 1}
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 42
}

func main() {
    var v1 int64 = 42
    var v2 int = 32768
    v3 := []string{"", "foo bar", ("\n\t" + "\n\t"), "a"}
    var v4 float64 = -2.25
    _ = len(v3)
    for fi13 := 0; fi13 < 12; fi13++ {
        _ = len(v3)
        if (true || (1 <= 33)) {
            v14 := GenS2{F0: -32768}
            for fi15 := 0; fi15 < 14; fi15++ {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                _ = fi15
            }
            switch "hello" {
            case "\n\t":
                if false {
                    fmt.Println(v1, v2, fi13)
                    var v16 uint8 = 1
                    _ = v16
                } else {
                    if false {
                        for fi17 := 0; fi17 < 7; fi17++ {
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                switch (fi13 ^ fi17) {
                                case -411:
                                    v1 = v1
                                case 121:
                                    v1 = v1
                                    fmt.Println(fi17, v4, v1)
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                v4 = (float64(0.1) + v4)
                                fmt.Println(v2, v2)
                                if (1e10 <= 1e10) {
                                    var v18 uint64 = 42
                                    _ = v18
                                } else {
                                    switch fi13 {
                                    case 4294967295:
                                        v4 = float64(1.5)
                                    case 127:
                                        var v19 uint32 = 2147483648
                                        _ = v19
                                    case -2147483648:
                                        if true {
                                            v3 = append(v3, "hello")
                                            fmt.Println(fi13, fi17)
                                        }
                                    default:
                                        v2 += (v2 ^ fi17)
                                    }
                                    v20 := GenS1{F0: 0}
                                    _ = v20
                                }
                            }
                            fmt.Println(v1, fi13)
                            _ = fi17
                        }
                    }
                    if true {
                        _ = len(v3)
                        var v21 uint16 = 0
                        _ = v21
                    }
                    if ((42 != 1) && false) {
                        var v22 float32 = (0.0 + (1e10 / (float32(v4) - ((0.0 - (0.0 - 3.14159)) / 1e10))))
                        _ = v22
                    } else {
                        v23 := GenS3{F0: (v2 % ((v2) | 1)), F1: 2}
                        v1 = v1
                        _ = v23
                    }
                }
            case "":
                v24 := GenS3{F0: 4294967295, F1: 37}
                _ = v24
            default:
                for fi25 := 0; fi25 < 12; fi25++ {
                    v1 = v1
                    v26 := GenS1{F0: 300}
                    _ = fi25
                    _ = v26
                }
            }
            _ = v14
        }
        _ = fi13
    }
    v1 = v1
    v2 = v2
    if (false && (v2 == (v2 << ((18446744073709551615) & 63)))) {
        v2 += (v2 - (v2 + v2))
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 int64 = -1
        for fi29 := 0; fi29 < 9; fi29++ {
            if ((true && true) || false) {
                if true {
                    if false {
                        fmt.Println(v1, fi29, v4)
                        _ = len(v3)
                        switch "" {
                        case "a":
                            fmt.Println(v1, v2, fi27)
                        case "":
                            var v30 bool = true
                            _ = v30
                        }
                    }
                    v31 := GenS3{F0: (v2 << ((1) & 63)), F1: 2}
                    _ = v31
                }
                var v32 uint = 42
                v33 := GenS3{F0: fi27, F1: 42}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v32
                _ = v33
            }
            v4 -= v4
            v34 := GenS1{F0: 395}
            _ = fi29
            _ = v34
        }
        _ = fi27
        _ = v28
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        var v36 int32 = 0
        var v37 float64 = ((v4 + float64((0.1 - -2.25))) - (v4 - float64(float32(v4))))
        if false {
            var v38 int32 = -779
            for fi39 := 0; fi39 < 9; fi39++ {
                v40 := GenS2{F0: -1}
                _ = fi39
                _ = v40
            }
            _ = v38
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    v2 += int(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
