package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 string; F1 uint; F2 uint8 }
func genFn0() int8 {
    var v7 int = 32768
    v8 := []int{(v7 + -594)}
    _ = v7
    _ = v8
    return -2
}
func genFn1(p0 string) uint16 {
    if (-32768 >= 1000) {
        switch p0 {
        case "foo bar":
            fmt.Println(p0)
        case "a":
            v9 := GenS1{F0: -975}
            _ = v9
        }
        var v10 int32 = 2147483647
        _ = v10
    } else {
        switch (("hello" + p0) + p0) {
        case "\n\t":
            for fi11 := 0; fi11 < 13; fi11++ {
                if false {
                    v12 := []int{(fi11 / ((128) | 1)), fi11, fi11}
                    switch p0 {
                    case "hello":
                        p0 = p0
                    case "\n\t":
                        var v13 uint8 = 128
                        _ = v13
                    }
                    var v14 uint8 = 200
                    _ = v12
                    _ = v14
                } else {
                    p0 = (p0 + p0)
                }
                switch (((p0 + p0) + p0) + p0) {
                case "hello":
                    p0 = p0
                case "foo bar":
                    p0 = ((("foo bar" + (p0 + p0)) + p0) + "a")
                default:
                    v15 := GenS1{F0: 1000}
                    _ = v15
                }
                _ = fi11
            }
        default:
            p0 = p0
        }
    }
    switch ((p0 + "a") + (p0 + p0)) {
    case "hello":
        v16 := GenS2{F0: 36, F1: 0}
        _ = v16
    case "":
        p0 = (p0 + (p0 + p0))
    }
    _ = p0
    return 65535
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    v4 := []int{}
    v5 := make(map[int]uint32)
    v6 := make(map[string]string)
    var v17 uint8 = 42
    if (false && (true || (18446744073709551615 > 42))) {
        for fi18 := 0; fi18 < 8; fi18++ {
            if (1 >= 785) {
                var v19 float64 = (-2.25 + -2.25)
                fmt.Println(v19, fi18, v1)
                for fi20 := 0; fi20 < 14; fi20++ {
                    if (v17 <= (v17 + v17)) {
                        v21 := GenS3{F0: "foo bar", F1: 593, F2: (v17 ^ v17)}
                        switch "a" {
                        case "":
                            v1 = (-(v1))
                        case "\n\t":
                            if len(v4) > 0 {
                                v4[0] = fi18
                            }
                        }
                        _ = v21
                    } else {
                        if false {
                            fmt.Println(fi18)
                            v17 -= (v17 << ((18446744073709551615) & 63))
                        } else {
                            fmt.Println(v1, v1, fi20)
                        }
                    }
                    var v22 uint64 = 4294967295
                    var v23 uint16 = 32767
                    _ = fi20
                    _ = v22
                    _ = v23
                }
                _ = v19
            } else {
                for fi24 := 0; fi24 < 4; fi24++ {
                    if (v17 > v17) {
                        var v25 uint = 65535
                        v17 = (v17 - (v17 % (((v17 * v17)) | 1)))
                        _ = v25
                    } else {
                        if (-57 > 710) {
                            v17 += 1
                            v1 -= v1
                            fmt.Println(v1)
                        }
                        _ = len(v4)
                        switch "foo bar" {
                        case "\n\t":
                            var v26 bool = (!true)
                            _ = v26
                        case "":
                            fmt.Println(fi18)
                        default:
                            var v27 uint8 = v17
                            _ = v27
                        }
                    }
                    var v28 uint32 = 141
                    _ = fi24
                    _ = v28
                }
            }
            switch ("hello" + ("hello" + (("foo bar" + ("a" + ("foo bar" + "a"))) + ("\n\t" + "")))) {
            case "foo bar":
                var v29 uint16 = 32767
                _ = v29
            case "":
                var v30 int8 = 82
                _ = v30
            default:
                v17 = (v17 | v17)
            }
            fmt.Println(v17, v1, v1)
            _ = fi18
        }
    } else {
        fmt.Println(v17, v1)
        v5[(v1 >> ((255) & 63))] = 19
        v1 = (v1 * v1)
    }
    _ = len(v2)
    v31 := GenS2{F0: 32767, F1: 0}
    v17 += 0
    v32 := GenS3{F0: ("foo bar" + "\n\t"), F1: 42, F2: 1}
    v5[v1] = 2
    switch v1 {
    case 42:
        for fi33 := 0; fi33 < 5; fi33++ {
            fmt.Println(v1)
            delete(v2, "")
            _ = fi33
        }
    }
    var v34 string = ("foo bar" + "foo bar")
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 string = (v34 + v34)
        _ = fi35
        _ = v36
    }
    var v37 string = "\n\t"
    for fi38 := 0; fi38 < 1; fi38++ {
        var v39 bool = (fi38 >= (v1 * 128))
        for fi40 := 0; fi40 < 17; fi40++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    fmt.Println(v34, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
