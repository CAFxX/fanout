package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int }
type GenS3 struct { F0 int }
func genFn0(p0 int, p1 int, p2 uint) uint8 {
    fmt.Println(p1, p0)
    p0 += p1
    var v4 uint32 = 65535
    var v5 float32 = (0.0 * 0.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 string, p1 string, p2 string) uint32 {
    switch p2 {
    case "a":
        v6 := GenS1{F0: p2, F1: p2, F2: 2}
        _ = v6
    case "hello":
        switch -440 {
        case 569:
            switch 662 {
            case 42:
                p2 = p0
            case -499:
                switch p2 {
                case "hello":
                    for fi7 := 0; fi7 < 8; fi7++ {
                        if ((!false) && (!(1e10 >= (1.5 * -2.25)))) {
                            if (!false) {
                                switch p0 {
                                case "hello":
                                    for fi8 := 0; fi8 < 18; fi8++ {
                                        var v9 float64 = ((-2.25 - 0.1) - 0.1)
                                        _ = fi8
                                        _ = v9
                                    }
                                case "\n\t":
                                    for fi10 := 0; fi10 < 11; fi10++ {
                                        v11 := []int{(-(fi10))}
                                        v11 = append(v11, fi10)
                                        _ = fi10
                                        _ = v11
                                    }
                                }
                            } else {
                                v12 := []int{fi7, 0, 2147483647, fi7}
                                if false {
                                    switch p2 {
                                    case "\n\t":
                                        v13 := GenS1{F0: p1, F1: p1, F2: 127}
                                        _ = v13
                                    }
                                    var v14 uint16 = 306
                                    _ = v14
                                }
                                _ = v12
                            }
                        }
                        _ = fi7
                    }
                case "":
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v16 := GenS1{F0: p1, F1: p0, F2: 53}
                        fmt.Println(fi15)
                        fmt.Println(p1, fi15, p0)
                        _ = fi15
                        _ = v16
                    }
                case "a":
                    p2 = (("hello" + "a") + (((p0 + "a") + (p2 + p2)) + ""))
                }
            case -194:
                if true {
                    v17 := GenS2{F0: 540, F1: 153, F2: 1}
                    var v18 uint = 18446744073709551615
                    v19 := []int{}
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    fmt.Println(p2, p1, p1)
                    v20 := []int{-206, -1, 32768}
                    var v21 bool = (!(162 <= -9223372036854775808))
                    _ = v20
                    _ = v21
                }
            }
        default:
            var v22 uint8 = 120
            _ = v22
        }
    case "foo bar":
        p0 = p0
    }
    var v23 uint16 = 282
    p0 = ("foo bar" + "hello")
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "a"
    v3 := make(map[int]string)
    v24 := GenS2{F0: 4294967295, F1: 2, F2: 1}
    v25 := []int{-9223372036854775808, -2147483648, 127}
    var v26 int32 = 2147483647
    if true {
        var v27 uint32 = 2
        _ = v27
    } else {
        for fi28 := 0; fi28 < 4; fi28++ {
            v2 = v2
            _ = fi28
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 string = "hello"
            v2 = "a"
            _ = fi29
            _ = v30
        }
    }
    if true {
        var v31 uint32 = 2147483648
        v2 = (v2 + ((v2 + "hello") + "\n\t"))
        var v32 int32 = -2147483648
        _ = v31
        _ = v32
    } else {
        var v33 int8 = -1
        if false {
            v2 = ((v2 + v2) + v2)
            for fi34 := 0; fi34 < 12; fi34++ {
                v2 = v2
                v35 := GenS2{F0: 18446744073709551615, F1: 19, F2: fi34}
                v33 = (v33 | (-(v33)))
                _ = fi34
                _ = v35
            }
        }
        _ = v33
    }
    v36 := GenS3{F0: -881}
    v2 = v2
    var v37 bool = true
    v37 = false
    _ = len(v3)
    v38 := v1[genFn1("\n\t", "foo bar", ("a" + v2))]
    v39 := GenS1{F0: (v2 + v2), F1: v2, F2: 127}
    v37 = (-2147483648 > 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
