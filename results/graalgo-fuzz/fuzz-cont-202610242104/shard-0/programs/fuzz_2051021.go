package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 bool; F1 int32 }
func genFn0(p0 int32, p1 uint16) int {
    if ((32767 <= 0) || (!false)) {
        fmt.Println(p0, p1)
    } else {
        v5 := []int{9223372036854775807, -9223372036854775808, -2147483648, 127}
        p1 += 255
        p0 -= 1
        _ = v5
    }
    fmt.Println(p0, p1)
    p0 += (p0 & p0)
    var v6 uint16 = (^(p1))
    _ = p0
    _ = p1
    _ = v6
    return 128
}

func main() {
    v1 := []uint32{2147483648, 2147483648, 523}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 1
    var v4 int = 613
    if len(v1) > 0 {
        v1[0] = 2
    }
    if (false && false) {
        var v7 string = "a"
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 string = v7
            if (18446744073709551615 < 0) {
                var v10 uint = 0
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        if true {
            switch (v4 | (v4 % ((v4) | 1))) {
            case 0:
                v7 = (v7 + v7)
            case -9223372036854775808:
                for fi11 := 0; fi11 < 7; fi11++ {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        if (v4 > 653) {
                            var v13 int32 = -157
                            if false {
                                var v14 uint16 = 458
                                v3 -= (v3 / (((v3 & v3)) | 1))
                                _ = v14
                            } else {
                                v4 = v4
                                v7 = v7
                                v15 := GenS1{F0: 2, F1: 2}
                                _ = v15
                            }
                            var v16 int = fi12
                            _ = v13
                            _ = v16
                        }
                        _ = fi12
                    }
                    v17 := GenS2{F0: true, F1: 42}
                    v7 = (v7 + ("foo bar" + ""))
                    _ = fi11
                    _ = v17
                }
            case 32768:
                v3 = (v3 >> ((65535) & 63))
            }
        }
        _ = v7
    }
    v3 = v3
    if (false || true) {
        v18 := GenS1{F0: 32767, F1: 2}
        _ = v18
    }
    v19 := GenS1{F0: 774, F1: 234}
    if false {
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 uint32 = 760
            v3 += 2
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 uint32 = 1
            v3 = 18446744073709551615
            v24 := v2[v23]
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    _ = len(v2)
    switch ((("foo bar" + "a") + "") + ("" + "foo bar")) {
    case "foo bar":
        if true {
            var v25 float64 = 0.0
            _ = v25
        } else {
            _ = len(v2)
        }
    case "hello":
        switch v4 {
        case -9223372036854775808:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 32768:
            v3 = 972
        }
    default:
        fmt.Println(v3, v3)
    }
    if false {
        if (42 >= 2) {
            var v26 float64 = 0.0
            for fi27 := 0; fi27 < 5; fi27++ {
                if (false || ((42 < -124) || (!true))) {
                    if (true || true) {
                        v26 = v26
                        v28 := GenS1{F0: 932, F1: 2}
                        var v29 uint16 = 255
                        _ = v28
                        _ = v29
                    }
                    v30 := GenS1{F0: 0, F1: 42}
                    for fi31 := 0; fi31 < 20; fi31++ {
                        if false {
                            var v32 int = v4
                            _ = len(v2)
                            v4 += 127
                            _ = v32
                        }
                        _ = fi31
                    }
                    _ = v30
                }
                var v33 int = fi27
                _ = fi27
                _ = v33
            }
            v3 = (v3 % ((v3) | 1))
            _ = v26
        }
    }
    fmt.Println(v4, v4)
    var v34 string = "hello"
    for fi35 := 0; fi35 < 15; fi35++ {
        var v36 uint32 = 42
        _ = fi35
        _ = v36
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
