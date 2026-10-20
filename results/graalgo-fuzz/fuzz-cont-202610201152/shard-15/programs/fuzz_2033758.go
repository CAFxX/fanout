package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint32; F2 int16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
func genFn0() int64 {
    var v5 int64 = 9223372036854775807
    v6 := GenS1{F0: "", F1: false, F2: 177}
    v7 := GenS1{F0: "\n\t", F1: (v5 < v5), F2: 2}
    v8 := GenS1{F0: "foo bar", F1: true, F2: 255}
    var v9 uint8 = 42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return v5
}
func genFn1(p0 uint8, p1 uint16) int64 {
    p0 = 255
    v10 := []int{1, -507, -9223372036854775808, -338}
    p0 += p0
    p0 = (p0 % ((255) | 1))
    _ = p0
    _ = p1
    _ = v10
    return -9223372036854775808
}

func main() {
    var v1 int8 = 100
    var v2 string = "foo bar"
    v3 := []int{}
    v4 := []int{-416, 651, 127, -1}
    if true {
        if (42 != (v1 ^ 22)) {
            for fi11 := 0; fi11 < 6; fi11++ {
                switch ("\n\t" + v2) {
                case "hello":
                    if false {
                        switch (^(fi11)) {
                        case -416:
                            var v12 int = fi11
                            _ = v12
                        }
                    } else {
                        for fi13 := 0; fi13 < 17; fi13++ {
                            var v14 int = (fi11 / ((fi13) | 1))
                            v14 += (-(fi13))
                            var v15 float32 = (1.5 + 0.0)
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        var v16 int = 4294967295
                        if false {
                            v2 = v2
                            for fi17 := 0; fi17 < 8; fi17++ {
                                var v18 int = (-(fi11))
                                v1 += (v1 >> ((2) & 63))
                                var v19 int16 = -678
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        }
                        _ = v16
                    }
                case "\n\t":
                    fmt.Println(v1, v2)
                case "":
                    var v20 uint8 = 42
                    _ = v20
                default:
                    var v21 int32 = 2147483647
                    _ = v21
                }
                if false {
                    if (42 >= 555) {
                        if false {
                            var v22 string = v2
                            var v23 int8 = (^(v1))
                            if true {
                                for fi24 := 0; fi24 < 5; fi24++ {
                                    if (588 != 42) {
                                        v23 = (v23 / ((-128) | 1))
                                        v22 = v22
                                    } else {
                                        v2 = "hello"
                                    }
                                    _ = fi24
                                }
                                v3 = append(v3, fi11)
                                var v25 uint8 = 255
                                _ = v25
                            }
                            _ = v22
                            _ = v23
                        }
                        switch "foo bar" {
                        case "hello":
                            if len(v4) > 0 {
                                v4[0] = (^(fi11))
                            }
                        }
                    } else {
                        v2 = (v2 + v2)
                    }
                }
                _ = fi11
            }
            v2 = v2
            var v26 uint8 = 1
            _ = v26
        }
    }
    if true {
        v27 := GenS1{F0: v2, F1: (4294967295 < 32767), F2: 649}
        v28 := GenS2{F0: 42, F1: 4294967295, F2: 1}
        v29 := GenS3{F0: -1, F1: (v1 / (((v1 << ((42) & 63))) | 1))}
        _ = v27
        _ = v28
        _ = v29
    }
    if true {
        for fi30 := 0; fi30 < 17; fi30++ {
            fmt.Println(v1)
            var v31 int32 = 0
            if true {
                v31 -= (v31 * (v31 - (v31 >> ((42) & 63))))
                var v32 uint8 = 1
                v33 := GenS4{F0: -580}
                _ = v32
                _ = v33
            }
            _ = fi30
            _ = v31
        }
    }
    v34 := GenS2{F0: v1, F1: 0, F2: 1}
    var v35 int8 = 44
    switch ((v2 + (v2 + v2)) + v2) {
    case "foo bar":
        _ = len(v3)
    default:
        var v36 int64 = 1
        _ = v36
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        v38 := GenS5{F0: (true || false)}
        v2 = ("hello" + (("foo bar" + "foo bar") + v2))
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        v1 = 127
        _ = fi39
    }
    if len(v4) > 0 {
        v4[0] = 843
    }
    if false {
        fmt.Println(v1)
    } else {
        v40 := GenS2{F0: v1, F1: 2147483648, F2: 0}
        var v41 uint64 = 514
        _ = v40
        _ = v41
    }
    fmt.Println(v35)
    for fi42 := 0; fi42 < 4; fi42++ {
        v1 += (v1 ^ (-(v1)))
        _ = fi42
    }
    var v43 string = v2
    v1 = v35
    switch ("a" + "hello") {
    case "foo bar":
        var v44 uint32 = 65535
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v43
}
