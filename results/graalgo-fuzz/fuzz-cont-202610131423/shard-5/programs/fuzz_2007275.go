package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int64; F2 string }
type GenS4 struct { F0 uint64; F1 uint; F2 int32 }
type GenS5 struct { F0 uint8; F1 int; F2 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 int16; F2 uint8 }
func genFn0(p0 bool, p1 uint64, p2 bool) int64 {
    var v5 uint = 42
    var v6 float64 = 3.14159
    var v7 int8 = -89
    p1 = (p1 & (^(p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 208
    var v2 uint8 = 127
    v3 := make(map[uint32]bool)
    var v4 uint32 = 0
    fmt.Println(v1)
    v8 := v3[v4]
    var v9 uint16 = (v1 + (v1 ^ v1))
    var v10 float64 = 0.0
    v2 = v2
    var v11 uint32 = v4
    for fi12 := 0; fi12 < 6; fi12++ {
        if (1.5 != 0.1) {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 int16 = int16(fi12)
                _ = fi13
                _ = v14
            }
        } else {
            for fi15 := 0; fi15 < 20; fi15++ {
                v1 += (v9 >> ((556) & 63))
                var v16 uint16 = v9
                _ = fi15
                _ = v16
            }
            switch fi12 {
            case 0:
                v2 = v2
            case -403:
                for fi17 := 0; fi17 < 12; fi17++ {
                    v18 := GenS1{F0: "foo bar", F1: -1, F2: 0}
                    if ((!(!false)) && (127 <= 127)) {
                        var v19 int32 = 42
                        v20 := []int{fi17}
                        switch ("a" + ((((("foo bar" + "hello") + "\n\t") + "\n\t") + "hello") + "\n\t")) {
                        case "\n\t":
                            fmt.Println(v19, v4)
                        case "foo bar":
                            if (("foo bar" + ("\n\t" + "\n\t")) != ("a" + "foo bar")) {
                                v11 = v11
                                var v21 int64 = int64(v11)
                                _ = v21
                            } else {
                                v22 := GenS1{F0: "", F1: -117, F2: 1}
                                v9 -= (v9 & v9)
                                switch "\n\t" {
                                case "":
                                    fmt.Println(v1, v9)
                                case "foo bar":
                                    v23 := GenS2{F0: 1}
                                    _ = v23
                                case "hello":
                                    v9 = (v1 % ((v1) | 1))
                                default:
                                    if v8 {
                                        v24 := GenS1{F0: "", F1: 0, F2: 42}
                                        _ = v24
                                    } else {
                                        v3[v11] = (v19 > v19)
                                    }
                                }
                                fmt.Println(fi17, fi12)
                                fmt.Println(v8)
                                _ = v22
                            }
                        default:
                            v1 = 1
                        }
                        _ = v19
                        _ = v20
                    } else {
                        switch fi12 {
                        case 42:
                            v25 := GenS1{F0: "foo bar", F1: 1, F2: -520}
                            _ = v25
                        case 2147483647:
                            fmt.Println(v8)
                        }
                    }
                    fmt.Println(v1, v8)
                    _ = fi17
                    _ = v18
                }
            }
            v26 := GenS3{F0: v2, F1: 715, F2: "foo bar"}
            _ = v26
        }
        if ((^(v9)) != (v9 >> ((65535) & 63))) {
            fmt.Println(v11, v10, v1)
            fmt.Println(fi12)
            v27 := GenS2{F0: (^(v11))}
            _ = v27
        }
        _ = fi12
    }
    v11 += v11
    delete(v3, 2)
    v28 := GenS4{F0: 0, F1: 255, F2: 1}
    v2 += 184
    for fi29 := 0; fi29 < 4; fi29++ {
        v2 = v2
        _ = fi29
    }
    var v30 uint = 673
    var v31 int32 = 42
    for fi32 := 0; fi32 < 2; fi32++ {
        v33 := GenS2{F0: v11}
        var v34 uint8 = (v2 * v2)
        if v8 {
            if (128 == v2) {
                var v35 uint16 = 32767
                _ = v35
            }
            v30 = v30
        } else {
            if v8 {
                var v36 bool = true
                v37 := GenS3{F0: (v34 * v2), F1: -9223372036854775808, F2: ("" + "\n\t")}
                _ = v36
                _ = v37
            } else {
                v38 := []int{(fi32 >> ((4294967295) & 63)), 456, 1, (^(fi32))}
                switch 9223372036854775807 {
                case -437:
                    v39 := GenS5{F0: (^(v34)), F1: (fi32 ^ (fi32 - fi32)), F2: 308}
                    _ = v39
                case 42:
                    switch (fi32 / ((fi32) | 1)) {
                    case 32767:
                        if v8 {
                            v40 := GenS6{F0: 4294967295}
                            v41 := v3[(v11 + 2147483647)]
                            var v42 int8 = -128
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                    }
                case 9223372036854775807:
                    for fi43 := 0; fi43 < 6; fi43++ {
                        v44 := GenS7{F0: "hello"}
                        var v45 int16 = 20
                        var v46 uint64 = 4294967295
                        _ = fi43
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                }
                _ = v38
            }
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    switch 128 {
    case 128:
        var v47 int16 = 0
        _ = v47
    case -1:
        fmt.Println(v1)
    case 158:
        var v48 int = 4294967295
        _ = v48
    }
    fmt.Println(v2)
    for fi49 := 0; fi49 < 13; fi49++ {
        v50 := []int{}
        v51 := GenS8{F0: v9, F1: 1, F2: 42}
        var v52 int32 = v31
        _ = fi49
        _ = v50
        _ = v51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v8
    _ = v9
}
