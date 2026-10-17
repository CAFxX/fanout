package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint8; F2 int32 }

func main() {
    var v1 int16 = 0
    v2 := []string{"hello"}
    v3 := make(map[uint32]string)
    v4 := []int{}
    var v5 string = ""
    v6 := GenS1{F0: 65535, F1: 9223372036854775807, F2: -125}
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint32 = 65535
            var v10 int64 = 9223372036854775807
            _ = fi8
            _ = v10
            _ = v9
        }
        switch "foo bar" {
        case "":
            if len(v4) > 0 {
                v4[0] = (fi7 % ((fi7) | 1))
            }
        case "hello":
            for fi11 := 0; fi11 < 16; fi11++ {
                if len(v2) > 0 {
                    v2[0] = (v5 + ((((v5 + "foo bar") + (((((v5 + v5) + v5) + v5) + "a") + ((("\n\t" + v5) + "foo bar") + v5))) + v5) + v5))
                }
                switch 42 {
                case -9223372036854775808:
                    v1 = 32767
                case 42:
                    fmt.Println(fi11, fi11)
                }
                _ = fi11
            }
        default:
            switch (fi7 + fi7) {
            case 0:
                var v12 int = 791
                _ = v12
            case -620:
                fmt.Println(v5, v5, v5)
            case -982:
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    for fi13 := 0; fi13 < 15; fi13++ {
                        switch fi13 {
                        case -9223372036854775808:
                            v1 -= (v1 & 42)
                            v3[2] = ("\n\t" + v5)
                        case 32767:
                            v1 += v1
                        case -2147483648:
                            v5 = "foo bar"
                            v2 = append(v2, "hello")
                            v1 -= (v1 | -1)
                        default:
                            switch "hello" {
                            case "hello":
                                _ = len(v3)
                            case "":
                                if (42 <= 4294967295) {
                                    var v14 bool = (true && ((-632 < (fi13 ^ -768)) && (2147483647 != 65535)))
                                    _ = len(v4)
                                    _ = len(v3)
                                    _ = v14
                                }
                            case "foo bar":
                                fmt.Println(fi13)
                            }
                        }
                        _ = fi13
                    }
                }
            }
        }
        _ = fi7
    }
    var v15 int64 = 1
    v1 = (v1 + -254)
    if len(v2) > 0 {
        v2[0] = v5
    }
    _ = len(v3)
    fmt.Println(v15, v15)
    if false {
        v16 := GenS2{F0: 42, F1: 270}
        v17 := GenS2{F0: 273, F1: 42}
        for fi18 := 0; fi18 < 16; fi18++ {
            delete(v3, 65535)
            if (119 >= 255) {
                v19 := GenS3{F0: 42, F1: 255, F2: -2147483648}
                fmt.Println(v5)
                _ = v19
            } else {
                v5 = v5
                var v20 int16 = (^(v1))
                _ = v20
            }
            if (!(!(1 != 4294967295))) {
                var v21 int16 = (v1 * v1)
                fmt.Println(v5, v15, v15)
                _ = v21
            }
            _ = fi18
        }
        _ = v16
        _ = v17
    } else {
        v1 = v1
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        v1 = v1
        _ = fi22
    }
    var v23 int = -1
    switch v23 {
    case 127:
        var v24 int = v23
        _ = v24
    case 42:
        v23 += (-(v23))
    case -290:
        var v25 uint64 = 927
        _ = v25
    }
    v15 = (v15 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
