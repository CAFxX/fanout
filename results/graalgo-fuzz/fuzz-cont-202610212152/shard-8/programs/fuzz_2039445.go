package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 int; F1 int16; F2 uint64 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 uint16 = 42
        var v6 int8 = 127
        v7 := []int{(^(fi4)), (fi4 % ((127) | 1))}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: -87, F1: 93}
    _ = v8
    return 2
}
func genFn1() int {
    var v9 int = 42
    switch (v9 * (-(v9))) {
    case 128:
        fmt.Println(v9)
    case 4294967295:
        v10 := GenS1{F0: 394, F1: 0}
        _ = v10
    default:
        fmt.Println(v9)
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        v11 := []int{-522}
        _ = v11
    default:
        if false {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 string = ("foo bar" + "a")
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 int32 = -488
                var v16 string = "a"
                _ = fi14
                _ = v15
                _ = v16
            }
            if (2 < 284) {
                v9 = v9
            }
        } else {
            v17 := GenS2{F0: 497, F1: 826, F2: 806}
            switch "foo bar" {
            case "":
                var v18 bool = ((!true) && (false || (-1 != 605)))
                _ = v18
            case "\n\t":
                for fi19 := 0; fi19 < 19; fi19++ {
                    if (1 > 1000) {
                        var v20 float64 = 1e10
                        _ = v20
                    } else {
                        var v21 uint8 = 209
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 11; fi22++ {
                        var v23 int64 = 1
                        v23 = (v23 % ((-240) | 1))
                        var v24 uint64 = 4294967295
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS2{F0: 1, F1: 2147483648, F2: 4294967295}
                    _ = fi19
                    _ = v25
                }
            case "foo bar":
                if true {
                    var v26 uint16 = 577
                    _ = v26
                }
            }
            _ = v17
        }
    }
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 string = "a"
    var v2 bool = ((true || true) || false)
    var v3 uint32 = 2147483647
    for fi27 := 0; fi27 < 5; fi27++ {
        v28 := GenS2{F0: 42, F1: (v3 + 2), F2: 4294967295}
        _ = fi27
        _ = v28
    }
    var v29 int64 = 1
    if v2 {
        v30 := []int{42, -621}
        fmt.Println(v3, v2)
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        _ = v30
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        switch v1 {
        case "\n\t":
            v32 := []int{}
            _ = v32
        case "hello":
            v3 = (v3 | v3)
        case "a":
            v33 := GenS1{F0: -1, F1: 0}
            _ = v33
        default:
            var v34 bool = true
            _ = v34
        }
        _ = fi31
    }
    v35 := GenS3{F0: -962, F1: 42, F2: 1}
    v1 = (v1 + ((((v1 + v1) + v1) + "") + "\n\t"))
    fmt.Println(v2)
    v36 := []int{42, 32767, 32768}
    v2 = (v3 > 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v36
}
