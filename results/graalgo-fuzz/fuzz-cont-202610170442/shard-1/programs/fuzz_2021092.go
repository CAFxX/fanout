package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 int32 }
func genFn0(p0 int16, p1 int16) int8 {
    v4 := []int{-1, -9223372036854775808}
    _ = len(v4)
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 float32 = (1.5 * 0.0)
        if true {
            var v7 string = ""
            v8 := GenS1{F0: v7}
            v7 = (v7 + "hello")
            _ = v7
            _ = v8
        } else {
            if (((fi5 & (^(fi5))) > (fi5 >> ((52) & 63))) && false) {
                p0 -= p1
            }
            v4 = append(v4, -9223372036854775808)
            if (!(false && (!((1e10 >= (1.5 - float64(v6))) || (255 >= 1))))) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi9 := 0; fi9 < 10; fi9++ {
                    v10 := GenS2{F0: 1, F1: -1, F2: (^(p1))}
                    _ = fi9
                    _ = v10
                }
                v11 := GenS3{F0: 2}
                _ = v11
            } else {
                v12 := GenS1{F0: ""}
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 float32 = (v6 / v6)
                    switch ("" + "") {
                    case "hello":
                        v4 = append(v4, (^(fi5)))
                    default:
                        switch fi5 {
                        case 2147483647:
                            v15 := GenS3{F0: 255}
                            _ = v15
                        case -571:
                            p1 -= p1
                        case 298:
                            p0 = p1
                        }
                    }
                    _ = fi13
                    _ = v14
                }
                v16 := GenS1{F0: ("" + "a")}
                _ = v12
                _ = v16
            }
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return -58
}

func main() {
    var v1 int16 = 1
    v2 := []int{-789, 32767}
    v3 := make(map[uint32]int)
    v2 = append(v2, 9223372036854775807)
    v17 := GenS4{F0: true}
    if true {
        fmt.Println(v1)
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := GenS5{F0: 4294967295, F1: -1}
            var v21 uint = 255
            var v22 uint32 = 2
            var v23 uint8 = 232
            v1 += (v1 * (v1 / (((^(v1))) | 1)))
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi18
    }
    var v24 float32 = (((0.1 / -2.25) - 1e10) - 3.14159)
    switch 128 {
    case -1:
        var v25 uint64 = 4294967295
        _ = v25
    case 2147483647:
        for fi26 := 0; fi26 < 5; fi26++ {
            v1 += v1
            fmt.Println(v1, v1, v1)
            _ = fi26
        }
    case 42:
        v27 := GenS2{F0: 2, F1: 356, F2: 0}
        _ = v27
    }
    v1 = v1
    v1 += v1
    var v28 string = "a"
    switch 32768 {
    case 2147483647:
        delete(v3, 695)
    case 0:
        for fi29 := 0; fi29 < 8; fi29++ {
            for fi30 := 0; fi30 < 1; fi30++ {
                switch fi30 {
                case 128:
                    if false {
                        if (!true) {
                            var v31 int8 = genFn0(32767, v1)
                            var v32 float32 = v24
                            switch ((v28 + ((v28 + (v28 + v28)) + (v28 + v28))) + "a") {
                            case "a":
                                var v33 int32 = -2147483648
                                _ = v33
                            }
                            if (!true) {
                                delete(v3, 2)
                            }
                            v31 -= v31
                            _ = v31
                            _ = v32
                        }
                        v34 := GenS1{F0: v28}
                        _ = v34
                    }
                }
                switch v28 {
                case "":
                    var v35 string = v28
                    _ = v35
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = (fi29 ^ fi30)
                    }
                case "hello":
                    v28 = ((v28 + v28) + v28)
                default:
                    v36 := GenS3{F0: 32767}
                    _ = v36
                }
                v24 = v24
                _ = fi30
            }
            _ = fi29
        }
    case -1:
        fmt.Println(v1, v1)
    default:
        v28 = v28
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        var v38 uint64 = 4294967295
        var v39 int16 = v1
        _ = fi37
        _ = v38
        _ = v39
    }
    fmt.Println(v1, v28, v24)
    var v40 int32 = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v40
}
