package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 4294967295
    var v3 uint8 = 255
    var v4 uint32 = 2147483648
    var v5 uint8 = v3
    if true {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 int8 = 127
            v8 := []int{(fi6 >> ((65535) & 63)), fi6}
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if true {
            if false {
                v10 := GenS1{F0: 4294967295}
                v11 := []int{fi9, (^(fi9)), (fi9 & (fi9 % ((fi9) | 1)))}
                _ = v10
                _ = v11
            } else {
                v12 := GenS1{F0: 637}
                _ = v12
            }
            _ = len(v1)
        }
        var v13 uint16 = 32767
        fmt.Println(v13, v5)
        _ = fi9
        _ = v13
    }
    v14 := []int{172, -83, 0, 127}
    _ = len(v14)
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(v5, v2, v5)
        v5 = (v3 + 1)
        v2 = uint64(v4)
        _ = fi15
    }
    if false {
        switch ("\n\t" + "a") {
        case "a":
            v16 := GenS1{F0: 18446744073709551615}
            _ = v16
        case "\n\t":
            v17 := GenS1{F0: 550}
            _ = v17
        }
        v18 := GenS1{F0: 58}
        _ = v18
    } else {
        var v19 bool = false
        if (true && v19) {
            if (255 > 32767) {
                for fi20 := 0; fi20 < 16; fi20++ {
                    for fi21 := 0; fi21 < 6; fi21++ {
                        v22 := GenS2{F0: (v2 * v2), F1: 1, F2: -282}
                        v23 := GenS2{F0: v2, F1: 1, F2: -1}
                        var v24 uint32 = (v4 % (((^(v4))) | 1))
                        var v25 uint = 153
                        _ = len(v14)
                        fmt.Println(v24, v5, fi21)
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = fi20
                }
                v19 = (v19 && true)
                var v26 uint8 = (^(v5))
                _ = v26
            } else {
                v27 := GenS2{F0: (v2 * 380), F1: 98, F2: 1000}
                v4 += 1
                _ = v27
            }
        } else {
            v5 = (v3 + (v3 / ((v5) | 1)))
            v28 := v1["a"]
            _ = v28
        }
        v29 := GenS3{F0: -792, F1: (v2 + 241)}
        _ = v19
        _ = v29
    }
    if false {
        v30 := GenS4{F0: 42, F1: 1}
        _ = v30
    } else {
        if (0 > 467) {
            var v31 int16 = 42
            switch 2147483647 {
            case 4294967295:
                v2 = 18446744073709551615
            }
            _ = v31
        }
        _ = len(v14)
        v32 := GenS1{F0: 645}
        _ = v32
    }
    v3 = v5
    _ = len(v1)
    if (("hello" + "a") == "hello") {
        fmt.Println(v2)
    }
    var v33 float64 = (0.0 - 0.0)
    var v34 uint32 = v4
    v3 = v5
    var v35 uint32 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
