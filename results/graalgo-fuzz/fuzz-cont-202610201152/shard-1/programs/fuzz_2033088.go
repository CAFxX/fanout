package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 string; F2 int }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0(p0 int32, p1 int, p2 uint) int64 {
    var v7 int = p1
    if (0 == 83) {
        v7 += v7
    } else {
        var v8 int8 = 0
        v8 = (v8 & v8)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 595
}

func main() {
    var v1 uint32 = 67
    var v2 uint8 = 127
    var v3 int16 = 1000
    v4 := make(map[int]string)
    var v5 int16 = (v3 | (v3 * 1))
    var v6 uint32 = 494
    if true {
        if true {
            v2 = uint8(v1)
            var v9 float32 = ((-2.25 * (1.5 - float32(1.5))) / 1.0)
            _ = v9
        } else {
            if (0 >= 0) {
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS1{F0: 65535}
                    if (true && (true && false)) {
                        for fi12 := 0; fi12 < 12; fi12++ {
                            v6 += (^(v6))
                            switch -2147483648 {
                            case 128:
                                if false {
                                    v13 := []int{}
                                    _ = v13
                                } else {
                                    var v14 bool = (42 < 0)
                                    if false {
                                        v5 += (v3 | (v5 | v5))
                                        var v15 uint32 = (v1 * (v1 % ((v6) | 1)))
                                        var v16 float64 = (1e10 / float64(1.5))
                                        _ = v15
                                        _ = v16
                                    }
                                    _ = v14
                                }
                            case 2147483647:
                                _ = len(v4)
                            }
                            fmt.Println(v3)
                            _ = fi12
                        }
                        v1 -= v6
                        v2 -= 183
                    } else {
                        var v17 int16 = -908
                        var v18 bool = ((true || (0 == 2)) || false)
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(v5, v1, fi10)
                    _ = fi10
                    _ = v11
                }
                v1 = (v6 ^ (v6 << ((741) & 63)))
                v19 := []int{}
                _ = v19
            }
            if true {
                for fi20 := 0; fi20 < 4; fi20++ {
                    v21 := []int{32767, (-(fi20)), (^(fi20))}
                    fmt.Println(v2, v3)
                    _ = fi20
                    _ = v21
                }
            } else {
                var v22 string = (("a" + "foo bar") + "")
                _ = v22
            }
        }
    }
    v23 := []int{}
    v3 -= 0
    for fi24 := 0; fi24 < 19; fi24++ {
        v23 = append(v23, (-(fi24)))
        for fi25 := 0; fi25 < 5; fi25++ {
            switch ("foo bar" + "foo bar") {
            case "":
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
            case "a":
                fmt.Println(v2, v1)
            case "foo bar":
                fmt.Println(fi25, fi24)
            }
            _ = fi25
        }
        _ = fi24
    }
    v1 += v6
    v3 += (v5 | v5)
    v6 += (v1 ^ v1)
    switch 1 {
    case 1:
        v5 += (v5 | v3)
    default:
        fmt.Println(v1)
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        v3 -= (v5 ^ int16(fi26))
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := GenS2{F0: -128, F1: (("hello" + (("foo bar" + "foo bar") + ("a" + "foo bar"))) + "hello"), F2: fi27}
            fmt.Println(v2)
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    v29 := GenS1{F0: 792}
    for fi30 := 0; fi30 < 1; fi30++ {
        v31 := GenS3{F0: 42, F1: 1}
        _ = fi30
        _ = v31
    }
    if true {
        v23 = append(v23, 19)
        v4[291] = "\n\t"
    }
    fmt.Println(v3)
    var v32 uint = 18446744073709551615
    _ = len(v23)
    if len(v23) > 0 {
        v23[0] = 2147483647
    }
    v33 := GenS3{F0: (v2 >> (((^(v32))) & 63)), F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
