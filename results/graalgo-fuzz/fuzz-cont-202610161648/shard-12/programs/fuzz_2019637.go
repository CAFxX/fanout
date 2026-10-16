package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint32 }
func genFn0(p0 uint64, p1 string, p2 string) uint32 {
    var v5 uint32 = 2147483648
    p2 = "a"
    var v6 float32 = 1e10
    var v7 uint = 635
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (v5 | 358)
}
func genFn1() int32 {
    var v8 float64 = float64(3.14159)
    v8 += 1.5
    _ = v8
    return -1
}

func main() {
    v1 := []int{-1}
    var v2 int8 = 127
    var v3 uint32 = 1
    var v4 uint = 42
    var v9 uint32 = 2147483647
    v3 -= v3
    switch (("foo bar" + "\n\t") + ("" + "")) {
    case "hello":
        _ = len(v1)
    }
    if (true && (!(0.0 == ((0.1 / 3.14159) + 1.5)))) {
        fmt.Println(v2)
        v2 = v2
        var v10 uint8 = 238
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(v9, v4)
        switch ("\n\t" + ("foo bar" + "a")) {
        case "hello":
            var v12 int16 = 372
            _ = v12
        case "a":
            v4 = 42
        case "foo bar":
            switch (-(fi11)) {
            case 191:
                switch int(v4) {
                case -512:
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v9 = v3
                        var v14 float64 = ((((float64(1.5) * float64((0.0 - ((0.0 + 1.5) * 1e10)))) - 0.1) / -2.25) + -2.25)
                        if (128 > 42) {
                            var v15 string = ((("a" + "a") + ("" + "hello")) + "\n\t")
                            var v16 uint16 = uint16(v4)
                            if false {
                                v2 -= (v2 >> ((v4) & 63))
                            }
                            _ = v15
                            _ = v16
                        }
                        _ = fi13
                        _ = v14
                    }
                case 32767:
                    for fi17 := 0; fi17 < 16; fi17++ {
                        v9 = 4294967295
                        var v18 int = fi17
                        _ = fi17
                        _ = v18
                    }
                default:
                    var v19 float64 = -2.25
                    _ = v19
                }
            }
        }
        if ((false && false) || false) {
            for fi20 := 0; fi20 < 6; fi20++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi20
            }
            if true {
                v1 = append(v1, fi11)
                for fi21 := 0; fi21 < 7; fi21++ {
                    v22 := GenS1{F0: 1, F1: (v3 / ((v3) | 1)), F2: 196}
                    _ = fi21
                    _ = v22
                }
            } else {
                fmt.Println(v4, v4, v9)
                v3 = v9
                if true {
                    var v23 int32 = 0
                    fmt.Println(fi11, v2)
                    for fi24 := 0; fi24 < 14; fi24++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi24
                    }
                    _ = v23
                }
            }
        }
        _ = fi11
    }
    if ((v9 << ((17) & 63)) < v3) {
        var v25 int64 = 0
        _ = v25
    }
    _ = len(v1)
    var v26 int64 = 9223372036854775807
    for fi27 := 0; fi27 < 3; fi27++ {
        v2 += (v2 % ((v2) | 1))
        if ((0 > genFn1()) && true) {
            var v28 uint64 = 0
            _ = v28
        } else {
            v2 = -54
            var v29 int32 = 1
            fmt.Println(v4)
            _ = v29
        }
        if ("foo bar" > "foo bar") {
            var v30 string = "a"
            v26 = 1
            _ = v30
        }
        _ = fi27
    }
    v3 = v3
    v4 += (^(v4))
    var v31 int = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v9
}
