package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int16; F2 int64 }
func genFn0() uint64 {
    v6 := GenS1{F0: 4294967295}
    var v7 string = ""
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 uint, p2 bool) uint {
    var v8 string = ("\n\t" + "")
    p0 += p0
    v9 := []int{2147483647, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (p1 & p1)
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]bool)
    v3 := []int{42, -769, 32767}
    v4 := []string{v1, ((("\n\t" + (("hello" + ("hello" + (("hello" + v1) + ((v1 + "foo bar") + (v1 + (("\n\t" + v1) + "\n\t")))))) + v1)) + "hello") + (v1 + "hello"))}
    var v5 uint16 = 42
    if (0.1 != 0.0) {
        var v10 uint16 = (v5 ^ 0)
        v11 := GenS1{F0: 255}
        switch ("foo bar" + (((v1 + (((v1 + "foo bar") + ("a" + v1)) + ((v1 + v1) + v1))) + "a") + (v1 + v1))) {
        case "a":
            _ = len(v3)
        default:
            var v12 uint8 = 128
            _ = v12
        }
        _ = v10
        _ = v11
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            switch "\n\t" {
            case "":
                v14 := v2["a"]
                _ = v14
            case "a":
                v4 = append(v4, (("hello" + ((v1 + "a") + "a")) + "foo bar"))
            case "hello":
                v5 = v5
            }
            v15 := v2[v1]
            _ = fi13
            _ = v15
        }
        var v16 int64 = -9223372036854775808
        for fi17 := 0; fi17 < 4; fi17++ {
            if false {
                if false {
                    v18 := GenS2{F0: 1, F1: 1}
                    _ = v18
                } else {
                    switch fi17 {
                    case 32767:
                        switch (fi17 - fi17) {
                        case 435:
                            v19 := GenS3{F0: 83}
                            _ = v19
                        case 9223372036854775807:
                            v16 -= v16
                        case 176:
                            v1 = "hello"
                        default:
                            v20 := GenS3{F0: 1}
                            _ = v20
                        }
                    }
                }
                switch v1 {
                case "foo bar":
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v1 = v1
                        v22 := GenS4{F0: 127, F1: -1, F2: -9223372036854775808}
                        v16 -= (v16 & (v16 | v16))
                        _ = fi21
                        _ = v22
                    }
                default:
                    v23 := GenS1{F0: 255}
                    _ = v23
                }
                v5 = 65535
            } else {
                for fi24 := 0; fi24 < 6; fi24++ {
                    switch 42 {
                    case 914:
                        v5 -= (v5 - 0)
                    }
                    _ = fi24
                }
                for fi25 := 0; fi25 < 11; fi25++ {
                    var v26 uint = 2
                    v16 += v16
                    _ = fi25
                    _ = v26
                }
                switch "\n\t" {
                case "":
                    v16 = -475
                case "\n\t":
                    var v27 int32 = 1
                    _ = v27
                default:
                    if (!true) {
                        v16 += -9223372036854775808
                    }
                }
            }
            v1 = v1
            _ = fi17
        }
        _ = v16
    }
    switch -829 {
    case 32768:
        if len(v4) > 0 {
            v4[0] = v1
        }
    case -2147483648:
        v1 = v1
    default:
        _ = len(v2)
    }
    var v28 bool = (97 <= 2)
    v5 = (^(v5))
    fmt.Println(v1, v1, v5)
    var v29 string = ("hello" + v1)
    for fi30 := 0; fi30 < 12; fi30++ {
        v5 = v5
        _ = fi30
    }
    if (127 != 185) {
        v1 = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
