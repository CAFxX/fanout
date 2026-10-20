package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 bool }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint; F1 string }
func genFn0(p0 int32, p1 uint8) bool {
    p1 = (p1 << ((0) & 63))
    p0 += (^(p0))
    _ = p0
    _ = p1
    return (817 > 764)
}

func main() {
    var v1 int8 = 127
    var v2 uint16 = 433
    v3 := make(map[uint32]string)
    var v4 float32 = -2.25
    v2 += (v2 ^ 2)
    var v5 int8 = 1
    v6 := []int{-1}
    v6 = append(v6, 706)
    var v7 uint8 = 127
    for fi8 := 0; fi8 < 6; fi8++ {
        if true {
            var v9 uint8 = (^(v7))
            v5 = (v5 >> ((981) & 63))
            v4 = -2.25
            _ = v9
        } else {
            v2 -= v2
            if (-604 < -699) {
                if true {
                    v1 -= (v5 % (((v5 % ((0) | 1))) | 1))
                    v2 -= v2
                } else {
                    switch 128 {
                    case -9223372036854775808:
                        v2 = (v2 * (v2 | 1))
                    case -717:
                        if (!(42 < 0)) {
                            if false {
                                var v10 float32 = (float32(((0.0 * -2.25) - float64(3.14159))) / float32(3.14159))
                                _ = v10
                            }
                            var v11 uint16 = (^(v2))
                            switch -95 {
                            case -226:
                                switch (fi8 / ((fi8) | 1)) {
                                case 0:
                                    var v12 uint = 42
                                    _ = v12
                                }
                            }
                            _ = v11
                        } else {
                            for fi13 := 0; fi13 < 17; fi13++ {
                                var v14 bool = false
                                _ = fi13
                                _ = v14
                            }
                            var v15 float64 = (0.0 - 1e10)
                            _ = v15
                        }
                    }
                    v16 := GenS1{F0: -1, F1: 0, F2: (((-9223372036854775808 <= -476) || (!false)) || (194 > 65535))}
                    _ = v16
                }
                v4 = ((float32(3.14159) - 3.14159) + v4)
                var v17 int32 = 42
                _ = v17
            }
        }
        fmt.Println(v2)
        v5 = (-(v5))
        _ = fi8
    }
    if true {
        switch (("foo bar" + "") + "foo bar") {
        case "hello":
            if (2 != 32767) {
                switch ("foo bar" + "a") {
                case "foo bar":
                    var v18 string = ("\n\t" + "")
                    _ = v18
                case "":
                    v7 -= (v7 | (v7 + v7))
                }
            } else {
                for fi19 := 0; fi19 < 6; fi19++ {
                    switch ("a" + "\n\t") {
                    case "hello":
                        if len(v6) > 0 {
                            v6[0] = (fi19 - fi19)
                        }
                    }
                    _ = fi19
                }
            }
        case "":
            if false {
                v2 -= v2
            }
        case "a":
            switch -2147483648 {
            case -43:
                var v20 uint8 = (v7 + 36)
                _ = v20
            }
        }
    } else {
        v4 = v4
        v21 := GenS2{F0: true, F1: -432}
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := GenS3{F0: 18446744073709551615, F1: ("hello" + "foo bar")}
        _ = fi22
        _ = v23
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
