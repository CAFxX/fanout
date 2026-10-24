package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int64; F1 uint32; F2 string }
func genFn0(p0 int, p1 bool) bool {
    v7 := GenS1{F0: (p0 % (((p0 ^ p0)) | 1))}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v7
    return p1
}
func genFn1(p0 int32) uint8 {
    p0 = -513
    v8 := GenS2{F0: 194}
    if true {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := []int{}
            p0 += (p0 >> ((0) & 63))
            fmt.Println(p0, p0)
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS2{F0: 4294967295}
    _ = p0
    _ = v11
    _ = v8
    return 123
}

func main() {
    var v1 bool = (2 <= 42)
    v2 := make(map[int]uint32)
    var v3 int64 = 1
    v4 := make(map[int]uint32)
    var v5 bool = ("\n\t" <= "\n\t")
    v6 := make(map[int]int)
    switch 60 {
    case 648:
        var v12 int = 42
        _ = v12
    case 9223372036854775807:
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS1{F0: fi13}
                var v16 uint32 = 42
                v5 = true
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi13
        }
    }
    fmt.Println(v5, v1)
    v17 := GenS3{F0: 238}
    if false {
        if false {
            fmt.Println(v1, v5)
        } else {
            v3 = (-(v3))
            for fi18 := 0; fi18 < 20; fi18++ {
                if v5 {
                    v19 := GenS1{F0: 42}
                    var v20 int8 = 1
                    _ = v19
                    _ = v20
                }
                var v21 uint8 = 132
                _ = fi18
                _ = v21
            }
        }
        v6[32768] = 32768
        if (!(("foo bar" + "hello") == ("\n\t" + "\n\t"))) {
            if (-2.25 <= float64(1.5)) {
                v3 -= (^(v3))
            } else {
                var v22 uint64 = 17
                if (!(!((((3.14159 / ((-2.25 - 0.0) / 3.14159)) / 3.14159) * -2.25) >= ((-2.25 + 1.5) - (-2.25 / (3.14159 + 0.1)))))) {
                    fmt.Println(v5, v5)
                    switch ((("foo bar" + "foo bar") + "hello") + "") {
                    case "hello":
                        for fi23 := 0; fi23 < 14; fi23++ {
                            v24 := []int{42, (fi23 % (((fi23 + fi23)) | 1))}
                            if v5 {
                                for fi25 := 0; fi25 < 9; fi25++ {
                                    fmt.Println(fi23, fi23, fi23)
                                    var v26 uint8 = genFn1(0)
                                    var v27 int64 = -1
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                            } else {
                                v1 = (2147483648 < 42)
                                if true {
                                    v3 = v3
                                    v1 = (false || ((-2147483648 > -414) && ("\n\t" > "foo bar")))
                                    v1 = (667 < 65535)
                                } else {
                                    for fi28 := 0; fi28 < 9; fi28++ {
                                        var v29 int32 = 2147483647
                                        var v30 int32 = 773
                                        _ = fi28
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 string = "\n\t"
                                    if v1 {
                                        fmt.Println(v5)
                                        var v32 int64 = v3
                                        if len(v24) > 0 {
                                            v24[0] = fi23
                                        }
                                        _ = v32
                                    }
                                    _ = v31
                                }
                                if v5 {
                                    v1 = ((fi23 << ((42) & 63)) <= fi23)
                                }
                            }
                            _ = fi23
                            _ = v24
                        }
                    default:
                        switch 1 {
                        case -9223372036854775808:
                            v3 = (v3 & (v3 / (((v3 ^ (v3 >> ((65535) & 63)))) | 1)))
                        default:
                            if v5 {
                                v5 = v1
                                var v33 float64 = (0.0 + (float64(1.5) - 0.1))
                                _ = v33
                            } else {
                                v34 := GenS1{F0: -677}
                                var v35 float32 = float32(1.5)
                                _ = v34
                                _ = v35
                            }
                        }
                    }
                    var v36 int32 = -2147483648
                    _ = v36
                }
                _ = v22
            }
            v37 := []int{}
            v38 := GenS4{F0: 200, F1: 633}
            _ = v37
            _ = v38
        } else {
            var v39 string = "hello"
            v5 = v1
            v4[32768] = 65535
            _ = v39
        }
    } else {
        switch 32768 {
        case 1:
            v5 = v5
        case -386:
            v40 := []int{244, 42}
            _ = v40
        case -1:
            switch ("hello" + "foo bar") {
            case "hello":
                v5 = (898 > 0)
            }
        default:
            var v41 int16 = -32768
            _ = v41
        }
        var v42 uint32 = 654
        _ = v42
    }
    if false {
        fmt.Println(v3)
        if (-1 == -304) {
            v43 := GenS5{F0: (v3 % ((-9223372036854775808) | 1)), F1: 0, F2: "a"}
            fmt.Println(v5, v1, v5)
            _ = v43
        }
        var v44 int64 = -212
        _ = v44
    }
    var v45 bool = (561 == 18446744073709551615)
    v45 = v45
    var v46 uint16 = 0
    v5 = (1 != 1)
    for fi47 := 0; fi47 < 13; fi47++ {
        var v48 int8 = -1
        v49 := []int{(fi47 ^ 577), fi47, 128, (-(fi47))}
        fmt.Println(v45, fi47)
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
