package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 uint8 }
func genFn0(p0 int16, p1 int32, p2 string) uint32 {
    fmt.Println(p0)
    p0 -= (p0 ^ -787)
    var v4 int32 = 1
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int32 = v4
        if (((^(fi7)) < 1) && ((1 == 1) || (!false))) {
            p2 = p2
            if true {
                p2 = p2
                switch fi7 {
                case 4294967295:
                    p1 += v8
                }
                if (42 < 18446744073709551615) {
                    v9 := GenS1{F0: 835}
                    p0 = (^(p0))
                    if (42 == 19) {
                        var v10 int64 = -9223372036854775808
                        v11 := GenS2{F0: p2, F1: -1}
                        _ = v10
                        _ = v11
                    }
                    _ = v9
                }
            } else {
                p2 = p2
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                if false {
                    var v13 uint = 4294967295
                    var v14 uint = (v13 | 255)
                    _ = v13
                    _ = v14
                }
                _ = fi12
            }
        } else {
            v8 = (^(p1))
            for fi15 := 0; fi15 < 20; fi15++ {
                p2 = "a"
                if true {
                    v16 := GenS2{F0: ((p2 + (p2 + p2)) + ""), F1: (v4 - (v4 ^ 42))}
                    _ = v16
                }
                v4 = 42
                _ = fi15
            }
        }
        v17 := []int{-1}
        _ = fi7
        _ = v17
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 uint32 = 0
    switch -83 {
    case 353:
        switch -1 {
        case -2147483648:
            switch -372 {
            case -9223372036854775808:
                var v18 int8 = -128
                _ = v18
            case 4294967295:
                v19 := GenS3{F0: -2147483648, F1: uint64(v3)}
                _ = v19
            default:
                var v20 bool = true
                _ = v20
            }
        default:
            if true {
                v3 = (^(v3))
                var v21 uint8 = 193
                var v22 float64 = 3.14159
                _ = v21
                _ = v22
            }
        }
    case 696:
        if (-1 < -128) {
            fmt.Println(v3)
        }
    case 2147483647:
        v23 := GenS2{F0: ("\n\t" + ""), F1: 42}
        _ = v23
    }
    fmt.Println(v3)
    var v24 uint64 = 0
    for fi25 := 0; fi25 < 1; fi25++ {
        if (388 >= 42) {
            v3 = v3
            switch (fi25 * fi25) {
            case -2147483648:
                var v26 uint16 = 32767
                _ = v26
            case -1:
                v27 := []int{(fi25 & fi25), (fi25 / ((0) | 1)), (fi25 % (((^(fi25))) | 1))}
                _ = v27
            }
            if (((42 != -97) || (!((v3 ^ genFn0(32767, 42, (("\n\t" + "foo bar") + ""))) < genFn0(487, -772, "\n\t")))) && (127 > 127)) {
                var v28 uint8 = 2
                for fi29 := 0; fi29 < 6; fi29++ {
                    _ = len(v2)
                    _ = fi29
                }
                _ = v28
            } else {
                if (!true) {
                    v30 := []int{0}
                    v2[1] = (v3 % (((v3 - (^(v3)))) | 1))
                    _ = v30
                } else {
                    var v31 int8 = 127
                    _ = v31
                }
            }
        } else {
            switch 0 {
            case -9223372036854775808:
                fmt.Println(fi25, v24)
            }
            fmt.Println(v24, fi25, fi25)
            if true {
                for fi32 := 0; fi32 < 11; fi32++ {
                    v33 := GenS4{F0: 128}
                    v3 = v3
                    var v34 uint8 = 90
                    if (!((!true) && true)) {
                        var v35 int = (fi25 * -278)
                        _ = v35
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                if (false || (3.14159 > float32(-2.25))) {
                    v1[genFn0(-526, 1, ("foo bar" + "hello"))] = fi25
                    var v36 float32 = 3.14159
                    for fi37 := 0; fi37 < 4; fi37++ {
                        var v38 string = "hello"
                        v39 := []int{(fi37 >> ((1) & 63)), -2147483648}
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    _ = v36
                } else {
                    v24 += 873
                    v3 = (v3 ^ (v3 - genFn0(0, 2147483647, "hello")))
                }
            }
        }
        _ = fi25
    }
    for fi40 := 0; fi40 < 2; fi40++ {
        v41 := GenS1{F0: 502}
        _ = fi40
        _ = v41
    }
    v24 = 358
    delete(v2, 695)
    switch ("\n\t" + "foo bar") {
    case "a":
        v42 := GenS5{F0: -76, F1: 42}
        _ = v42
    case "":
        v43 := []int{}
        _ = v43
    default:
        v3 += v3
    }
    delete(v2, 128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
