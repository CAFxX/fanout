package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 uint; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 32767
    var v3 int32 = 1
    v4 := []int{32767}
    var v5 string = ("a" + "a")
    v6 := []string{v5, (((v5 + "") + "\n\t") + "a"), v5, ""}
    v7 := GenS1{F0: 1}
    v2 += v2
    v8 := GenS1{F0: 4294967295}
    if (true || true) {
        _ = len(v1)
    }
    v5 = ("foo bar" + "hello")
    var v9 int8 = 100
    v5 = v5
    fmt.Println(v2, v9)
    v10 := GenS1{F0: 2147483648}
    v1["hello"] = v5
    var v11 int64 = -729
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 uint16 = 714
        switch fi12 {
        case -2147483648:
            v14 := GenS1{F0: 2147483648}
            _ = v14
        case 0:
            if (!(false || (-360 == -384))) {
                fmt.Println(v13)
                fmt.Println(v13, fi12)
            } else {
                if ("\n\t" < ("a" + ("\n\t" + "hello"))) {
                    var v15 int = fi12
                    if (fi12 > fi12) {
                        for fi16 := 0; fi16 < 19; fi16++ {
                            delete(v1, (((v5 + v5) + v5) + v5))
                            if len(v6) > 0 {
                                v6[0] = "foo bar"
                            }
                            switch v5 {
                            case "\n\t":
                                v17 := GenS1{F0: 2147483647}
                                _ = v17
                            case "foo bar":
                                if ((208 <= 4294967295) || ((((1e10 / 3.14159) / 0.1) < (1.5 / float32(0.1))) || ((1 < v9) || false))) {
                                    if true {
                                        v13 = v13
                                    } else {
                                        if true {
                                            v18 := GenS2{F0: 378, F1: 2, F2: "a"}
                                            _ = v18
                                        }
                                        fmt.Println(v3, v11)
                                        if ((v13 >= (v13 * v13)) && true) {
                                            var v19 string = "a"
                                            _ = v19
                                        }
                                    }
                                }
                            case "hello":
                                var v20 uint32 = 846
                                _ = v20
                            default:
                                fmt.Println(v11, v3, fi12)
                            }
                            fmt.Println(fi16, v2, v3)
                            _ = fi16
                        }
                        for fi21 := 0; fi21 < 3; fi21++ {
                            var v22 int16 = v2
                            var v23 uint32 = 2147483647
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        switch (v15 * 322) {
                        case -9223372036854775808:
                            delete(v1, "a")
                        }
                    }
                    if (745 >= uint32(v13)) {
                        v13 = v13
                        fmt.Println(fi12)
                    } else {
                        v13 = (v13 << ((65535) & 63))
                        for fi24 := 0; fi24 < 10; fi24++ {
                            switch v15 {
                            case 556:
                                if true {
                                    v9 -= v9
                                } else {
                                    v3 += (v3 - (v3 >> ((2) & 63)))
                                    v15 += v15
                                }
                            }
                            var v25 uint32 = 2147483648
                            _ = fi24
                            _ = v25
                        }
                        if (160 != 127) {
                            var v26 int64 = (v11 * 9223372036854775807)
                            v11 -= v11
                            v2 = (v2 ^ v2)
                            _ = v26
                        } else {
                            var v27 int8 = v9
                            if false {
                                v13 += (v13 / (((v13 + (v13 / ((1) | 1)))) | 1))
                                v28 := GenS3{F0: 65535, F1: (v13 & (v13 / ((v13) | 1)))}
                                v2 = (v2 & (v2 / ((v2) | 1)))
                                _ = v28
                            } else {
                                v2 += 1
                            }
                            switch 9223372036854775807 {
                            case -2147483648:
                                fmt.Println(fi12)
                            case 4294967295:
                                v13 -= 2
                            case 32767:
                                delete(v1, ((v5 + "\n\t") + (("a" + "") + v5)))
                            default:
                                if ((v2 > v2) && false) {
                                    v27 = v9
                                }
                            }
                            _ = v27
                        }
                    }
                    _ = v15
                } else {
                    fmt.Println(v3, v2)
                    _ = len(v4)
                }
                v2 += (-(v2))
                delete(v1, v5)
            }
        case -887:
            v29 := GenS1{F0: 832}
            _ = v29
        }
        if true {
            fmt.Println(v11, v3)
        }
        _ = fi12
        _ = v13
    }
    var v30 int64 = (-(v11))
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
