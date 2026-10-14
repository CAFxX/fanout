package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 uint32 = 1
    v2 := []int{1}
    v3 := make(map[int]string)
    var v4 int32 = 42
    if (18446744073709551615 < 255) {
        fmt.Println(v4, v4)
    } else {
        fmt.Println(v4)
        for fi5 := 0; fi5 < 13; fi5++ {
            if true {
                var v6 uint = 65535
                fmt.Println(v1, fi5, v1)
                _ = v6
            } else {
                if true {
                    if false {
                        v4 = 31
                        for fi7 := 0; fi7 < 15; fi7++ {
                            v4 -= v4
                            _ = fi7
                        }
                    }
                } else {
                    switch (fi5 >> ((832) & 63)) {
                    case 127:
                        for fi8 := 0; fi8 < 15; fi8++ {
                            v3[fi8] = "hello"
                            _ = fi8
                        }
                    case 949:
                        var v9 uint = 682
                        _ = v9
                    default:
                        if (true && (53 < -1)) {
                            v1 += (v1 >> ((202) & 63))
                        } else {
                            delete(v3, (fi5 - -2147483648))
                            v10 := v3[-9223372036854775808]
                            if len(v2) > 0 {
                                v2[0] = fi5
                            }
                            _ = v10
                        }
                    }
                    var v11 float32 = 0.0
                    _ = v11
                }
            }
            v4 += (-(v4))
            if true {
                fmt.Println(v1)
                v1 = (v1 & 0)
                v1 -= 2147483648
            } else {
                switch "\n\t" {
                case "hello":
                    v12 := GenS1{F0: 25, F1: 127}
                    _ = v12
                case "a":
                    v13 := GenS2{F0: 32767}
                    _ = v13
                case "foo bar":
                    _ = len(v3)
                }
            }
            _ = fi5
        }
        var v14 int = -888
        _ = v14
    }
    switch 127 {
    case 2147483647:
        v15 := GenS3{F0: 127, F1: 0}
        _ = v15
    case -902:
        _ = len(v3)
    }
    v4 += (v4 >> ((4294967295) & 63))
    v16 := GenS4{F0: 1, F1: 1}
    v3[-115] = ((("hello" + "") + "\n\t") + (((("hello" + ("a" + "")) + "\n\t") + "\n\t") + (((("" + (("a" + "a") + ("" + ("foo bar" + "hello")))) + "\n\t") + ("\n\t" + ("foo bar" + ((("hello" + ("hello" + "")) + (("foo bar" + "hello") + "foo bar")) + "\n\t")))) + (((("" + "foo bar") + "hello") + "foo bar") + ""))))
    v17 := GenS5{F0: -261}
    for fi18 := 0; fi18 < 1; fi18++ {
        switch ("hello" + "\n\t") {
        case "foo bar":
            v4 = (v4 >> ((1) & 63))
        case "":
            for fi19 := 0; fi19 < 19; fi19++ {
                if (!(true || (!false))) {
                    v3[(-(fi19))] = "\n\t"
                    v20 := GenS6{F0: -654}
                    var v21 uint = 1
                    _ = v20
                    _ = v21
                } else {
                    if (109 > 10) {
                        v4 = v4
                    } else {
                        var v22 int32 = v4
                        v1 += v1
                        var v23 uint8 = 127
                        _ = v22
                        _ = v23
                    }
                }
                var v24 int16 = 1000
                _ = fi19
                _ = v24
            }
        default:
            if (4294967295 != v1) {
                v4 = -2147483648
                fmt.Println(v4, v1)
                v4 += (^(v4))
            } else {
                v25 := GenS7{F0: 270}
                fmt.Println(v1, fi18, v1)
                fmt.Println(v1, fi18)
                _ = v25
            }
        }
        _ = fi18
    }
    _ = len(v3)
    v26 := GenS2{F0: 1}
    var v27 int32 = v4
    var v28 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
