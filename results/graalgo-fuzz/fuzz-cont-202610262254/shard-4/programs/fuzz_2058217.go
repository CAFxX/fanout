package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 bool; F2 uint16 }

func main() {
    var v1 uint64 = 1
    var v2 uint = 1
    var v3 int32 = -166
    var v4 int16 = -32768
    v5 := GenS1{F0: (v2 * 42), F1: 2, F2: (0.1 == 1e10)}
    v2 = 4294967295
    for fi6 := 0; fi6 < 8; fi6++ {
        v4 -= (v4 << ((1) & 63))
        v7 := []int{fi6, -303}
        fmt.Println(v2, v1, v2)
        _ = fi6
        _ = v7
    }
    if true {
        switch "a" {
        case "hello":
            var v8 uint8 = 219
            _ = v8
        }
        var v9 uint8 = uint8(v2)
        _ = v9
    }
    v10 := GenS2{F0: 80}
    v3 = v3
    v11 := GenS1{F0: v2, F1: 255, F2: true}
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            if (!(884 >= 42)) {
                v14 := []int{fi12}
                var v15 int16 = (-(v4))
                _ = v14
                _ = v15
            }
            if ((v4 << ((v2) & 63)) > -1) {
                var v16 string = "hello"
                _ = v16
            }
            v2 = (v2 % ((v2) | 1))
            _ = fi13
        }
        _ = fi12
    }
    v17 := []int{127, 770}
    v18 := GenS2{F0: 2147483647}
    switch 127 {
    case 128:
        var v19 uint16 = 301
        _ = v19
    case -575:
        if (false && false) {
            if ((^(v2)) >= (v2 / ((18446744073709551615) | 1))) {
                var v20 int = 32768
                if (((0.0 * (1.5 / 3.14159)) >= float32((3.14159 + 0.0))) || ((v2 - 255) != (v2 ^ 255))) {
                    switch ("foo bar" + (("\n\t" + ("a" + ("foo bar" + "a"))) + ("\n\t" + "a"))) {
                    case "a":
                        fmt.Println(v20)
                    }
                    v21 := GenS2{F0: 2}
                    _ = v21
                }
                fmt.Println(v20, v1)
                _ = v20
            } else {
                for fi22 := 0; fi22 < 14; fi22++ {
                    v4 = -32768
                    _ = fi22
                }
                switch 42 {
                case 438:
                    v23 := GenS2{F0: 943}
                    _ = v23
                case 2147483647:
                    v3 += (v3 % ((v3) | 1))
                }
                if false {
                    for fi24 := 0; fi24 < 14; fi24++ {
                        if (2147483647 < 992) {
                            if true {
                                v4 += 143
                            } else {
                                for fi25 := 0; fi25 < 7; fi25++ {
                                    var v26 uint8 = 1
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                            }
                        }
                        _ = fi24
                    }
                    var v27 int8 = -4
                    _ = v27
                } else {
                    switch int(v4) {
                    case 927:
                        _ = len(v17)
                    default:
                        v28 := GenS3{F0: (("" + "hello") != "foo bar"), F1: true, F2: 185}
                        _ = v28
                    }
                }
            }
        }
    case -988:
        if false {
            fmt.Println(v1)
        } else {
            v29 := GenS3{F0: (false || (42 != 0)), F1: false, F2: 255}
            v30 := GenS2{F0: 4294967295}
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 int8 = 3
        _ = fi31
        _ = v32
    }
    if (20 > v1) {
        v2 = v2
        if false {
            if len(v17) > 0 {
                v17[0] = 2147483647
            }
            var v33 uint8 = 232
            switch "foo bar" {
            case "a":
                if len(v17) > 0 {
                    v17[0] = int(v1)
                }
            default:
                v4 = (v4 + 42)
            }
            _ = v33
        }
    } else {
        if (!false) {
            var v34 int32 = (v3 | v3)
            _ = v34
        } else {
            fmt.Println(v4)
        }
        _ = len(v17)
        if len(v17) > 0 {
            v17[0] = 1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
