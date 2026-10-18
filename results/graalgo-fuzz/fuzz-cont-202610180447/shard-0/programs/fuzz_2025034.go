package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int }
type GenS2 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint; F1 uint64; F2 int32 }
type GenS5 struct { F0 int64; F1 int64; F2 uint8 }
type GenS6 struct { F0 bool; F1 uint64; F2 int32 }
func genFn0() int64 {
    var v7 uint16 = 851
    var v8 string = "foo bar"
    switch (("hello" + "") + "") {
    case "\n\t":
        v8 = "\n\t"
    default:
        var v9 int8 = -128
        _ = v9
    }
    fmt.Println(v8)
    if ("\n\t" < "a") {
        fmt.Println(v7, v8)
        for fi10 := 0; fi10 < 6; fi10++ {
            v7 += (v7 | 188)
            for fi11 := 0; fi11 < 5; fi11++ {
                if false {
                    v12 := []int{fi11, (fi10 << ((255) & 63)), fi10}
                    _ = v12
                }
                var v13 bool = false
                _ = fi11
                _ = v13
            }
            v8 = v8
            _ = fi10
        }
    }
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 bool = (4294967295 == 1)
    var v2 uint32 = 42
    var v3 int16 = -211
    var v4 uint16 = 567
    v5 := []uint32{}
    var v6 float32 = (float32((1.5 * float64(0.1))) * 1.5)
    v14 := GenS1{F0: genFn0(), F1: -128, F2: -554}
    _ = len(v5)
    if v1 {
        switch "" {
        case "hello":
            switch "\n\t" {
            case "foo bar":
                for fi15 := 0; fi15 < 6; fi15++ {
                    v6 = v6
                    _ = fi15
                }
            case "a":
                v6 += v6
            case "":
                if false {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        _ = len(v5)
                        _ = fi16
                    }
                    switch 2147483647 {
                    case 42:
                        for fi17 := 0; fi17 < 15; fi17++ {
                            switch fi17 {
                            case -9223372036854775808:
                                var v18 string = ""
                                _ = v18
                            }
                            _ = fi17
                        }
                    }
                } else {
                    if v1 {
                        _ = len(v5)
                    }
                }
            }
        }
    } else {
        for fi19 := 0; fi19 < 1; fi19++ {
            v1 = v1
            v6 = (v6 + float32(float64(v6)))
            _ = fi19
        }
    }
    var v20 string = ("hello" + ("" + "a"))
    v1 = (v1 && false)
    if false {
        if ((false || (v4 != (v4 - 1))) && v1) {
            fmt.Println(v4, v20, v6)
            var v21 string = (("a" + v20) + ((v20 + v20) + "a"))
            _ = v21
        } else {
            v6 += 0.0
        }
        v5 = append(v5, v2)
        v5 = append(v5, v2)
    } else {
        v3 = 291
    }
    if (-1 >= 605) {
        fmt.Println(v2, v1)
        if (-503 <= -1) {
            for fi22 := 0; fi22 < 3; fi22++ {
                v5 = append(v5, (v2 | 0))
                v23 := GenS2{F0: v3, F1: v4, F2: 279}
                var v24 uint16 = (v4 - v4)
                _ = fi22
                _ = v23
                _ = v24
            }
            v4 = v4
        }
        var v25 uint64 = 1
        _ = v25
    } else {
        fmt.Println(v2)
    }
    v26 := GenS1{F0: -1, F1: -1, F2: 1}
    for fi27 := 0; fi27 < 2; fi27++ {
        var v28 string = (v20 + "\n\t")
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v30 uint = 18446744073709551615
        _ = fi29
        _ = v30
    }
    fmt.Println(v6, v20)
    v31 := GenS3{F0: (v2 + (v2 << ((192) & 63))), F1: v4, F2: (v3 * 460)}
    if len(v5) > 0 {
        v5[0] = 481
    }
    v32 := GenS4{F0: 1, F1: 1, F2: -82}
    v5 = append(v5, (^(v2)))
    switch -1 {
    case -742:
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := GenS5{F0: -1, F1: 42, F2: 2}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 int32 = int32(v3)
                if v1 {
                    var v37 uint = 18446744073709551615
                    var v38 int = fi33
                    var v39 uint32 = v2
                    _ = v37
                    _ = v38
                    _ = v39
                }
                _ = fi35
                _ = v36
            }
            _ = fi33
            _ = v34
        }
    default:
        v2 = v2
    }
    _ = len(v5)
    v40 := GenS6{F0: (!v1), F1: 439, F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
