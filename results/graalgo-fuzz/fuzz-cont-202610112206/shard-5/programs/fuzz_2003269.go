package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int8; F1 bool; F2 bool }
func genFn0(p0 int8) uint16 {
    v6 := []int{102}
    if (false || (65535 >= 1)) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        p0 = -1
        var v7 uint = 130
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 int64 = 1
    var v2 uint = 177
    var v3 int = 32767
    var v4 string = "a"
    var v5 uint8 = 42
    v8 := []int{v3, v3}
    if false {
        v1 = (-(v1))
    } else {
        if true {
            v9 := GenS1{F0: 2147483647, F1: 1, F2: 4294967295}
            fmt.Println(v4)
            _ = v9
        } else {
            v1 += v1
            for fi10 := 0; fi10 < 2; fi10++ {
                if len(v8) > 0 {
                    v8[0] = v3
                }
                _ = fi10
            }
            _ = len(v8)
        }
        switch (v3 ^ (v3 % (((v3 / ((v3) | 1))) | 1))) {
        case 149:
            if (false && ((202 == (v2 * v2)) && (!true))) {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 float32 = 3.14159
                    _ = fi11
                    _ = v12
                }
                v8 = append(v8, v3)
            }
        case 128:
            var v13 uint = v2
            _ = v13
        case -9223372036854775808:
            var v14 string = v4
            _ = v14
        default:
            switch "foo bar" {
            case "\n\t":
                switch v4 {
                case "a":
                    v15 := GenS2{F0: v2, F1: (true || (1 > 23))}
                    _ = v15
                default:
                    v5 = v5
                }
            case "foo bar":
                for fi16 := 0; fi16 < 6; fi16++ {
                    if false {
                        v2 = 65535
                    }
                    _ = fi16
                }
            case "a":
                fmt.Println(v2, v1, v3)
            }
        }
        if (false || (!false)) {
            v2 -= (v2 + 1)
        }
    }
    v3 = (-(v3))
    v17 := GenS2{F0: (v2 % (((^(v2))) | 1)), F1: false}
    for fi18 := 0; fi18 < 11; fi18++ {
        if true {
            _ = len(v8)
        }
        _ = fi18
    }
    var v19 int8 = 1
    v4 = "a"
    if (false || (!false)) {
        var v20 float64 = float64(0.0)
        if len(v8) > 0 {
            v8[0] = (^(v3))
        }
        switch v4 {
        case "foo bar":
            v2 = 712
        case "":
            var v21 uint16 = genFn0((v19 ^ -100))
            _ = v21
        }
        _ = v20
    } else {
        if true {
            v1 = -218
            v22 := GenS3{F0: (-(v19)), F1: false, F2: true}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 uint64 = 143
                _ = fi23
                _ = v24
            }
            v8 = append(v8, (v3 & v3))
        }
    }
    v19 += (^(v19))
    var v25 int32 = -2147483648
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        switch v3 {
        case -2147483648:
            v27 := GenS1{F0: 42, F1: (v25 | 2147483647), F2: 1}
            _ = v27
        default:
            var v28 int32 = (v25 + 1)
            _ = v28
        }
        for fi29 := 0; fi29 < 13; fi29++ {
            if (!true) {
                if ((0 > 1) || false) {
                    var v30 uint32 = 0
                    switch v4 {
                    case "hello":
                        v8 = append(v8, fi26)
                    }
                    var v31 uint64 = 20
                    _ = v30
                    _ = v31
                } else {
                    fmt.Println(v2, v3)
                }
                fmt.Println(fi26)
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            } else {
                var v32 uint64 = 665
                var v33 uint32 = 42
                _ = v32
                _ = v33
            }
            _ = fi29
        }
        _ = fi26
    }
    var v34 int64 = v1
    for fi35 := 0; fi35 < 6; fi35++ {
        if false {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v36 uint64 = 65
            v2 += v2
            _ = v36
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v8
}
