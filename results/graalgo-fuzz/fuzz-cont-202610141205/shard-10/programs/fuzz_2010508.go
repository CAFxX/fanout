package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 string }
func genFn0(p0 string) int8 {
    fmt.Println(p0)
    var v5 string = p0
    v6 := []int{0, 32768, 345, -1}
    _ = p0
    _ = v5
    _ = v6
    return 100
}
func genFn1() bool {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint64 = 4294967295
        var v9 uint32 = 677
        fmt.Println(v9, v8, fi7)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 65535}
    v11 := GenS2{F0: 65535, F1: 255}
    var v12 float32 = 3.14159
    _ = v10
    _ = v11
    _ = v12
    return true
}

func main() {
    var v1 bool = (false && false)
    var v2 uint64 = 0
    v3 := []uint32{143, 1, 149, 65535}
    var v4 uint32 = 42
    v13 := GenS2{F0: 1, F1: 216}
    v3 = append(v3, 634)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v4
    if v1 {
        v2 -= v2
        for fi14 := 0; fi14 < 7; fi14++ {
            v4 -= (v4 * (^(v4)))
            v15 := GenS1{F0: 42}
            v4 += (v4 | v4)
            _ = fi14
            _ = v15
        }
        var v16 float32 = ((0.1 + -2.25) / -2.25)
        _ = v16
    }
    if len(v3) > 0 {
        v3[0] = uint32(v2)
    }
    v2 = v2
    switch ("hello" + "hello") {
    case "foo bar":
        var v17 float64 = 0.0
        _ = v17
    case "":
        for fi18 := 0; fi18 < 15; fi18++ {
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS3{F0: ((("" + "a") + "hello") + "foo bar")}
                var v21 uint32 = (v4 << ((4294967295) & 63))
                v22 := GenS3{F0: ("hello" + (("hello" + "hello") + "hello"))}
                _ = len(v3)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v4 = (v4 >> ((516) & 63))
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            _ = fi18
        }
    case "\n\t":
        fmt.Println(v4, v2)
    default:
        v3 = append(v3, v4)
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        v4 = v4
        v4 = (v4 + (v4 >> ((413) & 63)))
        _ = fi23
    }
    v24 := GenS2{F0: 42, F1: 42}
    if len(v3) > 0 {
        v3[0] = (^(v4))
    }
    v2 = v2
    for fi25 := 0; fi25 < 10; fi25++ {
        switch ("\n\t" + "a") {
        case "":
            v3 = append(v3, (v4 * v4))
        case "foo bar":
            if (!(0 > 0)) {
                if v1 {
                    v1 = v1
                } else {
                    v2 -= (v2 >> ((2) & 63))
                    v3 = append(v3, v4)
                    v3 = append(v3, v4)
                }
                for fi26 := 0; fi26 < 15; fi26++ {
                    v27 := GenS1{F0: 431}
                    v4 = (v4 ^ v4)
                    _ = fi26
                    _ = v27
                }
                v1 = (!(v1 && ((v1 || (!true)) && (!v1))))
            } else {
                var v28 uint8 = 42
                for fi29 := 0; fi29 < 20; fi29++ {
                    if (-429 < -697) {
                        var v30 uint8 = (v28 % ((128) | 1))
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v2 += (v2 * v2)
                        _ = v30
                    } else {
                        var v31 bool = (v1 && true)
                        var v32 uint = 42
                        v32 = (v32 >> (((v32 >> (((v32 % ((255) | 1))) & 63))) & 63))
                        var v33 uint = (^(v32))
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    if v1 {
                        if (fi29 == fi29) {
                            var v34 bool = (!(32767 >= 32767))
                            var v35 int = (fi29 | (^(fi29)))
                            _ = v34
                            _ = v35
                        } else {
                            _ = len(v3)
                        }
                    } else {
                        v36 := GenS1{F0: 255}
                        var v37 uint = 42
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            v38 := GenS1{F0: 42}
                            _ = v38
                        case "hello":
                            v1 = (fi29 > 32767)
                        case "a":
                            var v39 int = int(v37)
                            var v40 int8 = 89
                            var v41 uint = (v37 % ((uint(v2)) | 1))
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        _ = v36
                        _ = v37
                    }
                    _ = fi29
                }
                _ = v28
            }
        case "hello":
            var v42 string = ("foo bar" + "\n\t")
            _ = v42
        default:
            fmt.Println(v2, v1, v2)
        }
        _ = fi25
    }
    v43 := GenS2{F0: uint16(v2), F1: 65535}
    fmt.Println(v1, v1)
    var v44 string = "foo bar"
    v4 = 150
    v45 := GenS3{F0: ("\n\t" + (("foo bar" + (("foo bar" + ((v44 + (v44 + v44)) + v44)) + "")) + v44))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
