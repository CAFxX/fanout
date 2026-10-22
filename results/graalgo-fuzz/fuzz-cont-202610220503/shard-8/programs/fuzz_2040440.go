package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 string; F2 bool }
func genFn0(p0 int8, p1 uint32) bool {
    var v4 string = ("a" + ("foo bar" + ""))
    v5 := []int{530, 936, 2147483647}
    v4 = v4
    v5 = append(v5, 261)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return true
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]string)
    v3 := []uint32{282, 924, 4294967295, 727}
    if false {
        var v6 string = "hello"
        for fi7 := 0; fi7 < 11; fi7++ {
            v1 += (v1 + -32768)
            for fi8 := 0; fi8 < 10; fi8++ {
                fmt.Println(fi8, v6, fi8)
                var v9 int32 = 42
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: (fi7 | (fi7 + 900))}
            _ = fi7
            _ = v10
        }
        fmt.Println(v6)
        _ = v6
    } else {
        v11 := GenS1{F0: 127}
        _ = v11
    }
    v12 := GenS1{F0: 9223372036854775807}
    if false {
        for fi13 := 0; fi13 < 15; fi13++ {
            switch (^(fi13)) {
            case 9223372036854775807:
                v1 -= v1
            case -353:
                if (!false) {
                    switch 9223372036854775807 {
                    case 2147483647:
                        switch fi13 {
                        case -690:
                            v14 := GenS1{F0: 0}
                            _ = v14
                        case 42:
                            var v15 uint8 = 255
                            _ = v15
                        case 32767:
                            v1 -= -1
                        default:
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                        }
                    case 32767:
                        var v16 int32 = 42
                        _ = v16
                    }
                }
            case -737:
                v17 := GenS1{F0: fi13}
                _ = v17
            }
            _ = fi13
        }
    } else {
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 bool = (636 >= 88)
            v19 = v19
            _ = fi18
            _ = v19
        }
        var v20 int64 = 0
        v1 += 1
        _ = v20
    }
    v1 += -32768
    var v21 int64 = 9223372036854775807
    var v22 int64 = v21
    switch "foo bar" {
    case "a":
        v23 := GenS1{F0: 2147483647}
        _ = v23
    }
    switch "" {
    case "hello":
        fmt.Println(v22, v22)
    default:
        v24 := GenS1{F0: 2147483647}
        _ = v24
    }
    var v25 uint8 = 0
    switch ("\n\t" + "\n\t") {
    case "":
        v26 := GenS2{F0: (^(v1)), F1: ("\n\t" + "\n\t"), F2: "foo bar"}
        _ = v26
    case "hello":
        var v27 bool = false
        _ = v27
    }
    switch 128 {
    case -2147483648:
        v22 = int64(v25)
    case 10:
        fmt.Println(v22, v21)
    default:
        v28 := GenS1{F0: -791}
        _ = v28
    }
    var v29 float32 = 3.14159
    var v30 uint64 = 42
    for fi31 := 0; fi31 < 13; fi31++ {
        for fi32 := 0; fi32 < 13; fi32++ {
            v22 = 9223372036854775807
            v33 := GenS3{F0: (("\n\t" + ("foo bar" + "\n\t")) + "hello")}
            if true {
                if false {
                    switch fi31 {
                    case 42:
                        v34 := GenS4{F0: v30, F1: v30, F2: (v22 + v22)}
                        _ = v34
                    case -236:
                        fmt.Println(v21, fi31, fi32)
                    case 127:
                        v29 += v29
                    default:
                        _ = len(v2)
                        v35 := GenS5{F0: (^(v30))}
                        _ = v35
                    }
                }
                var v36 bool = (!true)
                _ = v36
            }
            _ = fi32
            _ = v33
        }
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        if (4294967295 > 0) {
            v37 := GenS5{F0: 362}
            var v38 int32 = -1
            _ = v37
            _ = v38
        }
        _ = fi31
    }
    var v39 uint8 = v25
    v40 := GenS6{F0: (v30 / (((v30 & 234)) | 1)), F1: ("" + ("a" + "\n\t")), F2: (!false)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v39
    _ = v40
}
