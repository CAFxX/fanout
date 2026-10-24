package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint16, p1 int8, p2 int32) uint {
    v6 := []int{4294967295}
    var v7 bool = false
    if false {
        switch ("" + "a") {
        case "\n\t":
            var v8 string = "\n\t"
            _ = v8
        case "hello":
            var v9 uint8 = 255
            _ = v9
        }
    } else {
        switch 1 {
        case -2147483648:
            if true {
                switch ("a" + ("hello" + "a")) {
                case "foo bar":
                    v6 = append(v6, 127)
                }
            }
        default:
            v6 = append(v6, 9223372036854775807)
        }
        switch -2147483648 {
        case 984:
            fmt.Println(p2)
        case 537:
            if v7 {
                for fi10 := 0; fi10 < 2; fi10++ {
                    if v7 {
                        p2 = p2
                    }
                    for fi11 := 0; fi11 < 9; fi11++ {
                        p2 = -970
                        p0 = p0
                        p2 = (p2 >> ((255) & 63))
                        _ = fi11
                    }
                    if len(v6) > 0 {
                        v6[0] = (fi10 % (((fi10 | fi10)) | 1))
                    }
                    _ = fi10
                }
                fmt.Println(v7, p0)
                if false {
                    if len(v6) > 0 {
                        v6[0] = -9223372036854775808
                    }
                    var v12 int = 250
                    _ = v12
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint64, p1 int, p2 int32) int64 {
    v13 := []int{42}
    p1 = p1
    if true {
        for fi14 := 0; fi14 < 6; fi14++ {
            p2 = (p2 ^ (p2 * p2))
            v13 = append(v13, fi14)
            _ = fi14
        }
        p2 -= 1
        _ = len(v13)
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 1e10
    var v3 float32 = float32(0.1)
    v4 := make(map[int]bool)
    v5 := []uint32{42, 40}
    var v15 uint = 255
    v15 -= v15
    var v16 uint16 = 1
    switch 692 {
    case 9223372036854775807:
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS1{F0: 42}
            v16 = (v16 % ((630) | 1))
            _ = fi17
            _ = v18
        }
    case 127:
        var v19 int8 = -1
        _ = v19
    }
    delete(v1, 238)
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 float64 = v2
        _ = fi20
        _ = v21
    }
    v22 := v1[1]
    var v23 int16 = 1
    for fi24 := 0; fi24 < 10; fi24++ {
        switch (v22 & fi24) {
        case 32768:
            v15 += v15
        }
        v2 += 0.0
        _ = fi24
    }
    v5 = append(v5, 42)
    var v25 int = v22
    switch "a" {
    case "\n\t":
        var v26 uint64 = 0
        _ = v26
    case "foo bar":
        v25 -= (v22 + -95)
    case "a":
        var v27 int32 = 1
        _ = v27
    }
    var v28 int16 = (v23 % (((v23 % ((v23) | 1))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
