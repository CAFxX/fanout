package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64; F1 uint8; F2 bool }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint64; F1 int; F2 bool }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint32) int32 {
    switch -9223372036854775808 {
    case 31:
        v7 := GenS1{F0: 42, F1: 2147483647}
        _ = v7
    case 0:
        var v8 int32 = -2147483648
        _ = v8
    }
    var v9 int64 = p0
    if true {
        fmt.Println(v9)
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
        var v12 int8 = 1
        _ = v12
    } else {
        for fi13 := 0; fi13 < 7; fi13++ {
            if true {
                v14 := GenS2{F0: v9, F1: 128, F2: true}
                _ = v14
            } else {
                v15 := GenS3{F0: 1, F1: 0}
                _ = v15
            }
            v16 := GenS4{F0: 42, F1: fi13, F2: true}
            var v17 uint32 = 42
            _ = fi13
            _ = v16
            _ = v17
        }
        var v18 float32 = (-2.25 * ((3.14159 / 1.5) * 0.0))
        _ = v18
    }
    var v19 uint32 = p1
    _ = p0
    _ = p1
    _ = v19
    _ = v9
    return -2147483648
}
func genFn1(p0 int64) uint32 {
    var v20 uint64 = 1
    var v21 int = 832
    v22 := []int{v21}
    v23 := GenS3{F0: 525, F1: 0}
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    return 2147483648
}

func main() {
    var v1 int16 = 1000
    v2 := []uint32{4294967295, 338}
    var v3 string = ("\n\t" + "foo bar")
    v4 := make(map[int]string)
    var v5 uint = 18446744073709551615
    var v6 string = v3
    if ((^(v5)) <= (v5 & 42)) {
        if ((127 != 42) || (1e10 > 1.5)) {
            if true {
                _ = len(v2)
                v5 -= 42
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 int64 = -9223372036854775808
                    if false {
                        v26 := v4[fi24]
                        v27 := GenS5{F0: 432}
                        _ = v26
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
            }
            var v28 bool = false
            fmt.Println(v28, v6, v6)
            _ = v28
        } else {
            v5 = v5
            for fi29 := 0; fi29 < 12; fi29++ {
                if (true && false) {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        var v31 string = (v3 + "foo bar")
                        v2 = append(v2, 2147483647)
                        _ = fi30
                        _ = v31
                    }
                    _ = len(v2)
                    var v32 uint = (v5 % (((v5 % (((v5 | v5)) | 1))) | 1))
                    _ = v32
                }
                _ = fi29
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v3 = ""
        v3 = v3
    } else {
        var v33 bool = (!(-2.25 < 0.0))
        _ = v33
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        fmt.Println(v3, v3)
        _ = fi34
    }
    v5 = v5
    v35 := GenS4{F0: 42, F1: 127, F2: (373 != 18446744073709551615)}
    var v36 int64 = 0
    fmt.Println(v6)
    if (0 < 255) {
        var v37 uint8 = 0
        _ = v37
    }
    if true {
        v38 := GenS4{F0: 890, F1: 4294967295, F2: false}
        _ = v38
    } else {
        fmt.Println(v1, v36)
        v3 = ((v3 + "foo bar") + (((v6 + ("hello" + v6)) + (v3 + "\n\t")) + (v3 + v6)))
        var v39 float64 = (1e10 + 1.5)
        _ = v39
    }
    var v40 int16 = v1
    v41 := GenS4{F0: 1, F1: 4294967295, F2: true}
    switch v6 {
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "hello":
        v6 = v6
    }
    var v42 uint64 = 746
    v6 = ("foo bar" + v3)
    var v43 int16 = -734
    var v44 int64 = (^(v36))
    if (!false) {
        v45 := v4[-201]
        v46 := GenS1{F0: 255, F1: -9223372036854775808}
        var v47 uint32 = 2147483647
        _ = v45
        _ = v46
        _ = v47
    } else {
        fmt.Println(v40, v43, v1)
        var v48 int = 42
        switch v48 {
        case 1:
            var v49 uint = (^(v5))
            _ = v49
        case -400:
            if len(v2) > 0 {
                v2[0] = 278
            }
        }
        _ = v48
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
