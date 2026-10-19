package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int; F1 int8; F2 int64 }
func genFn0(p0 int16, p1 uint16, p2 int64) uint16 {
    if (p2 <= (p2 / ((p2) | 1))) {
        var v4 uint = 0
        v5 := GenS1{F0: v4, F1: 82}
        _ = v4
        _ = v5
    }
    switch "" {
    case "":
        p2 += p2
    case "a":
        v6 := GenS1{F0: 623, F1: 127}
        _ = v6
    default:
        if ((p1 - p1) > p1) {
            for fi7 := 0; fi7 < 7; fi7++ {
                if true {
                    p0 += p0
                } else {
                    p1 -= (^(p1))
                    for fi8 := 0; fi8 < 3; fi8++ {
                        v9 := GenS2{F0: (fi8 << ((4294967295) & 63)), F1: false}
                        p0 += p0
                        _ = fi8
                        _ = v9
                    }
                }
                var v10 int16 = -52
                _ = fi7
                _ = v10
            }
        } else {
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(p2, p2, p0)
                v12 := []int{32767, (fi11 % ((-2147483648) | 1)), (fi11 & 0)}
                fmt.Println(fi11, fi11, p2)
                _ = fi11
                _ = v12
            }
        }
    }
    p0 -= (p0 & p0)
    v13 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 0
}

func main() {
    var v1 uint32 = 1
    var v2 int64 = -59
    v3 := make(map[uint32]string)
    var v14 bool = true
    fmt.Println(v14, v1, v1)
    if v14 {
        for fi15 := 0; fi15 < 5; fi15++ {
            v14 = v14
            _ = fi15
        }
        if false {
            switch 4294967295 {
            case 127:
                var v16 int64 = v2
                _ = v16
            }
            for fi17 := 0; fi17 < 9; fi17++ {
                delete(v3, (^(v1)))
                for fi18 := 0; fi18 < 15; fi18++ {
                    v19 := GenS1{F0: 558, F1: 0}
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        }
    } else {
        v14 = (65535 == 616)
    }
    _ = len(v3)
    fmt.Println(v1)
    fmt.Println(v14)
    var v20 uint32 = (v1 / (((v1 + v1)) | 1))
    var v21 uint = 280
    for fi22 := 0; fi22 < 9; fi22++ {
        v1 -= (v20 | v20)
        var v23 int8 = 66
        var v24 int16 = -32768
        _ = fi22
        _ = v23
        _ = v24
    }
    v21 = 65535
    v25 := []int{32767, 1}
    for fi26 := 0; fi26 < 5; fi26++ {
        if (467 == 0) {
            for fi27 := 0; fi27 < 12; fi27++ {
                v14 = v14
                var v28 int16 = 32767
                _ = fi27
                _ = v28
            }
        }
        _ = len(v3)
        _ = fi26
    }
    var v29 string = "a"
    delete(v3, v1)
    v30 := GenS1{F0: v21, F1: 255}
    v31 := GenS3{F0: 128, F1: -105, F2: 42}
    switch v29 {
    case "":
        _ = len(v25)
    case "a":
        if (!v14) {
            v32 := GenS1{F0: 2, F1: 128}
            var v33 int32 = -1
            _ = v32
            _ = v33
        }
    case "hello":
        v29 = "\n\t"
    }
    var v34 int = 332
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
}
