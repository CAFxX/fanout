package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int8; F1 int8; F2 int64 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int8, p1 uint) uint {
    v4 := []int{-2147483648, 32768, 127, 1}
    p0 -= p0
    for fi5 := 0; fi5 < 10; fi5++ {
        p1 -= (^(p1))
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 16
}

func main() {
    var v1 int16 = 1
    var v2 uint = 842
    v3 := make(map[string]uint32)
    var v6 uint8 = 0
    v7 := GenS1{F0: ((0.1 / (0.1 + 1e10)) < 3.14159), F1: genFn0(0, v2)}
    var v8 int8 = -22
    var v9 int = 42
    fmt.Println(v8, v9)
    fmt.Println(v1, v1, v8)
    v1 += (v1 - v1)
    v6 += v6
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint16 = 1
        v6 = 0
        v6 += 0
        _ = fi10
        _ = v11
    }
    var v12 int32 = -706
    v13 := GenS2{F0: 1, F1: 255}
    var v14 int8 = v8
    v1 = (v1 / ((-32768) | 1))
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 int16 = (v1 - v1)
            var v18 uint8 = (v6 % (((v6 | 0)) | 1))
            if true {
                v19 := GenS3{F0: 690, F1: "\n\t"}
                _ = v19
            } else {
                var v20 int64 = 0
                var v21 int16 = (v17 ^ v1)
                _ = v20
                _ = v21
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        if true {
            v22 := GenS4{F0: int8(v6), F1: (v8 % ((42) | 1)), F2: int64(v12)}
            if false {
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 uint16 = 947
                    v25 := GenS4{F0: (-(v8)), F1: int8(v12), F2: 1}
                    delete(v3, "a")
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                for fi26 := 0; fi26 < 16; fi26++ {
                    if ((32767 != 783) && ((v2 / ((v2) | 1)) < (v2 % (((v2 + genFn0((v8 << ((genFn0(11, (v2 - 623))) & 63)), 65535))) | 1)))) {
                        fmt.Println(v2, fi15, v6)
                        v27 := v3[("a" + "hello")]
                        v28 := GenS3{F0: 780, F1: "\n\t"}
                        delete(v3, ("foo bar" + ("hello" + ("hello" + ((("hello" + ("foo bar" + "a")) + "foo bar") + ("foo bar" + "\n\t"))))))
                        v29 := GenS4{F0: v8, F1: int8(v27), F2: 42}
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    var v30 uint = (v2 << ((982) & 63))
                    var v31 int8 = v14
                    _ = fi26
                    _ = v30
                    _ = v31
                }
                var v32 int = (v9 | v9)
                _ = v32
            } else {
                v8 = -1
                fmt.Println(v2, v8, v8)
                v2 = v2
            }
            var v33 uint64 = 2
            _ = v22
            _ = v33
        }
        _ = fi15
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        switch "\n\t" {
        case "foo bar":
            for fi35 := 0; fi35 < 4; fi35++ {
                v12 -= v12
                v6 -= v6
                _ = fi35
            }
        }
        v6 -= (v6 ^ (v6 ^ (^(v6))))
        v12 = v12
        _ = fi34
    }
    switch v9 {
    case -62:
        var v36 uint = v2
        _ = v36
    case -539:
        if false {
            if (1 > 2147483648) {
                v37 := GenS4{F0: (v14 % ((v8) | 1)), F1: 74, F2: -82}
                switch (-(v9)) {
                case 30:
                    var v38 int8 = (v8 & v8)
                    _ = v38
                case 32767:
                    v39 := GenS5{F0: 42, F1: 2}
                    _ = v39
                case 631:
                    var v40 int16 = (v1 % (((v1 % (((v1 / ((900) | 1))) | 1))) | 1))
                    _ = v40
                }
                if ("a" < (("foo bar" + ("a" + ("hello" + "\n\t"))) + "hello")) {
                    v8 -= (^(v8))
                } else {
                    fmt.Println(v9, v8, v9)
                    fmt.Println(v6)
                    fmt.Println(v9, v8)
                }
                _ = v37
            } else {
                if true {
                    v1 = v1
                } else {
                    v41 := GenS6{F0: v8}
                    v42 := []int{v9, v9, v9, (^(v9))}
                    v2 = 255
                    _ = v41
                    _ = v42
                }
                v43 := []int{}
                _ = v43
            }
        } else {
            var v44 uint8 = v6
            _ = v44
        }
    default:
        for fi45 := 0; fi45 < 3; fi45++ {
            v46 := []int{}
            _ = fi45
            _ = v46
        }
    }
    if ((3.14159 <= 1e10) && true) {
        var v47 int32 = (v12 * (v12 ^ v12))
        var v48 int32 = (v12 << (((v2 % ((4294967295) | 1))) & 63))
        _ = v47
        _ = v48
    }
    for fi49 := 0; fi49 < 18; fi49++ {
        var v50 int16 = -32768
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
