package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 bool; F1 uint16; F2 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 2
    var v3 int8 = -128
    var v4 string = "hello"
    var v5 float64 = 1e10
    v6 := GenS1{F0: 42, F1: 42}
    var v7 int32 = 428
    var v8 int16 = -720
    switch (v4 + "hello") {
    case "foo bar":
        v9 := GenS2{F0: 1}
        _ = v9
    case "a":
        var v10 float32 = 0.0
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS3{F0: 4294967295, F1: ("" + v4)}
        _ = fi11
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint8 = (v2 + v2)
            if (42 <= 42) {
                v15 := GenS4{F0: 241, F1: 361}
                var v16 uint64 = 42
                _ = v15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        switch (("\n\t" + ((v4 + "a") + v4)) + (v4 + "a")) {
        case "foo bar":
            delete(v1, 2147483648)
        case "":
            v17 := v1[65535]
            _ = v17
        }
    } else {
        if true {
            if ((-2.25 + 1e10) < (-2.25 - 0.0)) {
                var v18 int64 = 0
                if (!true) {
                    v19 := []int{2147483647}
                    _ = v19
                } else {
                    fmt.Println(v4)
                    v20 := GenS1{F0: 255, F1: 65535}
                    _ = v20
                }
                if ((-(v3)) != v3) {
                    v2 += (v2 | 127)
                } else {
                    switch -732 {
                    case 798:
                        v8 += (v8 * v8)
                    default:
                        fmt.Println(v2, v3)
                    }
                    v21 := GenS5{F0: (!true), F1: 0, F2: (v7 & v7)}
                    var v22 uint64 = 4294967295
                    _ = v21
                    _ = v22
                }
                _ = v18
            }
            v23 := []int{1}
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := GenS3{F0: 18446744073709551615, F1: ("a" + v4)}
                for fi26 := 0; fi26 < 20; fi26++ {
                    var v27 bool = (2 > 1)
                    v8 += 32767
                    _ = fi26
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        if (!false) {
            for fi28 := 0; fi28 < 1; fi28++ {
                v7 += (v7 * (v7 >> ((0) & 63)))
                for fi29 := 0; fi29 < 6; fi29++ {
                    var v30 uint64 = 525
                    if true {
                        switch fi29 {
                        case 4294967295:
                            v31 := GenS2{F0: v30}
                            _ = v31
                        default:
                            if ((!(0.1 == 1.5)) || ((v4 + v4) < v4)) {
                                var v32 int8 = (v3 | 2)
                                v33 := []int{32767}
                                _ = v32
                                _ = v33
                            } else {
                                v30 = v30
                                var v34 uint64 = 1
                                var v35 float64 = v5
                                _ = v34
                                _ = v35
                            }
                            if (!(!true)) {
                                v2 = v2
                            }
                        }
                        v36 := []int{32768}
                        _ = v36
                    }
                    fmt.Println(v8)
                    _ = fi29
                    _ = v30
                }
                for fi37 := 0; fi37 < 20; fi37++ {
                    v38 := GenS2{F0: 562}
                    _ = fi37
                    _ = v38
                }
                _ = fi28
            }
        } else {
            var v39 uint32 = 4294967295
            _ = v39
        }
    }
    v7 = 0
    var v40 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
