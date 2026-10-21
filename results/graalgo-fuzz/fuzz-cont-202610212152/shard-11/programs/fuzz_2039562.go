package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint8; F2 string }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint64 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 bool = true
        _ = fi4
        _ = v5
    }
    var v6 uint8 = 127
    _ = v6
    return 340
}
func genFn1() int16 {
    if (false && false) {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: 1}
            _ = fi7
            _ = v8
        }
        v9 := []int{1, 42, -644}
        _ = v9
    }
    var v10 bool = (763 != 255)
    v10 = (0.1 >= ((-2.25 / 1.0) * ((3.14159 + -2.25) + ((0.0 - float32((3.14159 / ((-2.25 - 0.0) * (float64(0.0) - (0.1 + 3.14159)))))) + 1e10))))
    _ = v10
    return 42
}

func main() {
    var v1 float64 = 0.1
    var v2 uint = 0
    v3 := make(map[uint32]uint32)
    fmt.Println(v2, v1)
    v11 := GenS1{F0: 128}
    var v12 uint8 = 42
    v13 := GenS2{F0: 272, F1: v12, F2: "hello"}
    fmt.Println(v1)
    var v14 int64 = 0
    fmt.Println(v14)
    for fi15 := 0; fi15 < 12; fi15++ {
        switch ((("\n\t" + ("\n\t" + ("" + ("foo bar" + "\n\t")))) + "hello") + "a") {
        case "\n\t":
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 uint32 = 2147483648
                fmt.Println(fi16)
                v2 = v2
                _ = fi16
                _ = v17
            }
        case "a":
            var v18 uint16 = 32767
            _ = v18
        case "hello":
            v19 := GenS3{F0: 977, F1: (0 <= 127)}
            _ = v19
        }
        v2 = 2
        if true {
            v14 = (v14 | int64(v2))
        } else {
            var v20 uint64 = 4294967295
            fmt.Println(v1)
            _ = v20
        }
        _ = fi15
    }
    if ("hello" != ((("foo bar" + ("\n\t" + "")) + (("foo bar" + "foo bar") + "hello")) + "hello")) {
        v21 := []int{}
        v12 = v12
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS2{F0: genFn1(), F1: v12, F2: ""}
        v12 += v12
        var v24 float64 = v1
        _ = fi22
        _ = v23
        _ = v24
    }
    v12 += (^(v12))
    _ = len(v3)
    fmt.Println(v12)
    fmt.Println(v1)
    switch ("foo bar" + "") {
    case "hello":
        if true {
            var v25 int8 = 100
            _ = v25
        } else {
            v2 = (v2 - 0)
        }
    case "":
        if true {
            for fi26 := 0; fi26 < 20; fi26++ {
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 float64 = (v1 * (v1 * ((0.1 / -2.25) - v1)))
                    v12 -= 43
                    if false {
                        var v29 uint16 = 2
                        v30 := GenS4{F0: 4294967295}
                        _ = v29
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                }
                if (127 >= 2) {
                    for fi31 := 0; fi31 < 1; fi31++ {
                        v2 += (v2 / ((1) | 1))
                        var v32 float32 = float32(float64(-2.25))
                        _ = fi31
                        _ = v32
                    }
                    switch "a" {
                    case "":
                        delete(v3, 2)
                    case "foo bar":
                        fmt.Println(v14)
                    case "\n\t":
                        if (false && (fi26 <= (fi26 + (fi26 - fi26)))) {
                            for fi33 := 0; fi33 < 11; fi33++ {
                                if (fi26 == 295) {
                                    v34 := []int{fi26, fi26}
                                    var v35 uint16 = 1
                                    _ = v34
                                    _ = v35
                                } else {
                                    var v36 int = fi33
                                    v14 = 598
                                    v12 += (v12 | 72)
                                    v37 := GenS2{F0: genFn1(), F1: v12, F2: "\n\t"}
                                    _ = v36
                                    _ = v37
                                }
                                if false {
                                    switch (-(fi33)) {
                                    case 42:
                                        v3[2] = 364
                                    default:
                                        var v38 uint32 = 978
                                        _ = v38
                                    }
                                    var v39 uint32 = 564
                                    _ = v39
                                }
                                var v40 int8 = 42
                                var v41 uint64 = 36
                                var v42 bool = true
                                _ = fi33
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            for fi43 := 0; fi43 < 20; fi43++ {
                                delete(v3, 0)
                                if true {
                                    v3[1] = 463
                                    fmt.Println(v2, v12, v14)
                                }
                                _ = fi43
                            }
                        }
                    }
                }
                _ = len(v3)
                _ = fi26
            }
            v2 += v2
        }
    case "\n\t":
        v44 := v3[2]
        _ = v44
    default:
        v45 := GenS2{F0: -1, F1: (v12 << (((v2 & uint(v12))) & 63)), F2: "a"}
        _ = v45
    }
    var v46 uint = (v2 | (v2 | v2))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v46
}
