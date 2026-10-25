package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int16; F1 int32; F2 int8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int = -2147483648
    var v2 uint16 = 0
    var v3 int64 = -736
    var v4 int = 483
    var v5 uint16 = (v2 + 652)
    var v6 int = -9223372036854775808
    var v7 float64 = (-2.25 / -2.25)
    switch ("" + "hello") {
    case "":
        v4 -= (v4 ^ v6)
    case "a":
        if true {
            if (!false) {
                fmt.Println(v5, v5)
                var v8 int64 = v3
                if false {
                    fmt.Println(v3)
                } else {
                    var v9 string = "a"
                    _ = v9
                }
                _ = v8
            } else {
                v10 := GenS1{F0: ((!(-2147483648 != -1)) || false)}
                for fi11 := 0; fi11 < 3; fi11++ {
                    if (2 != 335) {
                        v12 := []int{v1, 42, 2147483647}
                        _ = v12
                    } else {
                        v4 = (v6 / (((fi11 / ((v4) | 1))) | 1))
                        fmt.Println(v4, v1, fi11)
                        v13 := GenS1{F0: false}
                        _ = v13
                    }
                    v14 := GenS1{F0: true}
                    v5 -= (v5 - v2)
                    _ = fi11
                    _ = v14
                }
                _ = v10
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                v5 = v5
                for fi16 := 0; fi16 < 12; fi16++ {
                    fmt.Println(fi15, v3, v7)
                    _ = fi16
                }
                fmt.Println(v5, v6)
                _ = fi15
            }
            if true {
                v17 := GenS2{F0: 32767, F1: 32767}
                _ = v17
            }
        } else {
            var v18 uint64 = 42
            _ = v18
        }
    case "hello":
        fmt.Println(v3, v3, v2)
    }
    v1 = 32768
    v5 -= v2
    v19 := []int{v4}
    var v20 uint16 = 1
    switch 532 {
    case -9223372036854775808:
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 uint8 = 1
            fmt.Println(v1, fi21)
            for fi23 := 0; fi23 < 17; fi23++ {
                switch v4 {
                case -1:
                    if len(v19) > 0 {
                        v19[0] = fi21
                    }
                case 266:
                    v24 := GenS3{F0: 1, F1: -2147483648, F2: 1}
                    _ = v24
                default:
                    v25 := GenS3{F0: 42, F1: -586, F2: 0}
                    _ = v25
                }
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    default:
        if ((!false) || true) {
            v1 = v6
            if false {
                var v26 int8 = 127
                _ = v26
            }
        } else {
            switch ("\n\t" + ((("foo bar" + "") + "foo bar") + "foo bar")) {
            case "":
                for fi27 := 0; fi27 < 1; fi27++ {
                    v1 = (v4 % ((9223372036854775807) | 1))
                    v28 := GenS4{F0: 2}
                    _ = fi27
                    _ = v28
                }
            }
            var v29 int = -9223372036854775808
            if true {
                for fi30 := 0; fi30 < 15; fi30++ {
                    var v31 uint = 510
                    v7 = (float64(-2.25) - v7)
                    _ = len(v19)
                    _ = fi30
                    _ = v31
                }
            } else {
                v29 = v29
            }
            _ = v29
        }
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 bool = (42 != 1)
        fmt.Println(v3, v20)
        v5 = 1
        _ = fi32
        _ = v33
    }
    fmt.Println(v6, v5, v5)
    for fi34 := 0; fi34 < 6; fi34++ {
        if true {
            fmt.Println(v4)
        } else {
            var v35 int8 = 0
            _ = v35
        }
        var v36 bool = true
        _ = fi34
        _ = v36
    }
    switch 2147483647 {
    case 42:
        v37 := GenS4{F0: 1}
        _ = v37
    case -910:
        v3 = (v3 ^ v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
