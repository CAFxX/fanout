package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 string; F2 int32 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 uint; F1 int32; F2 uint64 }
func genFn0(p0 uint32, p1 uint8) uint {
    v7 := GenS1{F0: uint(p0), F1: p1}
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    return 826
}

func main() {
    var v1 uint64 = 0
    var v2 uint = 355
    var v3 int = -2147483648
    v4 := make(map[string]string)
    var v5 int32 = 2147483647
    v6 := make(map[uint32]int)
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{(v3 * 32767), int(v1)}
        v10 := GenS2{F0: 65535}
        v11 := GenS1{F0: v2, F1: 42}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 int = v3
    var v13 uint32 = 2147483647
    switch v3 {
    case 127:
        switch "foo bar" {
        case "a":
            v1 += (^(v1))
        case "\n\t":
            v14 := []int{(v3 % ((-908) | 1))}
            _ = v14
        case "":
            v15 := GenS3{F0: -541}
            _ = v15
        default:
            v1 = (^(v1))
        }
    case -1:
        if true {
            switch v12 {
            case 42:
                v2 -= v2
            case 32767:
                var v16 string = ("" + "\n\t")
                _ = v16
            case 2147483647:
                var v17 float64 = ((3.14159 / 1.0) - 3.14159)
                _ = v17
            default:
                v1 -= v1
            }
            fmt.Println(v3, v12)
        }
    default:
        for fi18 := 0; fi18 < 1; fi18++ {
            if (!(false || true)) {
                var v19 int = v3
                switch ("" + "foo bar") {
                case "":
                    var v20 uint = v2
                    _ = v20
                }
                _ = v19
            } else {
                var v21 int16 = 0
                v22 := GenS1{F0: (v2 >> (((v2 << (((v2 - (v2 - (v2 - 65535)))) & 63))) & 63)), F1: 128}
                v23 := GenS1{F0: 18446744073709551615, F1: 128}
                _ = v21
                _ = v22
                _ = v23
            }
            v12 -= v3
            var v24 uint8 = 0
            _ = fi18
            _ = v24
        }
    }
    v25 := GenS1{F0: v2, F1: 255}
    for fi26 := 0; fi26 < 8; fi26++ {
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := GenS3{F0: 230}
            v3 = fi26
            _ = fi27
            _ = v28
        }
        v5 = v5
        _ = fi26
    }
    if (true && true) {
        switch "hello" {
        case "":
            v29 := GenS4{F0: 42, F1: "", F2: v5}
            _ = v29
        case "a":
            if false {
                if false {
                    v3 = v3
                } else {
                    for fi30 := 0; fi30 < 17; fi30++ {
                        for fi31 := 0; fi31 < 17; fi31++ {
                            v32 := GenS5{F0: v5, F1: 42}
                            delete(v4, ("foo bar" + ("foo bar" + "hello")))
                            v5 = (^(v5))
                            _ = fi31
                            _ = v32
                        }
                        _ = fi30
                    }
                }
            } else {
                v33 := GenS6{F0: -30, F1: -9223372036854775808}
                var v34 string = ("foo bar" + (("\n\t" + ("foo bar" + "\n\t")) + "hello"))
                v1 += v1
                _ = v33
                _ = v34
            }
        default:
            if ((v1 ^ 42) <= uint64(v2)) {
                v4[("a" + (("" + "foo bar") + ("" + "foo bar")))] = ("\n\t" + ("hello" + (("hello" + "\n\t") + (("a" + ("" + "foo bar")) + "\n\t"))))
            }
        }
        var v35 int16 = 842
        v36 := GenS5{F0: v5, F1: (v1 << ((255) & 63))}
        _ = v35
        _ = v36
    } else {
        v37 := GenS7{F0: 352, F1: (v5 ^ v5), F2: v1}
        if false {
            for fi38 := 0; fi38 < 4; fi38++ {
                if false {
                    for fi39 := 0; fi39 < 12; fi39++ {
                        v4["\n\t"] = ((("hello" + (("hello" + "hello") + "hello")) + "") + (("" + "hello") + ("hello" + "")))
                        var v40 int = v3
                        _ = fi39
                        _ = v40
                    }
                    switch v12 {
                    case 32767:
                        var v41 int32 = (^(v5))
                        _ = v41
                    case 9223372036854775807:
                        var v42 int16 = 42
                        _ = v42
                    }
                }
                _ = fi38
            }
        } else {
            v13 -= 2147483647
            fmt.Println(v13)
            if (127 >= 80) {
                v3 = (v12 >> (((^(v2))) & 63))
                v43 := GenS2{F0: 959}
                _ = v43
            }
        }
        _ = v37
    }
    delete(v4, ("hello" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
