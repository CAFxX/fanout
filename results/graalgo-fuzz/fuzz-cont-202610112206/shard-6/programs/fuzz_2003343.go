package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int, p2 int) string {
    fmt.Println(p2)
    p0 += (p0 - 418)
    if false {
        v7 := []int{p2, p1, p1}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        fmt.Println(p2)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := []int{(fi8 << ((0) & 63)), fi8, p2}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := []int{p2, 32768}
            var v12 int = p2
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (((("hello" + "\n\t") + "foo bar") + (("hello" + "") + (("a" + ("a" + "hello")) + ""))) + ("foo bar" + ""))
}
func genFn1(p0 uint8, p1 int, p2 uint16) string {
    v13 := GenS1{F0: 32767}
    p2 = (p2 / ((p2) | 1))
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 uint = 18446744073709551615
        v15 = 977
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return "hello"
}

func main() {
    var v1 int16 = -506
    var v2 string = "\n\t"
    v3 := make(map[uint32]uint32)
    var v4 int32 = 1
    var v5 uint32 = 600
    var v6 bool = (!(v2 < (v2 + v2)))
    v16 := []int{-612, -2147483648, -9223372036854775808, 32767}
    v3[(^(v5))] = 778
    var v17 int16 = v1
    fmt.Println(v1, v1, v1)
    v18 := GenS1{F0: 1}
    var v19 uint = 0
    switch v2 {
    case "":
        v16 = append(v16, 0)
    case "hello":
        var v20 int64 = -646
        _ = v20
    }
    if (0 <= 0) {
        if false {
            switch (v2 + v2) {
            case "":
                var v21 int32 = v4
                _ = v21
            case "\n\t":
                var v22 uint32 = (v5 * (v5 & v5))
                _ = v22
            }
            v1 += v17
            v23 := GenS1{F0: 1}
            _ = v23
        } else {
            v17 = v1
            fmt.Println(v4)
            var v24 int64 = -9223372036854775808
            _ = v24
        }
        delete(v3, (^(v5)))
    } else {
        v3[(v5 % (((v5 / ((42) | 1))) | 1))] = (v5 * 471)
        for fi25 := 0; fi25 < 2; fi25++ {
            if ((v5 << ((v19) & 63)) > (v5 | (v5 >> ((v19) & 63)))) {
                var v26 uint8 = 181
                if v6 {
                    switch fi25 {
                    case 610:
                        v27 := v3[v5]
                        _ = v27
                    case 32767:
                        v4 -= v4
                    default:
                        if false {
                            v28 := GenS1{F0: 335}
                            _ = v28
                        } else {
                            if false {
                                v1 = 42
                                _ = len(v3)
                                v17 = (^(v1))
                            }
                            var v29 int = (fi25 - (fi25 / (((fi25 + fi25)) | 1)))
                            if len(v16) > 0 {
                                v16[0] = fi25
                            }
                            _ = v29
                        }
                    }
                    v26 -= uint8(v5)
                    v6 = v6
                } else {
                    v30 := GenS1{F0: 872}
                    fmt.Println(fi25)
                    _ = v30
                }
                _ = v26
            } else {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                v31 := GenS1{F0: 2}
                _ = v31
            }
            v32 := GenS2{F0: 330}
            _ = fi25
            _ = v32
        }
    }
    v33 := GenS3{F0: 255}
    if false {
        v2 = v2
    } else {
        v3[(v5 + (v5 + 698))] = v5
    }
    v2 = (("" + v2) + "a")
    v34 := GenS3{F0: 0}
    _ = len(v3)
    var v35 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
