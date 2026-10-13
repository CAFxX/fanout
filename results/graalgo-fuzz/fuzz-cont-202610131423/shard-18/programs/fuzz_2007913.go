package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint; F1 uint8; F2 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 int8; F2 int8 }
type GenS7 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS8 struct { F0 int8; F1 int8; F2 int8 }

func main() {
    var v1 uint = 412
    var v2 int16 = -704
    var v3 float64 = (1.5 - 0.1)
    v4 := make(map[uint32]bool)
    var v5 int16 = 42
    v6 := make(map[uint32]bool)
    delete(v6, 2147483647)
    switch int(v5) {
    case 42:
        for fi7 := 0; fi7 < 1; fi7++ {
            if true {
                var v8 uint16 = 0
                var v9 uint8 = 127
                v10 := GenS1{F0: -46, F1: 0}
                _ = v10
                _ = v8
                _ = v9
            } else {
                v3 = (v3 - (0.1 * 0.0))
                fmt.Println(v5, v2)
                if (2 <= 42) {
                    if true {
                        _ = len(v4)
                    } else {
                        delete(v6, 42)
                    }
                    v5 -= 0
                    v1 = v1
                } else {
                    v11 := GenS2{F0: 0, F1: 1, F2: (v5 ^ v5)}
                    var v12 int = (fi7 << ((255) & 63))
                    _ = v11
                    _ = v12
                }
            }
            v13 := GenS3{F0: v1, F1: 42, F2: "foo bar"}
            for fi14 := 0; fi14 < 19; fi14++ {
                fmt.Println(fi14)
                if false {
                    v15 := GenS1{F0: 42, F1: 1}
                    switch "a" {
                    case "\n\t":
                        if (927 < 0) {
                            v2 = 330
                            v1 = (v1 * 984)
                            v16 := GenS2{F0: 255, F1: 0, F2: (v5 % (((v2 % ((-32768) | 1))) | 1))}
                            v5 = 0
                            _ = v16
                        } else {
                            var v17 uint32 = 65535
                            _ = v17
                        }
                    case "a":
                        if (!((false && ((!true) && false)) && (!true))) {
                            fmt.Println(fi7, v5, v5)
                        }
                    case "":
                        if false {
                            v18 := GenS4{F0: v1}
                            v3 -= (-2.25 / v3)
                            if ((v1 % ((v1) | 1)) >= 241) {
                                if false {
                                    v19 := GenS3{F0: v1, F1: 128, F2: ("\n\t" + "a")}
                                    _ = v19
                                }
                                delete(v4, 4294967295)
                            } else {
                                var v20 uint8 = 42
                                v4[0] = ((66 <= 255) && false)
                                _ = v20
                            }
                            _ = v18
                        }
                    }
                    if false {
                        var v21 int8 = 1
                        var v22 string = "\n\t"
                        _ = v21
                        _ = v22
                    }
                    _ = v15
                } else {
                    v23 := v4[42]
                    var v24 int32 = 1
                    _ = v23
                    _ = v24
                }
                v25 := GenS2{F0: 255, F1: 2, F2: (v5 << ((v1) & 63))}
                var v26 int32 = -2147483648
                _ = fi14
                _ = v25
                _ = v26
            }
            _ = fi7
            _ = v13
        }
    case -299:
        v27 := GenS5{F0: 42}
        _ = v27
    default:
        v28 := GenS5{F0: 2}
        _ = v28
    }
    v2 = v2
    fmt.Println(v3)
    if (2147483647 >= 42) {
        fmt.Println(v1)
        v1 += (^(v1))
    } else {
        v29 := GenS6{F0: 42, F1: -82, F2: 0}
        _ = v29
    }
    v30 := GenS3{F0: v1, F1: 42, F2: ("\n\t" + ("\n\t" + "a"))}
    v31 := GenS7{F0: 2, F1: -478, F2: 1}
    if true {
        for fi32 := 0; fi32 < 12; fi32++ {
            var v33 int32 = 0
            if (true || (!false)) {
                v1 += (v1 >> (((v1 << ((v1) & 63))) & 63))
                var v34 uint16 = 2
                var v35 string = "a"
                _ = v34
                _ = v35
            } else {
                if (false || true) {
                    v1 += (v1 / ((65535) | 1))
                    v36 := v6[4294967295]
                    _ = v36
                } else {
                    v37 := GenS8{F0: 34, F1: 127, F2: 1}
                    v38 := []int{-468, 726}
                    _ = v37
                    _ = v38
                }
            }
            _ = fi32
            _ = v33
        }
        fmt.Println(v2)
        v39 := GenS1{F0: -1, F1: 634}
        _ = v39
    } else {
        v5 -= v2
        v4[2147483647] = ("" <= "")
    }
    v40 := []int{32767}
    fmt.Println(v2, v2, v1)
    v41 := GenS4{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
