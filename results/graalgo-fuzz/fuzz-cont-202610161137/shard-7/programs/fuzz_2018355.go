package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32) uint16 {
    p0 = (p0 / ((p0) | 1))
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 int16 = 1
        v6 := []int{(fi4 << ((134) & 63)), (^(fi4)), fi4, (fi4 << ((65535) & 63))}
        v7 := GenS1{F0: p0, F1: 127, F2: p0}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := []int{-804, 32768}
    _ = p0
    _ = v8
    return 255
}
func genFn1() int16 {
    var v9 int = 323
    fmt.Println(v9)
    _ = v9
    return 32767
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]uint32)
    var v3 uint16 = 543
    switch (("a" + ("a" + "foo bar")) + ((("hello" + "") + ("\n\t" + "hello")) + "hello")) {
    case "\n\t":
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 uint = 255
            _ = fi10
            _ = v11
        }
    case "hello":
        fmt.Println(v1, v1)
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v3, fi12, v1)
        _ = len(v2)
        _ = fi12
    }
    v13 := []int{-9223372036854775808, 32768}
    v3 = v3
    switch "\n\t" {
    case "hello":
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v3, v3)
        case "foo bar":
            var v14 int = -301
            _ = v14
        default:
            switch 127 {
            case -1:
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            case 1:
                var v15 uint8 = 128
                _ = v15
            case -9223372036854775808:
                v2[948] = 65535
            default:
                v16 := GenS1{F0: -2147483648, F1: 0, F2: 0}
                _ = v16
            }
        }
    case "":
        v17 := GenS2{F0: 1000}
        _ = v17
    }
    v1 = 0
    fmt.Println(v1)
    switch 570 {
    case -9223372036854775808:
        var v18 int32 = 2147483647
        _ = v18
    case 291:
        switch "\n\t" {
        case "":
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 int64 = 1
                for fi21 := 0; fi21 < 20; fi21++ {
                    _ = len(v13)
                    var v22 int32 = 1
                    _ = len(v2)
                    _ = fi21
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
        default:
            if true {
                if false {
                    var v23 int16 = genFn1()
                    v3 += genFn0(-2147483648)
                    _ = v23
                }
            }
        }
    default:
        v24 := GenS1{F0: -1, F1: 1, F2: 2147483647}
        _ = v24
    }
    var v25 uint64 = 1
    var v26 int16 = 427
    var v27 bool = false
    var v28 float64 = 3.14159
    var v29 bool = (!((v26 << ((255) & 63)) >= (^(v26))))
    if ((13 <= 2147483647) || (!(((18446744073709551615 == (^(v25))) && v29) || v27))) {
        for fi30 := 0; fi30 < 16; fi30++ {
            switch (fi30 | (fi30 & 32768)) {
            case 4294967295:
                if len(v13) > 0 {
                    v13[0] = (^(fi30))
                }
            case 42:
                fmt.Println(v29, v1, v26)
            case -828:
                v26 -= v26
            default:
                for fi31 := 0; fi31 < 14; fi31++ {
                    switch fi31 {
                    case 4294967295:
                        switch ("a" + ("" + "")) {
                        case "a":
                            if ((-128 != (-(v1))) || v29) {
                                v32 := GenS2{F0: v26}
                                _ = v32
                            } else {
                                delete(v2, fi30)
                            }
                        }
                    default:
                        v27 = (!(!(821 <= genFn0(631))))
                    }
                    var v33 int = int(v3)
                    v34 := GenS2{F0: 42}
                    _ = fi31
                    _ = v33
                    _ = v34
                }
            }
            _ = fi30
        }
        if v27 {
            _ = len(v13)
            fmt.Println(v28, v28)
        } else {
            var v35 uint16 = genFn0(-1)
            _ = v35
        }
        var v36 int8 = (^(v1))
        _ = v36
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        if v29 {
            var v38 int8 = (v1 + v1)
            _ = v38
        }
        if (65535 < 65535) {
            var v39 float64 = 1e10
            _ = v39
        }
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 float64 = 0.1
            v1 = (v1 | 127)
            if v29 {
                var v42 uint32 = 2147483648
                _ = v42
            }
            _ = fi40
            _ = v41
        }
        _ = fi37
    }
    var v43 uint32 = 42
    v2[32767] = 301
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v43
}
