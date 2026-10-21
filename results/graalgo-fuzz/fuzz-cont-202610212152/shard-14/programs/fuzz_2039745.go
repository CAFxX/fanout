package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 string; F1 uint16; F2 string }
type GenS3 struct { F0 int64; F1 int64; F2 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0(p0 int8, p1 int8) uint16 {
    v7 := []int{}
    v8 := GenS1{F0: p1, F1: 42}
    var v9 uint = 42
    switch -9223372036854775808 {
    case -2147483648:
        fmt.Println(p0, p0, p1)
    default:
        if false {
            switch "a" {
            case "a":
                fmt.Println(p1, p1)
            case "\n\t":
                fmt.Println(v9, v9)
            }
            var v10 string = "a"
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v10
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS1{F0: p0, F1: 2}
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi11
                _ = v12
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 938
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint64 = 765
    v3 := make(map[int]int)
    var v4 uint8 = 42
    var v5 uint8 = v4
    v6 := make(map[string]int)
    v13 := GenS1{F0: 127, F1: v2}
    if ((("hello" + "\n\t") + "foo bar") == "a") {
        v2 -= 1
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 uint = 719
            _ = fi14
            _ = v15
        }
    }
    if false {
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS2{F0: ("foo bar" + ""), F1: 255, F2: "\n\t"}
            v18 := GenS2{F0: ("\n\t" + "hello"), F1: 2, F2: "\n\t"}
            _ = fi16
            _ = v17
            _ = v18
        }
        if false {
            if false {
                fmt.Println(v2, v5)
            } else {
                switch 127 {
                case 759:
                    var v19 int8 = 0
                    _ = v19
                }
            }
        } else {
            v20 := v3[v1]
            _ = v20
        }
    } else {
        v21 := GenS1{F0: 13, F1: v2}
        if (4294967295 >= v1) {
            if true {
                for fi22 := 0; fi22 < 3; fi22++ {
                    for fi23 := 0; fi23 < 13; fi23++ {
                        v4 += (v5 + v5)
                        v2 = v2
                        _ = fi23
                    }
                    _ = fi22
                }
            } else {
                v2 = 42
            }
            v5 = v5
            v24 := []int{}
            _ = v24
        }
        v25 := GenS3{F0: -9223372036854775808, F1: 2, F2: false}
        _ = v21
        _ = v25
    }
    if false {
        v26 := []int{(v1 - v1), v1, v1}
        _ = v26
    } else {
        var v27 uint32 = 65535
        v28 := []int{v1, (v1 & 9223372036854775807), (^(v1))}
        v2 += v2
        _ = v27
        _ = v28
    }
    var v29 float32 = 1e10
    v30 := GenS4{F0: (^(v2))}
    switch "a" {
    case "hello":
        fmt.Println(v29, v4, v4)
    case "":
        _ = len(v6)
    case "foo bar":
        v31 := GenS3{F0: 29, F1: -9223372036854775808, F2: ((!true) || ("foo bar" == (("" + "\n\t") + "foo bar")))}
        _ = v31
    }
    v5 = (v4 - v4)
    switch (v1 + -9223372036854775808) {
    case -1:
        for fi32 := 0; fi32 < 12; fi32++ {
            fmt.Println(v2)
            var v33 float64 = (-2.25 / 1.5)
            v33 += 3.14159
            _ = fi32
            _ = v33
        }
    case 9223372036854775807:
        var v34 uint = 1
        _ = v34
    case 1:
        fmt.Println(v2, v1)
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        var v36 float64 = 0.1
        _ = fi35
        _ = v36
    }
    v37 := GenS2{F0: "hello", F1: 65535, F2: ("foo bar" + "foo bar")}
    v38 := GenS5{F0: 42, F1: 1, F2: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
