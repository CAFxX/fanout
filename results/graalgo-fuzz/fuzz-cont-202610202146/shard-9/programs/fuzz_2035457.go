package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint32; F1 int; F2 int8 }
func genFn0() uint8 {
    var v4 int = -870
    v5 := []int{(v4 - (v4 >> ((705) & 63))), 40, 1}
    v4 = (v4 | -192)
    v4 = (v4 ^ -2147483648)
    _ = v4
    _ = v5
    return 255
}
func genFn1(p0 uint16, p1 bool, p2 bool) int64 {
    fmt.Println(p2)
    if false {
        v6 := GenS1{F0: 4294967295, F1: 32767, F2: 1000}
        p0 = p0
        p0 = (p0 * (p0 + 32767))
        _ = v6
    } else {
        v7 := []int{-179, -362, 32767}
        var v8 int16 = -913
        fmt.Println(p0, p2)
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    var v1 uint = 759
    v2 := make(map[string]int)
    var v3 float32 = (float32((0.1 / -2.25)) + float32((3.14159 - -2.25)))
    v9 := GenS1{F0: 2, F1: 32767, F2: 1}
    v3 -= v3
    switch int(v1) {
    case 1:
        v10 := []int{}
        _ = v10
    case 0:
        switch "hello" {
        case "foo bar":
            var v11 bool = false
            _ = v11
        case "\n\t":
            switch 128 {
            case 2147483647:
                if (true && ((310 <= 2) || (false || false))) {
                    v1 -= (^(v1))
                    var v12 int = -2147483648
                    _ = v12
                }
            case -629:
                v1 = (^(v1))
            default:
                v1 += 40
            }
        }
    case 128:
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint16 = 1
            var v15 string = ""
            v15 = "\n\t"
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v16 := []int{1}
    for fi17 := 0; fi17 < 14; fi17++ {
        if ("a" != "a") {
            v18 := GenS2{F0: "\n\t"}
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            switch ("\n\t" + "") {
            case "":
                v19 := GenS3{F0: 42, F1: 1000}
                _ = v19
            case "a":
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
            default:
                _ = len(v16)
            }
            _ = v18
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                v3 = float32(float64((v3 * v3)))
                switch (fi20 * -2147483648) {
                case -1:
                    if false {
                        var v21 bool = false
                        v22 := GenS4{F0: 687, F1: (fi20 >> ((v1) & 63)), F2: -1}
                        _ = v21
                        _ = v22
                    }
                }
                if len(v16) > 0 {
                    v16[0] = fi20
                }
                _ = fi20
            }
            fmt.Println(v1, v1, v1)
        }
        _ = fi17
    }
    fmt.Println(v3)
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 int32 = 42
        var v25 uint64 = 4294967295
        switch fi23 {
        case 32767:
            var v26 uint32 = 4294967295
            _ = v26
        case -664:
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 uint = v1
                var v29 uint64 = (v25 * (v25 % ((485) | 1)))
                _ = fi27
                _ = v28
                _ = v29
            }
        case 9223372036854775807:
            var v30 float64 = 0.1
            _ = v30
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    var v31 uint16 = 42
    v32 := v2["\n\t"]
    var v33 bool = true
    var v34 int8 = 100
    for fi35 := 0; fi35 < 19; fi35++ {
        for fi36 := 0; fi36 < 19; fi36++ {
            fmt.Println(fi35, v3, fi36)
            v37 := GenS3{F0: v31, F1: -32768}
            _ = fi36
            _ = v37
        }
        for fi38 := 0; fi38 < 9; fi38++ {
            if false {
                v16 = append(v16, v32)
                if v33 {
                    v32 = (fi38 - v32)
                    var v39 uint32 = 4294967295
                    _ = v39
                } else {
                    if true {
                        v31 -= (v31 - 65535)
                    } else {
                        var v40 uint64 = 4294967295
                        fmt.Println(v31, fi35)
                        v32 -= (fi38 >> ((v1) & 63))
                        v32 = fi35
                        _ = v40
                    }
                    delete(v2, "\n\t")
                }
            }
            _ = len(v2)
            v41 := GenS2{F0: "a"}
            _ = fi38
            _ = v41
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v9
}
