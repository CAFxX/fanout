package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int, p1 int64) uint8 {
    var v6 bool = (!false)
    v6 = (v6 || false)
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 string = ("\n\t" + "\n\t")
        p0 += (p0 + -146)
        switch p0 {
        case 42:
            if (31 >= 100) {
                if (2 != 2) {
                    fmt.Println(fi7, v6)
                    if (p0 == p0) {
                        if (!v6) {
                            if v6 {
                                p1 += p1
                            }
                        } else {
                            v8 = v8
                        }
                        fmt.Println(fi7, v6)
                    } else {
                        var v9 bool = v6
                        _ = v9
                    }
                }
                v8 = "\n\t"
                switch "hello" {
                case "foo bar":
                    if v6 {
                        v10 := GenS1{F0: -1}
                        _ = v10
                    } else {
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v12 := GenS2{F0: 1}
                            _ = fi11
                            _ = v12
                        }
                        var v13 uint = 0
                        _ = v13
                    }
                case "hello":
                    if (int64(p0) < p1) {
                        p1 = 1
                        for fi14 := 0; fi14 < 4; fi14++ {
                            var v15 int64 = p1
                            v16 := []int{fi14}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 string = v8
                        _ = v17
                    }
                }
            }
        case 32767:
            var v18 uint64 = 18446744073709551615
            _ = v18
        }
        _ = fi7
        _ = v8
    }
    v19 := []int{(-(p0)), p0, p0, -2147483648}
    _ = p0
    _ = p1
    _ = v19
    _ = v6
    return 174
}

func main() {
    var v1 int16 = 1
    v2 := []int{2147483647}
    var v3 float64 = 0.1
    v4 := make(map[int]uint32)
    var v5 int16 = 42
    for fi20 := 0; fi20 < 19; fi20++ {
        if false {
            v5 = (v5 >> ((857) & 63))
            var v21 bool = false
            if v21 {
                fmt.Println(v21)
            }
            _ = v21
        }
        fmt.Println(v5, v1, v3)
        _ = fi20
    }
    if true {
        v22 := GenS1{F0: 100}
        var v23 int64 = -9223372036854775808
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: int8(v5)}
    v3 -= (float64(float32(v3)) + float64(0.1))
    if ((18446744073709551615 >= 42) && false) {
        v5 = (v5 << ((2) & 63))
        var v25 int64 = -9223372036854775808
        _ = v25
    }
    fmt.Println(v1, v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v26 int16 = (v5 % (((v1 % (((^(v5))) | 1))) | 1))
    for fi27 := 0; fi27 < 7; fi27++ {
        delete(v4, fi27)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v28 string = "foo bar"
        _ = fi27
        _ = v28
    }
    if false {
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 uint16 = 49
            delete(v4, fi29)
            v31 := GenS1{F0: -89}
            _ = fi29
            _ = v30
            _ = v31
        }
        v32 := GenS3{F0: 255, F1: 42}
        var v33 uint32 = 65535
        _ = v32
        _ = v33
    } else {
        var v34 float32 = -2.25
        v35 := v4[127]
        for fi36 := 0; fi36 < 10; fi36++ {
            v37 := GenS4{F0: 2, F1: 1}
            v4[4294967295] = v35
            _ = fi36
            _ = v37
        }
        _ = v34
        _ = v35
    }
    v26 = v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
