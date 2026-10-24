package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -220
    var v3 uint16 = 2
    var v4 int32 = 2147483647
    var v5 float32 = (1e10 / ((1.5 - -2.25) / ((0.1 + 1.5) + (-2.25 + 1e10))))
    fmt.Println(v3, v2, v2)
    if false {
        v5 -= (1.5 + v5)
        switch 42 {
        case 9223372036854775807:
            switch 127 {
            case 0:
                var v6 int64 = -322
                _ = v6
            case 1:
                v3 = v3
            case 617:
                if true {
                    switch (("a" + ("a" + "a")) + "hello") {
                    case "a":
                        switch 775 {
                        case 0:
                            switch ("\n\t" + "") {
                            case "a":
                                var v7 uint16 = v3
                                _ = v7
                            }
                        case -1:
                            var v8 uint64 = 574
                            _ = v8
                        case -2147483648:
                            v9 := GenS1{F0: 851, F1: (0 == 2)}
                            _ = v9
                        default:
                            switch "\n\t" {
                            case "":
                                v5 -= float32(1e10)
                            case "\n\t":
                                if false {
                                    var v10 float32 = 0.0
                                    var v11 uint32 = 2147483647
                                    var v12 int8 = -114
                                    _ = v10
                                    _ = v11
                                    _ = v12
                                }
                            case "hello":
                                if (!((!(1e10 != 0.1)) && ((!true) || (v4 >= v4)))) {
                                    v1["foo bar"] = 555
                                    v13 := []int{-2147483648, 9223372036854775807, -2147483648, -2147483648}
                                    fmt.Println(v3, v5)
                                    _ = v13
                                } else {
                                    fmt.Println(v2, v5, v4)
                                    v3 += 255
                                    v4 = v4
                                }
                            }
                        }
                    }
                    v14 := GenS2{F0: 42, F1: 1}
                    if (true && (true || (-27 <= 42))) {
                        v5 = v5
                    }
                    _ = v14
                }
            }
        default:
            var v15 int = 42
            _ = v15
        }
    } else {
        if true {
            v5 += v5
        } else {
            fmt.Println(v4)
            v16 := GenS3{F0: 339}
            for fi17 := 0; fi17 < 11; fi17++ {
                fmt.Println(v2, v3)
                v3 -= v3
                _ = fi17
            }
            _ = v16
        }
    }
    var v18 int64 = (v2 / ((v2) | 1))
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS4{F0: v18, F1: -1}
        v18 = v18
        _ = fi19
        _ = v20
    }
    switch 42 {
    case 128:
        v5 = (float32(0.0) * v5)
    case -626:
        fmt.Println(v18)
    case -30:
        var v21 string = "foo bar"
        _ = v21
    default:
        v22 := []int{int(v18)}
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := GenS1{F0: 2, F1: (!(v2 == (v18 - v2)))}
        for fi25 := 0; fi25 < 2; fi25++ {
            v26 := GenS1{F0: 0, F1: ((fi23 + fi25) != (fi23 | fi23))}
            _ = fi25
            _ = v26
        }
        if ((v18 ^ v18) != 61) {
            v27 := GenS4{F0: (v18 & (v2 >> ((0) & 63))), F1: -1}
            _ = v27
        } else {
            delete(v1, ("" + "a"))
            v2 += (v2 & v2)
        }
        _ = fi23
        _ = v24
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        if true {
            v29 := v1["\n\t"]
            v30 := []int{2147483647, fi28}
            v31 := GenS1{F0: 516, F1: (("foo bar" >= "foo bar") || ((v4 >> ((1) & 63)) > (v4 - -2147483648)))}
            _ = v29
            _ = v30
            _ = v31
        } else {
            var v32 uint64 = 2
            fmt.Println(v3)
            v33 := GenS1{F0: (v32 | v32), F1: false}
            _ = v32
            _ = v33
        }
        if true {
            v34 := []int{fi28, (-(fi28)), 128, (fi28 >> ((900) & 63))}
            _ = v34
        }
        v3 = v3
        _ = fi28
    }
    v35 := []int{36, 37, 0}
    for fi36 := 0; fi36 < 12; fi36++ {
        v3 -= v3
        if ((v3 ^ v3) != v3) {
            var v37 bool = true
            v2 -= (v18 + 0)
            delete(v1, "foo bar")
            _ = v37
        } else {
            v35 = append(v35, -9223372036854775808)
        }
        _ = fi36
    }
    var v38 uint8 = 152
    var v39 float32 = (3.14159 / ((v5 - 1.5) - v5))
    _ = len(v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
