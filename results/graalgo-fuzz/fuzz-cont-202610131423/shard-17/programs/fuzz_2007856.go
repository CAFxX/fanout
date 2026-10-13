package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint, p1 int64) uint8 {
    var v7 uint16 = 0
    var v8 uint = (p0 & (p0 >> ((p0) & 63)))
    switch 128 {
    case 1:
        fmt.Println(v7)
    case 659:
        if (true || (!(p0 <= (p0 + (p0 % ((v8) | 1)))))) {
            var v9 bool = false
            switch "hello" {
            case "hello":
                if v9 {
                    var v10 float32 = float32(0.0)
                    fmt.Println(v8, v9)
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 19; fi11++ {
                        v12 := GenS1{F0: v7, F1: v9}
                        fmt.Println(v7)
                        if (!(v8 > (v8 << ((v8) & 63)))) {
                            for fi13 := 0; fi13 < 1; fi13++ {
                                var v14 string = "hello"
                                _ = fi13
                                _ = v14
                            }
                            v15 := []int{fi11, fi11, 42}
                            if len(v15) > 0 {
                                v15[0] = (fi11 * 42)
                            }
                            _ = v15
                        }
                        _ = fi11
                        _ = v12
                    }
                    var v16 string = ""
                    v17 := []int{2147483647}
                    _ = v16
                    _ = v17
                }
            case "":
                v18 := GenS1{F0: (v7 + v7), F1: (!(v9 && (!(1000 > -32768))))}
                _ = v18
            case "a":
                var v19 uint = (p0 ^ (p0 + p0))
                _ = v19
            }
            _ = v9
        }
    }
    if false {
        for fi20 := 0; fi20 < 17; fi20++ {
            if ("" == "\n\t") {
                v21 := GenS1{F0: v7, F1: false}
                _ = v21
            } else {
                var v22 int8 = 0
                var v23 uint16 = (v7 | (v7 ^ (v7 * (v7 ^ 141))))
                if (!((^(v8)) == (p0 * 4294967295))) {
                    var v24 int = -1
                    v25 := GenS2{F0: -168}
                    _ = v24
                    _ = v25
                }
                _ = v22
                _ = v23
            }
            _ = fi20
        }
    } else {
        v26 := GenS2{F0: int64(v7)}
        switch 42 {
        case 2147483647:
            v27 := GenS1{F0: v7, F1: (!(p1 <= -1))}
            _ = v27
        case 4294967295:
            p1 += (p1 & (p1 | 0))
        }
        _ = v26
    }
    fmt.Println(v8, v8, v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint64, p1 bool) uint {
    fmt.Println(p1, p0)
    p1 = (4294967295 <= p0)
    p0 = 730
    p0 = (^(p0))
    p0 += (^(p0))
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "a"
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    v5 := []string{(v2 + "a"), v2}
    var v6 int8 = 42
    switch -9223372036854775808 {
    case -518:
        if (1 == -34) {
            v28 := GenS2{F0: -9223372036854775808}
            v2 = ""
            var v29 uint64 = 42
            _ = v28
            _ = v29
        } else {
            v30 := v1[(v2 + (v2 + "\n\t"))]
            for fi31 := 0; fi31 < 6; fi31++ {
                var v32 int16 = 32767
                var v33 uint16 = 584
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = v30
        }
    case 952:
        v6 += (v6 - v6)
    case 128:
        var v34 int32 = 418
        _ = v34
    default:
        var v35 int16 = -1
        _ = v35
    }
    var v36 string = "foo bar"
    v1[((v36 + v2) + v2)] = (101 < 32768)
    var v37 int8 = v6
    v36 = (v2 + "")
    switch 4294967295 {
    case 127:
        v5 = append(v5, v2)
    default:
        fmt.Println(v36, v6)
    }
    if true {
        v36 = v36
        v6 = -1
        v37 = (v37 >> ((genFn1(4294967295, (!true))) & 63))
    }
    v2 = ""
    var v38 int8 = 42
    v37 += (v37 & v6)
    if false {
        delete(v1, v2)
    } else {
        delete(v1, v36)
        var v39 uint8 = 1
        v37 -= v37
        _ = v39
    }
    var v40 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
