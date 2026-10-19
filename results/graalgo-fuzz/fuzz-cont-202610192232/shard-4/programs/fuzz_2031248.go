package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0(p0 uint, p1 int32) uint {
    var v5 float64 = (1.5 / 1e10)
    v6 := []int{127, 4294967295, int(p0)}
    var v7 int32 = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 604
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 42
    var v3 uint = 65535
    v4 := make(map[string]bool)
    if ((!false) || true) {
        if (!((("foo bar" + ("\n\t" + ("foo bar" + "\n\t"))) <= "a") || ((v2 - (v2 / ((65535) | 1))) < (v2 << (((v3 % ((v3) | 1))) & 63))))) {
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 int8 = int8(v2)
                switch "\n\t" {
                case "":
                    v10 := GenS1{F0: (v2 % (((v2 / ((662) | 1))) | 1))}
                    _ = v10
                case "foo bar":
                    if false {
                        var v11 uint = (^(v3))
                        v9 += (v9 + (-(v9)))
                        if false {
                            fmt.Println(v9, v2)
                            switch (fi8 * -595) {
                            case -769:
                                for fi12 := 0; fi12 < 9; fi12++ {
                                    switch ("" + "\n\t") {
                                    case "hello":
                                        v1["\n\t"] = "hello"
                                    case "foo bar":
                                        if false {
                                            fmt.Println(v3)
                                        }
                                    }
                                    if true {
                                        fmt.Println(v9, fi12, v3)
                                        if false {
                                            fmt.Println(fi12, v3, v9)
                                            fmt.Println(v2, v3, v3)
                                            v11 -= v11
                                        } else {
                                            v13 := GenS2{F0: 1}
                                            _ = v13
                                        }
                                        var v14 uint32 = 2147483648
                                        _ = v14
                                    }
                                    var v15 int32 = 1
                                    _ = fi12
                                    _ = v15
                                }
                            case 199:
                                v16 := GenS1{F0: 61}
                                _ = v16
                            }
                        }
                        _ = v11
                    }
                case "hello":
                    v17 := GenS2{F0: -914}
                    _ = v17
                default:
                    if true {
                        v3 += (^(v3))
                        v1["a"] = "foo bar"
                    }
                }
                _ = fi8
                _ = v9
            }
            v18 := GenS1{F0: (^(v2))}
            _ = v18
        }
    } else {
        var v19 uint32 = 332
        _ = v19
    }
    var v20 uint8 = 42
    v3 -= (v3 + (v3 & 762))
    if true {
        var v21 float32 = -2.25
        _ = v21
    } else {
        fmt.Println(v3)
        v22 := GenS3{F0: -1, F1: "a"}
        _ = v22
    }
    v23 := GenS4{F0: v3, F1: 1}
    var v24 int32 = 1
    if false {
        var v25 float64 = -2.25
        _ = v25
    } else {
        v3 = (v3 ^ (v3 >> ((genFn0(v3, (v24 >> ((v3) & 63)))) & 63)))
    }
    if false {
        if (-623 == -32768) {
            v24 += (v24 + v24)
            var v26 int16 = -1
            if (false || ((^(v20)) < (^(v20)))) {
                v3 = (^(v3))
                v27 := []int{4294967295, 127, 4294967295}
                switch 32768 {
                case -2147483648:
                    v28 := GenS3{F0: 1, F1: ("a" + "hello")}
                    _ = v28
                case -492:
                    delete(v1, ("hello" + "foo bar"))
                default:
                    delete(v4, "foo bar")
                }
                _ = v27
            }
            _ = v26
        }
        var v29 int = -9223372036854775808
        for fi30 := 0; fi30 < 17; fi30++ {
            for fi31 := 0; fi31 < 2; fi31++ {
                v32 := GenS4{F0: v3, F1: (-(v24))}
                fmt.Println(fi30, v29, fi31)
                _ = fi31
                _ = v32
            }
            fmt.Println(v20)
            _ = fi30
        }
        _ = v29
    }
    var v33 bool = (-32768 != 1000)
    v33 = (2 == 0)
    v33 = (302 != 4294967295)
    var v34 uint64 = 18446744073709551615
    v35 := []int{4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
