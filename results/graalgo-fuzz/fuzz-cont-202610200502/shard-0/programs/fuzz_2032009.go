package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]string)
    var v3 uint16 = 430
    var v4 uint64 = 42
    var v5 uint8 = 38
    var v6 int = 32768
    delete(v2, 122)
    delete(v2, 4294967295)
    var v7 uint64 = 661
    v6 -= (v6 / ((-2147483648) | 1))
    if true {
        var v8 int = (v6 % (((-(v6))) | 1))
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := []int{v6, v6, (-(v8)), -9223372036854775808}
            if len(v10) > 0 {
                v10[0] = (v6 % (((v8 - (fi9 << ((18446744073709551615) & 63)))) | 1))
            }
            _ = fi9
            _ = v10
        }
        if false {
            if true {
                fmt.Println(v5, v5)
                v3 -= v3
            }
            var v11 bool = true
            _ = v11
        } else {
            v1 = 18446744073709551615
            v12 := GenS1{F0: -46}
            var v13 int = v6
            _ = v12
            _ = v13
        }
        _ = v8
    } else {
        v2[uint32(v3)] = "\n\t"
        v1 -= (v1 * uint(v5))
        v14 := []int{int(v4), (v6 >> ((4294967295) & 63)), v6}
        _ = v14
    }
    var v15 uint = v1
    v1 += uint(v3)
    var v16 bool = true
    v7 = (v4 >> ((v15) & 63))
    v17 := GenS1{F0: -9223372036854775808}
    v3 = v3
    if (v6 == 128) {
        for fi18 := 0; fi18 < 10; fi18++ {
            switch "foo bar" {
            case "a":
                if (1e10 <= 0.0) {
                    var v19 float32 = ((1e10 + (float32(0.1) - 0.1)) * (1e10 / 0.1))
                    v4 = v4
                    _ = v19
                } else {
                    if v16 {
                        v15 = (^(v1))
                        v16 = false
                    } else {
                        v15 += v15
                        if v16 {
                            v20 := GenS1{F0: 754}
                            if true {
                                if true {
                                    v5 = v5
                                    fmt.Println(v7)
                                }
                                v4 -= v7
                            } else {
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    v22 := GenS1{F0: 748}
                                    var v23 string = "hello"
                                    var v24 float32 = 1.5
                                    _ = len(v2)
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                if false {
                                    v25 := GenS1{F0: 1}
                                    _ = v25
                                } else {
                                    v26 := []int{(v6 | (v6 / ((927) | 1))), -867, v6, -45}
                                    for fi27 := 0; fi27 < 17; fi27++ {
                                        if false {
                                            fmt.Println(fi27)
                                            fmt.Println(v4, fi27)
                                            v4 -= 18446744073709551615
                                        } else {
                                            v5 += 255
                                            _ = len(v26)
                                            var v28 int64 = -661
                                            _ = v28
                                        }
                                        _ = fi27
                                    }
                                    if false {
                                        var v29 int8 = 0
                                        _ = v29
                                    }
                                    _ = v26
                                }
                            }
                            var v30 int = 4294967295
                            _ = v20
                            _ = v30
                        } else {
                            v1 -= (v1 ^ v15)
                            fmt.Println(v16)
                        }
                        if (true && v16) {
                            v31 := GenS2{F0: v1, F1: -138}
                            _ = v31
                        } else {
                            for fi32 := 0; fi32 < 11; fi32++ {
                                fmt.Println(v16)
                                v7 += (v7 & (v4 / (((^(v7))) | 1)))
                                v33 := []int{(v6 << (((v15 - (v1 | (v15 + v15)))) & 63)), fi32, (v6 & (fi18 >> (((v1 ^ uint(v7))) & 63)))}
                                _ = fi32
                                _ = v33
                            }
                        }
                    }
                    v34 := GenS3{F0: -119}
                    var v35 string = ""
                    _ = v34
                    _ = v35
                }
            case "":
                var v36 int16 = -989
                _ = v36
            }
            var v37 uint32 = 2
            var v38 bool = (42 > v6)
            _ = fi18
            _ = v37
            _ = v38
        }
    }
    v15 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
