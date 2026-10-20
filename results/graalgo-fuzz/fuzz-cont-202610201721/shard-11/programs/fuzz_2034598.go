package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint16; F1 bool; F2 int }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 string; F1 int8 }
func genFn0(p0 int8, p1 uint32) int32 {
    p0 += p0
    if true {
        switch (("hello" + "hello") + "hello") {
        case "":
            if (!true) {
                var v4 int32 = -987
                var v5 int = 127
                if true {
                    p0 += (^(p0))
                }
                _ = v4
                _ = v5
            } else {
                v6 := []int{291, 1}
                p1 = p1
                if (uint(p1) <= 65535) {
                    if true {
                        if ((("foo bar" + ("hello" + "")) + "foo bar") <= ("foo bar" + "foo bar")) {
                            p0 += 127
                            _ = len(v6)
                        } else {
                            v7 := GenS1{F0: 0}
                            var v8 uint = 2
                            _ = v7
                            _ = v8
                        }
                        v9 := GenS1{F0: 911}
                        _ = len(v6)
                        _ = v9
                    } else {
                        var v10 uint64 = 1
                        p0 = p0
                        for fi11 := 0; fi11 < 17; fi11++ {
                            if false {
                                var v12 bool = (1 <= 65535)
                                var v13 uint = 255
                                _ = len(v6)
                                _ = v12
                                _ = v13
                            }
                            for fi14 := 0; fi14 < 15; fi14++ {
                                _ = len(v6)
                                _ = fi14
                            }
                            for fi15 := 0; fi15 < 9; fi15++ {
                                fmt.Println(fi15)
                                _ = fi15
                            }
                            _ = fi11
                        }
                        _ = v10
                    }
                } else {
                    if false {
                        switch 372 {
                        case 128:
                            var v16 bool = true
                            _ = v16
                        default:
                            var v17 int = -233
                            _ = v17
                        }
                        p1 = 240
                        p0 -= (p0 ^ (p0 / ((-1) | 1)))
                    }
                }
                _ = v6
            }
        }
        var v18 int32 = 42
        if (!("\n\t" <= "a")) {
            var v19 uint64 = 4294967295
            v19 = v19
            v19 = (v19 + uint64(p1))
            _ = v19
        } else {
            switch "hello" {
            case "\n\t":
                if true {
                    if (172 != 1) {
                        var v20 int = 2147483647
                        _ = v20
                    }
                }
            }
            if true {
                p1 = (p1 >> ((110) & 63))
            } else {
                var v21 bool = true
                _ = v21
            }
        }
        _ = v18
    }
    var v22 int = 4294967295
    v22 += v22
    _ = p0
    _ = p1
    _ = v22
    return 42
}
func genFn1(p0 int64) uint {
    v23 := GenS1{F0: 2}
    v24 := GenS2{F0: "", F1: 0}
    v25 := []int{}
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    _ = p0
    _ = v23
    _ = v24
    _ = v25
    return 932
}

func main() {
    var v1 uint16 = 418
    var v2 uint16 = v1
    var v3 uint32 = 65535
    v26 := GenS1{F0: 968}
    v27 := []int{480, 0, -2147483648}
    if true {
        switch 7 {
        case -9223372036854775808:
            switch "a" {
            case "a":
                fmt.Println(v3, v1, v1)
            }
        case 32767:
            v28 := GenS3{F0: 730, F1: 2}
            _ = v28
        }
    }
    var v29 int32 = -2147483648
    for fi30 := 0; fi30 < 16; fi30++ {
        for fi31 := 0; fi31 < 17; fi31++ {
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            var v32 float32 = ((float32(1e10) * -2.25) / 0.1)
            _ = fi31
            _ = v32
        }
        var v33 int16 = -294
        _ = fi30
        _ = v33
    }
    v29 -= (^(v29))
    v1 = v2
    if (!((("" + "") + ("foo bar" + "a")) <= (("\n\t" + "\n\t") + "a"))) {
        if (0.0 <= float32(((3.14159 / 0.1) * -2.25))) {
            v34 := GenS4{F0: v2, F1: ((!(uint64(v1) > 99)) || ((42 >= 1) && true)), F2: 9223372036854775807}
            v35 := GenS5{F0: 479, F1: v3, F2: v1}
            fmt.Println(v1)
            _ = v34
            _ = v35
        }
        for fi36 := 0; fi36 < 16; fi36++ {
            fmt.Println(v3, v1, v3)
            _ = fi36
        }
        v3 += v3
    }
    switch "a" {
    case "foo bar":
        v29 += (v29 & v29)
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        v38 := GenS6{F0: 939, F1: 65535}
        if (100 != 0) {
            v39 := GenS7{F0: ("foo bar" + "hello"), F1: 80}
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
}
