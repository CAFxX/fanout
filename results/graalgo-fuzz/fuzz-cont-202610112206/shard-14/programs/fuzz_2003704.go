package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int32; F2 int64 }
type GenS3 struct { F0 int16 }
func genFn0() int16 {
    var v6 int32 = -2147483648
    v7 := GenS1{F0: -681}
    var v8 uint8 = 1
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 bool = (1 <= 42)
    var v2 bool = (-827 > 42)
    var v3 int = 4294967295
    var v4 uint64 = 18446744073709551615
    v5 := make(map[uint32]uint32)
    var v9 uint8 = 2
    switch "a" {
    case "foo bar":
        var v10 int8 = 32
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 11; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                switch (("foo bar" + "hello") + "foo bar") {
                case "foo bar":
                    v13 := []int{}
                    _ = v13
                default:
                    v4 = 42
                }
                _ = fi12
            }
            switch "" {
            case "foo bar":
                var v14 int8 = 127
                _ = v14
            case "\n\t":
                if v1 {
                    _ = len(v5)
                    var v15 uint16 = 42
                    _ = v15
                } else {
                    fmt.Println(v9, v2, v4)
                    if true {
                        v4 = v4
                        v9 -= 127
                        if (v2 && (!(4294967295 > (v4 - v4)))) {
                            v3 += 2147483647
                        } else {
                            if v2 {
                                if v1 {
                                    v2 = (42 <= -1)
                                    switch "\n\t" {
                                    case "a":
                                        for fi16 := 0; fi16 < 16; fi16++ {
                                            v17 := GenS2{F0: 65535, F1: 389, F2: -1}
                                            var v18 float32 = (3.14159 / 0.1)
                                            _ = fi16
                                            _ = v17
                                            _ = v18
                                        }
                                    case "hello":
                                        for fi19 := 0; fi19 < 15; fi19++ {
                                            var v20 string = "foo bar"
                                            switch fi19 {
                                            case 2147483647:
                                                var v21 float64 = -2.25
                                                _ = v21
                                            case -942:
                                                var v22 float64 = 3.14159
                                                _ = v22
                                            default:
                                                fmt.Println(v2, fi11)
                                            }
                                            _ = fi19
                                            _ = v20
                                        }
                                    case "\n\t":
                                        var v23 int8 = 127
                                        _ = v23
                                    default:
                                        v9 += v9
                                    }
                                    v3 = (fi11 & fi11)
                                }
                                fmt.Println(v9, fi11)
                            }
                            for fi24 := 0; fi24 < 1; fi24++ {
                                if false {
                                    if v2 {
                                        v3 += v3
                                        v25 := GenS3{F0: 512}
                                        _ = v25
                                    }
                                    v5[0] = 619
                                }
                                if v2 {
                                    delete(v5, 2)
                                    v26 := []int{(fi11 | fi24)}
                                    fmt.Println(fi24, v4)
                                    _ = v26
                                }
                                _ = len(v5)
                                _ = fi24
                            }
                        }
                    } else {
                        v27 := GenS3{F0: 288}
                        if v1 {
                            if (v9 == v9) {
                                fmt.Println(v2, fi11, v1)
                                v1 = v2
                            }
                        }
                        _ = v27
                    }
                    var v28 uint16 = 464
                    _ = v28
                }
            case "":
                fmt.Println(fi11, fi11, v3)
            }
            _ = fi11
        }
    case "a":
        _ = len(v5)
    default:
        if (1 <= 2) {
            var v29 string = ""
            _ = v29
        }
    }
    if true {
        var v30 string = "a"
        _ = v30
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        var v32 int8 = -128
        _ = fi31
        _ = v32
    }
    v2 = v2
    var v33 int64 = 0
    v34 := GenS1{F0: -375}
    fmt.Println(v9, v2)
    var v35 bool = (2147483647 < -2147483648)
    if (v1 && v1) {
        delete(v5, 2147483648)
        var v36 uint64 = (v4 + 348)
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
