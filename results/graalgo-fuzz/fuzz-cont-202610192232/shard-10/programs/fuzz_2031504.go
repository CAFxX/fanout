package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 uint }
func genFn0(p0 int8, p1 int32) uint32 {
    p0 = (p0 - (p0 >> ((274) & 63)))
    switch "foo bar" {
    case "foo bar":
        switch "a" {
        case "foo bar":
            if (0.1 != (1.5 * -2.25)) {
                var v4 uint = 1
                _ = v4
            } else {
                if (0 == 2) {
                    if (42 == 42) {
                        if true {
                            var v5 uint = 4294967295
                            v6 := []int{9223372036854775807}
                            v7 := GenS1{F0: -9223372036854775808}
                            _ = v5
                            _ = v6
                            _ = v7
                        }
                        v8 := GenS1{F0: 4294967295}
                        v9 := []int{833, 1}
                        _ = v8
                        _ = v9
                    }
                }
            }
        case "":
            if true {
                var v10 bool = (255 <= 2)
                _ = v10
            }
        case "hello":
            v11 := GenS1{F0: 2147483647}
            _ = v11
        default:
            p0 -= p0
        }
    case "\n\t":
        p1 = (-(p1))
    case "a":
        if true {
            p0 += 100
            for fi12 := 0; fi12 < 11; fi12++ {
                p0 += 0
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 string = ("hello" + "a")
                    if true {
                        v14 = (((v14 + v14) + "hello") + v14)
                        if false {
                            v14 = v14
                            var v15 string = (("foo bar" + (("foo bar" + v14) + ((("hello" + v14) + v14) + "a"))) + v14)
                            var v16 uint = 42
                            _ = v15
                            _ = v16
                        } else {
                            if (!((!false) || (910 >= 0))) {
                                if (false || (877 > -1)) {
                                    fmt.Println(p1, fi13, fi12)
                                    fmt.Println(fi12, fi12, p1)
                                    if false {
                                        v17 := GenS1{F0: fi13}
                                        _ = v17
                                    } else {
                                        p0 += -128
                                        v18 := GenS1{F0: fi12}
                                        _ = v18
                                    }
                                    if ((4294967295 == 1) && (false && false)) {
                                        if true {
                                            v19 := []int{32768}
                                            p1 = 1
                                            p0 -= p0
                                            v14 = ""
                                            _ = v19
                                        }
                                    }
                                } else {
                                    switch v14 {
                                    case "foo bar":
                                        p0 += (p0 % ((p0) | 1))
                                    case "":
                                        v14 = "\n\t"
                                    }
                                    p0 = -128
                                }
                            }
                            p0 -= 100
                            switch ((v14 + v14) + "hello") {
                            case "hello":
                                fmt.Println(p0)
                            case "\n\t":
                                p1 = p1
                            case "foo bar":
                                fmt.Println(v14)
                            default:
                                if (1 == 42) {
                                    var v20 uint32 = 0
                                    fmt.Println(p1, p1, fi12)
                                    _ = v20
                                }
                            }
                        }
                        switch v14 {
                        case "a":
                            if true {
                                v21 := []int{(fi13 + fi12)}
                                v22 := GenS2{F0: (fi12 + fi13), F1: 255, F2: 6}
                                _ = v21
                                _ = v22
                            } else {
                                var v23 int16 = 1
                                var v24 int16 = (v23 ^ v23)
                                var v25 uint8 = 1
                                var v26 uint8 = v25
                                _ = v23
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                        case "foo bar":
                            if (1 != 9223372036854775807) {
                                var v27 float32 = 1e10
                                var v28 string = v14
                                _ = v27
                                _ = v28
                            }
                        case "":
                            if ((false && true) && (p1 <= -1)) {
                                v29 := GenS1{F0: (fi12 * 1)}
                                var v30 float32 = 0.0
                                v31 := []int{(fi13 * (fi13 << ((65535) & 63)))}
                                _ = v29
                                _ = v30
                                _ = v31
                            } else {
                                var v32 int = fi12
                                _ = v32
                            }
                        }
                    } else {
                        v33 := GenS3{F0: 731}
                        _ = v33
                    }
                    v34 := GenS4{F0: 32767}
                    var v35 int16 = 998
                    _ = fi13
                    _ = v14
                    _ = v34
                    _ = v35
                }
                _ = fi12
            }
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 int64 = -520
                _ = fi36
                _ = v37
            }
        }
    }
    if false {
        v38 := GenS2{F0: 32767, F1: 1, F2: 128}
        v39 := GenS3{F0: 927}
        _ = v38
        _ = v39
    } else {
        v40 := []int{4294967295, 315, -2147483648}
        _ = v40
    }
    var v41 int32 = (p1 | 1)
    _ = p0
    _ = p1
    _ = v41
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 float32 = float32(1e10)
    var v3 uint8 = 127
    var v42 float32 = v2
    v43 := GenS5{F0: 1, F1: uint(v3)}
    v1 += 1
    if true {
        fmt.Println(v42, v1, v42)
        v42 += v2
        fmt.Println(v2)
    }
    v2 -= v42
    var v44 float32 = v2
    var v45 uint = 174
    v46 := []int{4294967295, 128, -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
