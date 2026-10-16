package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 string }
func genFn0() uint {
    for fi7 := 0; fi7 < 1; fi7++ {
        if true {
            v8 := []int{}
            _ = v8
        }
        var v9 string = "hello"
        var v10 int32 = 42
        _ = fi7
        _ = v10
        _ = v9
    }
    switch "\n\t" {
    case "a":
        var v11 float32 = 0.0
        _ = v11
    case "":
        switch 42 {
        case 1:
            v12 := GenS1{F0: 100, F1: 679}
            _ = v12
        default:
            v13 := GenS1{F0: -128, F1: 0}
            _ = v13
        }
    }
    var v14 uint8 = 111
    _ = v14
    return 826
}
func genFn1(p0 bool, p1 uint) int {
    switch -144 {
    case 2147483647:
        var v15 uint64 = 911
        _ = v15
    }
    p0 = p0
    _ = p0
    _ = p1
    return 127
}

func main() {
    var v1 uint16 = 65535
    var v2 string = "foo bar"
    v3 := make(map[string]int)
    var v4 uint32 = 2
    var v5 int = 725
    var v6 uint8 = 1
    switch v2 {
    case "foo bar":
        v2 = (v2 + v2)
    case "\n\t":
        v16 := GenS1{F0: 0, F1: -1}
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := []int{(fi17 / (((fi17 - 128)) | 1)), fi17, genFn1(true, 565), (fi17 ^ fi17)}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 int8 = 100
        var v21 int32 = 0
        v4 -= v4
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v6)
    var v22 uint = 608
    if true {
        var v23 int8 = -107
        _ = v23
    } else {
        v24 := v3[v2]
        _ = v24
    }
    if false {
        v25 := []int{v5, v5}
        var v26 int = (v5 ^ (v5 << (((v22 % ((459) | 1))) & 63)))
        _ = v25
        _ = v26
    } else {
        switch (v5 | (-(v5))) {
        case 4294967295:
            v4 += 1
        }
        if (int8(v5) == 1) {
            var v27 string = v2
            for fi28 := 0; fi28 < 1; fi28++ {
                fmt.Println(v1, v2)
                var v29 int8 = 0
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 8; fi30++ {
                if false {
                    v31 := []int{(fi30 / ((127) | 1))}
                    if true {
                        v32 := v3[v27]
                        var v33 uint8 = v6
                        _ = v32
                        _ = v33
                    } else {
                        var v34 int8 = 0
                        if (int32(v4) == int32(v34)) {
                            v5 -= 636
                            var v35 int64 = 3
                            _ = v35
                        } else {
                            var v36 string = "hello"
                            var v37 string = v2
                            _ = v36
                            _ = v37
                        }
                        _ = v34
                    }
                    v2 = v27
                    _ = v31
                }
                v22 = (^(v22))
                v22 = v22
                _ = fi30
            }
            _ = v27
        } else {
            v1 -= (v1 & (v1 ^ uint16(v22)))
            v6 += (v6 / (((^(v6))) | 1))
        }
    }
    v38 := GenS2{F0: (v2 + v2)}
    v5 = -566
    v39 := GenS2{F0: (("" + v2) + ("hello" + (((v2 + v2) + ("a" + v2)) + ("hello" + ""))))}
    v4 -= (v4 - v4)
    v40 := []int{(v5 + v5), genFn1(((0 > 42) || false), v22), int(v4), -260}
    v1 = v1
    v5 += (v5 * (v5 << ((620) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
