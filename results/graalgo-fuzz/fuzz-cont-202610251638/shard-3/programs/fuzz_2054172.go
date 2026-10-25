package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int16) uint {
    if true {
        v7 := []int{32768, 4294967295, -9223372036854775808}
        if true {
            for fi8 := 0; fi8 < 20; fi8++ {
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 int32 = 64
                    v10 = 42
                    _ = fi9
                    _ = v10
                }
                v11 := GenS1{F0: 2, F1: -79}
                v12 := GenS1{F0: 0, F1: 30}
                _ = fi8
                _ = v11
                _ = v12
            }
            var v13 int = 128
            _ = v13
        }
        fmt.Println(p0)
        _ = v7
    } else {
        p0 = (p0 - (p0 * (p0 & p0)))
        p0 = p0
    }
    if (2 > 128) {
        fmt.Println(p0)
        if true {
            var v14 uint32 = 42
            _ = v14
        } else {
            for fi15 := 0; fi15 < 9; fi15++ {
                if true {
                    var v16 uint8 = 0
                    v16 += v16
                    _ = v16
                }
                if (!false) {
                    p0 += (^(p0))
                } else {
                    v17 := GenS1{F0: 198, F1: 0}
                    if true {
                        v18 := []int{}
                        v19 := GenS1{F0: 2, F1: -128}
                        p0 = p0
                        _ = v18
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 15; fi20++ {
                            v21 := GenS1{F0: 190, F1: 1}
                            v22 := GenS2{F0: false}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        v23 := []int{}
                        var v24 uint8 = 128
                        _ = v23
                        _ = v24
                    }
                    _ = v17
                }
                _ = fi15
            }
            for fi25 := 0; fi25 < 6; fi25++ {
                p0 = (p0 - (p0 | (^(p0))))
                p0 -= p0
                _ = fi25
            }
            if false {
                var v26 int16 = 32767
                switch 128 {
                case 127:
                    v27 := GenS3{F0: 42}
                    _ = v27
                case 4294967295:
                    for fi28 := 0; fi28 < 5; fi28++ {
                        p0 += p0
                        _ = fi28
                    }
                }
                _ = v26
            }
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            switch (fi29 << ((402) & 63)) {
            case 4294967295:
                var v30 float64 = -2.25
                _ = v30
            case 0:
                fmt.Println(fi29, fi29)
            case 32767:
                v31 := []int{(fi29 - (fi29 ^ (fi29 + (fi29 % ((fi29) | 1))))), (fi29 ^ (fi29 + fi29)), fi29}
                _ = v31
            default:
                v32 := []int{42, (fi29 + 1), fi29}
                _ = v32
            }
            if true {
                switch "a" {
                case "foo bar":
                    var v33 uint64 = 4294967295
                    _ = v33
                case "a":
                    p0 += (p0 & p0)
                default:
                    p0 = p0
                }
                var v34 string = ("\n\t" + ("hello" + "foo bar"))
                _ = v34
            }
            p0 = p0
            _ = fi29
        }
    } else {
        var v35 int = 4294967295
        _ = v35
    }
    v36 := GenS4{F0: p0}
    _ = p0
    _ = v36
    return 42
}

func main() {
    v1 := []uint32{0}
    var v2 int8 = 1
    v3 := make(map[uint32]uint32)
    var v4 float32 = (-2.25 + float32(((-2.25 * (0.0 - 0.0)) * (((1e10 / (((3.14159 - float64(-2.25)) + (3.14159 * (3.14159 - 3.14159))) + 1.5)) * -2.25) / 3.14159))))
    var v5 uint = 65535
    var v6 int = 32767
    fmt.Println(v2, v2)
    switch (v6 + (v6 + v6)) {
    case 1:
        switch (^(v6)) {
        case -1:
            fmt.Println(v2, v5, v2)
        case 0:
            switch (((("a" + (("hello" + ("foo bar" + ("" + "hello"))) + "foo bar")) + "hello") + "foo bar") + "a") {
            case "a":
                var v37 int = v6
                _ = v37
            case "":
                for fi38 := 0; fi38 < 9; fi38++ {
                    fmt.Println(v6)
                    _ = fi38
                }
            case "foo bar":
                var v39 int8 = (v2 ^ v2)
                _ = v39
            default:
                v4 = float32(1e10)
            }
        case 9223372036854775807:
            for fi40 := 0; fi40 < 19; fi40++ {
                v41 := GenS2{F0: ((v4 + 1.5) <= v4)}
                _ = fi40
                _ = v41
            }
        }
    default:
        var v42 float64 = 3.14159
        _ = v42
    }
    if ((true || (!(v2 == v2))) && true) {
        fmt.Println(v4, v2)
    }
    if (true || (4294967295 < 42)) {
        _ = len(v1)
        delete(v3, uint32(v5))
    }
    v6 = v6
    v5 -= v5
    v43 := v3[2147483648]
    var v44 float32 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
