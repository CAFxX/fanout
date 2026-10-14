package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 0
    var v3 uint8 = 0
    v4 := make(map[string]string)
    var v5 int16 = 967
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            switch 2147483647 {
            case -228:
                v5 = (v5 * 32767)
            default:
                var v9 uint16 = v2
                _ = v9
            }
            v10 := GenS1{F0: uint8(v2)}
            _ = fi8
            _ = v10
        }
        _ = fi7
    }
    var v11 float32 = 0.1
    var v12 uint8 = (v3 & (v3 - (v3 ^ (v3 / (((v3 >> ((1) & 63))) | 1)))))
    v13 := GenS1{F0: (v12 ^ (v12 ^ v12))}
    if ((v11 != v11) && false) {
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS2{F0: (42 <= 266), F1: 128}
            delete(v1, (fi14 << ((42) & 63)))
            _ = fi14
            _ = v15
        }
        if false {
            v2 += 32767
            for fi16 := 0; fi16 < 6; fi16++ {
                v5 -= v5
                v17 := []int{(fi16 ^ (-(fi16))), (fi16 << ((1) & 63))}
                for fi18 := 0; fi18 < 3; fi18++ {
                    v4[("foo bar" + "foo bar")] = ("hello" + "foo bar")
                    _ = fi18
                }
                _ = fi16
                _ = v17
            }
        }
    }
    v19 := GenS2{F0: true, F1: v3}
    if (0 == uint32(v3)) {
        v20 := GenS3{F0: -85}
        _ = v20
    }
    if ((v2 | v2) != v2) {
        fmt.Println(v2, v11)
    } else {
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 uint16 = (^(v2))
            v5 -= (v5 & v5)
            if true {
                var v23 int16 = (v5 / (((^(v5))) | 1))
                v4[((("foo bar" + "a") + (("hello" + "a") + "foo bar")) + "a")] = (("" + "hello") + "a")
                switch "\n\t" {
                case "a":
                    var v24 string = "foo bar"
                    _ = v24
                default:
                    switch (-(fi21)) {
                    case -856:
                        switch "\n\t" {
                        case "":
                            var v25 uint32 = 299
                            _ = v25
                        default:
                            if true {
                                var v26 int16 = (v5 % ((1000) | 1))
                                _ = v26
                            }
                        }
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            v27 := GenS4{F0: (v3 >> ((42) & 63))}
                            _ = v27
                        }
                    }
                }
                _ = v23
            } else {
                v28 := []int{}
                _ = v28
            }
            _ = fi21
            _ = v22
        }
    }
    if true {
        fmt.Println(v5, v11)
        v5 = (v5 ^ (v5 + 1000))
    } else {
        delete(v4, "foo bar")
        fmt.Println(v5)
        v5 = 703
    }
    var v29 uint8 = v12
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
