package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint16; F2 bool }
type GenS5 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]int)
    var v3 bool = (!((18446744073709551615 >= 2) || false))
    v4 := make(map[string]string)
    v5 := []int{127, 0, 0, 42}
    if (!v3) {
        for fi6 := 0; fi6 < 10; fi6++ {
            v1 = (v1 * (v1 & (^(v1))))
            v7 := GenS1{F0: 179, F1: v1, F2: 255}
            for fi8 := 0; fi8 < 16; fi8++ {
                if v3 {
                    if true {
                        v9 := GenS2{F0: 0}
                        if (14 != 0) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if ("a" == "hello") {
                                v3 = (!v3)
                                v10 := GenS2{F0: 42}
                                _ = v10
                            }
                        }
                        var v11 uint = 21
                        _ = v11
                        _ = v9
                    } else {
                        v12 := GenS1{F0: 312, F1: v1, F2: 255}
                        _ = v12
                    }
                    _ = len(v4)
                    fmt.Println(fi8, v1, v1)
                }
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
    } else {
        v13 := GenS3{F0: 62, F1: 1, F2: 127}
        if ((!(!(!(v3 || false)))) || (65535 > 541)) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v1, v3)
            var v14 int8 = 8
            _ = v14
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            if true {
                var v16 string = ""
                _ = v16
            } else {
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 float32 = 0.1
                    _ = fi17
                    _ = v18
                }
                var v19 uint8 = 42
                _ = v19
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 int8 = 12
                v3 = (!(791 != 65535))
                v3 = true
                _ = fi20
                _ = v21
            }
            _ = fi15
        }
        _ = v13
    }
    var v22 uint = 255
    v3 = v3
    _ = len(v2)
    v23 := GenS1{F0: v22, F1: v1, F2: v22}
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 uint16 = 2
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi26 := 0; fi26 < 14; fi26++ {
            v5 = append(v5, fi24)
            var v27 uint16 = (v25 & (v25 - (v25 | 1)))
            var v28 uint = v22
            if v3 {
                v29 := GenS4{F0: v1, F1: v27, F2: (0 < 0)}
                v30 := GenS1{F0: (v22 | (v28 & (v28 | (v28 - v22)))), F1: 4294967295, F2: uint(v1)}
                v27 += (v25 << ((546) & 63))
                _ = v29
                _ = v30
            } else {
                v31 := GenS5{F0: 128, F1: -32768}
                _ = v31
            }
            var v32 int32 = 0
            _ = fi26
            _ = v27
            _ = v28
            _ = v32
        }
        _ = fi24
        _ = v25
    }
    switch -691 {
    case 427:
        if len(v5) > 0 {
            v5[0] = 382
        }
    }
    v1 += v1
    switch ("foo bar" + ("foo bar" + ((((("foo bar" + "") + "a") + ("\n\t" + ("\n\t" + ("a" + "a")))) + "hello") + "a"))) {
    case "foo bar":
        for fi33 := 0; fi33 < 3; fi33++ {
            delete(v2, "\n\t")
            v22 = v22
            _ = fi33
        }
    case "":
        var v34 uint16 = 32767
        _ = v34
    case "\n\t":
        fmt.Println(v1)
    }
    fmt.Println(v3, v22, v3)
    v35 := GenS2{F0: 1}
    for fi36 := 0; fi36 < 5; fi36++ {
        for fi37 := 0; fi37 < 13; fi37++ {
            v22 = v22
            fmt.Println(fi37, v1, fi36)
            var v38 float32 = (1e10 - 1.5)
            _ = fi37
            _ = v38
        }
        fmt.Println(fi36)
        _ = fi36
    }
    v1 = (v1 >> ((v22) & 63))
    var v39 float32 = 1e10
    var v40 float32 = v39
    v40 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
