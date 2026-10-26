package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 float32 = 0.0
    var v3 int32 = -2147483648
    var v4 uint16 = 2
    var v5 string = "\n\t"
    v4 = 65535
    v6 := GenS1{F0: ("\n\t" + (v5 + (v5 + "hello"))), F1: 733}
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            v8 := []int{fi7}
            v8 = append(v8, (^(fi7)))
            _ = len(v8)
            _ = v8
        } else {
            if false {
                for fi9 := 0; fi9 < 5; fi9++ {
                    v10 := []int{}
                    var v11 int = -620
                    fmt.Println(v1, v4)
                    if (((v1 ^ v1) == v1) || (uint32(v1) <= 67)) {
                        var v12 uint = v1
                        switch 42 {
                        case 1:
                            var v13 int32 = v3
                            _ = v13
                        case -449:
                            var v14 uint = 255
                            _ = v14
                        }
                        v15 := GenS1{F0: v5, F1: -178}
                        _ = v12
                        _ = v15
                    } else {
                        if len(v10) > 0 {
                            v10[0] = 42
                        }
                        var v16 int32 = v3
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        _ = v16
                    }
                    var v17 int = v11
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v17
                }
            }
        }
        v18 := GenS1{F0: "foo bar", F1: -725}
        var v19 int8 = int8(v1)
        _ = fi7
        _ = v18
        _ = v19
    }
    var v20 int64 = 0
    var v21 uint = 633
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int64 = v20
        v23 -= v20
        v24 := []int{fi22, fi22}
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int = 42
    for fi26 := 0; fi26 < 6; fi26++ {
        if true {
            var v27 float64 = (0.0 - (0.1 - (1e10 / 3.14159)))
            if (1 != 42) {
                var v28 float64 = -2.25
                _ = v28
            }
            v1 -= v1
            _ = v27
        } else {
            fmt.Println(fi26)
            if false {
                v25 -= fi26
            } else {
                for fi29 := 0; fi29 < 13; fi29++ {
                    v30 := GenS2{F0: 951}
                    v31 := GenS1{F0: (v5 + ("hello" + v5)), F1: (v20 >> ((v21) & 63))}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                if (9223372036854775807 < fi26) {
                    v2 += float32(((0.0 - (0.0 - 0.1)) * 0.0))
                }
                v32 := []int{}
                _ = v32
            }
            v33 := []int{v25, (fi26 % (((v25 >> ((v1) & 63))) | 1))}
            _ = v33
        }
        for fi34 := 0; fi34 < 1; fi34++ {
            v35 := []int{}
            var v36 uint8 = 255
            v37 := GenS1{F0: ("foo bar" + v5), F1: v20}
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        for fi38 := 0; fi38 < 17; fi38++ {
            v39 := GenS3{F0: 255, F1: 979}
            v5 = v5
            _ = fi38
            _ = v39
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
