package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint }
func genFn0() int8 {
    if true {
        var v4 bool = ("" <= "hello")
        v5 := []int{-1, 9223372036854775807, 128, 32768}
        _ = v4
        _ = v5
    } else {
        var v6 bool = true
        fmt.Println(v6)
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: (-(fi8)), F1: ""}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 42, F1: "a"}
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(fi11)
            var v12 int64 = 231
            v13 := GenS2{F0: "a", F1: 42}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi7
        _ = v10
    }
    var v14 bool = true
    v14 = false
    _ = v14
    return 1
}

func main() {
    v1 := []uint32{2147483648, 2147483648, 2147483648, 725}
    var v2 uint = 42
    var v3 int8 = -51
    var v15 int8 = (v3 / ((v3) | 1))
    v3 += (v15 | 1)
    v3 -= (v3 * (v15 ^ genFn0()))
    var v16 float64 = (0.0 + float64(0.0))
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            fmt.Println(v2, v2)
            _ = len(v1)
            _ = fi18
        }
        v1 = append(v1, 2147483647)
        v19 := GenS1{F0: fi17, F1: ("a" + "foo bar")}
        _ = fi17
        _ = v19
    }
    switch 42 {
    case 611:
        if (!false) {
            if true {
                _ = len(v1)
            } else {
                var v20 uint32 = 989
                _ = v20
            }
            for fi21 := 0; fi21 < 6; fi21++ {
                v2 = (^(v2))
                var v22 float64 = -2.25
                _ = fi21
                _ = v22
            }
        } else {
            for fi23 := 0; fi23 < 16; fi23++ {
                if (0 >= -32768) {
                    var v24 uint8 = 255
                    var v25 uint16 = 833
                    var v26 int = (-(fi23))
                    _ = v24
                    _ = v25
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 5; fi27++ {
                        v28 := GenS2{F0: "hello", F1: 587}
                        var v29 string = ("hello" + ("hello" + ("" + ("" + ("a" + ("\n\t" + "foo bar"))))))
                        var v30 int16 = -273
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        switch v29 {
                        case "":
                            var v31 int16 = -1
                            _ = v31
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                }
                _ = fi23
            }
            _ = len(v1)
            v16 -= (3.14159 - (0.1 + v16))
        }
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        if true {
            v33 := GenS3{F0: 1}
            if (!(((("" + ("hello" + "foo bar")) + ((((("foo bar" + "\n\t") + "") + "a") + ("\n\t" + "foo bar")) + "foo bar")) <= ("foo bar" + "foo bar")) || (1 < 0))) {
                v2 += (v2 % (((v2 % ((v2) | 1))) | 1))
                fmt.Println(v15)
            }
            _ = v33
        }
        _ = fi32
    }
    v34 := GenS1{F0: 4294967295, F1: "\n\t"}
    fmt.Println(v15, v2)
    fmt.Println(v16)
    v3 -= (v15 << ((843) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v34
}
