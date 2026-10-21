package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint8; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 bool; F2 int }
type GenS6 struct { F0 int8; F1 uint; F2 uint32 }
type GenS7 struct { F0 int }
func genFn0(p0 int32, p1 uint64) bool {
    v5 := []int{9223372036854775807, -9223372036854775808}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v5
    return ("a" > ((("foo bar" + "foo bar") + "hello") + "foo bar"))
}
func genFn1() uint8 {
    switch "" {
    case "":
        v6 := []int{860, 32767, 127, -746}
        _ = v6
    case "foo bar":
        var v7 int32 = 1
        _ = v7
    }
    v8 := []int{766}
    var v9 uint64 = 0
    if false {
        var v10 int = 42
        v10 += 397
        _ = v10
    }
    switch -1 {
    case 526:
        v9 = v9
    case -9223372036854775808:
        for fi11 := 0; fi11 < 11; fi11++ {
            if true {
                v12 := GenS1{F0: fi11, F1: 985}
                _ = v12
            } else {
                for fi13 := 0; fi13 < 1; fi13++ {
                    if (true && true) {
                        fmt.Println(fi11, fi13, fi11)
                        _ = len(v8)
                    } else {
                        var v14 bool = true
                        if v14 {
                            var v15 int16 = 32767
                            _ = v15
                        } else {
                            if (2 > 2) {
                                if len(v8) > 0 {
                                    v8[0] = fi11
                                }
                                var v16 int64 = 42
                                v17 := GenS1{F0: (fi11 >> ((734) & 63)), F1: 1}
                                _ = v16
                                _ = v17
                            } else {
                                v18 := GenS2{F0: (fi11 + fi11)}
                                v19 := GenS2{F0: fi13}
                                _ = v18
                                _ = v19
                            }
                            v14 = v14
                        }
                        fmt.Println(fi13)
                        _ = v14
                    }
                    v20 := GenS2{F0: (fi13 / ((fi11) | 1))}
                    v9 = (v9 - v9)
                    _ = fi13
                    _ = v20
                }
            }
            _ = fi11
        }
    case 2:
        var v21 uint32 = 2147483647
        _ = v21
    default:
        if ("foo bar" == ("a" + "foo bar")) {
            v9 = 1
            switch -747 {
            case -32:
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 uint8 = 0
                    _ = fi22
                    _ = v23
                }
            case 0:
                for fi24 := 0; fi24 < 18; fi24++ {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        v8 = append(v8, fi25)
                        v26 := GenS3{F0: fi24, F1: 30, F2: (^(v9))}
                        _ = fi25
                        _ = v26
                    }
                    fmt.Println(v9, v9)
                    _ = fi24
                }
            case 127:
                v8 = append(v8, 4294967295)
            }
        }
    }
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float64 = 1e10
    var v2 float64 = (v1 / (-2.25 + v1))
    var v3 string = ((("hello" + "hello") + "hello") + ((("\n\t" + "a") + "a") + "\n\t"))
    var v4 int8 = -127
    var v27 float64 = float64(1.5)
    v28 := GenS4{F0: v3}
    switch 2147483647 {
    case -9223372036854775808:
        for fi29 := 0; fi29 < 16; fi29++ {
            v30 := []int{4294967295, (fi29 >> ((42) & 63)), (^(fi29))}
            for fi31 := 0; fi31 < 10; fi31++ {
                fmt.Println(fi31)
                v32 := GenS5{F0: 2147483647, F1: true, F2: (^(fi29))}
                _ = fi31
                _ = v32
            }
            if true {
                var v33 uint = 18446744073709551615
                _ = v33
            }
            _ = fi29
            _ = v30
        }
    case -2147483648:
        v34 := GenS6{F0: (v4 / (((v4 / ((v4) | 1))) | 1)), F1: 255, F2: 1}
        _ = v34
    default:
        if false {
            if (-2147483648 >= 2147483647) {
                v4 = 98
                var v35 int8 = 127
                _ = v35
            } else {
                var v36 int8 = 42
                v4 -= (v4 / (((v36 << ((18446744073709551615) & 63))) | 1))
                _ = v36
            }
            fmt.Println(v1, v3, v1)
            var v37 int16 = -32768
            _ = v37
        } else {
            var v38 float64 = (v1 * 0.1)
            if (!(-566 >= 1)) {
                for fi39 := 0; fi39 < 1; fi39++ {
                    v38 += float64(((-2.25 / 1.5) - 0.1))
                    v40 := []int{fi39, (fi39 ^ fi39)}
                    var v41 float64 = 3.14159
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            } else {
                if ((!(18446744073709551615 < 4294967295)) || ((true && (v4 < 0)) || false)) {
                    v42 := GenS6{F0: 42, F1: 4294967295, F2: 2}
                    v43 := GenS2{F0: 1}
                    _ = v42
                    _ = v43
                }
                v3 = "\n\t"
                var v44 uint32 = 0
                _ = v44
            }
            _ = v38
        }
    }
    v45 := GenS7{F0: 127}
    v4 = 42
    v46 := []int{127}
    fmt.Println(v27)
    fmt.Println(v2, v27)
    var v47 int32 = 2147483647
    v3 = v3
    v46 = append(v46, -2147483648)
    v46 = append(v46, -2147483648)
    if len(v46) > 0 {
        v46[0] = 42
    }
    fmt.Println(v4, v27, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
}
