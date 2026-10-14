package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 bool) int64 {
    var v7 uint16 = 2
    var v8 uint64 = 63
    if ((v7 % ((v7) | 1)) > (^(v7))) {
        v9 := []int{4294967295, 4294967295, 32767, -501}
        p0 = p0
        v8 = 0
        _ = v9
    }
    fmt.Println(p0, p0, v7)
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int8 = -128
    var v2 uint64 = 485
    var v3 uint16 = 32767
    var v4 int64 = 42
    v5 := []string{"hello"}
    var v6 uint64 = 2
    v1 = (v1 * 0)
    var v10 int32 = 872
    switch -2147483648 {
    case 2147483647:
        for fi11 := 0; fi11 < 7; fi11++ {
            if false {
                v12 := GenS1{F0: 0, F1: 1000}
                _ = v12
            } else {
                var v13 int64 = int64(fi11)
                v14 := GenS1{F0: 0, F1: 1}
                v2 -= v6
                _ = v13
                _ = v14
            }
            if false {
                v4 += 9223372036854775807
            } else {
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 float32 = 0.1
                    v17 := GenS1{F0: -85, F1: 1}
                    v5 = append(v5, "\n\t")
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v18 float32 = 1e10
                var v19 int = fi11
                _ = v18
                _ = v19
            }
            _ = fi11
        }
    case 688:
        v20 := GenS2{F0: (v1 & v1), F1: v6}
        _ = v20
    default:
        for fi21 := 0; fi21 < 14; fi21++ {
            for fi22 := 0; fi22 < 3; fi22++ {
                v23 := GenS1{F0: -1, F1: 32767}
                _ = fi22
                _ = v23
            }
            fmt.Println(v2, fi21, v10)
            fmt.Println(v10)
            _ = fi21
        }
    }
    if false {
        if false {
            var v24 uint8 = 164
            _ = v24
        }
    }
    v2 = v6
    if false {
        v25 := GenS2{F0: (v1 + (v1 + 127)), F1: 4294967295}
        for fi26 := 0; fi26 < 5; fi26++ {
            if len(v5) > 0 {
                v5[0] = "hello"
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                v3 = (v3 | (v3 - 958))
                _ = fi27
            }
            _ = len(v5)
            _ = fi26
        }
        _ = v25
    } else {
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 int64 = v4
            if true {
                if true {
                    for fi30 := 0; fi30 < 18; fi30++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(v4, fi30)
                        _ = fi30
                    }
                } else {
                    v6 -= v6
                }
                fmt.Println(v10, v2)
                v31 := GenS3{F0: (v1 % (((v1 - (v1 % (((^(v1))) | 1)))) | 1)), F1: v10}
                _ = v31
            }
            v32 := GenS1{F0: -647, F1: -32768}
            _ = fi28
            _ = v29
            _ = v32
        }
    }
    var v33 uint = 0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v34 int32 = -1
    if (42 >= v34) {
        var v35 int32 = (^(v10))
        _ = v35
    }
    v36 := GenS1{F0: -525, F1: 42}
    v37 := GenS4{F0: uint8(v3)}
    var v38 int64 = (v4 - (v4 - int64(v10)))
    for fi39 := 0; fi39 < 17; fi39++ {
        v40 := GenS2{F0: v1, F1: v6}
        _ = fi39
        _ = v40
    }
    if len(v5) > 0 {
        v5[0] = ("" + "hello")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
