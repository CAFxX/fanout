package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint; F1 bool }
func genFn0(p0 uint16) int16 {
    fmt.Println(p0)
    var v7 int = 0
    var v8 int = (v7 % (((v7 & (v7 + 0))) | 1))
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1() uint8 {
    var v9 bool = true
    v9 = v9
    switch -967 {
    case -1:
        if (0 <= 2) {
            v10 := []int{32767, -1}
            _ = v10
        } else {
            var v11 uint = 1
            v11 = v11
            _ = v11
        }
    }
    var v12 float64 = -2.25
    var v13 uint = 65535
    _ = v12
    _ = v13
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]bool)
    var v3 float32 = 0.0
    v4 := make(map[int]bool)
    var v5 uint64 = 2
    var v6 float32 = (3.14159 - (-2.25 + float32(1e10)))
    v14 := GenS1{F0: (v5 ^ v5), F1: 953}
    delete(v2, 1)
    v15 := GenS1{F0: v5, F1: 48}
    var v16 uint16 = 0
    for fi17 := 0; fi17 < 13; fi17++ {
        v3 = 0.0
        for fi18 := 0; fi18 < 7; fi18++ {
            switch "hello" {
            case "\n\t":
                var v19 int64 = 0
                _ = v19
            case "a":
                v3 = v6
            default:
                v5 = 2
            }
            if false {
                v20 := []int{}
                _ = v20
            }
            v1 = 65535
            v21 := GenS1{F0: 0, F1: uint(v1)}
            _ = fi18
            _ = v21
        }
        _ = fi17
    }
    if (1 <= 0) {
        var v22 uint = 248
        fmt.Println(v6)
        _ = v22
    } else {
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := GenS1{F0: (v5 % ((v5) | 1)), F1: 255}
            if (v16 <= 42) {
                v6 -= v3
                v25 := []int{fi23, (fi23 | fi23), int(v1)}
                _ = v25
            } else {
                var v26 int64 = int64(v1)
                for fi27 := 0; fi27 < 6; fi27++ {
                    if false {
                        var v28 string = ("hello" + (("hello" + "\n\t") + "\n\t"))
                        delete(v2, 0)
                        fmt.Println(v5, v16, v5)
                        _ = v28
                    }
                    _ = fi27
                }
                _ = v26
            }
            switch (fi23 * int(v16)) {
            case 32767:
                var v29 int64 = 0
                _ = v29
            case 9223372036854775807:
                v3 = float32(0.1)
            }
            _ = fi23
            _ = v24
        }
    }
    var v30 int32 = 1
    if false {
        v31 := GenS2{F0: false, F1: 42, F2: 59}
        _ = v31
    }
    v6 = float32(1.5)
    var v32 uint16 = uint16(v5)
    var v33 uint8 = genFn1()
    var v34 int8 = -9
    switch -2147483648 {
    case 4294967295:
        v35 := []int{int(v32)}
        _ = v35
    case 127:
        for fi36 := 0; fi36 < 16; fi36++ {
            switch (-(fi36)) {
            case 128:
                var v37 string = ("\n\t" + ("\n\t" + "hello"))
                _ = v37
            case 700:
                var v38 int16 = 0
                _ = v38
            default:
                if false {
                    if true {
                        _ = len(v4)
                    } else {
                        var v39 string = ((("" + ("foo bar" + ("" + "foo bar"))) + ("foo bar" + "a")) + ("\n\t" + ""))
                        for fi40 := 0; fi40 < 12; fi40++ {
                            v41 := v4[32767]
                            v5 += v5
                            _ = fi40
                            _ = v41
                        }
                        _ = v39
                    }
                    v1 = 86
                    if (!true) {
                        for fi42 := 0; fi42 < 12; fi42++ {
                            v43 := v2[v1]
                            _ = fi42
                            _ = v43
                        }
                        fmt.Println(v33, v3)
                    } else {
                        var v44 uint32 = (v1 * (^(v1)))
                        switch -14 {
                        case 1:
                            v45 := GenS3{F0: 127, F1: v44, F2: -422}
                            _ = v45
                        }
                        _ = v44
                    }
                }
            }
            v46 := GenS4{F0: 0, F1: (-9223372036854775808 > 1)}
            v47 := []int{(fi36 << ((65535) & 63))}
            _ = fi36
            _ = v46
            _ = v47
        }
    case -1:
        v3 -= v6
    default:
        switch ("" + (("\n\t" + ("a" + "a")) + (("" + "a") + ("a" + ("" + "a"))))) {
        case "hello":
            var v48 bool = true
            _ = v48
        case "\n\t":
            fmt.Println(v1, v34, v30)
        case "a":
            v2[915] = (!(!(!false)))
        }
    }
    fmt.Println(v5, v16, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
