package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint8 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]bool)
    v3 := []uint32{}
    var v4 int32 = (v1 << ((461) & 63))
    switch 745 {
    case 128:
        if false {
            var v5 int8 = 68
            fmt.Println(v1, v5, v4)
            v4 -= v1
            _ = v5
        } else {
            v1 -= v4
        }
    case 2147483647:
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 int = 42
            v7 -= fi6
            if true {
                var v8 uint8 = 2
                fmt.Println(v1, fi6)
                v4 += (v1 + 1)
                _ = v8
            } else {
                var v9 int = v7
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    case 944:
        var v10 uint32 = 0
        _ = v10
    }
    v11 := GenS1{F0: 1, F1: ("hello" + "hello"), F2: 2}
    v12 := GenS2{F0: "\n\t", F1: "foo bar"}
    v1 = v4
    fmt.Println(v4)
    v4 = v4
    v4 += (v4 & 2147483647)
    v13 := GenS2{F0: (("" + "") + ""), F1: ("\n\t" + "a")}
    var v14 bool = (-49 == 42)
    fmt.Println(v4, v4, v14)
    switch ("a" + (("\n\t" + ("" + "")) + ("hello" + ("" + "\n\t")))) {
    case "\n\t":
        v15 := v2[93]
        _ = v15
    default:
        if v14 {
            if len(v3) > 0 {
                v3[0] = 737
            }
            v14 = v14
            v1 = v4
        } else {
            var v16 int32 = v4
            _ = v16
        }
    }
    v1 = v1
    if (!("foo bar" >= "")) {
        v17 := GenS2{F0: ("foo bar" + "a"), F1: "a"}
        _ = v17
    }
    var v18 uint8 = 127
    if (!(255 == 18446744073709551615)) {
        var v19 float32 = 3.14159
        var v20 int8 = int8(v1)
        _ = v19
        _ = v20
    } else {
        _ = len(v3)
        for fi21 := 0; fi21 < 7; fi21++ {
            switch (fi21 - (fi21 << ((867) & 63))) {
            case 633:
                switch (("" + ("" + (("\n\t" + "\n\t") + ("foo bar" + ("hello" + "hello"))))) + ("" + ("hello" + "foo bar"))) {
                case "\n\t":
                    if (213 != 977) {
                        if v14 {
                            v22 := GenS3{F0: v14, F1: 229, F2: 18446744073709551615}
                            _ = v22
                        } else {
                            v14 = v14
                            var v23 float32 = (0.0 - float32(0.1))
                            _ = v23
                        }
                        var v24 int64 = 1
                        for fi25 := 0; fi25 < 13; fi25++ {
                            switch (-(fi25)) {
                            case 32768:
                                var v26 uint = 2
                                _ = v26
                            }
                            v27 := GenS3{F0: ("" > (("\n\t" + ("foo bar" + ("hello" + "foo bar"))) + ("\n\t" + ((("a" + ("a" + "a")) + ((("hello" + ("foo bar" + "")) + "") + "hello")) + "")))), F1: 0, F2: 682}
                            _ = fi25
                            _ = v27
                        }
                        _ = v24
                    } else {
                        if (-606 > -651) {
                            if (((false && (-128 != 113)) && (1000 <= 32767)) && v14) {
                                v28 := GenS4{F0: 646}
                                _ = v28
                            }
                            v4 -= 0
                        }
                    }
                default:
                    v29 := GenS5{F0: 655, F1: (v18 + (v18 ^ 255))}
                    _ = v29
                }
            case 1:
                var v30 int32 = (v4 * (v1 << ((0) & 63)))
                _ = v30
            default:
                if v14 {
                    var v31 int8 = 42
                    var v32 int = 127
                    _ = v31
                    _ = v32
                }
            }
            _ = fi21
        }
        v33 := GenS2{F0: "\n\t", F1: "hello"}
        _ = v33
    }
    fmt.Println(v18, v14)
    var v34 float64 = float64(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
