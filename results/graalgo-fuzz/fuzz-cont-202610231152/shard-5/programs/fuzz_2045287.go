package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint; F2 string }
func genFn0(p0 int64) uint32 {
    var v4 uint32 = 42
    v5 := GenS1{F0: -1}
    if (!true) {
        fmt.Println(v4, p0)
        if true {
            if true {
                if true {
                    p0 = (p0 | (p0 + (-(p0))))
                    v6 := []int{}
                    p0 = (p0 + p0)
                    _ = v6
                }
                switch ("" + ("a" + (("\n\t" + ("a" + ("a" + ("hello" + "foo bar")))) + "hello"))) {
                case "hello":
                    switch "\n\t" {
                    case "foo bar":
                        if true {
                            var v7 uint32 = (v4 % ((v4) | 1))
                            var v8 bool = false
                            var v9 uint8 = 42
                            _ = v7
                            _ = v8
                            _ = v9
                        }
                    case "a":
                        if (-1 > 127) {
                            for fi10 := 0; fi10 < 3; fi10++ {
                                var v11 int16 = 0
                                var v12 int = (fi10 << ((4294967295) & 63))
                                for fi13 := 0; fi13 < 4; fi13++ {
                                    if false {
                                        var v14 uint32 = 1
                                        v11 = v11
                                        _ = v14
                                    }
                                    _ = fi13
                                }
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                            v4 += v4
                            p0 = p0
                        }
                    case "hello":
                        p0 -= 42
                    default:
                        v4 += 1
                    }
                }
                for fi15 := 0; fi15 < 19; fi15++ {
                    p0 += -408
                    for fi16 := 0; fi16 < 19; fi16++ {
                        if (255 != 0) {
                            v17 := []int{-2147483648}
                            _ = v17
                        } else {
                            fmt.Println(fi15, fi16, fi15)
                        }
                        var v18 float64 = (3.14159 * (0.0 - 1e10))
                        fmt.Println(v4)
                        _ = fi16
                        _ = v18
                    }
                    v19 := []int{(fi15 & (-(fi15))), 127, 0}
                    _ = fi15
                    _ = v19
                }
            } else {
                fmt.Println(p0)
                v20 := GenS1{F0: 0}
                v21 := GenS2{F0: v4}
                _ = v20
                _ = v21
            }
            p0 -= p0
            for fi22 := 0; fi22 < 20; fi22++ {
                fmt.Println(v4, p0, fi22)
                _ = fi22
            }
        } else {
            v4 = v4
        }
    } else {
        p0 += (p0 | p0)
        v23 := []int{4294967295, 32767, int(v4)}
        v24 := GenS1{F0: 42}
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            p0 = (-(p0))
            _ = fi26
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            switch fi25 {
            case 127:
                if false {
                    fmt.Println(fi25, fi25, v4)
                    var v28 int8 = -128
                    fmt.Println(fi27, v28, v4)
                    _ = v28
                } else {
                    v29 := []int{fi25, fi27, fi27}
                    fmt.Println(fi27)
                    var v30 int8 = 49
                    v31 := GenS3{F0: 879, F1: 42, F2: "foo bar"}
                    _ = v29
                    _ = v30
                    _ = v31
                }
            }
            fmt.Println(fi25, p0)
            _ = fi27
        }
        _ = fi25
    }
    _ = p0
    _ = v4
    _ = v5
    return v4
}

func main() {
    v1 := []uint32{4294967295, 2, 42}
    var v2 int16 = 32767
    var v3 int = -9223372036854775808
    v2 += 320
    v32 := GenS1{F0: 32767}
    v3 = (v3 | 9223372036854775807)
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v33 int32 = 1
    for fi34 := 0; fi34 < 20; fi34++ {
        var v35 int8 = 42
        v36 := GenS2{F0: 2147483647}
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 float32 = 0.0
    var v38 uint = 4294967295
    switch v3 {
    case -353:
        v37 = v37
    }
    v1 = append(v1, 65535)
    for fi39 := 0; fi39 < 2; fi39++ {
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = fi39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
}
