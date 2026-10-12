package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int8; F2 bool }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int64; F1 bool; F2 int8 }
func genFn0(p0 uint8, p1 string, p2 int32) uint32 {
    switch 32768 {
    case -1:
        var v4 uint32 = 0
        _ = v4
    case 682:
        var v5 bool = true
        _ = v5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{}
        switch p1 {
        case "hello":
            v8 := GenS1{F0: false, F1: 0}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}
func genFn1() int32 {
    var v9 string = "hello"
    v9 = (((v9 + "") + v9) + (v9 + ((v9 + "\n\t") + "\n\t")))
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            v9 = "foo bar"
            switch fi10 {
            case 213:
                fmt.Println(fi11, v9, fi11)
                if true {
                    v12 := GenS1{F0: false, F1: 822}
                    v13 := GenS1{F0: false, F1: 261}
                    _ = v12
                    _ = v13
                }
                if (42 <= 127) {
                    if true {
                        var v14 int16 = 1
                        _ = v14
                    } else {
                        var v15 int8 = 1
                        switch (fi10 - (fi10 ^ fi10)) {
                        case -978:
                            var v16 uint8 = 0
                            _ = v16
                        case 32767:
                            switch (-(fi11)) {
                            case -886:
                                switch -9223372036854775808 {
                                case 32768:
                                    var v17 int16 = 42
                                    _ = v17
                                case -1:
                                    v18 := []int{}
                                    _ = v18
                                case 2147483647:
                                    var v19 int16 = 42
                                    _ = v19
                                default:
                                    v20 := []int{(fi11 % ((fi11) | 1)), 128, fi11}
                                    _ = v20
                                }
                            }
                        case 1000:
                            v15 -= v15
                        }
                        _ = v15
                    }
                } else {
                    var v21 uint64 = 0
                    v22 := []int{(fi10 | (fi11 - (^(fi10))))}
                    _ = v21
                    _ = v22
                }
            case 42:
                var v23 uint = 395
                _ = v23
            case 32767:
                var v24 uint64 = 1
                _ = v24
            default:
                v25 := GenS2{F0: true}
                _ = v25
            }
            _ = fi11
        }
        var v26 int16 = -32768
        v27 := GenS1{F0: false, F1: 1}
        _ = fi10
        _ = v26
        _ = v27
    }
    _ = v9
    return -1
}

func main() {
    var v1 uint8 = 203
    var v2 uint8 = (v1 & (v1 >> ((65535) & 63)))
    var v3 string = "foo bar"
    switch -1 {
    case -333:
        fmt.Println(v3, v3, v2)
    case 4294967295:
        var v28 float64 = float64(0.1)
        _ = v28
    case -9223372036854775808:
        v3 = "foo bar"
    }
    v29 := GenS1{F0: (true || (!true)), F1: 0}
    for fi30 := 0; fi30 < 14; fi30++ {
        if false {
            var v31 int8 = -128
            fmt.Println(v2, v2)
            _ = v31
        } else {
            var v32 int = fi30
            _ = v32
        }
        _ = fi30
    }
    v33 := GenS3{F0: 702, F1: 1, F2: (!false)}
    v2 = (v2 % ((v1) | 1))
    var v34 int32 = 9
    for fi35 := 0; fi35 < 1; fi35++ {
        var v36 int32 = v34
        _ = fi35
        _ = v36
    }
    v34 = 42
    fmt.Println(v2)
    if true {
        switch 127 {
        case -572:
            v37 := GenS4{F0: 42, F1: -155}
            _ = v37
        case 363:
            var v38 string = v3
            _ = v38
        case 0:
            v39 := GenS5{F0: -1, F1: (!false), F2: 100}
            _ = v39
        }
        v40 := []int{2147483647, 875, -9223372036854775808, -9223372036854775808}
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
}
