package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int32; F2 string }
func genFn0(p0 int, p1 int8, p2 uint16) int32 {
    p0 = (p0 | 4294967295)
    v4 := GenS1{F0: 436, F1: 484}
    p2 -= (^(p2))
    var v5 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return int32(p0)
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 int32 = 968
    var v6 int32 = genFn0(-410, 127, 32767)
    fmt.Println(v6, v6)
    switch 0 {
    case -577:
        v6 = genFn0(128, 100, 1)
    case 2147483647:
        fmt.Println(v3, v6)
    }
    switch "foo bar" {
    case "hello":
        v3 += (v6 + 531)
    case "":
        switch 42 {
        case -307:
            v7 := []int{2147483647, -9223372036854775808}
            _ = v7
        case -2147483648:
            v8 := GenS2{F0: 255, F1: 42}
            _ = v8
        default:
            if false {
                var v9 string = "a"
                fmt.Println(v9, v3)
                var v10 int8 = 100
                _ = v10
                _ = v9
            }
        }
    default:
        v11 := GenS1{F0: 4294967295, F1: 585}
        _ = v11
    }
    fmt.Println(v6)
    delete(v1, "a")
    for fi12 := 0; fi12 < 20; fi12++ {
        switch "hello" {
        case "\n\t":
            fmt.Println(fi12, v6)
        case "foo bar":
            if (-9223372036854775808 < 117) {
                v13 := GenS3{F0: 2147483647}
                v3 += (v6 % (((v6 % ((v6) | 1))) | 1))
                if false {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        v3 -= (v3 | -2147483648)
                        v3 -= v6
                        v6 += (v3 >> ((218) & 63))
                        _ = fi14
                    }
                    v3 -= genFn0(fi12, 67, 32767)
                    switch (-(fi12)) {
                    case -48:
                        v15 := GenS2{F0: 1, F1: -966}
                        _ = v15
                    }
                }
                _ = v13
            } else {
                var v16 int8 = -1
                _ = v16
            }
        }
        if (155 > 42) {
            switch -1 {
            case 2147483647:
                var v17 bool = (-1 >= (v6 - 552))
                _ = v17
            case 582:
                v6 = (^(v3))
            case 802:
                fmt.Println(v6)
            }
            if (true && (65535 != 255)) {
                for fi18 := 0; fi18 < 20; fi18++ {
                    v19 := []int{fi12}
                    fmt.Println(fi12, fi12)
                    _ = fi18
                    _ = v19
                }
                fmt.Println(v6)
                v6 = -264
            }
        }
        var v20 int = fi12
        _ = fi12
        _ = v20
    }
    switch 4294967295 {
    case -1:
        v6 = (v3 + genFn0(1, 0, 32767))
    case -670:
        for fi21 := 0; fi21 < 20; fi21++ {
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 uint16 = 32767
                _ = fi22
                _ = v23
            }
            var v24 uint8 = 168
            _ = fi21
            _ = v24
        }
    case 522:
        v3 += v3
    default:
        var v25 int64 = -1
        _ = v25
    }
    v6 -= -2147483648
    for fi26 := 0; fi26 < 14; fi26++ {
        if ((false && false) && true) {
            fmt.Println(fi26, v6, v3)
            for fi27 := 0; fi27 < 20; fi27++ {
                v28 := []int{300, -9223372036854775808, (fi27 / ((811) | 1)), -1}
                v3 += v6
                if false {
                    fmt.Println(fi27)
                }
                if true {
                    v6 = v6
                    v29 := GenS4{F0: 100, F1: v3, F2: "hello"}
                    fmt.Println(fi27, v6, v3)
                    _ = v29
                } else {
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                    var v30 uint16 = 467
                    _ = v30
                }
                fmt.Println(fi27)
                _ = fi27
                _ = v28
            }
            for fi31 := 0; fi31 < 5; fi31++ {
                v32 := v1["a"]
                _ = fi31
                _ = v32
            }
        }
        _ = fi26
    }
    var v33 uint32 = 4294967295
    v34 := GenS2{F0: 65535, F1: 1}
    var v35 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v6
}
