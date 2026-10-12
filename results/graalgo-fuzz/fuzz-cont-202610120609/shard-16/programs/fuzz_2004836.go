package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 int16; F2 int8 }
func genFn0() int16 {
    v4 := []int{-741, -936}
    var v5 int = 4294967295
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint16, p1 uint32) string {
    fmt.Println(p1, p1)
    if false {
        v6 := []int{-407, 32767}
        var v7 float32 = 0.1
        v8 := GenS1{F0: true, F1: -1}
        _ = v6
        _ = v7
        _ = v8
    }
    switch -9223372036854775808 {
    case 604:
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 float32 = 3.14159
                v12 := []int{(fi9 << ((2) & 63))}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    }
    var v13 string = "foo bar"
    var v14 uint8 = 255
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return v13
}

func main() {
    var v1 bool = true
    var v2 uint16 = 42
    var v3 int8 = 0
    var v15 uint8 = 255
    v15 -= v15
    var v16 uint32 = 441
    if v1 {
        v1 = (365 != 704)
        var v17 int32 = -295
        switch "foo bar" {
        case "a":
            fmt.Println(v1, v17, v16)
        case "foo bar":
            if false {
                for fi18 := 0; fi18 < 18; fi18++ {
                    for fi19 := 0; fi19 < 1; fi19++ {
                        fmt.Println(v3)
                        if (673 < 592) {
                            v20 := GenS1{F0: v1, F1: 0}
                            _ = v20
                        } else {
                            v21 := GenS1{F0: (("\n\t" + "\n\t") != ""), F1: 0}
                            var v22 float64 = ((3.14159 * 0.0) / (((0.1 - 1e10) * 1e10) - 1e10))
                            _ = v21
                            _ = v22
                        }
                        fmt.Println(v16, v2, fi19)
                        var v23 int64 = 0
                        v3 += (v3 | int8(v23))
                        _ = fi19
                        _ = v23
                    }
                    _ = fi18
                }
            } else {
                v24 := GenS2{F0: "a", F1: (v3 >> ((1) & 63)), F2: v3}
                _ = v24
            }
        case "hello":
            fmt.Println(v2)
        default:
            if v1 {
                if v1 {
                    v25 := GenS3{F0: (("a" + ("hello" + "a")) + "foo bar")}
                    var v26 uint8 = 127
                    v27 := []int{-1}
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    v28 := []int{-9223372036854775808, 42}
                    v29 := GenS4{F0: genFn0(), F1: 1000, F2: v3}
                    _ = v28
                    _ = v29
                }
                fmt.Println(v17)
            } else {
                v15 = v15
                var v30 float64 = (3.14159 * ((0.0 / 1e10) - 1.5))
                var v31 int = 842
                _ = v30
                _ = v31
            }
        }
        _ = v17
    } else {
        v32 := []int{368}
        v32 = append(v32, 2147483647)
        _ = v32
    }
    var v33 uint = 2
    var v34 uint16 = v2
    v34 -= 2
    var v35 uint8 = v15
    for fi36 := 0; fi36 < 17; fi36++ {
        v2 -= (v34 | v34)
        _ = fi36
    }
    if v1 {
        var v37 uint16 = v2
        _ = v37
    } else {
        for fi38 := 0; fi38 < 17; fi38++ {
            v2 = v2
            for fi39 := 0; fi39 < 1; fi39++ {
                var v40 uint32 = v16
                _ = fi39
                _ = v40
            }
            v41 := []int{fi38}
            _ = fi38
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        if v1 {
            v43 := []int{}
            _ = v43
        }
        _ = fi42
    }
    switch ("hello" + ("a" + "\n\t")) {
    case "foo bar":
        switch 0 {
        case 9223372036854775807:
            var v44 int8 = 32
            _ = v44
        default:
            var v45 int = 32768
            _ = v45
        }
    case "":
        v46 := []int{1, 126, 1}
        _ = v46
    }
    v47 := GenS2{F0: "a", F1: v3, F2: int8(v16)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v47
}
