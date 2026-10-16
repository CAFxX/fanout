package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int64; F2 int8 }
func genFn0(p0 bool, p1 int32, p2 uint32) uint16 {
    if p0 {
        for fi6 := 0; fi6 < 20; fi6++ {
            p0 = (!p0)
            _ = fi6
        }
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        p2 = p2
        v8 := GenS1{F0: -159, F1: p1}
        var v9 float32 = ((3.14159 / 0.1) + 3.14159)
        _ = fi7
        _ = v8
        _ = v9
    }
    if p0 {
        var v10 string = ("foo bar" + "a")
        if p0 {
            v11 := GenS2{F0: 42, F1: (p1 >> ((0) & 63))}
            _ = v11
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint32, p1 int16) uint8 {
    switch "\n\t" {
    case "foo bar":
        for fi12 := 0; fi12 < 4; fi12++ {
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 int = fi12
                fmt.Println(fi12, p1, fi13)
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 bool = false
                _ = fi15
                _ = v16
            }
            _ = fi12
        }
    case "hello":
        switch "foo bar" {
        case "":
            p1 = (-(p1))
        case "a":
            var v17 float64 = 3.14159
            _ = v17
        }
    case "a":
        fmt.Println(p1, p1)
    default:
        p1 = p1
    }
    switch 32768 {
    case 4294967295:
        p0 -= 944
    case 1:
        switch ("" + "a") {
        case "hello":
            p0 = (p0 << ((2) & 63))
        case "":
            var v18 uint16 = 864
            _ = v18
        case "a":
            p1 = p1
        }
    default:
        var v19 float32 = ((0.1 + 0.1) - float32(3.14159))
        _ = v19
    }
    if true {
        v20 := GenS3{F0: false, F1: 2}
        _ = v20
    } else {
        fmt.Println(p1, p0)
        var v21 uint64 = 817
        _ = v21
    }
    _ = p0
    _ = p1
    return 66
}

func main() {
    var v1 int16 = 826
    v2 := make(map[string]int)
    var v3 int16 = (v1 / ((42) | 1))
    var v4 int16 = v1
    var v5 uint16 = 65535
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := []int{fi22, -6}
        v24 := v2["hello"]
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 float32 = float32(1.5)
    for fi26 := 0; fi26 < 10; fi26++ {
        v27 := v2[("\n\t" + "\n\t")]
        for fi28 := 0; fi28 < 11; fi28++ {
            v29 := []int{fi26, (fi26 << ((42) & 63)), 9223372036854775807}
            var v30 uint = 785
            _ = fi28
            _ = v29
            _ = v30
        }
        v31 := GenS4{F0: 4294967295}
        _ = fi26
        _ = v27
        _ = v31
    }
    fmt.Println(v5, v1)
    v1 = 1000
    var v32 string = "hello"
    var v33 uint32 = 4294967295
    var v34 float32 = v25
    if ((-2.25 <= 0.1) && false) {
        if (v32 < v32) {
            for fi35 := 0; fi35 < 4; fi35++ {
                for fi36 := 0; fi36 < 17; fi36++ {
                    v37 := GenS1{F0: 9223372036854775807, F1: 674}
                    v38 := GenS4{F0: 585}
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                v39 := GenS1{F0: 42, F1: 0}
                _ = fi35
                _ = v39
            }
        } else {
            v4 = v4
            if true {
                v40 := GenS4{F0: 42}
                _ = v40
            }
        }
    }
    if true {
        fmt.Println(v1, v32)
        v41 := GenS5{F0: (2147483647 <= -2147483648), F1: 306, F2: 1}
        var v42 uint64 = 18446744073709551615
        _ = v41
        _ = v42
    }
    var v43 uint = 589
    var v44 float32 = float32(((3.14159 * (1.5 / (0.1 * ((3.14159 * 0.1) / 1e10)))) - (0.0 + 0.1)))
    if (v43 > (v43 >> (((v43 - 42)) & 63))) {
        fmt.Println(v5, v3)
        var v45 int8 = 0
        _ = v45
    } else {
        fmt.Println(v32, v34, v25)
        for fi46 := 0; fi46 < 7; fi46++ {
            v33 -= 946
            _ = fi46
        }
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        v4 += -32768
        _ = fi47
    }
    v48 := GenS4{F0: (^(v43))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v48
    _ = v5
}
