package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int) bool {
    p0 = 9223372036854775807
    v6 := GenS1{F0: false}
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(fi7)
        if (!false) {
            fmt.Println(fi7)
            p0 = -840
        }
        _ = fi7
    }
    v8 := []int{}
    v9 := GenS2{F0: 170, F1: 130}
    _ = p0
    _ = v6
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 float64 = (3.14159 * (3.14159 - 3.14159))
    var v2 int32 = 189
    var v3 int = -348
    v4 := make(map[int]uint32)
    var v5 uint64 = 414
    v10 := v4[v3]
    var v11 int8 = -1
    v12 := GenS1{F0: false}
    var v13 uint16 = 2
    v2 = v2
    v11 -= 97
    for fi14 := 0; fi14 < 10; fi14++ {
        v13 = (^(v13))
        var v15 bool = true
        _ = fi14
        _ = v15
    }
    _ = len(v4)
    switch "" {
    case "foo bar":
        fmt.Println(v10, v5)
    default:
        if ((^(v10)) >= (v10 * (^(v10)))) {
            for fi16 := 0; fi16 < 15; fi16++ {
                v1 -= v1
                var v17 uint = 736
                var v18 int16 = -208
                _ = fi16
                _ = v17
                _ = v18
            }
            v5 -= v5
            var v19 string = "a"
            _ = v19
        } else {
            switch ("hello" + "foo bar") {
            case "foo bar":
                var v20 float64 = v1
                _ = v20
            case "hello":
                for fi21 := 0; fi21 < 11; fi21++ {
                    v22 := []int{int(v5), v3, v3}
                    if false {
                        var v23 int = fi21
                        v23 += v3
                        fmt.Println(v23, v10, v10)
                        _ = v23
                    } else {
                        if true {
                            var v24 float64 = v1
                            _ = v24
                        } else {
                            _ = len(v4)
                            v25 := GenS1{F0: (9223372036854775807 < 0)}
                            _ = v25
                        }
                        var v26 uint = 702
                        _ = v26
                    }
                    var v27 uint16 = (v13 >> ((1) & 63))
                    _ = fi21
                    _ = v22
                    _ = v27
                }
            case "a":
                if ((v5 << ((1) & 63)) <= (v5 & (v5 / (((v5 * v5)) | 1)))) {
                    v28 := GenS3{F0: (v11 | v11)}
                    var v29 int = 642
                    _ = v28
                    _ = v29
                } else {
                    v3 = v3
                    v5 -= (v5 * (^(v5)))
                }
            }
        }
    }
    var v30 int16 = 0
    switch "" {
    case "\n\t":
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 bool = false
            _ = fi31
            _ = v32
        }
    case "a":
        switch (("\n\t" + "a") + "") {
        case "\n\t":
            var v33 int64 = int64(v3)
            _ = v33
        case "foo bar":
            var v34 uint8 = 1
            _ = v34
        default:
            v35 := v4[-9223372036854775808]
            _ = v35
        }
    }
    v2 += (v2 / (((v2 & v2)) | 1))
    v36 := GenS3{F0: -1}
    switch "hello" {
    case "a":
        v37 := GenS1{F0: true}
        _ = v37
    case "foo bar":
        var v38 int8 = 126
        _ = v38
    case "":
        v39 := GenS1{F0: ((-(v30)) >= (v30 | (v30 | 42)))}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
}
