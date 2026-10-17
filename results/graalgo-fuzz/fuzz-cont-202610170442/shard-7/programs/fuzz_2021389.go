package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
func genFn0() bool {
    var v5 int64 = -1
    var v6 uint8 = 42
    v6 = (v6 & 191)
    for fi7 := 0; fi7 < 6; fi7++ {
        if (0 > 0) {
            if false {
                fmt.Println(fi7, fi7)
                switch fi7 {
                case -1:
                    v8 := []int{(fi7 / ((fi7) | 1)), (^(fi7)), (fi7 & fi7)}
                    _ = v8
                case 32768:
                    var v9 bool = true
                    _ = v9
                }
                v5 -= (v5 % (((-(v5))) | 1))
            }
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v5)
        fmt.Println(v6, v5, fi10)
        _ = fi10
    }
    _ = v5
    _ = v6
    return (0 <= 570)
}
func genFn1(p0 int32, p1 uint) int16 {
    var v11 int8 = -1
    var v12 int = 0
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := []int{127, (v12 | fi13), -404, v12}
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int64 = -1
            var v17 int = fi15
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 6; fi18++ {
            fmt.Println(v12, fi13, p1)
            fmt.Println(p0)
            if (((!((p0 == (p0 + -1)) && false)) && (p1 <= 421)) && (!false)) {
                var v19 float32 = -2.25
                var v20 uint = (^(p1))
                _ = v19
                _ = v20
            }
            _ = fi18
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 1000
}

func main() {
    var v1 int32 = -307
    var v2 int32 = (v1 << ((67) & 63))
    var v3 uint = 688
    var v4 int64 = 0
    var v21 int32 = (^(v1))
    if (false && (false || false)) {
        var v22 string = "hello"
        v23 := []int{32767, -876}
        var v24 int8 = -1
        _ = v22
        _ = v23
        _ = v24
    } else {
        if true {
            switch 0 {
            case 1:
                switch 9223372036854775807 {
                case 128:
                    v25 := []int{}
                    _ = v25
                case 32768:
                    fmt.Println(v21)
                case 65:
                    switch 332 {
                    case -9223372036854775808:
                        v26 := GenS1{F0: (v3 > (v3 - v3)), F1: (v3 / ((v3) | 1))}
                        _ = v26
                    }
                }
            default:
                v4 = v4
            }
            v4 += v4
            var v27 uint32 = 633
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        fmt.Println(fi28)
        v29 := GenS1{F0: (fi28 <= (fi28 ^ fi28)), F1: (v3 / (((v3 >> ((v3) & 63))) | 1))}
        _ = fi28
        _ = v29
    }
    var v30 int = 980
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 uint64 = 420
        v33 := []int{v30}
        _ = fi31
        _ = v32
        _ = v33
    }
    if true {
        v34 := GenS1{F0: ((v4 >= (v4 | (v4 / (((v4 + v4)) | 1)))) || false), F1: 518}
        fmt.Println(v2, v4, v2)
        _ = v34
    }
    switch v30 {
    case 9223372036854775807:
        var v35 uint = (v3 + 255)
        _ = v35
    case 42:
        if (1000 <= -32768) {
            if true {
                var v36 bool = true
                switch ("" + "") {
                case "hello":
                    v37 := []int{v30, (^(v30)), (^(v30))}
                    _ = v37
                case "a":
                    var v38 uint64 = 98
                    _ = v38
                case "foo bar":
                    v39 := GenS1{F0: true, F1: v3}
                    _ = v39
                default:
                    var v40 float64 = ((1.5 - float64((1.5 + -2.25))) / 3.14159)
                    _ = v40
                }
                _ = v36
            }
            switch "hello" {
            case "foo bar":
                var v41 int64 = v4
                _ = v41
            }
        }
    case -517:
        v42 := []int{v30, (v30 * v30), 127}
        _ = v42
    default:
        for fi43 := 0; fi43 < 7; fi43++ {
            v44 := GenS1{F0: true, F1: 4294967295}
            v21 = -419
            v4 -= (v4 >> ((v3) & 63))
            _ = fi43
            _ = v44
        }
    }
    var v45 int8 = -1
    var v46 int16 = 32767
    v30 = (v30 + (v30 | (v30 + v30)))
    switch "a" {
    case "":
        var v47 float32 = 3.14159
        _ = v47
    }
    v48 := GenS1{F0: (-128 == v45), F1: v3}
    v49 := []int{}
    v46 = v46
    v1 -= v21
    var v50 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v46
    _ = v48
    _ = v49
    _ = v50
}
