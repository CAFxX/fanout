package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 bool }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0() int32 {
    v4 := []int{9223372036854775807}
    var v5 string = "foo bar"
    v5 = ""
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 uint32) int8 {
    var v6 int64 = 42
    v6 = (^(v6))
    _ = p0
    _ = v6
    return -114
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 104
    v3 := make(map[uint32]bool)
    if (3.14159 < float32(0.0)) {
        switch 127 {
        case -19:
            var v7 uint = 428
            _ = v7
        case 4294967295:
            var v8 uint32 = 42
            _ = v8
        case 1:
            v9 := GenS1{F0: 441}
            _ = v9
        default:
            v10 := []int{}
            _ = v10
        }
    }
    delete(v1, -781)
    switch -1 {
    case -1:
        fmt.Println(v2)
    default:
        fmt.Println(v2)
    }
    if (32767 != 653) {
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(v2, v2)
            _ = fi11
        }
        if ((-2.25 + (3.14159 - -2.25)) == 3.14159) {
            v12 := GenS2{F0: 169}
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS3{F0: 42, F1: 302, F2: true}
                _ = fi13
                _ = v14
            }
            _ = v12
        }
    } else {
        v2 = (v2 ^ -1)
        for fi15 := 0; fi15 < 16; fi15++ {
            v2 = 42
            _ = fi15
        }
        if true {
            v16 := GenS4{F0: 255, F1: -9223372036854775808}
            fmt.Println(v2)
            _ = v16
        }
    }
    var v17 int32 = genFn0()
    v17 += v17
    var v18 int = -9223372036854775808
    if (true && false) {
        v19 := GenS1{F0: 310}
        var v20 uint8 = 1
        _ = v19
        _ = v20
    }
    var v21 float32 = 1.5
    v22 := GenS5{F0: 42, F1: ((!((42 < 2) && ((v2 > (-(v2))) && ("a" != ("hello" + "\n\t"))))) || true), F2: 2}
    v23 := []int{-956, -1, v18, (v18 % (((v18 - 127)) | 1))}
    v3[4294967295] = true
    if (42 == 262) {
        for fi24 := 0; fi24 < 18; fi24++ {
            v21 += (v21 * 1e10)
            _ = fi24
        }
        if (false || false) {
            for fi25 := 0; fi25 < 19; fi25++ {
                if ((!false) && (65535 < 1)) {
                    var v26 int64 = -357
                    for fi27 := 0; fi27 < 13; fi27++ {
                        var v28 int16 = -32768
                        v17 -= 42
                        delete(v3, 764)
                        v2 = -36
                        v17 = 42
                        _ = fi27
                        _ = v28
                    }
                    var v29 int32 = genFn0()
                    _ = v26
                    _ = v29
                } else {
                    if false {
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        var v30 float32 = float32(-2.25)
                        switch "hello" {
                        case "\n\t":
                            v31 := GenS1{F0: 65535}
                            _ = v31
                        case "":
                            var v32 string = ("a" + ((("hello" + "foo bar") + ("a" + ("a" + ((("foo bar" + "\n\t") + "") + ("a" + "\n\t"))))) + "foo bar"))
                            _ = v32
                        }
                        _ = v30
                    }
                }
                switch fi25 {
                case -762:
                    switch ((("\n\t" + "hello") + "\n\t") + (("hello" + "a") + "\n\t")) {
                    case "\n\t":
                        v18 += 776
                    case "a":
                        switch "" {
                        case "":
                            v33 := GenS3{F0: 32767, F1: 962, F2: false}
                            _ = v33
                        default:
                            for fi34 := 0; fi34 < 15; fi34++ {
                                v18 -= (fi25 * fi34)
                                switch "hello" {
                                case "a":
                                    var v35 uint8 = 0
                                    _ = v35
                                case "":
                                    v17 -= genFn0()
                                default:
                                    fmt.Println(v18, fi25)
                                }
                                switch "" {
                                case "foo bar":
                                    switch (v18 >> ((255) & 63)) {
                                    case 32768:
                                        v21 += v21
                                    }
                                case "a":
                                    var v36 string = ("a" + "")
                                    _ = v36
                                }
                                _ = fi34
                            }
                        }
                    case "foo bar":
                        delete(v3, 2)
                    }
                case 684:
                    var v37 int32 = v17
                    _ = v37
                case -248:
                    v21 = 1e10
                }
                _ = fi25
            }
            for fi38 := 0; fi38 < 20; fi38++ {
                delete(v3, 2147483648)
                v21 -= (v21 + 1.5)
                var v39 float32 = -2.25
                _ = fi38
                _ = v39
            }
        } else {
            fmt.Println(v18)
            var v40 int = -1
            _ = v40
        }
    }
    v17 += int32(v18)
    var v41 int8 = -1
    var v42 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v41
    _ = v42
}
