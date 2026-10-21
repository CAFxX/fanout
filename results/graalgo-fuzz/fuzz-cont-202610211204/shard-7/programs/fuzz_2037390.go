package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int64; F1 int8; F2 uint64 }
type GenS5 struct { F0 int32; F1 bool; F2 bool }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS9 struct { F0 uint64; F1 int16 }
func genFn0(p0 int64) uint8 {
    var v6 uint32 = 677
    v7 := GenS1{F0: 208, F1: 65535, F2: 24}
    if (((1.5 - 1.5) != 3.14159) || false) {
        var v8 float64 = float64(0.1)
        v9 := GenS1{F0: 0, F1: 255, F2: 42}
        fmt.Println(p0, v8, v6)
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: p0}
    switch 127 {
    case 127:
        if (255 != 128) {
            v11 := []int{-849, -76}
            v6 += v6
            v12 := GenS3{F0: 344, F1: (v6 - (^(v6))), F2: 2}
            _ = v11
            _ = v12
        }
    case -1:
        for fi13 := 0; fi13 < 2; fi13++ {
            if (!(0 < -18)) {
                var v14 uint32 = 593
                v15 := []int{(fi13 >> ((0) & 63))}
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 string = ""
                    fmt.Println(p0, fi16, p0)
                    v17 = ("\n\t" + (("foo bar" + (v17 + (v17 + v17))) + v17))
                    _ = fi16
                    _ = v17
                }
                v18 := GenS4{F0: 0, F1: 48, F2: 18446744073709551615}
                _ = v18
            }
            _ = fi13
        }
    }
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    return 51
}
func genFn1(p0 int8) bool {
    p0 -= p0
    var v19 string = "foo bar"
    var v20 uint64 = 1
    for fi21 := 0; fi21 < 17; fi21++ {
        switch (-(fi21)) {
        case -1:
            v20 += v20
        case 352:
            for fi22 := 0; fi22 < 1; fi22++ {
                v23 := []int{(fi22 | fi22), (fi21 * fi21), (fi22 + fi22), fi21}
                var v24 int8 = p0
                v19 = v19
                _ = fi22
                _ = v23
                _ = v24
            }
        default:
            v19 = (v19 + "\n\t")
        }
        _ = fi21
    }
    v25 := GenS5{F0: 2147483647, F1: true, F2: true}
    _ = p0
    _ = v19
    _ = v20
    _ = v25
    return (32767 != 42)
}

func main() {
    var v1 int8 = 100
    v2 := []string{"", "a", ("\n\t" + "\n\t"), ("hello" + ("hello" + ("" + "hello")))}
    var v3 float64 = 3.14159
    var v4 uint = 200
    var v5 uint8 = 57
    v3 = v3
    switch ("" + ((("" + ((("a" + ("a" + ("" + "\n\t"))) + ("hello" + ("foo bar" + "hello"))) + "\n\t")) + ("a" + "hello")) + ("hello" + "hello"))) {
    case "a":
        v26 := GenS1{F0: 255, F1: 0, F2: (v1 | (v1 >> ((v4) & 63)))}
        _ = v26
    case "hello":
        v3 -= v3
    default:
        v5 -= genFn0(0)
    }
    switch (("\n\t" + "a") + "a") {
    case "a":
        v4 += v4
    case "foo bar":
        switch "a" {
        case "":
            v2 = append(v2, "foo bar")
        case "hello":
            var v27 int64 = 0
            _ = v27
        default:
            if (1 > 1000) {
                _ = len(v2)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                if true {
                    var v28 uint = v4
                    if (1 == 32767) {
                        var v29 uint16 = 490
                        switch 1 {
                        case 351:
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                        }
                        _ = v29
                    }
                    v1 += 127
                    _ = v28
                }
                fmt.Println(v5, v1, v3)
            }
        }
    }
    if (42 >= 42) {
        fmt.Println(v5, v5)
        fmt.Println(v1)
        v5 += (v5 % (((v5 | 1)) | 1))
    } else {
        var v30 uint16 = 255
        fmt.Println(v1)
        _ = v30
    }
    switch "" {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        for fi31 := 0; fi31 < 15; fi31++ {
            if len(v2) > 0 {
                v2[0] = "a"
            }
            _ = fi31
        }
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        v33 := GenS2{F0: 0}
        v34 := GenS6{F0: 42, F1: 1}
        _ = fi32
        _ = v33
        _ = v34
    }
    v1 = (v1 >> (((^(v4))) & 63))
    var v35 uint16 = uint16(v4)
    var v36 string = "foo bar"
    var v37 int = 42
    v38 := GenS7{F0: 0}
    v39 := GenS8{F0: (^(v35)), F1: 127, F2: 18446744073709551615}
    if len(v2) > 0 {
        v2[0] = (v36 + v36)
    }
    _ = len(v2)
    v40 := GenS9{F0: 396, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
