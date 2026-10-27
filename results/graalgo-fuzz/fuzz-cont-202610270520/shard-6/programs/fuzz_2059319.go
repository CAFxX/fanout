package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int64, p1 int16) int64 {
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: ((p0 * -1) >= 1), F1: 28}
        p1 += p1
        _ = fi4
        _ = v5
    }
    switch 984 {
    case 2147483647:
        var v6 int64 = 9223372036854775807
        _ = v6
    case 42:
        v7 := GenS2{F0: false, F1: 42, F2: 0}
        _ = v7
    case -716:
        p1 += p1
    }
    var v8 float32 = 0.0
    _ = p0
    _ = p1
    _ = v8
    return p0
}
func genFn1() int {
    var v9 int32 = 42
    if true {
        var v10 int64 = 1
        switch 127 {
        case 127:
            fmt.Println(v9, v10)
        case -81:
            fmt.Println(v10, v9)
        default:
            var v11 uint64 = 42
            _ = v11
        }
        v10 += (v10 * (v10 / ((v10) | 1)))
        _ = v10
    } else {
        for fi12 := 0; fi12 < 15; fi12++ {
            for fi13 := 0; fi13 < 17; fi13++ {
                switch ((("" + "foo bar") + ("foo bar" + "hello")) + "hello") {
                case "a":
                    var v14 uint32 = 65535
                    _ = v14
                case "foo bar":
                    if true {
                        var v15 uint64 = 4294967295
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            v16 := GenS2{F0: false, F1: 781, F2: 42}
                            _ = v16
                        case "":
                            fmt.Println(fi12, v9, fi13)
                        case "\n\t":
                            v15 = (v15 % (((v15 | v15)) | 1))
                            v17 := []int{fi13, (-(fi13))}
                            fmt.Println(fi12, fi12, v9)
                            _ = v17
                        }
                        _ = v15
                    }
                case "\n\t":
                    v18 := []int{(fi12 * 32767)}
                    v9 -= (v9 / (((v9 - v9)) | 1))
                    if (255 > 255) {
                        fmt.Println(fi12, v9)
                        v18 = append(v18, (-(fi12)))
                    }
                    _ = v18
                }
                _ = fi13
            }
            _ = fi12
        }
    }
    v19 := []int{9223372036854775807, 321, 4294967295}
    _ = v19
    _ = v9
    return -336
}

func main() {
    v1 := []string{}
    var v2 int = 127
    var v3 float64 = float64(-2.25)
    if (!true) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v20 int16 = -1
        for fi21 := 0; fi21 < 17; fi21++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v22 string = ""
            for fi23 := 0; fi23 < 18; fi23++ {
                fmt.Println(v22, fi21)
                var v24 uint16 = 1
                var v25 uint32 = 0
                _ = len(v1)
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi21
            _ = v22
        }
        fmt.Println(v20, v20)
        _ = v20
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        var v27 int32 = 955
        v28 := GenS1{F0: ((30 <= genFn0(605, -1)) && true), F1: 100}
        _ = fi26
        _ = v27
        _ = v28
    }
    v2 = (v2 % ((v2) | 1))
    switch (v2 ^ genFn1()) {
    case 867:
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 int64 = genFn0(9223372036854775807, 0)
            var v31 string = ((("a" + "a") + "\n\t") + "a")
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi29
            _ = v30
            _ = v31
        }
    case -54:
        var v32 int8 = 16
        _ = v32
    case 4294967295:
        if (!false) {
            var v33 int16 = -791
            if (42 < 236) {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                v34 := GenS2{F0: false, F1: 935, F2: 2147483648}
                _ = v34
            } else {
                v3 = v3
                switch ("foo bar" + ("" + "hello")) {
                case "\n\t":
                    if (-71 >= 0) {
                        v3 = 1e10
                    } else {
                        v3 = v3
                        _ = len(v1)
                        var v35 uint64 = 4294967295
                        _ = v35
                    }
                default:
                    v3 = v3
                }
                v36 := GenS2{F0: (-79 != 42), F1: 42, F2: 4294967295}
                _ = v36
            }
            var v37 int = v2
            _ = v33
            _ = v37
        }
    }
    v38 := GenS3{F0: 0}
    v2 -= v2
    v2 = 0
    var v39 int32 = 0
    var v40 uint32 = 743
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
}
