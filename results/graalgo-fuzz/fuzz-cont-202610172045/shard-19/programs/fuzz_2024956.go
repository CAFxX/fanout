package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int; F2 uint }
type GenS3 struct { F0 int64; F1 uint; F2 uint }
type GenS4 struct { F0 int; F1 int8 }
func genFn0(p0 int, p1 int64) bool {
    if true {
        p1 += p1
        fmt.Println(p0, p0)
        v5 := GenS1{F0: p1}
        _ = v5
    }
    fmt.Println(p0, p1)
    v6 := []int{}
    for fi7 := 0; fi7 < 17; fi7++ {
        _ = len(v6)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return false
}

func main() {
    var v1 int16 = 32767
    var v2 string = (("" + "foo bar") + "\n\t")
    var v3 uint64 = 996
    v4 := []string{}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint8 = 42
        _ = fi8
        _ = v9
    }
    v2 = "\n\t"
    v1 = (v1 % (((v1 << ((4294967295) & 63))) | 1))
    if false {
        v2 = (v2 + "hello")
    }
    var v10 uint32 = 1
    var v11 int32 = 42
    switch v2 {
    case "foo bar":
        v3 = 375
    case "a":
        var v12 uint32 = 42
        _ = v12
    case "\n\t":
        var v13 int16 = (v1 % (((v1 / ((v1) | 1))) | 1))
        _ = v13
    }
    var v14 bool = false
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            fmt.Println(fi15, v14)
            v1 = int16(fi15)
            v17 := GenS2{F0: (v2 + (v2 + v2)), F1: (fi16 + fi15), F2: 742}
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    if false {
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 float64 = 1e10
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 uint16 = 42
                v22 := GenS3{F0: -78, F1: 0, F2: 468}
                fmt.Println(fi18, v21, v3)
                _ = fi20
                _ = v21
                _ = v22
            }
            switch fi18 {
            case 32768:
                fmt.Println(v3)
            case 9223372036854775807:
                v23 := GenS2{F0: "hello", F1: int(v3), F2: 65535}
                _ = v23
            }
            _ = fi18
            _ = v19
        }
        if v14 {
            var v24 int16 = v1
            fmt.Println(v3, v10, v1)
            _ = v24
        } else {
            var v25 uint16 = 2
            var v26 uint32 = (v10 | 0)
            _ = v25
            _ = v26
        }
        switch 417 {
        case -1:
            for fi27 := 0; fi27 < 15; fi27++ {
                v3 -= 2
                if true {
                    if v14 {
                        v28 := GenS4{F0: fi27, F1: 42}
                        _ = v28
                    } else {
                        if true {
                            fmt.Println(fi27)
                            if len(v4) > 0 {
                                v4[0] = (v2 + (("a" + v2) + (v2 + "a")))
                            }
                        } else {
                            v14 = (!((v10 >> ((0) & 63)) == v10))
                        }
                        if v14 {
                            var v29 int8 = -128
                            v30 := GenS4{F0: (fi27 ^ (fi27 % ((fi27) | 1))), F1: -100}
                            v11 -= 42
                            _ = v29
                            _ = v30
                        } else {
                            var v31 int8 = -10
                            v14 = (!(!(!(-443 >= -9223372036854775808))))
                            fmt.Println(v1, fi27, v14)
                            _ = v31
                        }
                    }
                    var v32 float32 = (1e10 - -2.25)
                    _ = v32
                } else {
                    fmt.Println(v1)
                    switch fi27 {
                    case 32767:
                        var v33 string = v2
                        _ = v33
                    }
                }
                v14 = v14
                _ = fi27
            }
        case 9223372036854775807:
            fmt.Println(v10, v14)
        }
    } else {
        v3 += (v3 + v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
