package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int; F1 int64; F2 uint64 }
func genFn0(p0 uint32, p1 int16) string {
    fmt.Println(p0, p1)
    v6 := GenS1{F0: 602, F1: -866}
    if true {
        if true {
            for fi7 := 0; fi7 < 16; fi7++ {
                p1 = (p1 * (p1 | p1))
                _ = fi7
            }
            var v8 float32 = 0.0
            var v9 string = ("foo bar" + "")
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 12; fi10++ {
                if true {
                    v11 := GenS1{F0: 42, F1: fi10}
                    _ = v11
                } else {
                    p1 += 42
                }
                _ = fi10
            }
            v12 := GenS2{F0: 1}
            p0 = (^(p0))
            _ = v12
        }
    }
    v13 := []int{-991, 32767, 128}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    return ("\n\t" + (("\n\t" + ((("foo bar" + "foo bar") + ("a" + ((("\n\t" + ("\n\t" + "foo bar")) + "") + ((("a" + "\n\t") + "foo bar") + "")))) + ("hello" + ("\n\t" + ("a" + ("foo bar" + ("a" + ""))))))) + "foo bar"))
}
func genFn1(p0 string, p1 bool) uint16 {
    var v14 uint8 = 1
    var v15 bool = true
    v16 := GenS2{F0: 2}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return 255
}

func main() {
    var v1 uint8 = 2
    v2 := []int{9223372036854775807, 42}
    var v3 int8 = 127
    v4 := make(map[int]bool)
    var v5 uint64 = 0
    if (v3 == v3) {
        if (-1 < 4294967295) {
            v3 = (v3 & v3)
            switch 127 {
            case -2147483648:
                if true {
                    var v17 int = 2147483647
                    var v18 int32 = -1
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        switch v17 {
                        case 506:
                            if (!true) {
                                var v19 int8 = v3
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v19
                            }
                        case 45:
                            v5 += v5
                        case -9223372036854775808:
                            _ = len(v2)
                        default:
                            if false {
                                v18 = v18
                            }
                        }
                    }
                    _ = v17
                    _ = v18
                }
            case -420:
                if (((!(((float64(-2.25) * (1.5 - ((((-2.25 / 3.14159) - 0.1) / 1.0) / (0.1 / 3.14159)))) == (0.0 + (-2.25 - 1e10))) || true)) || true) || false) {
                    delete(v4, 466)
                } else {
                    fmt.Println(v1, v3, v5)
                    var v20 uint16 = 2
                    _ = v20
                }
            default:
                v4[4294967295] = (((("\n\t" + "foo bar") + "a") + "\n\t") > "a")
            }
        }
    } else {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int64 = 827
            _ = fi21
            _ = v22
        }
        v1 = v1
    }
    var v23 int64 = 1
    if (-2147483648 == -2147483648) {
        v24 := GenS3{F0: 166, F1: 0}
        var v25 int = 1
        _ = v24
        _ = v25
    } else {
        switch (("a" + "") + "hello") {
        case "hello":
            var v26 bool = false
            _ = v26
        case "\n\t":
            var v27 uint = 734
            _ = v27
        default:
            v5 += v5
        }
    }
    var v28 float64 = (-2.25 - float64(float32(0.0)))
    v29 := GenS1{F0: 4294967295, F1: -9223372036854775808}
    v28 += v28
    var v30 uint64 = 276
    _ = len(v2)
    if true {
        v31 := GenS3{F0: -9223372036854775808, F1: 32767}
        _ = v31
    } else {
        _ = len(v2)
        fmt.Println(v28, v28)
    }
    v32 := GenS4{F0: 427, F1: (v23 / (((v23 - 1)) | 1)), F2: 0}
    fmt.Println(v5)
    if true {
        v33 := GenS2{F0: v5}
        _ = v33
    }
    v30 = v5
    var v34 uint64 = (v30 ^ v5)
    delete(v4, 166)
    delete(v4, int(v34))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
