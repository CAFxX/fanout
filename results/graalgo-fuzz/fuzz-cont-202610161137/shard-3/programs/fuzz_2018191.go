package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint64; F1 int16; F2 int32 }
type GenS7 struct { F0 int }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := GenS1{F0: 18446744073709551615, F1: -754}
        v1 += v1
        _ = fi4
        _ = v5
    }
    v3[0] = 42
    v6 := GenS2{F0: (v1 % ((v1) | 1))}
    v1 -= v1
    v1 = v1
    var v7 int8 = 126
    v1 -= (v1 * (v1 << ((2) & 63)))
    delete(v2, (("\n\t" + ("\n\t" + "a")) + ""))
    v1 = v1
    for fi8 := 0; fi8 < 17; fi8++ {
        if (("hello" + "") != "\n\t") {
            _ = len(v3)
        } else {
            v7 = 92
            v1 += 65535
        }
        _ = fi8
    }
    if (-996 > -2147483648) {
        var v9 uint64 = 1
        if false {
            if (!(!true)) {
                v9 -= v9
                v1 = v1
                var v10 uint8 = 42
                _ = v10
            } else {
                for fi11 := 0; fi11 < 20; fi11++ {
                    v7 += v7
                    for fi12 := 0; fi12 < 8; fi12++ {
                        var v13 uint = v1
                        if ((((0 >= 42) || (true || false)) || false) && (!true)) {
                            if false {
                                if false {
                                    _ = len(v3)
                                } else {
                                    var v14 float32 = 1.5
                                    var v15 int16 = 770
                                    _ = v14
                                    _ = v15
                                }
                                if (1.5 > 0.0) {
                                    fmt.Println(v9)
                                    v1 += v1
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        switch "foo bar" {
                                        case "":
                                            v16 := GenS1{F0: 52, F1: -1}
                                            _ = v16
                                        case "a":
                                            fmt.Println(v1)
                                            v17 := []int{127}
                                            _ = v17
                                        }
                                    case "foo bar":
                                        if false {
                                            v3[2147483647] = 0
                                        } else {
                                            v18 := GenS1{F0: (^(v9)), F1: 42}
                                            _ = v18
                                        }
                                    }
                                } else {
                                    var v19 uint32 = 4294967295
                                    _ = v19
                                }
                            }
                            if (true && true) {
                                if ((!((v7 / (((v7 & (v7 / ((v7) | 1)))) | 1)) < int8(fi12))) || false) {
                                    var v20 uint16 = 385
                                    var v21 bool = (-1 != -91)
                                    _ = v20
                                    _ = v21
                                }
                            } else {
                                var v22 string = "foo bar"
                                v7 = (v7 & (v7 % ((v7) | 1)))
                                _ = v22
                            }
                            v23 := GenS2{F0: v13}
                            v24 := GenS3{F0: (("a" + (("hello" + "\n\t") + "foo bar")) + "a")}
                            v1 = 18446744073709551615
                            delete(v3, 65535)
                            _ = v23
                            _ = v24
                        } else {
                            v25 := GenS4{F0: 0, F1: (v9 + 874)}
                            var v26 uint16 = 42
                            v13 -= v1
                            v27 := GenS4{F0: 107, F1: 2}
                            if true {
                                v28 := GenS5{F0: fi12, F1: v13}
                                switch fi12 {
                                case 0:
                                    v13 -= v13
                                    v29 := GenS4{F0: 182, F1: (^(v9))}
                                    _ = v29
                                case 42:
                                    switch "hello" {
                                    case "a":
                                        v30 := GenS1{F0: v9, F1: 819}
                                        _ = v30
                                    case "\n\t":
                                        v31 := []int{-2147483648, (fi11 + -1), -808, fi11}
                                        _ = v31
                                    }
                                }
                                if true {
                                    var v32 uint = v1
                                    v33 := GenS4{F0: 1, F1: v9}
                                    v34 := GenS4{F0: uint8(v26), F1: v9}
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v35 string = ""
                                    fmt.Println(v7)
                                    _ = v35
                                }
                                _ = v28
                            }
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
            }
            fmt.Println(v1, v9)
        }
        var v36 int8 = v7
        _ = v36
        _ = v9
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        var v38 string = "foo bar"
        v39 := GenS6{F0: 1, F1: -1, F2: -2147483648}
        for fi40 := 0; fi40 < 18; fi40++ {
            delete(v2, v38)
            _ = fi40
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    for fi41 := 0; fi41 < 5; fi41++ {
        switch "hello" {
        case "\n\t":
            v42 := GenS7{F0: 270}
            _ = v42
        }
        _ = fi41
    }
    var v43 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v6
    _ = v7
}
