package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint16, p1 int, p2 int16) uint8 {
    var v6 int = -1
    v7 := GenS1{F0: (p1 ^ 127), F1: 555, F2: 228}
    fmt.Println(p0, p0, p1)
    var v8 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (v8 / ((42) | 1))
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 42
    var v3 uint32 = 2
    var v4 uint = 0
    var v5 string = ("a" + "a")
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := []int{(-(fi9))}
        if len(v10) > 0 {
            v10[0] = (fi9 + fi9)
        }
        var v11 uint16 = 512
        _ = fi9
        _ = v10
        _ = v11
    }
    switch 42 {
    case -1:
        switch -1 {
        case 9223372036854775807:
            var v12 int = int(v3)
            _ = v12
        case -534:
            fmt.Println(v2)
        default:
            fmt.Println(v5, v4)
        }
    case 127:
        fmt.Println(v2, v2)
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint64 = 500
        var v15 uint = (^(v4))
        v1[32768] = v5
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v3)
    var v16 int32 = -261
    _ = len(v1)
    var v17 uint8 = 110
    var v18 uint32 = (v3 & 0)
    v18 = v18
    if false {
        v18 += v3
        v1[313] = "\n\t"
        v19 := []int{}
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        if true {
            v21 := GenS1{F0: 360, F1: (v18 | 1), F2: uint8(v4)}
            fmt.Println(v17, v3)
            v1[fi20] = (v5 + ((v5 + (("foo bar" + v5) + ((v5 + "") + v5))) + ("a" + v5)))
            _ = v21
        } else {
            v17 += v17
        }
        v3 -= 42
        _ = fi20
    }
    v22 := GenS1{F0: -2147483648, F1: 678, F2: 2}
    v23 := GenS2{F0: 32767, F1: 1}
    if false {
        if true {
            delete(v1, 42)
        } else {
            switch 0 {
            case 972:
                var v24 int = 830
                _ = v24
            }
        }
        switch -2147483648 {
        case -2147483648:
            fmt.Println(v5, v5)
        case -401:
            if (42 == 4294967295) {
                var v25 int32 = int32(v2)
                _ = v25
            }
        }
    } else {
        if false {
            var v26 int32 = (v16 * -1)
            for fi27 := 0; fi27 < 9; fi27++ {
                fmt.Println(v5)
                fmt.Println(fi27, v2, fi27)
                fmt.Println(v26)
                _ = fi27
            }
            _ = v26
        } else {
            for fi28 := 0; fi28 < 12; fi28++ {
                switch (("a" + ((v5 + "a") + v5)) + "foo bar") {
                case "\n\t":
                    v29 := GenS2{F0: uint16(v4), F1: 796}
                    _ = v29
                case "foo bar":
                    switch fi28 {
                    case 2147483647:
                        var v30 uint64 = 499
                        _ = v30
                    case -2147483648:
                        v3 = v3
                    case -9223372036854775808:
                        var v31 int = fi28
                        _ = v31
                    }
                case "a":
                    var v32 int32 = -2147483648
                    _ = v32
                }
                var v33 float64 = (-2.25 - 0.1)
                switch v5 {
                case "a":
                    v34 := []int{(fi28 % ((fi28) | 1)), (fi28 + (fi28 - (fi28 >> ((v4) & 63)))), (fi28 | (fi28 >> (((^(v4))) & 63))), int(v16)}
                    _ = v34
                case "hello":
                    v35 := GenS1{F0: (fi28 - (fi28 % ((fi28) | 1))), F1: 2, F2: (^(v17))}
                    _ = v35
                }
                _ = fi28
                _ = v33
            }
        }
        switch ("a" + v5) {
        case "hello":
            v18 += v3
        }
    }
    if (true || true) {
        var v36 int16 = 32767
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
