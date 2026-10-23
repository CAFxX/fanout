package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 string }
func genFn0(p0 bool, p1 bool, p2 int32) uint {
    v4 := []int{127}
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int64 = -438
        var v7 int32 = p2
        fmt.Println(v6, p0, fi5)
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = (127 >= 32768)
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1(p0 int32, p1 string, p2 uint64) string {
    p1 = p1
    var v8 uint = 994
    var v9 int16 = -716
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := []int{(fi10 + fi10), fi10}
        if (705 > (v9 & (v9 * (v9 / ((0) | 1))))) {
            p1 = "\n\t"
            v12 := GenS1{F0: (v9 | 42)}
            _ = v12
        } else {
            v11 = append(v11, 325)
        }
        var v13 float32 = 3.14159
        _ = fi10
        _ = v11
        _ = v13
    }
    v9 = (v9 & (v9 - -431))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return ("" + "foo bar")
}

func main() {
    v1 := []uint32{65535}
    var v2 uint8 = 255
    var v3 float32 = 1.5
    var v14 uint16 = 839
    v15 := GenS2{F0: 0, F1: 0, F2: 42}
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 bool = false
            var v19 uint16 = 26
            v20 := GenS1{F0: -32768}
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi16
    }
    if (!false) {
        v21 := GenS3{F0: "a"}
        _ = v21
    }
    switch 4294967295 {
    case 374:
        for fi22 := 0; fi22 < 10; fi22++ {
            v3 += v3
            if (uint32(v2) >= 952) {
                if false {
                    if (1 != (v2 % (((v2 ^ v2)) | 1))) {
                        v14 = (v14 / ((0) | 1))
                    }
                    fmt.Println(v14, fi22, v2)
                    var v23 int = 4294967295
                    _ = v23
                }
                var v24 float64 = 0.1
                _ = v24
            } else {
                v25 := GenS2{F0: 605, F1: -1, F2: 858}
                for fi26 := 0; fi26 < 13; fi26++ {
                    if ((!(!false)) && (0 == 1)) {
                        v27 := GenS3{F0: "\n\t"}
                        var v28 float32 = -2.25
                        v29 := GenS4{F0: "hello", F1: "foo bar"}
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    var v30 int16 = 1000
                    _ = fi26
                    _ = v30
                }
                for fi31 := 0; fi31 < 18; fi31++ {
                    v14 += 0
                    _ = fi31
                }
                _ = v25
            }
            _ = fi22
        }
    }
    v2 = v2
    if len(v1) > 0 {
        v1[0] = 83
    }
    v14 -= (^(v14))
    _ = len(v1)
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 uint = genFn0((false || (true || false)), (false && (4294967295 <= 2147483648)), 42)
        _ = fi32
        _ = v33
    }
    fmt.Println(v14, v2, v14)
    v2 = (v2 ^ (v2 & 127))
    var v34 int64 = -295
    var v35 float64 = -2.25
    v2 += (v2 / (((v2 ^ v2)) | 1))
    switch 4294967295 {
    case -2147483648:
        v2 += (v2 + v2)
    }
    switch 384 {
    case 128:
        v34 = (v34 & v34)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
}
