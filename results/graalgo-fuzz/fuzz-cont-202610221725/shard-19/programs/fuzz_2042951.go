package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 343
    v3 := make(map[int]uint32)
    var v4 uint8 = 1
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]string)
    if (0 < 65535) {
        if false {
            var v7 int32 = 2147483647
            for fi8 := 0; fi8 < 9; fi8++ {
                _ = len(v1)
                _ = fi8
            }
            var v9 int32 = (-(v7))
            _ = v7
            _ = v9
        } else {
            v10 := []int{}
            _ = v10
        }
        fmt.Println(v4, v2)
    } else {
        fmt.Println(v2)
        v11 := []int{127, -198, int(v4), 0}
        _ = v11
    }
    var v12 uint8 = v4
    var v13 uint32 = 4294967295
    var v14 int8 = 1
    v15 := GenS1{F0: "hello", F1: (v13 ^ (v2 << ((4294967295) & 63)))}
    fmt.Println(v2, v13, v12)
    var v16 float64 = 1.5
    v2 = (v13 | 2147483648)
    if true {
        for fi17 := 0; fi17 < 20; fi17++ {
            _ = len(v6)
            _ = fi17
        }
        if true {
            v18 := GenS1{F0: ("foo bar" + ""), F1: v13}
            fmt.Println(v14, v4, v2)
            _ = v18
        } else {
            var v19 int32 = 42
            var v20 bool = (!true)
            _ = v19
            _ = v20
        }
        var v21 string = ("\n\t" + "\n\t")
        _ = v21
    } else {
        v16 -= (0.0 / 0.1)
        if false {
            if true {
                v22 := []int{int(v14), 128}
                _ = v22
            }
            if false {
                _ = len(v3)
            } else {
                fmt.Println(v4)
                v23 := GenS2{F0: "hello", F1: 335}
                _ = v23
            }
            var v24 uint16 = 0
            _ = v24
        } else {
            switch "hello" {
            case "a":
                v2 = v2
            case "":
                v25 := []int{32767, -9, 128, 9223372036854775807}
                _ = v25
            case "\n\t":
                v12 += (v4 / ((v12) | 1))
            default:
                v26 := GenS1{F0: (("a" + ("a" + "")) + "hello"), F1: 2}
                _ = v26
            }
        }
    }
    _ = len(v6)
    switch -582 {
    case 4294967295:
        v2 = (v13 - 804)
    case 912:
        v14 = 0
    case 797:
        if false {
            switch "\n\t" {
            case "hello":
                _ = len(v1)
            case "":
                for fi27 := 0; fi27 < 11; fi27++ {
                    for fi28 := 0; fi28 < 14; fi28++ {
                        if false {
                            v29 := GenS2{F0: "hello", F1: 1}
                            _ = v29
                        } else {
                            v14 -= (v14 >> ((255) & 63))
                            fmt.Println(v14, v16)
                            v13 = (v13 / ((v2) | 1))
                        }
                        _ = fi28
                    }
                    switch fi27 {
                    case -2147483648:
                        for fi30 := 0; fi30 < 5; fi30++ {
                            fmt.Println(v12, v4)
                            _ = fi30
                        }
                    default:
                        switch (fi27 % ((fi27) | 1)) {
                        case 2147483647:
                            var v31 int64 = 507
                            _ = v31
                        }
                    }
                    _ = fi27
                }
            default:
                fmt.Println(v16)
            }
            if ("hello" <= "") {
                var v32 uint16 = 1
                switch ("a" + "hello") {
                case "a":
                    v33 := GenS2{F0: ((("foo bar" + ("foo bar" + "a")) + "foo bar") + "foo bar"), F1: 255}
                    _ = v33
                case "hello":
                    v34 := GenS3{F0: false, F1: -1}
                    _ = v34
                case "foo bar":
                    v35 := GenS4{F0: 0}
                    _ = v35
                }
                v12 += 42
                _ = v32
            } else {
                v36 := GenS1{F0: "", F1: (^(v2))}
                var v37 int8 = (^(v14))
                switch 353 {
                case 188:
                    v14 += 127
                case -2147483648:
                    v16 = v16
                default:
                    v38 := GenS1{F0: "foo bar", F1: 42}
                    _ = v38
                }
                _ = v36
                _ = v37
            }
        } else {
            var v39 int32 = -1
            switch (("\n\t" + "a") + "a") {
            case "":
                v5[2147483648] = 128
            default:
                var v40 uint16 = 114
                _ = v40
            }
            _ = v39
        }
    }
    switch "hello" {
    case "foo bar":
        if false {
            var v41 uint64 = 2
            _ = v41
        }
    case "a":
        v42 := GenS5{F0: (true && true)}
        _ = v42
    default:
        fmt.Println(v2, v16)
    }
    var v43 string = "a"
    fmt.Println(v4)
    var v44 uint16 = 32767
    for fi45 := 0; fi45 < 8; fi45++ {
        v13 = (v13 / (((v2 - (^(v13)))) | 1))
        _ = fi45
    }
    var v46 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v46
    _ = v5
    _ = v6
}
