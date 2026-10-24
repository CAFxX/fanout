package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 uint8, p1 int16, p2 uint) bool {
    p2 += (p2 & (p2 / ((p2) | 1)))
    for fi4 := 0; fi4 < 9; fi4++ {
        switch "foo bar" {
        case "hello":
            p0 -= 57
        }
        _ = fi4
    }
    p1 = p1
    p1 = (p1 >> ((p2) & 63))
    _ = p0
    _ = p1
    _ = p2
    return ("hello" < ("" + "hello"))
}
func genFn1(p0 uint32, p1 int, p2 uint32) int16 {
    v5 := []int{p1, p1, p1}
    var v6 float32 = 0.1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(p1)
    var v7 float64 = (1e10 * 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 195
}

func main() {
    var v1 int32 = 1
    v2 := []int{}
    var v3 string = ""
    v3 = "foo bar"
    var v8 float32 = 0.0
    switch 4294967295 {
    case -2147483648:
        var v9 int64 = 1
        _ = v9
    case -9223372036854775808:
        v8 = v8
    case 1:
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 bool = true
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v1, v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v12 uint64 = 1
    switch 32767 {
    case 2147483647:
        var v13 string = v3
        _ = v13
    case -9223372036854775808:
        if false {
            var v14 string = ""
            _ = v14
        } else {
            var v15 uint64 = 107
            var v16 uint8 = 42
            _ = v15
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 9; fi17++ {
            switch -786 {
            case 434:
                fmt.Println(v3, v3, fi17)
            case 9223372036854775807:
                for fi18 := 0; fi18 < 10; fi18++ {
                    v19 := GenS1{F0: (680 >= 2147483648), F1: v12}
                    var v20 int64 = 9223372036854775807
                    _ = len(v2)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case 32768:
                if false {
                    var v21 float32 = float32(-2.25)
                    switch (^(fi17)) {
                    case 32767:
                        var v22 float32 = v8
                        _ = v22
                    }
                    _ = v21
                } else {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        switch (fi23 >> ((42) & 63)) {
                        case 0:
                            var v24 uint16 = 2
                            _ = v24
                        case -1:
                            var v25 float32 = 0.1
                            _ = v25
                        case -386:
                            if false {
                                var v26 uint64 = (v12 * (^(v12)))
                                if (false && false) {
                                    v3 = v3
                                    v1 = v1
                                }
                                _ = v26
                            }
                        default:
                            var v27 uint32 = 2147483647
                            _ = v27
                        }
                        v28 := GenS1{F0: (true && false), F1: v12}
                        var v29 uint64 = 0
                        v30 := GenS1{F0: false, F1: v12}
                        var v31 int16 = 32767
                        _ = fi23
                        _ = v28
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                }
            default:
                fmt.Println(fi17, v8)
            }
            var v32 int64 = 9223372036854775807
            _ = fi17
            _ = v32
        }
    }
    fmt.Println(v8, v3, v8)
    fmt.Println(v1)
    v33 := GenS1{F0: true, F1: (v12 / (((v12 % (((v12 ^ 0)) | 1))) | 1))}
    var v34 int16 = -32768
    for fi35 := 0; fi35 < 11; fi35++ {
        var v36 uint8 = 128
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        switch ("" + v3) {
        case "":
            fmt.Println(v3)
        }
        v2 = append(v2, fi37)
        var v38 uint32 = 702
        _ = fi37
        _ = v38
    }
    v39 := GenS2{F0: ((v3 + v3) + ("foo bar" + "foo bar")), F1: 65535}
    var v40 float64 = 3.14159
    v12 -= v12
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
    _ = v40
    _ = v8
}
