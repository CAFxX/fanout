package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }
func genFn0(p0 int16) string {
    var v6 uint32 = 4294967295
    if (true || true) {
        v7 := []int{-2147483648}
        var v8 uint32 = (v6 % ((v6) | 1))
        var v9 float32 = (0.1 * 1e10)
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 int = -2147483648
    _ = p0
    _ = v10
    _ = v6
    return ("hello" + "")
}
func genFn1(p0 int32, p1 uint16, p2 int32) uint {
    v11 := []int{}
    for fi12 := 0; fi12 < 1; fi12++ {
        switch ("" + "foo bar") {
        case "hello":
            p0 -= 958
        default:
            if ((fi12 | (fi12 & (fi12 ^ -9223372036854775808))) == fi12) {
                p2 -= (p0 & p0)
                fmt.Println(p1, p0)
            }
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 uint8 = 188
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 969
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(v3, v5, v3)
        for fi14 := 0; fi14 < 9; fi14++ {
            if (true && (2 >= 954)) {
                v5 -= v5
                switch fi14 {
                case 42:
                    v5 = 859
                }
            }
            v5 = v5
            v15 := []int{fi14, fi13}
            _ = fi14
            _ = v15
        }
        v16 := []int{fi13, 127, fi13, (fi13 * fi13)}
        _ = fi13
        _ = v16
    }
    if (((v3 % ((v3) | 1)) <= 128) || false) {
        if false {
            var v17 uint = 42
            v17 += (^(v17))
            switch 4294967295 {
            case 9223372036854775807:
                if false {
                    var v18 int = 247
                    if (("foo bar" >= "foo bar") && (32767 > int16(v3))) {
                        v3 += (v3 + v3)
                    }
                    var v19 uint16 = 42
                    _ = v18
                    _ = v19
                } else {
                    var v20 uint16 = 959
                    v5 = 2147483647
                    var v21 int64 = -9223372036854775808
                    _ = v20
                    _ = v21
                }
            default:
                for fi22 := 0; fi22 < 11; fi22++ {
                    v23 := GenS1{F0: "", F1: ("" + ("hello" + "\n\t")), F2: "\n\t"}
                    v24 := []int{}
                    if (1 == 1) {
                        switch "a" {
                        case "foo bar":
                            v17 = 1
                        case "hello":
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                        }
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            _ = v17
        }
        var v25 string = "foo bar"
        fmt.Println(v5)
        _ = v25
    } else {
        v26 := GenS1{F0: ("hello" + (("foo bar" + "") + "hello")), F1: ("hello" + (("\n\t" + ("" + "\n\t")) + "")), F2: ("\n\t" + ((("" + ("" + ("foo bar" + "a"))) + ("" + "a")) + ""))}
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 uint16 = 504
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    switch -9223372036854775808 {
    case 32768:
        v5 = 2147483648
    }
    var v29 int = 4294967295
    v3 += (v3 / ((2) | 1))
    v5 -= (v5 << ((818) & 63))
    fmt.Println(v3, v5)
    var v30 int32 = 311
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
