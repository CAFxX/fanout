package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint32 }
func genFn0() int {
    var v7 int16 = 0
    v8 := GenS1{F0: false, F1: 479}
    if false {
        var v9 int16 = v7
        _ = v9
    } else {
        if (65535 >= 42) {
            fmt.Println(v7)
            fmt.Println(v7)
            switch "hello" {
            case "":
                var v10 int32 = 0
                _ = v10
            case "foo bar":
                if (653 >= 32767) {
                    fmt.Println(v7)
                    switch ((((("a" + ("a" + "foo bar")) + ((("foo bar" + "") + (("\n\t" + "\n\t") + "")) + ("" + "hello"))) + (("hello" + "hello") + "\n\t")) + "a") + "foo bar") {
                    case "hello":
                        var v11 uint32 = 384
                        _ = v11
                    default:
                        switch "" {
                        case "\n\t":
                            if (!(!true)) {
                                v7 = 1
                            } else {
                                switch 761 {
                                case -9223372036854775808:
                                    v12 := GenS2{F0: 619, F1: 1}
                                    _ = v12
                                default:
                                    var v13 int64 = 0
                                    _ = v13
                                }
                                v14 := GenS1{F0: (330 == 4294967295), F1: 763}
                                _ = v14
                            }
                        case "hello":
                            if (true && true) {
                                switch 1 {
                                case 0:
                                    v15 := []int{-1, 32767, -2147483648, -9223372036854775808}
                                    _ = v15
                                default:
                                    var v16 float32 = float32((((0.1 - 0.1) - 0.1) - 3.14159))
                                    _ = v16
                                }
                            } else {
                                v7 += (v7 << ((65535) & 63))
                                var v17 uint64 = 511
                                _ = v17
                            }
                        }
                    }
                }
            case "a":
                if (-77 >= 42) {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        if (("\n\t" + "") <= (("" + ((("\n\t" + "hello") + ("hello" + ("hello" + ("foo bar" + "hello")))) + "a")) + ("a" + "hello"))) {
                            v7 = 0
                            var v19 int8 = 101
                            v20 := GenS1{F0: true, F1: 1}
                            _ = v19
                            _ = v20
                        } else {
                            var v21 int = (fi18 - fi18)
                            v21 += -9223372036854775808
                            _ = v21
                        }
                        v7 = (v7 + -45)
                        _ = fi18
                    }
                    if false {
                        for fi22 := 0; fi22 < 19; fi22++ {
                            v23 := GenS3{F0: 2}
                            v24 := GenS2{F0: (fi22 << ((65535) & 63)), F1: 1}
                            v7 -= 42
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        v7 -= (v7 << ((566) & 63))
                    }
                }
            default:
                v25 := GenS3{F0: 0}
                _ = v25
            }
        } else {
            v7 += (v7 | v7)
            if true {
                v26 := GenS2{F0: -9223372036854775808, F1: 1}
                if false {
                    fmt.Println(v7)
                }
                _ = v26
            }
            v7 -= v7
        }
    }
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 int16, p1 int64, p2 uint8) bool {
    v27 := GenS4{F0: p0, F1: 42, F2: 100}
    p1 += p1
    fmt.Println(p1, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return true
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 0
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 uint64 = 42
    var v6 int8 = 1
    v28 := GenS1{F0: (-322 != 1000), F1: 65535}
    v2 += v2
    var v29 int64 = 0
    v1 = (v1 << ((18446744073709551615) & 63))
    var v30 int64 = 1
    v3[-2147483648] = ((1000 == 1000) || (false && true))
    for fi31 := 0; fi31 < 10; fi31++ {
        for fi32 := 0; fi32 < 10; fi32++ {
            v6 += (-(v6))
            _ = fi32
        }
        for fi33 := 0; fi33 < 19; fi33++ {
            v34 := []int{}
            _ = fi33
            _ = v34
        }
        _ = fi31
    }
    var v35 uint64 = (^(v5))
    for fi36 := 0; fi36 < 16; fi36++ {
        for fi37 := 0; fi37 < 9; fi37++ {
            var v38 float64 = 0.0
            v2 += 166
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    v39 := GenS5{F0: (v5 | 90), F1: v1, F2: 813}
    _ = len(v3)
    if false {
        v40 := v4[""]
        v41 := []int{-687, genFn0()}
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
