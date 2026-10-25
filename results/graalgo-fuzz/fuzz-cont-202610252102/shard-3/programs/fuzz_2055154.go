package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }

func main() {
    var v1 int64 = -651
    v2 := make(map[int]int)
    var v3 int16 = 1
    var v4 uint16 = 1
    var v5 int = 4294967295
    v6 := GenS1{F0: 1, F1: -128, F2: v5}
    var v7 string = (("foo bar" + "hello") + "foo bar")
    v7 = "foo bar"
    switch (((((("" + "a") + "a") + v7) + (v7 + "foo bar")) + (v7 + "a")) + (v7 + "\n\t")) {
    case "\n\t":
        v8 := []int{(-(v5)), 9223372036854775807}
        _ = v8
    }
    var v9 int8 = 127
    if true {
        fmt.Println(v1)
        if true {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 string = v7
                var v12 uint32 = 2147483647
                if false {
                    var v13 uint32 = v12
                    _ = v13
                } else {
                    var v14 uint32 = (v12 + (v12 * (v12 / (((v12 << ((0) & 63))) | 1))))
                    _ = v14
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            v15 := GenS1{F0: 42, F1: 0, F2: int(v9)}
            var v16 float32 = 0.0
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        fmt.Println(v9, v4, fi17)
        for fi18 := 0; fi18 < 7; fi18++ {
            v1 -= v1
            _ = fi18
        }
        if true {
            v1 = v1
            v5 -= v5
            v9 -= v9
        } else {
            v19 := GenS2{F0: -835, F1: true, F2: v9}
            v7 = v7
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint = 982
                if (!(32767 <= v4)) {
                    v4 = 1
                    switch v7 {
                    case "foo bar":
                        v4 = v4
                    case "\n\t":
                        v22 := GenS1{F0: 696, F1: v9, F2: v5}
                        _ = v22
                    case "":
                        v9 -= (v9 - v9)
                    default:
                        var v23 float32 = (1.5 + float32((0.0 + 1e10)))
                        _ = v23
                    }
                    if false {
                        switch (v7 + v7) {
                        case "":
                            if false {
                                v24 := GenS1{F0: 2, F1: v9, F2: fi20}
                                v21 += (v21 % (((^(v21))) | 1))
                                _ = v24
                            }
                        default:
                            v25 := GenS1{F0: 0, F1: (v9 | 0), F2: (fi17 * (-(fi20)))}
                            _ = v25
                        }
                        v21 = 255
                        v9 += (v9 ^ (v9 << ((v21) & 63)))
                    } else {
                        v26 := []int{(fi17 * -161), fi20, -1, (fi17 % ((fi20) | 1))}
                        _ = v26
                    }
                } else {
                    var v27 int64 = 42
                    if (1.5 != float32(0.1)) {
                        v21 = v21
                        if (!(!(18446744073709551615 > 263))) {
                            v4 = (v4 << (((^(v21))) & 63))
                            var v28 float64 = 3.14159
                            _ = v28
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                var v29 uint16 = (v4 % (((v4 % ((1) | 1))) | 1))
                                _ = v29
                            case "":
                                var v30 string = ("hello" + (("hello" + v7) + v7))
                                _ = v30
                            case "a":
                                fmt.Println(v1, v27)
                                if false {
                                    if false {
                                        v31 := []int{(fi20 % ((2147483647) | 1)), fi17, fi17}
                                        _ = v31
                                    }
                                }
                                var v32 uint32 = 2
                                v4 = 0
                                _ = v32
                            }
                            var v33 int32 = 409
                            _ = v33
                        }
                    }
                    _ = v27
                }
                if false {
                    v34 := v2[42]
                    v35 := GenS2{F0: (v3 / ((v3) | 1)), F1: (!(2147483647 <= -724)), F2: 0}
                    var v36 int = int(v4)
                    v21 += (^(v21))
                    _ = v34
                    _ = v35
                    _ = v36
                }
                if ((((!(!true)) || (!(!(147 == (v1 % ((v1) | 1)))))) && false) || (("" + ("\n\t" + v7)) <= v7)) {
                    v37 := []int{fi17}
                    _ = v37
                }
                if (float64(1.5) == -2.25) {
                    var v38 int = v5
                    _ = v38
                }
                _ = fi20
                _ = v21
            }
            _ = v19
        }
        _ = fi17
    }
    var v39 bool = false
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
