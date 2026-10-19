package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int32; F2 int64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }
func genFn0(p0 bool, p1 int16, p2 string) string {
    if p0 {
        switch (("\n\t" + p2) + (p2 + p2)) {
        case "hello":
            fmt.Println(p1, p0, p2)
        case "foo bar":
            var v7 int = -2147483648
            _ = v7
        case "\n\t":
            v8 := GenS1{F0: 736}
            _ = v8
        default:
            if (1.5 >= 1e10) {
                p0 = true
                switch 1 {
                case -139:
                    var v9 float64 = ((1e10 / 0.1) - 1.5)
                    _ = v9
                case -9223372036854775808:
                    var v10 int8 = 1
                    _ = v10
                }
                if p0 {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        for fi12 := 0; fi12 < 4; fi12++ {
                            if p0 {
                                p2 = ("\n\t" + "\n\t")
                                if ((fi11 ^ (fi11 ^ fi11)) <= (fi12 % (((^(fi12))) | 1))) {
                                    p0 = (!(0 >= 0))
                                    var v13 uint16 = 32767
                                    var v14 int = 9223372036854775807
                                    _ = v13
                                    _ = v14
                                } else {
                                    v15 := GenS2{F0: fi11, F1: -994, F2: -1}
                                    _ = v15
                                }
                            }
                            p2 = (("hello" + ((p2 + p2) + p2)) + (("\n\t" + ((("" + p2) + p2) + "a")) + (p2 + p2)))
                            _ = fi12
                        }
                        v16 := []int{fi11}
                        _ = fi11
                        _ = v16
                    }
                } else {
                    switch "a" {
                    case "":
                        v17 := []int{}
                        _ = v17
                    case "foo bar":
                        var v18 float32 = (1.5 - (float32((float64(float32((1e10 * float64(1e10)))) - -2.25)) / 3.14159))
                        _ = v18
                    default:
                        p2 = ("foo bar" + (p2 + ((p2 + p2) + p2)))
                    }
                }
            }
        }
    } else {
        var v19 uint32 = 2
        _ = v19
    }
    fmt.Println(p0, p2)
    var v20 int64 = -528
    var v21 uint32 = 656
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return p2
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]int)
    var v3 int16 = -1
    v4 := make(map[uint32]int)
    v5 := make(map[int]int)
    v6 := make(map[uint32]string)
    var v22 uint16 = 213
    v22 = 2
    _ = len(v6)
    v1 -= v1
    for fi23 := 0; fi23 < 6; fi23++ {
        fmt.Println(v22, v3, v3)
        v3 = (v3 ^ 1)
        v24 := GenS3{F0: "foo bar", F1: 1}
        _ = fi23
        _ = v24
    }
    var v25 int8 = -128
    for fi26 := 0; fi26 < 5; fi26++ {
        for fi27 := 0; fi27 < 17; fi27++ {
            v3 += 42
            _ = fi27
        }
        for fi28 := 0; fi28 < 6; fi28++ {
            if (!false) {
                v25 -= v25
                switch (fi28 + fi28) {
                case 29:
                    if (fi28 <= -317) {
                        v3 -= (v3 - v3)
                        if false {
                            var v29 int8 = (v25 / ((0) | 1))
                            v30 := GenS4{F0: 18446744073709551615, F1: (fi26 % (((-(fi28))) | 1))}
                            _ = v29
                            _ = v30
                        }
                        v31 := GenS1{F0: v22}
                        _ = v31
                    } else {
                        v32 := GenS4{F0: 4294967295, F1: 4294967295}
                        _ = v32
                    }
                case 2147483647:
                    var v33 uint = 1
                    _ = v33
                case 128:
                    var v34 bool = true
                    _ = v34
                default:
                    v25 = v25
                }
                var v35 int64 = 42
                _ = v35
            } else {
                v25 = 1
                if (1 >= -108) {
                    delete(v6, 444)
                    v25 -= 42
                    v3 += (-(v3))
                    var v36 int32 = 42
                    delete(v5, (^(fi26)))
                    _ = v36
                } else {
                    fmt.Println(v1, v1)
                    v37 := GenS2{F0: (fi28 >> ((480) & 63)), F1: -460, F2: -9223372036854775808}
                    _ = v37
                }
            }
            _ = fi28
        }
        v38 := GenS5{F0: 956}
        _ = fi26
        _ = v38
    }
    fmt.Println(v22)
    v22 += 42
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v3 = v3
    case "hello":
        v39 := GenS1{F0: (v22 / ((v22) | 1))}
        _ = v39
    case "\n\t":
        switch "" {
        case "":
            var v40 string = ("hello" + (("hello" + "hello") + "a"))
            _ = v40
        case "foo bar":
            v1 += float32(1.5)
        default:
            v41 := GenS6{F0: 127}
            _ = v41
        }
    }
    var v42 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
