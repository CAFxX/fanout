package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = float32((0.1 * float64(1e10)))
    v3 := make(map[int]bool)
    var v4 int64 = -1
    v2 -= 0.0
    var v5 uint8 = 0
    v6 := GenS1{F0: false}
    if (42 > 444) {
        var v7 int8 = 100
        var v8 uint32 = 65535
        _ = v7
        _ = v8
    } else {
        v9 := []int{-650, 274, -9223372036854775808}
        var v10 int8 = -128
        var v11 int8 = v10
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := GenS2{F0: 0}
    if false {
        fmt.Println(v5)
        v13 := GenS3{F0: 1}
        _ = v13
    } else {
        switch -1 {
        case -709:
            v2 = (v2 - (3.14159 - v2))
        default:
            v14 := v3[4294967295]
            _ = v14
        }
    }
    v2 = (v2 + v2)
    for fi15 := 0; fi15 < 10; fi15++ {
        if false {
            v5 -= v5
        }
        if true {
            var v16 uint16 = 2
            _ = v16
        }
        fmt.Println(v4, v2)
        _ = fi15
    }
    if true {
        switch ("hello" + "\n\t") {
        case "a":
            var v17 string = "foo bar"
            _ = v17
        case "":
            delete(v1, ("foo bar" + "foo bar"))
        case "foo bar":
            if false {
                switch "\n\t" {
                case "":
                    switch 127 {
                    case 0:
                        for fi18 := 0; fi18 < 12; fi18++ {
                            for fi19 := 0; fi19 < 13; fi19++ {
                                var v20 float64 = 0.1
                                _ = fi19
                                _ = v20
                            }
                            _ = len(v1)
                            v21 := GenS3{F0: -232}
                            _ = fi18
                            _ = v21
                        }
                    case 32768:
                        var v22 bool = (32767 >= -861)
                        _ = v22
                    default:
                        v4 = 42
                    }
                default:
                    v23 := []int{}
                    _ = v23
                }
                v5 = 26
            }
        }
        if ("\n\t" <= ("a" + "")) {
            _ = len(v1)
            var v24 float64 = -2.25
            v5 = (v5 * v5)
            _ = v24
        } else {
            if true {
                var v25 bool = true
                _ = v25
            }
            v1["\n\t"] = -2147483648
            for fi26 := 0; fi26 < 13; fi26++ {
                fmt.Println(v4, v5, fi26)
                v5 = v5
                switch 2147483647 {
                case 9223372036854775807:
                    for fi27 := 0; fi27 < 6; fi27++ {
                        if (!true) {
                            var v28 uint = 0
                            if true {
                                var v29 uint64 = 2
                                fmt.Println(v5)
                                switch "" {
                                case "a":
                                    if true {
                                        var v30 uint64 = (v29 - v29)
                                        _ = v30
                                    }
                                case "hello":
                                    fmt.Println(v28, fi26, v2)
                                }
                                _ = v29
                            }
                            v4 -= -560
                            _ = v28
                        }
                        var v31 uint64 = 18446744073709551615
                        v31 = 18446744073709551615
                        _ = fi27
                        _ = v31
                    }
                }
                _ = fi26
            }
        }
        v32 := []int{9223372036854775807, -1, -1}
        _ = v32
    } else {
        v33 := GenS2{F0: 1}
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
