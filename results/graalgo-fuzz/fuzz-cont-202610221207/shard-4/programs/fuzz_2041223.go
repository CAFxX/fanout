package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
func genFn0(p0 int16, p1 uint8) int {
    for fi4 := 0; fi4 < 14; fi4++ {
        if ("hello" == "") {
            for fi5 := 0; fi5 < 6; fi5++ {
                var v6 int64 = -1
                v7 := []int{fi4, 127, (fi4 - fi4), -604}
                _ = fi5
                _ = v6
                _ = v7
            }
            var v8 float64 = -2.25
            _ = v8
        } else {
            v9 := GenS1{F0: 4294967295}
            var v10 int8 = -128
            _ = v10
            _ = v9
        }
        switch fi4 {
        case 32768:
            var v11 uint = 989
            _ = v11
        case 996:
            for fi12 := 0; fi12 < 5; fi12++ {
                p0 = -870
                if false {
                    if (!true) {
                        p0 += (p0 * 128)
                        v13 := GenS1{F0: 504}
                        if (!(true && (3.14159 >= -2.25))) {
                            v14 := []int{(fi12 + fi4)}
                            _ = v14
                        }
                        _ = v13
                    } else {
                        if false {
                            v15 := []int{fi4, (fi12 << ((2) & 63))}
                            if (!(false || (!true))) {
                                if (!((1e10 / 1.5) > (0.1 / ((0.0 - 1.5) - 1e10)))) {
                                    var v16 uint32 = 4294967295
                                    _ = v16
                                }
                            }
                            _ = v15
                        }
                        v17 := GenS2{F0: 2147483648, F1: 166, F2: ("" + ("foo bar" + ""))}
                        _ = v17
                    }
                    var v18 int64 = 1
                    _ = v18
                }
                var v19 int32 = 737
                _ = fi12
                _ = v19
            }
        case 9223372036854775807:
            p1 = p1
        }
        _ = fi4
    }
    var v20 int64 = 0
    var v21 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 0
}
func genFn1(p0 int8) uint8 {
    p0 -= -47
    var v22 int64 = 1
    p0 += 120
    _ = p0
    _ = v22
    return 128
}

func main() {
    var v1 int64 = -1
    var v2 float32 = 0.1
    var v3 int = -9223372036854775808
    if ("\n\t" == "a") {
        switch (v3 | v3) {
        case 770:
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS1{F0: 0}
                v1 -= v1
                _ = fi23
                _ = v24
            }
        }
    } else {
        for fi25 := 0; fi25 < 12; fi25++ {
            fmt.Println(v1, v2)
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 int = fi26
                switch "hello" {
                case "":
                    if (!false) {
                        v3 -= (fi25 + fi26)
                    } else {
                        switch "" {
                        case "\n\t":
                            var v28 int32 = 0
                            _ = v28
                        case "foo bar":
                            v29 := []int{-773}
                            _ = v29
                        }
                        if (1 == 2147483648) {
                            v2 += (v2 * float32(-2.25))
                            v30 := GenS1{F0: 2147483648}
                            v3 = (v3 % ((9223372036854775807) | 1))
                            _ = v30
                        } else {
                            v2 = v2
                            v27 += 32767
                        }
                    }
                case "a":
                    fmt.Println(fi26, v3, fi25)
                case "\n\t":
                    var v31 uint64 = 2
                    _ = v31
                }
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        v32 := GenS1{F0: 246}
        _ = v32
    }
    fmt.Println(v2, v3, v3)
    v33 := []int{0}
    v34 := GenS2{F0: 493, F1: 2, F2: "a"}
    v3 -= (v3 << ((2) & 63))
    v33 = append(v33, genFn0(577, 188))
    v33 = append(v33, v3)
    if len(v33) > 0 {
        v33[0] = 2147483647
    }
    switch genFn0(42, 141) {
    case -1:
        switch ("foo bar" + ((("\n\t" + "hello") + "foo bar") + "\n\t")) {
        case "foo bar":
            fmt.Println(v2, v1, v2)
        default:
            v2 += (0.1 / -2.25)
        }
    }
    switch v3 {
    case 0:
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
    case 857:
        if true {
            fmt.Println(v3, v2)
            var v35 uint16 = 255
            fmt.Println(v35, v1)
            _ = v35
        } else {
            v2 = v2
        }
    }
    var v36 uint16 = 564
    for fi37 := 0; fi37 < 7; fi37++ {
        v2 = v2
        if (2147483647 > 0) {
            _ = len(v33)
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
}
