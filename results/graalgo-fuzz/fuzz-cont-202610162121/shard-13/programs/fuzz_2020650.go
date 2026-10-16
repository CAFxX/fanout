package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[string]string)
    var v2 string = "hello"
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    if true {
        v2 = v2
        if true {
            v2 = "a"
        } else {
            v2 = v2
            v6 := []int{-420, 127, -9223372036854775808}
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 float64 = 3.14159
            fmt.Println(v8, v2)
            v2 = v2
            _ = fi7
            _ = v8
        }
    }
    delete(v4, 1)
    var v9 int32 = 1
    var v10 uint64 = 767
    var v11 uint8 = 0
    v12 := GenS1{F0: ("foo bar" + "\n\t")}
    switch v2 {
    case "foo bar":
        fmt.Println(v10, v9)
    case "hello":
        switch 32768 {
        case 32767:
            if false {
                v3[-2147483648] = v2
                switch 0 {
                case -1:
                    if false {
                        if false {
                            for fi13 := 0; fi13 < 4; fi13++ {
                                v14 := GenS1{F0: v2}
                                var v15 uint16 = 421
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                        }
                        var v16 string = (v2 + "hello")
                        fmt.Println(v2)
                        _ = v16
                    } else {
                        fmt.Println(v9)
                        var v17 uint16 = 673
                        fmt.Println(v9, v11)
                        _ = v17
                    }
                case 1:
                    fmt.Println(v2, v9)
                case 128:
                    if false {
                        fmt.Println(v10)
                        var v18 int64 = 897
                        var v19 int64 = v18
                        _ = v18
                        _ = v19
                    }
                }
            }
        }
    case "a":
        if false {
            v20 := []int{}
            v11 += (v11 >> ((uint(v10)) & 63))
            if ((v10 * v10) != v10) {
                var v21 bool = (float32((-2.25 + 1e10)) <= 3.14159)
                v22 := GenS1{F0: ""}
                var v23 uint8 = 42
                _ = v21
                _ = v22
                _ = v23
            } else {
                if (true && (!(4294967295 >= 2147483647))) {
                    var v24 int = 32768
                    v24 = 0
                    _ = v24
                } else {
                    fmt.Println(v2, v10, v2)
                }
                var v25 float64 = 1e10
                _ = v25
            }
            _ = v20
        } else {
            switch "a" {
            case "":
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 string = (v2 + (v2 + v2))
                    for fi28 := 0; fi28 < 14; fi28++ {
                        switch (v27 + (("\n\t" + "\n\t") + v2)) {
                        case "hello":
                            var v29 uint32 = 0
                            _ = v29
                        default:
                            v10 = 556
                        }
                        if false {
                            fmt.Println(v27, v9)
                        }
                        v30 := []int{(fi28 & (fi26 + (-(fi26)))), 0, fi26, fi26}
                        _ = fi28
                        _ = v30
                    }
                    _ = fi26
                    _ = v27
                }
            case "\n\t":
                v10 = v10
            default:
                var v31 uint64 = v10
                _ = v31
            }
            v9 = -1
            v11 -= (v11 ^ (v11 % (((v11 / ((v11) | 1))) | 1)))
        }
    }
    v2 = v2
    v32 := GenS1{F0: "foo bar"}
    fmt.Println(v9, v11)
    var v33 uint = 903
    var v34 string = "foo bar"
    switch -560 {
    case 697:
        v35 := GenS2{F0: 128}
        _ = v35
    case 9223372036854775807:
        v9 = (v9 << ((v33) & 63))
    case -452:
        if (!true) {
            var v36 string = "a"
            var v37 int = 9223372036854775807
            v38 := GenS3{F0: 0, F1: false}
            _ = v36
            _ = v37
            _ = v38
        }
    }
    var v39 uint32 = 119
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v9
}
