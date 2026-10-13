package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint }
func genFn0() uint16 {
    v6 := GenS1{F0: -1, F1: (("" + "\n\t") + "foo bar"), F2: 58}
    var v7 int32 = 995
    var v8 uint = 817
    _ = v6
    _ = v7
    _ = v8
    return 224
}
func genFn1() string {
    var v9 uint64 = 329
    v10 := GenS1{F0: 9223372036854775807, F1: (("hello" + (("a" + "hello") + "a")) + "a"), F2: 101}
    v11 := GenS2{F0: (34 != 128)}
    var v12 float64 = -2.25
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return ""
}

func main() {
    var v1 float64 = ((((3.14159 / (3.14159 - (float64(1.5) + 3.14159))) + (0.0 + ((float64(3.14159) + (1e10 + (0.0 - (0.0 / (3.14159 / -2.25))))) * 0.1))) / 1e10) - 0.1)
    v2 := []int{9223372036854775807, 42}
    var v3 uint = 953
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]int)
    for fi13 := 0; fi13 < 17; fi13++ {
        v1 = v1
        var v14 int16 = 32767
        _ = fi13
        _ = v14
    }
    fmt.Println(v3)
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 uint8 = 255
            v18 := GenS3{F0: 4294967295, F1: 0, F2: 1}
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := GenS4{F0: 1, F1: -1, F2: 0}
        _ = fi15
        _ = v19
    }
    v20 := GenS2{F0: false}
    v21 := GenS5{F0: 1}
    for fi22 := 0; fi22 < 4; fi22++ {
        v3 -= v3
        _ = fi22
    }
    if ((1e10 + v1) != v1) {
        var v23 uint8 = 127
        _ = v23
    } else {
        if (false || (2147483648 == 743)) {
            if false {
                fmt.Println(v1)
            }
            v1 = float64(0.1)
        } else {
            var v24 int = 2147483647
            var v25 float64 = ((float64(0.0) + 0.1) + v1)
            if false {
                v3 += (^(v3))
                fmt.Println(v3, v25)
                if false {
                    _ = len(v4)
                } else {
                    v26 := GenS6{F0: 1}
                    fmt.Println(v1, v24)
                    _ = v26
                }
            } else {
                v3 = v3
                v1 = (v1 / v25)
            }
            _ = v24
            _ = v25
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch "\n\t" {
    case "a":
        if (v1 >= v1) {
            var v27 uint16 = 32767
            _ = v27
        }
    case "foo bar":
        v4[531] = -656
    case "":
        v1 -= ((((3.14159 + (((v1 * float64((-2.25 - (1e10 / (-2.25 - 0.1))))) + ((v1 / v1) / v1)) / (0.0 / v1))) / float64(3.14159)) / (3.14159 - ((0.0 - ((0.1 * v1) * v1)) + float64(1.5)))) - (float64(0.1) * 1e10))
    default:
        if (true && (!(false && (18446744073709551615 != 18446744073709551615)))) {
            v3 += 65535
            if ((false || true) && (true || (!false))) {
                fmt.Println(v3, v1)
                fmt.Println(v3, v3)
                v28 := GenS3{F0: 4, F1: 745, F2: -9223372036854775808}
                _ = v28
            }
        } else {
            var v29 int32 = -1
            v30 := GenS7{F0: (v3 * (v3 ^ 2))}
            _ = v29
            _ = v30
        }
    }
    fmt.Println(v3, v1)
    _ = len(v5)
    v4[2147483648] = -753
    switch (("a" + "a") + ("" + ("a" + "\n\t"))) {
    case "hello":
        var v31 bool = (true && (((((((("foo bar" + "") + "") + "\n\t") + ("\n\t" + ("\n\t" + "a"))) + ("a" + "foo bar")) + "") + (((("\n\t" + "a") + ("foo bar" + "hello")) + "") + "")) <= ("" + (("" + "\n\t") + ("foo bar" + ((("a" + "\n\t") + "\n\t") + "hello"))))))
        _ = v31
    case "foo bar":
        switch 9223372036854775807 {
        case 4294967295:
            v3 -= (v3 >> ((v3) & 63))
        case -9223372036854775808:
            if (!(508 <= 1)) {
                var v32 string = (("foo bar" + "") + "a")
                _ = v32
            } else {
                for fi33 := 0; fi33 < 20; fi33++ {
                    if (65535 <= 374) {
                        for fi34 := 0; fi34 < 12; fi34++ {
                            if false {
                                v35 := GenS1{F0: int64(fi34), F1: ("" + "\n\t"), F2: 100}
                                _ = v35
                            }
                            _ = fi34
                        }
                        fmt.Println(v1, fi33)
                    } else {
                        fmt.Println(v1, fi33, v3)
                        v1 = (v1 + 1.5)
                        v1 += v1
                    }
                    _ = fi33
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
