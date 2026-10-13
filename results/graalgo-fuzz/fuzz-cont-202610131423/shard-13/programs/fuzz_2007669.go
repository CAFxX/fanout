package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int64; F1 int16 }
func genFn0(p0 int32, p1 bool) int64 {
    p0 += (^(p0))
    var v7 uint16 = 42
    v8 := GenS1{F0: 2, F1: 128}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint16) uint64 {
    var v9 int8 = 100
    for fi10 := 0; fi10 < 2; fi10++ {
        switch fi10 {
        case 66:
            var v11 uint32 = 0
            _ = v11
        }
        v12 := []int{(fi10 & 469)}
        _ = fi10
        _ = v12
    }
    if (!false) {
        v13 := GenS1{F0: 1, F1: 158}
        fmt.Println(v9)
        _ = v13
    }
    var v14 int8 = v9
    v15 := []int{1, int(p0), 860}
    _ = p0
    _ = v14
    _ = v15
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 38
    v3 := make(map[string]uint32)
    var v4 uint8 = 42
    var v5 uint16 = 65535
    v6 := make(map[uint32]string)
    v4 = 2
    var v16 int32 = 1
    var v17 int64 = -9223372036854775808
    if true {
        var v18 float32 = (float32(3.14159) / 1.0)
        _ = v18
    }
    var v19 int = 1
    for fi20 := 0; fi20 < 20; fi20++ {
        if (((-2147483648 < v19) && false) || ((v2 | v2) >= (^(v2)))) {
            if false {
                v4 += v4
            } else {
                if (genFn1((^(v5))) <= 18446744073709551615) {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        fmt.Println(v5)
                        v2 = v2
                        v22 := GenS2{F0: -32768, F1: false}
                        _ = fi21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if false {
                            fmt.Println(fi20, v5)
                            delete(v1, "a")
                            if true {
                                v17 -= genFn0(v16, ((^(v4)) != (v4 - uint8(v5))))
                                v24 := []int{fi23, (-(fi23))}
                                _ = v24
                            }
                        } else {
                            v25 := []int{fi20}
                            _ = v25
                        }
                        v2 = (v2 + v2)
                        var v26 string = ("hello" + "\n\t")
                        _ = fi23
                        _ = v26
                    }
                    fmt.Println(v4, v5)
                } else {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v28 := GenS3{F0: 401}
                        if false {
                            fmt.Println(fi27)
                            _ = len(v1)
                            v3["a"] = 2147483647
                        }
                        _ = fi27
                        _ = v28
                    }
                    if false {
                        switch "a" {
                        case "\n\t":
                            if true {
                                var v29 int = (fi20 % (((fi20 << (((^(v2))) & 63))) | 1))
                                var v30 uint = (v2 ^ v2)
                                fmt.Println(fi20, v4)
                                _ = v29
                                _ = v30
                            }
                        case "foo bar":
                            for fi31 := 0; fi31 < 2; fi31++ {
                                fmt.Println(fi20)
                                _ = fi31
                            }
                        case "a":
                            var v32 uint8 = v4
                            _ = v32
                        }
                    }
                    if (1.5 < 3.14159) {
                        var v33 float32 = (0.0 - 0.1)
                        v16 -= 64
                        var v34 uint64 = genFn1(v5)
                        _ = v33
                        _ = v34
                    } else {
                        v35 := []int{(fi20 << ((4294967295) & 63)), (fi20 + (fi20 | v19)), (v19 >> ((0) & 63)), (-(v19))}
                        _ = len(v35)
                        v17 = v17
                        _ = v35
                    }
                }
                var v36 uint8 = v4
                var v37 uint32 = 65535
                _ = v36
                _ = v37
            }
        } else {
            var v38 int64 = v17
            v17 = int64(v5)
            _ = v38
        }
        for fi39 := 0; fi39 < 13; fi39++ {
            v40 := GenS4{F0: (v4 & 128), F1: 4294967295}
            v41 := GenS3{F0: int16(fi39)}
            var v42 uint64 = 76
            _ = fi39
            _ = v40
            _ = v41
            _ = v42
        }
        _ = fi20
    }
    fmt.Println(v19)
    fmt.Println(v2, v4, v2)
    var v43 uint16 = v5
    var v44 int16 = -446
    var v45 bool = false
    v46 := v1[""]
    for fi47 := 0; fi47 < 4; fi47++ {
        v48 := GenS5{F0: -1, F1: (v44 - (^(v44)))}
        _ = fi47
        _ = v48
    }
    delete(v6, 443)
    var v49 uint = (v2 << ((v2) & 63))
    v19 = (v19 ^ (v19 * v19))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v49
    _ = v5
    _ = v6
}
