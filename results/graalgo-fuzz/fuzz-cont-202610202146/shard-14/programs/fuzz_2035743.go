package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int32; F1 string; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint16; F2 int64 }
type GenS5 struct { F0 int64; F1 bool; F2 uint32 }
type GenS6 struct { F0 int32; F1 uint8 }
func genFn0(p0 string, p1 uint16, p2 uint16) string {
    var v4 uint32 = 4294967295
    for fi5 := 0; fi5 < 8; fi5++ {
        p2 = (p2 | (p1 << ((203) & 63)))
        v4 = v4
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := []int{fi6}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := []int{128, 254, 9223372036854775807, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return ("hello" + "")
}
func genFn1(p0 uint8) bool {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return false
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]int)
    v3 := []string{"\n\t", "hello", (((((("\n\t" + "\n\t") + ("foo bar" + "")) + "a") + "hello") + ((("" + "a") + "") + "foo bar")) + "\n\t")}
    v1 = v1
    v9 := v2["a"]
    if (!(!(true && (0 == 0)))) {
        v1 = (v1 % ((v1) | 1))
    }
    v10 := v2[("foo bar" + "")]
    for fi11 := 0; fi11 < 14; fi11++ {
        fmt.Println(v10, v1, v10)
        _ = fi11
    }
    switch (v10 | 1) {
    case -860:
        if false {
            if true {
                if (773 >= 2147483648) {
                    fmt.Println(v1, v9)
                } else {
                    if (1000 != 398) {
                        v12 := GenS1{F0: 9223372036854775807, F1: 0, F2: 127}
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 5; fi13++ {
                            fmt.Println(v9, v10, fi13)
                            _ = fi13
                        }
                    }
                }
                switch (v10 << ((4294967295) & 63)) {
                case -576:
                    v14 := GenS1{F0: v10, F1: 1000, F2: 116}
                    _ = v14
                }
            } else {
                if true {
                    if true {
                        var v15 int = 4294967295
                        v9 += v9
                        _ = v15
                    } else {
                        v16 := v2[(("a" + ("" + "a")) + ((("a" + "") + "hello") + ("" + "")))]
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                        for fi17 := 0; fi17 < 7; fi17++ {
                            fmt.Println(v16)
                            var v18 float32 = (3.14159 / 1.0)
                            fmt.Println(v1)
                            _ = fi17
                            _ = v18
                        }
                        _ = v16
                    }
                } else {
                    fmt.Println(v9, v9)
                }
                for fi19 := 0; fi19 < 11; fi19++ {
                    switch "hello" {
                    case "a":
                        fmt.Println(v9, v1, v10)
                    }
                    v3 = append(v3, ("\n\t" + "foo bar"))
                    _ = fi19
                }
                for fi20 := 0; fi20 < 17; fi20++ {
                    v1 = v1
                    _ = len(v3)
                    v21 := GenS2{F0: v1, F1: 178, F2: (^(v1))}
                    _ = fi20
                    _ = v21
                }
            }
            v22 := GenS3{F0: 2147483647, F1: "\n\t", F2: 1}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 uint8 = 127
                for fi25 := 0; fi25 < 9; fi25++ {
                    switch -9223372036854775808 {
                    case -320:
                        v24 = (v24 ^ v24)
                    case -2147483648:
                        if (!(-1 > 0)) {
                            fmt.Println(v10)
                            fmt.Println(fi25)
                            v1 -= 42
                        } else {
                            v26 := GenS4{F0: 1000, F1: 2, F2: -900}
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                            v10 += -170
                            _ = v26
                        }
                    case 1:
                        switch fi25 {
                        case -417:
                            var v27 uint64 = v1
                            _ = v27
                        }
                    default:
                        switch "\n\t" {
                        case "\n\t":
                            v10 -= fi23
                            var v28 uint16 = 65535
                            _ = v28
                        case "foo bar":
                            var v29 uint32 = 2
                            v30 := GenS5{F0: 575, F1: false, F2: v29}
                            switch ("foo bar" + "\n\t") {
                            case "":
                                fmt.Println(fi25, v29)
                            }
                            fmt.Println(v29, v9)
                            _ = v29
                            _ = v30
                        case "":
                            v9 = (-(fi25))
                        }
                    }
                    _ = fi25
                }
                switch v9 {
                case -539:
                    if true {
                        v9 += v10
                        for fi31 := 0; fi31 < 19; fi31++ {
                            v32 := GenS6{F0: -1, F1: v24}
                            _ = fi31
                            _ = v32
                        }
                        v10 -= (-(v9))
                    }
                case 2147483647:
                    v1 -= 55
                }
                _ = fi23
                _ = v24
            }
            fmt.Println(v10, v9)
            var v33 int = (-(v9))
            _ = v33
        }
    default:
        var v34 bool = (true || (127 >= -1))
        _ = v34
    }
    if true {
        _ = len(v2)
        if true {
            var v35 float32 = 1.5
            v35 -= 3.14159
            _ = v35
        } else {
            for fi36 := 0; fi36 < 6; fi36++ {
                if (float64(1e10) <= (1.5 * (1.5 * 1e10))) {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    var v37 uint64 = (^(v1))
                    _ = v37
                }
                _ = fi36
            }
            fmt.Println(v9, v10, v1)
            var v38 int32 = 375
            _ = v38
        }
        _ = len(v2)
    }
    var v39 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v39
    _ = v9
}
