package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int32 }
func genFn0() uint64 {
    var v5 uint32 = 1
    v5 = v5
    var v6 int = 4294967295
    _ = v5
    _ = v6
    return 359
}

func main() {
    var v1 int64 = 42
    var v2 bool = false
    v3 := make(map[int]int)
    var v4 uint64 = 4294967295
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            if (413 == 2147483647) {
                var v9 int16 = 42
                var v10 uint16 = 0
                v10 -= (v10 + (v10 | v10))
                _ = v10
                _ = v9
            }
            fmt.Println(fi8, v2)
            fmt.Println(v4, v4, v1)
            _ = fi8
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        if ((("\n\t" + ((((("\n\t" + ("foo bar" + (("hello" + (("\n\t" + "") + (((("" + (("" + (("" + "a") + "foo bar")) + "a")) + ("" + "hello")) + ("\n\t" + "foo bar")) + "a"))) + ("foo bar" + "foo bar")))) + ("" + "foo bar")) + "\n\t") + "a") + ("foo bar" + (("a" + (("\n\t" + "hello") + "a")) + "a")))) + ("a" + "")) == "a") {
            for fi12 := 0; fi12 < 16; fi12++ {
                if v2 {
                    v13 := GenS1{F0: -2147483648, F1: uint(v4), F2: 18446744073709551615}
                    var v14 float32 = 1e10
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS2{F0: 0}
                    var v16 uint = 532
                    if (2 > v4) {
                        v4 -= v4
                    } else {
                        _ = len(v3)
                    }
                    _ = len(v3)
                    v2 = v2
                    _ = v15
                    _ = v16
                }
                v4 = v4
                var v17 uint64 = (v4 & (v4 & v4))
                if (("hello" + ("foo bar" + ("" + "foo bar"))) < "foo bar") {
                    var v18 uint16 = 65535
                    switch (fi11 << ((4294967295) & 63)) {
                    case 0:
                        fmt.Println(fi11, v4)
                        v19 := v3[fi11]
                        if v2 {
                            v20 := []int{9223372036854775807, (fi12 + 32767), (fi11 + 0)}
                            _ = v20
                        } else {
                            var v21 uint = uint(v18)
                            var v22 int8 = 100
                            v19 += (v19 * -9223372036854775808)
                            _ = v21
                            _ = v22
                        }
                        _ = v19
                    }
                    v4 = (v4 & v17)
                    _ = v18
                }
                _ = fi12
                _ = v17
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                v2 = (!(!(v2 || (!false))))
                var v24 uint64 = 1
                _ = fi23
                _ = v24
            }
            v25 := GenS3{F0: fi11, F1: 216}
            _ = v25
        }
        fmt.Println(v1, fi11, v2)
        fmt.Println(v4, v2, v2)
        _ = fi11
    }
    v4 = genFn0()
    v2 = true
    v4 -= (v4 % (((v4 % ((v4) | 1))) | 1))
    fmt.Println(v1, v4)
    v4 -= genFn0()
    switch (("" + "hello") + ("hello" + "a")) {
    case "":
        var v26 int16 = -447
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v2 = (65535 >= 272)
        switch "" {
        case "foo bar":
            var v28 int32 = 0
            _ = v28
        case "":
            var v29 float64 = 1e10
            _ = v29
        }
        fmt.Println(v4, v1, fi27)
        _ = fi27
    }
    var v30 bool = v2
    for fi31 := 0; fi31 < 7; fi31++ {
        v4 -= (^(v4))
        _ = len(v3)
        _ = fi31
    }
    v4 -= genFn0()
    v32 := []int{-1, 32767, 127}
    for fi33 := 0; fi33 < 14; fi33++ {
        delete(v3, -2147483648)
        if len(v32) > 0 {
            v32[0] = 32767
        }
        _ = fi33
    }
    fmt.Println(v2, v4)
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v4
}
