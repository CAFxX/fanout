package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int32 }
func genFn0() int64 {
    v5 := []int{984}
    var v6 int = 127
    for fi7 := 0; fi7 < 4; fi7++ {
        if false {
            switch "a" {
            case "":
                fmt.Println(v6, v6)
            case "\n\t":
                v6 -= -62
            default:
                switch "foo bar" {
                case "\n\t":
                    v6 = (v6 * 42)
                case "a":
                    v8 := GenS1{F0: 42, F1: 2147483647}
                    _ = v8
                default:
                    if (!(v6 == (-(fi7)))) {
                        v9 := GenS2{F0: ("\n\t" + ""), F1: false}
                        switch ("\n\t" + "") {
                        case "":
                            var v10 uint8 = 127
                            _ = v10
                        case "\n\t":
                            var v11 bool = (float32(1e10) != (float32(1.5) + 0.0))
                            _ = v11
                        default:
                            var v12 int8 = 61
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 10; fi13++ {
                            v14 := GenS1{F0: 4294967295, F1: 42}
                            v6 = v6
                            v15 := GenS1{F0: 1, F1: 144}
                            _ = len(v5)
                            v6 = (fi7 - fi13)
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        _ = v9
                    }
                }
            }
            v16 := GenS3{F0: -1}
            _ = v16
        } else {
            if ((1.5 + 1.5) >= 0.0) {
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 uint16 = 1
                    _ = fi17
                    _ = v18
                }
                if false {
                    var v19 int8 = 4
                    v20 := GenS2{F0: "a", F1: true}
                    if (42 < 0) {
                        var v21 int16 = 271
                        _ = v21
                    } else {
                        v6 -= (fi7 / ((32767) | 1))
                    }
                    _ = v19
                    _ = v20
                } else {
                    switch fi7 {
                    case 127:
                        v22 := GenS1{F0: 438, F1: 171}
                        _ = v22
                    case 0:
                        if (32767 < 255) {
                            v6 += (v6 << ((255) & 63))
                            v6 = fi7
                        }
                    default:
                        v23 := GenS1{F0: 396, F1: 65535}
                        _ = v23
                    }
                    if len(v5) > 0 {
                        v5[0] = 87
                    }
                    fmt.Println(fi7, fi7)
                }
            } else {
                var v24 uint32 = 65535
                v6 += -2147483648
                var v25 uint8 = 113
                _ = v24
                _ = v25
            }
        }
        if len(v5) > 0 {
            v5[0] = -969
        }
        switch ("\n\t" + (("foo bar" + "foo bar") + "")) {
        case "a":
            var v26 bool = ((!(0 != 128)) || true)
            _ = v26
        }
        _ = fi7
    }
    _ = len(v5)
    _ = v5
    _ = v6
    return -854
}
func genFn1() int16 {
    switch "" {
    case "":
        var v27 uint64 = 2
        _ = v27
    case "\n\t":
        v28 := []int{32767, 623, 0, 647}
        _ = v28
    default:
        var v29 int64 = 9223372036854775807
        _ = v29
    }
    v30 := []int{}
    _ = v30
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "a"
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]int)
    switch -797 {
    case 1:
        v31 := []int{}
        _ = v31
    }
    v32 := GenS3{F0: 235}
    for fi33 := 0; fi33 < 14; fi33++ {
        v34 := GenS2{F0: (("a" + "foo bar") + "hello"), F1: (v2 >= v2)}
        _ = fi33
        _ = v34
    }
    switch "" {
    case "":
        var v35 int32 = -1
        _ = v35
    }
    v36 := GenS2{F0: (((v2 + ("a" + (v2 + (v2 + "\n\t")))) + (v2 + (v2 + ("\n\t" + ("a" + v2))))) + v2), F1: false}
    var v37 string = v2
    var v38 int = -620
    v3 = (v3 & 998)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
