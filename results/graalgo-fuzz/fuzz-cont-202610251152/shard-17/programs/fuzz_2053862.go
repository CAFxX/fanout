package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint8 }
func genFn0(p0 int32) int {
    var v5 int64 = 42
    switch "hello" {
    case "foo bar":
        v5 -= (v5 % ((-1) | 1))
    case "":
        var v6 int16 = 42
        _ = v6
    }
    v7 := GenS1{F0: 1}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint16 = 42
        switch "hello" {
        case "":
            v10 := []int{(fi8 << ((4294967295) & 63)), fi8}
            _ = v10
        case "\n\t":
            for fi11 := 0; fi11 < 20; fi11++ {
                if (42 >= 18446744073709551615) {
                    v12 := GenS1{F0: v9}
                    _ = v12
                } else {
                    switch (-(fi8)) {
                    case -9223372036854775808:
                        v13 := GenS2{F0: ((!(!false)) || (1e10 != 0.0)), F1: 466}
                        _ = v13
                    case 32768:
                        v14 := []int{fi8}
                        _ = v14
                    case 612:
                        if (!(127 != 62)) {
                            fmt.Println(v9, fi11)
                            v15 := GenS3{F0: 42, F1: (v9 >> ((65535) & 63))}
                            _ = v15
                        } else {
                            var v16 int8 = 0
                            _ = v16
                        }
                    default:
                        v17 := []int{(fi11 << ((323) & 63)), (fi8 & 42), (^(fi8)), fi11}
                        _ = v17
                    }
                }
                var v18 int8 = 0
                _ = fi11
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := GenS4{F0: 1, F1: 255}
            _ = fi19
            _ = v20
        }
        _ = fi8
        _ = v9
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS5{F0: 596}
        if false {
            var v23 uint32 = 2
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = v5
    _ = v7
    return 42
}

func main() {
    var v1 uint = 65535
    var v2 uint8 = 42
    var v3 float32 = -2.25
    var v4 int16 = 707
    var v24 float32 = 0.1
    v1 -= (^(v1))
    var v25 int16 = v4
    var v26 int = 0
    for fi27 := 0; fi27 < 18; fi27++ {
        if true {
            switch "a" {
            case "hello":
                v28 := GenS6{F0: 42, F1: 4294967295, F2: 42}
                _ = v28
            default:
                fmt.Println(v3, v26, v2)
            }
            v29 := GenS3{F0: 42, F1: 42}
            v30 := GenS4{F0: v2, F1: 111}
            _ = v29
            _ = v30
        }
        _ = fi27
    }
    var v31 uint32 = 4294967295
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 uint = 1
        v3 = v3
        var v34 uint16 = 42
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 float64 = ((1e10 * 1e10) * 0.0)
    var v36 int8 = 1
    v4 -= v25
    for fi37 := 0; fi37 < 1; fi37++ {
        if false {
            v24 += (-2.25 + v3)
            if false {
                v35 += (((float64(((v3 * 0.1) - 1.5)) / float64(((v24 + v3) + ((v24 - ((v3 - (v3 / 0.1)) * 3.14159)) * -2.25)))) / (1e10 - float64(1.5))) + v35)
            }
        }
        for fi38 := 0; fi38 < 6; fi38++ {
            v36 = v36
            if false {
                v24 += v3
            }
            var v39 uint = (v1 | v1)
            _ = fi38
            _ = v39
        }
        _ = fi37
    }
    fmt.Println(v24, v25, v1)
    v40 := []int{v26}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v40
}
