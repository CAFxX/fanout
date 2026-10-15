package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int }
func genFn0() bool {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 uint64 = 18446744073709551615
        v6 := GenS1{F0: 118}
        v5 += (^(v5))
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int = 127
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := []int{fi8, (fi8 >> ((2) & 63))}
        if false {
            v7 += (^(fi8))
        } else {
            v10 := GenS1{F0: 0}
            if false {
                v7 += 350
                for fi11 := 0; fi11 < 10; fi11++ {
                    if true {
                        var v12 string = "\n\t"
                        _ = v12
                    }
                    _ = fi11
                }
            }
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    _ = v7
    return true
}
func genFn1(p0 uint, p1 int, p2 uint16) uint64 {
    v13 := GenS1{F0: -734}
    var v14 bool = true
    fmt.Println(p1)
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := []int{(p1 << ((18446744073709551615) & 63)), fi15, -9223372036854775808}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 4294967295
}

func main() {
    var v1 bool = (1 != 0)
    v2 := []string{"hello", "hello", "a"}
    v3 := []uint32{0, 42, 2147483648}
    v17 := GenS2{F0: 976, F1: 42, F2: ("\n\t" + ("a" + "hello"))}
    fmt.Println(v1)
    var v18 bool = (-128 > 1)
    if (127 != 42) {
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 uint = 2
            fmt.Println(v18, v1)
            var v21 int32 = 0
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v1 = v1
    switch "foo bar" {
    case "a":
        if (float32(1.5) >= (((-2.25 - float32(-2.25)) - -2.25) * 3.14159)) {
            if v1 {
                var v22 int32 = -400
                var v23 uint16 = 2
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                _ = len(v2)
                _ = fi24
            }
        }
    case "hello":
        v18 = ((false && ((3.14159 + 1e10) > 0.0)) && (!false))
    }
    var v25 uint8 = 40
    v26 := GenS3{F0: -431}
    if (!(222 < 2)) {
        switch -271 {
        case 32768:
            v27 := GenS3{F0: -62}
            _ = v27
        case -2147483648:
            for fi28 := 0; fi28 < 9; fi28++ {
                switch -521 {
                case 2147483647:
                    if true {
                        v29 := GenS4{F0: 65535, F1: -1}
                        switch (fi28 % (((fi28 * (fi28 * fi28))) | 1)) {
                        case -9223372036854775808:
                            var v30 int64 = -511
                            _ = v30
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = v29
                    } else {
                        switch ("" + "") {
                        case "a":
                            switch fi28 {
                            case -381:
                                var v31 uint8 = v25
                                _ = v31
                            case 0:
                                switch ("hello" + "foo bar") {
                                case "foo bar":
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(v25, v1)
                                    default:
                                        fmt.Println(fi28)
                                    }
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case "":
                                        fmt.Println(fi28)
                                    case "hello":
                                        v32 := GenS4{F0: 255, F1: 0}
                                        _ = v32
                                    default:
                                        v33 := GenS5{F0: fi28}
                                        _ = v33
                                    }
                                }
                            }
                        case "\n\t":
                            v1 = v1
                        }
                        var v34 uint32 = 2147483648
                        v34 = (v34 & v34)
                        _ = v34
                    }
                case -1:
                    v35 := GenS6{F0: fi28, F1: 100}
                    _ = v35
                default:
                    if v1 {
                        var v36 int16 = 1
                        v25 += v25
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        _ = v36
                    } else {
                        fmt.Println(v18, v1)
                    }
                }
                _ = fi28
            }
        case 42:
            for fi37 := 0; fi37 < 20; fi37++ {
                v38 := GenS7{F0: fi37}
                for fi39 := 0; fi39 < 13; fi39++ {
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                    var v40 float32 = float32((-2.25 * float64((1.5 / 0.1))))
                    v1 = (!(((32767 == (fi39 % ((int(v25)) | 1))) && v1) && (!false)))
                    if (fi39 >= (fi37 >> ((2) & 63))) {
                        v40 += float32((float64(v40) / (((0.0 * float64(v40)) + -2.25) - (float64(v40) * 0.1))))
                        v2 = append(v2, "foo bar")
                        if false {
                            var v41 int8 = -2
                            v41 += -48
                            v18 = (!("a" <= "hello"))
                            _ = v41
                        }
                    } else {
                        fmt.Println(v40, v18, fi37)
                    }
                    _ = fi39
                    _ = v40
                }
                _ = fi37
                _ = v38
            }
        default:
            fmt.Println(v1)
        }
        var v42 string = "a"
        _ = v42
    }
    if (!false) {
        v1 = v18
    } else {
        for fi43 := 0; fi43 < 4; fi43++ {
            switch "hello" {
            case "":
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            case "a":
                v25 = v25
            }
            v1 = ((!(!(v1 && (2 <= 2147483648)))) && (0.1 == (float32(0.1) * ((1.5 / 1.5) - -2.25))))
            _ = fi43
        }
        v3 = append(v3, 565)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    if (!(v1 || ((-2147483648 > 763) || false))) {
        var v44 int8 = -90
        v1 = (!(-1 == 0))
        _ = v44
    }
    v45 := GenS4{F0: 65535, F1: 1}
    _ = len(v2)
    switch -1 {
    case 10:
        var v46 bool = (0 != -117)
        _ = v46
    case 329:
        for fi47 := 0; fi47 < 1; fi47++ {
            v1 = false
            _ = fi47
        }
    case -286:
        v18 = v18
    default:
        fmt.Println(v25, v25, v25)
    }
    fmt.Println(v18, v18)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v18 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v45
}
