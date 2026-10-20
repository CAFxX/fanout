package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint64 }

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]string)
    var v3 uint64 = 4294967295
    v4 := make(map[int]int)
    v5 := GenS1{F0: "hello"}
    var v6 int8 = 100
    var v7 int16 = -32768
    v3 = v3
    for fi8 := 0; fi8 < 10; fi8++ {
        v7 = (v7 ^ v7)
        _ = fi8
    }
    v9 := GenS1{F0: "\n\t"}
    if true {
        v1 = (^(v1))
    } else {
        if false {
            v6 += v6
            _ = len(v4)
            var v10 string = "a"
            _ = v10
        } else {
            v11 := []int{-298}
            if false {
                if (!(2147483648 != 278)) {
                    v6 -= v6
                    if false {
                        var v12 float32 = (1e10 - 1e10)
                        fmt.Println(v1, v7)
                        v7 = v7
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 8; fi13++ {
                            var v14 uint8 = 2
                            _ = fi13
                            _ = v14
                        }
                        var v15 uint = 75
                        var v16 uint = v15
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(v1, v1, v7)
                }
                v3 = (v3 << ((0) & 63))
            } else {
                var v17 int8 = 42
                switch "a" {
                case "foo bar":
                    var v18 uint16 = uint16(v3)
                    _ = v18
                default:
                    fmt.Println(v7)
                }
                if (!true) {
                    v19 := GenS2{F0: v3}
                    switch (("foo bar" + "") + ("hello" + "")) {
                    case "a":
                        var v20 bool = (42 < uint64(v1))
                        _ = v20
                    case "":
                        for fi21 := 0; fi21 < 2; fi21++ {
                            v7 += 175
                            v22 := GenS3{F0: (v7 << ((556) & 63))}
                            _ = fi21
                            _ = v22
                        }
                    default:
                        for fi23 := 0; fi23 < 10; fi23++ {
                            var v24 uint64 = (v3 * v3)
                            var v25 uint16 = (v1 << ((65535) & 63))
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                    var v26 int = int(v3)
                    _ = v19
                    _ = v26
                } else {
                    v7 += v7
                    v27 := GenS4{F0: 971}
                    for fi28 := 0; fi28 < 7; fi28++ {
                        var v29 float64 = 0.1
                        for fi30 := 0; fi30 < 1; fi30++ {
                            _ = len(v4)
                            fmt.Println(v6)
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                v17 = (v6 ^ int8(fi28))
                                var v31 float32 = (1.5 * float32(v29))
                                _ = v31
                            case "":
                                if true {
                                    v32 := v4[(^(fi28))]
                                    _ = v32
                                } else {
                                    v33 := GenS1{F0: ("foo bar" + ("a" + ("foo bar" + "\n\t")))}
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    v34 := GenS2{F0: v3}
                                    _ = v33
                                    _ = v34
                                }
                            default:
                                var v35 float64 = v29
                                _ = v35
                            }
                            _ = fi30
                        }
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                }
                _ = v17
            }
            _ = v11
        }
        v7 = (^(v7))
    }
    if (4294967295 == 2147483648) {
        v3 = (^(v3))
        v6 -= (v6 + (v6 - int8(v7)))
        var v36 string = ""
        _ = v36
    } else {
        fmt.Println(v3, v1, v3)
        v1 = v1
        v37 := GenS4{F0: 1}
        _ = v37
    }
    if (-2147483648 == 2147483647) {
        v6 -= v6
        v38 := GenS2{F0: 18446744073709551615}
        _ = v38
    }
    var v39 uint8 = 127
    v6 = (v6 & -1)
    v40 := GenS5{F0: 42, F1: uint64(v39)}
    var v41 uint = 65535
    var v42 uint = (^(v41))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
