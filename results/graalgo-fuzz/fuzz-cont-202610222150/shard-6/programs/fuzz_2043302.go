package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int16; F1 int16; F2 int32 }
type GenS3 struct { F0 uint }
func genFn0() string {
    var v5 uint64 = 2
    var v6 uint = 4294967295
    v6 = (v6 / (((v6 ^ v6)) | 1))
    if true {
        v6 = v6
        v7 := []int{4294967295}
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(fi8, v5)
        v6 += v6
        _ = fi8
    }
    _ = v5
    _ = v6
    return ((("foo bar" + "hello") + "") + (("a" + "a") + ""))
}
func genFn1(p0 uint64, p1 uint16) uint8 {
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := []int{fi10, fi9, fi10}
            _ = fi10
            _ = v11
        }
        fmt.Println(p1, p1, fi9)
        if true {
            v12 := []int{(fi9 << ((1) & 63))}
            var v13 int64 = -879
            _ = v12
            _ = v13
        } else {
            v14 := []int{(fi9 + -476)}
            var v15 int16 = -1
            _ = len(v14)
            _ = v14
            _ = v15
        }
        _ = fi9
    }
    p1 = (p1 + 794)
    fmt.Println(p1)
    var v16 string = "\n\t"
    _ = p0
    _ = p1
    _ = v16
    return 255
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 127
    var v3 uint8 = 128
    v4 := make(map[int]uint32)
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS1{F0: (fi17 << ((255) & 63)), F1: false}
        for fi19 := 0; fi19 < 8; fi19++ {
            if (2 > 2147483648) {
                var v20 int32 = 42
                v21 := []int{fi19, (fi19 * (fi19 + 9223372036854775807)), 0, (fi19 >> ((4294967295) & 63))}
                _ = v20
                _ = v21
            }
            if true {
                v2 = 177
                if (fi19 > (fi17 / (((fi19 >> ((0) & 63))) | 1))) {
                    if false {
                        if ((false || false) && (1e10 == 0.1)) {
                            v22 := []int{}
                            var v23 uint16 = 32767
                            _ = v22
                            _ = v23
                        } else {
                            v3 = genFn1(555, 271)
                            fmt.Println(v2, v3, v3)
                        }
                        fmt.Println(v2, v2, v2)
                        delete(v1, (fi19 + fi17))
                    }
                    var v24 int32 = 115
                    v24 += v24
                    if false {
                        var v25 float32 = (-2.25 - -2.25)
                        _ = v25
                    }
                    var v26 bool = false
                    _ = v24
                    _ = v26
                } else {
                    v1[128] = false
                    v2 = 0
                }
            }
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    if true {
        v27 := GenS1{F0: -2147483648, F1: ((555 <= 111) && false)}
        if (65535 != 32767) {
            v2 -= v3
            v3 -= v3
            v2 = genFn1(675, 1)
        }
        fmt.Println(v3, v2)
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v2 -= (v2 | v3)
        var v29 float64 = 1.5
        _ = fi28
        _ = v29
    }
    var v30 uint16 = 776
    v2 -= v3
    fmt.Println(v3)
    var v31 int = -9223372036854775808
    v32 := GenS2{F0: 53, F1: -32768, F2: -2147483648}
    v33 := GenS3{F0: 992}
    var v34 int8 = -128
    var v35 uint = 445
    switch (("\n\t" + (("a" + "") + "a")) + "\n\t") {
    case "":
        var v36 uint16 = (v30 - (v30 / ((v30) | 1)))
        _ = v36
    case "\n\t":
        if true {
            var v37 uint = v35
            v30 = v30
            _ = v37
        } else {
            var v38 uint64 = 18446744073709551615
            _ = v38
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
