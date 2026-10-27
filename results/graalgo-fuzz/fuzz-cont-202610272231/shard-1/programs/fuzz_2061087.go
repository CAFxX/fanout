package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int16; F2 bool }
type GenS4 struct { F0 int8; F1 uint32; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 int8; F2 int16 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{32767, -540}
        v9 := GenS1{F0: 32767}
        _ = len(v8)
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        switch 472 {
        case -191:
            var v10 int64 = -9223372036854775808
            _ = v10
        case -927:
            if false {
                var v11 uint32 = 2
                var v12 uint = 836
                _ = v11
                _ = v12
            } else {
                if false {
                    var v13 int32 = 144
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v15 := []int{(fi14 >> ((18446744073709551615) & 63)), fi14, (-(fi14)), (-(fi14))}
                        var v16 int64 = 0
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = v13
                } else {
                    var v17 int = 1
                    if (437 > 1) {
                        switch ("" + ("" + ("a" + "\n\t"))) {
                        case "foo bar":
                            v17 -= -2147483648
                        }
                        v18 := []int{(v17 + 128), v17, (v17 ^ (v17 | v17))}
                        v18 = append(v18, (v17 * v17))
                        _ = v18
                    } else {
                        v19 := GenS1{F0: 349}
                        for fi20 := 0; fi20 < 17; fi20++ {
                            for fi21 := 0; fi21 < 16; fi21++ {
                                v22 := []int{(^(fi21)), -516, fi21, (fi21 ^ (fi20 - v17))}
                                v17 = fi21
                                _ = len(v22)
                                _ = fi21
                                _ = v22
                            }
                            _ = fi20
                        }
                        for fi23 := 0; fi23 < 2; fi23++ {
                            var v24 uint32 = 1
                            _ = fi23
                            _ = v24
                        }
                        _ = v19
                    }
                    fmt.Println(v17)
                    _ = v17
                }
            }
        case 1:
            if true {
                var v25 bool = false
                v26 := GenS1{F0: 65535}
                _ = v25
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 4; fi27++ {
            v28 := GenS2{F0: (fi27 % (((fi27 >> ((106) & 63))) | 1))}
            v29 := []int{(-(fi27)), (fi27 >> ((0) & 63)), -314, fi27}
            _ = fi27
            _ = v28
            _ = v29
        }
    } else {
        var v30 float64 = 0.0
        var v31 uint16 = 42
        v32 := []int{-2147483648, 0}
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 int64 = 564
    var v34 uint64 = 42
    if false {
        if true {
            fmt.Println(v33)
            switch "\n\t" {
            case "":
                v35 := GenS3{F0: 22, F1: -32768, F2: false}
                _ = v35
            }
        } else {
            var v36 int8 = 100
            _ = v36
        }
    } else {
        var v37 uint32 = 4294967295
        _ = v37
    }
    _ = v33
    _ = v34
    return 125
}

func main() {
    var v1 int8 = 11
    var v2 uint64 = 174
    var v3 uint8 = 127
    var v4 int = 1
    v5 := make(map[string]bool)
    var v6 uint64 = v2
    v38 := GenS4{F0: -128, F1: 4294967295, F2: 0}
    v39 := GenS4{F0: v1, F1: 991, F2: 512}
    delete(v5, "")
    var v40 float32 = 1e10
    if true {
        v2 += v6
    } else {
        v4 = (v4 << ((4294967295) & 63))
    }
    v41 := GenS2{F0: v4}
    v3 = 107
    v2 = (^(v2))
    var v42 uint32 = 806
    delete(v5, "a")
    v42 = (v42 - v42)
    v43 := GenS5{F0: ""}
    var v44 uint8 = v3
    v4 -= v4
    var v45 float64 = 0.1
    for fi46 := 0; fi46 < 17; fi46++ {
        v47 := GenS6{F0: 76, F1: (v1 >> ((255) & 63)), F2: -933}
        v48 := GenS5{F0: "a"}
        v4 += (fi46 - (^(fi46)))
        _ = fi46
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
