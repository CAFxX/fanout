package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int64; F2 uint8 }

func main() {
    var v1 int32 = -1
    var v2 bool = ((v1 * -1) <= (v1 * v1))
    var v3 string = "foo bar"
    v4 := make(map[string]int)
    var v5 uint8 = 186
    var v6 string = ((v3 + "a") + v3)
    v6 = ((v6 + v3) + ((v6 + (v3 + ((v3 + v6) + "\n\t"))) + "hello"))
    v7 := []int{4294967295, -9223372036854775808, 32767, 4294967295}
    v8 := GenS1{F0: 258}
    fmt.Println(v5, v3)
    switch "\n\t" {
    case "foo bar":
        var v9 uint16 = 334
        _ = v9
    case "":
        _ = len(v7)
    case "a":
        switch v3 {
        case "hello":
            var v10 bool = (v2 || v2)
            _ = v10
        case "foo bar":
            var v11 bool = v2
            _ = v11
        case "a":
            var v12 uint8 = 128
            _ = v12
        default:
            fmt.Println(v1, v2, v5)
        }
    default:
        var v13 uint8 = (v5 + v5)
        _ = v13
    }
    v3 = (v3 + v3)
    fmt.Println(v6)
    v14 := GenS1{F0: 0}
    if (v3 >= "") {
        var v15 int16 = 592
        switch -710 {
        case 162:
            switch "" {
            case "hello":
                var v16 uint16 = 0
                _ = v16
            }
        }
        var v17 int = 711
        _ = v15
        _ = v17
    } else {
        if false {
            v3 = (v3 + (v6 + v6))
        }
        v7 = append(v7, 127)
    }
    var v18 int8 = 100
    v6 = ("hello" + ((v6 + ((v6 + v3) + v3)) + (v3 + "foo bar")))
    for fi19 := 0; fi19 < 2; fi19++ {
        v1 = (v1 - v1)
        if (true || (-1 >= -216)) {
            v20 := GenS1{F0: 1}
            v3 = (v3 + v3)
            _ = v20
        }
        v1 -= (v1 << ((65535) & 63))
        _ = fi19
    }
    var v21 uint64 = 0
    if true {
        if v2 {
            if true {
                if v2 {
                    _ = len(v7)
                    for fi22 := 0; fi22 < 20; fi22++ {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        if v2 {
                            for fi23 := 0; fi23 < 9; fi23++ {
                                v3 = (v6 + "a")
                                _ = fi23
                            }
                        } else {
                            switch "a" {
                            case "foo bar":
                                switch (fi22 % (((-(fi22))) | 1)) {
                                case 9223372036854775807:
                                    v18 = v18
                                case 42:
                                    var v24 int = (-(fi22))
                                    _ = v24
                                default:
                                    var v25 int = (^(fi22))
                                    _ = v25
                                }
                            }
                            var v26 bool = (0 >= 376)
                            _ = v26
                        }
                        _ = fi22
                    }
                    var v27 uint8 = uint8(v21)
                    _ = v27
                }
                switch 127 {
                case 127:
                    var v28 uint64 = 2
                    _ = v28
                }
            } else {
                switch -2147483648 {
                case 0:
                    v2 = v2
                case 9223372036854775807:
                    var v29 string = v6
                    _ = v29
                case -2147483648:
                    var v30 uint8 = 255
                    _ = v30
                default:
                    var v31 uint16 = 941
                    _ = v31
                }
                v2 = v2
                if true {
                    v5 = v5
                    v5 -= v5
                    var v32 uint8 = v5
                    _ = v32
                }
            }
        } else {
            v33 := GenS1{F0: 4294967295}
            _ = v33
        }
        v34 := GenS2{F0: -1}
        _ = v34
    }
    if (42 <= 991) {
        if v2 {
            _ = len(v7)
            v35 := GenS3{F0: v3, F1: -365, F2: v5}
            _ = v35
        }
        fmt.Println(v2, v2, v1)
    } else {
        v5 = (v5 / (((v5 & (v5 >> ((354) & 63)))) | 1))
        var v36 int16 = -29
        v37 := GenS1{F0: (v21 + v21)}
        _ = v36
        _ = v37
    }
    if len(v7) > 0 {
        v7[0] = -983
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
