package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint16; F2 string }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }

func main() {
    var v1 int16 = 1
    var v2 uint64 = 0
    var v3 bool = true
    var v4 uint16 = 228
    if v3 {
        v5 := GenS1{F0: 346, F1: 4294967295, F2: "hello"}
        switch 128 {
        case 32767:
            for fi6 := 0; fi6 < 18; fi6++ {
                var v7 uint32 = 2147483648
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 string = (("\n\t" + "") + (("foo bar" + "hello") + ("" + ("hello" + "\n\t"))))
                    fmt.Println(v1, fi6)
                    if true {
                        fmt.Println(v7, v2)
                        v10 := GenS1{F0: 18446744073709551615, F1: 0, F2: ("a" + "\n\t")}
                        v11 := GenS2{F0: 255}
                        if true {
                            switch ("" + ("hello" + v9)) {
                            case "\n\t":
                                if v3 {
                                    var v12 float64 = -2.25
                                    var v13 bool = v3
                                    _ = v12
                                    _ = v13
                                }
                            default:
                                if (3.14159 < 1.5) {
                                    v3 = (65535 < 65535)
                                    v2 += v2
                                    if v3 {
                                        v9 = (("a" + v9) + "\n\t")
                                        switch "a" {
                                        case "\n\t":
                                            v9 = v9
                                        case "foo bar":
                                            v14 := []int{fi8}
                                            _ = v14
                                        case "":
                                            fmt.Println(fi8, v1, v7)
                                        }
                                        v3 = false
                                    } else {
                                        v15 := []int{fi8, 128, -257}
                                        _ = v15
                                    }
                                } else {
                                    v3 = v3
                                    switch fi6 {
                                    case 127:
                                        v1 += (^(v1))
                                    case 1:
                                        var v16 string = ((v9 + v9) + "")
                                        _ = v16
                                    case 42:
                                        v17 := []int{}
                                        _ = v17
                                    default:
                                        var v18 int16 = v1
                                        _ = v18
                                    }
                                    switch (fi6 << ((1) & 63)) {
                                    case -1:
                                        if (false && v3) {
                                            v19 := GenS3{F0: 0, F1: (v4 / (((v4 | v4)) | 1)), F2: v9}
                                            v2 = 4294967295
                                            var v20 int16 = -1
                                            fmt.Println(v7)
                                            v9 = v9
                                            var v21 float32 = (0.1 + (0.0 / (1e10 - (1e10 + (-2.25 + 1e10)))))
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        } else {
                                            if (!(!v3)) {
                                                if ((!true) || (65535 > (v4 ^ 255))) {
                                                    var v22 uint64 = (v2 & 42)
                                                    v2 = 4294967295
                                                    var v23 int64 = 9223372036854775807
                                                    _ = v22
                                                    _ = v23
                                                }
                                            } else {
                                                var v24 int64 = 512
                                                if (v4 != (v4 >> ((42) & 63))) {
                                                    var v25 uint64 = (v2 << ((65535) & 63))
                                                    if (255 != 42) {
                                                        v26 := GenS4{F0: 42, F1: v2}
                                                        _ = v26
                                                    }
                                                    var v27 int64 = -455
                                                    _ = v25
                                                    _ = v27
                                                } else {
                                                    var v28 int32 = 336
                                                    v29 := GenS4{F0: 18446744073709551615, F1: 1}
                                                    _ = v28
                                                    _ = v29
                                                }
                                                _ = v24
                                            }
                                        }
                                    default:
                                        v30 := []int{}
                                        _ = v30
                                    }
                                }
                            }
                            var v31 int64 = 445
                            v32 := GenS2{F0: (v4 >> ((1) & 63))}
                            fmt.Println(fi6)
                            _ = v31
                            _ = v32
                        } else {
                            fmt.Println(v1, fi6, v3)
                        }
                        _ = v10
                        _ = v11
                    }
                    v33 := GenS3{F0: 0, F1: (v4 + uint16(v2)), F2: v9}
                    var v34 uint64 = 1
                    _ = fi8
                    _ = v33
                    _ = v34
                    _ = v9
                }
                for fi35 := 0; fi35 < 7; fi35++ {
                    if v3 {
                        var v36 bool = (v3 || v3)
                        _ = v36
                    } else {
                        var v37 uint64 = (v2 * (v2 & v2))
                        v38 := GenS3{F0: -28, F1: 279, F2: (("hello" + "hello") + "")}
                        var v39 uint64 = (v2 + (v2 ^ (v37 / (((v2 >> ((65535) & 63))) | 1))))
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    _ = fi35
                }
                _ = fi6
                _ = v7
            }
        }
        fmt.Println(v4, v3)
        _ = v5
    }
    fmt.Println(v3, v4)
    if true {
        v3 = ((v4 % (((v4 >> ((65535) & 63))) | 1)) <= (v4 ^ 191))
        v4 -= (v4 << ((1) & 63))
        fmt.Println(v2)
    }
    v1 += (v1 + (v1 | v1))
    fmt.Println(v2, v1)
    v40 := GenS5{F0: 494}
    v3 = (158 < v2)
    switch ("" + ("foo bar" + "a")) {
    case "hello":
        v1 -= -32768
    default:
        v4 = v4
    }
    v3 = v3
    v41 := GenS5{F0: 68}
    var v42 uint32 = 0
    for fi43 := 0; fi43 < 7; fi43++ {
        v3 = v3
        v4 -= (v4 & (v4 + (v4 << ((259) & 63))))
        v44 := GenS6{F0: -1}
        _ = fi43
        _ = v44
    }
    var v45 uint64 = (v2 | 851)
    var v46 int16 = v1
    fmt.Println(v42, v4, v42)
    v47 := []int{0, int(v4), -1, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v46
    _ = v47
}
