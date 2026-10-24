package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 string }
func genFn0(p0 int8, p1 uint8, p2 uint) string {
    v7 := GenS1{F0: "foo bar"}
    var v8 int16 = -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    var v1 bool = true
    var v2 string = ""
    var v3 uint8 = 143
    var v4 int64 = 42
    var v5 int = 9223372036854775807
    v6 := []int{}
    v9 := GenS1{F0: v2}
    v5 = v5
    v10 := GenS1{F0: (v2 + (v2 + "foo bar"))}
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 string = "foo bar"
            fmt.Println(v3, fi11, fi12)
            v5 = (fi12 * (v5 * 959))
            v14 := GenS1{F0: (((((v13 + v2) + v2) + (v13 + ("a" + v13))) + v13) + (("hello" + ((v2 + "") + "a")) + (((v13 + v13) + "\n\t") + ("hello" + "a"))))}
            v15 := GenS1{F0: v13}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    switch v5 {
    case 0:
        if (v4 == (^(v4))) {
            _ = len(v6)
        }
    case 2147483647:
        v16 := GenS1{F0: (v2 + v2)}
        _ = v16
    case 32767:
        v3 = 3
    }
    v5 = (v5 & -2147483648)
    v5 = (v5 * (v5 - (v5 ^ (v5 - -833))))
    fmt.Println(v1)
    if ((v1 || (v1 || ((v3 & v3) >= 127))) && (!(v5 < v5))) {
        for fi17 := 0; fi17 < 10; fi17++ {
            if (v3 == v3) {
                v18 := GenS2{F0: 32767}
                _ = v18
            } else {
                var v19 uint64 = uint64(v3)
                _ = v19
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                if len(v6) > 0 {
                    v6[0] = (fi17 / ((-2147483648) | 1))
                }
                var v21 string = v2
                v3 -= v3
                _ = fi20
                _ = v21
            }
            if (!(!v1)) {
                switch (v5 & 32768) {
                case 0:
                    if (v1 || ((v1 || false) || false)) {
                        switch v2 {
                        case "\n\t":
                            switch 834 {
                            case 854:
                                var v22 uint64 = 1
                                _ = v22
                            case 42:
                                for fi23 := 0; fi23 < 17; fi23++ {
                                    v5 -= -420
                                    var v24 uint64 = 116
                                    v24 = v24
                                    _ = fi23
                                    _ = v24
                                }
                            case -9223372036854775808:
                                var v25 string = "hello"
                                _ = v25
                            }
                        }
                        v26 := GenS2{F0: 32767}
                        _ = v26
                    }
                case 707:
                    if (645 == 430) {
                        var v27 uint64 = 1
                        var v28 float32 = 0.1
                        _ = v27
                        _ = v28
                    }
                case 4294967295:
                    v3 += (^(v3))
                default:
                    v5 -= v5
                }
                v4 += (v4 % ((42) | 1))
                fmt.Println(v2, v5)
            } else {
                v29 := GenS1{F0: ((v2 + (v2 + v2)) + v2)}
                _ = v29
            }
            _ = fi17
        }
        var v30 int8 = 42
        if (!(v1 && false)) {
            v2 = ("" + (v2 + "foo bar"))
        } else {
            var v31 int64 = (v4 % (((v4 / ((v4) | 1))) | 1))
            if false {
                if len(v6) > 0 {
                    v6[0] = v5
                }
                var v32 uint32 = 2
                var v33 uint8 = v3
                _ = v32
                _ = v33
            } else {
                var v34 bool = (((-2.25 / 1e10) - 0.0) <= 1.5)
                v35 := GenS3{F0: 742, F1: v2}
                v31 += (v4 - v31)
                _ = v34
                _ = v35
            }
            _ = v31
        }
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
