package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint16 }
func genFn0(p0 string) bool {
    for fi5 := 0; fi5 < 5; fi5++ {
        if false {
            v6 := []int{fi5}
            switch (fi5 % ((fi5) | 1)) {
            case -582:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case 1:
                for fi7 := 0; fi7 < 7; fi7++ {
                    v6 = append(v6, 992)
                    v8 := GenS1{F0: 18446744073709551615, F1: 42}
                    var v9 string = p0
                    v9 = ""
                    v10 := GenS1{F0: 255, F1: 65535}
                    if true {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v11 uint16 = 65535
                        _ = v11
                    } else {
                        var v12 uint64 = 1
                        v13 := GenS1{F0: 255, F1: 2}
                        _ = v12
                        _ = v13
                    }
                    v9 = ""
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 uint32 = 4294967295
                _ = fi14
                _ = v15
            }
            _ = v6
        } else {
            p0 = ""
            p0 = ("a" + p0)
            if (false || true) {
                p0 = (p0 + p0)
                switch (fi5 >> ((65535) & 63)) {
                case 0:
                    v16 := GenS1{F0: 669, F1: 42}
                    _ = v16
                case 1:
                    v17 := []int{(-(fi5)), (fi5 - fi5)}
                    _ = v17
                case 32767:
                    p0 = "foo bar"
                default:
                    p0 = p0
                }
            } else {
                var v18 uint16 = 0
                v19 := []int{}
                p0 = p0
                _ = v18
                _ = v19
            }
        }
        var v20 uint8 = 226
        if true {
            for fi21 := 0; fi21 < 10; fi21++ {
                v22 := []int{int(v20), fi5, fi5}
                switch p0 {
                case "":
                    v20 -= v20
                case "\n\t":
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                case "hello":
                    var v23 int8 = 9
                    _ = v23
                default:
                    switch p0 {
                    case "":
                        v24 := GenS2{F0: -155, F1: 42, F2: 65535}
                        _ = v24
                    case "a":
                        var v25 int32 = 42
                        _ = v25
                    }
                }
                _ = fi21
                _ = v22
            }
            v26 := []int{}
            _ = v26
        }
        _ = fi5
        _ = v20
    }
    var v27 uint32 = 42
    _ = p0
    _ = v27
    return (v27 == (^(v27)))
}
func genFn1() string {
    if (1 == 0) {
        var v28 int32 = 2147483647
        v28 = 1
        fmt.Println(v28)
        _ = v28
    }
    v29 := GenS1{F0: 0, F1: 42}
    _ = v29
    return "\n\t"
}

func main() {
    v1 := []uint32{65535, 4294967295, 1, 911}
    var v2 string = (("a" + "") + "a")
    var v3 int16 = 42
    var v4 string = v2
    for fi30 := 0; fi30 < 19; fi30++ {
        var v31 uint8 = 2
        for fi32 := 0; fi32 < 8; fi32++ {
            v31 += (^(v31))
            fmt.Println(v2, fi32, v3)
            _ = fi32
        }
        _ = fi30
        _ = v31
    }
    v33 := GenS1{F0: 42, F1: 0}
    fmt.Println(v2, v2, v3)
    switch (v2 + v2) {
    case "":
        v2 = (v2 + v4)
    case "a":
        var v34 float64 = (0.0 + 1e10)
        _ = v34
    default:
        fmt.Println(v4, v3)
    }
    var v35 int = 128
    v35 = v35
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v35 = (v35 ^ -2147483648)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v4
}
